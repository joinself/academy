# SDK Version Updaters

Automatically update SDK versions across all examples and test that they build successfully. 

## Overview

| Tool | Updates | Target Files | Dependency Format |
|------|---------|--------------|-------------------|
| **Go Updater** | `github.com/joinself/self-go-sdk` | `go.mod` | `github.com/joinself/self-go-sdk v1.0.0` |
| **Java Updater** | `com.joinself:sdk-jvm` | `build.gradle.kts` in `java/` subdirs | `implementation("com.joinself:sdk-jvm:1.0.0")` |

Both tools provide:
- **Automatic Discovery** - Finds all relevant files
- **Version Updates** - Updates SDK dependencies  
- **Build Testing** - Tests compilation with new versions
- **Detailed Reports** - Success/failure reporting
- **Dry Run Mode** - Preview changes safely
- **CI/CD Ready** - JSON output for automation

## Quick Start

### Interactive Shell Scripts (Recommended for Local Use)

```bash
# Update Go examples
./update_sdk_version.sh -v v0.60.0

# Update Java examples  
./update_java_sdk_version.sh -v 1.0.1

# Dry run (preview changes)
./update_sdk_version.sh -v v0.60.0 --dry-run
./update_java_sdk_version.sh -v 1.0.1 --dry-run

# Verbose output
./update_sdk_version.sh -v v0.60.0 --verbose
./update_java_sdk_version.sh -v 1.0.1 --verbose
```

### Go Programs (Recommended for CI/CD)

```bash
# Build both tools
cd go-updater && ./build.sh && cd ..
cd java-updater && ./build.sh && cd ..

# Update with JSON output
./go-updater/update-sdk-version -version v0.60.0 -json -output go-results.json
./java-updater/update-java-sdk-version -version 1.0.1 -json -output java-results.json

# Dry run
./go-updater/update-sdk-version -version v0.60.0 -dry-run
./java-updater/update-java-sdk-version -version 1.0.1 -dry-run
```

## Command Reference

### Shell Scripts

```bash
# Go updater
./update_sdk_version.sh -v <version> [--dry-run] [--verbose]

# Java updater  
./update_java_sdk_version.sh -v <version> [--dry-run] [--verbose]
```

**Options:**
- `-v <version>` - SDK version to update to (required)
- `-d <directory>` - Examples directory (default: `../examples/server`)
- `--dry-run` - Preview changes without making them
- `--verbose` - Show detailed progress
- `-h, --help` - Show help

### Go Programs

```bash
# Go updater
./go-updater/update-sdk-version [options]

# Java updater
./java-updater/update-java-sdk-version [options]
```

**Options:**
- `-version, -v <version>` - SDK version to update to (required)
- `-dir, -d <directory>` - Examples directory (default: `../../examples/server`)
- `-dry-run` - Preview changes without making them
- `-verbose` - Show detailed progress  
- `-json` - Output results in JSON format
- `-output <file>` - Save results to file (default: stdout)

## What Gets Updated

### Go Examples
```
examples/server/
├── 00_setup/01_new_account/go/go.mod     # github.com/joinself/self-go-sdk v0.59.0 → v0.60.0
├── 01_connection/01_direct/go/go.mod     # github.com/joinself/self-go-sdk v0.59.0 → v0.60.0
└── ... (all go.mod files)
```

### Java Examples  
```
examples/server/
├── 00_setup/01_new_account/java/build.gradle.kts    # implementation("com.joinself:sdk-jvm:1.0.0") → 1.0.1
├── 01_connection/01_direct/java/build.gradle.kts    # implementation("com.joinself:sdk-jvm:1.0.0") → 1.0.1
└── ... (all build.gradle.kts in java/ subdirectories)
```

## CI/CD Integration

### GitHub Actions

```yaml
name: Update SDK Versions

on:
  workflow_dispatch:
    inputs:
      go_sdk_version:
        description: 'Go SDK version (e.g., v0.60.0)'
        required: true
      java_sdk_version:
        description: 'Java SDK version (e.g., 1.0.1)'
        required: true

jobs:
  update-sdks:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Set up JDK
        uses: actions/setup-java@v3
        with:
          java-version: '17'
          distribution: 'temurin'
      
      - name: Build updaters
        run: |
          cd scripts/go-updater && ./build.sh && cd ..
          cd java-updater && ./build.sh && cd ..
      
      - name: Update Go SDK
        run: |
          cd scripts
          ./go-updater/update-sdk-version \
            -version ${{ inputs.go_sdk_version }} \
            -json -output go-results.json
      
      - name: Update Java SDK
        run: |
          cd scripts  
          ./java-updater/update-java-sdk-version \
            -version ${{ inputs.java_sdk_version }} \
            -json -output java-results.json
      
      - name: Upload results
        uses: actions/upload-artifact@v3
        if: always()
        with:
          name: sdk-update-results
          path: scripts/*-results.json
```

### Example JSON Output

```json
{
  "timestamp": "2024-01-15T10:30:00Z",
  "target_version": "v0.60.0",
  "total_modules": 15,
  "successful_builds": 14,
  "failed_builds": 1,
  "modules": [
    {
      "path": "examples/server/00_setup/01_new_account/go/go.mod",
      "current_version": "v0.59.0",
      "new_version": "v0.60.0", 
      "update_success": true,
      "build_success": true
    }
  ],
  "success": false,
  "dry_run": false
}
```

## Troubleshooting

### Common Issues

| Issue | Go Solution | Java Solution |
|-------|-------------|---------------|
| **Build failures** | Check Go version compatibility | Check Gradle/Java version |
| **Version not found** | Verify go.mod contains SDK dependency | Verify build.gradle.kts contains implementation line |
| **Permission denied** | `chmod +x *.sh` | `chmod +x *.sh` |
| **Network issues** | Check Go proxy settings | Check Gradle repository access |

### Recovery

Both tools integrate with git for easy recovery:

```bash
# See what changed
git diff

# Restore all changes
git checkout .

# Restore specific files
git checkout -- path/to/specific/file
```

### Debug Tips

1. **Always start with dry-run:**
   ```bash
   ./update_sdk_version.sh -v v0.60.0 --dry-run
   ./update_java_sdk_version.sh -v 1.0.1 --dry-run
   ```

2. **Use verbose mode for details:**
   ```bash
   ./update_sdk_version.sh -v v0.60.0 --verbose
   ./update_java_sdk_version.sh -v 1.0.1 --verbose
   ```

3. **Check JSON output for automation:**
   ```bash
   ./go-updater/update-sdk-version -version v0.60.0 -json | jq .
   ./java-updater/update-java-sdk-version -version 1.0.1 -json | jq .
   ```

## Requirements

### Base Requirements
- Git (for change management)
- Access to `examples/server` directory

### Go Updater
- Go 1.21+ installed

### Java Updater  
- Go 1.21+ installed (to build the updater tool)
- Gradle installed and in PATH
- Java 17+ installed

## Directory Structure

```
scripts/
├── update_sdk_version.sh           # Go updater shell script
├── update_java_sdk_version.sh      # Java updater shell script
├── go-updater/                     # Go updater Go program
│   ├── build.sh                    # Build script
│   ├── go.mod                      # Go module
│   ├── update_sdk_version.go       # Source code
│   └── update-sdk-version          # Built binary
├── java-updater/                   # Java updater Go program
│   ├── build.sh                    # Build script
│   ├── go.mod                      # Go module
│   ├── update_java_sdk_version.go  # Source code
│   └── update-java-sdk-version     # Built binary
└── README.md                       # This file
```

---

**Ready to keep your SDK examples up to date?**

Choose the tool for your language and run it! Both shell scripts and Go programs provide the same functionality - shell scripts are great for interactive use, Go programs are perfect for CI/CD automation.
