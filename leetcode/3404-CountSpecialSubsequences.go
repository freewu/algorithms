package main

// 3404. Count Special Subsequences
// You are given an array nums consisting of positive integers.

// A special subsequence is defined as a subsequence of length 4, represented by indices (p, q, r, s), where p < q < r < s. 
// This subsequence must satisfy the following conditions:
//     1. nums[p] * nums[r] == nums[q] * nums[s]
//     2. There must be at least one element between each pair of indices. 
//        In other words, q - p > 1, r - q > 1 and s - r > 1.

// Return the number of different special subsequences in nums.

// Example 1:
// Input: nums = [1,2,3,4,3,6,1]
// Output: 1
// Explanation:
// There is one special subsequence in nums.
// (p, q, r, s) = (0, 2, 4, 6):
// This corresponds to elements (1, 3, 3, 1).
// nums[p] * nums[r] = nums[0] * nums[4] = 1 * 3 = 3
// nums[q] * nums[s] = nums[2] * nums[6] = 3 * 1 = 3

// Example 2:
// Input: nums = [3,4,3,4,3,4,3,4]
// Output: 3
// Explanation:
// There are three special subsequences in nums.
// (p, q, r, s) = (0, 2, 4, 6):
// This corresponds to elements (3, 3, 3, 3).
// nums[p] * nums[r] = nums[0] * nums[4] = 3 * 3 = 9
// nums[q] * nums[s] = nums[2] * nums[6] = 3 * 3 = 9
// (p, q, r, s) = (1, 3, 5, 7):
// This corresponds to elements (4, 4, 4, 4).
// nums[p] * nums[r] = nums[1] * nums[5] = 4 * 4 = 16
// nums[q] * nums[s] = nums[3] * nums[7] = 4 * 4 = 16
// (p, q, r, s) = (0, 2, 5, 7):
// This corresponds to elements (3, 3, 4, 4).
// nums[p] * nums[r] = nums[0] * nums[5] = 3 * 4 = 12
// nums[q] * nums[s] = nums[2] * nums[7] = 3 * 4 = 12

// Constraints:
//     7 <= nums.length <= 1000
//     1 <= nums[i] <= 1000

import "fmt"

func numberOfSubsequences(nums []int) int64 {
    res, n := 0, len(nums)
    matches := make(map[float32]int)
    for i := 2; i < n - 4; i++ {
        for j := 0; j < i - 1; j++ {
            matches[float32(nums[j]) / float32(nums[i])]++
        }
        r := i + 2
        for j := r + 2; j < n; j++ {
            res += matches[float32(nums[j]) / float32(nums[r])]
        }
    }
    return int64(res) 
}

func numberOfSubsequences1(nums []int) int64 {
    res, mx := 0, nums[0]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for i := range nums {
        mx = max(mx, nums[i])
    }
    dp := make([][]int, mx + 1)
    for i := range dp {
        dp[i] = make([]int, mx + 1)
    }
    for q := len(nums) - 5; q >= 2; q-- {
        r := q + 2
        for s := q + 4; s < len(nums); s++ {
            g := gcd(nums[r], nums[s])
            dp[nums[r]/g][nums[s]/g]++
        }
        for p := 0; p <= q-2; p++ {
            g := gcd(nums[p], nums[q])
            res += dp[nums[q]/g][nums[p]/g]
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4,3,6,1]
    // Output: 1
    // Explanation:
    // There is one special subsequence in nums.
    // (p, q, r, s) = (0, 2, 4, 6):
    // This corresponds to elements (1, 3, 3, 1).
    // nums[p] * nums[r] = nums[0] * nums[4] = 1 * 3 = 3
    // nums[q] * nums[s] = nums[2] * nums[6] = 3 * 1 = 3
    fmt.Println(numberOfSubsequences([]int{1,2,3,4,3,6,1})) // 1
    // Example 2:
    // Input: nums = [3,4,3,4,3,4,3,4]
    // Output: 3
    // Explanation:
    // There are three special subsequences in nums.
    // (p, q, r, s) = (0, 2, 4, 6):
    // This corresponds to elements (3, 3, 3, 3).
    // nums[p] * nums[r] = nums[0] * nums[4] = 3 * 3 = 9
    // nums[q] * nums[s] = nums[2] * nums[6] = 3 * 3 = 9
    // (p, q, r, s) = (1, 3, 5, 7):
    // This corresponds to elements (4, 4, 4, 4).
    // nums[p] * nums[r] = nums[1] * nums[5] = 4 * 4 = 16
    // nums[q] * nums[s] = nums[3] * nums[7] = 4 * 4 = 16
    // (p, q, r, s) = (0, 2, 5, 7):
    // This corresponds to elements (3, 3, 4, 4).
    // nums[p] * nums[r] = nums[0] * nums[5] = 3 * 4 = 12
    // nums[q] * nums[s] = nums[2] * nums[7] = 3 * 4 = 12
    fmt.Println(numberOfSubsequences([]int{3,4,3,4,3,4,3,4})) // 3

    fmt.Println(numberOfSubsequences([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(numberOfSubsequences([]int{9,8,7,6,5,4,3,2,1})) // 0

    fmt.Println(numberOfSubsequences1([]int{1,2,3,4,3,6,1})) // 1
    fmt.Println(numberOfSubsequences1([]int{3,4,3,4,3,4,3,4})) // 3
    fmt.Println(numberOfSubsequences1([]int{1,2,3,4,5,6,7,8,9})) // 0
    fmt.Println(numberOfSubsequences1([]int{9,8,7,6,5,4,3,2,1})) // 0
}