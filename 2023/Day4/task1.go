package day4

import (
	"fmt"
	"slices"
	"strings"

	"github.com/vas4oo/AoC/2023/helpers"
)

func CalcPoints(myNums, winningNums []string) (int, int) {
	cardTotal := 0
	count := 0
	for _, winningNum := range winningNums {
		if slices.Contains(myNums, winningNum) {
			count += 1
			if cardTotal == 0 {
				cardTotal += 1
				continue
			}
			cardTotal *= 2
		}
	}
	return cardTotal, count
}

func task1() {
	total := 0

	lines := helpers.ReadLines("day4.txt")
	for _, line := range lines {
		card := strings.Split(line, ": ")[1]
		cardNums := strings.Split(card, " | ")
		winningNums := strings.Fields(cardNums[0])
		myNums := strings.Fields(cardNums[1])
		c, _ := CalcPoints(myNums, winningNums)
		total += c
	}

	fmt.Println("Task 1:", total)
}
