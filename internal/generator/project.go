package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateProject(projectName string) error {
	root := projectName

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

	for _, d := range dirs {
		if err := os.MkdirAll(filepath.Join(root, d), 0755); err != nil {
			return err
		}
	}

	// go.mod
	goMod := fmt.Sprintf(`module %s

go 1.21
`, projectName)

	writeFileIfNotExist(filepath.Join(root, "go.mod"), goMod)

	// main.go
	mainGo := `package main

func main() {
	// TODO: bootstrap app
}
`
	writeFileIfNotExist(filepath.Join(root, "main.go"), mainGo)

	fmt.Println("✅ Project created:", projectName)
	return nil
}

func writeFileIfNotExist(path, content string) {
	if _, err := os.Stat(path); err == nil {
		return
	}
	_ = os.WriteFile(path, []byte(content), 0644)
}
