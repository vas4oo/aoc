package day5

import (
	"fmt"
	"slices"
	"sync"
)

func GetAll(start, length int, almanac Almanac, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	var seeds []int

	for i := start; i < start+length; i++ {
		seeds = append(seeds, i)
	}

	ch <- Solve(seeds, almanac)
}

func task2(seeds []int, almanac Almanac) {
	channel := make(chan int, len(seeds))
	var wg sync.WaitGroup
	for i := 0; i < len(seeds); i += 2 {
		wg.Add(1)
		go GetAll(seeds[i], seeds[i+1], almanac, channel, &wg)
		// if m < min {
		// 	min = m
		// }
	}
	wg.Wait()
	close(channel)

	var mins []int

	for i := range channel {
		mins = append(mins, i)
	}

	fmt.Println("Task 2", slices.Min(mins))
}
