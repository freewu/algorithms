package main

// 2401. Longest Nice Subarray
// You are given an array nums consisting of positive integers.

// We call a subarray of nums nice if the bitwise AND of every pair of elements that are in different positions in the subarray is equal to 0.

// Return the length of the longest nice subarray.

// A subarray is a contiguous part of an array.

// Note that subarrays of length 1 are always considered nice.

// Example 1:
// Input: nums = [1,3,8,48,10]
// Output: 3
// Explanation: The longest nice subarray is [3,8,48]. This subarray satisfies the conditions:
// - 3 AND 8 = 0.
// - 3 AND 48 = 0.
// - 8 AND 48 = 0.
// It can be proven that no longer nice subarray can be obtained, so we return 3.

// Example 2:
// Input: nums = [3,1,5,11,13]
// Output: 1
// Explanation: The length of the longest nice subarray is 1. Any subarray of length 1 can be chosen.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func longestNiceSubarray(nums []int) int {
    res, mask, l, r, n := 1, 0, 0, 0, len(nums)
    for r < n {
        for (l < r) && (mask & nums[r] != 0) {
            mask ^= nums[l]
            l++
        }
        mask |= nums[r]
        r++
        if res < r - l {
            res = r - l
        }
    }
    return res
}

func longestNiceSubarray1(nums []int) int {
    res, start := 1, -1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(nums); i++ {
        if start == -1 {
            if nums[i] & nums[i - 1] == 0 {
                start, res = i - 1, max(res, 2)
            }
            continue
        }
        j := start
        for ; j < i; j++ {
            if nums[j] & nums[i] != 0 { break }
        }
        if j == i {
            res = max(res, i - start + 1)
        } else { // 当前的 i 和 j 位置冲突，那就回溯到 j + 1 的位置开始继续检查
            i, start = j + 1, -1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,8,48,10]
    // Output: 3
    // Explanation: The longest nice subarray is [3,8,48]. This subarray satisfies the conditions:
    // - 3 AND 8 = 0.
    // - 3 AND 48 = 0.
    // - 8 AND 48 = 0.
    // It can be proven that no longer nice subarray can be obtained, so we return 3.
    fmt.Println(longestNiceSubarray([]int{1,3,8,48,10})) // 3
    // Example 2:
    // Input: nums = [3,1,5,11,13]
    // Output: 1
    // Explanation: The length of the longest nice subarray is 1. Any subarray of length 1 can be chosen.
    fmt.Println(longestNiceSubarray([]int{3,1,5,11,13})) // 1

    fmt.Println(longestNiceSubarray1([]int{1,3,8,48,10})) // 3
    fmt.Println(longestNiceSubarray1([]int{3,1,5,11,13})) // 1
}