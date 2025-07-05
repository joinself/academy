#!/bin/bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Default values
EXAMPLES_DIR="../examples/server"
SDK_PACKAGE="github.com/joinself/self-go-sdk"
NEW_VERSION=""
VERBOSE=false
DRY_RUN=false

# Arrays to track results
declare -a SUCCESSFUL_UPDATES=()
declare -a FAILED_BUILDS=()
declare -a CHANGES_MADE=()

# Function to print usage
usage() {
    cat << EOF
Usage: $0 -v <new_version> [options]

Update self-go-sdk version across all Go examples and test builds.

Options:
    -v <version>    New SDK version to update to (required)
    -d <directory>  Examples directory (default: examples/server)
    --verbose       Enable verbose output
    --dry-run       Show what would be changed without making changes
    -h, --help      Show this help message

Examples:
    $0 -v v0.60.0
    $0 -v v0.60.0 --verbose
    $0 -v v0.60.0 --dry-run
    $0 -v v0.60.0 -d examples/server --verbose

EOF
}

# Function to log messages
log() {
    local level=$1
    shift
    case $level in
        "INFO")
            echo -e "${BLUE}[INFO]${NC} $*"
            ;;
        "SUCCESS")
            echo -e "${GREEN}[SUCCESS]${NC} $*"
            ;;
        "WARNING")
            echo -e "${YELLOW}[WARNING]${NC} $*"
            ;;
        "ERROR")
            echo -e "${RED}[ERROR]${NC} $*"
            ;;
        "VERBOSE")
            if [[ "$VERBOSE" == "true" ]]; then
                echo -e "${BLUE}[VERBOSE]${NC} $*"
            fi
            ;;
    esac
}

# Function to find all go.mod files in examples directory
find_go_modules() {
    local base_dir=$1
    find "$base_dir" -name "go.mod" -type f | sort
}

# Function to get current SDK version from go.mod
get_current_sdk_version() {
    local go_mod_file=$1
    grep "$SDK_PACKAGE" "$go_mod_file" | awk '{print $2}' | head -1
}

# Function to update SDK version in go.mod
update_sdk_version() {
    local go_mod_file=$1
    local new_version=$2
    local current_version
    
    current_version=$(get_current_sdk_version "$go_mod_file")
    
    if [[ -z "$current_version" ]]; then
        log "WARNING" "No SDK dependency found in $go_mod_file"
        return 1
    fi
    
    if [[ "$current_version" == "$new_version" ]]; then
        log "VERBOSE" "SDK version already up to date in $go_mod_file ($current_version)"
        return 0
    fi
    
    log "VERBOSE" "Updating SDK version in $go_mod_file: $current_version -> $new_version"
    
    if [[ "$DRY_RUN" == "false" ]]; then
        # Update the version
        sed -i.tmp "s|$SDK_PACKAGE $current_version|$SDK_PACKAGE $new_version|g" "$go_mod_file"
        rm "$go_mod_file.tmp" 2>/dev/null || true
        
        # Record the change
        CHANGES_MADE+=("$go_mod_file: $current_version -> $new_version")
    else
        CHANGES_MADE+=("$go_mod_file: $current_version -> $new_version (dry-run)")
    fi
    
    return 0
}

# Function to build a Go module
build_go_module() {
    local module_dir=$1
    local build_output
    local build_exit_code
    local tidy_output
    
    log "VERBOSE" "Building module in $module_dir"
    
    if [[ "$DRY_RUN" == "true" ]]; then
        log "VERBOSE" "Skipping build (dry-run mode)"
        return 0
    fi
    
    cd "$module_dir"
    
    # Clean any existing builds
    go clean 2>/dev/null || true
    
    # Remove go.sum to force fresh download of dependencies
    if [[ -f "go.sum" ]]; then
        log "VERBOSE" "Removing go.sum to force fresh dependency resolution"
        rm -f go.sum
    fi
    
    # Update dependencies with verbose output for debugging
    log "VERBOSE" "Running go mod tidy to update dependencies"
    tidy_output=$(go mod tidy 2>&1)
    tidy_exit_code=$?
    
    if [[ $tidy_exit_code -ne 0 ]]; then
        log "ERROR" "go mod tidy failed in $module_dir"
        log "ERROR" "Tidy output: $tidy_output"
        cd - > /dev/null
        FAILED_BUILDS+=("$module_dir: go mod tidy failed - $tidy_output")
        return 1
    fi
    
    # Download dependencies explicitly
    log "VERBOSE" "Downloading dependencies"
    go mod download 2>/dev/null || true
    
    # Try to build
    log "VERBOSE" "Attempting to build"
    build_output=$(go build -o /tmp/test_build_$$ . 2>&1)
    build_exit_code=$?
    
    # Clean up temporary build file
    rm -f "/tmp/test_build_$$" 2>/dev/null || true
    
    cd - > /dev/null
    
    if [[ $build_exit_code -eq 0 ]]; then
        log "VERBOSE" "Build successful for $module_dir"
        return 0
    else
        log "ERROR" "Build failed for $module_dir"
        log "ERROR" "Build output: $build_output"
        FAILED_BUILDS+=("$module_dir: $build_output")
        return 1
    fi
}

# Function to show git status on failure
show_git_status() {
    log "INFO" "Build failures detected. Use 'git checkout .' to restore files or 'git diff' to see changes."
}

# Function to process a single example
process_example() {
    local go_mod_file=$1
    local module_dir
    
    module_dir=$(dirname "$go_mod_file")
    
    log "INFO" "Processing: $go_mod_file"
    
    # Update SDK version
    if update_sdk_version "$go_mod_file" "$NEW_VERSION"; then
        # Try to build
        if build_go_module "$module_dir"; then
            SUCCESSFUL_UPDATES+=("$go_mod_file")
            log "SUCCESS" "Successfully updated and built: $module_dir"
        else
            log "ERROR" "Build failed after update: $module_dir"
        fi
    else
        log "WARNING" "Failed to update SDK version: $go_mod_file"
    fi
}

# Function to generate final report
generate_report() {
    echo
    echo "================== UPDATE REPORT =================="
    echo
    
    if [[ ${#CHANGES_MADE[@]} -gt 0 ]]; then
        echo -e "${BLUE}Changes Made:${NC}"
        for change in "${CHANGES_MADE[@]}"; do
            echo "  - $change"
        done
        echo
    fi
    
    if [[ ${#SUCCESSFUL_UPDATES[@]} -gt 0 ]]; then
        echo -e "${GREEN}Successful Updates (${#SUCCESSFUL_UPDATES[@]}):${NC}"
        for success in "${SUCCESSFUL_UPDATES[@]}"; do
            echo "  ✓ $success"
        done
        echo
    fi
    
    if [[ ${#FAILED_BUILDS[@]} -gt 0 ]]; then
        echo -e "${RED}Failed Builds (${#FAILED_BUILDS[@]}):${NC}"
        for failure in "${FAILED_BUILDS[@]}"; do
            echo "  ✗ $failure"
        done
        echo
        
        echo -e "${RED}RESULT: FAILURE${NC}"
        echo "Some examples failed to build with the new SDK version."
        echo "Check the errors above for details."
        return 1
    else
        echo -e "${GREEN}RESULT: SUCCESS${NC}"
        echo "All examples successfully updated and built with the new SDK version."
        return 0
    fi
}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -v)
            NEW_VERSION="$2"
            shift 2
            ;;
        -d)
            EXAMPLES_DIR="$2"
            shift 2
            ;;
        --verbose)
            VERBOSE=true
            shift
            ;;
        --dry-run)
            DRY_RUN=true
            shift
            ;;
        -h|--help)
            usage
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            usage
            exit 1
            ;;
    esac
done

# Validate required arguments
if [[ -z "$NEW_VERSION" ]]; then
    echo "Error: New version is required"
    usage
    exit 1
fi

# Validate examples directory exists
if [[ ! -d "$EXAMPLES_DIR" ]]; then
    log "ERROR" "Examples directory does not exist: $EXAMPLES_DIR"
    exit 1
fi

# Main execution
main() {
    log "INFO" "Starting SDK version update process..."
    log "INFO" "Target version: $NEW_VERSION"
    log "INFO" "Examples directory: $EXAMPLES_DIR"
    log "INFO" "Dry run: $DRY_RUN"
    echo
    
    # Find all go.mod files
    local go_mod_files
    go_mod_files=$(find_go_modules "$EXAMPLES_DIR")
    
    if [[ -z "$go_mod_files" ]]; then
        log "ERROR" "No go.mod files found in $EXAMPLES_DIR"
        exit 1
    fi
    
    local total_modules
    total_modules=$(echo "$go_mod_files" | wc -l)
    log "INFO" "Found $total_modules Go modules to process"
    echo
    
    # Process each module
    while IFS= read -r go_mod_file; do
        process_example "$go_mod_file"
        echo
    done <<< "$go_mod_files"
    
    # Generate final report
    if generate_report; then
        exit 0
    else
        if [[ "$DRY_RUN" == "false" ]]; then
            show_git_status
        fi
        exit 1
    fi
}

# No cleanup needed - git handles version control

# Run main function
main 
