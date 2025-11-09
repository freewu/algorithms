package main

// 3738. Longest Non-Decreasing Subarray After Replacing at Most One Element
// You are given an integer array nums.

// You are allowed to replace at most one element in the array with any other integer value of your choice.

// Return the length of the longest non-decreasing subarray that can be obtained after performing at most one replacement.

// A subarray is a contiguous sequence of elements within an array.

// An array is said to be non-decreasing if each element is greater than or equal to its previous one (if it exists).

// Example 1:
// Input: nums = [1,2,3,1,2]
// Output: 4
// Explanation:
// Replacing nums[3] = 1 with 3 gives the array [1, 2, 3, 3, 2].
// The longest non-decreasing subarray is [1, 2, 3, 3], which has a length of 4.

// Example 2:
// Input: nums = [2,2,2,2,2]
// Output: 5
// Explanation:
// All elements in nums are equal, so it is already non-decreasing and the entire nums forms a subarray of length 5.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9​​​​​​​

import "fmt"
import "slices"

func longestSubarray(nums []int) int {
    n := len(nums)
    left, right := make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        left[i], right[i]  = 1, 1 // fill 1
    }
    for i := 1; i < n; i++ {
        if (nums[i - 1] <= nums[i]) {
            left[i] = left[i - 1] + 1
        }
    }
    for i := n - 2; i >= 0; i-- {
        if nums[i] <= nums[i + 1] {
            right[i] = right[i + 1] + 1
        }
    }
    res := min(n, slices.Max(left) + 1)
    for i := 1; i < n - 1; i++ {
        if nums[i - 1] <= nums[i + 1] {
            res = max(res, left[i - 1] + 1 + right[i + 1])
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,1,2]
    // Output: 4
    // Explanation:
    // Replacing nums[3] = 1 with 3 gives the array [1, 2, 3, 3, 2].
    // The longest non-decreasing subarray is [1, 2, 3, 3], which has a length of 4.
    fmt.Println(longestSubarray([]int{1,2,3,1,2})) // 4
    // Example 2:
    // Input: nums = [2,2,2,2,2]
    // Output: 5
    // Explanation:
    // All elements in nums are equal, so it is already non-decreasing and the entire nums forms a subarray of length 5.
    fmt.Println(longestSubarray([]int{2,2,2,2,2})) // 5

    fmt.Println(longestSubarray([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(longestSubarray([]int{9,8,7,6,5,4,3,2,1})) // 2
}