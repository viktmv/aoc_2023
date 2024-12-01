package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tiles = map[string][]int{
	"|": make([]int, 0),
	"-": make([]int, 0),
	"L": make([]int, 0),
	"J": make([]int, 0),
	"7": make([]int, 0),
	"F": make([]int, 0),
	".": make([]int, 0),
	"S": make([]int, 0),
}

type Tile struct {
	x        int
	y        int
	value    string
	visited  bool
	distance int
}

var grid [][]Tile

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var start Tile
	for scanner.Scan() {
		row := parseLine(scanner.Text())
		tileRow := make([]Tile, 0)
		for i, v := range row {
			tile := Tile{x: i, y: len(grid), value: v}
			if v == "S" {
				start = tile
			}
			tileRow = append(tileRow, tile)
		}
		grid = append(grid, tileRow)
	}

	start.visited = true
	start.distance = 0

	tiles := []*Tile{}
	tiles = append(tiles, findValidNeighbours(&start)...)
	maxDistance := 0
	var current *Tile

	for len(tiles) > 0 {
		current, tiles = tiles[0], tiles[1:]
		current.visited = true
		if current.distance > maxDistance {
			maxDistance = current.distance
		}
		fmt.Printf("current: %+v\n", current)
		neighbours := findValidNeighbours(current)
		tiles = append(tiles, neighbours...)
	}

	fmt.Println(maxDistance)
}

func findValidNeighbours(start *Tile) []*Tile {
	neighbors := getNeighbors(start.x, start.y)
	validNeighbours := make([]*Tile, 0)
	for _, n := range neighbors {
		if n.value == "." || n.visited {
			continue
		}

		if n.x == start.x {
			if n.y > start.y {
				// is below
				if n.value == "|" || n.value == "L" || n.value == "J" {
					validNeighbours = append(validNeighbours, n)
				}
			} else {
				// is above
				if n.value == "|" || n.value == "7" || n.value == "F" {
					validNeighbours = append(validNeighbours, n)
				}
			}
		}

		if n.y == start.y {
			if n.x > start.x {
				// is right
				if n.value == "-" || n.value == "7" || n.value == "J" {
					validNeighbours = append(validNeighbours, n)
				}
			} else {
				// is left
				if n.value == "-" || n.value == "L" || n.value == "F" {
					validNeighbours = append(validNeighbours, n)
				}
			}
		}
	}

	for _, n := range validNeighbours {
		n.distance = start.distance + 1
	}

	return validNeighbours
}

func getNeighbors(x, y int) []*Tile {
	var neighbors []*Tile
	if x > 0 {
		neighbors = append(neighbors, getTile(x-1, y))
	}
	if x < len(grid[y])-1 {
		neighbors = append(neighbors, getTile(x+1, y))
	}
	if y > 0 {
		neighbors = append(neighbors, getTile(x, y-1))
	}
	if y < len(grid)-1 {
		neighbors = append(neighbors, getTile(x, y+1))
	}
	return neighbors
}

func getTile(x, y int) *Tile {
	return &grid[y][x]
}

func parseLine(line string) []string {
	var list []string
	for _, v := range strings.Split(line, "") {
		list = append(list, v)
	}
	return list
}
