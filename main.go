package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"unicode"
	"unicode/utf8"
)

const (
	bufferSize = 4096
)

func main() {
	log.SetFlags(0)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("failed to read current working dir :", err)
	}

	file, err := os.Open(filepath.Join(wd, "words.txt"))
	if err != nil {
		log.Fatalln("failed to open file :", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln("failed to close file :", err)
		}
	}()

	wordCount := CountWordsUsingBufioScanner(file)
	fmt.Println("Word Count =>", wordCount)
}

// CountWordsUsingBufioScanner counts words using a buffered scanner.
func CountWordsUsingBufioScanner(reader io.Reader) int {
	var wordCount int
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		_ = scanner.Text()
		wordCount++
	}

	return wordCount
}

// CountWordsUsingBufioReader counts words using bufio.Reader and checking for whitespace.
func CountWordsUsingBufioReader(reader io.Reader) int {
	var wordCount int
	var isInsideWord bool
	bufReader := bufio.NewReaderSize(reader, bufferSize)

	for {
		r, _, err := bufReader.ReadRune()
		if err != nil {
			break
		}

		// If the rune is not a space and we're not inside a word, it's the start of a new word
		if !unicode.IsSpace(r) && !isInsideWord {
			wordCount++
		}

		isInsideWord = !unicode.IsSpace(r)
	}

	return wordCount
}

// CustomCountWords is a custom implementation for counting words.
func CustomCountWords(reader io.Reader) int {
	var wordCount int
	var isInsideWord bool

	// Buffers for reading input and storing leftover bytes.
	leftover := make([]byte, 0)
	buf := make([]byte, bufferSize)

	for {
		size, err := reader.Read(buf)
		if err != nil {
			break
		}

		// Append any leftover bytes from previous reads to the current buffer.
		subbuf := append(leftover, buf[:size]...)

		for len(subbuf) > 0 {
			r, rsize := utf8.DecodeRune(subbuf)
			if r == utf8.RuneError {
				break
			}

			// If the rune is not whitespace and we're not inside a word, it's a new word.
			subbuf = subbuf[rsize:]
			if !unicode.IsSpace(r) && !isInsideWord {
				wordCount++
			}

			isInsideWord = !unicode.IsSpace(r)
		}

		// Store any leftover bytes that didn't form a complete rune.
		leftover = leftover[:0]
		leftover = append(leftover, subbuf...)
	}

	return wordCount
}
