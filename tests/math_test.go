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
