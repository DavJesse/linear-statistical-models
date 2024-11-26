package main

import (
	"log"
	"os"
	"strings"

	"linear-stats/files"
)

func main() {
	args := os.Args[1:]

	// Handle insufficient or incorrect number of arguments
	if len(args) != 1 {
		log.Fatal("Please provide a single file path to a valid text file as an argument.")
	}

	// Define file path from arguments
	// Check file validity from its extenstion
	filePath := args[0]
	if !strings.HasPrefix(filePath, ".txt") {
		log.Fatalf("Invalid file format: %v is not a text file", filePath)
	}

	data := files.ReadFile(filePath)

	// Filter out empty files parsed as data
	if len(data) == 0 {
		log.Fatal("File parsed as data is empty.")
	}
}
