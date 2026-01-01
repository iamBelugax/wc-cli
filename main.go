package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	total := 0
	filenames := os.Args[1:]
	var hasErrorOccurred bool

	for _, filename := range filenames {
		wordCount, err := CountWordsInFile(filename)
		if err != nil {
			hasErrorOccurred = true
			fmt.Fprintln(os.Stderr, "wc:", err)
			continue
		}

		total += wordCount
		fmt.Println(wordCount, filename)
	}

	if len(filenames) == 0 {
		fmt.Println(CountWords(os.Stdin))
		os.Exit(0)
	}

	if len(filenames) > 1 {
		fmt.Println(total, "total")
	}

	if hasErrorOccurred {
		os.Exit(1)
	}
}
