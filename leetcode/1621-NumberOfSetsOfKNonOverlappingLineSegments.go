package main

// 1621. Number of Sets of K Non-Overlapping Line Segments
// Given n points on a 1-D plane, where the ith point (from 0 to n-1) is at x = i, 
// find the number of ways we can draw exactly k non-overlapping line segments such that each segment covers two or more points. 
// The endpoints of each segment must have integral coordinates. 
// The k line segments do not have to cover all n points, and they are allowed to share endpoints.

// Return the number of ways we can draw k non-overlapping line segments. 
// Since this number can be huge, return it modulo 10^9 + 7.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/07/ex1.png" />
// Input: n = 4, k = 2
// Output: 5
// Explanation: The two line segments are shown in red and blue.
// The image above shows the 5 different ways {(0,2),(2,3)}, {(0,1),(1,3)}, {(0,1),(2,3)}, {(1,2),(2,3)}, {(0,1),(1,2)}.

// Example 2:
// Input: n = 3, k = 1
// Output: 3
// Explanation: The 3 ways are {(0,1)}, {(0,2)}, {(1,2)}.

// Example 3:
// Input: n = 30, k = 7
// Output: 796297179
// Explanation: 
// The total number of possible ways to draw 7 line segments is 3796297200. 
// Taking this number modulo 10^9 + 7 gives us 796297179.

// Constraints:
//     2 <= n <= 1000
//     1 <= k <= n-1

import "fmt"

// dfs with memo
func numberOfSets(n int, k int) int {
    memo := make([][]int, n + 1)
    for i := 0; i < len(memo); i++ {
        memo[i] = make([]int, n)
    }
    var dfs func(n int, k int, memo [][]int) int
    dfs = func(n int, k int, memo [][]int) int {
        if k >= n { return 0 }
        if k == 1 { return n * (n - 1) / 2 }
        if memo[n][k] != 0 { return memo[n][k] }
        res := 0
        for i := 1; i < n; i++ {
            res += i  * dfs(n - i, k - 1, memo)
        }
        memo[n][k] = (res % 1_000_000_007)
        return memo[n][k]
    }
    return dfs(n, k, memo)
}

func numberOfSets1(n int, k int) int {
    n = n - 1 + k
    res, mod := n, 1_000_000_007
    power := func (x, y int) int {
        res := 1
        for ; y > 0; y >>= 1 {
            if y&1 > 0 {
                res = res * x % mod
            }
            x = x * x % mod
        }
        return res
    }
    // 2*n-2+k  n-1-k
    div := 1
    for i := 2; i <= k*2; i++ {
        res = res * (n - i + 1) % mod
        div = div * i % mod
    }
    return res * power(div, mod-2) % mod
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/07/ex1.png" />
    // Input: n = 4, k = 2
    // Output: 5
    // Explanation: The two line segments are shown in red and blue.
    // The image above shows the 5 different ways {(0,2),(2,3)}, {(0,1),(1,3)}, {(0,1),(2,3)}, {(1,2),(2,3)}, {(0,1),(1,2)}.
    fmt.Println(numberOfSets(4, 2)) // 5
    // Example 2:
    // Input: n = 3, k = 1
    // Output: 3
    // Explanation: The 3 ways are {(0,1)}, {(0,2)}, {(1,2)}.
    fmt.Println(numberOfSets(3, 1)) // 3
    // Example 3:
    // Input: n = 30, k = 7
    // Output: 796297179
    // Explanation: 
    // The total number of possible ways to draw 7 line segments is 3796297200. 
    // Taking this number modulo 10^9 + 7 gives us 796297179.
    fmt.Println(numberOfSets(30, 7)) // 796297179
    
    fmt.Println(numberOfSets(999, 998)) // 1
    fmt.Println(numberOfSets(999, 1)) // 498501

    fmt.Println(numberOfSets1(4, 2)) // 5
    fmt.Println(numberOfSets1(3, 1)) // 3
    fmt.Println(numberOfSets1(30, 7)) // 796297179
    fmt.Println(numberOfSets1(999, 998)) // 1
    fmt.Println(numberOfSets1(999, 1)) // 498501
}