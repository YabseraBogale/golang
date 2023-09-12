package main

import "fmt"

func sum(num ... int) int{
    // cann't use var total int:=0 must this
    var total int=0
    for _,n:= range num{
        total+=n
    }
    return total
}


func main(){
    num:=sum(1,2,3)
    fmt.Println("sum of 1 + 2 + 3 =",num)
}
