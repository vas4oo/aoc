package day6

import "fmt"

type Race struct {
	Time     int
	Distance int
}

type Races []Race

func Start() {
	// Part 1
	// races := Races{
	// 	{Time: 38, Distance: 234},
	// 	{Time: 67, Distance: 1027},
	// 	{Time: 76, Distance: 1157},
	// 	{Time: 73, Distance: 1236},
	// }
	// Part 2
	races := Races{
		{Time: 38677673, Distance: 234102711571236},
	}
	fmt.Println("ASDSAD")
	task1(races)
	//task2()
}
