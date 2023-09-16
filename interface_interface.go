package main

import "fmt"


type Shape interface{
    CountSide() (int,int)
}

type TwoD struct {
    width float64
    hiegth float64
}


func (two TwoD) CountSide() (float64,float64){
    return two.width,two.hiegth
}


func main(){
    two:=TwoD{width:5,hiegth:3}
    fmt.Println(two.CountSide())

}
