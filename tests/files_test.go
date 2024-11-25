package tests

import (
	"os"
	"testing"

	"linear-stats/files"
)

func TestReadFile_EmptyFile(t *testing.T) {
	// Create a temporary empty file
	tempFile, err := os.CreateTemp("", "empty.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	// Call ReadFile with the empty file
	result := files.ReadFile(tempFile.Name())

	// Check if the result is an empty slice
	if len(result) != 0 {
		t.Errorf("Expected empty slice, but got slice with length %d", len(result))
	}
}
