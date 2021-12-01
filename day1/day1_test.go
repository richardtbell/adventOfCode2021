package day1

import (
	"testing"
)

func TestGetNumberIncreases(t *testing.T) {
	input := []int{
		199, 200, 208, 210, 200, 207, 240, 269, 260, 263,
	}
	result := getNumberIncreases(input)
	if result != 7 {
		t.Fatalf("Expected 7, got %v", result)
	}
}

func TestGetNumberIncreasesWithSlidingWindow(t *testing.T) {
	input := []int{
		199, 200, 208, 210, 200, 207, 240, 269, 260, 263,
	}
	result := getNumberIncreasesWithSlidingWindow(input)
	if result != 5 {
		t.Fatalf("Expected 5, got %v", result)
	}
}

func TestConvertToInts(t *testing.T) {
	input := []string{
		"199", "200", "208", "210", "200", "207", "240", "269", "260", "263",
	}
	expected := []int{
		199, 200, 208, 210, 200, 207, 240, 269, 260, 263,
	}
	result := convertToInts(input)

	for i, val := range expected {
		if val != result[i] {
			t.Fatalf("Expected \n%+v\n, got \n%+v\n", val, result[i])
		}
	}
}
