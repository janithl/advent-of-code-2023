package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestCompare(t *testing.T) {
	testHands := []Hand {
		Hand{"KKKKK", 80},
		Hand{"AAAAA", 100},
		Hand{"AKKAA", 100},
		Hand{"AAKKK", 100},
	}

	slices.SortFunc(testHands, handCompare)
	if testHands[0].Hand != "AAAAA" {
		t.Errorf("Sort Failed: %+v", testHands)
	}

	fmt.Printf("Sorted: %+v", testHands)
}