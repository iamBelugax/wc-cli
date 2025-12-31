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

	file, err := os.Open(filepath.Join(wd, "utf-8.txt"))
	if err != nil {
		log.Fatalln("failed to open file :", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln("failed to close file :", err)
		}
	}()

	wordCount := CountWords(file)
	fmt.Println("Word Count =>", wordCount)
}

func CountWords(reader io.Reader) int {
	var wordCount int
	var isInsideWord bool
	bufReader := bufio.NewReaderSize(reader, bufferSize)

	for {
		r, _, err := bufReader.ReadRune()
		if err != nil {
			break
		}

		if !unicode.IsSpace(r) && !isInsideWord {
			wordCount++
		}
		isInsideWord = !unicode.IsSpace(r)
	}

	return wordCount
}

func CustomCountWords(reader io.Reader) int {
	var wordCount int
	var isInsideWord bool

	leftover := make([]byte, 0)
	buf := make([]byte, bufferSize)

	for {
		size, err := reader.Read(buf)
		if err != nil {
			break
		}

		subbuf := append(leftover, buf[:size]...)

		for len(subbuf) > 0 {
			r, rsize := utf8.DecodeRune(subbuf)
			if r == utf8.RuneError {
				break
			}

			subbuf = subbuf[rsize:]
			if !unicode.IsSpace(r) && !isInsideWord {
				wordCount++
			}

			isInsideWord = !unicode.IsSpace(r)
		}

		leftover = leftover[:0]
		leftover = append(leftover, subbuf...)
	}

	return wordCount
}
