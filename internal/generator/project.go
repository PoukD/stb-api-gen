package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

// ---------- helpers ----------

func mkdir(root string, paths ...string) error {
	return os.MkdirAll(filepath.Join(append([]string{root}, paths...)...), 0755)
}

func writeFile(root, path, content string) error {
	full := filepath.Join(root, filepath.FromSlash(path))
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		return err
	}
	return os.WriteFile(full, []byte(content), 0644)
}

// ---------- modules ----------

func createInternalCMS(root string) error {
	base := "cmd/routes/internal/cms"

	dirs := []string{
		"database/entity",
		"database/service",
		"domain/http",
		"domain/model",
		"model",
	}

	for _, d := range dirs {
		if err := mkdir(root, base, d); err != nil {
			return err
		}
	}

	files := map[string]string{
		base + "/cmsController.go":               "package cms\n",
		base + "/database/entity/cmsEntity.go":   "package entity\n",
		base + "/database/service/cmsService.go": "package service\n",
		base + "/domain/http/httpInterface.go":   "package http\n",
		base + "/domain/model/cmsHttp.go":        "package model\n",
		base + "/model/cmsModel.go":              "package model\n",
	}

	for p, c := range files {
		if err := writeFile(root, p, c); err != nil {
			return err
		}
	}
	return nil
}

func createExternalAuthorization(root string, name string) error {
	base := filepath.Join(
		"cmd/routes/external/controller/authorization",
		name,
	)

	dirs := []string{
		"database/entity",
		"database/service",
		"domain/api",
		"domain/http",
		"domain/model",
		"model",
	}

	for _, d := range dirs {
		if err := mkdir(root, base, d); err != nil {
			return err
		}
	}

	files := map[string]string{
		base + "/" + name + "Controller.go":               "package " + name + "\n",
		base + "/database/entity/" + name + "Entity.go":   "package entity\n",
		base + "/database/service/" + name + "Service.go": "package service\n",
		base + "/domain/api/api.go":                       "package api\n",
		base + "/domain/http/httpInterface.go":            "package http\n",
		base + "/domain/model/httpModel.go":               "package model\n",
		base + "/model/" + name + "Model.go":              "package model\n",
	}

	for p, c := range files {
		if err := writeFile(root, p, c); err != nil {
			return err
		}
	}
	return nil
}

// ---------- project ----------

func CreateProject(project string) error {
	root := project

	// root files
	rootFiles := map[string]string{
		".env":       "",
		".gitignore": "bin/\n*.log\n.env\n",
		"README.md":  "# " + project + "\n",
		"go.mod": fmt.Sprintf(`module %s

go 1.21
`, project),
	}

	for p, c := range rootFiles {
		if err := writeFile(root, p, c); err != nil {
			return err
		}
	}

	// base dirs
	baseDirs := []string{
		".vscode",
		"cmd/app",
		"cmd/middleware/health",
		"cmd/middleware/log",
		"cmd/middleware/request",
		"cmd/middleware/verify",
		"cmd/routes",
		"config/database",
		"config/http/model",
		"logs/app",
		"logs/http",
		"logs/request",
		"util/bcrypt",
		"util/jwt",
	}

	for _, d := range baseDirs {
		if err := mkdir(root, d); err != nil {
			return err
		}
	}

	// base files
	baseFiles := map[string]string{
		"cmd/app/main.go":                         "package main\n\nfunc main() {}\n",
		"cmd/routes/router.go":                    "package routes\n",
		"cmd/middleware/health/healthCheck.go":    "package health\n",
		"cmd/middleware/log/log_http.go":          "package log\n",
		"cmd/middleware/log/log_error.go":         "package log\n",
		"cmd/middleware/log/log_request.go":       "package log\n",
		"cmd/middleware/request/headerRequest.go": "package request\n",
		"config/database/databaseLog.go":          "package database\n",
		"config/http/httpConfig.go":               "package http\n",
		"config/http/model/httpModel.go":          "package model\n",
		"util/bcrypt/bcrypt.go":                   "package bcrypt\n",
		"util/jwt/jwt.go":                         "package jwt\n",
	}

	for p, c := range baseFiles {
		if err := writeFile(root, p, c); err != nil {
			return err
		}
	}

	// modules
	if err := createInternalCMS(root); err != nil {
		return err
	}

	if err := createExternalAuthorization(root, "register"); err != nil {
		return err
	}
	if err := createExternalAuthorization(root, "sign"); err != nil {
		return err
	}

	fmt.Println("✅ Project created:", project)
	return nil
}
