package echo

import "github.com/flywithbug/cmd"

func Echo(note, filePath string, cover bool) cmd.Status {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("/bin/bash", "-c", "git branch -d ")
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}
