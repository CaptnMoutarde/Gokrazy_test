package main

import (
	"fmt"
	"os"
	"os/exec"
)


func main() {
	cmd 		:= exec.Command("echo", os.Getenv("GOPATH"))
	cmd.Stdout	 = os.Stdout
	cmd.Stderr	 = os.Stderr
	err 			:= cmd.Run()
	if err != nil {
		fmt.Println("NIIIIIIIIIIIIIIIIIIIIIIIQUE !!!!!")
	}
}