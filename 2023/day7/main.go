package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	kind  string
	bid   int
	rank  int
}

var kindPowers = map[string]int{
	"five-of-a-kind":  7,
	"four-of-a-kind":  6,
	"full-house":      5,
	"three-of-a-kind": 4,
	"two-pair":        3,
	"one-pair":        2,
	"high-card":       1,
}

var cardPowers = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	// pt1
	// "J": 11,
	// pt2
	"J": 1,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cards, bid, _ := strings.Cut(line, " ")
		kind := DetermineKind(cards)
		hand := Hand{cards: cards, bid: mustAtoi(bid), kind: kind, rank: 1}
		hands = append(hands, hand)
	}

	rankHands(hands)

	winnings := 0
	for _, hand := range hands {
		winnings += hand.bid * hand.rank
	}

	fmt.Printf("winnings: %d\n", winnings)
}

func rankHands(hands []Hand) {
	for i, h := range hands {
		for j, other := range hands {
			if i == j {
				continue
			}
			if compareCards(h, other) == 1 {
				hands[i].rank++
			}
		}
	}
}

func countWithJoker(cards string) map[string]int {
	var counter = make(map[string]int)
	for _, c := range cards {
		counter[string(c)]++
	}

	var max string
	var maxCount int
	for c, v := range counter {
		if c == "J" {
			continue
		}
		if v > maxCount {
			max = c
			maxCount = v
		}
	}

	if counter["J"] > 0 {
		// use Joker to increment the max card
		counter[max] += counter["J"]
		// discount the used Joker
		counter["J"] -= counter["J"]
	}

	return counter
}

func count(cards string) map[string]int {
	var counter = make(map[string]int)
	for _, c := range cards {
		counter[string(c)]++
	}
	return counter
}

func DetermineKind(cards string) string {
	// pt1
	// counter := count(cards)
	// pt2
	counter := countWithJoker(cards)

	allUnique := true
	onePair := false
	threeOfAKind := false
	for _, v := range counter {
		if v > 1 {
			allUnique = false
		}
		if v == 5 {
			return "five-of-a-kind"
		}
		if v == 4 {
			return "four-of-a-kind"
		}
		if v == 3 {
			for _, c := range counter {
				if c == 2 {
					return "full-house"
				}
			}
			threeOfAKind = true
		}

		if v == 2 {
			if onePair {
				return "two-pair"
			}
			onePair = true
		}
	}

	if allUnique {
		return "high-card"
	}

	if threeOfAKind && !onePair {
		return "three-of-a-kind"
	}

	if onePair {
		return "one-pair"
	}

	return ""
}

func compareCards(hand Hand, other Hand) int {
	if kindPowers[hand.kind] > kindPowers[other.kind] {
		return 1
	}
	if kindPowers[hand.kind] < kindPowers[other.kind] {
		return -1
	}

	for i := 0; i < len(hand.cards); i++ {
		if cardPowers[string(hand.cards[i])] == cardPowers[string(other.cards[i])] {
			continue
		}
		if cardPowers[string(hand.cards[i])] > cardPowers[string(other.cards[i])] {
			return 1
		}
		if cardPowers[string(hand.cards[i])] < cardPowers[string(other.cards[i])] {
			return -1
		}
	}

	return 0
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
