package internal

import (
	"github.com/spf13/afero"
)

// Organize is the main function that organizes files and directories
// based on a given specification.
func organize() error {
	var AppFs = afero.NewOsFs()
	AppFs.Create("/tmp/test.txt")
	return nil
}
