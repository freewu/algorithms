package main

// 1685. Sum of Absolute Differences in a Sorted Array
// You are given an integer array nums sorted in non-decreasing order.

// Build and return an integer array result with the same length as nums 
// such that result[i] is equal to the summation of absolute differences between nums[i] 
// and all the other elements in the array.

// In other words, result[i] is equal to sum(|nums[i]-nums[j]|) where 0 <= j < nums.length and j != i (0-indexed).

// Example 1:
// Input: nums = [2,3,5]
// Output: [4,3,5]
// Explanation: Assuming the arrays are 0-indexed, then
// result[0] = |2-2| + |2-3| + |2-5| = 0 + 1 + 3 = 4,
// result[1] = |3-2| + |3-3| + |3-5| = 1 + 0 + 2 = 3,
// result[2] = |5-2| + |5-3| + |5-5| = 3 + 2 + 0 = 5.

// Example 2:
// Input: nums = [1,4,6,8,10]
// Output: [24,15,13,15,21]

// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i] <= nums[i + 1] <= 10^4

import "fmt"

func getSumAbsoluteDifferences(nums []int) []int {
    left, sum, n := 0, 0, len(nums)
    res := []int{}
    for _, v := range nums { 
        sum += v 
    }
    for i, v := range nums {
        right := sum - v - left
        res = append(res, (v * i - left) + (right - (n - i - 1) * v))
        left += v
    }
    return res
}

func getSumAbsoluteDifferences1(nums []int) []int {
    n, left, right := len(nums), 0, 0
    res := make([]int, n)
    for _, v := range nums { 
        right += v 
    }
    for i := 0; i < n; i++ {
        left += nums[i]
        if i == 0  {
            res[i] = right - nums[i] * n
        } else if i == n - 1 {
            res[i] = nums[i] * (i + 1) - left
        } else {
            res[i] = (right - nums[i] * (n - i)) + (nums[i] * (i + 1) - left)
        }
        right -= nums[i]
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,3,5]
    // Output: [4,3,5]
    // Explanation: Assuming the arrays are 0-indexed, then
    // result[0] = |2-2| + |2-3| + |2-5| = 0 + 1 + 3 = 4,
    // result[1] = |3-2| + |3-3| + |3-5| = 1 + 0 + 2 = 3,
    // result[2] = |5-2| + |5-3| + |5-5| = 3 + 2 + 0 = 5.
    fmt.Println(getSumAbsoluteDifferences([]int{2,3,5})) // [4,3,5]
    // Example 2:
    // Input: nums = [1,4,6,8,10]
    // Output: [24,15,13,15,21]
    fmt.Println(getSumAbsoluteDifferences([]int{1,4,6,8,10})) // [24,15,13,15,21]

    fmt.Println(getSumAbsoluteDifferences1([]int{2,3,5})) // [4,3,5]
    fmt.Println(getSumAbsoluteDifferences1([]int{1,4,6,8,10})) // [24,15,13,15,21]
}