package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("india_Database_Business.xlsx")
	if err != nil {
		log.Fatalln(err)
	}

	for _, i := range file {
		fmt.Println(i)

	}
}
