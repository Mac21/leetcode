package main

import (
	"cmp"
	"slices"
	"testing"
)

// Fill with example inputs
type Input struct {
	books      [][]int
	shelfWidth int
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

func TestLeetcodeExample1(t *testing.T) {
	input := Input{
		books:      [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}},
		shelfWidth: 4,
	}
	answer := 6
	result := minHeightShelves(input.books, input.shelfWidth)
	validateTestResults(t, input, answer, result)
}

func TestLeetcodeExample2(t *testing.T) {
	input := Input{
		books: [][]int{
			{1, 3}, {2, 4}, {3, 2},
		},
		shelfWidth: 6,
	}
	answer := 4
	result := minHeightShelves(input.books, input.shelfWidth)
	validateTestResults(t, input, answer, result)
}

func TestLeetcodeExample3(t *testing.T) {
	input := Input{
		books: [][]int{
			{7, 3}, {8, 7}, {2, 7}, {2, 5},
		},
		shelfWidth: 10,
	}
	answer := 15
	result := minHeightShelves(input.books, input.shelfWidth)
	validateTestResults(t, input, answer, result)
}

func TestLeetcodeExample7(t *testing.T) {
	input := Input{
		books: [][]int{
			{11, 83}, {170, 4}, {93, 80}, {155, 163}, {134, 118}, {75, 14}, {122, 192}, {123, 154}, {187, 29}, {160, 64}, {170, 152}, {113, 179}, {60, 102}, {28, 187}, {59, 95}, {187, 97}, {49, 193}, {67, 126}, {75, 45}, {130, 160}, {4, 102}, {116, 171}, {43, 170}, {96, 188}, {54, 15}, {167, 183}, {58, 158}, {59, 55}, {148, 183}, {89, 95}, {90, 113}, {51, 49}, {91, 28}, {172, 103}, {173, 3}, {131, 78}, {11, 199}, {77, 200}, {58, 65}, {77, 30}, {157, 58}, {18, 194}, {101, 148}, {22, 197}, {76, 181}, {21, 176}, {50, 45}, {80, 174}, {116, 198}, {138, 9}, {58, 125}, {163, 102}, {133, 175}, {21, 39}, {141, 156}, {34, 185}, {14, 113}, {11, 34}, {35, 184}, {16, 132}, {78, 147}, {85, 170}, {32, 149}, {46, 94}, {196, 3}, {155, 90}, {9, 114}, {117, 119}, {17, 157}, {94, 178}, {53, 55}, {103, 142}, {70, 121}, {9, 141}, {16, 170}, {92, 137}, {157, 30}, {94, 82}, {144, 149}, {128, 160}, {8, 147}, {153, 198}, {12, 22}, {140, 68}, {64, 172}, {86, 63}, {66, 158}, {23, 15}, {120, 99}, {27, 165}, {79, 174}, {46, 19}, {60, 98}, {160, 172}, {128, 184}, {63, 172}, {135, 54}, {40, 4}, {102, 171}, {29, 125}, {81, 9}, {111, 197}, {16, 90}, {22, 150}, {168, 126}, {187, 61}, {47, 190}, {54, 110}, {106, 102}, {55, 47}, {117, 134}, {33, 107}, {2, 10}, {18, 62}, {109, 188}, {113, 37}, {59, 159}, {120, 175}, {17, 147}, {112, 195}, {177, 53}, {148, 173}, {29, 105}, {196, 32}, {123, 51}, {29, 19}, {161, 178}, {148, 2}, {70, 124}, {126, 9}, {105, 87}, {41, 121}, {147, 10}, {78, 167}, {91, 197}, {22, 98}, {73, 33}, {148, 194}, {166, 64}, {33, 138}, {139, 158}, {160, 19}, {140, 27}, {103, 109}, {88, 16}, {99, 181}, {2, 140}, {50, 188}, {200, 77}, {73, 84}, {159, 130}, {115, 199}, {152, 79}, {1, 172}, {124, 136}, {117, 138}, {158, 86}, {193, 150}, {56, 57}, {150, 133}, {52, 186}, {21, 145}, {127, 97}, {108, 110}, {174, 44}, {199, 169}, {139, 200}, {66, 48}, {52, 190}, {27, 86}, {142, 191}, {191, 79}, {126, 114}, {125, 100}, {176, 95}, {104, 79}, {146, 189}, {144, 78}, {52, 106}, {74, 74}, {163, 128}, {34, 181}, {20, 178}, {15, 107}, {105, 8}, {66, 142}, {39, 126}, {95, 59}, {164, 69}, {138, 18}, {110, 145}, {128, 200}, {149, 150}, {149, 93}, {145, 140}, {90, 170}, {81, 127}, {57, 151}, {167, 127}, {95, 89},
		},
		shelfWidth: 200,
	}
	answer := 15672
	result := minHeightShelves(input.books, input.shelfWidth)
	validateTestResults(t, input, answer, result)
}
