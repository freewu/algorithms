package main

// 2219. Maximum Sum Score of Array
// You are given a 0-indexed integer array nums of length n.

// The sum score of nums at an index i where 0 <= i < n is the maximum of:
//     The sum of the first i + 1 elements of nums.
//     The sum of the last n - i elements of nums.

// Return the maximum sum score of nums at any index.

// Example 1:
// Input: nums = [4,3,-2,5]
// Output: 10
// Explanation:
// The sum score at index 0 is max(4, 4 + 3 + -2 + 5) = max(4, 10) = 10.
// The sum score at index 1 is max(4 + 3, 3 + -2 + 5) = max(7, 6) = 7.
// The sum score at index 2 is max(4 + 3 + -2, -2 + 5) = max(5, 3) = 5.
// The sum score at index 3 is max(4 + 3 + -2 + 5, 5) = max(10, 5) = 10.
// The maximum sum score of nums is 10.

// Example 2:
// Input: nums = [-3,-5]
// Output: -3
// Explanation:
// The sum score at index 0 is max(-3, -3 + -5) = max(-3, -8) = -3.
// The sum score at index 1 is max(-3 + -5, -5) = max(-8, -5) = -5.
// The maximum sum score of nums is -3.

// Constraints:
//     n == nums.length
//     1 <= n <= 10^5
//     -10^5 <= nums[i] <= 10^5

import "fmt"

func maximumSumScore(nums []int) int64 {
    n, sum := len(nums), 0
    for _, v := range nums {
        sum += v
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    left, right := nums[0], sum
    res := max(left, right)
    for i := 1; i < n; i++ {
        left += nums[i]
        right = sum - left + nums[i]
        res = max(res, max(left, right))
    }
    return int64(res)
}

func maximumSumScore1(nums []int) int64 {
    n := len(nums)
    left, right := make([]int, n), make([]int, n)
    for i := 1; i < n; i++ {
        left[i] = left[i - 1] + nums[i - 1]
    }
    for i := n-2; i >= 0; i-- {
        right[i] = right[i + 1] + nums[i + 1]
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res := max(left[0], right[0]) + nums[0]
    for i := 0; i < n; i++ {
        res = max(res, max(left[i], right[i]) + nums[i])
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [4,3,-2,5]
    // Output: 10
    // Explanation:
    // The sum score at index 0 is max(4, 4 + 3 + -2 + 5) = max(4, 10) = 10.
    // The sum score at index 1 is max(4 + 3, 3 + -2 + 5) = max(7, 6) = 7.
    // The sum score at index 2 is max(4 + 3 + -2, -2 + 5) = max(5, 3) = 5.
    // The sum score at index 3 is max(4 + 3 + -2 + 5, 5) = max(10, 5) = 10.
    // The maximum sum score of nums is 10.
    fmt.Println(maximumSumScore([]int{4,3,-2,5})) // 10
    // Example 2:
    // Input: nums = [-3,-5]
    // Output: -3
    // Explanation:
    // The sum score at index 0 is max(-3, -3 + -5) = max(-3, -8) = -3.
    // The sum score at index 1 is max(-3 + -5, -5) = max(-8, -5) = -5.
    // The maximum sum score of nums is -3.
    fmt.Println(maximumSumScore([]int{-3,-5})) // -3

    fmt.Println(maximumSumScore1([]int{4,3,-2,5})) // 10
    fmt.Println(maximumSumScore1([]int{-3,-5})) // -3
}