// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func sumOfprimes(limit int) int {
	sum := 0
	for i := 2; i < limit; i++ {
		yes_no := false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				yes_no = true
			}
		}
		if yes_no == false {
			//fmt.Println("Prime", i)
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(sumOfprimes(10))
	fmt.Println(sumOfprimes(100))
	fmt.Println(sumOfprimes(1000))
	fmt.Println(sumOfprimes(10000))
	fmt.Println(sumOfprimes(100000))
}
