package operation

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/theobori/a/filename"
	"github.com/theobori/a/ql"
)

func filesAtDepthOne(directory string) ([]string, error) {
	ans := []string{}

	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		path := filepath.Join(directory, entry.Name())
		ans = append(ans, path)
	}

	return ans, nil
}

func filesRecursive(directory string) ([]string, error) {
	ans := []string{}

	filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		myPath := filepath.Join(directory, info.Name())
		ans = append(ans, myPath)

		return nil
	})

	return ans, nil
}

type File struct {
	// Should always be absolute
	Path     string
	Metadata *filename.Metadata
}

func Query(directory string, query string, options *filename.Options) ([]File, error) {
	directoryAbs, err := filepath.Abs(directory)
	if err != nil {
		return nil, err
	}

	filePaths, err := filesAtDepthOne(directory)
	if err != nil {
		return nil, err
	}

	chunk, err := ql.NewChunkFromSource(query)
	if err != nil {
		return nil, err
	}

	ans := []File{}
	for _, filePath := range filePaths {
		metadata, err := filename.NewMetadataFromPath(filePath, options)
		if err != nil {
			return nil, err
		}

		ctx := ql.Context{Metadata: metadata}

		vm := ql.NewVM(chunk)
		match, err := vm.Interpret(&ctx)
		if err != nil {
			return nil, err
		}

		if match {
			filePathAbs := filepath.Join(directoryAbs, filePath)
			ans = append(ans, File{
				Path:     filePathAbs,
				Metadata: metadata,
			})
		}
	}

	return ans, nil
}
