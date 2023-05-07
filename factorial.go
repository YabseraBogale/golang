/*
Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.
Suppose the following input is supplied to the program: 8
Then, the output should be: 40320
*/
package main
import "fmt"

func factorial(num int) int{
	product:=1
	for i:=1;i<=num;i++ {
		product*=i
	}
	return product
}

func Howlarge(num int) {
	for i:=0;i<num;i++ {
		
		if factorial(i)<0 {
			fmt.Println("Overflows at ",i)
			break
		} else {
			fmt.Printf("factorail of %d is %d \n",i,factorial(i))
		}

	}
}

func main() {
	Howlarge(50)
		
}