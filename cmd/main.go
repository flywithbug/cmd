package main

import (
	"fmt"

	"github.com/flywithbug/cmd/git"
)

func main() {
	// Create Cmd, buffered output
	//envCmd := cmd.NewCmd("/bin/bash", "-c", "git clone ssh://git@gitlab.hellobike.cn:10022/MopedApp/JYBaseModel.git")
	//
	//// Run and wait for Cmd to return Status
	//status := <-envCmd.Start()
	//
	//if status.Error != nil {
	//	fmt.Println(status.Error)
	//}
	////// Print each line of STDOUT from Cmd
	//for _, line := range status.Stdout {
	//	fmt.Fprintf(os.Stdout, line)
	//	time.Sleep(time.Second * 1)
	//}
	sta := git.Clone("ssh://git@gitlab.hellobike.cn:10022/MopedApp/JYBaseModel.git")
	if sta.Error != nil {
		fmt.Println(sta.Error)
	}

	fmt.Println(sta.Stdout)

}
