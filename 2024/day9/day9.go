package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 6412390328451 --> too high
// 15822303544788
// 6412390328451
// 616982068530 --> too low

func main() {
	file, err := os.Open("input_test.txt")

	if err != nil {
		panic("no file?")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	discMap := make([]int, 0, 32)

	for scanner.Scan() {
		digit := scanner.Text()
		if digit != "\n" {
			discMap = append(discMap, mustAtoi(digit))
		}
	}

	memory := make([]int, 0, 32)
	id := 0
	for i, entry := range discMap {
		isFile := i%2 == 0

		if isFile {
			for range entry {
				memory = append(memory, id)
			}
			id++
		} else {
			for range entry {
				memory = append(memory, -1)
			}
		}
	}

	// Part 1
	// compactBlocks(&memory)

	// Part 2
	compactFiles(&memory)
}

func compactFiles(memory *[]int) {
	compacted := make([]int, len(*memory), len(*memory))
	copy(compacted, *memory)
	seen := make(map[int]bool)

	// fmt.Printf("before compacting %v\n", compacted)

	p2 := len(compacted) - 1
	for p2 >= 0 {
		// advance p2 to the end of the next id
		for compacted[p2] == -1 {
			p2--
			continue
		}

		fileSize := 0
		id := compacted[p2]

		// calculate the file size and advance p2
		// to the start of the id block - 1
		for p2 >= 0 && compacted[p2] == id {
			fileSize++
			p2--
		}

		// dont double process
		if v, ok := seen[id]; ok && v {
			continue
		}

		seen[id] = true
		p1 := 0

		// advance to the first free space
		for compacted[p1] != -1 {
			p1++
		}

		freeSpace := 0
		for p1 < p2 {
			// calculate the block of free space available
			if compacted[p1] == -1 {
				freeSpace++
				p1++
				continue
			}

			// try moving the file
			toStart := p1 - freeSpace
			toEnd := p1
			fromStart := p2 + 1
			fromEnd := p2 + fileSize + 1

			if freeSpace != 0 && fileSize <= freeSpace {
				moveFile(
					toStart,
					toEnd,
					fromStart,
					fromEnd,
					fileSize,
					&compacted,
				)
				p1++
			}

			// advance p1 to the next free spot
			for compacted[p1] != -1 {
				p1++
			}

			// reset free space
			freeSpace = 0
		}
	}

	checksum := 0
	for i, v := range compacted {
		if v == -1 {
			continue
		}
		checksum += (i * v)
	}

	fmt.Println(checksum)
}

func moveFile(toStart, toEnd, fromStart, fromEnd, fileSize int, memory *[]int) {
	j := 0
	for i := toStart; i < toEnd; i++ {
		if j >= fileSize {
			break
		}
		j++
		(*memory)[i] = (*memory)[fromStart]
	}
	// clean up old file
	for j := fromStart; j < fromEnd; j++ {
		(*memory)[j] = -1
	}
}

func compactBlocks(memory *[]int) {
	compacted := make([]int, len(*memory), len(*memory))
	copy(compacted, *memory)

	p1 := 0
	p2 := len(compacted) - 1
	for p1 != p2 {
		if compacted[p1] != -1 {
			p1++
			continue
		}
		if compacted[p2] == -1 {
			p2--
			continue
		}
		compacted[p1] = compacted[p2]
		compacted[p2] = -1
		p1++
		p2--
	}

	checksum := 0
	for i, b := range compacted {
		if b == -1 {
			break
		}
		checksum += (i * b)
	}

	fmt.Println(checksum)
}

func mustAtoi(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("could not atoi %s", s))
	}
	return number
}
