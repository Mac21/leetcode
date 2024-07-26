package main

/*
Problem description:
    Given an array of integers nums, sort the array in ascending order and return it.
    You must solve the problem with using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.
Constraints: 
    1 <= nums.length <= 5 * 10^4
    -5 * 10^4 <= nums[i] <= 5 * 10^4
*/

// [3 -1 2]
// [3 -1] [-1 2] 
func merge(temp, nums []int, start, middle, end int) {
    i := start
    j := middle

    for k := start; k < end; k++ {
        if i < middle && (j >= end || nums[i] <= nums[j]) {
            temp[k] = nums[i]
            i++
        } else {
            temp[k] = nums[j]
            j++
        }
    }
}

func mergeSort(temp, nums []int, start, end int) {
    if (end - start) <= 1 {
        return
    }

    middle := (start + end) / 2
    mergeSort(nums, temp, start, middle)
    mergeSort(nums, temp, middle, end)
    merge(nums, temp, start, middle, end)
}

func sortArray(nums []int) []int {
    n := len(nums)
    tempNums := make([]int, n)
    copy(tempNums, nums)

    mergeSort(tempNums, nums, 0, n)
    return nums
}
