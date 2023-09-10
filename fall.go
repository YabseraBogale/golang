package main

import "fmt"

func main(){
    grade:=2

    switch grade {
        case 2:
            fmt.Println("stops here")
            fallthrough
        case 6:
            fmt.Println("only works wit h \"fallthrogh\" here")

    }

}
