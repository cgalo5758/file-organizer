package internal

import (
	"path"

	"github.com/spf13/afero"
)

// Organize is the main function that organizes files and directories
// based on a given specification.
// Takes a path to a directory to organize and a path to a specification file.
// Returns an error if the directory or specification file cannot be found.
func Organize(dir string, spec string) error {
	var AppFs = afero.NewOsFs()
	var specPath = path.Join(dir, spec)
	var specFile, err = AppFs.Open(specPath)
	if err != nil {
		return err
	}
	defer specFile.Close()
	return nil
}
