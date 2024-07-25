package main

import (
	"cmp"
	"fmt"
	"slices"
)

func displayTestResults[T comparable](input any, answer, result T) {
    fmt.Printf("Test: %v success: %v? Expected %v got %v\n\n", input, answer == result, answer, result)
}

func displayOrderedTestResults[T cmp.Ordered](input any, answer, result T) {
	fmt.Printf("Test %v success: %v? Expected %v got %v\n\n", input, cmp.Compare(answer, result) == 0, answer, result)
}

func displaySliceTestResults[T ~[]E, E cmp.Ordered](input any, answer, result T) {
	fmt.Printf("Test %v success: %v? Expected %v got %v\n\n", input, slices.Compare(answer, result) == 0, answer, result)
}

func isOneBitCharacter(bits []int) bool {
    everythingPaired := false
    for i := 0; i < len(bits); {
        bit := bits[i]
        switch bit {
        case 0:
            everythingPaired = true
            i++
        case 1:
            everythingPaired = false
            i += 2
        }
    }

    return everythingPaired
}

/*
Problem description:
    We have two special characters:

    The first character can be represented by one bit 0
    The second character can be represented by two bits 10 or 11

    Given a binary array bits that ends with 0 return true if the last character must be a one-bit character

Constraints: 
    1 <= bits.length <= 1000
    bits[i] is either 0 or 1
*/
func main() {
    answer := true
    input := []int{0}
    result := isOneBitCharacter(input)
    displayTestResults(input, answer, result)

    answer = false
    input = []int{1, 0}
    result = isOneBitCharacter(input)
    displayTestResults(input, answer, result)

    answer = true
    input = []int{0, 0}
    result = isOneBitCharacter(input)
    displayTestResults(input, answer, result)


	answer = true
    input = []int{0, 0, 0}
	result = isOneBitCharacter(input)
	displayTestResults(input, answer, result)

	answer = true
    input = []int{1, 0, 0}
	result = isOneBitCharacter(input)
	displayTestResults(input, answer, result)

    answer = false
    input = []int{0, 1, 0}
    result = isOneBitCharacter(input)
    displayTestResults(input, answer, result)

    answer = true
    input = []int{1, 1, 0}
    result = isOneBitCharacter(input)
    displayTestResults(input, answer, result)

    answer = true
    input = []int{1, 1, 1, 1, 0}
    result = isOneBitCharacter(input)
    displayTestResults(input, answer, result)

    answer = false
    input = []int{1, 1, 1, 0}
    result = isOneBitCharacter(input)
    displayTestResults(input, answer, result)

    answer = true
    input = []int{0, 1, 1, 0}
    result = isOneBitCharacter(input)
    displayTestResults(input, answer, result)

    answer = false
    input = []int{0, 0, 1, 0}
    result = isOneBitCharacter(input)
    displayTestResults(input, answer, result)
}
