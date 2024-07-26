package main

import (
    "testing"
    "cmp"
    "slices"
)

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

func TestSortSimple(t *testing.T) {
    input := []int{3, -1, 2}
    answer := []int{-1, 2, 3}
    result := sortArray(input)
    validateSliceTestResults(t, input, answer, result)
}

func TestSortFirstFailure(t *testing.T) {
    input := []int{5, 2, 3, 1}
    answer := []int{1, 2, 3, 5}
    result := sortArray(input)
    validateSliceTestResults(t, input, answer, result)
}
