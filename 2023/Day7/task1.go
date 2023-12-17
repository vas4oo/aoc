package day7

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/vas4oo/AoC/2023/helpers"
)

func filterBySameCards(handPowers []HandPower, targetSameCards int) []HandPower {
	var filteredHandPowers []HandPower

	for _, hp := range handPowers {
		if hp.SameCards == targetSameCards {
			filteredHandPowers = append(filteredHandPowers, hp)
		}
	}

	return filteredHandPowers
}

func singleSort(oldHandsPower HandsPower) HandsPower {
	var handsPower HandsPower
	for i := 0; i <= 6; i++ {
		r := filterBySameCards(oldHandsPower, i)
		sort.Slice(r, func(i, j int) bool {
			iv, jv := r[i], r[j]
			ivN := helpers.GetNumber(strings.ReplaceAll(iv.CardsPower, " ", ""))
			jvN := helpers.GetNumber(strings.ReplaceAll(jv.CardsPower, " ", ""))
			return ivN < jvN
		})
		fmt.Println(len(r))
		handsPower = append(handsPower, r...)
	}

	return handsPower
}

func calculateScore(s string) int {
	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char] += 1
	}

	if len(charCount) == 5 || len(charCount) == 4 {
		return 5 - len(charCount)
	}

	if len(charCount) == 1 {
		return 6
	}

	biggestCount := 0
	for _, count := range charCount {
		if count > biggestCount {
			biggestCount = count
		}
	}

	if biggestCount == 2 && len(charCount) == 3 {
		return 2
	}

	if biggestCount == 3 && len(charCount) == 3 {
		return 3
	}

	if biggestCount == 3 && len(charCount) == 2 {
		return 4
	}

	if biggestCount == 4 {
		return 5
	}

	return 0
}

func task1(hands InitialHands) {
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
		if hp.Hand == "JJJJJ" {
			fmt.Println((i + 1) * hands[hp.Hand].Bid)
		}
	}
	fmt.Println("Task 1:", sum)
}
