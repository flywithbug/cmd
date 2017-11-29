package main

import (
	"github.com/flywithbug/file"
	"fmt"
	"regexp"
)

func main()  {
	path ,_ := file.CurrentDir()
	fmt.Println(path)
	str := "baabbb        a,a"
	str = regexp.MustCompile(`\s{2,}`).ReplaceAllString(str, " ")
	fmt.Println(str)

}