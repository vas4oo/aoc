package day5

import (
	"fmt"
	"sort"
)

type SeedRange struct {
	start int
	end   int
}

type SeedRanges []SeedRange

func (s SeedRanges) Len() int           { return len(s) }
func (s SeedRanges) Less(i, j int) bool { return s[i].start < s[j].start }
func (s SeedRanges) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func GetSeedRanges(seeds []int) []SeedRange {
	var seedRange []SeedRange
	for i := 0; i < len(seeds); i += 2 {
		var newRange SeedRange
		newRange.start = seeds[i]
		newRange.end = seeds[i] + seeds[i+1]
		seedRange = append(seedRange, newRange)
	}
	return seedRange
}

func pop(alist *[]SeedRange) SeedRange {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

type IntersectionMap map[string]SeedRanges

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetMapped(seedRanges []SeedRange, o string, almanac Almanac) SeedRanges {
	var newI SeedRanges
	for len(seedRanges) > 0 {
		seedRange := pop(&seedRanges)
		added := false
		for _, info := range almanac[o] {
			infoEnd := info.source + info.length
			intersectionStart := max(seedRange.start, info.source)
			intersectionEnd := min(seedRange.end, infoEnd)

			if intersectionStart < intersectionEnd {
				newI = append(newI, SeedRange{start: (intersectionStart - info.source + info.destination), end: (intersectionEnd - info.source + info.destination)})
				added = true
				if intersectionStart > seedRange.start {
					seedRanges = append(seedRanges, SeedRange{start: seedRange.start, end: intersectionStart})
				}
				if seedRange.end > intersectionEnd {
					seedRanges = append(seedRanges, SeedRange{start: intersectionEnd, end: seedRange.end})
				}

				continue
			}

		}
		if !added {
			newI = append(newI, SeedRange{start: seedRange.start, end: seedRange.end})
		}
	}

	return newI
}

func Solve2(seedRanges []SeedRange, almanac Almanac) IntersectionMap {
	intersections := make(IntersectionMap)
	newRanges := seedRanges
	for _, o := range Order {
		intersection := GetMapped(newRanges, o, almanac)
		newRanges = intersection
		intersections[o] = intersection
	}

	return intersections
}

func task2(seeds []int, almanac Almanac) {
	seedRanges := GetSeedRanges(seeds)
	intersections := Solve2(seedRanges, almanac)
	hum := intersections["humidity-to-location"]
	sort.Sort(hum)
	fmt.Println("Task 2", hum[0].start)
}
