package main

import (
	"log"
	"os"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}

	list, err := os.ReadDir(home)

	if err != nil {

	}
}
