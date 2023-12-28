package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards    [5]rune
	Bid      int
	HandType HandType
}

func cardValue(card string, wildcard bool) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		if wildcard {
			return 1
		} else {
			return 11
		}
	case "T":
		return 10
	default:
		return int(card[0] - '0')
	}
}

func categorizeHand(hand string, wildcard bool) HandType {
	counts := make(map[rune]int)
	for _, card := range hand {
		counts[card]++
	}

	j_count := counts[rune('J')]
	pairs, threes, fours, fives := 0, 0, 0, 0

	for _, count := range counts {
		switch count {
		case 2:
			pairs++
		case 3:
			threes++
		case 4:
			fours++
		case 5:
			fives++
		}
	}
	if !wildcard {
		switch {
		case fives == 1:
			return FiveOfAKind
		case fours == 1:
			return FourOfAKind
		case threes == 1 && pairs == 1:
			return FullHouse
		case threes == 1:
			return ThreeOfAKind
		case pairs == 2:
			return TwoPair
		case pairs == 1:
			return OnePair
		default:
			return HighCard
		}
	} else {
		switch {
		case fives == 1 || (fours == 1 && j_count >= 1) || (threes == 1 && j_count == 2) || (pairs == 1 && j_count == 3):
			return FiveOfAKind
		case fours == 1 || (threes == 1 && j_count >= 1) || (pairs == 2 && j_count >= 2):
			return FourOfAKind
		case (threes == 1 && pairs == 1) || (pairs == 2 && j_count == 1):
			return FullHouse
		case threes == 1 || (pairs >= 1 && j_count >= 1):
			return ThreeOfAKind
		case pairs == 2:
			return TwoPair
		case pairs == 1 || j_count >= 1:
			return OnePair
		default:
			return HighCard
		}
	}
}

func parseInput(raw_input string, wildcard bool) []Hand {
	var hands []Hand
	lines := strings.Split(raw_input, "\n")

	for _, line := range lines {
		var hand Hand
		parts := strings.Split(line, " ")
		hand.Bid, _ = strconv.Atoi(parts[1])
		cards := []rune(parts[0])
		copy(hand.Cards[:], cards)
		hand.HandType = categorizeHand(parts[0], wildcard)
		hands = append(hands, hand)
	}
	return hands
}

func compareHand(hand1 Hand, hand2 Hand, wildcard bool) bool {
	if hand1.HandType != hand2.HandType {
		return hand1.HandType > hand2.HandType
	}
	for i := range hand1.Cards {
		h1C := cardValue(string(hand1.Cards[i]), wildcard)
		h2C := cardValue(string(hand2.Cards[i]), wildcard)
		if h1C != h2C {
			return h1C > h2C
		}
	}
	return true
}

func Main(raw_input string) {
	hands := parseInput(raw_input, false)
	sort.Slice(hands, func(i, j int) bool {
		return compareHand(hands[j], hands[i], false)
	})
	s1 := 0
	for i, hand := range hands {
		s1 += hand.Bid * (i + 1)
	}
	fmt.Println("Solution:")
	fmt.Println("S1: ", s1)
	hands = parseInput(raw_input, true)
	sort.Slice(hands, func(i, j int) bool {
		return compareHand(hands[j], hands[i], true)
	})
	s2 := 0
	for i, hand := range hands {
		s2 += hand.Bid * (i + 1)
	}
	fmt.Println("S2: ", s2)
}
