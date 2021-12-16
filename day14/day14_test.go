package day14

import (
	"testing"
)

func TestCalculatePart1(t *testing.T) {
	result := CalculatePart1("./input_test.txt")
	if result != 1588 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 1588, result)
	}
}

func TestAddToPolymerChain(t *testing.T) {
	p := readFile("./input_test.txt")
	p.addToPolymerChain()
	if p.chain != "NCNBCHB" {
		t.Fatalf("Expected \n%+v\n got \n%+v", "NCNBCHB", p.chain)
	}
	p.addToPolymerChain()
	if p.chain != "NBCCNBBBCBHCB" {
		t.Fatalf("Expected \n%+v\n got \n%+v", "NBCCNBBBCBHCB", p.chain)
	}
	p.addToPolymerChain()
	if p.chain != "NBBBCNCCNBBNBNBBCHBHHBCHB" {
		t.Fatalf("Expected \n%+v\n got \n%+v", "NBBBCNCCNBBNBNBBCHBHHBCHB", p.chain)
	}
}

func TestAddNewPairs(t *testing.T) {
	p := readFile("./input_test.txt")
	p.addNewPairs()
	expected := map[string]int{
		"NC": 1,
		"CN": 1,
		"NB": 1,
		"BC": 1,
		"CH": 1,
		"HB": 1,
	}
	for k, v := range expected {
		if p.pairs[k] != v {
			t.Fatalf("Expected \n%+v\n got \n%+v", v, p.pairs[k])
		}
	}
	p.addNewPairs()
	expected2 := map[string]int{
		"NB": 2,
		"BC": 2,
		"CC": 1,
		"CN": 1,
		"BB": 2,
		"CB": 2,
		"BH": 1,
		"HC": 1,
	}
	for k, v := range expected2 {
		if p.pairs[k] != v {
			t.Fatalf("Expected \n%+v\n got \n%+v", v, p.pairs[k])
		}
	}
	p.addNewPairs()
	// NBBBCNCCNBBNBNBBC HBHHBCHB

	expected3 := map[string]int{
		"NB": 4,
		"BB": 4,
		"BC": 3,
		"CN": 2,
		"NC": 1,
		"CC": 1,
		"BN": 2,
		"CH": 2,
		"HB": 3,
		"BH": 1,
		"HH": 1,
	}
	for k, v := range expected3 {
		if p.pairs[k] != v {
			t.Fatalf("Expected \n%+v\n got \n%+v", v, p.pairs[k])
		}
	}
}

func TestGetMaxAndMin(t *testing.T) {
	p := readFile("./input_test.txt")
	max, min := p.getMaxAndMin()
	if max != 2 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 2, max)
	}
	if min != 1 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 1, min)
	}
}

// map[B:1 C:1 N:2]
func TestGetMaxAndMinFromPairs(t *testing.T) {
	p := readFile("./input_test.txt")
	max, min := p.getMaxAndMinFromPairs()
	if max != 2 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 2, max)
	}
	if min != 1 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 1, min)
	}
}

func TestCalculatePart2(t *testing.T) {
	result := CalculatePart2("./input_test.txt")
	if result != 2188189693529 {
		t.Fatalf("Expected \n%+v\n got \n%+v", 2188189693529, result)
	}
}

// After step 10, B occurs 1749 times, C occurs 298 times, H occurs 161 times, and N occurs 865 times;
