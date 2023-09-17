package main

import (
    "fmt"
    "time"
    "math/rand"
    "sync"
    "sync/atomic"
)

var balance int64
func odd(w *sync.WaitGroup){
    defer w.Done()
    for i:=0;i<5;i++{
        atomic.AddInt64(&balance,100)
       time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
    }
}

func even(w *sync.WaitGroup){
    defer w.Done()
    for i:=0;i<5;i++{
        atomic.AddInt64(&balance,-100)
        time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

    }

}
func main(){
    var w sync.WaitGroup
    balance=200
    w.Add(1)
    go odd(&w)
    w.Add(1)
    go even(&w)
    w.Wait()
    fmt.Println("final balance is",balance)
}
