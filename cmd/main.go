package main

import (
	"fmt"

	"github.com/flywithbug/cmd/dir"
	"github.com/flywithbug/file_manager"
)

func main() {

	filePath := "./a/a"
	//sta := git.Clone("ssh://git@gitlab.hellobike.cn:10022/MopedApp/JYBaseModel.git", filePath)
	//if sta.Error != nil {
	//	fmt.Println(sta.Error)
	//}
	//
	//fmt.Println(sta.ToJson())
	//
	err := dir.Chdir(filePath)
	if err != nil {
		fmt.Printf("ChdirError : %s\\n", err)
		return
	}
	////
	//sta = git.Branch("dev")
	//if sta.Error != nil {
	//	fmt.Println(sta.Error)
	//}
	//
	//fmt.Println(sta.ToJson())
	//
	//sta = git.Push()
	//if sta.Error != nil {
	//	fmt.Println(sta.Error)
	//}
	//
	//fmt.Println(sta.ToJson())
	//
	sta := dir.Ls()
	if sta.Error != nil {
		fmt.Println(sta.Error)
	}

	file_manager.AppendFileString("README.md", "test\n")

}
