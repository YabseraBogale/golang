package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Create("data.gz")
	if err != nil {

	}

	ListDir(home, file)
}

func ListDir(Name string, file *os.File) {

	list, err := os.ReadDir(Name)
	if err != nil {
		log.Println(err)
	}
	for _, i := range list {
		if !i.IsDir() && strings.Contains(i.Name(), ".txt") && i.Type().IsRegular() {
			err := os.Chdir(Name)
			if err != nil {

			}
			fs, err := os.Open(i.Name())
			if err != nil {

			}
			defer fs.Close()
			// Compress or copy ?
			dist_dir, err := os.Create()
		} else if i.IsDir() == true {
			ListDir(Name + "/" + i.Name())
		}
	}
}
