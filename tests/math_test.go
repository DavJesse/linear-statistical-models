package tests

import (
	"math"
	"testing"

	"linear-stats/maths"
)

func TestMean(t *testing.T) {
	subject := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 5.5
	got := maths.Mean(subject)

	if got != expected {
		t.Errorf("Expected: %f, Got: %f", expected, got)
		t.Errorf("TestMean_EmptyDataSet failed!")
	}
}

func TestPearsonCoefficient(t *testing.T) {
	input := []int{2, 4, 6, 8}
	output := []int{3, 5, 7, 9}
	got := maths.PearsonCoefficient(input, output)
	expected := float64(1)

	tolerance := 1e-9
	difference := math.Abs(got - expected)

	if difference > tolerance {
		t.Errorf("Expected: %f, Got: %f", expected, got)
		t.Error("TestVariance Failed")
	}
}
