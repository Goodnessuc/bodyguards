package architecture

import (
	"fmt"
	"os"
	"path/filepath"
)

func createHexagonal() {
	// create top level directories
	directories := []string{"cmd", "internal", "pkg", "vendor"}
	for _, dir := range directories {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Printf("Error creating directory %s: %s\n", dir, err)
		}
	}

	// create cmd/http directory
	err := os.MkdirAll(filepath.Join("cmd", "http"), 0755)
	if err != nil {
		fmt.Printf("Error creating directory %s: %s\n", filepath.Join("cmd", "http"), err)
	}

	// create internal directories
	internalDirectories := []string{
		"adapters", "app", filepath.Join("app", "domain"),
		filepath.Join("app", "ports"), filepath.Join("app", "ports", "input"),
		filepath.Join("app", "ports", "output"), filepath.Join("app", "usecases"),
	}
	for _, dir := range internalDirectories {
		err := os.MkdirAll(filepath.Join("internal", dir), 0755)
		if err != nil {
			fmt.Printf("Error creating directory %s: %s\n", filepath.Join("internal", dir), err)
		}
	}

	// create internal/adapters directories
	adaptersDirectories := []string{"api", "database"}
	for _, dir := range adaptersDirectories {
		err := os.MkdirAll(filepath.Join("internal", "adapters", dir), 0755)
		if err != nil {
			fmt.Printf("Error creating directory %s: %s\n", filepath.Join("internal", "adapters", dir), err)
		}
	}

	// create vendor directories
	vendorDirectories := []string{"module1", "module2"}
	for _, dir := range vendorDirectories {
		err := os.MkdirAll(filepath.Join("vendor", dir), 0755)
		if err != nil {
			fmt.Printf("Error creating directory %s: %s\n", filepath.Join("vendor", dir), err)
		}
	}

	fmt.Println("Directory structure created successfully.")
}

func createMicroserviceStructure() {

	// Create directories
	directories := []string{
		"cmd", "configs", "docs", "internal", "pkg",
		"local", "mocks", "proto", "pb", "test",
		".git",
	}

	for _, dir := range directories {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Printf("Error creating directory %s: %s\n", dir, err)
			return
		}
	}

	// Create internal/package1 and internal/package2 directories
	internalDirs := []string{"package1", "package2"}

	for _, dir := range internalDirs {
		err := os.MkdirAll(filepath.Join("internal", dir), 0755)
		if err != nil {
			fmt.Printf("Error creating directory internal/%s: %s\n", dir, err)
			return
		}
	}

	// Print success message
	fmt.Println("Directory structure created successfully!")
}

func createMonolithicDirectory() {
	// Create the directory structure
	directories := []string{
		"cmd/app",
		"cmd/admin",
		"internal/api/handlers",
		"internal/api/middleware",
		"internal/api/models",
		"internal/api/repositories",
		"internal/api/services",
		"internal/config",
		"internal/database/migrations",
		"internal/logging",
		"internal/utils",
		"web/static/css",
		"web/static/js",
		"web/templates",
		"config",
		"migrations",
		"scripts",
	}

	for _, dir := range directories {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Printf("Error creating directory %s: %s\n", dir, err)
		}
	}

	fmt.Println("Directory tree created successfully.")
}

func createMVCDirectories() {
	// Create the main directory
	err := os.Mkdir("myapp", 0755)
	if err != nil {
		return
	}

	// Create the app directory
	err = os.MkdirAll("myapp/app", 0755)
	if err != nil {
		return
	}

	// Create subdirectories inside the app directory
	subdirectories := []string{
		"myapp/app/controllers",
		"myapp/app/models",
		"myapp/app/views",
		"myapp/app/middlewares",
		"myapp/app/routers",
		"myapp/app/services",
		"myapp/app/utils",
	}
	for _, dir := range subdirectories {
		os.MkdirAll(dir, 0755)
	}

	// Create subdirectories inside the views directory
	viewSubdirectories := []string{
		"myapp/app/views/home",
		"myapp/app/views/user",
	}
	for _, dir := range viewSubdirectories {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return
		}
	}

	// Create subdirectories inside the static directory
	staticSubdirectories := []string{
		"myapp/static/css",
		"myapp/static/js",
		"myapp/static/img",
	}
	for _, dir := range staticSubdirectories {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return
		}
	}

	// Create subdirectories inside the templates directory
	err = os.MkdirAll("myapp/templates", 0755)
	if err != nil {
		return
	}
}
