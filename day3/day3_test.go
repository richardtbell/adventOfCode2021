package day3

import (
	"testing"
)

func TestGetPowerConsumption(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	result := getPowerConsumption(input)
	if result != 198 {
		t.Fatalf("Expected powerConsumption of 198, got %v", result)
	}
}

func TestGetMostCommonValueAtPosition(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	result1 := getMostCommonValueAtPosition(input, 0)
	if result1 != "1" {
		t.Fatalf("Expected MostCommonValueAtPosition of 1, got %v", result1)
	}
	result2 := getMostCommonValueAtPosition(input, 1)
	if result2 != "0" {
		t.Fatalf("Expected MostCommonValueAtPosition of 0, got %v", result2)
	}
	evenInput := []string{"1", "0"}
	result5 := getMostCommonValueAtPosition(evenInput, 0)
	if result5 != "1" {
		t.Fatalf("Expected MostCommonValueAtPosition of 1, got %v", result5)
	}
}
func TestGetLifeSupportRating(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	result := getLifeSupportRating(input)
	if result != 230 {
		t.Fatalf("Expected LifeSupportRating of 230, got %v", result)
	}
}
