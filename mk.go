package main

import (
	"fmt"
	"reflect"

)

func main(){

	var name string="YabseraBogale"
	var num=12

	fmt.Println(reflect.ValueOf(name).Kind())
	fmt.Printf("%T\n",num)

}
