package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadValuesFromFile(filepath string) []string {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	return strings.Split(string(contents), "\n\n")
}
func calculateScore(b Board) (total int) {
	for _, r := range b.rows {
		for _, n := range r {
			total += n
		}
	}
	return
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func hasBingo(b Board) bool {
	for _, row := range b.rows {
		if len(row) == 0 {
			return true
		}
	}

	for _, col := range b.cols {
		if len(col) == 0 {
			return true
		}
	}
	return false
}

type Row []int
type Col []int
type Board struct {
	rows [5]Row
	cols [5]Col
}

func getRowsAndColumns(board string) (b Board) {
	board = strings.ReplaceAll(board, "/t", "")
	rows := strings.Split(string(board), "\n")
	for ri, row := range rows {
		r := deleteEmpty(strings.Split(row, " "))
		for ci, num := range r {
			n, _ := strconv.Atoi(num)
			b.cols[ci] = append(b.cols[ci], n)
			b.rows[ri] = append(b.rows[ri], n)
		}
	}
	return
}
func removeFromArr(arr []int, num int) []int {
	newArr := []int{}
	for _, n := range arr {
		if n != num {
			newArr = append(newArr, n)
		}
	}
	return newArr
}
func markBoard(b Board, num int) Board {
	for i, r := range b.rows {
		b.rows[i] = removeFromArr(r, num)
	}
	for i, c := range b.cols {
		b.cols[i] = removeFromArr(c, num)
	}
	return b
}
func convertToBoard(input []string) (bs []Board) {
	boards := input[1:]
	for _, board := range boards {
		bs = append(bs, getRowsAndColumns(board))
	}
	return
}

func getScore(input []string) int {
	calledNumbers := strings.Split(input[0], ",")
	bs := convertToBoard(input)
	b, num, _ := getWinningBoard(bs, calledNumbers)
	score := calculateScore(b)
	return num * score
}

func getWinningBoard(bs []Board, calledNumbers []string) (Board, int, int) {
	for _, num := range calledNumbers {
		num, _ := strconv.Atoi(num)
		for i, b := range bs {
			bs[i] = markBoard(b, num)
			if hasBingo(bs[i]) {
				return bs[i], num, i
			}
		}
	}
	return Board{}, 0, 0
}

func getLosingBoardScore(input []string) int {
	calledNumbers := strings.Split(input[0], ",")
	bs := convertToBoard(input)
	for x := len(bs); x > 0; x-- {
		b, num, i := getWinningBoard(bs, calledNumbers)
		bs[i] = bs[len(bs)-1]   // Copy last element to index i.
		bs[len(bs)-1] = Board{} // Erase last element (write zero value).
		bs = bs[:len(bs)-1]
		if x == 1 {
			score := calculateScore(b)
			return num * score
		}
	}

	return 0
}

func CalculatePart1() int {
	input := ReadValuesFromFile("day4/input.txt")
	return getScore(input)
}

func CalculatePart2() int {
	input := ReadValuesFromFile("day4/input.txt")
	return getLosingBoardScore(input)
}
