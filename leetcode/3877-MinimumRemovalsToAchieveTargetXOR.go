package main

// 3877. Minimum Removals to Achieve Target XOR
// You are given an integer array nums and an integer target.

// You may remove any number of elements from nums (possibly zero).

// Return the minimum number of removals required so that the bitwise XOR of the remaining elements equals target. 
// If it is impossible to achieve target, return -1.

// The bitwise XOR of an empty array is 0.

// Example 1:
// Input: nums = [1,2,3], target = 2
// Output: 1
// Explanation:
// Removing nums[1] = 2 leaves [nums[0], nums[2]] = [1, 3].
// The XOR of [1, 3] is 2, which equals target.
// It is not possible to achieve XOR = 2 in less than one removal, therefore the answer is 1.

// Example 2:
// Input: nums = [2,4], target = 1
// Output: -1
// Explanation:
// It is impossible to remove elements to achieve target. Thus, the answer is -1.

// Example 3:
// Input: nums = [7], target = 7
// Output: 0
// Explanation:
// The XOR of all elements is nums[0] = 7, which equals target. Thus, no removal is needed.

// Constraints:
//     1 <= nums.length <= 40
//     0 <= nums[i] <= 10^4
//     0 <= target <= 10^4

import "fmt"
import "math/bits"
import "slices"

func minRemovals(nums []int, target int) int {
    m := bits.Len(uint(slices.Max(nums)))
    if 1 << m <= target {
        return -1
    }
    n, inf := len(nums), 1 << 61
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, 1 << m)
        for j := range dp[i] {
            dp[i][j] = -inf
        }
    }
    dp[0][0] = 0
    for i, x := range nums {
        for j := range 1 << m {
            dp[i + 1][j] = max(dp[i][j], dp[i][j^x]+1) // x 不选 or 选
        }
    }
    if dp[n][target] < 0 {
        return -1
    }
    return n - dp[n][target]
}

func minRemovals1(nums []int, target int) int {
    mx := target
    for _, v := range nums {
        mx |= v
    }
    dp := make([]int, mx + 1)
    for i := range dp {
        dp[i] = -1
    }
    dp[0] = 0
    for _,  v := range nums {
        tf := slices.Clone(dp)
        for i, y := range dp {
            if y == -1 { continue }
            tf[i ^ v] = max(tf[i ^ v], y + 1)
        }
        dp = tf
    }
    if dp[target] == -1 {
        return -1
    }
    return len(nums) - dp[target]
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3], target = 2
    // Output: 1
    // Explanation:
    // Removing nums[1] = 2 leaves [nums[0], nums[2]] = [1, 3].
    // The XOR of [1, 3] is 2, which equals target.
    // It is not possible to achieve XOR = 2 in less than one removal, therefore the answer is 1.
    fmt.Println(minRemovals([]int{1,2,3}, 2)) // 1
    // Example 2:
    // Input: nums = [2,4], target = 1
    // Output: -1
    // Explanation:
    // It is impossible to remove elements to achieve target. Thus, the answer is -1.
    fmt.Println(minRemovals([]int{2,4}, 1)) // -1
    // Example 3:
    // Input: nums = [7], target = 7
    // Output: 0
    // Explanation:
    // The XOR of all elements is nums[0] = 7, which equals target. Thus, no removal is needed.
    fmt.Println(minRemovals([]int{7}, 7)) // 0

    fmt.Println(minRemovals([]int{1,2,3,4,5,6,7,8,9}, 2)) // 1
    fmt.Println(minRemovals([]int{9,8,7,6,5,4,3,2,1}, 2)) // 1

    fmt.Println(minRemovals1([]int{1,2,3}, 2)) // 1
    fmt.Println(minRemovals1([]int{2,4}, 1)) // -1
    fmt.Println(minRemovals1([]int{7}, 7)) // 0
    fmt.Println(minRemovals1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 1
    fmt.Println(minRemovals1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 1
}