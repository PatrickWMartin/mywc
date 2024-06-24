package main

import (
    "fmt"
    "os"
    "errors"
)

func ParseCommandLineArgs(argsList []string) error{

    if len(argsList) == 0 {
        return errors.New("no arguments provided")
    }

    if argsList[0] == "-c"{
        file, err := os.Stat(argsList[1]) 
        if err != nil {
            return err
        }

        fmt.Println(file.Size(), argsList[1])
    }

    return nil
}

func main() {
    commandArgs := os.Args[1:]
    ParseCommandLineArgs(commandArgs)
}
