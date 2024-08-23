package main

import (
    "testing"
    "cmp"
    "slices"
)

type Input struct {
    Initial  int
    Expected string
}

func validateTestResults[T comparable](t *testing.T, input any, answer, result T) {
    if answer != result {
        t.Fatalf("Test: %v Expected %v got %v\n", input, answer, result)
    }
}

func validateOrderedTestResults[T cmp.Ordered](t *testing.T, input any, answer, result T) {
    if cmp.Compare(answer, result) != 0 {
        t.Fatalf("Test %v Expected %v got %v\n", input, answer, result)
    }
}

func validateSliceTestResults[T ~[]E, E cmp.Ordered](t *testing.T, input any, answer, result T) {
    if slices.Compare(answer, result) != 0 {
        t.Fatalf("Test %v Expected %v got %v\n", input, answer, result)
    }
}

func TestCase1(t *testing.T) {
    input := Input{
        Initial: 3749,
        Expected: "MMMDCCXLIX",
    }
    result := intToRoman(input.Initial)
    validateTestResults(t, input, input.Expected, result)
}

func TestCase2(t *testing.T) {
    input := Input{
        Initial: 58,
        Expected: "LVIII",
    }
    result := intToRoman(input.Initial)
    validateTestResults(t, input, input.Expected, result)
}

func TestCase3(t *testing.T) {
    input := Input{
        Initial: 1994,
        Expected: "MCMXCIV",
    }
    result := intToRoman(input.Initial)
    validateTestResults(t, input, input.Expected, result)
}
