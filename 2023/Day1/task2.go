package day1

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/vas4oo/AoC/2023/helpers"
)

func replaceWordsWithDigits(text string) string {
	replacements := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	result := text
	for word, digit := range replacements {
		result = strings.Replace(result, word, digit, -1)
	}

	return result
}

func task2() {
	total := 0
	re := regexp.MustCompile("[^1-9]+")
	lines := helpers.ReadLines("day3.txt")

	for _, line := range lines {
		replaced := replaceWordsWithDigits(line)
		numbers := re.ReplaceAllString(replaced, "")
		numLength := len(numbers)
		if numLength == 0 {
			continue
		}

		num := numbers[0:1] + numbers[numLength-1:]
		total += helpers.GetNumber(num)

	}

	fmt.Println("Task 2:", total)
}
