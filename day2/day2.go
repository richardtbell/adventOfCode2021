package day2

import (
	"adventOfCode2021/fileInteractions"
	"strconv"
	"strings"
)

func getPosition(input []string) int {
	horizontal := 0
	depth := 0
	for _, vector := range input {
		s := strings.Split(vector, " ")
		mag, _ := strconv.Atoi(s[1])
		switch s[0] {
		case "forward":
			horizontal = horizontal + mag
		case "down":
			depth = depth + mag
		case "up":
			depth = depth - mag
		}
	}
	return horizontal * depth
}

func getPositionWithAim(input []string) int {
	horizontal := 0
	depth := 0
	aim := 0
	for _, vector := range input {
		s := strings.Split(vector, " ")
		mag, _ := strconv.Atoi(s[1])
		switch s[0] {
		case "forward":
			horizontal = horizontal + mag
			depth = depth + (mag * aim)
		case "down":
			aim = aim + mag
		case "up":
			aim = aim - mag
		}
	}
	return horizontal * depth
}

func CalculatePart1() int {
	strVals := fileInteractions.ReadValuesFromFile("day2/input.txt")
	return getPosition(strVals)
}

func CalculatePart2() int {
	strVals := fileInteractions.ReadValuesFromFile("day2/input.txt")
	return getPositionWithAim(strVals)
}
