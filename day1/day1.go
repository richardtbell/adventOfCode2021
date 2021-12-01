package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func readValuesFromFile() []int {
	contents, err := os.ReadFile("day1/input.txt")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	strArr := strings.Split(string(contents), "\n")
	return convertToInts(strArr)
}

func CalculatePart1() int {
	return getNumberIncreases(readValuesFromFile())
}

func CalculatePart2() int {
	return getNumberIncreasesWithSlidingWindow(readValuesFromFile())
}
