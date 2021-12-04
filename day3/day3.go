package day3

import (
	"adventOfCode2021/fileInteractions"
	"strconv"
)

func getGammaAndEpsilon(input []string) (string, string) {
	arrLen := len(input[0])
	oneCount := make([]int, arrLen)
	for _, n := range input {
		for position, val := range n {
			if val == '1' {
				oneCount[position]++
			}
		}
	}
	gamma := ""
	epsilon := ""
	length := len(input)
	for _, val := range oneCount {
		if val > length/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	return gamma, epsilon
}

func getCount(input []string, position int) []int {
	count := [2]int{}
	for _, n := range input {
		if string(n[position]) == "0" {
			count[0]++
		}
		if string(n[position]) == "1" {
			count[1]++
		}
	}
	return count[:]
}
func getMostCommonValueAtPosition(input []string, position int) string {
	count := getCount(input, position)
	if count[0] > count[1] {
		return "0"
	} else {
		return "1"
	}
}
func getLeastCommonValueAtPosition(input []string, position int) string {
	count := getCount(input, position)
	if count[0] <= count[1] {
		return "0"
	} else {
		return "1"
	}
}

func getPowerConsumption(input []string) int {
	gamma, epsilon := getGammaAndEpsilon(input)
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	return int(gammaInt) * int(epsilonInt)
}

func getEntriesWithValueAtPosition(input []string, val string, position int) (entries []string) {
	for _, n := range input {
		if string(n[position]) == val {
			entries = append(entries, n)
		}
	}
	return
}
func getO2value(input []string) int {
	l := len(input[0])
	for i := 0; i < l; i++ {
		if len(input) > 1 {
			val := getMostCommonValueAtPosition(input, i)
			input = getEntriesWithValueAtPosition(input, val, i)
		}
	}
	s, _ := strconv.ParseInt(input[0], 2, 64)
	return int(s)
}

func getCO2value(input []string) int {
	l := len(input[0])
	for i := 0; i < l; i++ {
		if len(input) > 1 {
			val := getLeastCommonValueAtPosition(input, i)
			input = getEntriesWithValueAtPosition(input, val, i)
		}
	}
	s, _ := strconv.ParseInt(input[0], 2, 64)
	return int(s)
}

func getLifeSupportRating(input []string) int {
	o2 := getO2value(input)
	co2 := getCO2value(input)
	return o2 * co2
}

func CalculatePart1() int {
	strVals := fileInteractions.ReadValuesFromFile("day3/input.txt")
	return getPowerConsumption(strVals)
}

func CalculatePart2() int {
	strVals := fileInteractions.ReadValuesFromFile("day3/input.txt")
	return getLifeSupportRating(strVals)
}
