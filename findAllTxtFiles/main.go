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
	ListDir(home, 0)
}

func ListDir(Name string, size float32) {
	count := size
	list, err := os.ReadDir(Name)
	if err != nil {
		log.Println(err)
	}
	for _, i := range list {
		if !i.IsDir() && strings.Contains(i.Name(), ".txt") {
			info, err := os.ReadFile(Name + i.Name())
			if err != nil {

			}
			fmt.Println(info)
			count += 1
		} else if i.IsDir() == true {
			ListDir(Name+"/"+i.Name(), count)
		}
	}
}
