// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

func two(a int,b int) (int,int){
    //use this not this (a,b)
    return a,b
} 

func main() {
  a,b:=two(1,2)
  fmt.Println("a:",a)
  fmt.Println("b:",b)
}
