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
    case "-m":
        GetCharacterCount(argsList[1])
    }

	return nil
}

func open_file(path string) (*os.File, func()){
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

func get_count(file os.File, split bufio.SplitFunc) int  {
    fileScanner := bufio.NewScanner(&file)
    fileScanner.Split(split)
    count := 0

    for fileScanner.Scan() {
        count++
    }

    return count
}

func GetCharacterCount(path string) error {

    file, closeFn := open_file(path)
    
    defer closeFn()
    
    charCount := get_count(*file, bufio.ScanRunes)

    fmt.Println(charCount, path)
    return nil
}

func GetWordCount(path string) error {
    file, closeFn := open_file(path)
    
    defer closeFn()
    
    wordCount := get_count(*file, bufio.ScanWords)

    fmt.Println(wordCount, path)
    return nil
}

func GetFileLineCount(path string) error {
    file, closeFn := open_file(path)
    
    defer closeFn()
    
    lineCount := get_count(*file, bufio.ScanLines)

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
