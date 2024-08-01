package main

import (
	"cmp"
	"slices"
	"testing"
)

// Fill with example inputs
type Input struct {
	books      [][]int
	shelfWidth int
}

func validateTestResults[T comparable](t *testing.T, input any, answer, result T) {
	if answer != result {
		t.Fatalf("Test: %v success: %v? Expected %v got %v\n", input, answer == result, answer, result)
	}
}

func validateOrderedTestResults[T cmp.Ordered](t *testing.T, input any, answer, result T) {
	if cmp.Compare(answer, result) != 0 {
		t.Fatalf("Test %v success: %v? Expected %v got %v\n", input, cmp.Compare(answer, result) == 0, answer, result)
	}
}

func validateSliceTestResults[T ~[]E, E cmp.Ordered](t *testing.T, input any, answer, result T) {
	if slices.Compare(answer, result) != 0 {
		t.Fatalf("Test %v success: %v? Expected %v got %v\n", input, slices.Compare(answer, result) == 0, answer, result)
	}
}

func TestLeetcodeExample1(t *testing.T) {
	input := Input{
		books:      [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}},
		shelfWidth: 4,
	}
	answer := 6
	result := minHeightShelves(input.books, input.shelfWidth)
	validateTestResults(t, input, answer, result)
}

func TestLeetcodeExample2(t *testing.T) {
	input := Input{
		books: [][]int{
			{1, 3}, {2, 4}, {3, 2},
		},
		shelfWidth: 6,
	}
	answer := 4
	result := minHeightShelves(input.books, input.shelfWidth)
	validateTestResults(t, input, answer, result)
}

func TestLeetcodeExample3(t *testing.T) {
	input := Input{
		books: [][]int{
			{7, 3}, {8, 7}, {2, 7}, {2, 5},
		},
		shelfWidth: 10,
	}
	answer := 15
	result := minHeightShelves(input.books, input.shelfWidth)
	validateTestResults(t, input, answer, result)
}
