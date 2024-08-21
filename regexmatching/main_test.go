package main

import (
    "testing"
    "cmp"
    "slices"
)

type Input struct {
    Intial string
    Pattern string
    Answer bool
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
        Intial: "aa",
        Pattern: "a",
        Answer: false,
    }
    result := isMatch(input.Intial, input.Pattern)
    validateTestResults(t, input, input.Answer, result)
}

func TestCase2(t *testing.T) {
    input := Input{
        Intial: "aa",
        Pattern: "a*",
        Answer: true,
    }
    result := isMatch(input.Intial, input.Pattern)
    validateTestResults(t, input, input.Answer, result)
}

func TestCase3(t *testing.T) {
    input := Input{
        Intial: "aa",
        Pattern: "a*aa",
        Answer: true,
    }
    result := isMatch(input.Intial, input.Pattern)
    validateTestResults(t, input, input.Answer, result)
}

func TestCase4(t *testing.T) {
    input := Input{
        Intial: "aaaaaaaaaaaaaaaaaaab",
        Pattern: "a*a*a*a*a*a*a*a*a*a*",
        Answer: false,
    }
    result := isMatch(input.Intial, input.Pattern)
    validateTestResults(t, input, input.Answer, result)
}

//func TestCase5(t *testing.T) {
//    input := Input{
//        Intial: "abc",
//        Pattern: "a***abc",
//        Answer: true,
//    }
//    result := isMatch(input.Intial, input.Pattern)
//    validateTestResults(t, input, input.Answer, result)
//}
