package display

type Options struct {
	showWords bool
	showLines bool
	showBytes bool
}

func NewOptions(lines, words, bytes bool) Options {
	return Options{
		showWords: words,
		showLines: lines,
		showBytes: bytes,
	}
}

func (d *Options) ShowWords() bool {
	if !d.showBytes && !d.showLines && !d.showWords {
		return true
	}
	return d.showWords
}

func (d *Options) ShowLines() bool {
	if !d.showBytes && !d.showLines && !d.showWords {
		return true
	}
	return d.showLines
}

func (d *Options) ShowBytes() bool {
	if !d.showBytes && !d.showLines && !d.showWords {
		return true
	}
	return d.showBytes
}
