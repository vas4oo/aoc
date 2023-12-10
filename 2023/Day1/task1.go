package day1

import (
	"fmt"
	"regexp"

	"github.com/vas4oo/AoC/2023/helpers"
)

func task1() {
	total := 0
	re := regexp.MustCompile("[^0-9]+")
	lines := helpers.ReadLines("day1.txt")

	for _, line := range lines {
		numbers := re.ReplaceAllString(line, "")
		numLength := len(numbers)
		if numLength == 0 {
			continue
		}

		num := numbers[0:1] + numbers[numLength-1:]
		total += helpers.GetNumber(num)

	}

	fmt.Println("Task 1:", total)
}
