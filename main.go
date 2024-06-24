package main

import (
	"errors"
	"fmt"
	"os"
    "bufio"
)

func ParseCommandLineArgs(argsList []string) error {

	if len(argsList) == 0 {
		return errors.New("no arguments provided")
	}

    switch argsList[0] {
    case "-c":
		GetFileSize(argsList[1])
    case "-l":
        GetFileLineCount(argsList[1])
    case "-w":
        GetWordCount(argsList[1])
    }

	return nil
}

func GetWordCount(path string) error {

    file, err := os.Open(path)
    if err != nil {
        return err
    }

    fileScanner := bufio.NewScanner(file)
    fileScanner.Split(bufio.ScanWords)
    wordCount := 0

    for fileScanner.Scan() {
        wordCount++
    }

    fmt.Println(wordCount, path)
    return nil
}

func GetFileLineCount(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }

    fileScanner := bufio.NewScanner(file)
    lineCount := 0

    for fileScanner.Scan() {
        lineCount++
    }

    fmt.Println(lineCount, path)
    return nil
}

func GetFileSize(path string) error {
	file, err := os.Stat(path)
	if err != nil {
		return err
	}

	fmt.Println(file.Size(), path)
	return nil
}

func main() {
	commandArgs := os.Args[1:]
	ParseCommandLineArgs(commandArgs)
}
