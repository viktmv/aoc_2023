package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input string = "input.txt"

type Card struct {
	Numbers        []int
	WinningNumbers []int
	TotalPoints    int
	Matches        int
}

func main() {
	file, err := os.Open(input)

	if err != nil {
		log.Fatal("could not open input file")
	}

	scanner := bufio.NewScanner(file)

	var cards []Card
	for scanner.Scan() {
		card := parseCard(scanner.Text())
		cards = append(cards, card)
	}

  // pt 1
	var total int
	for _, card := range cards {
		total += card.TotalPoints
	}

	// pt 2
	cardsMap := make(map[int]int)
	for i, card := range cards {
		cardNum := i + 1
		cardsMap[cardNum] += 1

		for l := 0; l < cardsMap[cardNum]; l++ {
			numOfCardsWon := card.Matches
			for j := 1; j <= numOfCardsWon; j++ {
				cardsMap[cardNum+j] += 1
			}
		}
	}

	var sum int
	for _, k := range cardsMap {
		sum += k
	}

	fmt.Println(sum)
}

func parseCard(line string) Card {
	_, rawNumbers, _ := strings.Cut(line, ": ")
	myNumbers, winningNumbers, _ := strings.Cut(rawNumbers, " | ")

	var nums []int
	for _, n := range strings.Split(myNumbers, " ") {
		if d, err := strconv.Atoi(n); err == nil {
			nums = append(nums, d)
		}
	}

	var winningNums []int
	winningMap := make(map[int]int)
	for _, n := range strings.Split(winningNumbers, " ") {
		if d, err := strconv.Atoi(n); err == nil {
			winningNums = append(winningNums, d)
			winningMap[d] = d
		}
	}

	counter := 1
	var points int
	for _, n := range nums {
		if _, ok := winningMap[n]; ok {
			if counter == 1 {
				points += counter
			} else {
				points *= 2
			}
			counter += 1
		}
	}

	return Card{
    Matches: counter - 1, 
    Numbers: nums, 
    WinningNumbers: winningNums, 
    TotalPoints: points,
  }
}

func toDigit(str *string) int {
	d, err := strconv.Atoi(*str)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
