package operation

import (
	"os"
	"path/filepath"

	"github.com/theobori/a/filename"
)

func Rename(oldpath string, insecure bool, options *filename.Options) (*filename.Metadata, error) {
	if !insecure {
		_, err := filename.NewMetadataFromPath(oldpath, options)
		if err != nil {
			return nil, err
		}
	}

	fm, err := filename.NewMetadataFromStdin(options)
	if err != nil {
		return nil, err
	}

	directory := filepath.Dir(oldpath)
	newpath := filepath.Join(directory, fm.Filename())

	err = os.Rename(oldpath, newpath)
	if err != nil {
		return nil, err
	}

	return fm, nil
}
