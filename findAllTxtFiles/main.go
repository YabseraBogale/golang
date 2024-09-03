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
	ListDir(home)
}

func ListDir(Name string) {

	list, err := os.ReadDir(Name)
	if err != nil {
		log.Println(err)
	}
	for _, i := range list {
		if !i.IsDir() && strings.Contains(i.Name(), ".txt") {
			fmt.Println(Name + i.Name())
		} else if i.IsDir() == true {
			ListDir(Name + "/" + i.Name())
		}
	}
}
