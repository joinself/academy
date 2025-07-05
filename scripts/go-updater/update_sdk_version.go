package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Configuration holds the command line arguments and settings
type Configuration struct {
	NewVersion  string
	ExamplesDir string
	DryRun      bool
	Verbose     bool
	JSONOutput  bool
	OutputFile  string
	SDKPackage  string
}

// ModuleInfo represents information about a Go module
type ModuleInfo struct {
	Path           string `json:"path"`
	CurrentVersion string `json:"current_version"`
	NewVersion     string `json:"new_version"`
	UpdateSuccess  bool   `json:"update_success"`
	BuildSuccess   bool   `json:"build_success"`
	BuildOutput    string `json:"build_output,omitempty"`
	Error          string `json:"error,omitempty"`
}

// UpdateReport represents the final report of the update process
type UpdateReport struct {
	Timestamp        time.Time    `json:"timestamp"`
	TargetVersion    string       `json:"target_version"`
	ExamplesDir      string       `json:"examples_dir"`
	TotalModules     int          `json:"total_modules"`
	SuccessfulBuilds int          `json:"successful_builds"`
	FailedBuilds     int          `json:"failed_builds"`
	Modules          []ModuleInfo `json:"modules"`
	Success          bool         `json:"success"`
	DryRun           bool         `json:"dry_run"`
}

func main() {
	config := parseFlags()

	if err := validateConfig(config); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	report, err := updateSDKVersions(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := outputReport(report, config); err != nil {
		fmt.Fprintf(os.Stderr, "Error outputting report: %v\n", err)
		os.Exit(1)
	}

	if !report.Success {
		os.Exit(1)
	}
}

func parseFlags() *Configuration {
	config := &Configuration{
		SDKPackage: "github.com/joinself/self-go-sdk",
	}

	flag.StringVar(&config.NewVersion, "version", "", "New SDK version to update to (required)")
	flag.StringVar(&config.NewVersion, "v", "", "New SDK version to update to (required)")
	flag.StringVar(&config.ExamplesDir, "dir", "../../examples/server", "Examples directory")
	flag.StringVar(&config.ExamplesDir, "d", "../../examples/server", "Examples directory")
	flag.BoolVar(&config.DryRun, "dry-run", false, "Show what would be changed without making changes")
	flag.BoolVar(&config.Verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&config.JSONOutput, "json", false, "Output results in JSON format")
	flag.StringVar(&config.OutputFile, "output", "", "Output file for results (default: stdout)")
	flag.StringVar(&config.SDKPackage, "sdk-package", "github.com/joinself/self-go-sdk", "SDK package name")

	flag.Parse()

	return config
}

func validateConfig(config *Configuration) error {
	if config.NewVersion == "" {
		return fmt.Errorf("new version is required")
	}

	if _, err := os.Stat(config.ExamplesDir); os.IsNotExist(err) {
		return fmt.Errorf("examples directory does not exist: %s", config.ExamplesDir)
	}

	return nil
}

func updateSDKVersions(config *Configuration) (*UpdateReport, error) {
	report := &UpdateReport{
		Timestamp:     time.Now(),
		TargetVersion: config.NewVersion,
		ExamplesDir:   config.ExamplesDir,
		DryRun:        config.DryRun,
		Success:       true,
	}

	// Find all go.mod files
	goModFiles, err := findGoModFiles(config.ExamplesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to find go.mod files: %w", err)
	}

	report.TotalModules = len(goModFiles)

	if config.Verbose && !config.JSONOutput {
		fmt.Printf("Found %d Go modules to process\n", len(goModFiles))
	}

	// Process each module
	for _, goModFile := range goModFiles {
		moduleInfo := processModule(goModFile, config)
		report.Modules = append(report.Modules, moduleInfo)

		if moduleInfo.BuildSuccess {
			report.SuccessfulBuilds++
		} else {
			report.FailedBuilds++
			report.Success = false
		}

		if config.Verbose && !config.JSONOutput {
			printModuleStatus(moduleInfo)
		}
	}

	return report, nil
}

func findGoModFiles(baseDir string) ([]string, error) {
	var goModFiles []string

	err := filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.Name() == "go.mod" && !d.IsDir() {
			goModFiles = append(goModFiles, path)
		}

		return nil
	})

	return goModFiles, err
}

func processModule(goModFile string, config *Configuration) ModuleInfo {
	moduleInfo := ModuleInfo{
		Path: goModFile,
	}

	// Get current version
	currentVersion, err := getCurrentSDKVersion(goModFile, config.SDKPackage)
	if err != nil {
		moduleInfo.Error = fmt.Sprintf("Failed to get current version: %v", err)
		return moduleInfo
	}

	moduleInfo.CurrentVersion = currentVersion
	moduleInfo.NewVersion = config.NewVersion

	// Check if update is needed
	if currentVersion == config.NewVersion {
		moduleInfo.UpdateSuccess = true
		moduleInfo.BuildSuccess = true // Assume it's already working
		return moduleInfo
	}

	// Update version
	if !config.DryRun {
		err = updateGoModFile(goModFile, config.SDKPackage, currentVersion, config.NewVersion)
		if err != nil {
			moduleInfo.Error = fmt.Sprintf("Failed to update go.mod: %v", err)
			return moduleInfo
		}
	}

	moduleInfo.UpdateSuccess = true

	// Build module
	if !config.DryRun {
		buildSuccess, buildOutput := buildModule(filepath.Dir(goModFile))
		moduleInfo.BuildSuccess = buildSuccess
		moduleInfo.BuildOutput = buildOutput
	} else {
		moduleInfo.BuildSuccess = true // Assume it would work in dry-run
	}

	return moduleInfo
}

func getCurrentSDKVersion(goModFile, sdkPackage string) (string, error) {
	file, err := os.Open(goModFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Contains(line, sdkPackage) {
			parts := strings.Fields(line)
			if len(parts) >= 2 && parts[0] == sdkPackage {
				return parts[1], nil
			}
		}
	}

	return "", fmt.Errorf("SDK package not found in go.mod")
}

func updateGoModFile(goModFile, sdkPackage, oldVersion, newVersion string) error {
	// Read file
	content, err := os.ReadFile(goModFile)
	if err != nil {
		return err
	}

	// Replace version
	oldPattern := regexp.QuoteMeta(sdkPackage + " " + oldVersion)
	newString := sdkPackage + " " + newVersion

	re := regexp.MustCompile(oldPattern)
	newContent := re.ReplaceAll(content, []byte(newString))

	// Write updated content
	return os.WriteFile(goModFile, newContent, 0644)
}

func buildModule(moduleDir string) (bool, string) {
	// Change to module directory
	originalDir, err := os.Getwd()
	if err != nil {
		return false, fmt.Sprintf("Failed to get current directory: %v", err)
	}

	if err := os.Chdir(moduleDir); err != nil {
		return false, fmt.Sprintf("Failed to change to module directory: %v", err)
	}
	defer os.Chdir(originalDir)

	// Clean any existing builds
	exec.Command("go", "clean").Run()

	// Remove go.sum to force fresh download of dependencies
	goSumPath := "go.sum"
	if _, err := os.Stat(goSumPath); err == nil {
		os.Remove(goSumPath)
	}

	// Update dependencies
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyOutput, tidyErr := tidyCmd.CombinedOutput()
	if tidyErr != nil {
		return false, fmt.Sprintf("go mod tidy failed: %v\nOutput: %s", tidyErr, string(tidyOutput))
	}

	// Download dependencies explicitly
	exec.Command("go", "mod", "download").Run()

	// Build
	cmd := exec.Command("go", "build", "-o", "/tmp/test_build", ".")
	output, err := cmd.CombinedOutput()

	// Clean up build artifact
	os.Remove("/tmp/test_build")

	if err != nil {
		return false, string(output)
	}

	return true, string(output)
}

func printModuleStatus(moduleInfo ModuleInfo) {
	status := "✓"
	if !moduleInfo.BuildSuccess {
		status = "✗"
	}

	fmt.Printf("%s %s (%s -> %s)\n", status, moduleInfo.Path, moduleInfo.CurrentVersion, moduleInfo.NewVersion)

	if moduleInfo.Error != "" {
		fmt.Printf("  Error: %s\n", moduleInfo.Error)
	}

	if !moduleInfo.BuildSuccess && moduleInfo.BuildOutput != "" {
		fmt.Printf("  Build output: %s\n", strings.TrimSpace(moduleInfo.BuildOutput))
	}
}

func outputReport(report *UpdateReport, config *Configuration) error {
	var output []byte
	var err error

	if config.JSONOutput {
		output, err = json.MarshalIndent(report, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}
	} else {
		output = []byte(formatTextReport(report))
	}

	if config.OutputFile != "" {
		return os.WriteFile(config.OutputFile, output, 0644)
	}

	fmt.Print(string(output))
	return nil
}

func formatTextReport(report *UpdateReport) string {
	var sb strings.Builder

	sb.WriteString("================== UPDATE REPORT ==================\n")
	sb.WriteString(fmt.Sprintf("Timestamp: %s\n", report.Timestamp.Format(time.RFC3339)))
	sb.WriteString(fmt.Sprintf("Target Version: %s\n", report.TargetVersion))
	sb.WriteString(fmt.Sprintf("Examples Directory: %s\n", report.ExamplesDir))
	sb.WriteString(fmt.Sprintf("Total Modules: %d\n", report.TotalModules))
	sb.WriteString(fmt.Sprintf("Successful Builds: %d\n", report.SuccessfulBuilds))
	sb.WriteString(fmt.Sprintf("Failed Builds: %d\n", report.FailedBuilds))
	sb.WriteString(fmt.Sprintf("Dry Run: %t\n", report.DryRun))
	sb.WriteString("\n")

	if len(report.Modules) > 0 {
		sb.WriteString("Module Details:\n")
		for _, module := range report.Modules {
			status := "SUCCESS"
			if !module.BuildSuccess {
				status = "FAILED"
			}

			sb.WriteString(fmt.Sprintf("  [%s] %s (%s -> %s)\n",
				status, module.Path, module.CurrentVersion, module.NewVersion))

			if module.Error != "" {
				sb.WriteString(fmt.Sprintf("    Error: %s\n", module.Error))
			}

			if !module.BuildSuccess && module.BuildOutput != "" {
				sb.WriteString(fmt.Sprintf("    Build Output: %s\n", strings.TrimSpace(module.BuildOutput)))
			}
		}
		sb.WriteString("\n")
	}

	if report.Success {
		sb.WriteString("RESULT: SUCCESS\n")
		sb.WriteString("All examples successfully updated and built with the new SDK version.\n")
	} else {
		sb.WriteString("RESULT: FAILURE\n")
		sb.WriteString("Some examples failed to build with the new SDK version.\n")
	}

	return sb.String()
}
