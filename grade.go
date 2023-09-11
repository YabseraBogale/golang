package main

import (
    "fmt"
    "strconv"
)

func main(){
    var input string
    fmt.Print("Enter Your Grade: ")
    fmt.Scanf("%s",&input)
    grade,err:=strconv.Atoi(input)
    if err!=nil{
        fmt.Println(err)
    }
    switch {
        case grade<0: fmt.Println("Enter the grade correctly")
        case grade>0 && grade <=30: fmt.Println("F")
        case grade>=31 && grade<=40: fmt.Println("D")
        case grade>=41 && grade<=70: fmt.Println("C")
        case grade>=71 && grade<=80: fmt.Println("B")
        case grade>=81 && grade<=100: fmt.Println("A")
        default: fmt.Println("Are you a student ?")

    }

}
