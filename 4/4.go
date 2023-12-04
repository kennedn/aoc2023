package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kennedn/aoc2023/common"
)

func one(lines []string) {
	total := 0
	for _, l := range lines {
		l = strings.Split(l, ": ")[1]
		raw := strings.Split(l, " | ")
		var win, have []int
		for i := 0; i < 2; i++ {
			for _, n := range strings.Fields(raw[i]) {
				if i == 0 {
					val, _ := strconv.Atoi(n)
					win = append(win, val)
				} else {
					val, _ := strconv.Atoi(n)
					have = append(have, val)
				}
			}
		}
		var score int
		multiplier := 0
		for _, h := range have {
			for _, w := range win {
				if h == w {
					score = 1 << multiplier
					multiplier++
				}
			}
		}
		total += score
	}
	fmt.Printf("Score: %d\n", total)
}

func two(lines []string) {
	card_queue := []int{1}
	cards := 0
	for _, l := range lines {
		l = strings.Split(l, ": ")[1]
		var win, have []int
		for i, r := range strings.Split(l, " | ") {
			for _, n := range strings.Fields(r) {
				if i == 0 {
					val, _ := strconv.Atoi(n)
					win = append(win, val)
				} else {
					val, _ := strconv.Atoi(n)
					have = append(have, val)
				}
			}
		}
		cards += card_queue[0]
		for i := 0; i < card_queue[0]; i++ {
			score := 1
			for _, h := range have {
				for _, w := range win {
					if h == w {
						if len(card_queue)-1 < score {
							card_queue = append(card_queue, 1)
						}
						card_queue[score] += 1
						score++
					}
				}
			}
		}

		// Rotate queue left
		card_queue = card_queue[1:]
		card_queue = append(card_queue, 1)
	}

	fmt.Printf("Score: %d\n", cards)
}

func main() {
	lines, err := common.AdventDownloader("4", "")
	// lines, err := common.AdventOpener("day_4_sample_1.txt")

	if err != nil {
		fmt.Printf("%v, %s", lines, err.Error())
	}

	one(lines)
	two(lines)
}
