package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
}
