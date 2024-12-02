package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const filename = "input.txt"

func assert(condition bool, message string) {
	if !condition {
		log.Fatalf(message)
	}
}

func main() {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	first := make([]int, 0, 8)
	second := make([]int, 0, 8)

	for scanner.Scan() {
		text := scanner.Text()
		parsed := strings.Split(text, "   ")

		assert(len(parsed) == 2, "should only be two numbers in a row")

		for i, stringNum := range parsed {
			number := mustAtoi(stringNum)
			if i == 0 {
				first = append(first, number)
			} else if i == 1 {
				second = append(second, number)
			}
		}
	}

	sort.Ints(first)
	sort.Ints(second)

	// Part 1
	assert(len(first) == len(second), "got two different lengths")

	deltas := make([]int, 0, 8)
	for i := range first {
		deltas = append(deltas, absDiff(first[i], second[i]))
	}

	log.Printf("Sum is %d\n", sumList(deltas))

	// Part 2
	counter := make(map[int]int, 8)
	for _, num := range second {
		if _, isOk := counter[num]; isOk {
			counter[num] += 1
		} else {
			counter[num] = 1
		}
	}

	similarityScores := make([]int, 0, 8)
	for _, num := range first {
		if value, isOk := counter[num]; isOk {
			similarityScores = append(similarityScores, num*value)
		}
	}

	log.Printf("Sum of similarity scores %d\n", sumList(similarityScores))
}

func mustAtoi(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("is this a number?", err)
	}
	return number
}

func absDiff(n1, n2 int) int {
	if n1 > n2 {
		return n1 - n2
	}
	return n2 - n1
}

func sumList(list []int) int {
	sum := 0
	for _, num := range list {
		sum += num
	}
	return sum
}
