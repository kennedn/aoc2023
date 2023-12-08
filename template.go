package main

import (
	"fmt"

	"github.com/kennedn/aoc2023/common"
)

func one(lines []string) {
}

func two(lines []string) {
}

func main() {
	day := "8"
	lines, err := common.AdventDownloader(day, "")
	// lines, err := common.AdventOpener(fmt.Sprintf("day_%s_sample_2.txt", day))

	if err != nil {
		fmt.Printf("%v, %s", lines, err.Error())
	}

	one(lines)
	two(lines)
}
