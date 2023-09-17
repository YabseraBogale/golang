package main

import (
    "fmt"
    "math/rand"
)


func main(){

    arr1:=[]int{}
    arr2:=[]int{}
    for i:=0;i<6;i++{
        arr1=append(arr1,rand.Intn(6))
        arr2=append(arr2,rand.Intn(6))
    }
    fmt.Println(arr1)
    fmt.Println(arr2)
    fmt.Println(arr2==arr1)

}
