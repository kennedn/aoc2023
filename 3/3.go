package main

import (
	"fmt"
	"strconv"

	"github.com/kennedn/aoc2023/common"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func one(lines []string) {
	var part_sum int
	for i := range lines {
		for j := 0; j < len(lines); j++ {
			number_block := ""
			if lines[i][j] > '9' || lines[i][j] < '0' {
				continue
			}
			start := j
			for {
				if j >= len(lines[0]) || lines[i][j] > '9' || lines[i][j] < '0' {
					break
				}
				number_block += string(lines[i][j])
				j++
			}
			value, _ := strconv.Atoi(number_block)

			start_idx := max(0, start-1)
			end_idx := min(len(lines)-1, j)
			number_block += string(lines[i][start_idx])                                  // Left char
			number_block += string(lines[i][end_idx])                                    // Right char
			number_block += string(lines[max(0, i-1)][start_idx : end_idx+1])            // Above line
			number_block += string(lines[min(len(lines)-1, i+1)][start_idx : end_idx+1]) // Below line

			for _, c := range number_block {
				if (c > '9' || c < '0') && c != '.' {
					part_sum += value
					break
				}
			}
		}
	}
	fmt.Printf("Sum: %d\n", part_sum)
}

func two(lines []string) {
	var gears [][]int
	var gear_sum int
	for i := range lines {
		for j := 0; j < len(lines); j++ {
			number_block := ""
			if lines[i][j] > '9' || lines[i][j] < '0' {
				continue
			}
			start := j
			for {
				if j >= len(lines[0]) || lines[i][j] > '9' || lines[i][j] < '0' {
					break
				}
				number_block += string(lines[i][j])
				j++
			}
			value, _ := strconv.Atoi(number_block)

			start_idx := max(0, start-1)
			end_idx := min(len(lines)-1, j)

			for l := i - 1; l <= i+1; l++ {
				for k := start_idx; k < end_idx+1; k++ {
					if l < 0 || l > len(lines)-1 {
						continue
					}
					if lines[l][k] == '*' {
						gears = append(gears, []int{value, l, k})
					}
				}
			}

		}
	}
	for i, g1 := range gears {
		for _, g2 := range gears[i+1:] {
			if g1[1] == g2[1] && g1[2] == g2[2] {
				gear_sum += g1[0] * g2[0]
				break
			}
		}
	}

	fmt.Printf("Sum: %d\n", gear_sum)
}

func main() {
	lines, err := common.AdventDownloader("3", "")
	// lines, err := common.AdventOpener("day_3_sample_1.txt")

	if err != nil {
		fmt.Printf("%v, %s", lines, err.Error())
	}

	one(lines)
	two(lines)
}
