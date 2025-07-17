// Package file provides helpers for common file system interactions such as
// creating, reading and removing files.
package file

import (
	"bufio"
	"bytes"
	"fmt"

	"io"
	"log"
	"os"
	"path/filepath"
)

// RemoveAll recursively deletes the directory at tempDir. Any error encountered
// is logged but not returned.
func RemoveAll(tempDir string) {
	err := os.RemoveAll(tempDir)
	if err != nil {
		log.Printf("failed to remove folder %s: %v", tempDir, err)
	}
}

// CloseFile closes the provided io.Closer and logs an error if closing fails.
func CloseFile(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.Printf("failed to close file descriptor: %v", err)
	}
}

// RemoveFile deletes the file located at filePath. Errors are logged and not
// returned to the caller.
func RemoveFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("failed to delete the file: %v", err)
	}
}

// ReadLines reads a text file line by line and returns a slice of strings. The
// file is opened for reading and closed automatically. An error is returned if
// the file cannot be opened or scanned.
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s, error: %w", filename, err)
	}
	defer CloseFile(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filename, err)
	}

	return lines, nil
}

// SaveFile writes fileContent to a file named fileName.html inside outputDir.
// The directory is created if it does not already exist.
func SaveFile(outputDir string, fileName string, fileContent []byte) error {
	directoryError := os.MkdirAll(outputDir, 0755)
	if directoryError != nil {
		return directoryError
	}

	outputFilePath := filepath.Join(outputDir, fmt.Sprintf("%s.html", fileName))
	writeError := os.WriteFile(outputFilePath, fileContent, 0644)
	if writeError != nil {
		return writeError
	}

	return nil
}

// ReadFile reads a file from disk and returns a bytes.Reader containing the
// file contents.
func ReadFile(filePath string) (*bytes.Reader, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(fileContent), nil
}
