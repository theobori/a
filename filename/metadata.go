package filename

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/theobori/a/internal"
)

type Metadata struct {
	Title   string
	Tags    *internal.Hashset[string]
	Ext     string
	options *Options
}

func (m *Metadata) Filename() string {
	ans := ""

	ans += m.Title
	if m.Tags.Len() > 0 {
		ans += m.options.partsDelimitor

		tagsList := m.Tags.List()
		slices.Sort(tagsList)

		ans += strings.Join(tagsList, m.options.tagsDelimitor)
	}

	if len(m.Ext) > 0 {
		ans += m.Ext
	}

	return ans
}

func (m *Metadata) String() string {
	return m.Filename()
}

func (m *Metadata) AbsJoin(elem ...string) (string, error) {
	// The nested join is a little hacky but needed because of elem...
	return filepath.Abs(filepath.Join(filepath.Join(elem...), m.Filename()))
}

func NewMetadataWithOptions(title string, tags []string, ext string, options *Options) (*Metadata, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, fmt.Errorf("Empty title.")
	}
	if strings.Contains(title, options.partsDelimitor) {
		return nil, fmt.Errorf("'%s' is not allowed in the title.", options.partsDelimitor)
	}
	if ext != "" && !options.IsFileExtAllowed(ext) {
		return nil, fmt.Errorf("The file extension name '%s' is not allowed.", ext)
	}

	tagsHs := internal.NewEmptyHashset[string]()
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			return nil, fmt.Errorf("Empty of whitespace tags are not allowed.")
		}
		if strings.Contains(tag, options.partsDelimitor) {
			return nil, fmt.Errorf(
				"The tag '%s' contains the illegal string '%s'.",
				tag,
				options.partsDelimitor,
			)
		}
		tagsHs.Set(tag)
	}

	ext = strings.TrimSpace(ext)

	return &Metadata{
		Title:   title,
		Tags:    tagsHs,
		Ext:     ext,
		options: options,
	}, nil
}

func NewMetadata(title string, tags []string, ext string) (*Metadata, error) {
	options := NewOptionsDefault()

	return NewMetadataWithOptions(title, tags, ext, options)
}

func NewMetadataFromPath(path string, options *Options) (*Metadata, error) {
	// doing this because filepath.Base returns '.' for empty path
	if path == "" {
		return nil, fmt.Errorf("The path should not be empty.")
	}

	filebase := filepath.Base(path)
	ext := filepath.Ext(path)
	filename := strings.TrimSuffix(filebase, ext)

	parts := strings.Split(filename, options.partsDelimitor)

	var tags []string
	switch len(parts) {
	case 1:
		tags = []string{}
	case 2:
		for tag := range strings.SplitSeq(parts[1], options.tagsDelimitor) {
			tags = append(tags, tag)
		}
	default:
		return nil, fmt.Errorf("Invalid parts amount in the filename '%s'.", filename)
	}

	title := parts[0]

	m, err := NewMetadataWithOptions(title, tags, ext, options)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func NewMetadataFromStdin(options *Options) (*Metadata, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter title: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	fmt.Print("Enter tags splitted by ',' (press enter for none): ")
	tagsExpr, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	fmt.Print("Enter the file extension (press enter for none): ")
	ext, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	tags := strings.Split(tagsExpr, ",")
	// Prevention filter
	filteredTags := []string{}
	for _, tag := range tags {
		tag := strings.TrimSpace(tag)
		if tag == "" {
			continue
		}

		filteredTags = append(filteredTags, tag)
	}

	f, err := NewMetadataWithOptions(title, filteredTags, ext, options)
	if err != nil {
		return nil, err
	}

	return f, nil
}
