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
			err := os.Chdir(Name)
			if err != nil {

			}
			fs, err := os.Open(i.Name())
			if err != nil {

			}
			defer fs.Close()
			info, err := fs.Stat()
			if err != nil {

			}
			fmt.Println(Name+"/"+i.Name(), info.Name(), info.Size())
		} else if i.IsDir() == true {
			ListDir(Name+"/"+i.Name(), count)
		}
	}
}
