package day4

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetScore(t *testing.T) {
	input := ReadValuesFromFile("input_test.txt")
	result := getScore(input)

	if result != 4512 {
		t.Fatalf("Expected score of 4512, got %+v", result)
	}
}

func TestGetLosingBoardScore(t *testing.T) {
	input := ReadValuesFromFile("input_test.txt")
	result := getLosingBoardScore(input)

	if result != 1924 {
		t.Fatalf("Expected score of 1924, got %+v", result)
	}
}

func TestGetRowsAndColumns(t *testing.T) {
	input :=
		`22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19`
	b := getRowsAndColumns(input)

	if !reflect.DeepEqual(b.rows[0], []int{22, 13, 17, 11, 0}) {
		t.Fatalf("Expected rows[0] to be []int{22, 13, 17, 11, 0}, got %+v", b.rows[0])
	}
	if !reflect.DeepEqual(b.cols[0], []int{22, 8, 21, 6, 1}) {
		t.Fatalf("Expected rows[0] to be []int{22, 8, 21, 6, 1}, got %+v", b.cols[0])
	}
}

func TestMarkBoard(t *testing.T) {
	input :=
		`22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19`
	b := getRowsAndColumns(input)
	result := markBoard(b, 22)
	if !reflect.DeepEqual(result.rows[0], []int{13, 17, 11, 0}) {
		t.Fatalf("Expected rows[0] to be []int{13, 17, 11, 0}, got %+v", result.rows[0])
	}
	if !reflect.DeepEqual(result.cols[0], []int{8, 21, 6, 1}) {
		t.Fatalf("Expected rows[0] to be []int{8, 21, 6, 1}, got %+v", result.cols[0])
	}
}
func TestCalculateScore(t *testing.T) {
	input1 := Board{
		rows: [5]Row{Row{}, Row{}, Row{}, Row{}, Row{19}},
		cols: [5]Col{Col{}, Col{}, Col{}, Col{}, Col{19}},
	}
	result1 := calculateScore(input1)

	if result1 != 19 {
		t.Fatalf("Expected hasBingo of 19, got %+v", result1)
	}
	input2 := Board{
		rows: [5]Row{Row{}, Row{6}, Row{}, Row{}, Row{19}},
		cols: [5]Col{Col{}, Col{}, Col{6}, Col{}, Col{19}},
	}
	result2 := calculateScore(input2)

	if result2 != 25 {
		t.Fatalf("Expected hasBingo of 25, got %+v", result2)
	}

	input3 := Board{
		rows: [5]Row{Row{}, Row{75, 6}, Row{}, Row{}, Row{19}},
		cols: [5]Col{Col{}, Col{}, Col{6}, Col{}, Col{19}},
	}
	result3 := calculateScore(input3)

	if result3 != 100 {
		t.Fatalf("Expected hasBingo of 100, got %+v", result3)
	}
}

func TestHasBingo(t *testing.T) {
	input := Board{
		rows: [5]Row{Row{22, 13, 17, 11, 0}, Row{8, 2, 23, 4, 24}, Row{21, 9, 14, 16, 7}, Row{6, 10, 3, 18, 5}, Row{1, 12, 20, 15, 19}},
		cols: [5]Col{Col{22, 8, 21, 6, 1}, Col{13, 2, 9, 10, 12}, Col{17, 23, 14, 3, 20}, Col{11, 4, 16, 18, 15}, Col{0, 24, 7, 5, 19}},
	}
	fmt.Printf("%+v", input)
	result1 := hasBingo(input)

	if result1 != false {
		t.Fatalf("Expected hasBingo of false, got %+v", result1)
	}

	input2 := Board{
		rows: [5]Row{Row{13, 17, 11, 0}, Row{8, 2, 23, 4, 24}, Row{21, 9, 14, 16, 7}, Row{6, 10, 3, 18, 5}, Row{1, 12, 20, 15, 19}},
		cols: [5]Col{Col{8, 21, 6, 1}, Col{13, 2, 9, 10, 12}, Col{17, 23, 14, 3, 20}, Col{11, 4, 16, 18, 15}, Col{0, 24, 7, 5, 19}},
	}
	result2 := hasBingo(input2)

	if result2 != false {
		t.Fatalf("Expected hasBingo of false, got %+v", result2)
	}

	input3 := Board{
		rows: [5]Row{Row{}, Row{8, 2, 23, 4, 24}, Row{21, 9, 14, 16, 7}, Row{6, 10, 3, 18, 5}, Row{1, 12, 20, 15, 19}},
		cols: [5]Col{Col{22, 8, 21, 6, 1}, Col{13, 2, 9, 10, 12}, Col{17, 23, 14, 3, 20}, Col{11, 4, 16, 18, 15}, Col{0, 24, 7, 5, 19}},
	}
	result3 := hasBingo(input3)

	if result3 != true {
		t.Fatalf("Expected hasBingo of true, got %+v", result3)
	}

	input4 := Board{
		rows: [5]Row{Row{22, 13, 17, 11, 0}, Row{8, 2, 23, 4, 24}, Row{21, 9, 14, 16, 7}, Row{6, 10, 3, 18, 5}, Row{1, 12, 20, 15, 19}},
		cols: [5]Col{Col{}, Col{13, 2, 9, 10, 12}, Col{17, 23, 14, 3, 20}, Col{11, 4, 16, 18, 15}, Col{0, 24, 7, 5, 19}},
	}
	result4 := hasBingo(input4)

	if result4 != true {
		t.Fatalf("Expected hasBingo of true, got %+v", result4)
	}
}
