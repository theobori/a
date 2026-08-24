package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kingpin/v2"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/theobori/a/filename"
	"github.com/theobori/a/operation"
)

var (
	app = kingpin.New("a", "My opinionated tool to manage notes.")

	create            = app.Command("create", "Create a new note file.")
	createDirectory   = create.Arg("directory", "Base directory path.").Default("./").String()
	createDisableOpen = create.Flag("disable-open", "Disable editor file opening after the file creation.").Bool()

	edit     = app.Command("edit", "Edit a note file.")
	editPath = edit.Arg("path", "File path.").Required().String()

	rename         = app.Command("rename", "Rename a note file.")
	renamePath     = rename.Arg("path", "File path.").Required().String()
	renameInsecure = rename.Flag("insecure", "Skip the source filename verification.").Short('i').Bool()

	query          = app.Command("query", "Filter files in the base directory using a custom query language.")
	queryDirectory = query.Arg("directory", "Base directory path.").Default("./").String()
	queryQuery     = query.Flag("query", "The query to filter file metadatas.").Default("\"all\"").Short('q').String()
	queryRecursive = query.Flag("recursive", "Recursively walk into the base directory.").Short('r').Bool()

	// Default for the moment
	options = filename.NewOptionsDefault()
)

func Create() error {
	directory, err := filepath.Abs(*createDirectory)
	if err != nil {
		return err
	}

	fm, err := operation.Create(directory, *createDisableOpen, options)
	if err != nil {
		return err
	}

	fmt.Printf("You created the file '%s'.\n", fm.Filename())

	return nil
}

func Edit() error {
	err := operation.Edit(*editPath)
	if err != nil {
		return err
	}

	return nil
}

func Rename() error {
	fm, err := operation.Rename(*renamePath, *renameInsecure, options)
	if err != nil {
		return err
	}

	fmt.Printf("You renamed the file '%s' to '%s.'\n", *renamePath, fm.Filename())

	return nil
}

func Query() error {
	files, err := operation.Query(*queryDirectory, *queryRecursive, *queryQuery, options)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		fmt.Println("There are no matches.")
		return nil
	}

	idx, err := fuzzyfinder.Find(
		files,
		func(i int) string {
			return files[i].Path
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}

			joinedTags := strings.Join(files[i].Metadata.Tags.List(), ", ")

			return fmt.Sprintf("Title: %s\nTags: %s\nExtension: %s",
				files[i].Metadata.Title,
				joinedTags,
				files[i].Metadata.Ext)
		}))
	if err != nil {
		return err
	}

	fmt.Println(files[idx].Path)

	return nil
}

func runCommand() error {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case create.FullCommand():
		return Create()
	case edit.FullCommand():
		return Edit()
	case rename.FullCommand():
		return Rename()
	case query.FullCommand():
		return Query()
	default:
		return fmt.Errorf("invalid command.")
	}
}

func main() {
	err := runCommand()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
