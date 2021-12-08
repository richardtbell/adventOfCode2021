package day8

import (
	"adventOfCode2021/fileInteractions"
	"sort"
	"strconv"
	"strings"
)

// 0:      1:      2:      3:      4:
// aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
// ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
// gggg    ....    gggg    gggg    ....

//  5:      6:      7:      8:      9:
// aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
// dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
// gggg    gggg    ....    gggg    gggg

// a:8
// b:6 unique
// c:8
// d:7
// e:4 unique
// f:9 unique
// g:7
type Mapping map[rune]rune

func determineMapping(signal []string) Mapping {
	m := Mapping{}
	joined := strings.Join(signal, "")
	count := map[rune]int{}
	for _, l := range joined {
		count[l]++
	}
	sort.Slice(signal, func(i, j int) bool {
		return len(signal[i]) < len(signal[j])
	})
	for _, val := range signal {
		switch len(val) {
		case 2: //1
			for _, l := range val {
				if count[l] == 8 {
					m[l] = 'c'
				} else {
					m[l] = 'f'
				}
			}
		case 3: // 7
			for _, l := range val {
				if count[l] == 8 && m[l] != 'c' {
					m[l] = 'a'
				}
			}
		case 4: // 4
			for _, l := range val {
				if count[l] == 6 {
					m[l] = 'b'
				}
				if count[l] == 7 {
					m[l] = 'd'
				}
			}
		case 7: // 8
			for _, l := range val {
				if count[l] == 4 {
					m[l] = 'e'
				}
				if count[l] == 7 && m[l] != 'd' {
					m[l] = 'g'
				}
			}

		}
	}
	return m
}

func translateOutput(output []string, m Mapping) string {
	numMap := map[string]string{
		"abcefg":  "0",
		"cf":      "1",
		"acdeg":   "2",
		"acdfg":   "3",
		"bcdf":    "4",
		"abdfg":   "5",
		"abdefg":  "6",
		"acf":     "7",
		"abcdefg": "8",
		"abcdfg":  "9",
	}
	num := ""
	for _, outs := range output {
		newString := ""
		for _, char := range outs {
			newString += string(m[char])
		}
		strarr := strings.Split(newString, "")
		sort.Strings(strarr)
		newString = strings.Join(strarr, "")
		num += numMap[newString]
	}
	return num

}

func readFile(path string) (signals [][]string, outputs [][]string) {
	input := fileInteractions.ReadValuesFromFile(path)
	for i := 0; i < len(input); i++ {
		line := strings.Split(input[i], " | ")
		signals = append(signals, strings.Split(line[0], " "))
		outputs = append(outputs, strings.Split(line[1], " "))
	}
	return
}

func CalculatePart1() int {
	_, outputs := readFile("day8/input.txt")
	count := 0
	for _, output := range outputs {
		for _, val := range output {
			l := len(val)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}

	return count
}

func CalculatePart2(path string) (total int) {
	signals, outputs := readFile(path)
	for i, s := range signals {
		m := determineMapping(s)
		val, _ := strconv.Atoi(translateOutput(outputs[i], m))
		total += val
	}
	return
}
