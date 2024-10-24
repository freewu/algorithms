package main

// 1918. Kth Smallest Subarray Sum
// Given an integer array nums of length n and an integer k, return the kth smallest subarray sum.

// A subarray is defined as a non-empty contiguous sequence of elements in an array. 
// A subarray sum is the sum of all elements in the subarray.

// Example 1:
// Input: nums = [2,1,3], k = 4
// Output: 3
// Explanation: The subarrays of [2,1,3] are:
// - [2] with sum 2
// - [1] with sum 1
// - [3] with sum 3
// - [2,1] with sum 3
// - [1,3] with sum 4
// - [2,1,3] with sum 6 
// Ordering the sums from smallest to largest gives 1, 2, 3, 3, 4, 6. The 4th smallest is 3.

// Example 2:
// Input: nums = [3,3,5,5], k = 7
// Output: 10
// Explanation: The subarrays of [3,3,5,5] are:
// - [3] with sum 3
// - [3] with sum 3
// - [5] with sum 5
// - [5] with sum 5
// - [3,3] with sum 6
// - [3,5] with sum 8
// - [5,5] with sum 10
// - [3,3,5], with sum 11
// - [3,5,5] with sum 13
// - [3,3,5,5] with sum 16
// Ordering the sums from smallest to largest gives 3, 3, 5, 5, 6, 8, 10, 11, 13, 16. The 7th smallest is 10.

// Constraints:
//     n == nums.length
//     1 <= n <= 2 * 10^4
//     1 <= nums[i] <= 5 * 10^4
//     1 <= k <= n * (n + 1) / 2

import "fmt"

func kthSmallestSubarraySum(nums []int, k int) int {
    res, start, end := 0, 1 << 31, 0
    preSum := make([]int, len(nums) + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < len(preSum); i++ {
        end += nums[i-1]
        start = min(start, nums[i-1])
        preSum[i] = preSum[i-1] + nums[i-1]
    }
    kthSmall := func(nums []int, mid, k int, preSum []int) bool {
        count, low := 0, 0
        for right := range preSum {
            for preSum[right] - preSum[low] > mid { // 表示以right为右边界条件下满足和小于等于mid的子数组个数
                low++
            }
            count += (right - low)
        }
        return count >= k
    }
    for start <= end {
        mid := (start + end) / 2
        if kthSmall(nums, mid, k, preSum) {
            res, end = mid, mid - 1
        } else {
            start = mid + 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,1,3], k = 4
    // Output: 3
    // Explanation: The subarrays of [2,1,3] are:
    // - [2] with sum 2
    // - [1] with sum 1
    // - [3] with sum 3
    // - [2,1] with sum 3
    // - [1,3] with sum 4
    // - [2,1,3] with sum 6 
    // Ordering the sums from smallest to largest gives 1, 2, 3, 3, 4, 6. The 4th smallest is 3.
    fmt.Println(kthSmallestSubarraySum([]int{2,1,3}, 4)) // 3
    // Example 2:
    // Input: nums = [3,3,5,5], k = 7
    // Output: 10
    // Explanation: The subarrays of [3,3,5,5] are:
    // - [3] with sum 3
    // - [3] with sum 3
    // - [5] with sum 5
    // - [5] with sum 5
    // - [3,3] with sum 6
    // - [3,5] with sum 8
    // - [5,5] with sum 10
    // - [3,3,5], with sum 11
    // - [3,5,5] with sum 13
    // - [3,3,5,5] with sum 16
    // Ordering the sums from smallest to largest gives 3, 3, 5, 5, 6, 8, 10, 11, 13, 16. The 7th smallest is 10.
    fmt.Println(kthSmallestSubarraySum([]int{3,3,5,5}, 7)) // 10
}