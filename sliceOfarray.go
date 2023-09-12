package main

import "fmt"

func main() {
	q:=[3]int{1,2,3}
	fmt.Println(q)
	q=append(q,4)
	fmt.Println(q)
	//will not work b/c it similar to make(int,size)	
	
}
