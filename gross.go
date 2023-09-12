package main

import "fmt"

func sum(num ... int) (total int){
    for _,n:=range num{
        total+=n
    }
    return
}

func main(){
    num:=sum(1,2,3)
    fmt.Println("sum of 1 + 2 + 3 = ",num)
}
