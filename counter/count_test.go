package counter_test

import (
	"strings"
	"testing"

	"github.com/iamBelugax/wc-cli/counter"
)

const testText = `I was walking through the park, feeling the cool breeze, when I noticed a small group of friends chatting. One of them was holding a book titled "Мечта" (Russian for "Dream"). The girl beside him smiled and said, "오늘 날씨 정말 좋아요!" (Korean for "The weather is really nice today!"). Nearby, a boy was practicing his "空手" (Japanese for "karate") moves, his movements smooth and precise. It felt like the perfect day, where different cultures and languages came together in harmony.
`

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 Words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "Empty Input",
			input: "",
			wants: 0,
		},
		{
			name:  "Single Space",
			input: " ",
			wants: 0,
		},
		{
			name:  "Single New Line",
			input: "one two three\nfour five",
			wants: 5,
		},
		{
			name:  "Multi New Lines",
			input: "one two three \n four \n five",
			wants: 5,
		},
		{
			name:  "Multi Spaces",
			input: "This is a sentence.  This is another one.",
			wants: 8,
		},
		{
			name:  "Prefixed Multi Spaces",
			input: "    This is a sentence.  This is another one.",
			wants: 8,
		},
		{
			name:  "Suffixed Multi Spaces",
			input: "This is a sentence.  This is another one.    ",
			wants: 8,
		},
		{
			name:  "Tab Character",
			input: "This is\ta sentence.\tThis is \tanother one.",
			wants: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			res := counter.CountWords(reader)
			if res != tc.wants {
				t.Logf("expected %d, got %d", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func TestCountWordsBuf(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 Words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "Empty Input",
			input: "",
			wants: 0,
		},
		{
			name:  "Single Space",
			input: " ",
			wants: 0,
		},
		{
			name:  "Single New Line",
			input: "one two three\nfour five",
			wants: 5,
		},
		{
			name:  "Multi New Lines",
			input: "one two three \n four \n five",
			wants: 5,
		},
		{
			name:  "Multi Spaces",
			input: "This is a sentence.  This is another one.",
			wants: 8,
		},
		{
			name:  "Prefixed Multi Spaces",
			input: "    This is a sentence.  This is another one.",
			wants: 8,
		},
		{
			name:  "Suffixed Multi Spaces",
			input: "This is a sentence.  This is another one.    ",
			wants: 8,
		},
		{
			name:  "Tab Character",
			input: "This is\ta sentence.\tThis is \tanother one.",
			wants: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			res := counter.CountWordsBuf(reader)
			if res != tc.wants {
				t.Logf("expected %d, got %d", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func TestCountWordsRaw(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 Words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "Empty Input",
			input: "",
			wants: 0,
		},
		{
			name:  "Single Space",
			input: " ",
			wants: 0,
		},
		{
			name:  "Single New Line",
			input: "one two three\nfour five",
			wants: 5,
		},
		{
			name:  "Multi New Lines",
			input: "one two three \n four \n five",
			wants: 5,
		},
		{
			name:  "Multi Spaces",
			input: "This is a sentence.  This is another one.",
			wants: 8,
		},
		{
			name:  "Prefixed Multi Spaces",
			input: "    This is a sentence.  This is another one.",
			wants: 8,
		},
		{
			name:  "Suffixed Multi Spaces",
			input: "This is a sentence.  This is another one.    ",
			wants: 8,
		},
		{
			name:  "Tab Character",
			input: "This is\ta sentence.\tThis is \tanother one.",
			wants: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			res := counter.CountWordsRaw(reader)
			if res != tc.wants {
				t.Logf("expected %d, got %d", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Empty Input",
			input: "",
			wants: 0,
		},
		{
			name:  "Single New Line",
			input: "This is a sentence\n",
			wants: 1,
		},
		{
			name:  "Multi New Lines",
			input: "This is a sentence\n\n\n",
			wants: 3,
		},
		{
			name:  "No New Line",
			input: "This is a sentence",
			wants: 0,
		},
		{
			name:  "No New Line at the end",
			input: "This is a sentence\n. This is another sentence.",
			wants: 1,
		},
		{
			name:  "New Line at the beginning",
			input: "\nThis is a sentence\n. This is another sentence.",
			wants: 2,
		},
		{
			name:  "Multi New lines",
			input: "\n\n\n\n\n.",
			wants: 5,
		},
		{
			name:  "Multi Word New lines",
			input: "one\ntwo\nthree\nfour\nfive\n",
			wants: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			res := counter.CountLines(reader)
			if res != tc.wants {
				t.Logf("expected %d, got %d", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func TestCountBytes(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "Empty Input",
			input: "",
			wants: 0,
		},
		{
			name:  "5 Empty Spaces",
			input: "     ",
			wants: 5,
		},
		{
			name:  "Five Words",
			input: "one two three four five",
			wants: 23,
		},
		{
			name:  "Five Words With New Lines",
			input: "one two three \nfour five \n\n\n",
			wants: 28,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			res := counter.CountBytes(reader)
			if res != tc.wants {
				t.Logf("expected %d, got %d", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func TestCountAll(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		wants counter.Counts
	}{
		{
			name:  "Empty Input",
			input: "",
			wants: counter.Counts{Words: 0, Lines: 0, Bytes: 0},
		},
		{
			name:  "5 Spaces",
			input: "     ",
			wants: counter.Counts{Words: 0, Lines: 0, Bytes: 5},
		},
		{
			name:  "5 Spaces and Single New Line",
			input: "     \n",
			wants: counter.Counts{Words: 0, Lines: 1, Bytes: 6},
		},
		{
			name:  "Five Words With New Lines",
			input: "one two three \nfour five \n\n\n",
			wants: counter.Counts{Words: 5, Lines: 4, Bytes: 28},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			res := counter.CountAll(reader)

			if res.Bytes != tc.wants.Bytes {
				t.Logf("expected %d bytes, got %d bytes", tc.wants, res)
				t.Fail()
			}

			if res.Lines != tc.wants.Lines {
				t.Logf("expected %d lines, got %d lines", tc.wants, res)
				t.Fail()
			}

			if res.Words != tc.wants.Words {
				t.Logf("expected %d words, got %d words", tc.wants, res)
				t.Fail()
			}
		})
	}
}

func BenchmarkCountWords(b *testing.B) {
	data := strings.Repeat(testText, 10000)
	for b.Loop() {
		reader := strings.NewReader(data)
		_ = counter.CountWords(reader)
	}
}

func BenchmarkCountWordsBuf(b *testing.B) {
	data := strings.Repeat(testText, 10000)
	for b.Loop() {
		reader := strings.NewReader(data)
		_ = counter.CountWordsBuf(reader)
	}
}

func BenchmarkCountWordsRaw(b *testing.B) {
	data := strings.Repeat(testText, 10000)
	for b.Loop() {
		reader := strings.NewReader(data)
		_ = counter.CountWordsRaw(reader)
	}
}

func BenchmarkCountLines(b *testing.B) {
	data := strings.Repeat(testText, 10000)
	for b.Loop() {
		reader := strings.NewReader(data)
		_ = counter.CountLines(reader)
	}
}

func BenchmarkCountBytes(b *testing.B) {
	data := strings.Repeat(testText, 10000)
	for b.Loop() {
		reader := strings.NewReader(data)
		_ = counter.CountBytes(reader)
	}
}

func BenchmarkCountAll(b *testing.B) {
	data := strings.Repeat(testText, 10000)
	for b.Loop() {
		reader := strings.NewReader(data)
		_ = counter.CountAll(reader)
	}
}
