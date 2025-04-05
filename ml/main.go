package main

import (
	"fmt"

	"github.com/YabseraBogale/ml/function"
)

func main() {
	result := function.SquareFloat64([]float64{0, 1, 1, 2.3, 3, 5, 8})
	fmt.Println(result)
}
