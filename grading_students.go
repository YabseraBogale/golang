package main

import (
    "fmt"
    "strconv"
)

func main(){

    var input string
    fmt.Print("Enter your grade: ")
    fmt.Scanf("%s",&input)
    grade,err:= strconv.Atoi(input)
    if err!=nil{
        fmt.Println(err)
    }
    switch grade{
        case 50, 49:
            fmt.Println("Passed")
        case 51:
            fmt.Println("falied")

    }
}
