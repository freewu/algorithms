package main

// 3976. Maximum Subarray Sum After Multiplier
// You are given an integer array nums and a positive integer k.

// You must choose exactly one subarray of nums and perform exactly one of the following operations:
//     1. Multiply each number in the chosen subarray by k.
//     2. Divide each number in the chosen subarray by k.
//         2.1 When dividing a positive number by k, use the floor value of the division result.
//         2.2 When dividing a negative number by k, use the ceiling value of the division result.

// Return the maximum possible sum of a non-empty subarray in the resulting array.

// Note that the subarray chosen for the operation and the subarray chosen for the sum may be different.

// Example 1:
// Input: nums = [1,-2,3,4,-5], k = 2
// Output: 14
// Explanation:
// Multiply each number in the subarray [3, 4] by 2.
// This results in nums = [1, -2, 6, 8, -5].
// The subarray with the largest sum is [6, 8], so the output is 6 + 8 = 14.

// Example 2:
// Input: nums = [-5,-4,-3], k = 2
// Output: -1
// Explanation:
// Divide each number in the subarray [-3] by 2.
// This results in nums = [-5, -4, -1].
// The subarray with the largest sum is [-1], so the output is -1.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^5 <= nums[i] <= 10^5
//     1 <= k <= 10^5

import "fmt"
import "slices"

func maxSubarraySum(nums []int, k int) int64 {
    helper := func(isMul bool) int64 {
        res, n := int64(-1 << 61), len(nums)
        // f[i+1][0] 表示右端点为 i 的最大子数组和，且不修改任何元素
        // f[i+1][1] 表示右端点为 i 的最大子数组和，且修改了 nums[i]
        // f[i+1][2] 表示右端点为 i 的最大子数组和，且在 nums[i] 的左边发生了修改（没有修改 nums[i]）
        f := make([][3]int64, n+1)
        for i, x := range nums {
            x := int64(x)
            y := x
            if isMul {
                y *= int64(k)
            } else {
                y /= int64(k)
            }
            // 不修改 x，和 f[i][0] 拼起来，或者 x 是子数组的第一个数
            f[i+1][0] = max(f[i][0], 0) + x
            // 修改 x，和 f[i][0] 或者 f[i][1] 拼起来，或者 y 是子数组的第一个数
            f[i+1][1] = max(f[i][0], f[i][1], 0) + y
            // 不修改 x，和 f[i][1] 或者 f[i][2] 拼起来
            f[i+1][2] = max(f[i][1], f[i][2]) + x
            // 枚举子数组的右端点为 i
            res = max(res, f[i+1][1], f[i+1][2])
        }
        return res
    }
    return max(helper(true), helper(false))
}

func maxSubarraySum1(nums []int, k int) int64 {
    n, c, m, d := len(nums), int64(nums[0]), int64(nums[0] * k), int64(nums[0] / k)
    dp := [5]int64 { c, m, m, d, d }
    res := max(max(c, d), m)
    for i := 1; i < n; i++ {
        c, m, d = int64(nums[i]), int64(nums[i] * k), int64(nums[i] / k)
        next := [5]int64{}
        next[0] = max(c, dp[0] + c)
        next[1] = max(m, max(dp[1] + m, dp[0] + m))
        next[2] = max(c, max(dp[1] + c, dp[2] + c))
        next[3] = max(d, max(dp[3] + d, dp[0] + d))
        next[4] = max(c, max(dp[3] + c, dp[4] + c))
        dp = next
        for j := range next {
            res = max(res, next[j])
        }
    }
    return res
}

func maxSubarraySum2(nums []int, k int) int64 {
    f0, f1, f2, f3 := 0, 0, 0, 0
    res := slices.Max(nums)
    for _, v := range nums {
        f0, f1, f2, f3 = max(v, f0 + v) ,max(v*k, f0+v*k, f1+v*k) ,max(v/k, f0+v/k, f2+v/k) ,max(v, f1+v, f2+v, f3+v)
        res = max(res, f0, f1, f2, f3)
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,-2,3,4,-5], k = 2
    // Output: 14
    // Explanation:
    // Multiply each number in the subarray [3, 4] by 2.
    // This results in nums = [1, -2, 6, 8, -5].
    // The subarray with the largest sum is [6, 8], so the output is 6 + 8 = 14.
    fmt.Println(maxSubarraySum([]int{1,-2,3,4,-5}, 2)) // 14
    // Example 2:
    // Input: nums = [-5,-4,-3], k = 2
    // Output: -1
    // Explanation:
    // Divide each number in the subarray [-3] by 2.
    // This results in nums = [-5, -4, -1].
    // The subarray with the largest sum is [-1], so the output is -1.
    fmt.Println(maxSubarraySum([]int{-5,-4,-3}, 2)) // -1

    fmt.Println(maxSubarraySum([]int{1,2,3,4,5,6,7,8,9}, 2)) // 90
    fmt.Println(maxSubarraySum([]int{9,8,7,6,5,4,3,2,1}, 2)) // 90
    fmt.Println(maxSubarraySum([]int{1,2,3,4,5,6,7,8,9}, 1)) // 45
    fmt.Println(maxSubarraySum([]int{9,8,7,6,5,4,3,2,1}, 1)) // 45
    fmt.Println(maxSubarraySum([]int{1,2,3,4,5,6,7,8,9}, 100_000)) // 4500000
    fmt.Println(maxSubarraySum([]int{9,8,7,6,5,4,3,2,1}, 100_000)) // 4500000

    fmt.Println(maxSubarraySum1([]int{1,-2,3,4,-5}, 2)) // 14
    fmt.Println(maxSubarraySum1([]int{-5,-4,-3}, 2)) // -1
    fmt.Println(maxSubarraySum1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 90
    fmt.Println(maxSubarraySum1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 90
    fmt.Println(maxSubarraySum1([]int{1,2,3,4,5,6,7,8,9}, 1)) // 45
    fmt.Println(maxSubarraySum1([]int{9,8,7,6,5,4,3,2,1}, 1)) // 45
    fmt.Println(maxSubarraySum1([]int{1,2,3,4,5,6,7,8,9}, 100_000)) // 4500000
    fmt.Println(maxSubarraySum1([]int{9,8,7,6,5,4,3,2,1}, 100_000)) // 4500000

    fmt.Println(maxSubarraySum2([]int{1,-2,3,4,-5}, 2)) // 14
    fmt.Println(maxSubarraySum2([]int{-5,-4,-3}, 2)) // -1
    fmt.Println(maxSubarraySum2([]int{1,2,3,4,5,6,7,8,9}, 2)) // 90
    fmt.Println(maxSubarraySum2([]int{9,8,7,6,5,4,3,2,1}, 2)) // 90
    fmt.Println(maxSubarraySum2([]int{1,2,3,4,5,6,7,8,9}, 1)) // 45
    fmt.Println(maxSubarraySum2([]int{9,8,7,6,5,4,3,2,1}, 1)) // 45
    fmt.Println(maxSubarraySum2([]int{1,2,3,4,5,6,7,8,9}, 100_000)) // 4500000
    fmt.Println(maxSubarraySum2([]int{9,8,7,6,5,4,3,2,1}, 100_000)) // 4500000
}