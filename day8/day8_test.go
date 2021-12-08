package day8

import (
	"reflect"
	"testing"
)

// acedgfb: 8
// cdfbe: 5
// gcdfa: 2
// fbcad: 3
// dab: 7
// cefabd: 9
// cdfgeb: 6
// eafb: 4
// cagedb: 0
// ab: 1

func TestReadFile(t *testing.T) {
	signals, outputs := readFile("./input_test.txt")
	expectedSignals := []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"}
	if reflect.DeepEqual(signals[0][0], expectedSignals) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expectedSignals, signals[0][0])
	}

	expectedOutput := []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}
	if reflect.DeepEqual(outputs[0][0], expectedOutput) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expectedOutput, outputs[0][0])
	}
}

func TestDetermineCoding(t *testing.T) {
	input := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	result := determineMapping(input)
	expected := Mapping{
		'd': 'a',
		'e': 'b',
		'a': 'c',
		'f': 'd',
		'g': 'e',
		'b': 'f',
		'c': 'g',
	}
	// fmt.Println(expected)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected, result)
	}
}

func TestTranslateOutput(t *testing.T) {
	input := []string{"cdfeb", "fcadb", "cdfeb", "cdbaf"}
	m := Mapping{
		'd': 'a',
		'e': 'b',
		'a': 'c',
		'f': 'd',
		'g': 'e',
		'b': 'f',
		'c': 'g',
	}
	result := translateOutput(input, m)
	expected := "5353"
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected, result)
	}
}

func TestCalculatePart2(t *testing.T) {
	result := CalculatePart2("./input_test.txt")
	expected := 61229
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected, result)
	}
}
