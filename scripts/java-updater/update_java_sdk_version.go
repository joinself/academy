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

// ModuleInfo represents information about a Java/Kotlin module
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
		SDKPackage: "com.joinself:sdk-jvm",
	}

	flag.StringVar(&config.NewVersion, "version", "", "New SDK version to update to (required)")
	flag.StringVar(&config.NewVersion, "v", "", "New SDK version to update to (required)")
	flag.StringVar(&config.ExamplesDir, "dir", "../../examples/server", "Examples directory")
	flag.StringVar(&config.ExamplesDir, "d", "../../examples/server", "Examples directory")
	flag.BoolVar(&config.DryRun, "dry-run", false, "Show what would be changed without making changes")
	flag.BoolVar(&config.Verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&config.JSONOutput, "json", false, "Output results in JSON format")
	flag.StringVar(&config.OutputFile, "output", "", "Output file for results (default: stdout)")
	flag.StringVar(&config.SDKPackage, "sdk-package", "com.joinself:sdk-jvm", "SDK package name")

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

	// Find all build.gradle.kts files in java subdirectories
	gradleFiles, err := findGradleFiles(config.ExamplesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to find build.gradle.kts files: %w", err)
	}

	report.TotalModules = len(gradleFiles)

	if config.Verbose && !config.JSONOutput {
		fmt.Printf("Found %d Java modules to process\n", len(gradleFiles))
	}

	// Process each module
	for _, gradleFile := range gradleFiles {
		moduleInfo := processModule(gradleFile, config)
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

func findGradleFiles(baseDir string) ([]string, error) {
	var gradleFiles []string

	err := filepath.WalkDir(baseDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Look for build.gradle.kts files in java subdirectories
		if d.Name() == "build.gradle.kts" && !d.IsDir() {
			// Check if this is in a java subdirectory
			if strings.Contains(path, "/java/") {
				gradleFiles = append(gradleFiles, path)
			}
		}

		return nil
	})

	return gradleFiles, err
}

func processModule(gradleFile string, config *Configuration) ModuleInfo {
	moduleInfo := ModuleInfo{
		Path: gradleFile,
	}

	// Get current version
	currentVersion, err := getCurrentSDKVersion(gradleFile, config.SDKPackage)
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
		err = updateGradleFile(gradleFile, config.SDKPackage, currentVersion, config.NewVersion)
		if err != nil {
			moduleInfo.Error = fmt.Sprintf("Failed to update build.gradle.kts: %v", err)
			return moduleInfo
		}
	}

	moduleInfo.UpdateSuccess = true

	// Build module
	if !config.DryRun {
		buildSuccess, buildOutput := buildModule(filepath.Dir(gradleFile))
		moduleInfo.BuildSuccess = buildSuccess
		moduleInfo.BuildOutput = buildOutput
	} else {
		moduleInfo.BuildSuccess = true // Assume it would work in dry-run
	}

	return moduleInfo
}

func getCurrentSDKVersion(gradleFile, sdkPackage string) (string, error) {
	file, err := os.Open(gradleFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Pattern to match: implementation("com.joinself:sdk-jvm:1.0.0")
	pattern := regexp.MustCompile(`implementation\(\s*"` + regexp.QuoteMeta(sdkPackage) + `:([^"]+)"\s*\)`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.FindStringSubmatch(line)
		if len(matches) > 1 {
			return matches[1], nil
		}
	}

	return "", fmt.Errorf("SDK package not found in build.gradle.kts")
}

func updateGradleFile(gradleFile, sdkPackage, oldVersion, newVersion string) error {
	// Read file
	content, err := os.ReadFile(gradleFile)
	if err != nil {
		return err
	}

	// Pattern to match and replace: implementation("com.joinself:sdk-jvm:oldVersion")
	oldPattern := `implementation\(\s*"` + regexp.QuoteMeta(sdkPackage) + `:` + regexp.QuoteMeta(oldVersion) + `"\s*\)`
	newString := `implementation("` + sdkPackage + `:` + newVersion + `")`

	re := regexp.MustCompile(oldPattern)
	newContent := re.ReplaceAll(content, []byte(newString))

	// Write updated content
	return os.WriteFile(gradleFile, newContent, 0644)
}

func buildModule(moduleDir string) (bool, string) {
	// Change to module directory
	cmd := exec.Command("gradle", "build")
	cmd.Dir = moduleDir

	// Capture both stdout and stderr
	output, err := cmd.CombinedOutput()

	if err != nil {
		return false, string(output)
	}

	return true, string(output)
}

func printModuleStatus(moduleInfo ModuleInfo) {
	status := "SUCCESS"
	if moduleInfo.Error != "" || !moduleInfo.BuildSuccess {
		status = "FAILED"
	}

	fmt.Printf("  [%s] %s (%s -> %s)\n",
		status,
		moduleInfo.Path,
		moduleInfo.CurrentVersion,
		moduleInfo.NewVersion)

	if moduleInfo.Error != "" {
		fmt.Printf("    Error: %s\n", moduleInfo.Error)
	}

	if !moduleInfo.BuildSuccess && moduleInfo.BuildOutput != "" {
		// Show only the first few lines of build output to avoid spam
		lines := strings.Split(strings.TrimSpace(moduleInfo.BuildOutput), "\n")
		maxLines := 3
		if len(lines) > maxLines {
			for _, line := range lines[:maxLines] {
				fmt.Printf("    Build Output: %s\n", line)
			}
			fmt.Printf("    ... (truncated)\n")
		} else {
			for _, line := range lines {
				fmt.Printf("    Build Output: %s\n", line)
			}
		}
	}
}

func outputReport(report *UpdateReport, config *Configuration) error {
	var output string
	var err error

	if config.JSONOutput {
		jsonData, jsonErr := json.MarshalIndent(report, "", "  ")
		if jsonErr != nil {
			return fmt.Errorf("failed to marshal JSON: %w", jsonErr)
		}
		output = string(jsonData)
	} else {
		output = formatTextReport(report)
	}

	if config.OutputFile != "" {
		err = os.WriteFile(config.OutputFile, []byte(output), 0644)
	} else {
		_, err = fmt.Print(output)
	}

	return err
}

func formatTextReport(report *UpdateReport) string {
	var result strings.Builder

	result.WriteString("================== UPDATE REPORT ==================\n")
	result.WriteString(fmt.Sprintf("Timestamp: %s\n", report.Timestamp.Format(time.RFC3339)))
	result.WriteString(fmt.Sprintf("Target Version: %s\n", report.TargetVersion))
	result.WriteString(fmt.Sprintf("Examples Directory: %s\n", report.ExamplesDir))
	result.WriteString(fmt.Sprintf("Total Modules: %d\n", report.TotalModules))
	result.WriteString(fmt.Sprintf("Successful Builds: %d\n", report.SuccessfulBuilds))
	result.WriteString(fmt.Sprintf("Failed Builds: %d\n", report.FailedBuilds))
	result.WriteString(fmt.Sprintf("Dry Run: %t\n", report.DryRun))
	result.WriteString("\nModule Details:\n")

	for _, module := range report.Modules {
		status := "SUCCESS"
		if module.Error != "" || !module.BuildSuccess {
			status = "FAILED"
		}

		result.WriteString(fmt.Sprintf("  [%s] %s (%s -> %s)\n",
			status,
			module.Path,
			module.CurrentVersion,
			module.NewVersion))

		if module.Error != "" {
			result.WriteString(fmt.Sprintf("    Error: %s\n", module.Error))
		}

		if !module.BuildSuccess && module.BuildOutput != "" {
			// Show first few lines of build output
			lines := strings.Split(strings.TrimSpace(module.BuildOutput), "\n")
			maxLines := 3
			if len(lines) > maxLines {
				for _, line := range lines[:maxLines] {
					result.WriteString(fmt.Sprintf("    Build Output: %s\n", line))
				}
				result.WriteString("    ... (truncated)\n")
			} else {
				for _, line := range lines {
					result.WriteString(fmt.Sprintf("    Build Output: %s\n", line))
				}
			}
		}
	}

	result.WriteString("\n")

	if report.Success {
		result.WriteString("RESULT: SUCCESS\n")
		result.WriteString("All examples successfully updated and built with the new SDK version.\n")
	} else {
		result.WriteString("RESULT: FAILURE\n")
		result.WriteString("Some examples failed to build with the new SDK version.\n")
	}

	return result.String()
}
