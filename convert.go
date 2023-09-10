package main

import (
    "fmt"
    "strconv"
)

func main(){
    var input string
    fmt.Println("Enter Age")
    fmt.Scanf("%s",&input)
    // you need to use ":=" when you try to use err or _ for multi err and muti assigement
    age,err :=strconv.Atoi(input)
    if err!=nil {
        fmt.Println(err)
    } else{
        fmt.Println(age)
    }
}
