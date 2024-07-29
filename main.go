package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

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

func main(){
    var bytes []byte
    if checkStdIn() {
        bytes,_ = io.ReadAll(os.Stdin)
    } else {
        bytes,_ = os.ReadFile("test.txt") // just pass the file name
    }
    test := string(bytes)
    fmt.Println(len(bytes))
    fmt.Println(len(strings.Split(test, "")))
    fmt.Println(len(strings.Fields(test)))
    fmt.Println(len(strings.Split(test, "\n"))-1)
}
