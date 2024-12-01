package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var times []int
	var distances []int
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time: ") {
			// times = buildList(line)
			times = append(times, buildNum(line))
		}
		if strings.HasPrefix(line, "Distance: ") {
			// distances = buildList(line)
			distances = append(distances, buildNum(line))
		}
	}

	total := 1

	for i := 0; i < len(times); i++ {
		var numberOfWays int
		for j := 0; j < times[i]; j++ {
			// possible distance
			d := calculateDistance(j, times[i])
			if d > distances[i] {
				numberOfWays += 1
			}
		}
		total *= numberOfWays
	}

	fmt.Println(total)
}

// pt1
func buildList(line string) []int {
	list := []int{}
	_, raw, _ := strings.Cut(line, ": ")

	for _, rawString := range strings.Split(raw, " ") {
		if d, ok := toDigit(rawString); ok {
			list = append(list, d)
		}
	}

	return list
}

// pt 2
func buildNum(line string) int {
	var numberString string
	_, raw, _ := strings.Cut(line, ": ")

	for _, rawString := range strings.Split(raw, " ") {
		if _, ok := toDigit(rawString); ok {
			numberString += rawString
		}
	}

	if d, ok := toDigit(numberString); ok {
		return d
	}

	return 0
}

func calculateDistance(timeToHold, timeTotal int) int {
	// (total time - time to hold) * resulting speed (speed == time to hold)
	return (timeTotal - timeToHold) * timeToHold
}

func toDigit(str string) (int, bool) {
	d, err := strconv.Atoi(str)
	if err != nil {
		return 0, false
	}
	return d, true
}
