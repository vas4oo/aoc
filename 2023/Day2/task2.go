package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vas4oo/AoC/2023/helpers"
)

func task2() {
	total := 0
	lines := helpers.ReadLines("day2.txt")

	for _, line := range lines {
		game := strings.Split(line, ": ")

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

		mult := 1
		for _, count := range biggest {
			mult *= count
		}

		total += mult
	}

	fmt.Println("Task 2:", total)
}
