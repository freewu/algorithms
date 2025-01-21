package main

// 2218. Maximum Value of K Coins From Piles
// There are n piles of coins on a table. Each pile consists of a positive number of coins of assorted denominations.
// In one move, you can choose any coin on top of any pile, remove it, and add it to your wallet.
// Given a list piles, where piles[i] is a list of integers denoting the composition of the ith pile from top to bottom, 
// and a positive integer k, return the maximum total value of coins you can have in your wallet if you choose exactly k coins optimally.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/11/09/e1.png" />
// [1]    [7] | 1     [7]  | [1]     7
// 100     8  | 100   [8]  | [100]   8
//  3      9  | 3      9   |   3     9 
//     8      |     15     |      101
// Input: piles = [[1,100,3],[7,8,9]], k = 2
// Output: 101
// Explanation:
// The above diagram shows the different ways we can choose k coins.
// The maximum total we can obtain is 101.

// Example 2:
// Input: piles = [[100],[100],[100],[100],[100],[100],[1,1,1,1,1,1,700]], k = 7
// Output: 706
// Explanation:
// The maximum total can be obtained if we choose all coins from the last pile.
 
// Constraints:
//     n == piles.length
//     1 <= n <= 1000
//     1 <= piles[i][j] <= 10^5
//     1 <= k <= sum(piles[i].length) <= 2000

import "fmt"

// dfs  77 / 123  超出时间限制
func maxValueOfCoins1(piles [][]int, k int) int {
    n := len(piles)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(int, int) int
    dfs = func(i int, used int) int {
        if i == n { return 0 }
        accum := 0
        res := dfs(i+1, used)
        for j:=0; j< min(len(piles[i]), used); j++ {
            accum += piles[i][j]
            res = max(res, accum + dfs(i+1, used - j - 1))
        }
        return res
    }
    return dfs(0, k)
}

// dp dfs + memo
func maxValueOfCoins(piles [][]int, k int) int {
    n := len(piles)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, k+1)
    }
    var helper func(i, k int, piles [][]int, dp [][]int) int
    helper = func (i, k int, piles [][]int, dp [][]int) int {
        if dp[i][k] > 0 {
            return dp[i][k]
        }
        if i == len(piles) || k == 0 {
            return 0
        }
        res, cur := helper(i+1, k, piles, dp), 0
        for j := 0; j < len(piles[i]) && j < k; j++ {
            cur += piles[i][j]
            res = max(res, helper(i+1, k-j-1, piles, dp)+cur)
        }
        dp[i][k] = res
        return res
    }
    return helper(0, k, piles, dp)
}

func maxValueOfCoins2(piles [][]int, k int) int {
    n:= len(piles)
    dp := make([]int, k + 1)
    totalPiles := 0
    for i := n - 1; i >= 0; i-- {
        totalPiles += len(piles[i]) // 
        for count := min(totalPiles, k); count >= 0; count-- {
            // dp[i][count] = dp[i+1][count]
            sum := 0
            for j := 0; i < len(piles[i]) && j < count; j++ {
                sum += piles[i][j]
                dp[count] = max(dp[count], dp[count-j-1] + sum)
            }
        }
    }
    return dp[k]
}

func maxValueOfCoins3(piles [][]int, k int) int {
    sum, dp := 0, make([]int, k + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, pile := range piles {
        n := len(pile)
        for i := 1; i < n; i++ {
            pile[i] += pile[i-1] // 提前计算 pile 的前缀和
        }
        sum = min(sum + n, k)
        for j := sum; j > 0; j-- { // 优化：j 从前 i 个栈的大小之和开始枚举
            for w, v := range pile[:min(n, j)] {
                dp[j] = max(dp[j], dp[j-w-1] + v) // w 从 0 开始，物品体积为 w + 1
            }
        }
    }
    return dp[k]
}

func main() {
    // The above diagram shows the different ways we can choose k coins.
    // The maximum total we can obtain is 101.
    fmt.Println(maxValueOfCoins([][]int{{1,100,3},{7,8,9}}, 2)) // 101
    // The maximum total can be obtained if we choose all coins from the last pile.
    fmt.Println(maxValueOfCoins([][]int{{100},{100},{100},{100},{100},{100},{1,1,1,1,1,1,700}}, 7)) // 706

    fmt.Println(maxValueOfCoins1([][]int{{1,100,3},{7,8,9}}, 2)) // 101
    fmt.Println(maxValueOfCoins1([][]int{{100},{100},{100},{100},{100},{100},{1,1,1,1,1,1,700}}, 7)) // 706

    // fmt.Println(maxValueOfCoins2([][]int{{1,100,3},{7,8,9}}, 2)) // 101
    // fmt.Println(maxValueOfCoins2([][]int{{100},{100},{100},{100},{100},{100},{1,1,1,1,1,1,700}}, 7)) // 706

    fmt.Println(maxValueOfCoins3([][]int{{1,100,3},{7,8,9}}, 2)) // 101
    fmt.Println(maxValueOfCoins3([][]int{{100},{100},{100},{100},{100},{100},{1,1,1,1,1,1,700}}, 7)) // 706
}