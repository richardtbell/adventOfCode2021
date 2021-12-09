package day9

import (
	"adventOfCode2021/fileInteractions"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func readFile(path string) (matrix [][]int) {
	input := fileInteractions.ReadValuesFromFile(path)
	for i := 0; i < len(input); i++ {
		line := strings.Split(input[i], "")
		numsArr := []int{}
		for _, val := range line {
			num, _ := strconv.Atoi(val)
			numsArr = append(numsArr, num)
		}
		matrix = append(matrix, numsArr)
	}
	return
}

type Coords struct {
	x int
	y int
}

func findLowPoints(matrix [][]int) (lowPoints []Coords) {
	leny := len(matrix)
	lenx := len(matrix[0])
	for y := 0; y < leny; y++ {
		for x := 0; x < lenx; x++ {
			if y == 0 {
				if x == 0 {
					if matrix[y][x] < matrix[y][x+1] && matrix[y][x] < matrix[y+1][x] {
						lowPoints = append(lowPoints, Coords{x, y})
					}
				} else if x == lenx-1 {
					if matrix[y][x] < matrix[y][x-1] && matrix[y][x] < matrix[y+1][x] {
						lowPoints = append(lowPoints, Coords{x, y})
					}
				} else {
					if matrix[y][x] < matrix[y][x-1] && matrix[y][x] < matrix[y+1][x] && matrix[y][x] < matrix[y][x+1] {
						lowPoints = append(lowPoints, Coords{x, y})

					}
				}
			} else if y == leny-1 {
				if x == 0 {
					if matrix[y][x] < matrix[y][x+1] && matrix[y][x] < matrix[y-1][x] {
						lowPoints = append(lowPoints, Coords{x, y})

					}
				} else if x == lenx-1 {
					if matrix[y][x] < matrix[y][x-1] && matrix[y][x] < matrix[y-1][x] {
						lowPoints = append(lowPoints, Coords{x, y})

					}
				} else {
					if matrix[y][x] < matrix[y][x+1] && matrix[y][x] < matrix[y-1][x] && matrix[y][x] < matrix[y][x-1] {
						lowPoints = append(lowPoints, Coords{x, y})

					}
				}
			} else {
				if x == 0 {
					if matrix[y][x] < matrix[y][x+1] && matrix[y][x] < matrix[y-1][x] && matrix[y][x] < matrix[y+1][x] {
						lowPoints = append(lowPoints, Coords{x, y})

					}
				} else if x == lenx-1 {
					if matrix[y][x] < matrix[y][x-1] && matrix[y][x] < matrix[y-1][x] && matrix[y][x] < matrix[y+1][x] {
						lowPoints = append(lowPoints, Coords{x, y})

					}
				} else {
					if matrix[y][x] < matrix[y][x-1] && matrix[y][x] < matrix[y][x+1] && matrix[y][x] < matrix[y-1][x] && matrix[y][x] < matrix[y+1][x] {
						lowPoints = append(lowPoints, Coords{x, y})

					}
				}
			}
		}
	}
	return
}

func CalculatePart1(path string) (total int) {
	matrix := readFile(path)
	lowPoints := findLowPoints(matrix)
	for _, lp := range lowPoints {
		total += matrix[lp.y][lp.x] + 1
	}
	return
}

func getBoundsCoordinates(matrix [][]int) (bounds [][]int) {
	leny := len(matrix)
	lenx := len(matrix[0])

	for y := 0; y < leny; y++ {
		for x := 0; x < lenx; x++ {
			if matrix[y][x] == 9 {
				bounds = append(bounds, []int{x, y})
			}
		}
	}
	return
}
func coordNotRecorded(c Coords, inBasin *[]Coords) bool {
	for _, coord := range *inBasin {
		if reflect.DeepEqual(c, coord) {
			return false
		}
	}
	return true
}

func moveUp(c Coords, matrix [][]int, inBasin *[]Coords) {
	if c.y > 0 {
		c.y = c.y - 1
		if matrix[c.y][c.x] != 9 && coordNotRecorded(c, inBasin) {
			(*inBasin) = append((*inBasin), c)
			moveLeft(c, matrix, inBasin)
			moveRight(c, matrix, inBasin)
			moveUp(c, matrix, inBasin)
		}
	}
}
func moveLeft(c Coords, matrix [][]int, inBasin *[]Coords) {
	if c.x > 0 {
		c.x = c.x - 1
		if matrix[c.y][c.x] != 9 && coordNotRecorded(c, inBasin) {
			(*inBasin) = append((*inBasin), c)
			moveLeft(c, matrix, inBasin)
			moveUp(c, matrix, inBasin)
			moveDown(c, matrix, inBasin)
		}
	}
}
func moveRight(c Coords, matrix [][]int, inBasin *[]Coords) {
	if c.x < len(matrix[0])-1 {
		c.x = c.x + 1
		if matrix[c.y][c.x] != 9 && coordNotRecorded(c, inBasin) {
			(*inBasin) = append((*inBasin), c)
			moveRight(c, matrix, inBasin)
			moveUp(c, matrix, inBasin)
			moveDown(c, matrix, inBasin)
		}
	}
}
func moveDown(c Coords, matrix [][]int, inBasin *[]Coords) {
	if c.y < len(matrix)-1 {
		c.y = c.y + 1
		if matrix[c.y][c.x] != 9 && coordNotRecorded(c, inBasin) {
			(*inBasin) = append((*inBasin), c)
			moveLeft(c, matrix, inBasin)
			moveRight(c, matrix, inBasin)
			moveDown(c, matrix, inBasin)
		}
	}
}

func getAreas(lowPoints []Coords, matrix [][]int) (areas []int) {
	for _, coord := range lowPoints {
		inBasin := []Coords{coord}
		moveUp(coord, matrix, &inBasin)
		moveLeft(coord, matrix, &inBasin)
		moveRight(coord, matrix, &inBasin)
		moveDown(coord, matrix, &inBasin)
		areas = append(areas, len(inBasin))
	}
	return
}

func CalculatePart2(path string) int {
	matrix := readFile(path)
	lps := findLowPoints(matrix)
	areas := getAreas(lps, matrix)
	sort.Ints(areas)
	top3 := areas[len(areas)-3:]

	return top3[0] * top3[1] * top3[2]
}
