package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// should be 79026871

type Grid = [][]string

type Position struct {
	X int
	Y int
}

var Input string = "input.txt"

func main() {
	file, err := os.Open(Input)

	if err != nil {
		log.Fatalf("could not read file, %s", err)
	}

	defer file.Close()

	var grid Grid
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	FindPartsSum(grid)
}

type PartNumber struct {
	Value    string
	Valid    bool
	Position Position
}

type Gear struct {
	Position Position
}

type Symbol struct {
	Value    string
	Position Position
}

func FindPartsSum(grid Grid) int {
	var currentNum string
	var currentNumValid bool
	var parts []PartNumber
	var gears []Gear

	for y, row := range grid {
		for x, v := range row {
			// gear ?
			if v == "*" {
				gears = append(gears, Gear{Position: Position{X: x, Y: y}})
			}

			if currentNum != "" && !checkIfDigit(v) {
				// we're at symbol, reset the current number
				parts = append(parts, PartNumber{
					Value:    currentNum,
					Valid:    currentNumValid,
					Position: Position{X: x - 1, Y: y},
				})
				currentNum = ""
				currentNumValid = false
			} else if checkIfDigit(v) {
				// only check if not yet valid
				if !currentNumValid {
					valid, _ := checkAround(Position{X: x, Y: y}, grid)
					currentNumValid = valid
				}
				currentNum += v
			}
		}
	}

	// part 1
	var sum int
	for _, n := range parts {
		if n.Valid {
			sum += toDigit(&n.Value)
		}
	}

	var powerSum int = 0
	for _, n := range parts {
		if !n.Valid {
			continue
		}
	}

	fmt.Println("part 1")
	fmt.Println(sum)

	// part 2
	for i := range gears {
		one, two, match := findAdjacentParts(&gears[i], parts)
		if match {
			powerSum += toDigit(&one.Value) * toDigit(&two.Value)
		}
	}

	fmt.Println("part 2")
	fmt.Println(powerSum)
	return sum
}

func findAdjacentParts(gear *Gear, parts []PartNumber) (*PartNumber, *PartNumber, bool) {
	var matches []*PartNumber

	for i := range parts {
		part := &parts[i]

		// diff by Y
		if diff(gear.Position.Y, part.Position.Y) <= 1 {
			// slide by X to the left
			// X represent the end of the part number
			// need to check all X points of the part

			for i := 0; i <= len(part.Value)-1; i++ {
				posX := part.Position.X - i

				if diff(gear.Position.X, posX) <= 1 {
					matches = append(matches, part)
					break
				}
			}
		}
	}

	if len(matches) == 2 {
		return matches[0], matches[1], true
	}

	return &PartNumber{}, &PartNumber{}, false
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func checkIfDigit(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}

func toDigit(str *string) int {
	d, err := strconv.Atoi(*str)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

func checkAround(pos Position, grid Grid) (bool, Symbol) {
  for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
      valid, symbol := checkPosition(Position{X: pos.X + x, Y: pos.Y + y}, grid)
      if valid {
        return valid, Symbol{Value: symbol, Position: Position{X: pos.X + x, Y: pos.Y + y}}
      }
    }
  }
	return false, Symbol{Value: "", Position: Position{}}
}

func checkPosition(pos Position, grid Grid) (valid bool, symbol string) {
	if pos.X < 0 || pos.Y < 0 {
		return false, ""
	}
	if pos.X > len(grid[0])-1 {
		return false, ""
	}
	if pos.Y > len(grid)-1 {
		return false, ""
	}

	value := grid[pos.Y][pos.X]
	re := regexp.MustCompile(`\@|\&|\*|\=|\#|\%|\/|\+|\-|\$`)
	match := re.Find([]byte(value))
	if match == nil {
		return false, ""
	}

	return true, string(match)
}
