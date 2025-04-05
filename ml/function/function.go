package function

func Square(x []int) []int {
	result := []int{}

	for _, i := range x {
		result = append(result, i*i)
	}
	return result
}
