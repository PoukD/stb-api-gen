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

go 1.24.3

require github.com/gin-gonic/gin v1.11.0

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	gorm.io/gorm v1.30.0 // indirect
)

require (
	github.com/bytedance/sonic v1.14.0 // indirect
	github.com/bytedance/sonic/loader v0.3.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/gin-contrib/sse v1.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.27.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/goccy/go-yaml v1.18.0 // indirect
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/quic-go/qpack v0.5.1 // indirect
	github.com/quic-go/quic-go v0.54.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.3.0 // indirect
	go.uber.org/mock v0.5.0 // indirect
	go.uber.org/zap v1.27.1
	golang.org/x/arch v0.20.0 // indirect
	golang.org/x/crypto v0.40.0 // indirect
	golang.org/x/mod v0.25.0 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	golang.org/x/tools v0.34.0 // indirect
	google.golang.org/protobuf v1.36.9 // indirect
	gorm.io/driver/mysql v1.6.0
)

`, project),
	}

	for p, c := range rootFiles {
		if err := writeFile(root, p, c); err != nil {
			return err
		}
	}

	// base dirs
	baseDirs := []string{
		"cmd/app",
		"cmd/middleware/health",
		"cmd/middleware/log",
		"cmd/middleware/request",
		"cmd/middleware/verify",
		"cmd/routes/external/controller/authorization/sign/database/entity",
		"cmd/routes/external/controller/authorization/sign/database/service",
		"cmd/routes/external/controller/authorization/sign/domain/api",
		"cmd/routes/external/controller/authorization/sign/domain/http",
		"cmd/routes/external/controller/authorization/sign/domain/model",
		"cmd/routes/external/controller/authorization/sign/model",
		"cmd/routes/internal/controller/authorization/sign/database/entity",
		"cmd/routes/internal/controller/authorization/sign/database/service",
		"cmd/routes/internal/controller/authorization/sign/domain/api",
		"cmd/routes/internal/controller/authorization/sign/domain/http",
		"cmd/routes/internal/controller/authorization/sign/domain/model",
		"cmd/routes/internal/controller/authorization/sign/model",
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
		"cmd/routes/external/controller/authorization/sign/database/entity/signEntity.go":   "package entity\n",
		"cmd/routes/external/controller/authorization/sign/database/service/signService.go": "package service\n",
		"cmd/routes/external/controller/authorization/sign/domain/api/api.go":               "package api\n",
		"cmd/routes/external/controller/authorization/sign/domain/http/httpInterface.go":    "package http\n",
		"cmd/routes/external/controller/authorization/sign/domain/model/signHttp.go":        "package model\n",
		"cmd/routes/external/controller/authorization/sign/model/signModel.go":              "package model\n",
		"cmd/routes/internal/controller/authorization/sign/database/entity/signEntity.go":   "package entity\n",
		"cmd/routes/internal/controller/authorization/sign/database/service/signService.go": "package service\n",
		"cmd/routes/internal/controller/authorization/sign/domain/api/api.go":               "package api\n",
		"cmd/routes/internal/controller/authorization/sign/domain/http/httpInterface.go":    "package http\n",
		"cmd/routes/internal/controller/authorization/sign/domain/model/signHttp.go":        "package model\n",
		"cmd/routes/internal/controller/authorization/sign/model/signModel.go":              "package model\n",
		"config/configStatus.go":         "package database\n",
		"config/database/databaseLog.go": "package database\n",
		"config/http/httpConfig.go":      "package http\n",
		"config/http/model/httpModel.go": "package model\n",
		"util/bcrypt/bcrypt.go":          "package bcrypt\n",
		"util/jwt/jwt.go":                "package jwt\n",
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
