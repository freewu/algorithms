package main

// 1478. Allocate Mailboxes
// Given the array houses where houses[i] is the location of the ith house along a street and an integer k, 
// allocate k mailboxes in the street.

// Return the minimum total distance between each house and its nearest mailbox.

// The test cases are generated so that the answer fits in a 32-bit integer.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/05/07/sample_11_1816.png" />
// Input: houses = [1,4,8,10,20], k = 3
// Output: 5
// Explanation: Allocate mailboxes in position 3, 9 and 20.
// Minimum total distance from each houses to nearest mailboxes is |3-1| + |4-3| + |9-8| + |10-9| + |20-20| = 5 

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/05/07/sample_2_1816.png" />
// Input: houses = [2,3,5,12,18], k = 2
// Output: 9
// Explanation: Allocate mailboxes in position 3 and 14.
// Minimum total distance from each houses to nearest mailboxes is |2-3| + |3-3| + |5-3| + |12-14| + |18-14| = 9.

// Constraints:
//     1 <= k <= houses.length <= 100
//     1 <= houses[i] <= 10^4
//     All the integers of houses are unique.

import "fmt"
import "sort"

func minDistance(houses []int, k int) int {
    sort.Ints(houses)
    n := len(houses)
    dp := make([]int, n)
    for i := 1; i < n; i++ {
        dp[i] = dp[i-1] + houses[i] - houses[i/2]
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < k - 1; i++ {
        for j := n-1; j >= 0; j-- {
            sum := 0
            for m := j; m >= 0; m-- {
                sum += houses[(m+j+1) >> 1] - houses[m]
                if m == 0 {
                    dp[j] = min(dp[j], sum)
                } else {
                    dp[j] = min(dp[j], dp[m-1] + sum)
                }
            }
        }
    }
    return dp[n-1]
}

func minDistance1(houses []int, k int) int {
    n, inf := len(houses), 1 << 31
    sort.Ints(houses)
    // 预处理所有子串和中位数的距离
    dis := make([][]int, n)
    for i := n-1; i >= 0; i-- {
        dis[i] = make([]int, n)
        for j := i+1; j < n; j++ {
            dis[i][j] = dis[i+1][j-1] - houses[i] + houses[j]
        }
    }
    // 对于每个社区, 将邮箱放在中位数位置可保证距离最小
    dp := make([][]int, n + 1)
    for i := range dp {
        dp[i] = make([]int, k + 1)
        for j := range dp[i] {
            dp[i][j] = inf
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    dp[0][0] = 0
    for i := 0; i < n; i++ {
        for j := 0; j < min(k, i + 1); j++ {
            for c := i; c >= 0; c-- {
                dp[i+1][j+1] = min(dp[i+1][j+1], dp[c][j] + dis[c][i])
            }
        }
    }
    return dp[n][k]
}

func minDistance2(houses []int, k int) int {
    n, inf := len(houses), 1 << 31
    sort.Ints(houses)
    cost := make([][]int, n) // Calculate the minimum cost of placing a single mailbox for houses[i:j+1]
    for i := range cost {
        cost[i] = make([]int, n)
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            mid := (i + j) / 2
            for m := i; m <= j; m++ {
                cost[i][j] += abs(houses[m] - houses[mid])
            }
        }
    }
    dp := make([][]int, k+1) // Create DP table
    for i := range dp {
        dp[i] = make([]int, n+1)
        for j := range dp[i] {
            dp[i][j] = inf
        }
    }
    dp[0][0] = 0 // Base case: dp[0][0] is 0 (no cost for no houses and no mailboxes)
    for i := 1; i <= k; i++ { // Fill the DP table
        for j := 1; j <= n; j++ {
            for p := 0; p < j; p++ {
                dp[i][j] = min(dp[i][j], dp[i-1][p] + cost[p][j-1])
            }
        }
    }
    return dp[k][n] // The answer is the minimum cost to place k mailboxes for all n houses
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/05/07/sample_11_1816.png" />
    // Input: houses = [1,4,8,10,20], k = 3
    // Output: 5
    // Explanation: Allocate mailboxes in position 3, 9 and 20.
    // Minimum total distance from each houses to nearest mailboxes is |3-1| + |4-3| + |9-8| + |10-9| + |20-20| = 5 
    fmt.Println(minDistance([]int{1,4,8,10,20}, 3)) // 5
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/05/07/sample_2_1816.png" />
    // Input: houses = [2,3,5,12,18], k = 2
    // Output: 9
    // Explanation: Allocate mailboxes in position 3 and 14.
    // Minimum total distance from each houses to nearest mailboxes is |2-3| + |3-3| + |5-3| + |12-14| + |18-14| = 9.
    fmt.Println(minDistance([]int{2,3,5,12,18}, 2)) // 9

    fmt.Println(minDistance1([]int{1,4,8,10,20}, 3)) // 5
    fmt.Println(minDistance1([]int{2,3,5,12,18}, 2)) // 9

    fmt.Println(minDistance2([]int{1,4,8,10,20}, 3)) // 5
    fmt.Println(minDistance2([]int{2,3,5,12,18}, 2)) // 9
}