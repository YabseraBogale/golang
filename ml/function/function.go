package function

func SquareInt(x []int) []int {
	result := []int{}

	for _, i := range x {
		result = append(result, i*i)
	}
	return result
}

func SquareFloat64(x []float64) []float64 {
	result := []float64{}

	for _, i := range x {
		result = append(result, i*i)
	}
	return result
}
