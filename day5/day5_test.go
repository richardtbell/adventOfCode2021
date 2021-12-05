package day5

import (
	"adventOfCode2021/fileInteractions"
	"reflect"
	"testing"
)

func TestGetNumberRange(t *testing.T) {
	result := getNumberRange(1, 2)
	if !reflect.DeepEqual(result, []int{1, 2}) {
		t.Fatalf("Expected [1,2] got %+v", result)
	}

	result2 := getNumberRange(2, 1)
	if !reflect.DeepEqual(result2, []int{1, 2}) {
		t.Fatalf("Expected [1,2] got %+v", result2)
	}

	result3 := getNumberRange(1, 1)
	if !reflect.DeepEqual(result3, []int{1}) {
		t.Fatalf("Expected [1] got %+v", result3)
	}

	result4 := getNumberRange(9, 0)
	if !reflect.DeepEqual(result4, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fatalf("Expected [0, 1, 2, 3, 4, 5, 6, 7, 8, 9] got %+v", result4)
	}
}

func TestConvertToMatrix(t *testing.T) {
	input := fileInteractions.ReadValuesFromFile("input_test.txt")
	result := convertToMatrix(input)
	expected := Matrix{Row{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, Row{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, Row{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, Row{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, Row{0, 1, 1, 2, 1, 1, 1, 2, 1, 1}, Row{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Row{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Row{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Row{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Row{2, 2, 2, 1, 1, 1, 0, 0, 0, 0}}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected \n%+v\ngot \n%+v", expected, result)
	}
}

func TestCountOverlaps(t *testing.T) {
	input := Matrix{Row{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, Row{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, Row{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, Row{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, Row{0, 1, 1, 2, 1, 1, 1, 2, 1, 1}, Row{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Row{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Row{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Row{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Row{2, 2, 2, 1, 1, 1, 0, 0, 0, 0}}
	result := countOverlaps(input)
	if result != 5 {
		t.Fatalf("Expected 5, got %v", result)
	}
}
