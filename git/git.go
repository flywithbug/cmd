package git

import (
	"fmt"

	"strings"

	"github.com/flywithbug/cmd"
)

func Init(path string) cmd.Status {
	// Create Cmd, buffered output

	envCmd := cmd.NewCmd("/bin/bash", "-c", "git init "+path)
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}

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
	envCmd := cmd.NewCmd("/bin/bash", "-c", "git branch "+branch)
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}

func DelBranch(branch string) cmd.Status {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("/bin/bash", "-c", "git branch -d "+branch)
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}

func CheckOut(branch string) cmd.Status {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("/bin/bash", "-c", "git checkout "+branch)
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}

func Status() cmd.Status {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("/bin/bash", "-c", "git status -s")
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}

func Add(args ...string) cmd.Status {
	addS := "."
	if len(args) != 0 {
		addS = strings.Join(args, " ")
	}
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("/bin/bash", "-c", "git add "+addS)
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}

func Commit(note string) cmd.Status {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("/bin/bash", "-c", "git commit -m "+note)
	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()
	return status
}
