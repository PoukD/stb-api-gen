package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateProject(projectName string) error {
	root := projectName

	// ---------- DIRECTORIES ----------
	dirs := []string{
		".vscode",

		"cmd/app",
		"cmd/middleware/health",
		"cmd/middleware/log",
		"cmd/middleware/request",
		"cmd/middleware/verify",

		"cmd/routes/external/controller",
		"cmd/routes/internal",

		"config/database",
		"config/http/model",

		"logs/app",
		"logs/http",
		"logs/request",

		"util/bcrypt",
		"util/jwt",
	}

	// create all directories first
	for _, d := range dirs {
		fullPath := filepath.Join(root, d)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			return fmt.Errorf("failed to create dir %s: %w", fullPath, err)
		}
	}

	// ---------- FILES ----------
	files := map[string]string{
		"go.mod": fmt.Sprintf(`module %s

go 1.21
`, projectName),

		"main.go": `package main

func main() {
	// TODO: bootstrap application
}
`,

		"cmd/app/app.go":                               "package app\n",
		"cmd/middleware/health/health.go":              "package health\n",
		"cmd/middleware/log/log.go":                    "package log\n",
		"cmd/middleware/request/request.go":            "package request\n",
		"cmd/middleware/verify/verify.go":              "package verify\n",
		"cmd/routes/external/controller/controller.go": "package controller\n",
		"cmd/routes/internal/internal.go":              "package internal\n",

		"config/database/database.go": "package database\n",
		"config/http/model/http.go":   "package model\n",

		"util/bcrypt/bcrypt.go": "package bcrypt\n",
		"util/jwt/jwt.go":       "package jwt\n",
	}

	// create files safely
	for path, content := range files {
		fullPath := filepath.Join(root, filepath.FromSlash(path))

		// ensure parent folder exists (Windows fix)
		parentDir := filepath.Dir(fullPath)
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			return fmt.Errorf("failed to create parent dir for %s: %w", fullPath, err)
		}

		// create file if it doesn't exist
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
				return fmt.Errorf("failed to write file %s: %w", fullPath, err)
			}
		}
	}

	fmt.Println("✅ Go project created with files:", projectName)
	return nil
}
