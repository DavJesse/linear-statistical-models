package tests

import (
	"os"
	"reflect"
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

func TestReadFile_SingleDigitIntegers(t *testing.T) {
	// Create a temporary file with single-digit integers
	tempFile, err := os.CreateTemp("", "single_digits.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write single-digit integers to the file
	content := []byte("1\n2\n3\n4\n5\n")
	if _, err := tempFile.Write(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	// Call ReadFile with the temporary file
	result := files.ReadFile(tempFile.Name())

	// Check if the result matches the expected output
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
