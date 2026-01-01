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

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(filenames))

	var totals counter.Counts
	var hasErrorOccurred bool

	for _, filename := range filenames {
		go func() {
			defer wg.Done()

			counts, err := counter.CountFile(filename)
			if err != nil {
				mu.Lock()
				defer mu.Unlock()

				hasErrorOccurred = true
				fmt.Fprintln(os.Stderr, "wc:", err)
				return
			}

			mu.Lock()
			defer mu.Unlock()

			totals.Add(counts)
			counts.Print(tw, opts, filename)
		}()
	}

	wg.Wait()

	totals.Print(tw, opts, "total")
	tw.Flush()

	if hasErrorOccurred {
		os.Exit(1)
	}
}
