package main

import (

    "fmt"
)

func send(c chan string) {

    fmt.Println("Sending a String Channel ... ")
    c <- "Hello World"

}

func recvied(c chan string){

    fmt.Println("String retrieved from channel:", <-c)
}

func main(){
    c:=make(chan string)
    go send(c)
    go recvied(c)

    fmt.Scanln()
}

