package counter_test

import (
	"bytes"
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
		{"EmptyInput", "", 0},
		{"OnlySpaces", "     ", 0},
		{"OnlyTabs", "\t\t\t", 0},
		{"OnlyNewlines", "\n\n\n", 0},
		{"SingleWord", "hello", 1},
		{"MultipleSpaces", "hello   world", 2},
		{"MixedWhitespace", "hello\tworld\nagain", 3},
		{"Punctuation", "hello, world!", 2},
		{"UnicodeWords", "こんにちは 世界", 2},
		{"MixedScripts", "hello 世界 today", 3},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			reader := strings.NewReader(testCase.input)
			actualCount := counter.CountWords(reader)

			if actualCount != testCase.wants {
				t.Fatalf(
					"expected %d words, got %d",
					testCase.wants,
					actualCount,
				)
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
		{"EmptyInput", "", 0},
		{"OnlySpaces", "   ", 0},
		{"OnlyNewlines", "\n\n", 0},
		{"OnlyTabs", "\t\t", 0},
		{"SingleWord", "word", 1},
		{"MixedWhitespace", "one \t two \n three", 3},
		{"UnicodeWords", "空手 오늘", 2},
		{"MixedUnicodeAndAscii", "Go 언어 is fun", 4},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			reader := strings.NewReader(testCase.input)
			actualCount := counter.CountWordsBuf(reader)

			if actualCount != testCase.wants {
				t.Fatalf(
					"expected %d words, got %d",
					testCase.wants,
					actualCount,
				)
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
		{"EmptyInput", "", 0},
		{"OnlySpaces", "   ", 0},
		{"OnlyNewlines", "\n\n", 0},
		{"SingleWord", "hello", 1},
		{"WordsSeparatedByNewline", "hello\nworld", 2},
		{"UnicodeWords", "Мечта сегодня", 2},
		{"EmojiWords", "hello 👋 world", 3},
		{"RepeatedUnicode", strings.Repeat("界 ", 100), 100},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			reader := strings.NewReader(testCase.input)
			actualCount := counter.CountWordsRaw(reader)

			if actualCount != testCase.wants {
				t.Fatalf(
					"expected %d words, got %d",
					testCase.wants,
					actualCount,
				)
			}
		})
	}
}

func TestCountLines(t *testing.T) {
	testCases := []struct {
		input string
		wants int
	}{
		{"", 0},
		{"\n", 1},
		{"\n\n\n", 3},
		{"NoNewline", 0},
		{"First\nSecond", 1},
		{"First\nSecond\n", 2},
		{"\nStart\nMiddle\nEnd\n", 4},
	}

	for _, testCase := range testCases {
		reader := strings.NewReader(testCase.input)
		actualLineCount := counter.CountLines(reader)

		if actualLineCount != testCase.wants {
			t.Fatalf(
				"expected %d lines, got %d",
				testCase.wants,
				actualLineCount,
			)
		}
	}
}

func TestCountBytes(t *testing.T) {
	testCases := []struct {
		input string
		wants int
	}{
		{"", 0},
		{"     ", 5},
		{"abc", 3},
		{"hello\n", 6},
		{"こんにちは", len([]byte("こんにちは"))},
		{"emoji 🚀", len([]byte("emoji 🚀"))},
	}

	for _, testCase := range testCases {
		reader := strings.NewReader(testCase.input)
		actualByteCount := counter.CountBytes(reader)

		if actualByteCount != testCase.wants {
			t.Fatalf(
				"expected %d bytes, got %d",
				testCase.wants,
				actualByteCount,
			)
		}
	}
}

func TestCountAll(t *testing.T) {
	testCases := []struct {
		input string
		wants counter.Counts
	}{
		{"", counter.Counts{Words: 0, Lines: 0, Bytes: 0}},
		{"   ", counter.Counts{Words: 0, Lines: 0, Bytes: 3}},
		{"hello", counter.Counts{Words: 1, Lines: 0, Bytes: 5}},
		{"hello\nworld\n", counter.Counts{Words: 2, Lines: 2, Bytes: 12}},
		{
			"こんにちは 世界\n",
			counter.Counts{
				Words: 2,
				Lines: 1,
				Bytes: len([]byte("こんにちは 世界\n")),
			},
		},
	}

	for _, testCase := range testCases {
		reader := strings.NewReader(testCase.input)
		actualCounts := counter.CountAll(reader)

		if actualCounts != testCase.wants {
			t.Fatalf(
				"expected %+v, got %+v",
				testCase.wants,
				actualCounts,
			)
		}
	}
}

func TestCountsAdd(t *testing.T) {
	testCases := []struct {
		initialCounts  counter.Counts
		additionalData counter.Counts
		wants          counter.Counts
	}{
		{
			counter.Counts{},
			counter.Counts{},
			counter.Counts{},
		},
		{
			counter.Counts{Words: 1, Lines: 2, Bytes: 3},
			counter.Counts{Words: 4, Lines: 5, Bytes: 6},
			counter.Counts{Words: 5, Lines: 7, Bytes: 9},
		},
		{
			counter.Counts{Words: 0, Lines: 1, Bytes: 0},
			counter.Counts{Words: 5, Lines: 0, Bytes: 5},
			counter.Counts{Words: 5, Lines: 1, Bytes: 5},
		},
	}

	for _, testCase := range testCases {
		currentCounts := testCase.initialCounts
		currentCounts.Add(testCase.additionalData)

		if currentCounts.Bytes != testCase.wants.Bytes ||
			currentCounts.Lines != testCase.wants.Lines ||
			currentCounts.Words != testCase.wants.Words {
			t.Fatalf(
				"expected %+v, got %+v",
				testCase.wants,
				currentCounts,
			)
		}
	}
}

func TestCountsPrint(t *testing.T) {
	testCases := []struct {
		name   string
		counts counter.Counts
		suffix []string
		wants  string
	}{
		{
			"NoSuffix",
			counter.Counts{Words: 1, Lines: 2, Bytes: 3},
			nil,
			"2 1 3\n",
		},
		{
			"WithFilename",
			counter.Counts{Words: 1, Lines: 2, Bytes: 3},
			[]string{"file.txt"},
			"2 1 3 file.txt\n",
		},
		{
			"ZeroValues",
			counter.Counts{},
			[]string{"empty"},
			"0 0 0 empty\n",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var buffer bytes.Buffer
			testCase.counts.Print(&buffer, testCase.suffix...)

			if buffer.String() != testCase.wants {
				t.Fatalf(
					"expected %q, got %q",
					testCase.wants,
					buffer.String(),
				)
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
