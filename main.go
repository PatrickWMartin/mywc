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

func getBytes(input []byte) int {
	return len(input)
}

func getWords(input string) int {
	return len(strings.Fields(input))
}

func getChatacters(input string) int {
	return len(strings.Split(input, ""))
}

func getLines(input string) int {
	return len(strings.Split(input, "\n")) - 1
}

func main() {
	var bytes []byte
	if checkStdIn() {
		bytes, _ = io.ReadAll(os.Stdin)
	} else {
		bytes, _ = os.ReadFile("test.txt") // just pass the file name
	}

	fmt.Println(getBytes(bytes))
	input := string(bytes)
	fmt.Println(getWords(input))
	fmt.Println(getChatacters(input))
	fmt.Println(getLines(input))
}
