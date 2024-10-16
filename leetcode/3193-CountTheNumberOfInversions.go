package main

// 3193. Count the Number of Inversions
// You are given an integer n and a 2D array requirements, 
// where requirements[i] = [endi, cnti] represents the end index and the inversion count of each requirement.

// A pair of indices (i, j) from an integer array nums is called an inversion if:
//     i < j and nums[i] > nums[j]

// Return the number of permutations perm of [0, 1, 2, ..., n - 1] such that for all requirements[i], perm[0..endi] has exactly cnti inversions.

// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: n = 3, requirements = [[2,2],[0,0]]
// Output: 2
// Explanation:
// The two permutations are:
//     [2, 0, 1]
//         Prefix [2, 0, 1] has inversions (0, 1) and (0, 2).
//         Prefix [2] has 0 inversions.
//     [1, 2, 0]
//         Prefix [1, 2, 0] has inversions (0, 2) and (1, 2).
//         Prefix [1] has 0 inversions.

// Example 2:
// Input: n = 3, requirements = [[2,2],[1,1],[0,0]]
// Output: 1
// Explanation:
// The only satisfying permutation is [2, 0, 1]:
//     Prefix [2, 0, 1] has inversions (0, 1) and (0, 2).
//     Prefix [2, 0] has an inversion (0, 1).
//     Prefix [2] has 0 inversions.

// Example 3:
// Input: n = 2, requirements = [[0,0],[1,0]]
// Output: 1
// Explanation:
// The only satisfying permutation is [0, 1]:
//     Prefix [0] has 0 inversions.
//     Prefix [0, 1] has an inversion (0, 1).

// Constraints:
//     2 <= n <= 300
//     1 <= requirements.length <= n
//     requirements[i] = [endi, cnti]
//     0 <= endi <= n - 1
//     0 <= cnti <= 400
//     The input is generated such that there is at least one i such that endi == n - 1.
//     The input is generated such that all endi are unique.

import "fmt"
import "slices"

// 解答错误 719 / 720
// func numberOfPermutations(n int, requirements [][]int) int {
//     mod := 1_000_000_007
//     mp := make(map[int]int)
//     for _, v := range requirements {
//         mp[v[0] + 1] = v[1]
//     }
//     dp := make([][]int, n + 1)
//     for i := range dp {
//         dp[i] = make([]int, 401)
//     }
//     dp[1][0] = 1
//     for i := 2; i <= n; i++ {
//         for j := 0; j < 401; j++ {
//             for k := 0; k < i; k++ {
//                 count := k + j
//                 if count <= 400 {
//                     dp[i][count] += dp[i - 1][j]
//                     dp[i][count] %= mod
//                 }
//             }
//         }
//         if v, ok := mp[i]; ok {
//             for j := 0; j < 401; j++ {
//                 if j != v {
//                     dp[i][j] = 0
//                 }
//             }
//         }
//     }
//     return dp[n][mp[n]]
// }

func numberOfPermutations(n int, requirements [][]int) int {
    req := make([]int, n+1)
    for i := 0; i < n + 1; i++ { // init req -1
        req[i] = -1
    }
    target, mod := 400, 1_000_000_007
    for _, r := range requirements {
        req[r[0] + 1] = r[1]
    }
    dp := make([][]int, n + 1)
    for i := 0; i < n + 1; i++ {
        dp[i] = make([]int, target + 1)
    }
    dp[1][0] = 1
    for i := 0; i < n; i++ {
        idx := i + 1
        for j := 0; j < target + 1; j++ {
            for k := 0; k < idx && j - k >= 0; k++ {
                dp[idx][j] += dp[i][j-k] % mod 
            }
        }
        r := req[idx] 
        if r != -1 {
            for j := 0; j < target + 1; j++ {
                if r != j {
                    dp[idx][j] = 0
                }
            }
        }
    }
    res := 0
    for i := 0; i < target + 1; i++ {
        res += dp[n][i] % mod
    }
    return res
}

func numberOfPermutations1(n int, requirements [][]int) int {
    mod := 1_000_000_007
    req := make([]int, n)
    for i := 1; i < n; i++ {
        req[i] = -1
    }
    for _, v := range requirements {
        req[v[0]] = v[1]
    }
    if req[0] > 0 {
        return 0
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    mx := slices.Max(req)
    dp := make([]int, mx + 1)
    dp[0] = 1
    for i := 1; i < n; i++ {
        t := mx
        if req[i] >= 0 { t = req[i] }
        if r := req[i-1]; r >= 0 {
            clear(dp[:r])
            for j := r + 1; j <= min(i + r, t); j++ {
                dp[j] = dp[r]
            }
            clear(dp[min(i + r, t) + 1:])
        } else {
            for j := 1; j <= t; j++ {
                dp[j] = (dp[j] + dp[j-1]) % mod
            }
            for j := t; j > i; j-- {
                dp[j] = (dp[j] - dp[j-i-1] + mod) % mod
            }
        }
    }
    return dp[req[n-1]]
}

func main() {
    // Example 1:
    // Input: n = 3, requirements = [[2,2],[0,0]]
    // Output: 2
    // Explanation:
    // The two permutations are:
    //     [2, 0, 1]
    //         Prefix [2, 0, 1] has inversions (0, 1) and (0, 2).
    //         Prefix [2] has 0 inversions.
    //     [1, 2, 0]
    //         Prefix [1, 2, 0] has inversions (0, 2) and (1, 2).
    //         Prefix [1] has 0 inversions.
    fmt.Println(numberOfPermutations(3, [][]int{{2,2},{0,0}})) // 2
    // Example 2:
    // Input: n = 3, requirements = [[2,2],[1,1],[0,0]]
    // Output: 1
    // Explanation:
    // The only satisfying permutation is [2, 0, 1]:
    //     Prefix [2, 0, 1] has inversions (0, 1) and (0, 2).
    //     Prefix [2, 0] has an inversion (0, 1).
    //     Prefix [2] has 0 inversions.
    fmt.Println(numberOfPermutations(3, [][]int{{2,2},{1,1},{0,0}})) // 1
    // Example 3:
    // Input: n = 2, requirements = [[0,0],[1,0]]
    // Output: 1
    // Explanation:
    // The only satisfying permutation is [0, 1]:
    //     Prefix [0] has 0 inversions.
    //     Prefix [0, 1] has an inversion (0, 1).
    fmt.Println(numberOfPermutations(2, [][]int{{0,0},{1,0}})) // 1

    fmt.Println(numberOfPermutations(3, [][]int{{2,2},{0,1}})) // 0

    fmt.Println(numberOfPermutations1(3, [][]int{{2,2},{0,0}})) // 2
    fmt.Println(numberOfPermutations1(3, [][]int{{2,2},{1,1},{0,0}})) // 1
    fmt.Println(numberOfPermutations1(2, [][]int{{0,0},{1,0}})) // 1
    fmt.Println(numberOfPermutations1(3, [][]int{{2,2},{0,1}})) // 0
}