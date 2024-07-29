package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var printBytes, printLines, printWords, printChars bool

	flag.BoolVar(&printBytes, "c", false, "Count bytes")
	flag.BoolVar(&printLines, "l", false, "Count lines")
	flag.BoolVar(&printWords, "w", false, "Count words")
	flag.BoolVar(&printChars, "m", false, "Count characters")
	flag.Parse()

	if !printBytes && !printLines && !printWords && !printChars {
		printBytes = true
		printWords = true
		printLines = true
	}

	files := flag.CommandLine.Args()
	if len(files) == 0 {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		printStats(bytes)
	} else {
		bytes, err := os.ReadFile("test.txt") // just pass the file name
		if err != nil {
			log.Fatal(err)
		}
		printStats(bytes)
	}
}

func countBytes(input []byte) int {
	return len(input)
}

func countWords(input string) int {
	return len(strings.Fields(input))
}

func countChatacters(input string) int {
	return len(strings.Split(input, ""))
}

func countLines(input string) int {
	return len(strings.Split(input, "\n")) - 1
}

func printStats(input []byte) {
	fmt.Println(countBytes(input))
	inputString := string(input)
	fmt.Println(countWords(inputString))
	fmt.Println(countChatacters(inputString))
	fmt.Println(countLines(inputString))
}
