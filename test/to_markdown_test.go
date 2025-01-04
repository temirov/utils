package test

import (
	"github.com/temirov/utils/pkg/file"
	"github.com/temirov/utils/pkg/render"
	"github.com/temirov/utils/pkg/text"
	"os"
	"testing"
)

func TestRenderAndSave(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		render   func(string) (string, error) // Flexible renderer for each test
	}{
		{
			name: "Clean HTML with inline comment",
			input: `<h1>Hello World</h1>
<p>This is <strong>bold</strong> text.</p>
<!-- This is a comment -->
<div class="highlight">This is styled</div>
<ul>
  <li>List item 1</li>
  <li>List item 2</li>
</ul>`,
			expected: `# Hello World

This is **bold** text.
This is styled

- List item 1
- List item 2
`,
			render: render.HTML, // Use HTML renderer
		},
		{
			name: "Markdown with comments and metadata",
			input: `
# Divider

<!--*
# Document freshness
*-->

A [divider](https://example.com) is a thin line.

- Item 1
- Item 2
`,
			expected: `# Divider

A [divider](https://example.com) is a thin line.

- Item 1
- Item 2
`,
			render: render.Markdown, // Use Markdown renderer
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temp input and output files
			inputFile, err := os.CreateTemp("", "input-*")
			if err != nil {
				t.Fatalf("Failed to create input file: %v", err)
			}
			defer file.RemoveFile(inputFile.Name())

			outputFile, err := os.CreateTemp("", "output-*")
			if err != nil {
				t.Fatalf("Failed to create output file: %v", err)
			}
			defer file.RemoveFile(outputFile.Name())

			// Write input content to temp file
			if _, err := inputFile.Write([]byte(tt.input)); err != nil {
				t.Fatalf("Failed to write to input file: %v", err)
			}
			file.CloseFile(inputFile)

			// Render input
			rendered, err := tt.render(inputFile.Name())
			if err != nil {
				t.Fatalf("Render failed: %v", err)
			}

			// Save rendered output as Markdown
			err = render.SaveHTMLAsMarkdown(rendered, outputFile.Name())
			if err != nil {
				t.Fatalf("Failed to save markdown: %v", err)
			}

			// Read output file
			output, err := os.ReadFile(outputFile.Name())
			if err != nil {
				t.Fatalf("Failed to read output file: %v", err)
			}

			// Compare output with expected result
			if text.Normalize(string(output)) != text.Normalize(tt.expected) {
				t.Errorf("Output mismatch.\nExpected:\n%s\nGot:\n%s", tt.expected, string(output))
			}
		})
	}
}
