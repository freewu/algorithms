package main

// 4072. Maximum Alternating Subarray Sum With One Deletion
// You are given an integer array nums.

// You may delete at most one element from nums, then choose a subarray of the resulting array.

// Return the maximum possible alternating sum of the chosen subarray.

// The alternating sum of an array is the sum of its elements at even indices minus the sum of its elements at odd indices. 
// The chosen subarray is reindexed starting from 0 before calculating its alternating sum.

// Example 1:
// Input: nums = [5,-5,1]
// Output: 11
// Explanation:
// Choose not to delete an element and select the entire array. 
// Its alternating sum is 5 - (-5) + 1 = 11, which is the maximum possible.

// Example 2:
// Input: nums = [10,-5,-100]
// Output: 110
// Explanation:
// Delete nums[1] = -5 to obtain [10,-100], then select the entire resulting array. 
// Its alternating sum is 10 - (-100) = 110, which is the maximum possible.

// Example 3:
// Input: nums = [4,7]
// Output: 7
// Explanation:
// Choose not to delete an element and select the subarray [7]. 
// Its alternating sum is 7, which is the maximum possible.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^5 <= nums[i] <= 10^5

import "fmt"

func maxAlternatingSum(nums []int) int64 {
    const INF int64 = 1 << 61
    res, d00, d01, d10, d11 := -INF, -INF, -INF, -INF, -INF
    for _, v := range nums {
        x := int64(v)
        n00 := max(x, d01+x)
        n01 := d00 - x
        n10 := max(d11+x, d00)
        n11 := max(d10-x, d01)
        d00, d01, d10, d11 = n00, n01, n10, n11
        res = max(res, max(max(d00, d01), max(d10, d11)))
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,-5,1]
    // Output: 11
    // Explanation:
    // Choose not to delete an element and select the entire array. 
    // Its alternating sum is 5 - (-5) + 1 = 11, which is the maximum possible.
    fmt.Println(maxAlternatingSum([]int{5,-5,1})) // 11
    // Example 2:
    // Input: nums = [10,-5,-100]
    // Output: 110
    // Explanation:
    // Delete nums[1] = -5 to obtain [10,-100], then select the entire resulting array. 
    // Its alternating sum is 10 - (-100) = 110, which is the maximum possible.
    fmt.Println(maxAlternatingSum([]int{10,-5,-100})) // 110
    // Example 3:
    // Input: nums = [4,7]
    // Output: 7
    // Explanation:
    // Choose not to delete an element and select the subarray [7]. 
    // Its alternating sum is 7, which is the maximum possible.
    fmt.Println(maxAlternatingSum([]int{4,7})) // 7

    fmt.Println(maxAlternatingSum([]int{1,2,3,4,5,6,7,8,9})) // 9
    fmt.Println(maxAlternatingSum([]int{9,8,7,6,5,4,3,2,1})) // 9
}