package day7

import (
	"strings"

	"github.com/vas4oo/AoC/2023/helpers"
)

var CardsPower = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
var CardsPowerLength = 13

type HandOrder struct {
	Bid       int
	CardPower string
}
type Hands map[string]int

type HandPower struct {
	Hand      string
	SameCards int
}

type HandsPower []HandPower

func (s HandsPower) Len() int           { return len(s) }
func (s HandsPower) Less(i, j int) bool { return s[i].SameCards < s[j].SameCards }
func (s HandsPower) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func Start() {
	lines := helpers.ReadLines("day7.txt")
	hands := make(Hands)
	for _, line := range lines {
		split := strings.Fields(line)
		hands[split[0]] = helpers.GetNumber(split[1])
	}

	task1(hands)
	//task2()
}
