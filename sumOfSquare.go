package main

import (
	
	"fmt"
	"math"
	
	
	)

func Sum(number float64) float64{
	
	SquareOfSum:=-1*(1-math.Pow(number,2))
	fmt.Println(SquareOfSum)
	SumOfSquare:=math.Pow(((number*(number+1))/2),2)
	
	return SumOfSquare
	
	}

func main() {
	
	
	fmt.Println(Sum(10))
	
	}
