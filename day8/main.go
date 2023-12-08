package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"day8/lcm"
)

type Node string

var pathMap = map[string][]string{}
var stepToIndex = map[string]int{
	"L": 0,
	"R": 1,
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		if scanner.Text() != "" {
			lines = append(lines, scanner.Text())
		}
	}

	var steps []string
	var nodes []string
	for i, line := range lines {
		if i == 0 {
			steps = strings.Split(line, "")
			continue
		}
		node, next, _ := strings.Cut(line, " = ")
		left, right, _ := strings.Cut(next, ", ")
		left = strings.Trim(left, "(")
		right = strings.Trim(right, ")")
		pathMap[node] = []string{left, right}
		nodes = append(nodes, node)
	}

	// pt 1
	// currentNode := Node("AAA")
	// var count int
	// for i := 0; i < len(steps); i++ {
	//     count++
	//     nextNodes := pathMap[currentNode]
	//     step := steps[i]
	//     currentNode = nextNodes[stepToIndex[step]]
	//     if currentNode == "ZZZ" {
	//         fmt.Println(count)
	//         break
	//     }
	//     if i + 1 == len(steps) {
	//         i = -1
	//     }
	// }

	// pt 2
	var currentNodes []string
	for _, n := range nodes {
		if strings.HasSuffix(n, "A") {
			currentNodes = append(currentNodes, n)
		}
	}

	counts := make([]int, len(currentNodes))
	for j, n := range currentNodes {
		var count int
		currentNode := n
		for i := 0; i < len(steps); i++ {
			count++
			nextNodes := pathMap[currentNode]
			step := steps[i]
			currentNode = nextNodes[stepToIndex[step]]
			if strings.HasSuffix(currentNode, "Z") {
				counts[j] = count
				break
			}
			if i+1 == len(steps) {
				i = -1
			}
		}
	}

	fmt.Println("result:", lcm.FindLCM(counts))
}
