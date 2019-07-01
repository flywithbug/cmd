package git

import (
	"fmt"

	"github.com/flywithbug/cmd"
)

func Clone(url, name string) cmd.Status {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("/bin/bash", "-c", fmt.Sprintf("git clone %s %s", url, name))
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}

func Pull() cmd.Status {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("/bin/bash", "-c", "git pull")
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}

func Push() cmd.Status {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("/bin/bash", "-c", "git push")
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}

func PushF() cmd.Status {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("/bin/bash", "-c", "git push -f")
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}

func Branch(branch string) cmd.Status {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("/bin/bash", "-c", "git branch ", branch)
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}

func Init(path string) cmd.Status {
	// Create Cmd, buffered output

	envCmd := cmd.NewCmd("/bin/bash", "-c", "git init ", path)
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}
