package main

import (
	"archive/zip"
	"fmt"
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
	data, err := os.Create("data.zip")
	if err != nil {

	}
	defer data.Close()
	ListDir(home, data)
}

func ListDir(Name string, data *os.File) {

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
			zipwriter := zip.NewWriter(data)
			defer zipwriter.Close()
			fsInfo, err := fs.Stat()
			if err != nil {
				log.Fatalln()
			}
			header, err := zip.FileInfoHeader(fsInfo)
			if err != nil {
				log.Fatalln()
			}
			header.Name = i.Name()
			header.Method = zip.Deflate
			writer, err := zipwriter.CreateHeader(header)
			if err != nil {
				log.Fatalln(err)
			}
			io.Copy(writer, fs)
			fmt.Println(i.Name())

		} else if i.IsDir() == true {
			ListDir(Name+"/"+i.Name(), data)
		}
	}

}
