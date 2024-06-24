package main

import (
	"errors"
	"fmt"
	"os"
)

func ParseCommandLineArgs(argsList []string) error {

	if len(argsList) == 0 {
		return errors.New("no arguments provided")
	}

	if argsList[0] == "-c" {
		GetFileSize(argsList[1])
	}

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
