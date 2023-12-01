package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var filename string = "input.txt"

func main() {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	list := make([]int, 8)

	for scanner.Scan() {
		value := ReadLine(scanner.Text())
		list = append(list, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var result int = 0
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
		result += list[i]
	}

	fmt.Println("Result: ", result)
}

var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func ReadLine(line string) int {
	var first string = ""
	var last string = ""

	rune := []rune(line)

	for i := 0; i < len(rune); i++ {
		letter := rune[i]

		if _, err := strconv.Atoi(string(letter)); err == nil {
			if first == "" {
				first = string(letter)
			}
			last = string(letter)
		} else {
			limit := i + 5

			if i+limit > len(rune) {
				limit = len(rune)
			}

			if p := peek(rune[i:limit]); p != "" {
				if first == "" {
					first = string(p)
				}
				last = string(p)
				i += len(p) - 1
			}
		}
	}

	result, err := strconv.Atoi(first + last)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func peek(str []rune) string {
	var result string = ""

	for i := 0; i <= len(str)-1; i++ {
		result += string(str[i])
		if numbers[result] != "" {
			return numbers[result]
		}
	}

	return ""
}
