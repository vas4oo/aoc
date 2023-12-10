package day5

import (
	"regexp"

	"github.com/vas4oo/AoC/2023/helpers"
)

type Info struct {
	destination int
	source      int
	length      int
}

type Almanac map[string][]Info

var Order = []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

func Start() {
	lines := helpers.ReadLines("day5.txt")
	numsRe := regexp.MustCompile(`\d+`)
	seedString := numsRe.FindAllString(lines[0], -1)
	var seeds []int
	instructionNameRE := regexp.MustCompile(`^\s*([a-z\-]+)\s*map:`)
	almanac := make(Almanac)
	for _, k := range seedString {
		seeds = append(seeds, helpers.GetNumber(k))
	}

	mapName := ""
	for _, line := range lines[1:] {
		if line == "\n" {
			continue
		}
		name := instructionNameRE.FindAllStringSubmatch(line, -1)
		if len(name) > 0 {
			mapName = name[0][1]
			continue
		}

		nums := numsRe.FindAllString(line, -1)
		if len(nums) == 3 {
			var newAlm Info
			newAlm.destination = helpers.GetNumber(nums[0])
			newAlm.source = helpers.GetNumber(nums[1])
			newAlm.length = helpers.GetNumber(nums[2])
			almanac[mapName] = append(almanac[mapName], newAlm)
		}
	}

	//task1(seeds, almanac)
	task2(seeds, almanac)
}
