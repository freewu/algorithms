package main

// 2036. Maximum Alternating Subarray Sum
// A subarray of a 0-indexed integer array is a contiguous non-empty sequence of elements within an array.

// The alternating subarray sum of a subarray that ranges from index i to j (inclusive, 0 <= i <= j < nums.length) is nums[i] - nums[i+1] + nums[i+2] - ... +/- nums[j].

// Given a 0-indexed integer array nums, return the maximum alternating subarray sum of any subarray of nums.

// Example 1:
// Input: nums = [3,-1,1,2]
// Output: 5
// Explanation:
// The subarray [3,-1,1] has the largest alternating subarray sum.
// The alternating subarray sum is 3 - (-1) + 1 = 5.

// Example 2:
// Input: nums = [2,2,2,2,2]
// Output: 2
// Explanation:
// The subarrays [2], [2,2,2], and [2,2,2,2,2] have the largest alternating subarray sum.
// The alternating subarray sum of [2] is 2.
// The alternating subarray sum of [2,2,2] is 2 - 2 + 2 = 2.
// The alternating subarray sum of [2,2,2,2,2] is 2 - 2 + 2 - 2 + 2 = 2.

// Example 3:
// Input: nums = [1]
// Output: 1
// Explanation:
// There is only one non-empty subarray, which is [1].
// The alternating subarray sum is 1.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^5 <= nums[i] <= 10^5

import "fmt"

func maximumAlternatingSubarraySum(nums []int) int64 {
    res, n := nums[0], len(nums)
    arr1, arr2 := make([]int, n), make([]int, n)
    arr1[0], arr2[0] = nums[0], 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < n; i++ {
        arr1[i] = max(arr2[i - 1] + nums[i], nums[i])
        arr2[i] = arr1[i - 1] - nums[i]
        res = max(res, max(arr1[i], arr2[i]))
    }
    return int64(res)
}

func maximumAlternatingSubarraySum1(nums []int) int64 {
    res, pos, neg := nums[0], nums[0], 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(nums); i++ {
         npos, nneg := max(nums[i], neg+nums[i]), pos-nums[i]
         pos, neg = npos, nneg
         res = max(res, max(pos, neg))
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [3,-1,1,2]
    // Output: 5
    // Explanation:
    // The subarray [3,-1,1] has the largest alternating subarray sum.
    // The alternating subarray sum is 3 - (-1) + 1 = 5.
    fmt.Println(maximumAlternatingSubarraySum([]int{3,-1,1,2})) // 5
    // Example 2:
    // Input: nums = [2,2,2,2,2]
    // Output: 2
    // Explanation:
    // The subarrays [2], [2,2,2], and [2,2,2,2,2] have the largest alternating subarray sum.
    // The alternating subarray sum of [2] is 2.
    // The alternating subarray sum of [2,2,2] is 2 - 2 + 2 = 2.
    // The alternating subarray sum of [2,2,2,2,2] is 2 - 2 + 2 - 2 + 2 = 2.
    fmt.Println(maximumAlternatingSubarraySum([]int{2,2,2,2,2})) // 2
    // Example 3:
    // Input: nums = [1]
    // Output: 1
    // Explanation:
    // There is only one non-empty subarray, which is [1].
    // The alternating subarray sum is 1.
    fmt.Println(maximumAlternatingSubarraySum([]int{1})) // 1

    fmt.Println(maximumAlternatingSubarraySum1([]int{3,-1,1,2})) // 5
    fmt.Println(maximumAlternatingSubarraySum1([]int{2,2,2,2,2})) // 2
    fmt.Println(maximumAlternatingSubarraySum1([]int{1})) // 1
}