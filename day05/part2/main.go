package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Mapping struct {
	Source int
	Dest int
	Range int
}

type Seed struct {
	Start int
	Length int
}

func main() {
	numMatcher := regexp.MustCompile("[0-9]+")

	scanner := bufio.NewScanner(os.Stdin)
	const mapLength = 8
	var seeds []Seed
	
	maps := make([][]Mapping, mapLength)

	index := 0
	mapNum := 0
	for scanner.Scan() {
		text := scanner.Text()
	
		if mapNum > mapLength {
			break
		}
		
		if index == 0 {
			seedNums := numMatcher.FindAllString(text, -1)
			for i := 0; i < len(seedNums); i += 2 {
				seedStart, _ := strconv.Atoi(seedNums[i])
				seedLength, _ := strconv.Atoi(seedNums[i + 1])
				seeds = append(seeds, Seed{ seedStart, seedLength })
			}
		} else if index > 1 && text == "" {
			mapNum++
		} else {
			mappingRow := numMatcher.FindAllString(text, -1)
			if len(mappingRow) == 3 {
				mapDest, _ := strconv.Atoi(mappingRow[0])
				mapSource, _ := strconv.Atoi(mappingRow[1])
				mapRange, _ := strconv.Atoi(mappingRow[2])
				maps[mapNum] = append(maps[mapNum], Mapping{ mapSource, mapDest, mapRange })
			}
		}

		index++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(maps[0])

	var seedLocations []int
	for index, seed := range seeds {
		fmt.Println("Processing Seed ", seed, " --- ", index + 1, "/", len(seeds))
		for j := seed.Start; j < seed.Start + seed.Length; j++ {
			targetLoc := j
			for _, mapLine := range maps {
				for _, mapItem := range mapLine {
					if targetLoc >= mapItem.Source  && targetLoc < mapItem.Source + mapItem.Range {
						targetLoc = (targetLoc - mapItem.Source) + mapItem.Dest
						break
					}
				}
				// fmt.Print(" > ", targetLoc)
			}
			seedLocations = append(seedLocations, targetLoc)
		}
	}

	fmt.Println("")
	fmt.Println("Seeds: ", seeds)
	// fmt.Println("Locs : ", seedLocations)

	minLoc := 0
	for i, dest := range seedLocations {
		if i == 0 || dest < minLoc {
			minLoc = dest
		}
	}

	fmt.Println("Min Loc : ", minLoc)
}