package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func increment(vals []int) (newVals []int) {
	newCount := 0
	for _, v := range vals {
		newVal := v - 1
		if newVal == -1 {
			newVals = append(newVals, 6)
			newCount++
		} else {
			newVals = append(newVals, newVal)
		}
	}
	for i := 0; i < newCount; i++ {
		newVals = append(newVals, 8)

	}
	return
}

func countAfter(vals []int, iters int) int {
	for i := 0; i < iters; i++ {
		vals = increment(vals)
	}
	return len(vals)
}

func arrAfter(vals []int, iters int) []int {
	for i := 0; i < iters; i++ {
		vals = increment(vals)
	}
	return vals
}

func readFile(filepath string) (ints []int) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	strs := strings.Split(string(contents), ",")
	for _, s := range strs {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return
}
func nextDay(vals [9]int) (newVals [9]int) {
	first := vals[0]
	for i := 0; i < 8; i++ {
		newVals[i] = vals[i+1]
	}
	newVals[6] += first
	newVals[8] = first
	return
}

func countAfterSmart(vals []int, iters int) (count int) {
	fishWithDaysToBirth := [9]int{}
	for _, daysToBirth := range vals {
		fishWithDaysToBirth[daysToBirth]++
	}
	for i := 0; i < iters; i++ {
		fishWithDaysToBirth = nextDay(fishWithDaysToBirth)
	}
	for _, fish := range fishWithDaysToBirth {
		count += fish
	}

	return
}

func CalculatePart1() int {
	input := readFile("day6/input.txt")

	return countAfter(input, 80)
}

func CalculatePart2() int {
	input := readFile("day6/input.txt")

	return countAfterSmart(input, 256)
}
