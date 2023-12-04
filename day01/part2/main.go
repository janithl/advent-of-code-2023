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
	total := 0

	// regex to match only letters
	letterMatcher := regexp.MustCompile("[a-zA-Z]")
	// replace words for numbers with actual numbers, and leave the end chars
	// so that other partial number names can be matched
	replacer := strings.NewReplacer(
		"one", "o1e",
		"two", "t2o",
		"three", "t3e",
		"four", "f4r",
		"five", "f5e",
		"six", "s6x",
		"seven", "s7n",
		"eight", "e8t",
		"nine", "n9e",
	)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		// run the replacer twice
		text = replacer.Replace(text)
		text = replacer.Replace(text)
		number := letterMatcher.ReplaceAll([]byte(text), []byte(""))
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