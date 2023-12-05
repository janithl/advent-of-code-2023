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
	symbolMatcher := regexp.MustCompile("[^0-9.]")

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
	total := 0
	totalNums := 0
	symbolNums := 0

	for i := 0; i < len(input); i++ {
		nums := numMatcher.FindAllIndex([]byte(input[i]), -1)
		totalNums += len(nums)

		if(len(nums) > 0) {
			for _, num := range nums {
				if findSymbolsNearby(symbols, num[0], num[1] - 1, i) {
					validNum, err := strconv.Atoi(input[i][num[0]:num[1]])
					if err == nil {
						total += validNum
						symbolNums += 1
					}
				}
			}
		}
	}

	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Selected nums: %d / %d\n", symbolNums, totalNums)
}

func findSymbolsNearby(symbols [][]bool, startX int, endX int, y int) bool {
	for i := startX - 1; i <= endX + 1; i++ {
		for j := y - 1; j <= y + 1; j++ {
			if isSymbolNearby(symbols, i, j) {
				return true
			}
		}
	}
	return false
}

func isSymbolNearby(symbols [][]bool, x int, y int) bool {
	if x < 0 || y < 0 || x > len(symbols) - 1 || y > len(symbols) - 1 {
		return false
	}

	return symbols[x][y]
}