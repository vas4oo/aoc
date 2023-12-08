package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vas4oo/AoC/2023/helpers"
)

func task1() {
	total := 0
	available := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	lines := helpers.ReadLines("day2.txt")

	for _, line := range lines {
		game := strings.Split(line, ": ")
		gameId := strings.Split(game[0], " ")[1]

		tries := strings.Split(game[1], "; ")

		biggest := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, try := range tries {
			cubes := strings.Split(try, ", ")
			for _, cube := range cubes {
				res := strings.Split(cube, " ")
				cubeType := res[1]
				count, _ := strconv.Atoi(res[0])
				if count > biggest[cubeType] {
					biggest[cubeType] = count
				}
			}
		}

		conv := true
		for cubeType, count := range available {
			if biggest[cubeType] > count {
				conv = false
			}
		}

		if conv {
			fmt.Println(gameId)
			converted, _ := strconv.Atoi(gameId)
			total += converted
		}
	}

	fmt.Println("Task 1:", total)
}
