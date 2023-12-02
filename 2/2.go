package main

import (
	"fmt"
	"strings"

	"github.com/kennedn/aoc2023/common"
)

func one(lines []string) {
	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	count := 0

	for i, line := range lines {
		line = strings.Split(line, ": ")[1]
		sets := strings.Split(line, "; ")
		line_id := i + 1
		for _, set := range sets {
			for _, pair := range strings.Split(set, ", ") {
				var color string
				var n int
				_, err := fmt.Sscanf(pair, "%d %s", &n, &color)
				if err != nil {
					break
				}
				if n > maxCubes[color] {
					line_id = 0
					break
				}
			}
			if line_id == 0 {
				break
			}
		}
		count += line_id
	}
	fmt.Printf("Count: %d\n", count)

}

func two(lines []string) {

	count := 0

	for _, line := range lines {
		line = strings.Split(line, ": ")[1]
		sets := strings.Split(line, "; ")
		power := 0
		minCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, set := range sets {
			for _, pair := range strings.Split(set, ", ") {
				var color string
				var n int
				_, err := fmt.Sscanf(pair, "%d %s", &n, &color)
				if err != nil {
					break
				}
				if n > minCubes[color] {
					minCubes[color] = n
				}
			}
			power = minCubes["red"] * minCubes["green"] * minCubes["blue"]
		}
		count += power
	}
	fmt.Printf("Count: %d\n", count)
}

func main() {
	lines, err := common.AdventDownloader("2", "")
	// lines, err := common.AdventOpener("day_2_sample_1.txt")

	if err != nil {
		fmt.Printf("%v, %s", lines, err.Error())
	}

	one(lines)
	two(lines)
}
