package main

import "fmt"

func hello()string {
	return "hello world"
}


func sub(x int, y int) int{
	return x-y
}

func main() {
	fmt.Println(hello())
	//fmt.Println(sub(1,1))
}
