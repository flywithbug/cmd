package git

import "github.com/flywithbug/cmd"

func Clone(url string) cmd.Status {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("/bin/bash", "-c", "git clone "+url)
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

func Init(path string) cmd.Status {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("/bin/bash", "-c", "git init ", path)
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}
