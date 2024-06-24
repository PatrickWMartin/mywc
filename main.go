package main

import (
    "fmt"
    "os"
)
func main() {
    commandArgs := os.Args[1:]
    fmt.Println(commandArgs) 
}
