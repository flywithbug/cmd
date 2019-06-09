package main

import (
	"fmt"

	"github.com/flywithbug/cmd"
)

func main() {
	b, _ := cmd.Exec("ls")
	fmt.Println(string(b))
}
