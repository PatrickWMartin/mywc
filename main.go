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
	var options printOptions

	flag.BoolVar(&options.printBytes, "c", false, "Count bytes")
	flag.BoolVar(&options.printLines, "l", false, "Count lines")
	flag.BoolVar(&options.printWords, "w", false, "Count words")
	flag.BoolVar(&options.printChars, "m", false, "Count characters")
	flag.Parse()

	if !options.printBytes && !options.printLines && !options.printWords && !options.printChars {
		options.printBytes = true
		options.printWords = true
		options.printLines = true
	}

	files := flag.CommandLine.Args()
	if len(files) == 0 {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		printStats(bytes, options, "")
	} else {
		for _, file := range files {
			bytes, err := os.ReadFile(file) // just pass the file name
			if err != nil {
				log.Fatal(err)
			}
			printStats(bytes, options, file)
		}
	}
}

type printOptions struct {
	printBytes bool
	printLines bool
	printWords bool
	printChars bool
}

func countBytes(input []byte) int {
	return len(input)
}

func countWords(input string) int {
	return len(strings.Fields(input))
}

func countCharacters(input string) int {
	return len(strings.Split(input, ""))
}

func countLines(input string) int {
	return len(strings.Split(input, "\n")) - 1
}

func printStats(input []byte, options printOptions, fileName string) {
	output := ""
	inputString := string(input)

	if options.printLines {
		output += fmt.Sprint(countLines(inputString), " ")
	}

	if options.printWords {
		output += fmt.Sprint(countWords(inputString), " ")
	}

	if options.printChars {
		output += fmt.Sprint(countCharacters(inputString), " ")
	}

	if options.printBytes {
		output += fmt.Sprint(countBytes(input), " ")
	}

	fmt.Println(output, fileName)
}
