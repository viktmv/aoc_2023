package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var inputFile = "input.txt"

type Round struct {
	Red   int
	Blue  int
	Green int
}

type Game struct {
	Id       string
	RedMax   int
	BlueMax  int
	GreenMax int
	Rounds   []Round
}

var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatalf("errored with %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var games []Game

	for scanner.Scan() {
		line := scanner.Text()
		game := ParseGame(line)
		games = append(games, game)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var ids []string
	var powers []int
	for _, game := range games {
		shouldWork := true
		for _, round := range game.Rounds {
			if round.Green > limits["green"] ||
				round.Blue > limits["blue"] ||
				round.Red > limits["red"] {
				shouldWork = false
			}
		}
		if shouldWork {
			ids = append(ids, game.Id)
		}

		power := game.RedMax * game.BlueMax * game.GreenMax
		powers = append(powers, power)
	}

	var sum int
	for _, id := range ids {
		idInt, err := strconv.Atoi(id)

		if err != nil {
			log.Fatal("could not convert game id to int")
		}

		sum += idInt
	}

	var powerSum int
	for _, power := range powers {
		powerSum += power
	}

	fmt.Println(ids)
	fmt.Println(sum)
	fmt.Println(powerSum)
}

func ParseGame(line string) Game {
	gameString, roundsString, _ := strings.Cut(line, ":")
	_, gameID, _ := strings.Cut(gameString, " ")
	var rounds []Round

	game := Game{
		Id: gameID,
	}

	roundsRaw := strings.Split(roundsString, ";")
	for i := 0; i < len(roundsRaw); i++ {
		round := ParseRound(roundsRaw[i])

		if round.Red > game.RedMax {
			game.RedMax = round.Red
		}
		if round.Blue > game.BlueMax {
			game.BlueMax = round.Blue
		}
		if round.Green > game.GreenMax {
			game.GreenMax = round.Green
		}
		if round.Red > game.RedMax {
			game.RedMax = round.Red
		}
		rounds = append(rounds, round)
	}

	game.Rounds = rounds
	return game
}

func ParseRound(roundString string) Round {
	roundCubes := strings.Split(roundString, ",")
	round := Round{}

	for _, cube := range roundCubes {
		cube = strings.Trim(cube, " ")
		num, colour, _ := strings.Cut(cube, " ")
		numInt, err := strconv.Atoi(num)

		if err != nil {
			log.Fatal("could not get cube number")
		}

		switch colour {
		case "red":
			round.Red = numInt
		case "green":
			round.Green = numInt
		case "blue":
			round.Blue = numInt
		}
	}

	return round
}
