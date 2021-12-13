package day12

import (
	"adventOfCode2021/fileInteractions"
	"strings"
	"unicode"
)

func convertToMap(input []string) map[string][]string {
	m := map[string][]string{}
	for _, line := range input {
		vals := strings.Split(line, "-")
		m[vals[0]] = append(m[vals[0]], vals[1])
		m[vals[1]] = append(m[vals[1]], vals[0])
	}
	return m
}
func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func hasVisitedTwice(caves []string) bool {
	smallCaves := map[string]bool{}
	for _, c := range caves {
		if smallCaves[c] {
			return true
		}
		smallCaves[c] = true
	}
	return false
}
func (c *Caves) shouldProceed(nextCave string, smallCaves []string) bool {
	if nextCave == "start" {
		return false
	}
	hasVisitedCaveAlready := contains(smallCaves, nextCave)
	isLowercase := isLower(nextCave)
	hasVisitedTwice := hasVisitedTwice(smallCaves)
	if c.part2 {
		return !(hasVisitedTwice && hasVisitedCaveAlready && isLowercase)
	}
	return !(isLowercase && hasVisitedCaveAlready)
}
func (c *Caves) getNextCave(current string, smallCaves []string) {
	if current == "end" {
		c.countComplete++
		return
	}
	if isLower(current) && current != "start" {
		smallCaves = append(smallCaves, current)
	}

	for _, nextCave := range c.connections[current] {
		if c.shouldProceed(nextCave, smallCaves) {
			c.getNextCave(nextCave, smallCaves)
		}
	}
}

type Caves struct {
	connections   map[string][]string
	countComplete int
	part2         bool
}

func countPaths(input []string, part2 bool) int {
	connections := convertToMap(input)
	c := Caves{connections: connections, part2: part2}
	c.getNextCave("start", []string{})
	return c.countComplete
}

func CalculatePart1(path string) int {
	input := fileInteractions.ReadValuesFromFile(path)
	return countPaths(input, false)
}

func CalculatePart2(path string) int {
	input := fileInteractions.ReadValuesFromFile(path)
	return countPaths(input, true)
}
