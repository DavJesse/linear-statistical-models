package tests

import (
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

func TestVariance(t *testing.T) {
	intSlc := []int{1, 2, 3, 4, 5, 6, 7}
	got := maths.Variance(intSlc)
	expected := float64(4)
	// Compare 'got' and 'expected'
	if got != expected {
		t.Errorf("Expected: %f, Got: %f", expected, got)
		t.Error("TestVariance Failed")
	}
}
