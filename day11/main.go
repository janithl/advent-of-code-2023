package main

import (
	"bufio"
	"fmt"
	"os"
)

type Galaxy struct {
	Index int
	X int
	Y int
}

type Distance struct {
	From int
	To int
	Distance int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var textMatrix []string
	var galaxies []Galaxy

	y := 0
	count := 1
	for scanner.Scan() {
		text := scanner.Text()
	
		if text == "" {
			break
		}

		textMatrix = append(textMatrix, text)
		for x, char := range text {
			if char == '#' {
				galaxies = append(galaxies, Galaxy{count, x, y})
				count++
			}
		}

		y++
	}

	calculatePart1(textMatrix, galaxies)
	calculatePart2(textMatrix, galaxies)
}

func calculatePart1(matrix []string, galaxies []Galaxy) {
	emptyX, emptyY := findEmpties(matrix)
	fmt.Println("emptyX", emptyX)
	fmt.Println("emptyY", emptyY)

	var distances []Distance
	sum := 0
	for i, galaxyA := range galaxies {
		for j, galaxyB := range galaxies {
			if i < j {
				abDistance := calculateDistance(galaxyA.X, galaxyB.X, emptyX, 2) + calculateDistance(galaxyA.Y, galaxyB.Y, emptyY, 2)
				distances = append(distances, Distance{galaxyA.Index, galaxyB.Index, abDistance})
				sum += abDistance
			}
		}
	}

	fmt.Println(distances)
	fmt.Println("Sum of paths:", sum)
}

func calculatePart2(matrix []string, galaxies []Galaxy) {
	emptyX, emptyY := findEmpties(matrix)

	sum := 0
	for i, galaxyA := range galaxies {
		for j, galaxyB := range galaxies {
			if i < j {
				abDistance := calculateDistance(galaxyA.X, galaxyB.X, emptyX, 1000000) + calculateDistance(galaxyA.Y, galaxyB.Y, emptyY, 1000000)
				sum += abDistance
			}
		}
	}

	fmt.Println("Sum of paths:", sum)
}

func calculateDistance(start int, end int, empties []int, expansion int) int {
	expanded := 0
	for _, empty := range empties {
		if (empty < start && empty > end) || (empty > start && empty < end) {
			expanded += expansion - 1
		}
	}
	return abs(end - start) + expanded
}

func findEmpties(matrix []string) ([]int,[]int) {
	var emptyX []int
	var emptyY []int

	for i := 0; i < len(matrix); i++ {
		emptyRow := true
		emptyColumn := true
		for j := 0; j < len(matrix); j++ {
			emptyRow = emptyRow && matrix[i][j] == '.'
			emptyColumn = emptyColumn && matrix[j][i] == '.'
		}
		if emptyRow {
			emptyY = append(emptyY, i)
		}
		if emptyColumn {
			emptyX = append(emptyX, i)
		}
	}

	return emptyX, emptyY
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}