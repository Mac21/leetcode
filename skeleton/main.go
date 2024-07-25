package main

import (
	"cmp"
	"fmt"
	"slices"
    "github.com/mac21/gods"
)

func displayTestResults[T cmp.Ordered](answer, result T) {
	fmt.Printf("Test success: %v? Expected %v got %v\n\n", cmp.Compare(answer, result) == 0, answer, result)
}

func displaySliceTestResults[T ~[]E, E cmp.Ordered](answer, result T) {
	fmt.Printf("Test success: %v? Expected %v got %v\n\n", slices.Compare(answer, result) == 0, answer, result)
}

func main() {
	result := 0
	answer := 1
	displayTestResults(answer, result)
}
