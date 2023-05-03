package main

import "fmt"

func test(number int) {

	for number>0 {
		number-=1
		fmt.Println("Number:",number)
	}

}

func main() {

	
	 Prime(20)
	 go test(10)
}