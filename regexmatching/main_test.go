package main

import (
	"cmp"
	"slices"
	"testing"
)

type Input struct {
	Intial  string
	Pattern string
	Answer  bool
}

func validateTestResults[T comparable](t *testing.T, input any, answer, result T) {
	if answer != result {
		t.Fatalf("Test: %v expected %v got %v\n", input, answer, result)
	}
}

func validateOrderedTestResults[T cmp.Ordered](t *testing.T, input any, answer, result T) {
	if cmp.Compare(answer, result) != 0 {
		t.Fatalf("Test %v expected %v got %v\n", input, answer, result)
	}
}

func validateSliceTestResults[T ~[]E, E cmp.Ordered](t *testing.T, input any, answer, result T) {
	if slices.Compare(answer, result) != 0 {
		t.Fatalf("Test %v expected %v got %v\n", input, answer, result)
	}
}

func TestCase1(t *testing.T) {
	input := Input{
		Intial:  "aa",
		Pattern: "a",
		Answer:  false,
	}
	result := isMatch(input.Intial, input.Pattern)
	validateTestResults(t, input, input.Answer, result)
}

func TestCase2(t *testing.T) {
	input := Input{
		Intial:  "aa",
		Pattern: "a*",
		Answer:  true,
	}
	result := isMatch(input.Intial, input.Pattern)
	validateTestResults(t, input, input.Answer, result)
}

func TestCase3(t *testing.T) {
	input := Input{
		Intial:  "aa",
		Pattern: "a*aa",
		Answer:  true,
	}
	result := isMatch(input.Intial, input.Pattern)
	validateTestResults(t, input, input.Answer, result)
}

func TestCase4(t *testing.T) {
	input := Input{
		Intial:  "aaaaaaaaaaaaaaaaaaab",
		Pattern: "a*a*a*a*a*a*a*a*a*a*",
		Answer:  false,
	}
	result := isMatch(input.Intial, input.Pattern)
	validateTestResults(t, input, input.Answer, result)
}

func TestCase5(t *testing.T) {
	input := Input{
		Intial:  "abc",
		Pattern: "a***abc",
		Answer:  true,
	}
	result := isMatch(input.Intial, input.Pattern)
	validateTestResults(t, input, input.Answer, result)
}

func TestCase6(t *testing.T) {
	input := Input{
		Intial:  "aab",
		Pattern: "c*a*b",
		Answer:  true,
	}
	result := isMatch(input.Intial, input.Pattern)
	validateTestResults(t, input, input.Answer, result)
}

func TestCase7(t *testing.T) {
	input := Input{
		Intial:  "aa",
		Pattern: "aa",
		Answer:  true,
	}
	result := isMatch(input.Intial, input.Pattern)
	validateTestResults(t, input, input.Answer, result)
}

func TestCase8(t *testing.T) {
	input := Input{
		Intial:  "aaa",
		Pattern: "aaa",
		Answer:  true,
	}
	result := isMatch(input.Intial, input.Pattern)
	validateTestResults(t, input, input.Answer, result)
}

func BenchmarkCase4(b *testing.B) {
	for range b.N {
		isMatch("aaaaaaaaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*")
	}
}

func BenchmarkCase5(b *testing.B) {
	for range b.N {
		isMatch("abc", "a***abc")
	}
}

func BenchmarkCase6(b *testing.B) {
	for range b.N {
		isMatch("aab", "c*a*b")
	}
}
