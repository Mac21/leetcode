package main

import (
    "testing"
    "cmp"
    "slices"
)

type Input struct {
    Nums []int
    Expected []int
    Answer int
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

func TestCase1(t *testing.T) {
    input := Input{
        Nums: []int{1, 1, 2},
        Expected: []int{1,2},
        Answer: 2,
    }
    removeDuplicates(input.Nums)
    validateSliceTestResults(t, input, input.Expected, input.Nums)
}

func TestCase2(t *testing.T) {
    input := Input{
        Nums: []int{0,0,1,1,1,2,2,3,3,4},
        Expected: []int{0, 1, 2, 3, 4},
        Answer: 4,
    }
    removeDuplicates(input.Nums)
    validateSliceTestResults(t, input, input.Expected, input.Nums)
}
