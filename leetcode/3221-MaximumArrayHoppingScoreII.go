package main

// 3221. Maximum Array Hopping Score II
// Given an array nums, you have to get the maximum score starting from index 0 and hopping until you reach the last element of the array.
// In each hop, you can jump from index i to an index j > i, and you get a score of (j - i) * nums[j].
// Return the maximum score you can get.

// Example 1:
// Input: nums = [1,5,8]
// Output: 16
// Explanation:
// There are two possible ways to reach the last element:
// 0 -> 1 -> 2 with a score of (1 - 0) * 5 + (2 - 1) * 8 = 13.
// 0 -> 2 with a score of (2 - 0) * 8 = 16.

// Example 2:
// Input: nums = [4,5,2,8,9,1,3]
// Output: 42
// Explanation:
// We can do the hopping 0 -> 4 -> 6 with a score of (4 - 0) * 9 + (6 - 4) * 3 = 42.

// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"

// stack
func maxScore(nums []int) int64 {
    stack, n, res := []int{}, len(nums), 0
    for i := 1; i < n; i++ {
        for len(stack) != 0 && nums[stack[len(stack) - 1]] <= nums[i] {
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, i)
    }
    for len(stack) != 0 {
        cur := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        prev := 0
        if len(stack) != 0 {
            prev = stack[len(stack)-1]
        }
        res += (cur - prev) * nums[cur]
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,5,8]
    // Output: 16
    // Explanation:
    // There are two possible ways to reach the last element:
    // 0 -> 1 -> 2 with a score of (1 - 0) * 5 + (2 - 1) * 8 = 13.
    // 0 -> 2 with a score of (2 - 0) * 8 = 16.
    fmt.Println(maxScore([]int{1,5,8})) // 16
    // Example 2:
    // Input: nums = [4,5,2,8,9,1,3]
    // Output: 42
    // Explanation:
    // We can do the hopping 0 -> 4 -> 6 with a score of (4 - 0) * 9 + (6 - 4) * 3 = 42.
    fmt.Println(maxScore([]int{4,5,2,8,9,1,3})) // 42
}