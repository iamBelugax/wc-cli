package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"text/tabwriter"

	counter "github.com/iamBelugax/wc-cli"
	"github.com/iamBelugax/wc-cli/display"
)

type Result struct {
	filename string
	counts   counter.Counts
}

func main() {
	showBytes := flag.Bool("c", false, "Used to toggle whether to show bytes")
	showWords := flag.Bool("w", false, "Used to toggle whether to show word count")
	showLines := flag.Bool("l", false, "Used to toggle whether to show lines count")
	flag.Parse()

	opts := display.NewOptions(*showLines, *showWords, *showBytes)
	log.SetFlags(0)

	filenames := flag.Args()
	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)

	if len(filenames) == 0 {
		counts := counter.CountAll(os.Stdin)
		counts.Print(tw, opts)
		tw.Flush()
		os.Exit(0)
	}

	var wg sync.WaitGroup
	wg.Add(len(filenames))

	var totals counter.Counts
	var hasErrorOccurred bool
	resultCh := make(chan Result, len(filenames))

	for _, filename := range filenames {
		go func() {
			defer wg.Done()

			counts, err := counter.CountFile(filename)
			if err != nil {
				hasErrorOccurred = true
				fmt.Fprintln(os.Stderr, "wc:", err)
				return
			}

			resultCh <- Result{counts: counts, filename: filename}
		}()
	}

	wg.Wait()
	close(resultCh)

	for res := range resultCh {
		totals.Add(res.counts)
		res.counts.Print(tw, opts, res.filename)
	}

	totals.Print(tw, opts, "total")
	tw.Flush()

	if hasErrorOccurred {
		os.Exit(1)
	}
}
