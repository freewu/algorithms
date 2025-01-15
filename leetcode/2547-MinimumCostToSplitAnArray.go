package main

// 2547. Minimum Cost to Split an Array
// You are given an integer array nums and an integer k.

// Split the array into some number of non-empty subarrays. 
// The cost of a split is the sum of the importance value of each subarray in the split.

// Let trimmed(subarray) be the version of the subarray where all numbers which appear only once are removed.
//     For example, trimmed([3,1,2,4,3,4]) = [3,4,3,4].

// The importance value of a subarray is k + trimmed(subarray).length.
//     For example, if a subarray is [1,2,3,3,3,4,4], then trimmed([1,2,3,3,3,4,4]) = [3,3,3,4,4].
//     The importance value of this subarray will be k + 5.

// Return the minimum possible cost of a split of nums.

// A subarray is a contiguous non-empty sequence of elements within an array.

// Example 1:
// Input: nums = [1,2,1,2,1,3,3], k = 2
// Output: 8
// Explanation: We split nums to have two subarrays: [1,2], [1,2,1,3,3].
// The importance value of [1,2] is 2 + (0) = 2.
// The importance value of [1,2,1,3,3] is 2 + (2 + 2) = 6.
// The cost of the split is 2 + 6 = 8. It can be shown that this is the minimum possible cost among all the possible splits.

// Example 2:
// Input: nums = [1,2,1,2,1], k = 2
// Output: 6
// Explanation: We split nums to have two subarrays: [1,2], [1,2,1].
// The importance value of [1,2] is 2 + (0) = 2.
// The importance value of [1,2,1] is 2 + (2) = 4.
// The cost of the split is 2 + 4 = 6. It can be shown that this is the minimum possible cost among all the possible splits.

// Example 3:
// Input: nums = [1,2,1,2,1], k = 5
// Output: 10
// Explanation: We split nums to have one subarray: [1,2,1,2,1].
// The importance value of [1,2,1,2,1] is 5 + (3 + 2) = 10.
// The cost of the split is 10. It can be shown that this is the minimum possible cost among all the possible splits.

// Constraints:
//     1 <= nums.length <= 1000
//     0 <= nums[i] < nums.length
//     1 <= k <= 10^9

import "fmt"

func minCost(nums []int, k int) int {
    onceCount, n := 1, len(nums)
    dp, count, tmp := make([]int, n), make([]int, n), make([]int, n)
    dp[0], count[nums[0]] = k, 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < n; i++ {
        v := nums[i]
        switch count[v] {
        case 0:
            onceCount++
        case 1:
            onceCount--
        }
        count[v]++
        dp[i] = (i + 1 - onceCount + k)
        if onceCount <= i {
            copy(tmp, count)
            onceCountTmp := onceCount
            for j := 0; j < i; j++ {
                v = nums[j]
                switch tmp[v] {
                case 1:
                    onceCountTmp--
                case 2:
                    onceCountTmp++
                }
                tmp[v]--
                dp[i] = min(dp[i], dp[j] + int(i - j - onceCountTmp + k))
            }
        }
    }
    return int(dp[n - 1])
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1,2,1,3,3], k = 2
    // Output: 8
    // Explanation: We split nums to have two subarrays: [1,2], [1,2,1,3,3].
    // The importance value of [1,2] is 2 + (0) = 2.
    // The importance value of [1,2,1,3,3] is 2 + (2 + 2) = 6.
    // The cost of the split is 2 + 6 = 8. It can be shown that this is the minimum possible cost among all the possible splits.
    fmt.Println(minCost([]int{1,2,1,2,1,3,3}, 2)) // 8
    // Example 2:
    // Input: nums = [1,2,1,2,1], k = 2
    // Output: 6
    // Explanation: We split nums to have two subarrays: [1,2], [1,2,1].
    // The importance value of [1,2] is 2 + (0) = 2.
    // The importance value of [1,2,1] is 2 + (2) = 4.
    // The cost of the split is 2 + 4 = 6. It can be shown that this is the minimum possible cost among all the possible splits.
    fmt.Println(minCost([]int{1,2,1,2,1}, 2)) // 6
    // Example 3:
    // Input: nums = [1,2,1,2,1], k = 5
    // Output: 10
    // Explanation: We split nums to have one subarray: [1,2,1,2,1].
    // The importance value of [1,2,1,2,1] is 5 + (3 + 2) = 10.
    // The cost of the split is 10. It can be shown that this is the minimum possible cost among all the possible splits.
    fmt.Println(minCost([]int{1,2,1,2,1}, 5)) // 10
}