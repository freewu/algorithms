package main

// 3336. Find the Number of Subsequences With Equal GCD
// You are given an integer array nums.

// Your task is to find the number of pairs of non-empty subsequences (seq1, seq2) of nums that satisfy the following conditions:
//     1. The subsequences seq1 and seq2 are disjoint, meaning no index of nums is common between them.
//     2. The GCD of the elements of seq1 is equal to the GCD of the elements of seq2.

// Return the total number of such pairs.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: 10
// Explanation:
// The subsequence pairs which have the GCD of their elements equal to 1 are:
// ([1, 2, 3, 4], [1, 2, 3, 4])
// ([1, 2, 3, 4], [1, 2, 3, 4])
// ([1, 2, 3, 4], [1, 2, 3, 4])
// ([1, 2, 3, 4], [1, 2, 3, 4])
// ([1, 2, 3, 4], [1, 2, 3, 4])
// ([1, 2, 3, 4], [1, 2, 3, 4])
// ([1, 2, 3, 4], [1, 2, 3, 4])
// ([1, 2, 3, 4], [1, 2, 3, 4])
// ([1, 2, 3, 4], [1, 2, 3, 4])
// ([1, 2, 3, 4], [1, 2, 3, 4])

// Example 2:
// Input: nums = [10,20,30]
// Output: 2
// Explanation:
// The subsequence pairs which have the GCD of their elements equal to 10 are:
// ([10, 20, 30], [10, 20, 30])
// ([10, 20, 30], [10, 20, 30])

// Example 3:
// Input: nums = [1,1,1,1]
// Output: 50

// Constraints:
//     1 <= nums.length <= 200
//     1 <= nums[i] <= 200

import "fmt"
import "slices"

func subsequencePairCount(nums []int) int {
    n, m, mod := len(nums), slices.Max(nums), 1_000_000_007
    memo := make([][][]int, n)
    for i := range memo {
        memo[i] = make([][]int, m+1)
        for j := range memo[i] {
            memo[i][j] = make([]int, m+1)
            for k := range memo[i][j] {
                memo[i][j][k] = -1 // -1 表示没有计算过
            }
        }
    }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    var dfs func(int, int, int) int
    dfs = func(i, j, k int) int {
        if i < 0 {
            if j == k { return 1}
            return 0
        }
        if memo[i][j][k] < 0 {
            memo[i][j][k] = (dfs(i-1, j, k) + dfs(i-1, gcd(j, nums[i]), k) + dfs(i-1, j, gcd(k, nums[i]))) % mod
        }
        return memo[i][j][k]
    }
    // 减去两个子序列都是空的情况
    return (dfs(n-1, 0, 0) - 1 + mod) % mod // +mod 防止减一后变成负数
}

func subsequencePairCount1(nums []int) int {
    n, m, mod := len(nums), slices.Max(nums), 1_000_000_007
    dp := make([][][]int, n + 1)
    for i := range dp {
        dp[i] = make([][]int, m+1)
        for j := range dp[i] {
            dp[i][j] = make([]int, m+1)
        }
    }
    for j := 1; j <= m; j++ {
        dp[0][j][j] = 1
    }
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    for i, x := range nums {
        for j := 0; j <= m; j++ {
            for k := 0; k <= m; k++ {
                dp[i+1][j][k] = (dp[i][j][k] + dp[i][gcd(j, x)][k] + dp[i][j][gcd(k, x)]) % mod
            }
        }
    }
    return dp[n][0][0]
}

const mod = 1_000_000_007
const mx = 201

var lcms [mx][mx]int
var pow2, pow3, mu [mx]int

func gcd(a, b int) int {
    for a != 0 { a, b = b % a, a }
    return b
}

func lcm(a, b int) int {
    return a / gcd(a, b) * b
}

func init() {
    for i := 1; i < mx; i++ {
        for j := 1; j < mx; j++ {
            lcms[i][j] = lcm(i, j)
        }
    }
    pow2[0], pow3[0] = 1, 1
    for i := 1; i < mx; i++ {
        pow2[i] = pow2[i-1] * 2 % mod
        pow3[i] = pow3[i-1] * 3 % mod
    }
    mu[1] = 1
    for i := 1; i < mx; i++ {
        for j := i * 2; j < mx; j += i {
            mu[j] -= mu[i]
        }
    }
}

func subsequencePairCount2(nums []int) int {
    m := slices.Max(nums)
    count := make([]int, m + 1) // count[i] 表示 nums 中的 i 的倍数的个数
    for _, v := range nums {
        count[v]++
    }
    for i := 1; i <= m; i++ {
        for j := i * 2; j <= m; j += i {
            count[i] += count[j] // 统计 i 的倍数的个数
        }
    }
    dp := make([][]int, m + 1)
    for g1 := 1; g1 <= m; g1++ {
        dp[g1] = make([]int, m + 1)
        for g2 := 1; g2 <= m; g2++ {
            l, c := lcms[g1][g2], 0
            if l <= m {
                c = count[l]
            }
            c1, c2 := count[g1], count[g2]
            dp[g1][g2] = (pow3[c] * pow2[c1 + c2 - c * 2] - pow2[c1] - pow2[c2] + 1) % mod
        }
    }
    // 倍数容斥
    res := 0
    for i := 1; i <= m; i++ {
        for j := 1; j <= m / i; j++ {
            for k := 1; k <= m / i; k++ {
                res += mu[j] * mu[k] * dp[j * i][k * i]
            }
        }
    }
    return (res % mod + mod) % mod // 保证 res 非负
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4]
    // Output: 10
    // Explanation:
    // The subsequence pairs which have the GCD of their elements equal to 1 are:
    // ([1, 2, 3, 4], [1, 2, 3, 4])
    // ([1, 2, 3, 4], [1, 2, 3, 4])
    // ([1, 2, 3, 4], [1, 2, 3, 4])
    // ([1, 2, 3, 4], [1, 2, 3, 4])
    // ([1, 2, 3, 4], [1, 2, 3, 4])
    // ([1, 2, 3, 4], [1, 2, 3, 4])
    // ([1, 2, 3, 4], [1, 2, 3, 4])
    // ([1, 2, 3, 4], [1, 2, 3, 4])
    // ([1, 2, 3, 4], [1, 2, 3, 4])
    // ([1, 2, 3, 4], [1, 2, 3, 4])
    fmt.Println(subsequencePairCount([]int{1,2,3,4})) // 10
    // Example 2:
    // Input: nums = [10,20,30]
    // Output: 2
    // Explanation:
    // The subsequence pairs which have the GCD of their elements equal to 10 are:
    // ([10, 20, 30], [10, 20, 30])
    // ([10, 20, 30], [10, 20, 30])
    fmt.Println(subsequencePairCount([]int{10,20,30})) // 2
    // Example 3:
    // Input: nums = [1,1,1,1]
    // Output: 50
    fmt.Println(subsequencePairCount([]int{1,1,1,1})) // 50

    fmt.Println(subsequencePairCount([]int{1,2,3,4,5,6,7,8,9})) // 11888
    fmt.Println(subsequencePairCount([]int{9,8,7,6,5,4,3,2,1})) // 11888

    fmt.Println(subsequencePairCount1([]int{1,2,3,4})) // 10
    fmt.Println(subsequencePairCount1([]int{10,20,30})) // 2
    fmt.Println(subsequencePairCount1([]int{1,1,1,1})) // 50
    fmt.Println(subsequencePairCount1([]int{1,2,3,4,5,6,7,8,9})) // 11888
    fmt.Println(subsequencePairCount1([]int{9,8,7,6,5,4,3,2,1})) // 11888
}