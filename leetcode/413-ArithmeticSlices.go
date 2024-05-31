package main

// 413. Arithmetic Slices
// An integer array is called arithmetic if it consists of at least three elements 
// and if the difference between any two consecutive elements is the same.
//     For example, [1,3,5,7,9], [7,7,7,7], and [3,-1,-5,-9] are arithmetic sequences.

// Given an integer array nums, return the number of arithmetic subarrays of nums.
// A subarray is a contiguous subsequence of the array.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: 3
// Explanation: We have 3 arithmetic slices in nums: [1, 2, 3], [2, 3, 4] and [1,2,3,4] itself.

// Example 2:
// Input: nums = [1]
// Output: 0
 
// Constraints:
//     1 <= nums.length <= 5000
//     -1000 <= nums[i] <= 1000

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
    n := len(nums)
    if n < 3 { return 0 }

    res, dp := 0, 0
    for i := 2; i < n; i++ {
        // 每判断一组 3 个连续的数列，只需要用一个变量累加前面已经有多少个满足题意的连续元素，
        // 只要满足题意的等差数列就加上这个累加值
        if nums[i] - nums[i-1] == nums[i-1] - nums[i-2] {
            dp++
            res += dp
        } else { // 一旦不满足等差的条件，累加值置 0
            dp = 0
        }
    }
    return res
}

func numberOfArithmeticSlices1(nums []int) int {
    n := len(nums)
    if n < 3 { return 0 }

    res, first, second := 0, 0, 1
    diff := nums[second] - nums[first]
    for third := 2; third < n; third++ {
        if nums[third] - nums[third - 1] == diff {
            res += third - second
        } else {
            first, second, diff = third - 1, third, nums[third] - nums[third - 1]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4]
    // Output: 3
    // Explanation: We have 3 arithmetic slices in nums: [1, 2, 3], [2, 3, 4] and [1,2,3,4] itself.
    fmt.Println(numberOfArithmeticSlices([]int{1,2,3,4})) // 3
    // Example 2:
    // Input: nums = [1]
    // Output: 0
    fmt.Println(numberOfArithmeticSlices([]int{1})) // 0

    fmt.Println(numberOfArithmeticSlices1([]int{1,2,3,4})) // 3
    fmt.Println(numberOfArithmeticSlices1([]int{1})) // 0
}