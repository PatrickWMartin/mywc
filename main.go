package main

import (
	"bufio"
	"fmt"
	"os"
    "io/fs"
)

func checkForFlag(input string) bool {
    if len(input) > 1 && input[0] == '-'{
        return true
    }
    return false
}

func checkStdIn() bool {
    stat, err := os.Stdin.Stat()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error checking stdin:", err)
        os.Exit(1)
    }
    if (stat.Mode() & os.ModeCharDevice) == 0 {
        return true
    }

    return false
}

func ParseCommandLineArgs(argsList []string)  {

    // has to have input from stdin or its invaled 
    if len(argsList) == 0 {
        if checkStdIn() == false {
            fmt.Fprintln(os.Stderr, "Error no argument provided")
        }
        fileSize := GetFileSize("")
		lineCount := GetFileLineCount("")
		wordCount := GetWordCount("")
		fmt.Println(lineCount, wordCount, fileSize)
        os.Exit(0)
    }

    if len(argsList) == 1{
        if checkForFlag(argsList[0]) && checkStdIn(){
            executeFlag(argsList[0], "") 
            os.Exit(0) 
        }

        fileSize := GetFileSize(argsList[0])
		lineCount := GetFileLineCount(argsList[0])
		wordCount := GetWordCount(argsList[0])
		fmt.Println(lineCount, wordCount, fileSize, argsList[0])
        os.Exit(0)
    }

    executeFlag(argsList[0], argsList[1]) 

}

func executeFlag(flag string, path string){
    switch flag {
	case "-c":
		fileSize := GetFileSize(path)
		fmt.Println(fileSize, path)
	case "-l":
		lineCount := GetFileLineCount(path)
		fmt.Println(lineCount, path)
	case "-w":
		wordCount := GetWordCount(path)
		fmt.Println(wordCount, path)
	case "-m":
		charCount := GetCharacterCount(path)
		fmt.Println(charCount, path)
    default:
        fmt.Println("wc: illegal option -- ", flag)
	}

}

func open_file(path string) (*os.File, func()) {
    var file *os.File
    var err error

    if path == ""{
        file = os.Stdin
        err = nil
    } else {
        file, err = os.Open(path)
    }
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
    
	closeFn := func() {
        if path != ""{
            err := file.Close()
            if err != nil {
                panic("failed to close input file")
            }
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
    var file fs.FileInfo
    var err error
    if path == ""{
        file, err = os.Stdin.Stat()
    } else {
        file, err = os.Stat(path)
    }
	if err != nil {
		panic("Error getting file size")
	}

	return file.Size()
}


func main() {
	commandArgs := os.Args[1:]
	ParseCommandLineArgs(commandArgs)
}
