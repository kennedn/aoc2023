package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kennedn/aoc2023/common"
)

func one(lines []string) {
	var firstDigit, lastDigit rune
	count := 0

	for i, l := range lines {

		re := regexp.MustCompile(`(?:one|two|three|four|five|six|seven|eight|nine|\d)`)
		matches := re.FindAllString(l, -1)
		fmt.Printf("%d: %v\n", i, matches)

		lastDigit = -1
		for _, i := range l {
			if i >= '0' && i <= '9' {
				if lastDigit == -1 {
					firstDigit = i
				}
				lastDigit = i
			}
		}
		val, _ := strconv.Atoi(fmt.Sprintf("%c%c", firstDigit, lastDigit))
		count += val
	}

	fmt.Printf("Count: %d\n", count)

}

func two(lines []string) {

	numberStringMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"0":     "0",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	count := 0
	for _, l := range lines {
		var matches []string
		for j := range l {
			for k, v := range numberStringMap {
				if strings.HasPrefix(l[j:], k) {
					matches = append(matches, v)
				}
			}
		}

		val, _ := strconv.Atoi(fmt.Sprintf("%s%s", numberStringMap[matches[0]], numberStringMap[matches[len(matches)-1]]))
		count += val
	}

	fmt.Printf("Count: %d\n", count)
}

func main() {
	lines, err := common.AdventDownloader("1", "")
	// lines, err := common.AdventOpener("day_1_s2.txt")

	if err != nil {
		fmt.Printf("%v, %s", lines, err.Error())
	}

	one(lines)
	two(lines)
}
