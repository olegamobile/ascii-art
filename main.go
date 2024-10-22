package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	font := "standard.txt"
	argList := os.Args
	if len(argList) < 2 {
		fmt.Println("Not enough arguments.\nUsage: go run main.go <text to convert into ascii art>")
	}

	if len(argList) == 3 {
		font = os.Args[2]
	}
	text := os.Args[1]

	content, err := readFileContent(font)
	if err != nil {
		log.Fatal(err)
		return
	}

	symbols := strings.Split(content, "\n")
	textSlice := strings.Split(text, "\\n")
	
	// reduce textSlice if it is empty in order not to print additional line
	if isSliceEmpty(textSlice) {
		textSlice = textSlice[1:]
	}
	for _, textLine := range textSlice {
		if len(textLine) > 0 {
			for i := 1; i <= 8; i++ {
				for _, char := range textLine {
					fmt.Print(symbols[(int(char)-32)*9+i])
				}
				fmt.Println()
			}
		} else {
			fmt.Println()
		}
	}
}

func isSliceEmpty(inputSlice []string) bool {
	for _, textLine := range inputSlice {
		if len(textLine) > 0 {
			return false
		}
	}
	return true
}

func readFileContent(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: \n%v", filename, err)
	}
	return string(content), nil
}
