package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	const maxRed = 12
	const maxGreen = 13
	const maxBlue = 14

	// regex to match only letters
	letterMatcher := regexp.MustCompile("[a-zA-Z\\s]")
	// refex to match coloured cubes
	redMatcher := regexp.MustCompile("([0-9]+) red")
	greenMatcher := regexp.MustCompile("([0-9]+) green")
	blueMatcher := regexp.MustCompile("([0-9]+) blue")

	var gameIds []int
	var impossibleGameIds []int

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		lineComp1 := strings.Split(text, ": ")
		gameIdText := letterMatcher.ReplaceAll([]byte(lineComp1[0]), []byte(""))
		gameId, err := strconv.Atoi(string(gameIdText))
		if err != nil {
			fmt.Println("Error:", err)
		}

		gameResSet := strings.Split(lineComp1[1], "; ")
		gameIds = append(gameIds, gameId)

		for _, gameRes := range gameResSet {
			redMatches := redMatcher.FindStringSubmatch(gameRes)
			if len(redMatches) > 1 {
				redMatchNum, matchErr := strconv.Atoi(string(redMatches[1]))
				if matchErr == nil && redMatchNum > maxRed {
					impossibleGameIds = append(impossibleGameIds, gameId)
				}
			}

			greenMatches := greenMatcher.FindStringSubmatch(gameRes)
			if len(greenMatches) > 1 {
				greenMatchNum, matchErr := strconv.Atoi(string(greenMatches[1]))
				if matchErr == nil && greenMatchNum > maxGreen {
					impossibleGameIds = append(impossibleGameIds, gameId)
				}
			}

			blueMatches := blueMatcher.FindStringSubmatch(gameRes)
			if len(blueMatches) > 1 {
				blueMatchNum, matchErr := strconv.Atoi(string(blueMatches[1]))
				if matchErr == nil && blueMatchNum > maxBlue {
					impossibleGameIds = append(impossibleGameIds, gameId)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	total := 0
	for _, game := range gameIds {
		possible := true
		for _, impossible := range impossibleGameIds {
			if game == impossible {
				possible = false
			}
		}
		if possible {
			fmt.Printf("Game ID %d is possible.\n", game)
			total += game
		}
	}

	fmt.Printf("Total of possible game ids: %d\n", total)
}