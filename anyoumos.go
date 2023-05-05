package main

import "fmt"

func main()  {
	sum:=func(x int,y int) int{
		return x+y
	}
	func(x int,y int){
		fmt.Println(x*y)
	}(1,2)
	fmt.Println(sum(2,2))
}