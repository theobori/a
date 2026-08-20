package operation

import (
	"errors"
	"fmt"
	"os"

	"github.com/theobori/a/filename"
)

func Create(directory string, disableOpen bool, options *filename.Options) (*filename.Metadata, error) {
	fm, err := filename.NewMetadataFromStdin(options)
	if err != nil {
		return nil, err
	}

	path, _ := fm.AbsJoin(directory)

	_, err = os.Stat(path)
	if !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("The file '%s' already exist.", path)
	}
	_, err = os.Create(path)
	if err != nil {
		return nil, err
	}

	if !disableOpen {
		err = Edit(path)
		if err != nil {
			return nil, err
		}
	}

	return fm, nil
}
