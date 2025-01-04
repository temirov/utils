package render

import (
	"bytes"
	"fmt"
	"os"

	htmltomd "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	renderer "github.com/yuin/goldmark/renderer/html"
	"golang.org/x/net/html"
)

// HTML - Renders and formats raw HTML as valid HTML.
func HTML(inputFile string) (string, error) {
	// Read input HTML file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input HTML file: %v", err)
	}

	// Parse and render HTML using net/html
	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %v", err)
	}

	var buf bytes.Buffer
	if err := html.Render(&buf, doc); err != nil {
		return "", fmt.Errorf("failed to render HTML: %v", err)
	}

	return buf.String(), nil
}

// Markdown - Converts Markdown to HTML.
func Markdown(inputFile string) (string, error) {
	// Read input Markdown file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		return "", fmt.Errorf("failed to read input Markdown file: %v", err)
	}

	// Render Markdown as HTML using Goldmark
	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM), // GitHub Flavored Markdown
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		goldmark.WithRendererOptions(renderer.WithUnsafe()), // Allow unsafe HTML
	)

	err = md.Convert(content, &buf)
	if err != nil {
		return "", fmt.Errorf("failed to render Markdown to HTML: %v", err)
	}
	return buf.String(), nil
}

// SaveHTMLAsMarkdown rendered HTML as clean Markdown
func SaveHTMLAsMarkdown(htmlContent, outputFile string) error {
	// Use html-to-markdown to clean and convert HTML to Markdown
	converter := htmltomd.NewConverter("", true, nil) // Initialize the converter
	markdown, err := converter.ConvertString(htmlContent)
	if err != nil {
		return fmt.Errorf("failed to convert HTML to Markdown: %v", err)
	}

	// Write the cleaned Markdown to output file
	err = os.WriteFile(outputFile, []byte(markdown), 0644)
	if err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}

	return nil
}
