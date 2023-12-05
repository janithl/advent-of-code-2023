package main

import (
	"bufio"
	"fmt"
	"os"
    "regexp"
	"strconv"
	"strings"
)

type TotalHolder struct {
	Count int
	Total int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	index := 0
	symbolCount := 0

	// matrix to hold stuff
	const n = 150
	var inputSlice []string

	symbolMatrix := make([][]bool, n)
	for i := 0; i < n; i++ {
		symbolMatrix[i] = make([]bool, n)
	}


	// regex to match symbols
	symbolMatcher := regexp.MustCompile("[*]")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		inputSlice = append(inputSlice, text)
		// find symbols and store them
		matches := symbolMatcher.FindAllIndex([]byte(strings.TrimSpace(text)), -1)
		for _, match := range matches {
			symbolMatrix[match[0]][index] = true
			symbolCount += 1
		}

		index += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Total symbols: ", symbolCount)
	findNearbyTotal(symbolMatrix, inputSlice)
}

func findNearbyTotal(symbols [][]bool, input []string) {
	numMatcher := regexp.MustCompile("[0-9]+")

	n := len(symbols)
	totalMatrix := make([][]TotalHolder, n)
	for i := 0; i < n; i++ {
		totalMatrix[i] = make([]TotalHolder, n)
	}
	

	for i := 0; i < len(input); i++ {
		nums := numMatcher.FindAllIndex([]byte(input[i]), -1)

		if(len(nums) > 0) {
			for _, num := range nums {
				if coords := findSymbolsNearby(symbols, num[0], num[1] - 1, i); coords != "" {
					coordXY := strings.Split(coords, ",")
					validNum, err := strconv.Atoi(input[i][num[0]:num[1]])
					if err == nil && len(coordXY) == 2 {
						coordX, _ := strconv.Atoi(coordXY[0])
						coordY, _ := strconv.Atoi(coordXY[1])
						total := 0
						if(totalMatrix[coordX][coordY].Total == 0) {
							total = validNum
						} else {
							total = totalMatrix[coordX][coordY].Total * validNum
						}

						totalMatrix[coordX][coordY] = TotalHolder{totalMatrix[coordX][coordY].Count + 1, total}
					}
				}
			}
		}
	}

	finalTotal := 0
	symbolCount := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if(totalMatrix[i][j].Count == 2) {
				finalTotal += totalMatrix[i][j].Total
				symbolCount++
			}
		}
	}

	fmt.Println("Total: ", finalTotal, " with ", symbolCount, " valid gears")
}

func findSymbolsNearby(symbols [][]bool, startX int, endX int, y int) string {
	for i := startX - 1; i <= endX + 1; i++ {
		for j := y - 1; j <= y + 1; j++ {
			if coords := isSymbolNearby(symbols, i, j); coords != "" {
				return coords;
			}
		}
	}
	return ""
}

func isSymbolNearby(symbols [][]bool, x int, y int) string {
	if x < 0 || y < 0 || x > len(symbols) - 1 || y > len(symbols) - 1 {
		return ""
	}

	if symbols[x][y] {
		return fmt.Sprintf("%d,%d", x, y)
	}
	return ""
}