package main

import (
    "cmp"
    "slices"
    "sort"
    "fmt"
)

func mapNum(num int, mapping []int) int {
    if num == 0 {
        return mapping[num]
    }

    res := 0
    multiplier := 1
    for num != 0 {
        res += mapping[num % 10] * multiplier
        multiplier *= 10
        num /= 10
    }
    return res
}

func sortJumbledSlices(mapping []int, nums []int) []int {
    slices.SortStableFunc(nums, func (a, b int) int {
        return cmp.Compare(mapNum(a, mapping), mapNum(b, mapping))
    })

    return nums
}

func sortJumbledSort(mapping []int, nums []int) []int {
    sort.SliceStable(nums, func (i, j int) bool {
        mappedI, mappedJ := mapNum(nums[i], mapping), mapNum(nums[j], mapping)

        return mappedI < mappedJ
    })

    return nums
}

func main() {
    mapping := []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}
    nums := []int{991, 338, 38}
    fmt.Printf("Input mapping: %v, nums: %v\n", mapping, nums)
    fmt.Println("Result sortJumbledSlices", sortJumbledSlices(mapping, nums))

    mapping = []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}
    nums = []int{991, 338, 38}
    fmt.Println("Result sortJumbledSort", sortJumbledSort(mapping, nums))

    mapping = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
    nums = []int{0,1,2,3,4,5,6,7,8,9}
    fmt.Printf("Input mapping: %v, nums: %v\n", mapping, nums)
    fmt.Println("Result sortJumbledSlices", sortJumbledSlices(mapping, nums))

    mapping = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
    nums = []int{0,1,2,3,4,5,6,7,8,9}
    fmt.Println("Result sortJumbledSort", sortJumbledSort(mapping, nums))
}
