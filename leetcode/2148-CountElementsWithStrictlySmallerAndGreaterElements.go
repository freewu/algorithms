package main

// 2148. Count Elements With Strictly Smaller and Greater Elements
// Given an integer array nums, 
// return the number of elements that have both a strictly smaller and a strictly greater element appear in nums.

// Example 1:
// Input: nums = [11,7,2,15]
// Output: 2
// Explanation: The element 7 has the element 2 strictly smaller than it and the element 11 strictly greater than it.
// Element 11 has element 7 strictly smaller than it and element 15 strictly greater than it.
// In total there are 2 elements having both a strictly smaller and a strictly greater element appear in nums.

// Example 2:
// Input: nums = [-3,3,3,90]
// Output: 2
// Explanation: The element 3 has the element -3 strictly smaller than it and the element 90 strictly greater than it.
// Since there are two elements with the value 3, in total there are 2 elements having both a strictly smaller and a strictly greater element appear in nums.

// Constraints:
//     1 <= nums.length <= 100
//     -10^5 <= nums[i] <= 10^5

import "fmt"
import "sort"

func countElements(nums []int) int {
    sort.Ints(nums)
    mn, mx, mnc, mxc := nums[0], nums[len(nums) - 1], 0, 0
    if mn == mx { // 处理所有元素都一样的情况 [1,1,1,1,1,1]
        return 0 
    }
    for _, v := range nums {
        if v == mn { mnc++ }
        if v == mx { mxc++ }
    }
    return len(nums) - mnc - mxc
}

func main() {
    // Example 1:
    // Input: nums = [11,7,2,15]
    // Output: 2
    // Explanation: The element 7 has the element 2 strictly smaller than it and the element 11 strictly greater than it.
    // Element 11 has element 7 strictly smaller than it and element 15 strictly greater than it.
    // In total there are 2 elements having both a strictly smaller and a strictly greater element appear in nums.
    fmt.Println(countElements([]int{11,7,2,15})) // 2
    // Example 2:
    // Input: nums = [-3,3,3,90]
    // Output: 2
    // Explanation: The element 3 has the element -3 strictly smaller than it and the element 90 strictly greater than it.
    // Since there are two elements with the value 3, in total there are 2 elements having both a strictly smaller and a strictly greater element appear in nums.
    fmt.Println(countElements([]int{-3,3,3,90})) // 2

    fmt.Println(countElements([]int{1,1,1,1,1,1})) // 0
}