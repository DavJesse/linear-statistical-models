package maths

import "math"

func Mean(nums []int) float64 {
	var result float64

	// Find cummulative sum of all numbers
	for _, num := range nums {
		result += float64(num)
	}

	return result / float64(len(nums))
}

func Variance(nums []int) float64 {
	var result float64
	mean := Mean(nums)

	// Calculate cumulative sum of squared deviations
	for _, num := range nums {
		deviation := (float64(num) - mean) * (float64(num) - mean)
		result += deviation
	}

	return result / float64(len(nums))
}

func StandardDeviation(nums []int) float64 {
	variance := Variance(nums) // Calculate variance
	return math.Sqrt(variance)
}
