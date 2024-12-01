package main

import (
	"testing"
)

func TestDetermineKind(t *testing.T) {
	output := DetermineKind("77373")
	if output != "full-house" {
		t.Errorf("DetermineKind(77373) = %s; want full-house", output)
	}

	output = DetermineKind("32T3K")
	if output != "one-pair" {
		t.Errorf("DetermineKind(32T3K) = %s; want two-pair", output)
	}

	output = DetermineKind("KK677")
	if output != "two-pair" {
		t.Errorf("DetermineKind(KK677) = %s; want two-pair", output)
	}

	output = DetermineKind("KTJJT")
	if output != "two-pair" {
		t.Errorf("DetermineKind(KTJJT) = %s; want two-pair", output)
	}

	output = DetermineKind("QQQJA")
	if output != "three-of-a-kind" {
		t.Errorf("DetermineKind(QQQJA) = %s; want three-of-a-kind", output)
	}

	output = DetermineKind("T55J5")
	if output != "three-of-a-kind" {
		t.Errorf("DetermineKind(T55J5) = %s; want three-of-a-kind", output)
	}
}

func TestDetermineKindWithJoker(t *testing.T) {
	output := DetermineKindWithJoker("77373")
	if output != "full-house" {
		t.Errorf("DetermineKind(77373) = %s; want full-house", output)
	}

	output = DetermineKindWithJoker("32T3K")
	if output != "one-pair" {
		t.Errorf("DetermineKind(32T3K) = %s; want two-pair", output)
	}

	output = DetermineKindWithJoker("KK677")
	if output != "two-pair" {
		t.Errorf("DetermineKind(KK677) = %s; want two-pair", output)
	}

	output = DetermineKindWithJoker("KTJJT")
	if output != "four-of-a-kind" {
		t.Errorf("DetermineKind(KTJJT) = %s; want four-of-a-kind", output)
	}

	output = DetermineKindWithJoker("QQQJA")
	if output != "four-of-a-kind" {
		t.Errorf("DetermineKind(QQQJA) = %s; want four-of-a-kind", output)
	}

	output = DetermineKindWithJoker("T55J5")
	if output != "four-of-a-kind" {
		t.Errorf("DetermineKind(T55J5) = %s; want four-of-a-kind", output)
	}

	output = DetermineKindWithJoker("9JTQJ")
	if output != "three-of-a-kind" {
		t.Errorf("DetermineKind(9JTQJ) = %s; want three-of-a-kind", output)
	}
    
}
