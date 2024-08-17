package main

import (
    "testing"
    "cmp"
    "slices"
)

type Input struct {
    Nums []int
    Target int
    Answer []int
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

func TestGivenCase1(t *testing.T) {
    input := Input{
        Nums: []int{2,7,11,15},
        Target: 9,
        Answer: []int{0, 1},
    }
    result := twoSum(input.Nums, input.Target)
    validateSliceTestResults(t, input, input.Answer, result)
}

func TestGivenCase2(t *testing.T) {
    input := Input{
        Nums: []int{3,2,4},
        Target: 6,
        Answer: []int{1, 2},
    }
    result := twoSum(input.Nums, input.Target)
    validateSliceTestResults(t, input, input.Answer, result)
}

func TestGivenCase3(t *testing.T) {
    input := Input{
        Nums: []int{3,3},
        Target: 6,
        Answer: []int{0, 1},
    }
    result := twoSum(input.Nums, input.Target)
    validateSliceTestResults(t, input, input.Answer, result)
}
