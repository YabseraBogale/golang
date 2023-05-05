package main
import "fmt"

func main() {
	
	result:=make(chan int, 1)
	go func(x int,y int,res chan int){
		sum:=x+y
		res<-sum
	}(1,2,result)
	value:=<-result
	fmt.Println(value)
	
}