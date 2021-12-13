package day10

import (
	"adventOfCode2021/fileInteractions"
	"fmt"
	"sort"
)

// ): 3 points.
// ]: 57 points.
// }: 1197 points.
// >: 25137 points.

func CalculatePart1(path string) int {
	paredBracket := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	lines := fileInteractions.ReadValuesFromFile(path)
	unclosed := []rune{}
	illegals := map[rune]int{}
	for _, line := range lines {
		foundIllegal := false
		for _, char := range line {
			if !foundIllegal {
				if char == '(' || char == '[' || char == '{' || char == '<' {
					unclosed = append(unclosed, char)
				}

				if char == ')' || char == ']' || char == '}' || char == '>' {

					if unclosed[len(unclosed)-1] == paredBracket[char] {
						unclosed = unclosed[:len(unclosed)-1]
					} else {
						illegals[char]++
						foundIllegal = true
					}

				}
			}
		}
	}
	for key, element := range illegals {
		fmt.Println("Key:", string(key), "=>", "Element:", element)
	}
	return illegals[')']*3 + illegals[']']*57 + illegals['}']*1197 + illegals['>']*25137
}

func CalculatePart2(path string) int {
	paredOpenBracket := map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
	lines := fileInteractions.ReadValuesFromFile(path)
	unclosed := map[int][]rune{}
	illegals := map[rune]int{}
	for i, line := range lines {
		foundIllegal := false
		for _, char := range line {
			if !foundIllegal {
				if char == '(' || char == '[' || char == '{' || char == '<' {
					unclosed[i] = append(unclosed[i], char)
				}
				if char == ')' || char == ']' || char == '}' || char == '>' {
					if unclosed[i][len(unclosed[i])-1] == paredOpenBracket[char] {
						unclosed[i] = unclosed[i][:len(unclosed[i])-1]
					} else {
						illegals[char]++
						foundIllegal = true
						delete(unclosed, i)
					}
				}
			}
		}
	}
	paredClosingBracket := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	scoreMap := map[rune]int{')': 1,
		']': 2,
		'}': 3,
		'>': 4}
	scores := []int{}
	for _, line := range unclosed {

		lineScore := 0
		for i := len(line) - 1; i >= 0; i-- {

			closingBracket := paredClosingBracket[line[i]]
			lineScore = lineScore*5 + scoreMap[closingBracket]
		}
		scores = append(scores, lineScore)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}
