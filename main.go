package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	font := os.Args[2]
	text := os.Args[1]

	content, err := readFileContent(font)
	if err != nil {
		fmt.Println(err)
		return
	}

	symbols := strings.Split(content, "\n")

	for i := 1; i <= 8; i++ {
		for _, char := range text {
			fmt.Print(symbols[(int(char)-32)*9+i])
		}
		fmt.Println()
	}

}

func readFileContent(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("Error reading file %s: %v", filename, err)
	}
	return string(content), nil
}
