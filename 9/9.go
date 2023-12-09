package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kennedn/aoc2023/common"
)

func one(lines []string) {
	rows := [][]int{}
	for i, l := range lines {
		rows = append(rows, []int{})
		for _, val := range strings.Fields(l) {
			v, _ := strconv.Atoi(val)
			rows[i] = append(rows[i], v)
		}
	}

	extSum := 0
	for _, r := range rows {
		last := []int{}
		for j := len(r) - 1; j >= 0; j-- {
			allZero := true
			var i int
			for i = 0; i < j; i++ {
				r[i] = r[i+1] - r[i]
				if r[i] != 0 {
					allZero = false
				}
			}
			last = append(last, r[j])
			if allZero {
				break
			}
		}
		ext := 0
		for _, l := range last {
			ext += l
		}
		extSum += ext
	}
	fmt.Printf("Sum: %d\n", extSum)
}

func two(lines []string) {
	rows := [][]int{}
	for i, l := range lines {
		rows = append(rows, []int{})
		vals := strings.Fields(l)
		for j := range vals {
			val := vals[len(vals)-j-1]
			v, _ := strconv.Atoi(val)
			rows[i] = append(rows[i], v)
		}
	}

	extSum := 0
	for _, r := range rows {
		last := []int{}
		for j := len(r) - 1; j >= 0; j-- {
			allZero := true
			var i int
			for i = 0; i < j; i++ {
				r[i] = r[i+1] - r[i]
				if r[i] != 0 {
					allZero = false
				}
			}
			last = append(last, r[j])
			if allZero {
				break
			}
		}
		ext := 0
		for _, l := range last {
			ext += l
		}
		extSum += ext
	}
	fmt.Printf("Sum: %d\n", extSum)
}

func main() {
	day := "9"
	lines, err := common.AdventDownloader(day, "")
	// lines, err := common.AdventOpener(fmt.Sprintf("day_%s_sample_1.txt", day))

	if err != nil {
		fmt.Printf("%v, %s", lines, err.Error())
	}

	one(lines)
	two(lines)
}
