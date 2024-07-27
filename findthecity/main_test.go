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
	// Explain
	// the neighboring cities at a DistanceThreshold = 2 for each city are:
	// City 0 -> [City 1]
	// City 1 -> [City 0, City 4]
	// City 2 -> [City 3, City 4]
	// City 3 -> [City 2, City 4]
	// City 4 -> [City 1, City 2, City 3]
	// Cities 0 and 3 have 2 neighboring cities at a DistanceThreshold = 2, but we have to return city 3 since it has the greatest number.

	// City 0 -> (7)[City 1:9258 City 4:5966 City 6:6 City 7:2904 City 5:3399 City 2:3996 City 3:7495]
	// City 1 -> (7)[City 2:1079 City 3:8040 City 0:9258 City 5:6327 City 7:7782 City 6:2458 City 4:2262]
	// City 2 -> (7)[City 1:1079 City 5:8604 City 0:3996 City 6:1390 City 3:9119 City 7:8861 City 4:3341]
	// City 3 -> (7)[City 5:9558 City 1:8040 City 4:7284 City 6:8575 City 7:2336 City 0:7495 City 2:9119]
	// City 4 -> (7)[City 7:7558 City 3:7284 City 0:5966 City 6:2857 City 5:5055 City 1:2262 City 2:3341]
	// City 5 -> (7)[City 3:9558 City 6:8196 City 1:6327 City 2:8604 City 7:2870 City 4:5055 City 0:3399]
	// City 6 -> (7)[City 5:8196 City 3:8575 City 4:2857 City 0:6 City 1:2458 City 7:2202 City 2:1390]
	// City 7 -> (7)[City 4:7558 City 1:7782 City 3:2336 City 5:2870 City 0:2904 City 6:2202 City 2:8861]
	result := findTheCity(input.N, input.Edges, input.DistanceThreshold)
	validateTestResults(t, input, answer, result)
}
