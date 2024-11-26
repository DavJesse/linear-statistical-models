package tests

import (
	"testing"

	"linear-stats/maths"
)

func TestMean_EmptyDataSet(t *testing.T) {
	subject := []int{}
	expected := 0.0
	got := maths.Mean(subject)

	if got != expected {
		t.Errorf("Expected: %f\nGot: %f", expected, got)
		t.Errorf("TestMean_EmptyDataSet failed!")
	}
}
