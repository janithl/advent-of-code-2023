package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
    "regexp"
	"strings"
)

type CardVal struct {
	Name string
	Wins int
	Instances int
}

func main() {
	numMatcher := regexp.MustCompile("[0-9]+")
	scanner := bufio.NewScanner(os.Stdin)
	total := 0.0

	var cards []CardVal

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		total += part1(text, numMatcher)
		cards = part2(text, numMatcher, cards)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Part 1 total:", total)

	// part 2
	totalCards := 0
	cards = setInstances(cards)
	fmt.Println(cards)

	for _, card := range(cards) {
		totalCards += card.Instances
	}
	fmt.Println("Part 2 total:", totalCards)
}

func part1(text string, numMatcher *regexp.Regexp) float64 {
	cardComp := strings.Split(text, ": ")
	cardNumsets := strings.Split(cardComp[1], " | ")
	
	
	winningNums := numMatcher.FindAllString(cardNumsets[0], -1)
	ourNums := numMatcher.FindAllString(cardNumsets[1], -1)
	wins := 0.0
	for i := 0; i < len(ourNums); i++ {
		for j := 0; j < len(winningNums); j++ {
			if ourNums[i] == winningNums[j] {
				wins += 1.0
			}
		}
	}

	total := 0.0
	if wins > 0 {
		cardValue := math.Pow(2, wins - 1.0)
		fmt.Println("Card ", cardComp[0], " is valued at ", cardValue)
		total += cardValue
	}
	return total
}

func part2(text string, numMatcher *regexp.Regexp, cards []CardVal) []CardVal {
	cardComp := strings.Split(text, ": ")
	cardNumsets := strings.Split(cardComp[1], " | ")

	winningNums := numMatcher.FindAllString(cardNumsets[0], -1)
	ourNums := numMatcher.FindAllString(cardNumsets[1], -1)
	wins := 0
	for i := 0; i < len(ourNums); i++ {
		for j := 0; j < len(winningNums); j++ {
			if ourNums[i] == winningNums[j] {
				wins += 1
			}
		}
	}
	
	return append(cards, CardVal{cardComp[0], wins, 1})
}

func setInstances(cards []CardVal) []CardVal {
	for i, card := range(cards) {
		for k := 0; k < card.Instances; k++ {
			for j := i + 1; j < len(cards); j++ {
				if j <= i + card.Wins {
					cards[j].Instances += 1
				}
			}
		}
	}
	return cards
}