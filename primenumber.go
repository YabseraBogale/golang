package main

import "fmt"

func Prime(number int) {

	
	for i:=2;i<number;i++ {
		yes_no:=false
		for j:=2;j<i;j++ {
			
			if i%j==0 {
				yes_no=true
			}
		}
		if yes_no!=true {
			fmt.Println("Prime Number:",i)
		}
	}
	
}
