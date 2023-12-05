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

var input string = "input.txt"

type Almanac struct {
	Seeds      []int
	SeedToSoil map[int]int
}

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

	var seedToSoil [][]int
	var soilToFertilizer [][]int
	var fertilizerToWater [][]int
	var waterToLight [][]int
	var lightToTemperature [][]int
	var temperatureToHumidity [][]int
	var humidityToLocation [][]int
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
					if len(currPair) < 2 {
						currPair = append(currPair, d)
						if len(currPair) == 2 {
							pairs = append(pairs, currPair)
							currPair = nil
						}
					}
				}
			}
		}

		if strings.HasPrefix(line, "seed-to-soil") {
			marker = "seed-to-soil"
		}

		if strings.HasPrefix(line, "soil-to-fertilizer") {
			marker = "soil-to-fertilizer"
		}

		if strings.HasPrefix(line, "fertilizer-to-water") {
			marker = "fertilizer-to-water"
		}

		if strings.HasPrefix(line, "water-to-light") {
			marker = "water-to-light"
		}

		if strings.HasPrefix(line, "light-to-temperature") {
			marker = "light-to-temperature"
		}

		if strings.HasPrefix(line, "temperature-to-humidity") {
			marker = "temperature-to-humidity"
		}

		if strings.HasPrefix(line, "humidity-to-location") {
			marker = "humidity-to-location"
		}

		if marker == "seed-to-soil" {
			if r, ok := buildRange(line); ok {
				seedToSoil = append(seedToSoil, r)
			}
		}

		if marker == "soil-to-fertilizer" {
			if r, ok := buildRange(line); ok {
				soilToFertilizer = append(soilToFertilizer, r)
			}
		}

		if marker == "fertilizer-to-water" {
			if r, ok := buildRange(line); ok {
				fertilizerToWater = append(fertilizerToWater, r)
			}
		}

		if marker == "water-to-light" {
			if r, ok := buildRange(line); ok {
				waterToLight = append(waterToLight, r)
			}
		}

		if marker == "light-to-temperature" {
			if r, ok := buildRange(line); ok {
				lightToTemperature = append(lightToTemperature, r)
			}
		}

		if marker == "temperature-to-humidity" {
			if r, ok := buildRange(line); ok {
				temperatureToHumidity = append(temperatureToHumidity, r)
			}
		}

		if marker == "humidity-to-location" {
			if r, ok := buildRange(line); ok {
				humidityToLocation = append(humidityToLocation, r)
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

	r := processRange(seeds, seedToSoil)
	r1 := processRange(r, soilToFertilizer)
	r2 := processRange(r1, fertilizerToWater)
	r3 := processRange(r2, waterToLight)
	r4 := processRange(r3, lightToTemperature)
	r5 := processRange(r4, temperatureToHumidity)
	r6 := processRange(r5, humidityToLocation)

	min := int(math.Inf(1))
	for _, l := range r6 {
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
