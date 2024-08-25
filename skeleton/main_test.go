package main

import (
    "testing"
    "cmp"
    "slices"
)

type Input struct {
    Initial  any
    Expected any
    Answer any
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

func TestExample(t *testing.T) {
    input := Input{
    }
    result := problem(input.Initial)
    validateSliceTestResults(t, input, input.Expected, result)
}
