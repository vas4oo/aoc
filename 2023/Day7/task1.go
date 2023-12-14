package day7

import (
	"fmt"
	"sort"
	"strings"
)

func task1(hands Hands) {
	sum := 0
	mapHands := make(Hands) //allRanks := len(hands)
	for _, card := range CardsPower {
		for hand := range hands {
			i := strings.Index(hand, card)
			fmt.Println(i)
			count := strings.Count(hand, card)
			if power, ok := mapHands[hand]; ok {
				if power > count {
					count = power
				}
			}

			mapHands[hand] = count

		}
	}

	var handsPower HandsPower
	for k, v := range mapHands {
		handsPower = append(handsPower, HandPower{k, v})
	}

	sort.Sort(handsPower)

	fmt.Println(handsPower)
	fmt.Println("Task 1:", sum)
}
