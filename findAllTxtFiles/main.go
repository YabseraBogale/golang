package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}

	list, err := os.ReadDir(home)

	if err != nil {

	}
	for _, i := range list {
		if !i.IsDir() && strings.Contains(i.Name(), ".txt") {
			fmt.Println(i.Info())
		} else if i.IsDir() {
			fmt.Println(i.Name())
		}
	}
}
