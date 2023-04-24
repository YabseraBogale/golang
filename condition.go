package main

import "fmt"

func equal(x int, y int) bool{
	return x==y
}

func main() {
	a:= 0
	if a==0 {
	fmt.Println(equal(a,0))
} else{
	fmt.Println("False")
}
}
