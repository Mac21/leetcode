package main

import (
    "testing"
    "cmp"
    "slices"
)

type TestInput struct {
    N int
    Edges [][]int
    DistanceThreshold int
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
