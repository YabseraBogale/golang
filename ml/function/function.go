package function

import "errors"

func SquareFloat64(x []float64) []float64 {
	result := []float64{}

	for _, i := range x {
		result = append(result, i*i)
	}
	return result
}

func MatrixSum(x []float64, y []float64) ([]float64, error) {
	if len(x) != len(y) {
		return []float64{}, errors.New("the matrix are not equal")
	}
	result := []float64{}
	for i := 0; i < len(y); i++ {
		result = append(result, y[i]+x[i])
	}
	return result, nil
}

func MatrixMinus(x []float64, y []float64) ([]float64, error) {
	if len(x) != len(y) {
		return []float64{}, errors.New("the matrix are not equal")
	}
	result := []float64{}
	for i := 0; i < len(y); i++ {
		result = append(result, y[i]-x[i])
	}
	return result, nil
}

func MatrixMultiple(x []float64, y []float64) ([]float64, error) {
	if len(x) != len(y) {
		return []float64{}, errors.New("the matrix are not equal")
	}
	result := []float64{}
	for i := 0; i < len(y); i++ {
		result = append(result, y[i]*x[i])
	}
	return result, nil
}

func DervativeFunc(f func(x []int, d float64) []int, input []int, delta float64) []int {
	return []int{}
}
