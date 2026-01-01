package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/iamBelugax/wc-cli/counter"
)

func main() {
	log.SetFlags(0)

	var displayOpts counter.DisplayOpts
	flag.BoolVar(&displayOpts.ShowBytes, "c", false, "Used to toggle whether to show bytes")
	flag.BoolVar(&displayOpts.ShowWords, "w", false, "Used to toggle whether to show word count")
	flag.BoolVar(&displayOpts.ShowLines, "l", false, "Used to toggle whether to show lines count")
	flag.Parse()

	var total counter.Counts
	filenames := flag.Args()
	var hasErrorOccurred bool

	for _, filename := range filenames {
		counts, err := counter.CountFile(filename)
		if err != nil {
			hasErrorOccurred = true
			fmt.Fprintln(os.Stderr, "wc:", err)
			continue
		}

		total.Add(counts)
		counts.Print(os.Stdout, displayOpts, filename)
	}

	if len(filenames) == 0 {
		counts := counter.CountAll(os.Stdin)
		counts.Print(os.Stdout, displayOpts)
		os.Exit(0)
	}

	if len(filenames) > 1 {
		total.Print(os.Stdout, displayOpts, "total")
	}

	if hasErrorOccurred {
		os.Exit(1)
	}
}
