package main

// 3621. Number of Integers With Popcount-Depth Equal to K I
// You are given two integers n and k.

// For any positive integer x, define the following sequence:
//     1. p0 = x
//     2. pi+1 = popcount(pi) for all i >= 0, where popcount(y) is the number of set bits (1's) in the binary representation of y.

// This sequence will eventually reach the value 1.

// The popcount-depth of x is defined as the smallest integer d >= 0 such that pd = 1.

// For example, if x = 7 (binary representation "111"). 
// Then, the sequence is: 7 → 3 → 2 → 1, so the popcount-depth of 7 is 3.

// Your task is to determine the number of integers in the range [1, n] whose popcount-depth is exactly equal to k.

// Return the number of such integers.

// Example 1:
// Input: n = 4, k = 1
// Output: 2
// Explanation:
// The following integers in the range [1, 4] have popcount-depth exactly equal to 1:
// x	Binary	Sequence
// 2	"10"	2 → 1
// 4	"100"	4 → 1
// Thus, the answer is 2.

// Example 2:
// Input: n = 7, k = 2
// Output: 3
// Explanation:
// The following integers in the range [1, 7] have popcount-depth exactly equal to 2:
// x	Binary	Sequence
// 3	"11"	3 → 2 → 1
// 5	"101"	5 → 2 → 1
// 6	"110"	6 → 2 → 1
// Thus, the answer is 3.

// Constraints:
//     1 <= n <= 10^15
//     0 <= k <= 5

import "fmt"
import "strconv"
import "math/bits"

func popcountDepth(n int64, k int) int64 {
    if k == 0 { return 1 }
    // 注：也可以不转成字符串，下面 dfs 用位运算取出 n 的第 i 位 
    // 但转成字符串的通用性更好
    s := strconv.FormatInt(n, 2)
    res, m := int64(0), len(s)
    if k == 1 { return int64(m - 1) }
    memo := make([][]int64, m)
    for i := range memo {
        memo[i] = make([]int64, m+1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    var dfs func(i, left1 int, isLimit bool) int64
    dfs = func(i, left1 int, isLimit bool) int64 {
        val, up := int64(0), 1
        if i == m {
            if left1 == 0 { return 1 }
            return val
        }
        if !isLimit {
            p := &memo[i][left1]
            if *p >= 0 {
                return *p
            }
            defer func() { *p = val }()
        }
        if isLimit {
            up = int(s[i] - '0')
        }
        for d := 0; d <= min(up, left1); d++ {
            val += dfs(i + 1, left1 - d, isLimit && d == up)
        }
        return val
    }
    f := make([]int, m + 1)
    for i := 1; i <= m; i++ {
        f[i] = f[bits.OnesCount(uint(i))] + 1
        if f[i] == k {
            // 计算有多少个二进制数恰好有 i 个 1
            res += dfs(0, i, true)
        }
    }
    return res
}

func popcountDepth1(n int64, k int) int64 {
    if k == 0 { return 1 }
    f, arr := make([]int, 61), []int{}
    f[1] = 0
    for i := 2; i <= 60; i++ {
        j := bits.OnesCount(uint(i))
        f[i] = f[j] + 1
        if f[i] == k-1 {
            arr = append(arr, i)
        }
    }
    if k == 1 {
        arr = append(arr, 1)
    }
    h := bits.Len(uint(n))
    grid := make([][]int, h)
    for i := 0; i < h; i++ {
        grid[i] = make([]int, i + 1)
        grid[i][0], grid[i][i] = 1, 1
        for j := 1; j < i; j++ {
            grid[i][j] = grid[i-1][j-1] + grid[i-1][j]
        }
    }
    res, count := 0, 0
    for i := h - 1; i >= 0; i-- {
        x := (n >> i) & 1
        if x == 1 {
            // 如果这里放置0, 然后在后i位，放置 w - count 个 1
            for _, w := range arr {
                // count + y = w
                if count <= w && w - count <= i {
                    res += grid[i][w - count]
                }
                // 将1舍弃掉
                if w == 1 && count == 0 {
                    res--
                }
            }
            count++
        }
    }
    if f[count] == k - 1 {
        res++
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: n = 4, k = 1
    // Output: 2
    // Explanation:
    // The following integers in the range [1, 4] have popcount-depth exactly equal to 1:
    // x	Binary	Sequence
    // 2	"10"	2 → 1
    // 4	"100"	4 → 1
    // Thus, the answer is 2.
    fmt.Println(popcountDepth(4,1)) // 2
    // Example 2:
    // Input: n = 7, k = 2
    // Output: 3
    // Explanation:
    // The following integers in the range [1, 7] have popcount-depth exactly equal to 2:
    // x	Binary	Sequence
    // 3	"11"	3 → 2 → 1
    // 5	"101"	5 → 2 → 1
    // 6	"110"	6 → 2 → 1
    // Thus, the answer is 3.
    fmt.Println(popcountDepth(7, 2)) // 3

    fmt.Println(popcountDepth(1,0)) // 1
    fmt.Println(popcountDepth(1,5)) // 0
    fmt.Println(popcountDepth(1_000_000_000_000_000, 0)) // 1
    fmt.Println(popcountDepth(1_000_000_000_000_000, 5)) // 0

    fmt.Println(popcountDepth1(4,1)) // 2
    fmt.Println(popcountDepth1(7, 2)) // 3
    fmt.Println(popcountDepth1(1,0)) // 1
    fmt.Println(popcountDepth1(1,5)) // 0
    fmt.Println(popcountDepth1(1_000_000_000_000_000, 0)) // 1
    fmt.Println(popcountDepth1(1_000_000_000_000_000, 5)) // 0
}