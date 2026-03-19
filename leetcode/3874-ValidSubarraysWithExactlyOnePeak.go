package main

// 3874. Valid Subarrays With Exactly One Peak
// You are given an integer array nums of length n and an integer k.

// An index i is a peak if:
//     1. 0 < i < n - 1
//     2. nums[i] > nums[i - 1] and nums[i] > nums[i + 1]
    
// A subarray [l, r] is valid if:
//     1. It contains exactly one peak at index i from nums
//     2. i - l <= k and r - i <= k

// Return an integer denoting the number of valid subarrays in nums.

// A subarray is a contiguous non-empty sequence of elements within an array.
 
// Example 1:
// Input: nums = [1,3,2], k = 1
// Output: 4
// Explanation:
// Index i = 1 is a peak because nums[1] = 3 is greater than nums[0] = 1 and nums[2] = 2.
// Any valid subarray must include index 1, and the distance from the peak to both ends of the subarray must not exceed k = 1.
// The valid subarrays are [3], [1, 3], [3, 2], and [1, 3, 2], so the answer is 4.

// Example 2:
// Input: nums = [7,8,9], k = 2
// Output: 0
// Explanation:
// There is no index i such that nums[i] is greater than both nums[i - 1] and nums[i + 1].
// Therefore, the array contains no peak. Thus, the number of valid subarrays is 0.

// Example 3:
// Input: nums = [4,3,5,1], k = 2
// Output: 6
// Explanation:
// Index i = 2 is a peak because nums[2] = 5 is greater than nums[1] = 3 and nums[3] = 1.
// Any valid subarray must contain this peak, and the distance from the peak to both ends of the subarray must not exceed k = 2.
// The valid subarrays are [5], [3, 5], [5, 1], [3, 5, 1], [4, 3, 5], and [4, 3, 5, 1], so the answer is 6.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     -10^5 <= nums[i] <= 10^5
//     1 <= k <= n

import "fmt"

func validSubarrays(nums []int, k int) int64 {
    res, n := 0, len(nums)
    if n < 3 { return 0 }
    // 先找出所有的峰值位置
    peaks := make([]int, 0)
    for i := 1; i < n-1; i++ {
        if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
            peaks = append(peaks, i)
        }
    }
    if len(peaks) == 0 { return 0 } // 如果没有峰值，直接返回 0
    
    // 预处理峰值的区间，避免重复计算（确保每个子数组只包含一个峰值）
    // 为每个峰值确定其"专属区间"：左边到前一个峰值的下一个位置，右边到后一个峰值的前一个位置
    peakRanges := make([][2]int, len(peaks))
    for i := range peaks {
        // 左边界：前一个峰值的下一个位置（如果没有前一个峰值，则为0）
        left := 0
        if i > 0 {
            left = peaks[i-1] + 1
        }
        // 右边界：后一个峰值的前一个位置（如果没有后一个峰值，则为n-1）
        right := n - 1
        if i < len(peaks)-1 {
            right = peaks[i+1] - 1
        }
        peakRanges[i] = [2]int{left, right}
    }
    // 计算每个峰值对应的有效子数组数量
    for i, peak := range peaks {
        left, right := peakRanges[i][0], peakRanges[i][1] 
        // 计算当前峰值i的有效左选择数：l的可选范围
        // l的范围：max(left, peak - k) ≤ l ≤ peak
        l := peak - max(left, peak - k) + 1
        // 计算当前峰值i的有效右选择数：r的可选范围
        // r的范围：peak ≤ r ≤ min(right, peak + k)
        r := min(right, peak + k) - peak + 1
        // 有效子数组数量 = 左选择数 * 右选择数
        res += (l * r)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2], k = 1
    // Output: 4
    // Explanation:
    // Index i = 1 is a peak because nums[1] = 3 is greater than nums[0] = 1 and nums[2] = 2.
    // Any valid subarray must include index 1, and the distance from the peak to both ends of the subarray must not exceed k = 1.
    // The valid subarrays are [3], [1, 3], [3, 2], and [1, 3, 2], so the answer is 4.
    fmt.Println(validSubarrays([]int{1,3,2}, 1)) // 4
    // Example 2:
    // Input: nums = [7,8,9], k = 2
    // Output: 0
    // Explanation:
    // There is no index i such that nums[i] is greater than both nums[i - 1] and nums[i + 1].
    // Therefore, the array contains no peak. Thus, the number of valid subarrays is 0.
    fmt.Println(validSubarrays([]int{7,8,9}, 2)) // 0
    // Example 3:
    // Input: nums = [4,3,5,1], k = 2
    // Output: 6
    // Explanation:
    // Index i = 2 is a peak because nums[2] = 5 is greater than nums[1] = 3 and nums[3] = 1.
    // Any valid subarray must contain this peak, and the distance from the peak to both ends of the subarray must not exceed k = 2.
    // The valid subarrays are [5], [3, 5], [5, 1], [3, 5, 1], [4, 3, 5], and [4, 3, 5, 1], so the answer is 6.
    fmt.Println(validSubarrays([]int{4,3,5,1}, 2)) // 6

    fmt.Println(validSubarrays([]int{-12,20,-20,15,-12}, 2)) // 8
    fmt.Println(validSubarrays([]int{1,2,3,4,5,6,7,8,9}, 2)) // 0
    fmt.Println(validSubarrays([]int{9,8,7,6,5,4,3,2,1}, 2)) // 0
}