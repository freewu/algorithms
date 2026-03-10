package main

// 3865. Reverse K Subarrays
// You are given an integer array nums of length n and an integer k.

// You must partition the array into k contiguous subarrays of equal length and reverse each subarray.

// It is guaranteed that n is divisible by k.

// Return the resulting array after performing the above operation.

// Example 1:
// Input: nums = [1,2,4,3,5,6], k = 3
// Output: [2,1,3,4,6,5]
// Explanation:
// The array is partitioned into k = 3 subarrays: [1, 2], [4, 3], and [5, 6].
// After reversing each subarray: [2, 1], [3, 4], and [6, 5].
// Combining them gives the final array [2, 1, 3, 4, 6, 5].

// Example 2:
// Input: nums = [5,4,4,2], k = 1
// Output: [2,4,4,5]
// Explanation:
// The array is partitioned into k = 1 subarray: [5, 4, 4, 2].
// Reversing it produces [2, 4, 4, 5], which is the final array.

// Constraints:
//     1 <= n == nums.length <= 1000
//     1 <= nums[i] <= 1000
//     1 <= k <= n
//     n is divisible by k.

import "fmt"

func reverseSubarrays(nums []int, k int) []int {
    n := len(nums) / k // 计算每个子数组的长度
    reverse := func(s []int) {
        for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
            s[l], s[r] = s[r], s[l]
        }
    }
    for i := 0; i < k; i++ {  // 遍历每个子数组的起始索引 循环次数为k次，而非i < len(nums)
        // 确定当前子数组的起始和结束索引
        start := i * n
        end := start + n
        reverse(nums[start:end])  // 反转当前子数组 [start, end)
    }
    return nums
}

func main() {
    // Example 1:
    // Input: nums = [1,2,4,3,5,6], k = 3
    // Output: [2,1,3,4,6,5]
    // Explanation:
    // The array is partitioned into k = 3 subarrays: [1, 2], [4, 3], and [5, 6].
    // After reversing each subarray: [2, 1], [3, 4], and [6, 5].
    // Combining them gives the final array [2, 1, 3, 4, 6, 5].
    fmt.Println(reverseSubarrays([]int{1,2,4,3,5,6}, 3)) // [2,1,3,4,6,5]
    // Example 2:
    // Input: nums = [5,4,4,2], k = 1
    // Output: [2,4,4,5]
    // Explanation:
    // The array is partitioned into k = 1 subarray: [5, 4, 4, 2].
    // Reversing it produces [2, 4, 4, 5], which is the final array.
    fmt.Println(reverseSubarrays([]int{5,4,4,2}, 1)) // [2,4,4,5]

    fmt.Println(reverseSubarrays([]int{1,2,3,4,5,6,7,8,9}, 2)) // [4 3 2 1 8 7 6 5 9]
    fmt.Println(reverseSubarrays([]int{9,8,7,6,5,4,3,2,1}, 2)) // [6 7 8 9 2 3 4 5 1]
}