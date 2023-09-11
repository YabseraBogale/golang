package main

import "fmt"

func main(){
    stop:
    for i:=0;i<5;i++{

        for j:=0;j<5;j++{
                if i==3{
                break stop
            }
        }
        fmt.Println(i)

    }
}
