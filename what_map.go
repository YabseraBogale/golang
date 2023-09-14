// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

var data map[string] int
func main() {
    data=make(map[string] int)
    data["one"]=1
    data["two"]=2
    fmt.Println(data["one"])
}
