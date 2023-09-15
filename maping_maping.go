package main

import "fmt"

func main(){
    data:=map[string] int{
        "One": 1,
        "Two": 2,
    }

    if _,ok:=data["One"]; ok{
        delete(data,"One")
    } else{
        fmt.Println("doesn't exist")
    }
    fmt.Println(data["One"])
}
