package day7

import (
	"strings"

	"github.com/vas4oo/AoC/2023/helpers"
)

// var CardsPower = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
var CardsPower = []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}
var CardsPowerLength = 13

type HandOrder struct {
	Bid       int
	CardPower string
}
type InitialHands map[string]HandOrder

type Hands map[string]int

type HandPower struct {
	Hand       string
	SameCards  int
	CardsPower string
}

type HandsPower []HandPower

func (s HandsPower) Len() int {
	return len(s)
}

func (s HandsPower) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s HandsPower) Less(i, j int) bool {
	if s[i].SameCards == s[j].SameCards {
		return s[i].CardsPower < s[j].CardsPower
	}
	return s[i].SameCards < s[j].SameCards
}

func Start() {
	lines := helpers.ReadLines("day7.txt")
	hands := make(InitialHands)
	for _, line := range lines {
		split := strings.Fields(line)
		hands[split[0]] = HandOrder{Bid: helpers.GetNumber(split[1]), CardPower: "@0 @1 @2 @3 @4"}
	}

	//task1(hands)
	task2(hands)
}
