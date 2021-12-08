package day6

import (
	"reflect"
	"testing"
)

func TestIncrement(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	day1 := []int{2, 3, 2, 0, 1}
	day2 := []int{1, 2, 1, 6, 0, 8}
	day17 := []int{0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 0, 1, 2, 2, 2, 3, 3, 4, 4, 5, 7, 8}
	day18 := []int{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8}
	if !reflect.DeepEqual(increment(input), day1) {
		t.Fatalf("Expected \n%+v\n got \n%+v", day1, increment(input))
	}
	if !reflect.DeepEqual(increment(day1), day2) {
		t.Fatalf("Expected \n%+v\n got \n%+v", day2, increment(day1))
	}
	if !reflect.DeepEqual(increment(day17), day18) {
		t.Fatalf("Expected \n%+v\n got \n%+v", day18, increment(day17))
	}
}

func TestCountAfter(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}
	result := countAfter(input, 2)
	if result != 6 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 6, result)
	}
	result1 := countAfter(input, 18)
	if result1 != 26 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 26, result1)
	}
	result2 := countAfter(input, 80)
	if result2 != 5934 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 5934, result2)
	}
}
func TestNextDay(t *testing.T) {
	input := [9]int{1, 0, 0, 0, 0, 0, 0, 0, 0}
	result := nextDay(input)
	expected := [9]int{0, 0, 0, 0, 0, 0, 1, 0, 1}
	if result != expected {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected, result)
	}
}
func TestCountAfterSmart(t *testing.T) {
	input := []int{3, 4, 3, 1, 2}

	result1 := countAfterSmart(input, 18)
	if result1 != 26 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 26, result1)
	}
	result2 := countAfter(input, 80)
	if result2 != 5934 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 5934, result2)
	}
}
