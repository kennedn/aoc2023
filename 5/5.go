package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/kennedn/aoc2023/common"
)

func one(lines []string) {
	i := -1
	var seeds []int
	var tmp_seeds []int
	for _, l := range lines {
		if strings.HasSuffix(l, ":") {
			if i == -1 {
				tmp_seeds = make([]int, len(seeds))
				copy(tmp_seeds, seeds)
			} else {

				copy(seeds, tmp_seeds)
			}
			i++
		} else if l == "" {
			continue
		} else {
			if i == -1 {
				strs := strings.Fields(strings.Split(l, ":")[1])
				for _, s := range strs {
					v, _ := strconv.Atoi(s)
					seeds = append(seeds, v)
				}
			} else {
				strs := strings.Fields(l)
				var vals []int
				for _, s := range strs {
					v, _ := strconv.Atoi(s)
					vals = append(vals, v)
				}

				for j := 0; j < len(tmp_seeds); j++ {
					// If seed value is in range
					if tmp_seeds[j] == seeds[j] && tmp_seeds[j] < vals[1]+vals[2] && tmp_seeds[j] >= vals[1] {
						tmp_seeds[j] = tmp_seeds[j] - vals[1] + vals[0]
					}
				}
			}
		}
	}
	copy(seeds, tmp_seeds)

	min := -1
	for _, s := range seeds {
		if min > s || min == -1 {
			min = s
		}
	}
	fmt.Printf("1 Min: %d\n", min)
}

func findMinSeed(i int, base int, rng int, converters [][][]int, wg *sync.WaitGroup, resultChan chan<- int, printLock *sync.Mutex) {
	defer wg.Done()
	var minSeed int
	minSeed = -1
	for seed := base; seed < base+rng-1; seed++ {
		printLock.Lock()
		if seed%1000000 == 0 {
			fmt.Printf("\033[%d;0H", i+1)
			fmt.Printf("Seed %.2d: %010d (%.0f%%)      ", i, minSeed, float64(seed)/(float64(base+rng-1))*100)
		}
		printLock.Unlock()
		var seedCandidate int
		seedCandidate = seed
		for _, section := range converters {
			for _, converter := range section {
				if seedCandidate < converter[1]+converter[2] && seedCandidate >= converter[1] {
					seedCandidate = seedCandidate - converter[1] + converter[0]
					break
				}

			}
		}
		if minSeed == -1 || seedCandidate < minSeed {
			minSeed = seedCandidate
		}
	}
	printLock.Lock()
	fmt.Printf("\033[%d;0H", i+1)
	fmt.Printf("Seed %.2d: %d                                                     ", i, minSeed)
	printLock.Unlock()
	resultChan <- minSeed
}

func two(lines []string) {
	i := -1
	var seeds []int
	var converters [][][]int
	for _, l := range lines {
		if strings.HasSuffix(l, ":") {
			converters = append(converters, [][]int{})
			i++
		} else if l == "" {
			continue
		} else {
			if i == -1 {
				strs := strings.Fields(strings.Split(l, ":")[1])
				for _, s := range strs {
					v, _ := strconv.Atoi(s)
					seeds = append(seeds, v)
				}
			} else {
				strs := strings.Fields(l)
				var vals []int
				for _, s := range strs {
					v, _ := strconv.Atoi(s)
					vals = append(vals, v)
				}
				converters[i] = append(converters[i], vals)
			}
		}
	}

	// Clear screen
	fmt.Print("\033[H\033[2J")

	var wg sync.WaitGroup
	var printLock sync.Mutex
	resultChan := make(chan int, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		base := seeds[i]
		rng := seeds[i+1]
		wg.Add(1)
		go findMinSeed(i/2, base, rng, converters, &wg, resultChan, &printLock)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var minSeed int
	minSeed = -1
	for val := range resultChan {
		if minSeed == -1 || val < minSeed {
			minSeed = val
		}
	}

	fmt.Printf("\033[%d;0H", 11)
	fmt.Printf("2 Min: %d\n", minSeed)
}

func main() {
	lines, err := common.AdventDownloader("5", "")
	// lines, err := common.AdventOpener("day_5_sample_1.txt")

	if err != nil {
		fmt.Printf("%v, %s", lines, err.Error())
	}

	one(lines)
	two(lines)
}
