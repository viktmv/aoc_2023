package main

import (
	"testing"
)

var testCases = map[string]int{
	"1abc2":            12,
	"pqr3stu8vwx":      38,
	"a1b2c3d4e5f":      15,
	"treb7uchet":       77,
	"two1nine":         29,
	"eightwothree":     83,
	"abcone2threexyz":  13,
	"xtwone3four":      24,
	"4nineeightseven2": 42,
	"zoneight234":      14,
	"7pqrstsixteen":    76,
}

func TestReadLine(t *testing.T) {
	for k, v := range testCases {
		num := ReadLine(k)
		if num != v {
			t.Errorf("readLine(%s) = %d, want %d", k, num, v)
		}
	}
}
