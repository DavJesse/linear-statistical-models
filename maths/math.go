package maths

import (
	"log"
	"math"
)

func Mean(nums []int) float64 {
	var result float64

	// Find cummulative sum of all numbers
	for _, num := range nums {
		result += float64(num)
	}

	return result / float64(len(nums))
}

func PearsonCoefficient(input, output []int) float64 {
	var numerator, totalSquaredInputDeviation, totalSquaredOutputDeviation float64
	var i int

	// input and output arrays must have the same size
	if len(input) != len(output) {
		log.Fatal("Input and output arrays must have the same length.")
	}

	meanInput := Mean(input)
	meanOutput := Mean(output)

	// Calculate cumulative sum of multiple deviations
	for i < len(input) {
		inputDeviation := float64(input[i]) - meanInput
		outputDeviation := float64(output[i]) - meanOutput

		// accumulate sums of covariance component(numerator), squared deviations
		numerator += inputDeviation * outputDeviation
		totalSquaredInputDeviation += inputDeviation * inputDeviation
		totalSquaredOutputDeviation += outputDeviation * outputDeviation

		i++
	}

	// Establish denominator of equation
	denominator := math.Sqrt(totalSquaredInputDeviation * totalSquaredOutputDeviation)
	if denominator == 0 {
		log.Fatal("Cannot calculate Pearson correlation coefficient, the standard deviations are zero.")
	}

	return numerator / denominator
}
