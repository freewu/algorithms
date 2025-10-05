package main

// 3702. Longest Subsequence With Non-Zero Bitwise XOR
// You are given an integer array nums.

// Return the length of the longest subsequence in nums whose bitwise XOR is non-zero. 
// If no such subsequence exists, return 0.

// A subsequence is a non-empty array that can be derived from another array by deleting some or 
// no elements without changing the order of the remaining elements.

// Example 1:
// Input: nums = [1,2,3]
// Output: 2
// Explanation:
// One longest subsequence is [2, 3]. The bitwise XOR is computed as 2 XOR 3 = 1, which is non-zero.

// Example 2:
// Input: nums = [2,3,4]
// Output: 3
// Explanation:
// The longest subsequence is [2, 3, 4]. The bitwise XOR is computed as 2 XOR 3 XOR 4 = 5, which is non-zero.

// Constraints:
//     1 <= nums.length <= 10^5
//     0 <= nums[i] <= 10^9

import "fmt"

func longestSubsequence(nums []int) int {
    xor, zero := 0, true
    for _, v := range nums {
        xor ^= v  // 按位异或（XOR）计算结果
        if v != 0 {
            zero = false
        }
    }
    if zero { return 0 } // 全部是 0 时，返回 0
    if xor != 0 { return len(nums) } // 异或值不为 0，返回数组长度
    return len(nums) - 1 // 异或值为 0，返回数组长度减 1 (去掉一个就不为 0 了)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 2
    // Explanation:
    // One longest subsequence is [2, 3]. The bitwise XOR is computed as 2 XOR 3 = 1, which is non-zero.
    fmt.Println(longestSubsequence([]int{1,2,3})) // 2
    // Example 2:
    // Input: nums = [2,3,4]
    // Output: 3
    // Explanation:
    // The longest subsequence is [2, 3, 4]. The bitwise XOR is computed as 2 XOR 3 XOR 4 = 5, which is non-zero.
    fmt.Println(longestSubsequence([]int{2,3,4})) // 3

    fmt.Println(longestSubsequence([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(longestSubsequence([]int{9,8,7,6,5,4,3,2,1})) // 9
}