package main

import (
	"flag"
	"fmt"
	"os"
	"ccwc/internal/files"
)

func main() {
	// Define flags
	charFlag := flag.Bool("c", false, "Count characters")
	lineFlag := flag.Bool("l", false, "Count lines")
	wordFlag := flag.Bool("w", false, "Count words")
	flag.Parse()

	// Ensure a file path is provided
	if len(flag.Args()) == 0 {
		fmt.Println("Usage: wc [-c] [-l] [-w] <file_path>")
		os.Exit(1)
	}

	// Open the file
	filePath := flag.Args()[0]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Process file to count characters, words, and lines
	fileInfo, err := files.FileInfo(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing file: %v\n", err)
		os.Exit(1)
	}

	// Output results based on flags
	if *charFlag {
		fmt.Printf("Characters: %d\n", fileInfo.CharCount)
	}
	if *lineFlag {
		fmt.Printf("Lines: %d\n", fileInfo.LineCount)
	}
	if *wordFlag {
		fmt.Printf("Words: %d\n", fileInfo.WordCount)
	}

	// Default behavior: If no flags are set, print all counts
	if !*charFlag && !*lineFlag && !*wordFlag {
		fmt.Printf("Lines: %d  Words: %d  Characters: %d\n", fileInfo.LineCount, fileInfo.WordCount, fileInfo.CharCount)
	}
}