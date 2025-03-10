package main

// 3410. Maximize Subarray Sum After Removing All Occurrences of One Element
// You are given an integer array nums.

// You can do the following operation on the array at most once:
//     1. Choose any integer x such that nums remains non-empty on removing all occurrences of x.
//     2. Remove all occurrences of x from the array.

// Return the maximum subarray sum across all possible resulting arrays.

// Example 1:
// Input: nums = [-3,2,-2,-1,3,-2,3]
// Output: 7
// Explanation:
// We can have the following arrays after at most one operation:
// The original array is nums = [-3, 2, -2, -1, 3, -2, 3]. The maximum subarray sum is 3 + (-2) + 3 = 4.
// Deleting all occurences of x = -3 results in nums = [2, -2, -1, 3, -2, 3]. The maximum subarray sum is 3 + (-2) + 3 = 4.
// Deleting all occurences of x = -2 results in nums = [-3, 2, -1, 3, 3]. The maximum subarray sum is 2 + (-1) + 3 + 3 = 7.
// Deleting all occurences of x = -1 results in nums = [-3, 2, -2, 3, -2, 3]. The maximum subarray sum is 3 + (-2) + 3 = 4.
// Deleting all occurences of x = 3 results in nums = [-3, 2, -2, -1, -2]. The maximum subarray sum is 2.
// The output is max(4, 4, 7, 4, 2) = 7.

// Example 2:
// Input: nums = [1,2,3,4]
// Output: 10
// Explanation:
// It is optimal to not perform any operations.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^6 <= nums[i] <= 10^6

import "fmt"

func maxSubarraySum(nums []int) int64 {
    res, sum, nonDelMinS, allMin := -1 << 31, 0, 0, 0
    mp := make(map[int]int)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        sum += v
        res = max(res, sum - allMin)
        if v < 0 {
            mp[v] = min(mp[v], nonDelMinS) + v
            allMin, nonDelMinS = min(allMin, mp[v]), min(nonDelMinS, sum)
        }
    }
    return int64(res)
}

func maxSubarraySum1(nums []int) int64 {
    prev := make(map[int]int)
    res, acc, lowest := -1 << 31, 0, 0
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        acc += v
        res = max(res, acc - lowest) // Update the result based on current accumulator and lowest seen
        if v < 0 {
            // Since Go maps return 0 for missing keys, this mimics defaultdict behavior
            cand := v + min(prev[v], prev[0])
            prev[v] = cand
            lowest = min(lowest, cand)
        }
        lowest = min(prev[0], lowest)
        prev[0] = min(prev[0], acc)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [-3,2,-2,-1,3,-2,3]
    // Output: 7
    // Explanation:
    // We can have the following arrays after at most one operation:
    // The original array is nums = [-3, 2, -2, -1, 3, -2, 3]. The maximum subarray sum is 3 + (-2) + 3 = 4.
    // Deleting all occurences of x = -3 results in nums = [2, -2, -1, 3, -2, 3]. The maximum subarray sum is 3 + (-2) + 3 = 4.
    // Deleting all occurences of x = -2 results in nums = [-3, 2, -1, 3, 3]. The maximum subarray sum is 2 + (-1) + 3 + 3 = 7.
    // Deleting all occurences of x = -1 results in nums = [-3, 2, -2, 3, -2, 3]. The maximum subarray sum is 3 + (-2) + 3 = 4.
    // Deleting all occurences of x = 3 results in nums = [-3, 2, -2, -1, -2]. The maximum subarray sum is 2.
    // The output is max(4, 4, 7, 4, 2) = 7.
    fmt.Println(maxSubarraySum([]int{-3,2,-2,-1,3,-2,3})) // 7
    // Example 2:
    // Input: nums = [1,2,3,4]
    // Output: 10
    // Explanation:
    // It is optimal to not perform any operations.
    fmt.Println(maxSubarraySum([]int{1,2,3,4})) // 10

    fmt.Println(maxSubarraySum([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxSubarraySum([]int{9,8,7,6,5,4,3,2,1})) // 45

    fmt.Println(maxSubarraySum1([]int{-3,2,-2,-1,3,-2,3})) // 7
    fmt.Println(maxSubarraySum1([]int{1,2,3,4})) // 10
    fmt.Println(maxSubarraySum1([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxSubarraySum1([]int{9,8,7,6,5,4,3,2,1})) // 45
}