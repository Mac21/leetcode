package main

import (
	"cmp"
	"fmt"
	"slices"
    "github.com/mac21/gods"
)

func displayTestResults[T comparable](input any, answer, result T) {
    fmt.Printf("Test: %v success: %v? Expected %v got %v\n\n", input, answer == result, answer, result)
}

func displayOrderedTestResults[T cmp.Ordered](input any, answer, result T) {
	fmt.Printf("Test %v success: %v? Expected %v got %v\n\n", input, cmp.Compare(answer, result) == 0, answer, result)
}

func displaySliceTestResults[T ~[]E, E cmp.Ordered](input any, answer, result T) {
	fmt.Printf("Test %v success: %v? Expected %v got %v\n\n", input, slices.Compare(answer, result) == 0, answer, result)
}

/*
Problem description:

Constraints: 
*/
func main() {
    input := 1
	result := 0
	answer := 1
	displayTestResults(input, answer, result)
}
