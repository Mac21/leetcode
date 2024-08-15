package main

import (
	"cmp"
	"slices"
	"testing"
)

// Fill with example inputs
type Input struct {
	n int
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
		t.Fatalf("Test %v success: %v? Expected %#v got %#v\n", input, slices.Compare(answer, result) == 0, answer, result)
	}
}

func TestLeetcodeExample1(t *testing.T) {
	answer := []string{"()"}
    n := 1
	result := generateParenthesis(n)
	validateSliceTestResults(t, n, answer, result)
}

func TestLeetcodeExampleN2(t *testing.T) {
	answer := []string{"()()", "(())"}
    n := 2
	result := generateParenthesis(n)
	validateSliceTestResults(t, n, answer, result)
}

func TestLeetcodeExample2(t *testing.T) {
	answer := []string{"((()))","(()())","(())()","()(())","()()()",}
    n := 3
	result := generateParenthesis(n)
	validateSliceTestResults(t, n, answer, result)
}
