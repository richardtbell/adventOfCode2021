package fileInteractions

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
	return strings.Split(string(contents), "\n")
}

func ReadFileIntoMatrix(path string) (matrix [][]int) {
	input := ReadValuesFromFile(path)
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
