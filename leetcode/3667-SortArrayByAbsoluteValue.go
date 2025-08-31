package main

// 3667. Sort Array By Absolute Value
// You are given an integer array nums.

// Rearrange elements of nums in non-decreasing order of their absolute value.

// Return any rearranged array that satisfies this condition.

// Note: The absolute value of an integer x is defined as:
//     x if x >= 0
//     -x if x < 0

// Example 1:
// Input: nums = [3,-1,-4,1,5]
// Output: [-1,1,3,-4,5]
// Explanation:
// The absolute values of elements in nums are 3, 1, 4, 1, 5 respectively.
// Rearranging them in increasing order, we get 1, 1, 3, 4, 5.
// This corresponds to [-1, 1, 3, -4, 5]. Another possible rearrangement is [1, -1, 3, -4, 5].

// Example 2:
// Input: nums = [-100,100]
// Output: [-100,100]
// Explanation:
// The absolute values of elements in nums are 100, 100 respectively.
// Rearranging them in increasing order, we get 100, 100.
// This corresponds to [-100, 100]. Another possible rearrangement is [100, -100].

// Constraints:
//     1 <= nums.length <= 100
//     -100 <= nums[i] <= 100

import "fmt"
import "slices"

func sortByAbsoluteValue(nums []int) []int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    slices.SortFunc(nums, func(i, j int) int {
        a, b := abs(i), abs(j)
        if a == b {
            return i - j
        }
        return a - b
    })
    return nums;
}

func main() {
    // Example 1:
    // Input: nums = [3,-1,-4,1,5]
    // Output: [-1,1,3,-4,5]
    // Explanation:
    // The absolute values of elements in nums are 3, 1, 4, 1, 5 respectively.
    // Rearranging them in increasing order, we get 1, 1, 3, 4, 5.
    // This corresponds to [-1, 1, 3, -4, 5]. Another possible rearrangement is [1, -1, 3, -4, 5].
    fmt.Println(sortByAbsoluteValue([]int{3,-1,-4,1,5}))  // [-1,1,3,-4,5]
    // Example 2:
    // Input: nums = [-100,100]
    // Output: [-100,100]
    // Explanation:
    // The absolute values of elements in nums are 100, 100 respectively.
    // Rearranging them in increasing order, we get 100, 100.
    // This corresponds to [-100, 100]. Another possible rearrangement is [100, -100].
    fmt.Println(sortByAbsoluteValue([]int{-100,100}))  // [-100,100]

    fmt.Println(sortByAbsoluteValue([]int{1,2,3,4,5,6,7,8,9}))  // [1,2,3,4,5,6,7,8,9]
    fmt.Println(sortByAbsoluteValue([]int{9,8,7,6,5,4,3,2,1}))  // [1,2,3,4,5,6,7,8,9]
}