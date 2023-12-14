package day6

import (
	"fmt"
	"math"
)

func task1(races Races) {
	num := 1
	for _, race := range races {
		a := float64(-1)
		b := float64(race.Time)
		c := float64(-(race.Distance + 1))

		x1 := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
		x2 := (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)

		num *= int(math.Floor(x2) - math.Ceil(x1) + 1)
	}

	fmt.Println("Task 1:", num)
}
