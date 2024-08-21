package main

import (
    "testing"
    "cmp"
    "slices"
)

type Input[T ~[]comparable | ~comparable] struct {
    Initial  T
    Expected T
    Answer T
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

func TestExample(t *testing.T) {
    input := Input{
    }
    result := problem(input.Initial)
    validateSliceTestResults(t, input, input.Expected, result)
}
