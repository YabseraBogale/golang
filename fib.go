package main

import "fmt"

func fib() func() int{
    f1:=0
    f2:=1
    return func() int{
        f1,f2:=f2,(f1+f2)
        _=f2
        return f1
    }
}

func main(){
    gen :=fib()
    for i:=0;i<5;i++{
        fmt.Println(gen())
    }
}
