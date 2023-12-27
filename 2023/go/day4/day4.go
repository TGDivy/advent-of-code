package day4

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ScratchCard struct {
	WinningNumbers []int
	YourNumbers    []int
	ID             int
}

func parseCard(cardString string) ScratchCard {
	var card ScratchCard
	parts := strings.Split(cardString, ":")
	ID, _ := strconv.Atoi(strings.Trim(parts[0], "Card "))
	card.ID = ID

	numberParts := strings.Split(parts[1], "|")

	for _, v := range strings.Split(numberParts[0], " ") {
		nstr := strings.TrimSpace(v)
		if nstr == "" {
			continue
		}
		num, _ := strconv.Atoi(nstr)
		card.WinningNumbers = append(card.WinningNumbers, num)
	}

	for _, v := range strings.Split(numberParts[1], " ") {
		nstr := strings.TrimSpace(v)
		if nstr == "" {
			continue
		}
		num, _ := strconv.Atoi(nstr)
		card.YourNumbers = append(card.YourNumbers, num)
	}
	return card
}

func cardMatches(card ScratchCard) int {
	count := 0
	for _, num := range card.YourNumbers {
		for _, num2 := range card.WinningNumbers {
			if num2 == num {
				count += 1
				break
			}
		}
	}

	return count
}

func cardPoints(card ScratchCard) int {
	count := cardMatches(card)
	if count == 0 {
		return 0
	}
	val := int(math.Pow(2, float64(count-1)))
	return val
}

func findTotalCards(cards []ScratchCard) int {
	l := len(cards)

	m := make(map[int]int)
	for i := 0; i < l; i++ {
		m[i] = 1
	}

	for i := 0; i < l; i++ {
		count := cardMatches(cards[i])
		for j := 1; j <= count && (i+j) < l; j++ {
			m[i+j] += m[i]
		}
	}
	// fmt.Println(m)
	total := 0
	for _, v := range m {
		total += v
	}
	return total
}

func Main(raw_input string) {
	sum1 := 0
	var cards []ScratchCard
	for _, cardString := range strings.Split(raw_input, "\n") {
		card := parseCard(cardString)
		points := cardPoints(card)
		sum1 += points
		fmt.Println(card, points)
		cards = append(cards, card)
	}

	sum2 := findTotalCards(cards)

	fmt.Println("Solution:")
	fmt.Println("1: ", sum1)
	fmt.Println("2: ", sum2)
}
