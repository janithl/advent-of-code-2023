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

	// regex to match only letters
	letterMatcher := regexp.MustCompile("[a-zA-Z\\s]")
	// refex to match coloured cubes
	redMatcher := regexp.MustCompile("([0-9]+) red")
	greenMatcher := regexp.MustCompile("([0-9]+) green")
	blueMatcher := regexp.MustCompile("([0-9]+) blue")

	var gameIds []int
	total := 0

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

		minRed := 0
		minGreen := 0
		minBlue := 0

		for _, gameRes := range gameResSet {
			redMatches := redMatcher.FindStringSubmatch(gameRes)
			if len(redMatches) > 1 {
				redMatchNum, matchErr := strconv.Atoi(string(redMatches[1]))
				if matchErr == nil && redMatchNum > minRed {
					minRed = redMatchNum
				}
			}

			greenMatches := greenMatcher.FindStringSubmatch(gameRes)
			if len(greenMatches) > 1 {
				greenMatchNum, matchErr := strconv.Atoi(string(greenMatches[1]))
				if matchErr == nil && greenMatchNum > minGreen {
					minGreen = greenMatchNum
				}
			}

			blueMatches := blueMatcher.FindStringSubmatch(gameRes)
			if len(blueMatches) > 1 {
				blueMatchNum, matchErr := strconv.Atoi(string(blueMatches[1]))
				if matchErr == nil && blueMatchNum > minBlue {
					minBlue = blueMatchNum
				}
			}
		}

		power := minRed * minGreen * minBlue
		total += power
		fmt.Printf("Game ID %02d: Min %02d red, %02d green, %02d blue = Power %d\n", gameId, minRed, minGreen, minBlue, power)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Total is %d\n", total)
}