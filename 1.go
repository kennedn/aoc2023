package main

import (
	"fmt"
)

func main() {
	lines, err := common.advent_downloader("1", "")

	if err != nil {
		fmt.Printf("%v, %s", lines, err.Error())
	} else {
		fmt.Printf("%v", lines)
	}
}
