package main

// 325. Maximum Size Subarray Sum Equals k
// Given an integer array nums and an integer k, return the maximum length of a subarray that sums to k. 
// If there is not one, return 0 instead.

// Example 1:
// Input: nums = [1,-1,5,-2,3], k = 3
// Output: 4
// Explanation: The subarray [1, -1, 5, -2] sums to 3 and is the longest.

// Example 2:
// Input: nums = [-2,-1,2,1], k = 1
// Output: 2
// Explanation: The subarray [-1, 2] sums to 1 and is the longest.
 
// Constraints:
//     1 <= nums.length <= 2 * 10^5
//     -10^4 <= nums[i] <= 10^4
//     -10^9 <= k <= 10^9

import "fmt"

// 前缀和
func maxSubArrayLen(nums []int, k int) int {
    res, n, sum, m := 0, len(nums), 0, make(map[int]int)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        num := nums[i]
        sum += num
        if sum == k {
            res = i + 1
        } else if index, ok := m[sum-k]; ok {
            res = max(res, i - index)
        }
        if _, ok := m[sum]; !ok {
            m[sum] = i
        }
    }
    return res
}

func maxSubArrayLen1(nums []int, k int) int {
    res, m, preSum := 0, make(map[int]int, len(nums)), 0
    m[0] = -1
    for i := 0; i < len(nums); i++ {
        preSum += nums[i]
        if _, ok := m[preSum]; !ok {
            m[preSum] = i
        }
        need := preSum - k
        if preIndex , ok := m[need]; ok {
            res = max(res, i - preIndex)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,-1,5,-2,3], k = 3
    // Output: 4
    // Explanation: The subarray [1, -1, 5, -2] sums to 3 and is the longest.
    fmt.Println(maxSubArrayLen([]int{1,-1,5,-2,3}, 3)) // 4
    // Example 2:
    // Input: nums = [-2,-1,2,1], k = 1
    // Output: 2
    // Explanation: The subarray [-1, 2] sums to 1 and is the longest.
    fmt.Println(maxSubArrayLen([]int{-2,-1,2,1}, 1)) // 2

    fmt.Println(maxSubArrayLen1([]int{1,-1,5,-2,3}, 3)) // 4
    fmt.Println(maxSubArrayLen1([]int{-2,-1,2,1}, 1)) // 2
}