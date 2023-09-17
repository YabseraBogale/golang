package main

import (
    "fmt"
    "time"
)

func say(s string, num int){

    for i:=0;i<num;i++{
        time.Sleep(100*time.Millisecond)
        fmt.Println(i,s)
    }

}

func main(){
    go say("Hello",2)
    go say("world",3)
    fmt.Scanln()
}
