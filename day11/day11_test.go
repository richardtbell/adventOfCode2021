package day11

import (
	"reflect"
	"testing"
)

func TestGetSurroundingCoords(t *testing.T) {
	result0 := getSurroundingCoords(Coord{x: 0, y: 0})
	expected0 := []Coord{{x: 0, y: 1}, {x: 1, y: 1}, {x: 1, y: 0}}
	if !reflect.DeepEqual(result0, expected0) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected0, result0)

	}
	//  012345
	// 0 x
	// 1
	result1 := getSurroundingCoords(Coord{x: 1, y: 0})
	expected1 := []Coord{{x: 1, y: 1}, {x: 2, y: 1}, {x: 2, y: 0}, {x: 0, y: 0}, {x: 0, y: 1}}
	if !reflect.DeepEqual(result1, expected1) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected1, result1)

	}
	//  012345
	// 0
	// 1x
	result2 := getSurroundingCoords(Coord{x: 0, y: 1})
	expected2 := []Coord{{x: 0, y: 2}, {x: 1, y: 2}, {x: 1, y: 1}, {x: 0, y: 0}, {x: 1, y: 0}}
	if !reflect.DeepEqual(result2, expected2) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected2, result2)
	}
	//  012345
	// 0
	// 1 x
	// 2
	result3 := getSurroundingCoords(Coord{x: 1, y: 1})
	expected3 := []Coord{{x: 1, y: 2}, {x: 2, y: 2}, {x: 2, y: 1}, {x: 0, y: 1}, {x: 0, y: 2}, {x: 1, y: 0}, {x: 2, y: 0}, {x: 0, y: 0}}
	if !reflect.DeepEqual(result3, expected3) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected3, result3)
	}
}

func TestGetFlashCoords(t *testing.T) {
	input := [][]int{{1, 1, 1, 1, 1},
		{1, 9, 9, 9, 1},
		{1, 9, 1, 9, 1},
		{1, 9, 9, 9, 1},
		{1, 1, 1, 1, 1},
	}
	result0, newGrid := getFlashCoords(input)
	expected0 := []Coord{{x: 1, y: 1}, {x: 2, y: 1}, {x: 3, y: 1},
		{x: 1, y: 2}, {x: 3, y: 2},
		{x: 1, y: 3}, {x: 2, y: 3}, {x: 3, y: 3}, {x: 2, y: 2}}
	if !reflect.DeepEqual(result0, expected0) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected0, result0)
	}
	expectedGrid :=
		[][]int{{3, 4, 5, 4, 3},
			{4, 0, 0, 0, 4},
			{5, 0, 0, 0, 5},
			{4, 0, 0, 0, 4},
			{3, 4, 5, 4, 3}}
	if !reflect.DeepEqual(newGrid, expectedGrid) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expectedGrid, newGrid)
	}

}

func TestCountFlashesAfterSteps(t *testing.T) {
	input := [][]int{{1, 1, 1, 1, 1},
		{1, 9, 9, 9, 1},
		{1, 9, 1, 9, 1},
		{1, 9, 9, 9, 1},
		{1, 1, 1, 1, 1},
	}
	result0 := countFlashesAfterSteps(input, 1)
	expected0 := 9
	if !reflect.DeepEqual(result0, expected0) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected0, result0)
	}

	input1 := [][]int{{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6}}
	result1 := countFlashesAfterSteps(input1, 10)
	expected1 := 204
	if !reflect.DeepEqual(result1, expected1) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected1, result1)
	}

	result2 := countFlashesAfterSteps(input1, 100)
	expected2 := 1656
	if !reflect.DeepEqual(result2, expected2) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected2, result2)
	}
}

func TestFindStepForSimultaneousFlash(t *testing.T) {
	input1 := [][]int{{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6}}
	result1 := findStepForSimultaneousFlash(input1)
	expected1 := 195
	if !reflect.DeepEqual(result1, expected1) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected1, result1)
	}
}
