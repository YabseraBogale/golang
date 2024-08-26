package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")
	err := cmd.Run()
	if err != nil {
		fmt.Println("here 1", err)
	}
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("here 2", err)
	}
	fmt.Println(out)
}
