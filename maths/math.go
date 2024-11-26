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

func PearsonCoefficient(input, output []int) float64 {
	var result float64
	var in, out int
	meanInput := Mean(input)
	meanOutput := Mean(output)

	// Calculate cumulative sum of squared deviations
	for in < len(input) && out < len(output) {
		deviation := (float64(input[in]) - meanInput) * (float64(output[out]) - meanOutput)
		result += deviation
		in++
		out++
	}

	result = result / (float64(len(input)) * (math.Sqrt(result / float64(len(input)))) * (math.Sqrt(result / float64(len(output)))))

	return result
}
