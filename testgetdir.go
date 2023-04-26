
package main

import "fmt"
import "os"



func main(){

path, err := os.Getwd()
if err != nil {
	fmt.Println(err)
}

fmt.Println(path)
}
