package main

import "fmt"

type car struct{
	
	carWheel int
	carName string
	
	
	}

func main() {
	
	mycar:=car{4,"volovo"}
	
	var arr [10] int
	for i:=0;i<cap(arr);i++ {
		arr[i]=i
	}
	fmt.Println(arr)
	fmt.Println(mycar)
	fmt.Println("My car Name:%s",mycar.carName)
}
