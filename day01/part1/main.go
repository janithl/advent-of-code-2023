package main

import (
	"bufio"
	"fmt"
	"os"
    "regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := 0

	// regex to match only numbers
	numMatcher := regexp.MustCompile("[a-zA-Z]")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		number := numMatcher.ReplaceAll([]byte(text), []byte(""))
		output := string(number[0]) + string(number[len(number) - 1])

		outputNum, _ := strconv.Atoi(output)
		fmt.Println(text, outputNum)
		total += outputNum
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(total)
}