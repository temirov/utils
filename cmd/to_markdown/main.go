package main

import (
	"fmt"
	"github.com/temirov/utils/pkg/render"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go input.[md|html] output.md")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Check if input file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		log.Fatalf("Input file does not exist: %s\n", inputFile)
	}

	// Determine file type based on extension
	ext := strings.ToLower(filepath.Ext(inputFile))

	var rendered string
	var err error

	switch ext {
	case ".html":
		rendered, err = render.HTML(inputFile) // Render HTML
	case ".md":
		rendered, err = render.Markdown(inputFile) // Render Markdown
	default:
		log.Fatalf("Unsupported file type: %s\n", ext)
	}
	if err != nil {
		log.Fatalf("Error rendering input file: %v\n", err)
	}

	// Save the rendered content as clean Markdown
	err = render.SaveHTMLAsMarkdown(rendered, outputFile)
	if err != nil {
		log.Fatalf("Error saving Markdown: %v\n", err)
	}

	fmt.Printf("Processed and saved clean Markdown to: %s\n", outputFile)
}
