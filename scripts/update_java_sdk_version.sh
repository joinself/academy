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
SDK_PACKAGE="com.joinself:sdk-jvm"
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

Update com.joinself:sdk-jvm version across all Java examples and test builds.

Options:
    -v <version>    New SDK version to update to (required)
    -d <directory>  Examples directory (default: ../examples/server)
    --verbose       Enable verbose output
    --dry-run       Show what would be changed without making changes
    -h, --help      Show this help message

Examples:
    $0 -v 1.0.1
    $0 -v 1.0.1 --verbose
    $0 -v 1.0.1 --dry-run
    $0 -v 1.0.1 -d examples/server --verbose

Note:
    This script looks for build.gradle.kts files in java/ subdirectories
    and updates the implementation("com.joinself:sdk-jvm:VERSION") dependency.

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

# Function to find all build.gradle.kts files in java subdirectories
find_gradle_files() {
    local base_dir=$1
    find "$base_dir" -path "*/java/build.gradle.kts" -type f | sort
}

# Function to get current SDK version from build.gradle.kts
get_current_sdk_version() {
    local gradle_file=$1
    # Look for implementation("com.joinself:sdk-jvm:VERSION")
    grep -E 'implementation\(\s*"com\.joinself:sdk-jvm:' "$gradle_file" | \
        sed -E 's/.*implementation\(\s*"com\.joinself:sdk-jvm:([^"]+)".*/\1/' | \
        head -1
}

# Function to update SDK version in build.gradle.kts
update_sdk_version() {
    local gradle_file=$1
    local new_version=$2
    local current_version
    
    current_version=$(get_current_sdk_version "$gradle_file")
    
    if [[ -z "$current_version" ]]; then
        log "WARNING" "No SDK dependency found in $gradle_file"
        return 1
    fi
    
    if [[ "$current_version" == "$new_version" ]]; then
        log "VERBOSE" "SDK version already up to date in $gradle_file ($current_version)"
        return 0
    fi
    
    log "VERBOSE" "Updating SDK version in $gradle_file: $current_version -> $new_version"
    
    if [[ "$DRY_RUN" == "false" ]]; then
        # Update the version using sed
        sed -i.tmp "s|implementation(\"com\.joinself:sdk-jvm:$current_version\")|implementation(\"com.joinself:sdk-jvm:$new_version\")|g" "$gradle_file"
        rm "$gradle_file.tmp" 2>/dev/null || true
        
        # Record the change
        CHANGES_MADE+=("$gradle_file: $current_version -> $new_version")
    else
        CHANGES_MADE+=("$gradle_file: $current_version -> $new_version (dry-run)")
    fi
    
    return 0
}

# Function to build a Java module using Gradle
build_gradle_module() {
    local module_dir=$1
    local build_output
    local build_exit_code
    
    log "VERBOSE" "Building module in $module_dir"
    
    if [[ "$DRY_RUN" == "true" ]]; then
        log "VERBOSE" "Skipping build (dry-run mode)"
        return 0
    fi
    
    cd "$module_dir"
    
    # Clean any existing builds
    if [[ -f gradlew ]]; then
        ./gradlew clean 2>/dev/null || true
    else
        gradle clean 2>/dev/null || true
    fi
    
    # Remove build directory to force fresh build
    if [[ -d "build" ]]; then
        log "VERBOSE" "Removing build directory to force fresh compilation"
        rm -rf build
    fi
    
    # Remove .gradle directory to force fresh dependency resolution
    if [[ -d ".gradle" ]]; then
        log "VERBOSE" "Removing .gradle directory to force fresh dependency resolution"
        rm -rf .gradle
    fi
    
    # Try to build
    log "VERBOSE" "Attempting to build with Gradle"
    if [[ -f gradlew ]]; then
        build_output=$(./gradlew build 2>&1)
    else
        build_output=$(gradle build 2>&1)
    fi
    build_exit_code=$?
    
    cd - > /dev/null
    
    if [[ $build_exit_code -eq 0 ]]; then
        log "VERBOSE" "Build successful for $module_dir"
        return 0
    else
        log "ERROR" "Build failed for $module_dir"
        # Show first few lines of error to avoid overwhelming output
        echo "$build_output" | head -10 | while read -r line; do
            log "ERROR" "Build output: $line"
        done
        if [[ $(echo "$build_output" | wc -l) -gt 10 ]]; then
            log "ERROR" "... (output truncated)"
        fi
        FAILED_BUILDS+=("$module_dir: Build failed")
        return 1
    fi
}

# Function to show git status on failure
show_git_status() {
    log "INFO" "Build failures detected. Use 'git checkout .' to restore files or 'git diff' to see changes."
}

# Function to process a single example
process_example() {
    local gradle_file=$1
    local module_dir
    
    module_dir=$(dirname "$gradle_file")
    
    log "INFO" "Processing: $gradle_file"
    
    # Update SDK version
    if update_sdk_version "$gradle_file" "$NEW_VERSION"; then
        # Try to build
        if build_gradle_module "$module_dir"; then
            SUCCESSFUL_UPDATES+=("$gradle_file")
            log "SUCCESS" "Successfully updated and built: $module_dir"
        else
            log "ERROR" "Build failed after update: $module_dir"
        fi
    else
        log "WARNING" "Failed to update SDK version: $gradle_file"
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

# Validate required parameters
if [[ -z "$NEW_VERSION" ]]; then
    log "ERROR" "New version is required. Use -v <version>."
    usage
    exit 1
fi

if [[ ! -d "$EXAMPLES_DIR" ]]; then
    log "ERROR" "Examples directory does not exist: $EXAMPLES_DIR"
    exit 1
fi

# Show configuration
log "INFO" "Java SDK Version Updater"
log "INFO" "========================"
log "INFO" "Target version: $NEW_VERSION"
log "INFO" "Examples directory: $EXAMPLES_DIR"
log "INFO" "SDK package: $SDK_PACKAGE"
if [[ "$DRY_RUN" == "true" ]]; then
    log "INFO" "Mode: DRY RUN (no changes will be made)"
else
    log "INFO" "Mode: UPDATE (files will be modified)"
fi
if [[ "$VERBOSE" == "true" ]]; then
    log "INFO" "Verbose output: enabled"
fi
echo

# Find all Java modules
log "INFO" "Searching for Java modules..."
gradle_files=($(find_gradle_files "$EXAMPLES_DIR"))

if [[ ${#gradle_files[@]} -eq 0 ]]; then
    log "WARNING" "No build.gradle.kts files found in java/ subdirectories"
    exit 1
fi

log "INFO" "Found ${#gradle_files[@]} Java modules to process"
echo

# Process each module
for gradle_file in "${gradle_files[@]}"; do
    process_example "$gradle_file"
    echo
done

# Generate final report
if generate_report; then
    exit 0
else
    show_git_status
    exit 1
fi 
