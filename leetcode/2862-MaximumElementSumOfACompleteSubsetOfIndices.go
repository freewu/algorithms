package main

// 2862. Maximum Element-Sum of a Complete Subset of Indices
// You are given a 1-indexed array nums. 
// Your task is to select a complete subset from nums where every pair of selected indices multiplied is a perfect square,
// . i. e. if you select ai and aj, i * j must be a perfect square.

// Return the sum of the complete subset with the maximum sum.

// Example 1:
// Input: nums = [8,7,3,5,7,2,4,9]
// Output: 16
// Explanation:
// We select elements at indices 2 and 8 and 2 * 8 is a perfect square.

// Example 2:
// Input: nums = [8,10,3,8,1,13,7,9,4]
// Output: 20
// Explanation:
// We select elements at indices 1, 4, and 9. 1 * 4, 1 * 9, 4 * 9 are perfect squares.

// Constraints:
//     1 <= n == nums.length <= 10^4
//     1 <= nums[i] <= 10^9

import "fmt"

func maximumSum(nums []int) int64 {
    res, n := 0, len(nums)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        sum := 0
        for j := 1; j * j * i <= n; j++ { 
            sum += nums[j * j * i - 1]
        }
        res = max(res, sum)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [8,7,3,5,7,2,4,9]
    // Output: 16
    // Explanation:
    // We select elements at indices 2 and 8 and 2 * 8 is a perfect square.
    fmt.Println(maximumSum([]int{8,7,3,5,7,2,4,9})) // 16
    // Example 2:
    // Input: nums = [8,10,3,8,1,13,7,9,4]
    // Output: 20
    // Explanation:
    // We select elements at indices 1, 4, and 9. 1 * 4, 1 * 9, 4 * 9 are perfect squares.
    fmt.Println(maximumSum([]int{8,10,3,8,1,13,7,9,4})) // 20

    fmt.Println(maximumSum([]int{1,2,3,4,5,6,7,8,9})) // 14
    fmt.Println(maximumSum([]int{9,8,7,6,5,4,3,2,1})) // 16
}