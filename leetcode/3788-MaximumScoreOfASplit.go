package main

// 3788. Maximum Score of a Split
// You are given an integer array nums of length n.

// Choose an index i such that 0 <= i < n - 1.

// For a chosen split index i:
//     Let prefixSum(i) be the sum of nums[0] + nums[1] + ... + nums[i].
//     Let suffixMin(i) be the minimum value among nums[i + 1], nums[i + 2], ..., nums[n - 1].

// The score of a split at index i is defined as:

//     score(i) = prefixSum(i) - suffixMin(i)

// Return an integer denoting the maximum score over all valid split indices.

// Example 1:
// Input: nums = [10,-1,3,-4,-5]
// Output: 17
// Explanation:
// The optimal split is at i = 2, score(2) = prefixSum(2) - suffixMin(2) = (10 + (-1) + 3) - (-5) = 17.

// Example 2:
// Input: nums = [-7,-5,3]
// Output: -2
// Explanation:
// The optimal split is at i = 0, score(0) = prefixSum(0) - suffixMin(0) = (-7) - (-5) = -2.

// Example 3:
// Input: nums = [1,1]
// Output: 0
// Explanation:
// The only valid split is at i = 0, score(0) = prefixSum(0) - suffixMin(0) = 1 - 1 = 0.

// Constraints:
//     2 <= nums.length <= 10^5
//     -10^9​​​​​​​ <= nums[i] <= 10^9

import "fmt"

// time: O(n), space: O(1)
func maximumScore(nums []int) int64 {
    sum, res, mn := 0, -1 << 61, 1 << 61
    for _, v := range nums {
        sum += v
    }
    for i := len(nums) - 1; i > 0; i-- {
        sum -= nums[i]
        mn = min(mn, nums[i])
        res = max(res, sum - mn)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [10,-1,3,-4,-5]
    // Output: 17
    // Explanation:
    // The optimal split is at i = 2, score(2) = prefixSum(2) - suffixMin(2) = (10 + (-1) + 3) - (-5) = 17.
    fmt.Println(maximumScore([]int{10,-1,3,-4,-5})) // 17
    // Example 2:
    // Input: nums = [-7,-5,3]
    // Output: -2
    // Explanation:
    // The optimal split is at i = 0, score(0) = prefixSum(0) - suffixMin(0) = (-7) - (-5) = -2.
    fmt.Println(maximumScore([]int{-7,-5,3})) // -2
    // Example 3:
    // Input: nums = [1,1]
    // Output: 0
    // Explanation:
    // The only valid split is at i = 0, score(0) = prefixSum(0) - suffixMin(0) = 1 - 1 = 0.
    fmt.Println(maximumScore([]int{1,1})) // 0

    fmt.Println(maximumScore([]int{-7,-5,3})) // -2

    fmt.Println(maximumScore([]int{1,2,3,4,5,6,7,8,9})) // 27
    fmt.Println(maximumScore([]int{9,8,7,6,5,4,3,2,1})) // 43
}