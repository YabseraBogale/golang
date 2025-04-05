package function

import (
	"errors"
)

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

func MatrixMinus(y []float64, x []float64) ([]float64, error) {
	if len(x) != len(y) {
		return []float64{}, errors.New("the matrix are not equal")
	}
	result := []float64{}
	for i := 0; i < len(y); i++ {
		result = append(result, y[i]-x[i])
	}
	return result, nil
}

func MatrixMultiple(y []float64, x []float64) ([]float64, error) {
	if len(x) != len(y) {
		return []float64{}, errors.New("the matrix are not equal")
	}
	result := []float64{}
	for i := 0; i < len(y); i++ {
		result = append(result, y[i]*x[i])
	}
	return result, nil
}

func MatrixDivide(y []float64, x []float64) ([]float64, error) {
	if len(x) != len(y) {
		return []float64{}, errors.New("the matrix are not equal")
	}

	result := []float64{}
	for i := 0; i < len(y); i++ {
		if x[i] == 0 {
			return []float64{}, errors.New("divide by zero")
		}
		result = append(result, y[i]/x[i])
	}
	return result, nil

}

func DervativePostive(input []float64, delta float64) []float64 {
	result := []float64{}
	for i := 0; i < len(input); i++ {
		result = append(result, input[i]+delta)
	}
	return result

}

func DervativeNegative(input []float64, delta float64) []float64 {
	result := []float64{}
	for i := 0; i < len(input); i++ {
		result = append(result, input[i]-delta)
	}
	return result

}

func DervativeAtPoint(input []float64, point float64) ([]float64, error) {
	if point == 0 {
		return []float64{}, errors.New("dervatation at point is almost infinity try again")
	}
	f, err := MatrixMinus(DervativePostive(input, point), DervativeNegative(input, point))
	if err != nil {
		return []float64{}, errors.New("input and point produced unequal length of array")
	}
	result := []float64{}
	for _, i := range f {
		result = append(result, i/(2*point))
	}
	return result, nil
}
