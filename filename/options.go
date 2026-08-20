package filename

import (
	"fmt"
	"strings"

	"github.com/theobori/a/internal"
)

var AllowedFileExtsDefault = []string{
	".txt",
	".md",
	".markdown",
	".org",
	".rst",
	".rtf",
	".log",
}

type Options struct {
	partsDelimitor  string
	tagsDelimitor   string
	allowedFileExts *internal.Hashset[string]
}

func (o *Options) PartsDelimitor() string {
	return o.partsDelimitor
}

func (o *Options) TagsDelimitor() string {
	return o.tagsDelimitor
}

func NewOptions(partsDelimitor string, tagsDelimitor string, allowedFileExts *internal.Hashset[string]) (*Options, error) {
	if partsDelimitor == tagsDelimitor {
		return nil, fmt.Errorf("Every delimitor must be different.")
	}
	if strings.Contains(tagsDelimitor, partsDelimitor) {
		return nil, fmt.Errorf("'%s' contains '%s'.", tagsDelimitor, partsDelimitor)
	}

	return &Options{
		partsDelimitor:  partsDelimitor,
		tagsDelimitor:   tagsDelimitor,
		allowedFileExts: allowedFileExts,
	}, nil
}

func NewOptionsDefault() *Options {
	allowedFileExts := internal.NewHashset(AllowedFileExtsDefault)
	options, _ := NewOptions("__", "_", allowedFileExts)

	return options
}
