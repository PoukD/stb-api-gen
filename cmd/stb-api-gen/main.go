package main

import (
	"fmt"
	"os"

	"github.com/PoukD/stb-api-gen/internal/generator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  stb-api-gen <project-name>")
		os.Exit(1)
	}

	projectName := os.Args[1]

	if err := generator.CreateProject(projectName); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}
