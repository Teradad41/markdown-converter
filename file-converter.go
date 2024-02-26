package main

import (
	"fmt"
	"os"

	"github.com/gomarkdown/markdown"
)

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func convertMarkdownToHTML(inputFile, outputFile string) {
	md, err := os.ReadFile(inputFile)
	handleError(err)

	html := markdown.ToHTML(md, nil, nil)

	outFile, err := os.Create(outputFile)
	handleError(err)
	defer outFile.Close()

	_, err = outFile.Write(html)
	handleError(err)

	fmt.Println("Conversion completed successfully.")
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run file-converter.go markdown inputFile outputFile")
		os.Exit(1)
	}

	command := os.Args[1]
	inputFile := os.Args[2]
	outputFile := os.Args[3]

	if command == "markdown" {
		convertMarkdownToHTML(inputFile, outputFile)
	} else {
		fmt.Println("Invalid command:", command)
		os.Exit(1)
	}
}
