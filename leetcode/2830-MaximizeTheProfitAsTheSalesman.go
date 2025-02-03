package main

// 2830. Maximize the Profit as the Salesman
// You are given an integer n representing the number of houses on a number line, numbered from 0 to n - 1.

// Additionally, you are given a 2D integer array offers where offers[i] = [starti, endi, goldi], 
// indicating that ith buyer wants to buy all the houses from starti to endi for goldi amount of gold.

// As a salesman, your goal is to maximize your earnings by strategically selecting and selling houses to buyers.

// Return the maximum amount of gold you can earn.

// Note that different buyers can't buy the same house, and some houses may remain unsold.

// Example 1:
// Input: n = 5, offers = [[0,0,1],[0,2,2],[1,3,2]]
// Output: 3
// Explanation: There are 5 houses numbered from 0 to 4 and there are 3 purchase offers.
// We sell houses in the range [0,0] to 1st buyer for 1 gold and houses in the range [1,3] to 3rd buyer for 2 golds.
// It can be proven that 3 is the maximum amount of gold we can achieve.

// Example 2:
// Input: n = 5, offers = [[0,0,1],[0,2,10],[1,3,2]]
// Output: 10
// Explanation: There are 5 houses numbered from 0 to 4 and there are 3 purchase offers.
// We sell houses in the range [0,2] to 2nd buyer for 10 golds.
// It can be proven that 10 is the maximum amount of gold we can achieve.

// Constraints:
//     1 <= n <= 10^5
//     1 <= offers.length <= 10^5
//     offers[i].length == 3
//     0 <= starti <= endi <= n - 1
//     1 <= goldi <= 10^3

import "fmt"
import "sort"

func maximizeTheProfit(n int, offers [][]int) int {
    sort.Slice(offers, func(i, j int) bool {
        return offers[i][0] < offers[j][0]
    })
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp, j := make([]int, n + 1), 0
    for i := 1; i <= n; i++ {
        dp[i] = max(dp[i], dp[i-1])
        for ; j < len(offers) && offers[j][0] == i - 1; j++ {
            start, end, gold := offers[j][0], offers[j][1], offers[j][2]
            dp[end + 1] = max(dp[end + 1], dp[start] + gold)
        }
    }
    return dp[n]
}

func maximizeTheProfit1(n int, offers [][]int) int {
    type Pair struct{ start, gold int }
    groups := make([][]Pair, n)
    for _, offer := range offers {
        start, end, gold := offer[0], offer[1], offer[2]
        groups[end] = append(groups[end], Pair{ start, gold })
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp := make([]int, n + 1)
    for end, group := range groups {
        dp[end + 1] = dp[end]
        for _, pair := range group {
            dp[end + 1] = max(dp[end + 1], dp[pair.start] + pair.gold)
        }
    }
    return dp[n]
}

func main() {
    // Example 1:
    // Input: n = 5, offers = [[0,0,1],[0,2,2],[1,3,2]]
    // Output: 3
    // Explanation: There are 5 houses numbered from 0 to 4 and there are 3 purchase offers.
    // We sell houses in the range [0,0] to 1st buyer for 1 gold and houses in the range [1,3] to 3rd buyer for 2 golds.
    // It can be proven that 3 is the maximum amount of gold we can achieve.
    fmt.Println(maximizeTheProfit(5, [][]int{{0,0,1},{0,2,2},{1,3,2}})) // 3
    // Example 2:
    // Input: n = 5, offers = [[0,0,1],[0,2,10],[1,3,2]]
    // Output: 10
    // Explanation: There are 5 houses numbered from 0 to 4 and there are 3 purchase offers.
    // We sell houses in the range [0,2] to 2nd buyer for 10 golds.
    // It can be proven that 10 is the maximum amount of gold we can achieve.
    fmt.Println(maximizeTheProfit(5, [][]int{{0,0,1},{0,2,10},{1,3,2}})) // 10

    fmt.Println(maximizeTheProfit1(5, [][]int{{0,0,1},{0,2,2},{1,3,2}})) // 3
    fmt.Println(maximizeTheProfit1(5, [][]int{{0,0,1},{0,2,10},{1,3,2}})) // 10
}