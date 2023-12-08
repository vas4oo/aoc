package day3

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/vas4oo/AoC/2023/helpers"
)

type Pos struct {
	StartPosition int
	EndPosition   int
	Value         int
}

func findAllNumberSubstrings(input string) []Pos {
	re := regexp.MustCompile("\\d+")
	matches := re.FindAllStringIndex(input, -1)
	var positions []Pos
	for _, match := range matches {
		var pos Pos
		pos.StartPosition = match[0]
		pos.EndPosition = match[1]
		val, _ := strconv.Atoi(input[match[0]:match[1]])
		pos.Value = val
		positions = append(positions, pos)
	}

	return positions
}

func findSymbolIndexes(input string) []int {
	// Define the regular expression pattern to match symbols except dot and numbers
	re := regexp.MustCompile(`[^.\d]+`)

	// Find all matches in the input string
	matches := re.FindAllStringIndex(input, -1)
	var positions []int
	for _, match := range matches {
		positions = append(positions, match[0])
	}

	return positions
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}

func task1() {
	total := 0
	var oldLineNumber []Pos
	var oldLineSymbol []int
	lineIndx := 0
	var added []string

	lines := helpers.ReadLines("day3.txt")
	for _, line := range lines {
		nums := findAllNumberSubstrings(line)
		symbols := findSymbolIndexes(line)
		for _, num := range nums {
			numId := strconv.Itoa(lineIndx) + strconv.Itoa(num.StartPosition) + strconv.Itoa(num.EndPosition)
			for _, symbol := range symbols {
				if (num.StartPosition-1 == symbol || num.EndPosition == symbol) && !contains(added, numId) {
					total += num.Value
					added = append(added, numId)
				}
			}

			for _, symbol := range oldLineSymbol {
				if InBetween(symbol, num.StartPosition-1, num.EndPosition) && !contains(added, numId) {
					total += num.Value
					added = append(added, numId)
				}
			}

		}

		for _, num := range oldLineNumber {
			numId := strconv.Itoa(lineIndx) + strconv.Itoa(num.StartPosition) + strconv.Itoa(num.EndPosition)

			for _, symbol := range symbols {
				if InBetween(symbol, num.StartPosition-1, num.EndPosition) && !contains(added, numId) {
					total += num.Value
					added = append(added, numId)
				}
			}

		}
		oldLineNumber = nums
		oldLineSymbol = symbols
		lineIndx += 1
	}

	fmt.Println("Task 1:", total)
}
