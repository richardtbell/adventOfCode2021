package day7

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getAverage(input []int) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	av := float64(sum) / float64(len(input))
	return int(math.Floor(av + 0.5))
}

func calculateFuelRequired(start int, finish int,isPart2 bool) float64 {
	diff := math.Abs(float64(start - finish))

	if isPart2 {
		count := 0
		for i := 1; i <= int(diff); i++ {
			count += i
		}
		return float64(count)
	}
	return diff
}
func getFuelAtSurroundingPoints(input []int, pos int, isPart2 bool) (prev float64, curr float64, next float64) {
	for i := 0; i < len(input); i++ {
		prev += calculateFuelRequired(input[i], (pos - 1), isPart2)
		curr += calculateFuelRequired(input[i], pos, isPart2)
		next += calculateFuelRequired(input[i], (pos + 1), isPart2)
	}
	return
}
func isCurrLowest(prev float64, curr float64, next float64) bool {
	return curr < prev && curr < next
}

func getNextValueToTry(prev float64, curr float64, next float64, pos int, length int, prevPos int) int {
	shouldDecreaseGuess := prev < curr && next > curr
	shouldIncreaseGuess := prev > curr && next < curr
	if shouldDecreaseGuess {
		return pos - int(math.Floor(math.Abs(float64(prevPos-pos)/2)+0.5))
	}
	if shouldIncreaseGuess {
		if prevPos < pos {
			prevPos = length
		}
		return pos + int(math.Floor(math.Abs(float64(prevPos-pos)/2)+0.5))
	}
	return pos
}

func findBestPosition(input []int, isPart2 bool) (int, float64) {
	pos := getAverage(input)
	prevPos := 0
	length := len(input)
	prev, curr, next := getFuelAtSurroundingPoints(input, pos, isPart2)
	for ok := true; ok; ok = !isCurrLowest(prev, curr, next) {
		nextPos := getNextValueToTry(prev, curr, next, pos, length, prevPos)
		prevPos = pos
		pos = nextPos
		prev, curr, next = getFuelAtSurroundingPoints(input, pos, isPart2)
	}
	return pos, curr
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

func CalculatePart1() int {
	input := readFile("day7/input.txt")

	_, fuel := findBestPosition(input, false)
	return int(fuel)
}

func CalculatePart2() int {
	input := readFile("day7/input.txt")

	_, fuel := findBestPosition(input, true)
	return int(fuel)
}
