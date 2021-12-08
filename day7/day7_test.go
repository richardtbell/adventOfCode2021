package day7

import (
	"testing"
)

func TestGetAverage(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	result := getAverage(input)
	if result != 5 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 5, result)
	}
}

func TestGetNextValueToTry(t *testing.T) {
	prev := float64(10)
	curr := float64(20)
	next := float64(30)
	prevPos := 0
	pos := 4
	length := 10
	result := getNextValueToTry(prev, curr, next, pos, length, prevPos)
	if result != 2 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 2, result)
	}
}

func TestGetFuelAtSurroundingPoints(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	prev, curr, next := getFuelAtSurroundingPoints(input, 2, false)
	if prev != 41 {
		t.Fatalf("Expected prev = \n%+v\n got \n%+v", 41, prev)
	}
	if curr != 37 {
		t.Fatalf("Expected curr = \n%+v\n got \n%+v", 37, curr)
	}
	if next != 39 {
		t.Fatalf("Expected next = \n%+v\n got \n%+v", 39, next)
	}

	_, curr2, _ := getFuelAtSurroundingPoints(input, 2, true)
	if curr2 != 206 {
		t.Fatalf("Expected curr2 = \n%+v\n got \n%+v", 206, curr2)
	}
	_, curr3, _ := getFuelAtSurroundingPoints(input, 5, true)
	if curr3 != 168 {
		t.Fatalf("Expected curr3 = \n%+v\n got \n%+v", 168, curr3)
	}
}


func TestFindBestPosition(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	pos, fuel := findBestPosition(input, false)
	if pos != 2 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 2, pos)
	}
	if fuel != 37 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 37, fuel)
	}

	pos2, fuel2 := findBestPosition(input, true)
	if pos2 != 5 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 5, pos2)
	}
	if fuel2 != 168 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 168, fuel2)
	}
}


func TestCalculateFuelRequired(t *testing.T) {
	fuel := calculateFuelRequired(16,2, false)
	if fuel != 14 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 14, fuel)
	}
	fuel2 := calculateFuelRequired(16,5, true)
	if fuel2 != 66 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 66, fuel2)
	}
	fuel3 := calculateFuelRequired(1,5, true)
	if fuel3 != 10 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 10, fuel3)
	}

}


