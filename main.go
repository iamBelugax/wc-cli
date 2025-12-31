package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("failed to read current working dir :", err)
	}

	filepath := filepath.Join(wd, "words.txt")
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalln("failed to read file :", err)
	}

	fmt.Println("Word Count =>", CountWords(data))
}

func CountWords(data []byte) int {
	if len(data) == 0 {
		return 0
	}

	var wordCount int
	for _, char := range data {
		if char == ' ' {
			wordCount++
		}
	}

	return wordCount + 1
}
