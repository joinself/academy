#!/bin/bash

set -e

echo "Building Go SDK Version Updater..."

# Build the Go program
go build -o update-sdk-version update_sdk_version.go

echo "✅ Built: update-sdk-version"
echo ""
echo "Usage examples:"
echo "  ./update-sdk-version -version v0.60.0"
echo "  ./update-sdk-version -version v0.60.0 -dry-run"
echo "  ./update-sdk-version -version v0.60.0 -verbose"
echo "  ./update-sdk-version -version v0.60.0 -json -output results.json" 
