package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func isSliceEmpty(inputSlice []string) bool {
	for _, textLine := range inputSlice {
		if len(textLine) > 0 {
			return false
		}
	}
	return true
}

func isTextPrintable(text string) bool {
	for _, char := range text {
		if char < 32 || char > 127 {
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

func convertToAsciiArt(inputText, font string) string {
	result := ""
	bannerContent, err := readFileContent(font)
	if err != nil {
		log.Fatal(err)
	}
	bannerLines := strings.Split(bannerContent, "\n")

	// split input text into the separate lines
	textSlice := strings.Split(inputText, "\\n")

	// reduce textSlice if it is empty in order not to print additional line
	if isSliceEmpty(textSlice) {
		textSlice = textSlice[1:]
	}
	for _, textLine := range textSlice {
		if len(textLine) > 0 {
			for i := 1; i <= 8; i++ {
				for _, char := range textLine {
					result += bannerLines[(int(char)-32)*9+i]
				}
				result += "\n"
			}
		} else {
			result += "\n"
		}
	}
	return result
}
