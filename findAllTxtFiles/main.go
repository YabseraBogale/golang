package main

import (
	"bytes"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("locate", "~/*.txt")
	err := cmd.Run()
	if err != nil {
		println(1)
		log.Println(err)
	}
	var out bytes.Buffer
	out = &cmd.Stdout

	println(out)
}
