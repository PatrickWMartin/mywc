package main

import (
	"fmt"
	// "io"
	// "os"
	// "strings"
	"flag"
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

    fmt.Println(files) 
}

// func getBytes(input []byte) int {
// 	return len(input)
// }
//
// func getWords(input string) int {
// 	return len(strings.Fields(input))
// }
//
// func getChatacters(input string) int {
// 	return len(strings.Split(input, ""))
// }
//
// func getLines(input string) int {
// 	return len(strings.Split(input, "\n")) - 1
// }
