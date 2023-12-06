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

func main() {
	numMatcher := regexp.MustCompile("[0-9]+")

	scanner := bufio.NewScanner(os.Stdin)
	const mapLength = 8
	var seeds []string
	
	maps := make([][]Mapping, mapLength)

	index := 0
	mapNum := 0
	for scanner.Scan() {
		text := scanner.Text()
	
		if mapNum > mapLength {
			break
		}
		
		if index == 0 {
			seeds = numMatcher.FindAllString(text, -1)
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

	seedLocations := make([]int, len(seeds))
	for i, seed := range seeds {
		fmt.Print("\nSeed ", seed)
		targetLoc, _ := strconv.Atoi(seed)
		for _, mapLine := range maps {
			for _, mapItem := range mapLine {
				if targetLoc >= mapItem.Source  && targetLoc < mapItem.Source + mapItem.Range {
					targetLoc = (targetLoc - mapItem.Source) + mapItem.Dest
					break
				}
			}
			fmt.Print(" > ", targetLoc)
		}
		seedLocations[i] = targetLoc
	}

	fmt.Println("")
	fmt.Println("Seeds: ", seeds)
	fmt.Println("Locs : ", seedLocations)

	minLoc := 0
	for i, dest := range seedLocations {
		if i == 0 || dest < minLoc {
			minLoc = dest
		}
	}

	fmt.Println("Min Loc : ", minLoc)
}