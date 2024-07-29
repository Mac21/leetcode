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

func TestExample1(t *testing.T) {
    input := []int{2,5,3,4,1}
    answer := 3
    result := numTeams(input)
    validateTestResults(t, input, answer, result)
}

func TestExample2(t *testing.T) {
    input := []int{2, 1, 3}
    answer := 0
    result := numTeams(input)
    validateTestResults(t, input, answer, result)
}

func TestExample3(t *testing.T) {
    input := []int{1,2,3,4}
    answer := 4
    result := numTeams(input)
    validateTestResults(t, input, answer, result)
}
