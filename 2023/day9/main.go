package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var histories [][]int
	for scanner.Scan() {
		history := parseLine(scanner.Text())
		histories = append(histories, history)
	}

	var sum int
	for _, history := range histories {
		var diffs [][]int
		current := history
		for !allZeros(current) {
			current = generateDiffs(current)
			diffs = append(diffs, current)
		}

		// pt 1
		// projection := projectForward(history, diffs)
		// sum += projection[len(projection)-1]

		// pt2
		projection := projectBackwards(history, diffs)
		sum += projection[0]
	}

	fmt.Println(sum)
}

func projectForward(history []int, diffs [][]int) []int {
	diffs = append([][]int{history}, diffs...)
	diffs[len(diffs)-1] = append(diffs[len(diffs)-1], 0)

	for i := len(diffs) - 2; i >= 0; i-- {
		prevDiff := diffs[i+1]
		lastElement := diffs[i][len(diffs[i])-1]
		prevLastElement := prevDiff[len(prevDiff)-1]
		diffs[i] = append(diffs[i], lastElement+prevLastElement)
	}

	return diffs[0]
}

func projectBackwards(history []int, diffs [][]int) []int {
	diffs = append([][]int{history}, diffs...)
	diffs[len(diffs)-1] = append([]int{0}, diffs[len(diffs)-1]...)

	for i := len(diffs) - 2; i >= 0; i-- {
		prevDiff := diffs[i+1]
		firstElement := diffs[i][0]
		prevFirstElement := prevDiff[0]
		diffs[i] = append([]int{firstElement - prevFirstElement}, diffs[i]...)
	}

	return diffs[0]
}

func parseLine(line string) []int {
	var list []int
	for _, v := range strings.Split(line, " ") {
		list = append(list, mustAtoi(v))
	}
	return list
}

func generateDiffs(sequence []int) []int {
	var diffs []int
	for i := 1; i < len(sequence); i++ {
		diffs = append(diffs, sequence[i]-sequence[i-1])
	}
	return diffs
}

func allZeros(sequence []int) bool {
	for _, v := range sequence {
		if v != 0 {
			return false
		}
	}
	return true
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
