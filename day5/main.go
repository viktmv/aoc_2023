package main

import (
    "bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const input string = "input_test.txt"

var markers = [...]string{
	"seed-to-soil",
    "soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

var rangesMap = map[string][][]int{
	"seed-to-soil":            {},
	"soil-to-fertilizer":      {},
	"fertilizer-to-water":     {},
	"water-to-light":          {},
	"light-to-temperature":    {},
	"temperature-to-humidity": {},
	"humidity-to-location":    {},
}

// pt2 - 100165128

func main() {
	file, err := os.Open(input)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	defer file.Close()

	start := time.Now()
	t := time.Now()
	parseAlmanac(file)
	elapsed := t.Sub(start)
	fmt.Printf("took %s\n", elapsed.String())
}

func parseAlmanac(file *os.File) {
	scanner := bufio.NewScanner(file)
	var seeds []int
	var marker string
	var pairs [][]int

	for scanner.Scan() {
		line := scanner.Text()

		// pt 1
		//		if strings.HasPrefix(line, "seeds:") {
		//			_, seedsString, _ := strings.Cut(line, ": ")
		//			seedsList := strings.Split(seedsString, " ")
		//			for _, seed := range seedsList {
		//				if d, ok := toDigit(seed); ok {
		//					seeds = append(seeds, d)
		//				}
		//			}
		//		}

		// pt 2
		if strings.HasPrefix(line, "seeds:") {
			_, seedsString, _ := strings.Cut(line, ": ")
			seedsList := strings.Split(seedsString, " ")

			var currPair []int
			for _, seed := range seedsList {
				if d, ok := toDigit(seed); ok {
					if len(currPair) > 2 {
                        continue
					}
                    currPair = append(currPair, d)
                    if len(currPair) == 2 {
                        pairs = append(pairs, currPair)
                        currPair = nil
                    }
				}
			}
		}

		for _, m := range markers {
			if strings.HasPrefix(line, m) {
				marker = m
			}
		}

		if rangesMap[marker] != nil {
			if r, ok := buildRange(line); ok {
				rangesMap[marker] = append(rangesMap[marker], r)
			}
		}
	}

	for _, pair := range pairs {
		start := pair[0]
		count := pair[1]

		for i := 0; i < count; i++ {
			seeds = append(seeds, start+i)
		}
	}

	var r []int
	for i, m := range markers {
		if i == 0 {
			r = processRange(seeds, rangesMap[m])
		} else {
			r = processRange(r, rangesMap[m])
		}
	}

	min := int(math.Inf(1))
	for _, l := range r {
		if l < min {
			min = l
		}
	}

	fmt.Printf("min: %d\n", min)
}

func processRange(seeds []int, seedRange [][]int) []int {
	var targets []int

	for _, seed := range seeds {
		var found bool
		var target int

		for _, r := range seedRange {
			dest := r[0]
			source := r[1]
			length := r[2]

			if seed >= source+length {
				continue
			}

			if seed >= source && seed <= source+length {
				target = dest + (seed - source) // diff
				found = true
			}
		}

		if !found {
			target = seed
		}

		targets = append(targets, target)
	}

	return targets
}

func buildRange(line string) ([]int, bool) {
	var r []int

	list := strings.Split(line, " ")
	if len(list) == 3 {
		for _, s := range list {
			if d, ok := toDigit(s); ok {
				r = append(r, d)
			}
		}
	}

	return r, len(r) == 3
}

func toDigit(str string) (int, bool) {
	d, err := strconv.Atoi(str)
	if err != nil {
		return 0, false
	}
	return d, true
}
