package day7

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func calculateScore2(p string) int {
	charCount := make(map[rune]int)
	jCount := strings.Count(p, "J")
	s := strings.Replace(p, "J", "", -1)

	for _, char := range s {
		charCount[char] += 1
	}

	biggestCount := 0
	for _, count := range charCount {
		if count > biggestCount {
			biggestCount = count
		}
	}

	charCountLength := len(charCount)

	if len(s) == 0 || charCountLength == 1 {
		return 6
	}

	if biggestCount == 4 || jCount == 4 || (biggestCount == 3 && jCount == 1) || (biggestCount == 2 && jCount == 2) || (biggestCount == 1 && jCount == 3) {
		return 5
	}

	if (biggestCount == 3 && charCountLength == 2) || (jCount == 3 && biggestCount == 2) || (biggestCount == 2 && charCountLength == 2 && jCount == 1) {
		return 4
	}

	if (biggestCount == 3 && charCountLength == 3) || (jCount == 3 && biggestCount == 1) || (biggestCount == 2 && jCount == 1) || (biggestCount == 1 && jCount == 2) {
		return 3
	}

	if (biggestCount == 2 && charCountLength == 3) || (biggestCount == 2 && jCount == 1) {
		return 2
	}

	if (jCount == 0 && charCountLength == 4) || (jCount == 1 && charCountLength == 4) {
		return 1
	}

	return 0
}

func task2(hands InitialHands) {
	sum := 0
	mapHands := make(Hands)
	for cardIndex, card := range CardsPower {
		for hand, s := range hands {
			re := regexp.MustCompile(card)
			indexes := re.FindAllStringIndex(hand, -1)
			str := s.CardPower
			for _, i := range indexes {
				cP := strconv.Itoa(CardsPowerLength - cardIndex)
				if len(cP) == 1 {
					cP = fmt.Sprintf("0%s", cP)
				}
				newPower := strings.Replace(str, fmt.Sprintf("@%d", i[0]), cP, 1)
				str = newPower

			}
			hands[hand] = HandOrder{Bid: s.Bid, CardPower: str}
			// count := strings.Count(hand, card)
			// fmt.Println(count)
			// if power, ok := mapHands[hand]; ok {
			// 	count += power
			// }

			mapHands[hand] = calculateScore2(hand)

		}
	}

	var handsPower HandsPower
	for k, v := range mapHands {
		handsPower = append(handsPower, HandPower{k, v, hands[k].CardPower})
	}

	//sort.Sort(handsPower)
	sort.Slice(handsPower, func(i, j int) bool {
		iv, jv := handsPower[i], handsPower[j]
		return iv.SameCards < jv.SameCards
	})

	handsPower = singleSort(handsPower)

	fmt.Println(len(handsPower))
	for i, hp := range handsPower {
		sum += (i + 1) * hands[hp.Hand].Bid
		fmt.Println(hp)

	}
	fmt.Println("Task 2:", sum)
}
