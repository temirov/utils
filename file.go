package utils

import (
	"bufio"
	"bytes"
	"fmt"

	"io"
	"log"
	"os"
	"path/filepath"
)

func RemoveAll(tempDir string) {
	err := os.RemoveAll(tempDir)
	if err != nil {
		log.Printf("failed to remove folder %s: %v", tempDir, err)
	}
}

func CloseFile(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.Printf("failed to close file descriptor: %v", err)
	}
}

func RemoveFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("failed to delete the file: %v", err)
	}
}

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

func ReadFile(filePath string) (*bytes.Reader, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(fileContent), nil
}
