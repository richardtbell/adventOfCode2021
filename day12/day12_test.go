package day12

import (
	"reflect"
	"testing"
)

func TestConvertToMap(t *testing.T) {
	input := []string{"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end"}
	result0 := convertToMap(input)
	expected0 := map[string][]string{
		"start": []string{"A", "b"},
		"A":     []string{"start", "c", "b", "end"},
		"b":     []string{"start", "A", "d", "end"},
		"c":     []string{"A"},
		"d":     []string{"b"},
		"end":   []string{"A", "b"},
	}
	if !reflect.DeepEqual(result0, expected0) {
		t.Fatalf("Expected \n%+v\n got \n%+v", expected0, result0)
	}
}

func TestIsLower(t *testing.T) {
	result := isLower("end")
	if result != true {
		t.Fatalf("Expected \n%+v\n got \n%+v", true, result)
	}
}

func TestContains(t *testing.T) {
	result := contains([]string{"end"}, "end")
	if result != true {
		t.Fatalf("Expected \n%+v\n got \n%+v", true, result)
	}
}

func TestShouldProceed(t *testing.T) {
	c := Caves{}
	result := c.shouldProceed("A", []string{"A"})
	if result != true {
		t.Fatalf("Expected \n%+v\n got \n%+v", true, result)
	}

	result1 := c.shouldProceed("b", []string{"b"})
	if result1 != false {
		t.Fatalf("Expected \n%+v\n got \n%+v", false, result1)
	}

	c.part2 = true
	result2 := c.shouldProceed("b", []string{"b"})
	if result2 != true {
		t.Fatalf("Expected \n%+v\n got \n%+v", true, result2)
	}

	result3 := c.shouldProceed("A", []string{"A", "A"})
	if result3 != true {
		t.Fatalf("Expected \n%+v\n got \n%+v", true, result3)
	}

	result4 := c.shouldProceed("b", []string{"b", "b"})
	if result4 != false {
		t.Fatalf("Expected \n%+v\n got \n%+v", false, result4)
	}

	result5 := c.shouldProceed("start", []string{"b", "b"})
	if result5 != false {
		t.Fatalf("Expected \n%+v\n got \n%+v", false, result5)
	}
}

func TestCountPaths(t *testing.T) {
	result1 := countPaths([]string{
		"start-A",
		"start-b",
		"A-b",
		"A-c",
		"b-d",
		"b-end",
		"A-end",
	}, false)
	if result1 != 10 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 10, result1)
	}

	result2 := countPaths([]string{
		"fs-end",
		"he-DX",
		"fs-he",
		"start-DX",
		"pj-DX",
		"end-zg",
		"zg-sl",
		"zg-pj",
		"pj-he",
		"RW-he",
		"fs-DX",
		"pj-RW",
		"zg-RW",
		"start-pj",
		"he-WI",
		"zg-he",
		"pj-fs",
		"start-RW",
	}, false)
	if result2 != 226 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 226, result2)
	}

	result3 := countPaths([]string{
		"start-A",
		"start-b",
		"A-b",
		"A-c",
		"b-d",
		"b-end",
		"A-end",
	}, true)
	if result3 != 36 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 36, result3)
	}
}
