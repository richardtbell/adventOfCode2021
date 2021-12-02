package day1

import (
	"adventOfCode2021/fileInteractions"
	"strconv"
)

func getNumberIncreases(inputArr []int) int {
	increases := 0
	for i, value := range inputArr {
		if i > 0 {
			if value > inputArr[i-1] {
				increases = increases + 1
			}
		}
	}
	return increases
}

func getNumberIncreasesWithSlidingWindow(inputArr []int) int {
	summedArray := []int{}
	for i, value := range inputArr {
		if i < len(inputArr)-2 {
			summedArray = append(summedArray, value+inputArr[i+1]+inputArr[i+2])
		}
	}
	return getNumberIncreases(summedArray)
}

func convertToInts(strArr []string) (values []int) {
	for _, val := range strArr {
		val, _ := strconv.Atoi(val)
		values = append(values, val)
	}
	return
}

func CalculatePart1() int {
	strVals := fileInteractions.ReadValuesFromFile("day1/input.txt")
	return getNumberIncreases(convertToInts(strVals))
}

func CalculatePart2() int {
	strVals := fileInteractions.ReadValuesFromFile("day1/input.txt")
	return getNumberIncreasesWithSlidingWindow(convertToInts(strVals))
}
