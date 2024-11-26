package maths

func Mean(nums []int) float64 {
	var result float64

	// Find cummulative sum of all numbers
	for _, num := range nums {
		result += float64(num)
	}

	return result / float64(len(nums))
}
