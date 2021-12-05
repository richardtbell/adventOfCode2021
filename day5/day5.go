package day5

import (
	"adventOfCode2021/fileInteractions"
	"strconv"
	"strings"
)

type Row [1000]int
type Matrix [1000]Row

func getNumberRange(i1 int, i2 int) (r []int) {
	if i1 < i2 {
		for i := i1; i <= i2; i++ {
			r = append(r, i)
		}
		return
	} else {
		for i := i2; i <= i1; i++ {
			r = append(r, i)
		}
		return
	}
}

func convertToMatrix(input []string) (m Matrix) {
	for _, coordinateRange := range input {
		fromTo := strings.Split(coordinateRange, " -> ")
		from := strings.Split(fromTo[0], ",")
		to := strings.Split(fromTo[1], ",")
		x1, _ := strconv.Atoi(from[0])
		y1, _ := strconv.Atoi(from[1])
		x2, _ := strconv.Atoi(to[0])
		y2, _ := strconv.Atoi(to[1])
		if x1 == x2 {
			ys := getNumberRange(y1, y2)
			for _, y := range ys {
				m[y][x1]++
			}
		}
		if y1 == y2 {
			xs := getNumberRange(x1, x2)
			for _, x := range xs {
				m[y1][x]++
			}
		}
	}
	return
}

func countOverlaps(m Matrix) (count int) {
	for _, row := range m {
		for _, val := range row {
			if val > 1 {
				count++
			}
		}
	}
	return
}

func CalculatePart1() int {
	input := fileInteractions.ReadValuesFromFile("day5/input.txt")
	m := convertToMatrix(input)
	return countOverlaps(m)
}

// func CalculatePart2() int {
// 	input := ReadValuesFromFile("day5/input.txt")
// 	return getLosingBoardScore(input)
// }
