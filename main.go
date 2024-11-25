package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		log.Fatal("Please provide a single file path to a valid text file as an argument.")
	}
	filePath := args[0]
}
