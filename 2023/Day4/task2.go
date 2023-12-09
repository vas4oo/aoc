package day4

import (
	"fmt"
	"strings"

	"github.com/vas4oo/AoC/2023/helpers"
)

type WinningCard struct {
	instances int
	nums      int
}

func task2() {
	total := 0
	allCards := make(map[int]WinningCard)
	lines := helpers.ReadLines("day4.txt")
	for i, line := range lines {
		game := strings.Split(line, ": ")
		card := game[1]
		cardNums := strings.Split(card, " | ")
		winningNums := strings.Fields(cardNums[0])
		myNums := strings.Fields(cardNums[1])
		_, count := CalcPoints(myNums, winningNums)
		var wCard WinningCard
		wCard.nums = count
		if val, ok := allCards[i]; ok {
			wCard.instances = val.instances + 1
		} else {
			wCard.instances = 1
		}

		allCards[i] = wCard
		if count == 0 {
			continue
		}
		for k := 0; k < wCard.instances; k++ {
			for j := i + 1; j < i+1+count; j++ {
				var newCard WinningCard
				if val, ok := allCards[j]; ok {
					newCard.instances += val.instances + 1
				} else {
					newCard.instances = 1
				}

				allCards[j] = newCard
			}
		}
	}

	for _, card := range allCards {

		total += card.instances

	}
	//13820385
	fmt.Println("Task 2:", total)
}
