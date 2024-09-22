package main

// 3247. Number of Subsequences with Odd Sum
// Given an array nums, return the number of subsequences with an odd sum of elements.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [1,1,1]
// Output: 4
// Explanation:
// The odd-sum subsequences are: [1, 1, 1], [1, 1, 1], [1, 1, 1], [1, 1, 1].

// Example 2:
// Input: nums = [1,2,2]
// Output: 4
// Explanation:
// The odd-sum subsequences are: [1, 2, 2], [1, 2, 2], [1, 2, 2], [1, 2, 2].

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"

func subsequenceCount(nums []int) int {
    n, mod := len(nums), 1_000_000_007
    dp := make([][]int, n + 1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, 2)
    }
    dp[0][0] = 1
    for i := 1; i <= n; i++ {
        if nums[i - 1] % 2 == 0 {
            dp[i][0] = dp[i - 1][0] * 2 % mod
            dp[i][1] = dp[i - 1][1] * 2 % mod
        } else { // 分别计算数组 nums 的每个前缀的元素和为奇数的子序列的数量
            dp[i][0] = (dp[i - 1][0] + dp[i - 1][1]) % mod
            dp[i][1] = (dp[i - 1][1] + dp[i - 1][0]) % mod
        }
    }
    return dp[n][1]
}

func subsequenceCount1(nums []int) int {
    odd, even, mod := 0, 0, int64(1_000_000_007)
    for _, v := range nums { // 分别统计 奇数 & 偶数个数
        if v % 2 != 0 { 
            odd++
        } else {
            even++
        }
    }
    pow := func (x int, n int) int64 {
        if n == 0 { return 1 }
        if x == 0 { return 0 }
        res, base := int64(1), int64(x)
        for n != 0 {
            if n % 2 != 0 { res = res * base % mod }
            base = base * base % mod
            n /= 2
        }
        return res
    }
    if odd > 0 {
        return int(pow(2, odd - 1) * pow(2, even) % mod)
    } else {
        return 0
    }
}

func main() {
    // Example 1:
    // Input: nums = [1,1,1]
    // Output: 4
    // Explanation:
    // The odd-sum subsequences are: [1, 1, 1], [1, 1, 1], [1, 1, 1], [1, 1, 1].
    fmt.Println(subsequenceCount([]int{1,1,1})) // 4
    // Example 2:
    // Input: nums = [1,2,2]
    // Output: 4
    // Explanation:
    // The odd-sum subsequences are: [1, 2, 2], [1, 2, 2], [1, 2, 2], [1, 2, 2].
    fmt.Println(subsequenceCount([]int{1,2,2})) // 4

    fmt.Println(subsequenceCount1([]int{1,1,1})) // 4
    fmt.Println(subsequenceCount1([]int{1,2,2})) // 4
}