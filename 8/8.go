package main

import (
	"fmt"
	"strings"

	"github.com/kennedn/aoc2023/common"
)

func one(lines []string) {
	inst := lines[0]
	dirs := make(map[string][]string)
	for _, l := range lines[2:] {
		raw := strings.Split(l, " = ")
		lr := raw[1]
		for _, r := range []string{"(", ")"} {
			lr = strings.Replace(lr, r, "", 1)
		}
		val := strings.Split(lr, ", ")
		dirs[raw[0]] = append(dirs[raw[0]], val...)
	}

	ptr := "AAA"
	count := 1
	for {
		for _, i := range inst {
			if i == 'L' {
				ptr = dirs[ptr][0]
			} else {
				ptr = dirs[ptr][1]
			}
			if ptr == "ZZZ" {
				break
			}
			count++
		}
		if ptr == "ZZZ" {
			break
		}
	}
	fmt.Printf("Count: %d\n", count)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func LCM(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}
	return result
}

func two(lines []string) {
	inst := lines[0]
	dirs := map[string][]string{}
	ptrs := []string{}
	for _, l := range lines[2:] {
		raw := strings.Split(l, " = ")
		lr := raw[1]
		for _, r := range []string{"(", ")"} {
			lr = strings.Replace(lr, r, "", 1)
		}
		val := strings.Split(lr, ", ")
		if raw[0][2] == 'A' {
			ptrs = append(ptrs, raw[0])
		}
		dirs[raw[0]] = append(dirs[raw[0]], val...)
	}

	firstZs := make([]int, len(ptrs))
	for i := 0; i < len(firstZs); i++ {
		firstZs[i] = -1
	}

	count := 0
	allZ := true
	for {
		for _, i := range inst {
			for j, ptr := range ptrs {
				if ptrs[j][2] == 'Z' && firstZs[j] == -1 {
					firstZs[j] = count
				}
				if i == 'L' {
					ptrs[j] = dirs[ptr][0]
				} else {
					ptrs[j] = dirs[ptr][1]
				}
			}
			count++
			allZ = true
			for _, f := range firstZs {
				if f == -1 {
					allZ = false
					break
				}
			}
		}
		if allZ {
			break
		}
	}

	fmt.Printf("Count: %d\n", LCM(firstZs...))
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
