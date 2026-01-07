package main

import (
	"archive/zip"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const version = "1.0.0"

type Config struct {
	sourceDir   string
	outputPath  string
	packageAll  bool
	ignoreRules []string
}

func main() {
	var packageAll bool
	var showVersion bool
	
	flag.BoolVar(&packageAll, "a", false, "Package everything, ignore no files")
	flag.BoolVar(&showVersion, "v", false, "Show version")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.Parse()

	if showVersion {
		fmt.Printf("Corazon v%s - Browser Extension Packager\n", version)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) < 2 {
		printUsage()
		os.Exit(1)
	}

	config := Config{
		sourceDir:  args[0],
		outputPath: args[1],
		packageAll: packageAll,
	}

	if err := run(config); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Corazon - Browser Extension Packager (c)2026 (github.com/abraham-ny, https://abrahamonline.netlify.app) Abraham Moruri")
	fmt.Println("\nUsage:")
	fmt.Println("  crx <source_dir> <output_file>")
	fmt.Println("  crx -a <source_dir> <output_file>")
	fmt.Println("\nOptions:")
	fmt.Println("  -a          Package all files, ignore .gitignore and .zignore")
	fmt.Println("  -v, -version Show version")
	fmt.Println("\nExamples:")
	fmt.Println("  crx ./my-extension extension.zip")
	fmt.Println("  crx ./my-extension extension.crx")
	fmt.Println("  crx -a ./my-extension extension-full.zip")
}

func run(config Config) error {
	// Validate source directory
	info, err := os.Stat(config.sourceDir)
	if err != nil {
		return fmt.Errorf("source directory error: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("source path is not a directory: %s", config.sourceDir)
	}

	// Check for manifest.json
	manifestPath := filepath.Join(config.sourceDir, "manifest.json")
	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		return fmt.Errorf("manifest.json not found in source directory")
	}

	// Load ignore rules if not packaging all
	if !config.packageAll {
		config.ignoreRules = loadIgnoreRules(config.sourceDir)
	}

	// Create output file
	outFile, err := os.Create(config.outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	// Create zip writer
	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	// Walk directory and add files
	fileCount := 0
	err = filepath.Walk(config.sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get relative path
		relPath, err := filepath.Rel(config.sourceDir, path)
		if err != nil {
			return err
		}

		// Skip root directory
		if relPath == "." {
			return nil
		}

		// Check if should ignore
		if !config.packageAll && shouldIgnore(relPath, info.IsDir(), config.ignoreRules) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip directories (they're created automatically with files)
		if info.IsDir() {
			return nil
		}

		// Add file to zip
		if err := addFileToZip(zipWriter, path, relPath); err != nil {
			return fmt.Errorf("failed to add %s: %w", relPath, err)
		}

		fileCount++
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to package extension: %w", err)
	}

	fmt.Printf("✓ Successfully packaged %d files to %s\n", fileCount, config.outputPath)
	return nil
}

func loadIgnoreRules(sourceDir string) []string {
	var rules []string

	// Load .gitignore
	gitignorePath := filepath.Join(sourceDir, ".gitignore")
	if gitRules := readIgnoreFile(gitignorePath); len(gitRules) > 0 {
		rules = append(rules, gitRules...)
	}

	// Load .zignore (takes precedence)
	zignorePath := filepath.Join(sourceDir, ".zignore")
	if zRules := readIgnoreFile(zignorePath); len(zRules) > 0 {
		rules = append(rules, zRules...)
	}

	// Add common sensitive files by default
	defaultIgnores := []string{
		".git/",
		".gitignore",
		".zignore",
		"node_modules/",
		".DS_Store",
		"Thumbs.db",
		"*.log",
		".env",
		".env.local",
		"*.pem",
		"*.key",
	}
	rules = append(rules, defaultIgnores...)

	return rules
}

func readIgnoreFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var rules []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		rules = append(rules, line)
	}

	return rules
}

func shouldIgnore(path string, isDir bool, rules []string) bool {
	// Normalize path separators
	path = filepath.ToSlash(path)

	for _, rule := range rules {
		if rule == "" {
			continue
		}

		// Handle directory-specific rules (ending with /)
		if strings.HasSuffix(rule, "/") {
			if isDir && matchPattern(path+"/", rule) {
				return true
			}
			continue
		}

		// Handle negation rules (starting with !)
		if strings.HasPrefix(rule, "!") {
			negatedRule := strings.TrimPrefix(rule, "!")
			if matchPattern(path, negatedRule) {
				return false
			}
			continue
		}

		// Check if path matches rule
		if matchPattern(path, rule) {
			return true
		}

		// Check if any parent directory matches
		if isDir && matchPattern(path+"/", rule) {
			return true
		}
	}

	return false
}

func matchPattern(path, pattern string) bool {
	// Handle wildcard patterns
	if strings.Contains(pattern, "*") {
		matched, _ := filepath.Match(pattern, filepath.Base(path))
		if matched {
			return true
		}
		// Try matching full path
		matched, _ = filepath.Match(pattern, path)
		return matched
	}

	// Exact match or prefix match
	pattern = strings.TrimSuffix(pattern, "/")
	path = strings.TrimSuffix(path, "/")

	if path == pattern {
		return true
	}

	// Check if path starts with pattern (for directory matching)
	if strings.HasPrefix(path, pattern+"/") {
		return true
	}

	return false
}

func addFileToZip(zipWriter *zip.Writer, filePath, zipPath string) error {
	// Open source file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file info
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Create header
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Use forward slashes in zip (required by zip spec)
	header.Name = filepath.ToSlash(zipPath)
	header.Method = zip.Deflate

	// Create writer
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	// Copy file content
	_, err = io.Copy(writer, file)
	return err
}