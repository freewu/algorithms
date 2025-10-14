package main

// 3349. Adjacent Increasing Subarrays Detection I
// Given an array nums of n integers and an integer k, 
// determine whether there exist two adjacent subarrays of length k such that both subarrays are strictly increasing.
// Specifically, check if there are two subarrays starting at indices a and b (a < b), where:
//     1. Both subarrays nums[a..a + k - 1] and nums[b..b + k - 1] are strictly increasing.
//     2. The subarrays must be adjacent, meaning b = a + k.

// Return true if it is possible to find two such subarrays, and false otherwise.

// Example 1:
// Input: nums = [2,5,7,8,9,2,3,4,3,1], k = 3
// Output: true
// Explanation:
// The subarray starting at index 2 is [7, 8, 9], which is strictly increasing.
// The subarray starting at index 5 is [2, 3, 4], which is also strictly increasing.
// These two subarrays are adjacent, so the result is true.

// Example 2:
// Input: nums = [1,2,3,4,4,4,4,5,6,7], k = 5
// Output: false

// Constraints:
//     2 <= nums.length <= 100
//     1 < 2 * k <= nums.length
//     -1000 <= nums[i] <= 1000

import "fmt"

func hasIncreasingSubarrays(nums []int, k int) bool {
    n := len(nums)
    if n < 2 * k { return false }
    isStrictlyInc := func(nums []int, start int, k int) bool {
        for i := start; i < start + k - 1; i++ {
            if nums[i] >= nums[i + 1] {
                return false
            }
        }
        return true
    }
    for i := 0; i <= n-2*k; i++ {
        if isStrictlyInc(nums, i, k) && isStrictlyInc(nums, i+k, k) {
            return true
        }
    }
    return false
}

func hasIncreasingSubarrays1(nums []int, k int) bool {
    n := len(nums)
    dp := make([]int, n)
    dp[0] = 1
    for i := 1; i < n; i++ {
        dp[i] = 1
        if nums[i] > nums[i-1] {
            dp[i] += dp[i-1]
        }
        if i + 1 >= 2* k && dp[i] >= k && dp[i - k] >= k {
            return true
        }
    }
    return false
}

func hasIncreasingSubarrays2(nums []int, k int) bool {
    if k == 1 { return true }
    n := len(nums)
    dp := make([]int, n)
    dp[0] = 1
    for i := 1; i < n; i++ {
        if nums[i] > nums[i - 1] {
            dp[i] = dp[i - 1] + 1
        } else {
            dp[i] = 1
        }
    }
    for i := 0; i + k + k - 1 < n; i++ {
        if dp[i + k - 1]-dp[i]+1 == k && dp[i+k-1]-dp[i] == dp[i+k+k-1]-dp[i+k] { 
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: nums = [2,5,7,8,9,2,3,4,3,1], k = 3
    // Output: true
    // Explanation:
    // The subarray starting at index 2 is [7, 8, 9], which is strictly increasing.
    // The subarray starting at index 5 is [2, 3, 4], which is also strictly increasing.
    // These two subarrays are adjacent, so the result is true.
    fmt.Println(hasIncreasingSubarrays([]int{2,5,7,8,9,2,3,4,3,1}, 3)) // true
    // Example 2:
    // Input: nums = [1,2,3,4,4,4,4,5,6,7], k = 5
    // Output: false
    fmt.Println(hasIncreasingSubarrays([]int{1,2,3,4,4,4,4,5,6,7}, 5)) // false

    fmt.Println(hasIncreasingSubarrays([]int{1,2,3,4,5,6,7,8,9}, 2)) // true
    fmt.Println(hasIncreasingSubarrays([]int{9,8,7,6,5,4,3,2,1}, 2)) // false

    fmt.Println(hasIncreasingSubarrays1([]int{2,5,7,8,9,2,3,4,3,1}, 3)) // true
    fmt.Println(hasIncreasingSubarrays1([]int{1,2,3,4,4,4,4,5,6,7}, 5)) // false
    fmt.Println(hasIncreasingSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 2)) // true
    fmt.Println(hasIncreasingSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 2)) // false

    fmt.Println(hasIncreasingSubarrays2([]int{2,5,7,8,9,2,3,4,3,1}, 3)) // true
    fmt.Println(hasIncreasingSubarrays2([]int{1,2,3,4,4,4,4,5,6,7}, 5)) // false
    fmt.Println(hasIncreasingSubarrays2([]int{1,2,3,4,5,6,7,8,9}, 2)) // true
    fmt.Println(hasIncreasingSubarrays2([]int{9,8,7,6,5,4,3,2,1}, 2)) // false
}