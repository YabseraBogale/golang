package main

import (
	"fmt"
	"net"
)

func main() {
	i:=80
	address:=fmt.Sprintf("%d",i)
	_,err:=net.Dial("tcp",address)
	if err==nil {
		fmt.Println("Succfull")
	}
	fmt.Println(address)
	}
