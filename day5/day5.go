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

func convertToMatrix(input []string, part2 bool) (m Matrix) {
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
		} else if y1 == y2 {
			xs := getNumberRange(x1, x2)
			for _, x := range xs {
				m[y1][x]++
			}
		} else {
			if part2 {

				grad := (y2 - y1) / (x2 - x1)
				if grad == 1 {
					xs := getNumberRange(x1, x2)
					ys := getNumberRange(y1, y2)
					for i := 0; i < len(xs); i++ {
						m[ys[i]][xs[i]]++
					}
				}
				if grad == -1 {
					xs := getNumberRange(x1, x2)
					ys := getNumberRange(y1, y2)
					for i, j := 0, len(ys)-1; i < j; i, j = i+1, j-1 {
						ys[i], ys[j] = ys[j], ys[i]
					}
					for i := 0; i < len(xs); i++ {
						m[ys[i]][xs[i]]++
					}
				}
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
	m := convertToMatrix(input, false)
	return countOverlaps(m)
}

func CalculatePart2() int {
	input := fileInteractions.ReadValuesFromFile("day5/input.txt")
	m := convertToMatrix(input, true)
	return countOverlaps(m)
}
