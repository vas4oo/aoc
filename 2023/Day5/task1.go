package day5

import (
	"fmt"
	"math"

	"github.com/vas4oo/AoC/2023/helpers"
)

func getMapped(seed int, mapName string, almanac Almanac) int {
	loc := seed
	for _, info := range almanac[mapName] {

		max := info.source + info.length
		if !helpers.InBetween(seed, info.source, max) {
			continue
		}

		dest := seed - info.source
		loc = info.destination + dest
		break
	}

	return loc
}

func Solve(seeds []int, almanac Almanac) int {
	mappedSeed := make(GenMap)
	for _, seed := range seeds {
		m := make(SeedMapping)
		prevVal := seed
		for _, o := range Order {
			loc := getMapped(prevVal, o, almanac)
			m[o] = loc
			prevVal = loc
		}

		mappedSeed[seed] = m
	}

	min := math.MaxInt
	for _, set := range mappedSeed {
		if set["humidity-to-location"] < min {
			min = set["humidity-to-location"]
		}
	}
	return min
}

type SeedMapping map[string]int

type GenMap map[int]SeedMapping

func task1(seeds []int, almanac Almanac) {
	min := Solve(seeds, almanac)
	fmt.Println("Task 1:", min)
}
