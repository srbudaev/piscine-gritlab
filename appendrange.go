package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}

	length := max - min
	result := []int{}

	for i := 0; i < length; i++ {
		result = append(result, min+i)
	}

	return result
}
