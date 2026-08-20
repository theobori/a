package operation

import (
	"os"
	"os/exec"

	"github.com/theobori/a/internal"
)

func Edit(path string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = internal.DefaultEditor
	}

	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
