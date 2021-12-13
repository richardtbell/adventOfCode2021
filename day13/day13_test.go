package day13

import "testing"

func TestCalculatePart1(t *testing.T) {
	result := CalculatePart1("./input_test.txt")
	if result != 17 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 17, result)
	}
}
