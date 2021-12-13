package day11

import "adventOfCode2021/fileInteractions"

type Coord struct {
	x int
	y int
}

func getSurroundingCoords(c Coord) (sur []Coord) {
	sur = append(sur, Coord{x: c.x, y: c.y + 1})
	sur = append(sur, Coord{x: c.x + 1, y: c.y + 1})
	sur = append(sur, Coord{x: c.x + 1, y: c.y})
	if c.x > 0 {
		sur = append(sur, Coord{x: c.x - 1, y: c.y})
		sur = append(sur, Coord{x: c.x - 1, y: c.y + 1})
	}
	if c.y > 0 {
		sur = append(sur, Coord{x: c.x, y: c.y - 1})
		sur = append(sur, Coord{x: c.x + 1, y: c.y - 1})
	}
	if c.x > 0 && c.y > 0 {
		sur = append(sur, Coord{x: c.x - 1, y: c.y - 1})
	}

	return
}

func getFlashCoords(m [][]int) ([]Coord, [][]int) {
	flashes := []Coord{}
	leny := len(m)
	lenx := len(m[0])
	for y := 0; y < leny; y++ {
		for x := 0; x < lenx; x++ {
			m[y][x]++
			if m[y][x] == 10 {
				flashes = append(flashes, Coord{x, y})
			}
		}
	}
	for i := 0; i < len(flashes); i++ {
		surs := getSurroundingCoords(flashes[i])
		for _, c := range surs {
			if c.y < leny && c.x < lenx {
				m[c.y][c.x]++
				if m[c.y][c.x] == 10 {
					flashes = append(flashes, c)
				}
			}
		}
	}
	for _, c := range flashes {
		m[c.y][c.x] = 0
	}
	return flashes, m
}

func countFlashesAfterSteps(m [][]int, steps int) (count int) {
	for i := 0; i < steps; i++ {
		flashes, newm := getFlashCoords(m)
		m = newm
		count += len(flashes)
	}
	return
}

func CalculatePart1(path string) int {
	m := fileInteractions.ReadFileIntoMatrix(path)
	return countFlashesAfterSteps(m, 100)
}

func findStepForSimultaneousFlash(m [][]int) int {
	length := len(m) * len(m[0])
	flashes := []Coord{}
	step := 0
	for step = 0; len(flashes) < length; step++ {
		flashes, m = getFlashCoords(m)
	}
	return step
}
func CalculatePart2(path string) int {
	m := fileInteractions.ReadFileIntoMatrix(path)
	return findStepForSimultaneousFlash(m)
}
