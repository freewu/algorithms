package main

// 1434. Number of Ways to Wear Different Hats to Each Other
// There are n people and 40 types of hats labeled from 1 to 40.
// Given a 2D integer array hats, where hats[i] is a list of all hats preferred by the ith person.
// Return the number of ways that the n people wear different hats to each other.
// Since the answer may be too large, return it modulo 109 + 7.

// Example 1:
// Input: hats = [[3,4],[4,5],[5]]
// Output: 1
// Explanation: There is only one way to choose hats given the conditions. 
// First person choose hat 3, Second person choose hat 4 and last one hat 5.

// Example 2:
// Input: hats = [[3,5,1],[3,5]]
// Output: 4
// Explanation: There are 4 ways to choose hats:
// (3,5), (5,3), (1,3) and (1,5)

// Example 3:
// Input: hats = [[1,2,3,4],[1,2,3,4],[1,2,3,4],[1,2,3,4]]
// Output: 24
// Explanation: Each person can choose hats labeled from 1 to 4.
// Number of Permutations of (1,2,3,4) = 24.

// Constraints:
//     n == hats.length
//     1 <= n <= 10
//     1 <= hats[i].length <= 40
//     1 <= hats[i][j] <= 40
//     hats[i] contains a list of unique integers.

import "fmt"

func numberWays(hats [][]int) int {
    pb, nn := make([]int, len(hats)), 0
    for i := range pb {
        pb[i] = 1 << i
        nn |= pb[i]
    }
    hm := make([]map[int]struct{}, len(pb))
    for i := range hm {
        hm[i] = make(map[int]struct{})
        for j := range hats[i] {
            hm[i][hats[i][j]] = struct{}{}
        }
    }
    memo := func(fn func(int, int) int) func(int, int) int {
        m := make(map[string]int)
        d := func(i int, f int) int {
            k := fmt.Sprintf("%d-%d", i, f)
            if v, ok := m[k]; ok {
                return v
            }

            v := fn(i, f)
            m[k] = v
            return v
        }
        return d
    }
    var dp func(int, int) int
    dp = func(pm int, h int) int {
        if pm == nn {
            return 1
        }
        if h > 40 {
            return 0
        }
        a := 0
        for i := range pb {
            _, ok := hm[i][h]
            if pm&pb[i] == 0 && ok {
                a += dp(pm|pb[i], h+1)
            }
        }
        return a + dp(pm, h+1)
    }
    dp = memo(dp)
    return dp(0,1) % 1_000_000_007
}

func numberWays1(hats [][]int) int {
    n, hatToPeople := len(hats), make([][]int, 40)
    mask := 1 << n
    for people, hatList := range hats {
        for _, hat := range hatList {
            hatToPeople[hat-1] = append(hatToPeople[hat-1], people)
        }
    }
    dp := make([]int, mask) // 定义 dp[i][j] 为 [0, i] 号帽子已经分配给的人的集合为 j 的方案数
    dp[0] = 1 // base case，不分配任何帽子的方案数
    for _, id := range hatToPeople[0] { // base case，0 号帽子的分配方案数
        dp[1 << id] = 1
    }
    for i := 1; i < 40; i++ {
        // dp[i][j] 其实要延续 dp[i-1][j]，表示当前帽子不分配给任何人，
        // 我们这里采用滚动数组的写法就隐式帮我们做到这一点。
        // 另外我们计算 j 时需要用到上一轮的 <j 的状态，因此我们 j 要
        // 倒着遍历，防止 < j 的状态已经更新为本轮。
        for j := mask - 1; j > 0; j-- {
            for _, id := range hatToPeople[i] {
                if j >> id & 1 > 0 {
                    dp[j] = (dp[j] + dp[j-(1<<id)]) % 1_000_000_007
                }
            }
        }
    }
    return dp[mask-1]
}

func main() {
    // Example 1:
    // Input: hats = [[3,4],[4,5],[5]]
    // Output: 1
    // Explanation: There is only one way to choose hats given the conditions. 
    // First person choose hat 3, Second person choose hat 4 and last one hat 5.
    fmt.Println(numberWays([][]int{{3,4},{4, 5},{5}})) // 1
    // Example 2:
    // Input: hats = [[3,5,1],[3,5]]
    // Output: 4
    // Explanation: There are 4 ways to choose hats:
    // (3,5), (5,3), (1,3) and (1,5)
    fmt.Println(numberWays([][]int{{3,5,1},{3, 5}})) // 4
    // Example 3:
    // Input: hats = [[1,2,3,4],[1,2,3,4],[1,2,3,4],[1,2,3,4]]
    // Output: 24
    // Explanation: Each person can choose hats labeled from 1 to 4.
    // Number of Permutations of (1,2,3,4) = 24.
    fmt.Println(numberWays([][]int{{1,2,3,4},{1,2,3,4},{1,2,3,4},{1,2,3,4}})) // 24

    fmt.Println(numberWays1([][]int{{3,4},{4, 5},{5}})) // 1
    fmt.Println(numberWays1([][]int{{3,5,1},{3, 5}})) // 4
    fmt.Println(numberWays1([][]int{{1,2,3,4},{1,2,3,4},{1,2,3,4},{1,2,3,4}})) // 24
}