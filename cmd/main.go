package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	counter "github.com/iamBelugax/wc-cli"
	"github.com/iamBelugax/wc-cli/display"
)

func main() {
	log.SetFlags(0)

	var opts display.Options
	flag.BoolVar(&opts.ShowBytes, "c", false, "Used to toggle whether to show bytes")
	flag.BoolVar(&opts.ShowWords, "w", false, "Used to toggle whether to show word count")
	flag.BoolVar(&opts.ShowLines, "l", false, "Used to toggle whether to show lines count")
	flag.Parse()

	var total counter.Counts
	filenames := flag.Args()
	var hasErrorOccurred bool

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)

	for _, filename := range filenames {
		counts, err := counter.CountFile(filename)
		if err != nil {
			hasErrorOccurred = true
			fmt.Fprintln(os.Stderr, "wc:", err)
			continue
		}

		total.Add(counts)
		counts.Print(tw, opts, filename)
	}

	if len(filenames) == 0 {
		counts := counter.CountAll(os.Stdin)
		counts.Print(tw, opts)
		tw.Flush()
		os.Exit(0)
	}

	if len(filenames) > 1 {
		total.Print(tw, opts, "total")
	}

	tw.Flush()

	if hasErrorOccurred {
		os.Exit(1)
	}
}
