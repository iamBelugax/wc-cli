package counter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	offsetStart = 0
	bufSize     = 4096
)

type Counts struct {
	Words int
	Lines int
	Bytes int
}

func (c Counts) Print(w io.Writer, suffix ...string) {
	fmt.Fprint(w, c.Lines, c.Words, c.Bytes)
	if len(suffix) > 0 {
		fmt.Fprint(w, " ", suffix[len(suffix)-1])
	}
	fmt.Fprint(w, "\n")
}

func (c *Counts) Add(other Counts) {
	c.Bytes += other.Bytes
	c.Lines += other.Lines
	c.Words += other.Words
}

func CountAll(r io.ReadSeeker) Counts {
	words := CountWords(r)

	r.Seek(offsetStart, io.SeekStart)
	lines := CountLines(r)

	r.Seek(offsetStart, io.SeekStart)
	bytes := CountBytes(r)

	return Counts{
		Words: words,
		Lines: lines,
		Bytes: bytes,
	}
}

func CountFile(path string) (Counts, error) {
	f, err := os.Open(path)
	if err != nil {
		return Counts{}, err
	}
	defer f.Close()

	return CountAll(f), nil
}

// CountLines counts newline characters.
func CountLines(r io.Reader) int {
	br := bufio.NewReaderSize(r, bufSize)
	n := 0

	for {
		ch, _, err := br.ReadRune()
		if err != nil {
			break
		}

		if ch == '\n' {
			n++
		}
	}

	return n
}

// CountBytes counts the total number of bytes in the reader.
func CountBytes(r io.Reader) int {
	n, _ := io.Copy(io.Discard, r)
	return int(n)
}

// CountWords counts words using a buffered scanner.
func CountWords(r io.Reader) int {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	n := 0
	for s.Scan() {
		_ = s.Text()
		n++
	}

	return n
}

// CountWordsBuf counts words using bufio.Reader and checking for whitespace.
func CountWordsBuf(r io.Reader) int {
	br := bufio.NewReaderSize(r, bufSize)

	n := 0
	isInsideWord := false

	for {
		ch, _, err := br.ReadRune()
		if err != nil {
			break
		}

		// If the rune is not a space and we're not inside a word, it's the start of a new word
		if !unicode.IsSpace(ch) && !isInsideWord {
			n++
		}

		isInsideWord = !unicode.IsSpace(ch)
	}

	return n
}

// CountWordsRaw is a custom implementation for counting words.
func CountWordsRaw(r io.Reader) int {
	buf := make([]byte, bufSize)
	leftover := make([]byte, 0)

	n := 0
	isInsideWord := false

	// Buffers for reading input and storing leftover bytes.
	for {
		k, err := r.Read(buf)
		if err != nil {
			break
		}

		// Append any leftover bytes from previous reads to the current buffer.
		data := append(leftover, buf[:k]...)

		for len(data) > 0 {
			ch, size := utf8.DecodeRune(data)
			if ch == utf8.RuneError {
				break
			}

			// If the rune is not whitespace and we're not inside a word, it's a new word.
			if !unicode.IsSpace(ch) && !isInsideWord {
				n++
			}

			isInsideWord = !unicode.IsSpace(ch)
			data = data[size:]
		}

		// Store any leftover bytes that didn't form a complete rune.
		leftover = leftover[:0]
		leftover = append(leftover, data...)
	}

	return n
}
