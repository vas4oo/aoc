package day8

import (
	"fmt"
)

func Solve(maps Maps, instructions string, element Map, step int) int {
	for _, instruction := range instructions {
		//fmt.Println(step)
		step += 1
		result := ""
		if string(instruction) == "L" {
			result = element.Left
		} else {
			result = element.Right
		}

		if result == "ZZZ" {
			return step
		}

		element = maps[result]
	}
	fmt.Println(element)
	return step + Solve(maps, instructions, element, step)

}

func task1(maps Maps, instructions string, element Map) {
	steps := Solve(maps, instructions, element, 0)

	fmt.Println("Task 1:", steps)
}
