package day10

import "testing"

func TestCalculatePart1(t *testing.T) {
	result := CalculatePart1("./input_test.txt")
	expected := 26397
	if result != expected {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected, result)
	}
}
func TestCalculatePart2(t *testing.T) {
	result := CalculatePart2("./input_test.txt")
	expected := 288957
	if result != expected {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected, result)
	}
}
