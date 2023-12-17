package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	numMatcher := regexp.MustCompile("[\\-0-9]+")
	sum1 := 0
	sum2 := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
	
		if text == "" {
			break
		}
		
		numStrs := numMatcher.FindAllString(text, -1)
		fmt.Println("0 :", numStrs)

		nums := make([]int, len(numStrs))
		numsPart2 := make([]int, len(numStrs))
		for i, str := range numStrs {
			num, _ := strconv.Atoi(str)
			nums[i] = num
			numsPart2[len(numStrs) - i - 1] = num
		}

		sum1 += extrapolate(nums, 0)
		sum2 += extrapolate(numsPart2, 0)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("sum part 1:", sum1)
	fmt.Println("sum part 2:", sum2)
}

func extrapolate(line []int, depth int) int {
	nextLine := make([]int, len(line) - 1)

	diff := 0
	zeroDiff := 0

	for i, _ := range nextLine {
		diff = line[i + 1] - line[i]
		nextLine[i] = diff
		if diff == 0 {
			zeroDiff++
		}
	}

	if zeroDiff == len(nextLine) {
		return line[len(line) - 1]
	}

	res := extrapolate(nextLine, depth + 1)
	fmt.Println(depth, ":", res, diff)
	return res + line[len(line) - 1]
}