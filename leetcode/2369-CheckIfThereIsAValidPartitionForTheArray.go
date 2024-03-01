package main

// 2369. Check if There is a Valid Partition For The Array
// You are given a 0-indexed integer array nums. You have to partition the array into one or more contiguous subarrays.
// We call a partition of the array valid if each of the obtained subarrays satisfies one of the following conditions:
//         The subarray consists of exactly 2, equal elements. For example, the subarray [2,2] is good.
//         The subarray consists of exactly 3, equal elements. For example, the subarray [4,4,4] is good.
//         The subarray consists of exactly 3 consecutive increasing elements, that is, the difference between adjacent elements is 1. For example, the subarray [3,4,5] is good, but the subarray [1,3,5] is not.

// Return true if the array has at least one valid partition. Otherwise, return false.

// Example 1:
// Input: nums = [4,4,4,5,6]
// Output: true
// Explanation: The array can be partitioned into the subarrays [4,4] and [4,5,6].
// This partition is valid, so we return true.

// Example 2:
// Input: nums = [1,1,1,2]
// Output: false
// Explanation: There is no valid partition for this array.

// Constraints:
//         2 <= nums.length <= 10^5
//         1 <= nums[i] <= 10^6

import "fmt"

func validPartition(nums []int) bool {
    n := len(nums)
    if n == 1 {
        return false
    }
    dp := []bool{ true, false, n > 1 && nums[0] == nums[1]}
    for i := 2; i < n; i++ {
        current_dp := false
        if nums[i] == nums[i-1] && dp[1] {
            current_dp = true
        } else if nums[i] == nums[i-1] && nums[i] == nums[i-2] && dp[0] {
            current_dp = true
        } else if nums[i] - nums[i-1] == 1 && nums[i-1] - nums[i-2] == 1 && dp[0] {
            current_dp = true
        }
        dp[0], dp[1], dp[2] = dp[1], dp[2], current_dp
    }
    return dp[2]
}

func validPartition1(nums []int) bool {
    n := len(nums)
    // 用动态规划来解决，用一个长度为 (n+1)(n+1)(n+1) 的数组来记录 nums 是否存在有效划分，
    // dp[i]表示前 i 个元素组成的数组是否至少存在一个有效划分。
    // 边界情况 dp[0] 恒为 true，而 dp[n] 即为结果
    dp := make([]bool, n+1)
    dp[0] = true
    validTwo := func (num1, num2 int) bool {
        return num1 == num2
    }
    validThree := func (num1, num2, num3 int) bool {
        return (num1 == num2 && num1 == num3) || (num1+1 == num2 && num2+1 == num3)
    }
    for i := 1; i <= n; i++ {
        // 从第 2 位开始 判读是否能组成 [n，n] 的 
        if i >= 2 {
            // dp[i - 2] &&  关键
            dp[i] = dp[i - 2] && validTwo(nums[i - 2], nums[i - 1])
        }
        // 从第 3 位开始 判读是否能组成 [n, n + 1, n + 2] 的
        if i >= 3 {
            // 关键 dp[i - 3] && 
            dp[i] = dp[i] || (dp[i - 3] && validThree(nums[i - 3], nums[i - 2], nums[i - 1]))
        }
        // fmt.Println(dp)
    }
    return dp[n]
}

func main() {
    fmt.Println(validPartition([]int{4,4,4,5,6})) // true
    fmt.Println(validPartition([]int{1,1,1,2})) // false

    fmt.Println(validPartition1([]int{4,3,4,5,6})) // false
    fmt.Println(validPartition1([]int{4,4,4,5,6})) // true
    fmt.Println(validPartition1([]int{1,1,1,2})) // false
}