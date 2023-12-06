package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kennedn/aoc2023/common"
)

func one(lines []string) {
	var data [][]int
	for i, l := range lines {
		data = append(data, []int{})
		for _, d := range strings.Fields(strings.Split(l, ": ")[1]) {
			val, _ := strconv.Atoi(d)
			data[i] = append(data[i], val)
		}
	}

	sum := 1
	for j := 0; j < len(data[0]); j++ {
		count := 0
		time := data[0][j]
		length := data[1][j]
		for i := 0; i < time; i++ {
			if i*(time-i) > length {
				count++
			}
		}
		sum *= count
	}
	fmt.Printf("Count: %d\n", sum)
}

func two(lines []string) {
	var data []int
	for _, l := range lines {
		d := strings.ReplaceAll(strings.Split(l, ": ")[1], " ", "")
		val, _ := strconv.Atoi(d)

		data = append(data, val)
	}

	count := 0
	time := data[0]
	length := data[1]
	for i := 0; i < time; i++ {
		if i*(time-i) > length {
			count++
		}
	}
	fmt.Printf("Count: %d\n", count)
}

func main() {
	lines, err := common.AdventDownloader("6", "")
	// lines, err := common.AdventOpener("day_6_sample_1.txt")

	if err != nil {
		fmt.Printf("%v, %s", lines, err.Error())
	}

	one(lines)
	two(lines)
}
