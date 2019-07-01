package dir

import (
	"os"

	"github.com/flywithbug/cmd"
)

// Chdir changes the current working directory to the named directory.
// If there is an error, it will be of type *PathError.
func Chdir(path string) error {
	err := os.Chdir(path)
	if err != nil {
		return err
	}
	return nil
}

func Ls() cmd.Status {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("ls")
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}
