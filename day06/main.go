package main

import (
	"bufio"
	"fmt"
	"os"
)

type Race struct {
	Time int
	Distance int
}

const length = 4

func main() {
	var title string
	scanner := bufio.NewScanner(os.Stdin)

	races := make([]Race, length)
	texts := make([]int, length)

	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		fmt.Sscanf(text, "%10s%d%d%d%d", &title, &texts[0], &texts[1], &texts[2], &texts[3])
		for j := 0; j < length; j++ {
			if i == 0 {
				races[j].Time = texts[j]
			} else {
				races[j].Distance = texts[j]
			}
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	part1(races)
	part2()
}

func part1(races []Race) {
	totalWaysToWin := 0
	for i := 0; i < len(races); i++ {
		waysToWin := 0
		for j := 0; j < races[i].Time; j++ {
			if races[i].Distance < j * (races[i].Time - j) {
				waysToWin++
			}
		}
		fmt.Println("Race ", i, " ways to win = ", waysToWin)
		if totalWaysToWin == 0 {
			totalWaysToWin = waysToWin
		} else if waysToWin > 0  {
			totalWaysToWin *= waysToWin
		}
	}
	fmt.Println("Ways to win: ", totalWaysToWin)
}

func part2() {
	fmt.Println("Part 2 Test:")
	races := make([]Race, 1)
	races[0] = Race{71530, 940200}
	part1(races)

	fmt.Println("Part 2 Input:")
	races[0] = Race{56717999, 334113513502430}
	part1(races)
}