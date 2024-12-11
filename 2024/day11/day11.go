package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache map[int]int

func main() {
	file, err := os.Open("input.txt")
	cache = make(map[int]int)

	if err != nil {
		panic("where's the file?")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}

	stones := make([]string, 0, 1000)
	for _, s := range strings.Split(line, " ") {
		if s != "\n" && s != " " {
			fmt.Println(s)
			stones = append(stones, s)
		}
	}

	fmt.Printf("%v\n", stones)
	i := 1
	for range 75 {
		stones = blink(stones)
		fmt.Printf("%d: %v\n", i, len(stones))
		i++
	}
}

func blink(stones []string) []string {
	i := 0
	length := len(stones)

	for i < length {
		stone := stones[i]
		if stone == "0" {
			stones[i] = "1"
			i++
			continue
		}

		if len(stone)%2 == 0 {
			splitIdx := len(stone) / 2
			first := stone[:splitIdx]
			second := rmTrailingZeros(stone[splitIdx:])
			if i+1 >= length {
				stones[i] = first
				stones = append(stones, second)
			} else {
				stones = append(stones[:i+1], stones[i:]...)
				stones[i] = first
				stones[i+1] = second
			}
			length = len(stones)
			i++
			i++
			continue
		}

		if v, ok := cache[mustAtoi(stones[i])]; ok {
			stones[i] = strconv.Itoa(v)
		} else {
			v := mustAtoi(stones[i]) * 2024
			cache[mustAtoi(stones[i])] = v
			stones[i] = strconv.Itoa(v)
		}
		i++
	}

	return stones
}

func mustAtoi(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("could not atoi %v\n", s))
	}
	return number
}

func rmTrailingZeros(s string) string {
	return strconv.Itoa(mustAtoi(s))
}
