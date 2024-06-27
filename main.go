package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ParseCommandLineArgs(argsList []string) error {

	if len(argsList) == 0 {
		return errors.New("no arguments provided")
	}

	if len(argsList) == 1 {
		fileSize := GetFileSize(argsList[0])
		lineCount := GetFileLineCount(argsList[0])
		wordCount := GetWordCount(argsList[0])
		fmt.Println(lineCount, wordCount, fileSize, argsList[0])
		return nil
	}

	switch argsList[0] {
	case "-c":
		fileSize := GetFileSize(argsList[1])
		fmt.Println(fileSize, argsList[1])
	case "-l":
		lineCount := GetFileLineCount(argsList[1])
		fmt.Println(lineCount, argsList[1])
	case "-w":
		wordCount := GetWordCount(argsList[1])
		fmt.Println(wordCount, argsList[1])
	case "-m":
		charCount := GetCharacterCount(argsList[1])
		fmt.Println(charCount, argsList[1])
	}

	return nil
}

func open_file(path string) (*os.File, func()) {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	closeFn := func() {
		err := file.Close()
		if err != nil {
			panic("failed to close input file")
		}
	}

	return file, closeFn
}

func get_count(file os.File, split bufio.SplitFunc) int {
	fileScanner := bufio.NewScanner(&file)
	fileScanner.Split(split)
	count := 0

	for fileScanner.Scan() {
		count++
	}

	return count
}

func GetCharacterCount(path string) int {

	file, closeFn := open_file(path)

	defer closeFn()

	charCount := get_count(*file, bufio.ScanRunes)

	return charCount
}

func GetWordCount(path string) int {
	file, closeFn := open_file(path)

	defer closeFn()

	wordCount := get_count(*file, bufio.ScanWords)

	return wordCount
}

func GetFileLineCount(path string) int {
	file, closeFn := open_file(path)

	defer closeFn()

	lineCount := get_count(*file, bufio.ScanLines)

	return lineCount
}

func GetFileSize(path string) int64 {
	file, err := os.Stat(path)
	if err != nil {
		panic("Error getting file size")
	}

	return file.Size()
}

func main() {
	commandArgs := os.Args[1:]
	ParseCommandLineArgs(commandArgs)
}
