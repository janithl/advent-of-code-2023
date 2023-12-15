package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var cards = [...]byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2', '1'}

func isFullHouse(hand string) bool {
	for _, c := range []byte(hand) {
		if c != hand[0] {
			return false
		}
	}

	return true
}

type Hand struct {
	Hand string
	Bid int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var hands []Hand

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		var hand string
		var bid int

		fmt.Sscanf(text, "%5s%d", &hand, &bid)
		hands = append(hands, Hand{ hand, bid })
	}

	slices.SortFunc(hands, handCompare)

	fmt.Println(hands)
}

func handCompare(a, b Hand) int {
	if isFullHouse(a.Hand) {
		if isFullHouse(b.Hand) {
			return getStrength(b.Hand) - getStrength(a.Hand)
		}
		return -1
	}
	if isFullHouse(b.Hand) {
		return 1
	}
	return getStrength(b.Hand) - getStrength(a.Hand)
}

func getStrength(hand string) int {
	strength := 0
	for _, c := range []byte(hand) {
		for i := range cards {
			if cards[i] == c {
				strength += len(cards) - i
				break
			}
		}
	}
	
	return strength
}