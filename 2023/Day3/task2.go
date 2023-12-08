package day3

import (
	"fmt"
	"regexp"
	"strconv"

	helpers "github.com/vas4oo/AoC/2023/helpers"
)

func findStarsSymbolIndexes(input string) []int {
	// Define the regular expression pattern to match symbols except dot and numbers
	re := regexp.MustCompile(`[*]+`)

	// Find all matches in the input string
	matches := re.FindAllStringIndex(input, -1)
	var positions []int
	for _, match := range matches {
		positions = append(positions, match[0])
	}

	return positions
}

func task2() {
	total := 0
	var oldLineNumber []Pos
	var oldLineSymbol []int
	added := make(map[string]int)

	lines := helpers.ReadLines("day3.txt")

	for lineIndx, line := range lines {
		nums := findAllNumberSubstrings(line)
		symbols := findStarsSymbolIndexes(line)

		for _, num := range nums {
			for _, symbol := range symbols {
				symbolId := strconv.Itoa(lineIndx) + strconv.Itoa(symbol)
				if num.StartPosition-1 == symbol || num.EndPosition == symbol {
					if val, ok := added[symbolId]; ok {
						total += val * num.Value
						delete(added, symbolId)
					} else {
						added[symbolId] = num.Value
					}
				}
			}

			for _, symbol := range oldLineSymbol {
				symbolId := strconv.Itoa(lineIndx-1) + strconv.Itoa(symbol)
				if InBetween(symbol, num.StartPosition-1, num.EndPosition) {
					if val, ok := added[symbolId]; ok {
						total += val * num.Value
						delete(added, symbolId)
					} else {
						added[symbolId] = num.Value
					}
				}
			}

		}

		for _, num := range oldLineNumber {
			for _, symbol := range symbols {
				symbolId := strconv.Itoa(lineIndx) + strconv.Itoa(symbol)
				if InBetween(symbol, num.StartPosition-1, num.EndPosition) {
					if val, ok := added[symbolId]; ok {
						total += val * num.Value
						delete(added, symbolId)
					} else {
						added[symbolId] = num.Value
					}
				}
			}

		}
		oldLineNumber = nums
		oldLineSymbol = symbols
		lineIndx += 1
	}

	fmt.Println("Task 2:", total)
}
