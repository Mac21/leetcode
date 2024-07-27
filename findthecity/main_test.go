package main

import (
	"cmp"
	"slices"
	"testing"
)

type TestInput struct {
	N                 int
	Edges             [][]int
	DistanceThreshold int
}

func validateTestResults[T comparable](t *testing.T, input any, answer, result T) {
	if answer != result {
		t.Fatalf("Test: %v success: %v? Expected %v got %v\n", input, answer == result, answer, result)
	}
	t.Logf("Test: %v success: %v? Expected %v got %v\n", input, answer == result, answer, result)
}

func validateOrderedTestResults[T cmp.Ordered](t *testing.T, input any, answer, result T) {
	if cmp.Compare(answer, result) != 0 {
		t.Fatalf("Test %v success: %v? Expected %v got %v\n", input, cmp.Compare(answer, result) == 0, answer, result)
	}
	t.Logf("Test: %v success: %v? Expected %v got %v\n", input, cmp.Compare(answer, result), answer, result)
}

func validateSliceTestResults[T ~[]E, E cmp.Ordered](t *testing.T, input any, answer, result T) {
	if slices.Compare(answer, result) != 0 {
		t.Fatalf("Test %v success: %v? Expected %v got %v\n", input, slices.Compare(answer, result) == 0, answer, result)
	}
	t.Logf("Test: %v success: %v? Expected %v got %v\n", input, slices.Compare(answer, result), answer, result)
}

func TestExample1(t *testing.T) {
	input := TestInput{
		N: 4,
		Edges: [][]int{
			{0, 1, 3},
			{1, 2, 1},
			{1, 3, 4},
			{2, 3, 1},
		},
		DistanceThreshold: 4,
	}
	answer := 3
	// Explain
	// the neighboring cities at a DistanceThreshold = 4 for each city are:
	// City 0 -> [City 1, City 2]
	// City 1 -> [City 0, City 2, City 3]
	// City 2 -> [City 0, City 1, City 3]
	// City 3 -> [City 1, City 2]
	// Cities 0 and 3 have 2 neighboring cities at a DistanceThreshold = 4, but we have to return city 3 since it has the greatest number.
	result := findTheCity(input.N, input.Edges, input.DistanceThreshold)
	validateTestResults(t, input, answer, result)
}

func TestExample2(t *testing.T) {
	input := TestInput{
		N: 5,
		Edges: [][]int{
			{0, 1, 2},
			{0, 4, 8},
			{1, 2, 3},
			{1, 4, 2},
			{2, 3, 1},
			{3, 4, 1},
		},
		DistanceThreshold: 2,
	}
	answer := 0
	// Explain
	// the neighboring cities at a DistanceThreshold = 2 for each city are:
	// City 0 -> [City 1]
	// City 1 -> [City 0, City 4]
	// City 2 -> [City 3, City 4]
	// City 3 -> [City 2, City 4]
	// City 4 -> [City 1, City 2, City 3]
	// Cities 0 and 3 have 2 neighboring cities at a DistanceThreshold = 2, but we have to return city 3 since it has the greatest number.
	result := findTheCity(input.N, input.Edges, input.DistanceThreshold)
	validateTestResults(t, input, answer, result)
}

func TestExample3(t *testing.T) {
	input := TestInput{
		N: 6,
		Edges: [][]int{
			{0, 3, 5},
			{2, 3, 7},
			{0, 5, 2},
			{0, 2, 5},
			{1, 2, 6},
			{1, 4, 7},
			{3, 4, 4},
			{2, 5, 5},
			{1, 5, 8},
		},
		DistanceThreshold: 8279,
	}
	answer := 5
	result := findTheCity(input.N, input.Edges, input.DistanceThreshold)
	validateTestResults(t, input, answer, result)
}

func TestExample28(t *testing.T) {
	input := TestInput{
		N: 5,
		Edges: [][]int{
			{0, 1, 2},
			{0, 4, 8},
			{1, 2, 10000},
			{1, 4, 2},
			{2, 3, 10000},
			{3, 4, 1},
		},
		DistanceThreshold: 10000,
	}
	answer := 2
	result := findTheCity(input.N, input.Edges, input.DistanceThreshold)
	validateTestResults(t, input, answer, result)
}

func TestExample48(t *testing.T) {
	input := TestInput{
		N: 9,
		Edges: [][]int{
			{0, 6, 7036},
			{2, 3, 6792},
			{6, 8, 5813},
			{0, 2, 5816},
			{5, 7, 8383},
			{1, 8, 2429},
			{0, 4, 2932},
			{3, 7, 3703},
			{5, 8, 3760},
			{4, 5, 973},
			{3, 6, 5989},
			{0, 1, 1098},
			{3, 4, 7315},
			{5, 6, 5274},
			{2, 8, 3566},
			{3, 8, 3590},
			{4, 7, 4223},
			{1, 3, 2540},
			{4, 8, 8536},
			{1, 2, 6250},
			{1, 7, 1757},
			{2, 6, 7826},
			{3, 5, 31},
			{1, 5, 7736},
			{2, 4, 5115},
			{0, 3, 4301},
			{0, 7, 4771},
			{4, 6, 3417},
			{0, 8, 2176},
		},
		DistanceThreshold: 6235,
	}
	answer := 6
	result := findTheCity(input.N, input.Edges, input.DistanceThreshold)
	validateTestResults(t, input, answer, result)
}

func TestExample51(t *testing.T) {
	input := TestInput{
		N: 8,
		Edges: [][]int{
			{3, 5, 9558},
			{1, 2, 1079},
			{1, 3, 8040},
			{0, 1, 9258},
			{4, 7, 7558},
			{5, 6, 8196},
			{3, 4, 7284},
			{1, 5, 6327},
			{0, 4, 5966},
			{3, 6, 8575},
			{2, 5, 8604},
			{1, 7, 7782},
			{4, 6, 2857},
			{3, 7, 2336},
			{0, 6, 6},
			{5, 7, 2870},
			{4, 5, 5055},
			{0, 7, 2904},
			{1, 6, 2458},
			{0, 5, 3399},
			{6, 7, 2202},
			{0, 2, 3996},
			{0, 3, 7495},
			{1, 4, 2262},
			{2, 6, 1390},
		},
		DistanceThreshold: 7937,
	}
	answer := 7
	result := findTheCity(input.N, input.Edges, input.DistanceThreshold)
	validateTestResults(t, input, answer, result)
}
