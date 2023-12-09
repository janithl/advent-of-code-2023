package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Element struct {
	Loc string
	Left string
	Right string
}

func main() {
	var instructions, loc, ldest, rdest string
	var elems []Element

	scanner := bufio.NewScanner(os.Stdin)

	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" && i > 1 {
			break
		}

		if i == 0 {
			instructions = text
		} else if i > 1 {
			fmt.Sscanf(text, "%3s = (%3s, %3s)", &loc, &ldest, &rdest)
			elems = append(elems, Element{ loc, ldest, rdest })
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Instructions :", instructions)
	fmt.Println("Elements     :", elems)
	part1(elems, instructions)
	part2(elems, instructions)
}

func part1(elems []Element, instructions string) {
	steps := 0
	elem, count := lookup(elems, "AAA")

	for elem.Loc != "ZZZ" {
		for _, ins := range instructions {
			if ins == 'L' {
				elem, count = lookup(elems, elem.Left)
			} else {
				elem, count = lookup(elems, elem.Right)
			}

			if count == 0 {
				return
			}

			steps++
		}
	}

	fmt.Println("Steps     :", steps)
}

func lookup(elems []Element, loc string) (Element, int) {
	for _, el := range elems {
		if el.Loc == loc {
			return el, 1
		}
	}

	return Element{}, 0
}

func part2(elems []Element, instructions string) {
	steps := 0
	lookupElems := lookupPartial(elems, "A")
	fmt.Println("Looked up partial:", lookupElems)

	for !allElemsReached(lookupElems) {
		for _, ins := range instructions {
			for i := 0; i < len(lookupElems); i++ {
				elem := lookupElems[i]
				var count int
				if ins == 'L' {
					elem, count = lookup(elems, elem.Left)
				} else {
					elem, count = lookup(elems, elem.Right)
				}

				if count == 0 {
					fmt.Println("Failed to lookup:", string(ins))
					return
				}

				lookupElems[i] = elem
			}
			steps++
		}
	}

	fmt.Println("Steps      :", steps)
}

func lookupPartial(elems []Element, loc string) []Element {
	var results []Element
	for _, el := range elems {
		if strings.HasSuffix(el.Loc, loc) {
			results = append(results, el)
		}
	}

	return results
}

func allElemsReached(lookupElems []Element) bool {
	endsWithZ := lookupPartial(lookupElems, "Z")
	return len(endsWithZ) == len(lookupElems)
}