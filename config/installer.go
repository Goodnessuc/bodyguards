package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// SetupDependencies rewrite to install dependencies based on yaml and sets up permissions for architecture bash scripts
func SetupDependencies() error {

	// Install golangCI-lint
	msg, err := installGolangCILint()
	if err != nil {
		fmt.Println("Error: failed to install Golangci-lint:", err)
		return err
	}
	fmt.Println(msg)

	// Install Godoc
	msg, err = installGodoc()
	if err != nil {
		fmt.Println("Error: failed to install Godoc:", err)
		return err
	}
	fmt.Println(msg)

	err = CreateBodyguardsYAML()
	if err != nil {
		log.Fatalf("Failed to create bodyguards.yaml: %v", err)
	}

	fmt.Println("bodyguards.yaml created successfully!")

	fmt.Sprintln("Please wait while we install Go report card.")
	fmt.Sprintln("This may take a while...")

	// Install Goreportcard-cli
	msg, err = installGoReportCard()
	if err != nil {
		fmt.Println("Error: failed to install Goreportcard-cli:", err)
		return err
	}
	fmt.Println(msg)
	// All operations completed successfully
	return nil
}

func installGolangCILint() (string, error) {
	// Command 1: go install golang-ci-lint
	cmd := exec.Command("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1")
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to install Golangci-lint: %w", err)
	}
	return "GolangCI-lint installed successfully", nil
}

func installGodoc() (string, error) {
	// Install godoc command
	cmd := exec.Command("go", "install", "golang.org/x/tools/cmd/godoc@latest")
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to install Godoc: %w", err)
	}

	return "Godoc installed successfully", nil
}

func installGoReportCard() (string, error) {

	// Git clone
	cmd := exec.Command("git", "clone", "https://github.com/gojp/goreportcard.git")
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to clone repository: %w", err)
	}

	// Change directory
	err = os.Chdir("goreportcard")
	if err != nil {
		return "", fmt.Errorf("failed to change directory: %w", err)
	}

	// Make install
	cmd = exec.Command("make", "install")
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute 'make install': %w", err)
	}

	// Go install
	cmd = exec.Command("go", "install", "./cmd/goreportcard-cli")
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute 'go install': %w", err)
	}

	err = os.Chdir("..")
	if err != nil {
		fmt.Println("Error:", err)
	}
	cmd = exec.Command("rm", "-rf", "goreportcard")
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to delete goreportcard repository after installation  %w", err)
	}

	return "Goreportcard-cli successfully installed", nil
}
