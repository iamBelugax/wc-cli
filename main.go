package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iamBelugax/wc-cli/counter"
)

func main() {
	log.SetFlags(0)

	var total counter.Counts
	filenames := os.Args[1:]
	var hasErrorOccurred bool

	for _, filename := range filenames {
		counts, err := counter.CountFile(filename)
		if err != nil {
			hasErrorOccurred = true
			fmt.Fprintln(os.Stderr, "wc:", err)
			continue
		}

		total.Add(counts)
		counts.Print(os.Stdout, filename)
	}

	if len(filenames) == 0 {
		counter.CountAll(os.Stdin).Print(os.Stdout)
		os.Exit(0)
	}

	if len(filenames) > 1 {
		total.Print(os.Stdout, "total")
	}

	if hasErrorOccurred {
		os.Exit(1)
	}
}
