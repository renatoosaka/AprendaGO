package main

import (
	"testing"
)

func TestAverage(t *testing.T) {
	avg := Average([]float64{1, 2, 3, 4, 5, 6})
	if avg != 3.2 {
		t.Error("Expected 3.5 got", avg)
	}
}
