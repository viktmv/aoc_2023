package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	x        int
	y        int
	v        string
	n        int
	isGalaxy bool
	visited  bool
}

type Grid [][]*Point

var grid Grid

const expansionRate int = 2

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var galaxies []*Point
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := parseLine(scanner.Text())
		pointRow := make([]*Point, 0)
		for i, v := range row {
			isGalaxy := v == "#"
			point := Point{
				x:        i,
				y:        len(grid),
				v:        v,
				isGalaxy: isGalaxy,
			}
			if isGalaxy {
				galaxies = append(galaxies, &point)
				point.n = len(galaxies)
			}
			pointRow = append(pointRow, &point)
		}
		grid = append(grid, pointRow)
		// empand row
		if isRowEmpty(pointRow) {
			for i := 0; i < expansionRate-1; i++ {
				newRow := make([]*Point, len(pointRow))
				copy(newRow, pointRow)
				grid = append(grid, newRow)
			}
		}
	}

	// expand columns
	var emptyColumns []int
	for x := range grid[0] {
		isColumnEmpty := true
		for _, row := range grid {
			if row[x].v == "#" {
				isColumnEmpty = false
				break
			}
		}
		if isColumnEmpty {
			emptyColumns = append(emptyColumns, x)
		}
	}

	shift := 0
	for _, idx := range emptyColumns {
		// columns are shifted
		for i := 0; i < expansionRate-1; i++ {
			x := idx + shift
			for y := range grid {
				if len(grid[y]) == x {
					point := Point{x: x, y: y, v: ".", isGalaxy: false}
					grid[y] = append(grid[y], &point)
				} else {
					point := Point{x: x, y: y, v: ".", isGalaxy: false}
					grid[y] = append(grid[y][:x+1], grid[y][x:]...)
					grid[y][x] = &point
				}
			}
			shift++
		}
	}

	// normalize point coordinates after all the shifting around
	for y, row := range grid {
		for x, point := range row {
			point.x = x
			point.y = y
            fmt.Printf("%s", point.v)
		}
        fmt.Println()
	}

	var paths []int
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			paths = append(paths, manhattanDistance(galaxies[i], galaxies[j]))
		}
	}

	var sum int
	for _, p := range paths {
		sum += p
	}

	fmt.Printf("result %d\n", sum)
}

var distances map[Point]map[Point]int

func manhattanDistance(a *Point, b *Point) int {
    diffX := float64(a.x)-float64(b.x)
    diffY := float64(a.y)-float64(b.y)
	return int(math.Abs(diffX) + math.Abs(diffY))
}

func parseLine(line string) []string {
	var list []string
	for _, v := range strings.Split(line, "") {
		list = append(list, v)
	}
	return list
}

func isRowEmpty(row []*Point) bool {
	for _, v := range row {
		if v.isGalaxy {
			return false
		}
	}
	return true
}

func getPoint(x, y int) *Point {
	return grid[y][x]
}

func getNeighbors(x, y int) []*Point {
	var neighbors []*Point
	if x > 0 {
		neighbors = append(neighbors, getPoint(x-1, y))
	}
	if x < len(grid[y])-1 {
		fmt.Printf("got %+v\n", x+1)
		fmt.Printf("got point %+v\n", getPoint(x+1, y))
		neighbors = append(neighbors, getPoint(x+1, y))
	}
	if y > 0 {
		neighbors = append(neighbors, getPoint(x, y-1))
	}
	if y < len(grid)-1 {
		neighbors = append(neighbors, getPoint(x, y+1))
	}
	for _, n := range neighbors {
		fmt.Printf("neighbour %+v\n", n)
	}
	return neighbors
}
