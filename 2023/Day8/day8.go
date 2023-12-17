package day8

import (
	"strings"

	"github.com/vas4oo/AoC/2023/helpers"
)

type Map struct {
	Left  string
	Right string
}
type Maps map[string]Map

func Start() {
	lines := helpers.ReadLines("day8.txt")
	instructions := lines[0]
	maps := make(Maps)

	for _, line := range lines[2:] {
		split := strings.Split(line, "=")
		mapName := strings.TrimSpace(split[0])
		instruction := strings.TrimSpace(split[1])

		instruction = instruction[1 : len(instruction)-1]
		directions := strings.Split(instruction, ",")
		maps[mapName] = Map{Left: strings.TrimSpace(directions[0]), Right: strings.TrimSpace(directions[1])}

	}
	task1(maps, instructions, maps["AAA"])
	//task2(maps, instructions, start, end)
}
