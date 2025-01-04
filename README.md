# Utils

### **`to_markdown` Utility**

The **`to_markdown`** utility is a command-line tool designed to **clean and convert HTML or Markdown files** into **clean Markdown format**. It processes raw HTML and Markdown files, removing unnecessary comments, metadata, and unsupported tags, ensuring the output is well-structured and ready for use in documentation or publishing.

---

### **Features**

- **Supports Both HTML and Markdown Input**
    - Handles `.html` and `.md` files as input, making it versatile for mixed-content repositories.

- **Cleans Inline and Block Comments**
    - Removes HTML comments (`<!-- ... -->`), including multiline metadata blocks.

- **Preserves Formatting and Content**
    - Maintains semantic structure while cleaning redundant elements like custom tags or styling.

- **Markdown Rendering**
    - Converts Markdown files to HTML and then back to clean Markdown, ensuring consistent formatting.

- **HTML Rendering**
    - Processes raw HTML into valid, structured Markdown output.

- **Modern Go Implementation**
    - Written in **Go**, leveraging libraries like **Goldmark** for Markdown parsing and **html-to-markdown** for HTML cleaning.

---

### **Installation**

1. **Build the Binary:**

```bash
go build -o bin/to_markdown ./cmd/to_markdown
```

2. **Move to Path (Optional):**

```bash
sudo mv bin/to_markdown /usr/local/bin/
```

3. **Verify Installation:**

```bash
to_markdown --help
```

---

### **Usage**

```bash
to_markdown input.[html|md] output.md
```

- **Input File:**
    - Can be an HTML or Markdown file.
- **Output File:**
    - The resulting clean Markdown file.

**Example 1: HTML to Clean Markdown**

```bash
to_markdown input.html output.md
```

**Example 2: Markdown Cleanup**

```bash
to_markdown input.md output.md
```

---

### **Example Input and Output**

**Input (HTML):**

```html
<h1>Hello World</h1>
<p>This is <strong>bold</strong> text.</p>
<!-- This is a comment -->
<div class="highlight">This is styled</div>

This is styled
<ul>
  <li>List item 1</li>
  <li>List item 2</li>
</ul>
```

**Output (Markdown):**

```markdown
# Hello World

This is **bold** text.

This is styled

- List item 1
- List item 2
```

---

### **Testing**

The tool includes **table-driven tests** to ensure consistent behavior for a variety of inputs.

**Run Tests:**

```bash
go test ./test -v
```

---

### **Dependencies**

- **[Goldmark](https://github.com/yuin/goldmark)** - Markdown rendering and parsing.
- **[html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown)** - HTML-to-Markdown conversion and cleaning.
- **[net/html](https://pkg.go.dev/golang.org/x/net/html)** - HTML parsing and rendering.

---

### **Contributing**

Contributions are welcome!

1. Fork the repository.
2. Create a new branch (`feature/my-feature`).
3. Commit changes and submit a pull request.

---

### **License**

This project is licensed under the **MIT License**. See the **LICENSE** file for details.



