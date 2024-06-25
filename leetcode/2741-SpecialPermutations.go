package main

// 2741. Special Permutations
// You are given a 0-indexed integer array nums containing n distinct positive integers.
// A permutation of nums is called special if:
//     For all indexes 0 <= i < n - 1, either nums[i] % nums[i+1] == 0 or nums[i+1] % nums[i] == 0.

// Return the total number of special permutations. As the answer could be large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [2,3,6]
// Output: 2
// Explanation: [3,6,2] and [2,6,3] are the two special permutations of nums.

// Example 2:
// Input: nums = [1,4,3]
// Output: 2
// Explanation: [3,1,4] and [4,1,3] are the two special permutations of nums.

// Constraints:
//     2 <= nums.length <= 14
//     1 <= nums[i] <= 10^9

import "fmt"
import "math/bits"

// // 超出时间限制 587 / 590
// func specialPerm(nums []int) int {
//     n, mod := len(nums), 1_000_000_007
//     m := int(1 << n)
//     dp := make([][]int, n)
//     for i := range dp {
//         dp[i] = make([]int, m)
//     }
//     var dfs func(pre int, nums []int, mask int) int
//     dfs = func(pre int, nums []int, mask int) int {
//         if (1 << n) - mask == 1 { return 1; }
//         if dp[pre][mask] != 0 { return dp[pre][mask]; }
//         count := 0
//         for i := 0; i < n; i++ {
//             if mask & (1 << i) != 0 { continue; }
//             if mask == 0 || nums[pre] % nums[i] == 0 || nums[i] % nums[pre] == 0 {
//                 count = count % mod + dfs(i, nums, (mask | (1 << i))) % mod
//             }
//         }
//         dp[pre][mask] = count % mod
//         return dp[pre][mask]
//     }
//     return dfs(0,nums,0) % mod
// }

// 状态压缩动态规划
func specialPerm(nums []int) int {
    n, mod := len(nums), 1_000_000_007
    dp := make([][]int, 1 << n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for i := 0; i < n; i++ {
        dp[1 << i][i] = 1
    }
    for state := 1; state < (1 << n); state++ {
        for i := 0; i < n; i++ {
            if state >> i & 1 == 0 { continue }
            for j := 0; j < n; j++ {
                if i == j || state >> j & 1 == 0 { continue }
                x, y := nums[i], nums[j]
                if x % y != 0 && y % x != 0 { continue }
                dp[state][i] = (dp[state][i] + dp[state ^ (1<<i)][j]) % mod
            }
        }
    }
    res := 0
    for i := 0; i < n; i++ {
        res = (res + dp[(1 << n) - 1][i]) % mod
    }
    return res
}

func specialPerm1(nums []int) int {
    res, n := 0, len(nums)
    d := make([]int, n)
    for i, x := range nums {
        for j, y := range nums {
            if j != i && (x%y == 0 || y%x == 0) {
                d[i] |= 1 << j
            }
        }
    }
    f := make([][]int, 1<<n)
    for i := range f {
        f[i] = make([]int, n)
    }
    for j := range f[0] {
        f[1<<j][j] = 1
    }
    for s, dr := range f {
        for _s := uint(s); _s > 0; _s &= _s - 1 {
            i := bits.TrailingZeros(_s)
            if dr[i] == 0 {
                continue
            }
            pre := nums[i]
            for cus, lb := (len(f)-1^s)&d[i], 0; cus > 0; cus ^= lb {
                lb = cus & -cus
                j := bits.TrailingZeros(uint(lb))
                cur := nums[j]
                if pre%cur == 0 || cur%pre == 0 {
                    f[s|lb][j] += dr[i]
                }
            }
        }
    }
    for _, v := range f[len(f)-1] {
        res += v
    }
    return res % 1_000_000_007
}

func main() {
    // Example 1:
    // Input: nums = [2,3,6]
    // Output: 2
    // Explanation: [3,6,2] and [2,6,3] are the two special permutations of nums.
    fmt.Println(specialPerm([]int{2,3,6})) // 2
    // Example 2:
    // Input: nums = [1,4,3]
    // Output: 2
    // Explanation: [3,1,4] and [4,1,3] are the two special permutations of nums.
    fmt.Println(specialPerm([]int{1,4,3})) // 2

    fmt.Println(specialPerm1([]int{2,3,6})) // 2
    fmt.Println(specialPerm1([]int{1,4,3})) // 2
}