package day2

import (
	"testing"
)

func TestGetPosition(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	result := getPosition(input)
	if result != 150 {
		t.Fatalf("Expected position of 150, got %v", result)
	}
}

func TestGetPositionWithAim(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	result := getPositionWithAim(input)
	if result != 900 {
		t.Fatalf("Expected position of 900, got %v", result)
	}
}
