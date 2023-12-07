package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/kennedn/aoc2023/common"
)

func fiveOfAKind(s string) bool {
	seen := map[rune]int{}
	for _, v := range s {
		seen[v]++
	}

	if len(seen) != 1 {
		return false
	}

	for _, v := range seen {
		if v != 5 {
			return false
		}
	}
	return true
}

func fourOfAKind(s string) bool {
	seen := map[rune]int{}
	for _, v := range s {
		seen[v]++
	}

	if len(seen) != 2 {
		return false
	}

	for _, v := range seen {
		if v != 4 && v != 1 {
			return false
		}
	}
	return true
}

func fullHouse(s string) bool {
	seen := map[rune]int{}
	for _, v := range s {
		seen[v]++
	}

	if len(seen) != 2 {
		return false
	}

	for _, v := range seen {
		if v != 3 && v != 2 {
			return false
		}
	}
	return true
}

func threeOfAKind(s string) bool {
	seen := map[rune]int{}
	for _, v := range s {
		seen[v]++
	}

	if len(seen) != 3 {
		return false
	}

	for _, v := range seen {
		if v != 3 && v != 1 {
			return false
		}
	}
	return true
}

func twoPair(s string) bool {
	seen := map[rune]int{}
	for _, v := range s {
		seen[v]++
	}

	if len(seen) != 3 {
		return false
	}

	for _, v := range seen {
		if v != 2 && v != 1 {
			return false
		}
	}
	return true
}

func onePair(s string) bool {
	seen := map[rune]int{}
	for _, v := range s {
		seen[v]++
	}

	if len(seen) != 4 {
		return false
	}

	for _, v := range seen {
		if v != 2 && v != 1 {
			return false
		}
	}
	return true
}

func highCard(s string) bool {
	seen := map[rune]int{}
	for _, v := range s {
		seen[v]++
	}

	return len(seen) == 5
}

func multiCard(card string, cardFunc func(s string) bool) bool {
	multiCards := []string{card}
	for _, v := range card {
		if v == 'J' {
			for _, mc := range multiCards {
				for _, v2 := range []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'} {
					multiCards = append(multiCards, strings.Replace(mc, "J", string(v2), 1))
				}
			}
		}
	}
	for _, mc := range multiCards {
		if cardFunc(mc) {
			return true
		}

	}
	return false
}

func one(lines []string) {

	strengthMap := map[rune]int{
		'A': 12,
		'K': 11,
		'Q': 10,
		'J': 9,
		'T': 8,
		'9': 7,
		'8': 6,
		'7': 5,
		'6': 4,
		'5': 3,
		'4': 2,
		'3': 1,
		'2': 0,
	}

	var cards [][]any
	for _, l := range lines {
		card := strings.Fields(l)
		val, _ := strconv.Atoi(card[1])
		cards = append(cards, []any{card[0], val})
	}

	for i, c := range cards {
		card := c[0].(string)
		val := 0
		if fiveOfAKind(card) {
			val = 6
		} else if fourOfAKind(card) {
			val = 5
		} else if fullHouse(card) {
			val = 4
		} else if threeOfAKind(card) {
			val = 3
		} else if twoPair(card) {
			val = 2
		} else if onePair(card) {
			val = 1
		} else if highCard(card) {
			val = 0
		}
		cards[i] = append(cards[i], val)
	}

	sort.Slice(cards, func(i, j int) bool {
		if cards[i][2].(int) == cards[j][2].(int) {
			s1 := cards[i][0].(string)
			s2 := cards[j][0].(string)
			for k := 0; k < len(s1); k++ {
				r1 := rune(s1[k])
				r2 := rune(s2[k])
				if strengthMap[r1] != strengthMap[r2] {
					return strengthMap[r1] < strengthMap[r2]
				}
			}
		}
		return cards[i][2].(int) < cards[j][2].(int)
	})

	score := 0
	for i, c := range cards {
		score += c[1].(int) * (i + 1)
	}

	fmt.Printf("Score: %d\n", score)
}

func two(lines []string) {

	strengthMap := map[rune]int{
		'A': 12,
		'K': 11,
		'Q': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
		'J': 0,
	}

	var cards [][]any
	for _, l := range lines {
		card := strings.Fields(l)
		val, _ := strconv.Atoi(card[1])
		cards = append(cards, []any{card[0], val})
	}

	for i, c := range cards {
		card := c[0].(string)
		val := 0
		if multiCard(card, fiveOfAKind) {
			val = 6
		} else if multiCard(card, fourOfAKind) {
			val = 5
		} else if multiCard(card, fullHouse) {
			val = 4
		} else if multiCard(card, threeOfAKind) {
			val = 3
		} else if multiCard(card, twoPair) {
			val = 2
		} else if multiCard(card, onePair) {
			val = 1
		} else if multiCard(card, highCard) {
			val = 0
		}
		cards[i] = append(cards[i], val)
	}

	sort.Slice(cards, func(i, j int) bool {
		if cards[i][2].(int) == cards[j][2].(int) {
			s1 := cards[i][0].(string)
			s2 := cards[j][0].(string)
			for k := 0; k < len(s1); k++ {
				r1 := rune(s1[k])
				r2 := rune(s2[k])
				if strengthMap[r1] != strengthMap[r2] {
					return strengthMap[r1] < strengthMap[r2]
				}
			}
		}
		return cards[i][2].(int) < cards[j][2].(int)
	})

	score := 0
	for i, c := range cards {
		score += c[1].(int) * (i + 1)
	}

	fmt.Printf("Score: %d\n", score)
}

func main() {
	lines, err := common.AdventDownloader("7", "")
	// lines, err := common.AdventOpener("day_7_sample_1.txt")

	if err != nil {
		fmt.Printf("%v, %s", lines, err.Error())
	}

	one(lines)
	two(lines)
}
