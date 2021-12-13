package day13

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord map[string]int

type Fold struct {
	axis string
	val  int
}
type Input struct {
	coords []Coord
	folds  []Fold
}

func readFromFile(filepath string) (input Input) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	both := strings.Split(string(contents), "\n\n")
	coords := strings.Split(both[0], "\n")
	for _, c := range coords {
		b := strings.Split(c, ",")
		x, _ := strconv.Atoi(b[0])
		y, _ := strconv.Atoi(b[1])
		input.coords = append(input.coords, Coord{"x": x, "y": y})
	}
	folds := strings.Split(both[1], "\n")
	for _, f := range folds {
		b := strings.Split(f, "=")
		axis := strings.ReplaceAll(b[0], "fold along ", "")
		val, _ := strconv.Atoi(b[1])
		input.folds = append(input.folds, Fold{axis, val})
	}
	return
}
func (i *Input) fold() {
	f := i.folds[0]
	for _, c := range i.coords {
		if c[f.axis] > f.val {
			diff := c[f.axis] - f.val
			c[f.axis] = f.val - diff
		}
	}
	i.folds = i.folds[1:]
}

func (i Input) countUnique() int {
	m := make(map[string]int, 0)
	for _, c := range i.coords {
		key := strconv.Itoa(c["x"]) + "," + strconv.Itoa(c["y"])
		m[key]++
	}
	return len(m)
}

func (i Input) draw() {
	canvas := [7][40]string{}
	for y := 0; y < 7; y++ {
		for x := 0; x < 40; x++ {
			canvas[y][x] = " "
		}
	}
	for _, c := range i.coords {
		canvas[c["y"]][c["x"]] = "#"
	}
	for _, l := range canvas {
		fmt.Println(l)
	}
}

func CalculatePart1(path string) int {
	input := readFromFile(path)
	input.fold()
	return input.countUnique()
}

func CalculatePart2(path string) {
	input := readFromFile(path)
	l := len(input.folds)
	for i := 0; i < l; i++ {
		input.fold()
	}
	input.draw()
}
