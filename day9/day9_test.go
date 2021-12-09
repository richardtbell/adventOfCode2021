package day9

import (
	"reflect"
	"testing"
)

func TestCalculatePart1(t *testing.T) {
	result := CalculatePart1("./input_test.txt")
	expected := 15
	if result != expected {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected, result)
	}
}

func TestCalculatePart2(t *testing.T) {
	result := CalculatePart2("./input_test.txt")
	expected := 1134
	if result != expected {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected, result)
	}
}

//  0123456789
//0 2199943210
//1 3987894921
//2 9856789892
//3 8767896789
//4 9899965678
//
func TestGetBoundsCoordinates(t *testing.T) {
	result := getBoundsCoordinates(readFile("./input_test.txt"))
	expected := [][]int{{2, 0}, {3, 0}, {4, 0}, {1, 1}, {5, 1}, {7, 1}, {0, 2}, {6, 2}, {8, 2}, {5, 3}, {9, 3}, {0, 4}, {2, 4}, {3, 4}, {4, 4}}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected, result)
	}
}

func TestFindLowPoints(t *testing.T) {
	input := readFile("./input_test.txt")
	result := findLowPoints(input)
	expected := []Coords{Coords{x: 1, y: 0}, Coords{y: 0, x: 9}, Coords{x: 2, y: 2}, Coords{y: 4, x: 6}}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected, result)
	}
}

func TestCoordNotRecorded(t *testing.T) {
	cs := []Coords{Coords{x: 1, y: 0}, Coords{y: 0, x: 9}, Coords{x: 2, y: 2}, Coords{y: 4, x: 6}}
	c := Coords{x: 1, y: 0}
	result := coordNotRecorded(c, &cs)
	expected := false
	if result != expected {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected, result)
	}
}

func TestGetAreas(t *testing.T) {
	input := readFile("./input_test.txt")
	lps := findLowPoints(input)
	result := getAreas(lps, input)
	expected := []int{3, 9, 14, 9}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected, result)
	}
}
