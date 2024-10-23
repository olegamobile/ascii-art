package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	font := "standard.txt"
	argList := os.Args
	if len(argList) < 2 {
		log.Fatal("ERROR: not enough arguments.\nUsage: go run ascii-art 'text to convert into ascii art'")
	}

	userText := os.Args[1]
	if !isTextPrintable(userText) {
		log.Fatal("ERROR: input text may contain only printable characters (ASCII codes 32-127)")
	}
	result := convertToAsciiArt(userText, font)
	fmt.Print(result)
}
