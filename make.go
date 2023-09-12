package main

import "fmt"

func main() {
	q := make([]int,5)
	q[0]=1
	q[1]=2
	q[2]=3
	q[3]=4
	q[4]=5
	fmt.Println(q)
	// can't append new data to array that is made with make b/c it have a limited right
	// it similar with q:= [5] int{1,2,3,4,5}
	q[5]=6
	fmt.Println(q)
}
