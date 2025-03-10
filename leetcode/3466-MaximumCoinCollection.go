package main

// 3466. Maximum Coin Collection 
// Mario drives on a two-lane freeway with coins every mile. 
// You are given two integer arrays, lane1 and lane2, where the value at the ith index represents the number of coins he gains or loses in the ith mile in that lane.
//     1. If Mario is in lane 1 at mile i and lane1[i] > 0, Mario gains lane1[i] coins.
//     2. If Mario is in lane 1 at mile i and lane1[i] < 0, Mario pays a toll and loses abs(lane1[i]) coins.
//     3. The same rules apply for lane2.

// Mario can enter the freeway anywhere and exit anytime after traveling at least one mile. 
// Mario always enters the freeway on lane 1 but can switch lanes at most 2 times.

// A lane switch is when Mario goes from lane 1 to lane 2 or vice versa.

// Return the maximum number of coins Mario can earn after performing at most 2 lane switches.

// Note: Mario can switch lanes immediately upon entering or just before exiting the freeway.

// Example 1:
// Input: lane1 = [1,-2,-10,3], lane2 = [-5,10,0,1]
// Output: 14
// Explanation:
// Mario drives the first mile on lane 1.
// He then changes to lane 2 and drives for two miles.
// He changes back to lane 1 for the last mile.
// Mario collects 1 + 10 + 0 + 3 = 14 coins.

// Example 2:
// Input: lane1 = [1,-1,-1,-1], lane2 = [0,3,4,-5]
// Output: 8
// Explanation:
// Mario starts at mile 0 in lane 1 and drives one mile.
// He then changes to lane 2 and drives for two more miles. He exits the freeway before mile 3.
// He collects 1 + 3 + 4 = 8 coins.

// Example 3:
// Input: lane1 = [-5,-4,-3], lane2 = [-1,2,3]
// Output: 5
// Explanation:
// Mario enters at mile 1 and immediately switches to lane 2. He stays here the entire way.
// He collects a total of 2 + 3 = 5 coins.

// Example 4:
// Input: lane1 = [-3,-3,-3], lane2 = [9,-2,4]
// Output: 11
// Explanation:
// Mario starts at the beginning of the freeway and immediately switches to lane 2. He stays here the whole way.
// He collects a total of 9 + (-2) + 4 = 11 coins.

// Example 5:
// Input: lane1 = [-10], lane2 = [-2]
// Output: -2
// Explanation:
// Since Mario must ride on the freeway for at least one mile, he rides just one mile in lane 2.
// He collects a total of -2 coins.

// Constraints:
//     1 <= lane1.length == lane2.length <= 10^5
//     -10^9 <= lane1[i], lane2[i] <= 10^9

import "fmt"

func maxCoins(lane1 []int, lane2 []int) int64 {
    n := len(lane1)
    dp := make([][3]int64, n + 1)
    res := int64(-1 << 63)// int64(math.MinInt64)
    max := func (x, y int64) int64 { if x > y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        for j := 0; j < 3; j++ {
            if dp[i - 1][j] > 0 {
                dp[i][j] = dp[i - 1][j]
            }
            if j > 0 {
                dp[i][j] = max(dp[i][j], dp[i - 1][j - 1])
            }
            if j == 1 {
                dp[i][j] += int64(lane2[i-1])
            } else {
                dp[i][j] += int64(lane1[i-1])
            }
            res = max(res, dp[i][j])
        }
    }
    return res
}

func maxCoins1(lane1 []int, lane2 []int) int64 {
    n := len(lane1)
    dp := make([][2][3]int64, n)
    for i := range dp { // fill -1
        for j := range dp[i] {
            for k := range dp[i][j] {
                dp[i][j][k] = -1
            }
        }
    }
    max := func (x, y int64) int64 { if x > y { return x; }; return y; }
    var dfs func(i, j, k int) int64
    dfs = func(i, j, k int) int64 {
        if i >= n { return 0 }
        if dp[i][j][k] != -1 { return dp[i][j][k] }
        x := int64(lane1[i])
        if j == 1 {
            x = int64(lane2[i])
        }
        mx := max(x, dfs(i+1, j, k)+x)
        if k > 0 {
            mx = max(mx, max(dfs(i+1, j^1, k-1) + x,  dfs(i, j^1, k-1)))
        }
        dp[i][j][k] = mx
        return mx
    }
    res := int64(-1e18)
    for i := range lane1 {
        res = max(res, dfs(i, 0, 2))
    }
    return res
}

func maxCoins2(lane1 []int, lane2 []int) int64 {
   var maxCoins, zeroSw, oneSw, twoSw int = -1e10, -1e10, -1e10, -1e10
    for i := 0; i < len(lane1); i++ {
        twoSw = max(twoSw + lane1[i], oneSw + lane1[i])
        oneSw = max(oneSw + lane2[i], zeroSw + lane2[i], lane2[i])
        zeroSw = max(lane1[i], zeroSw + lane1[i])
        maxCoins = max(maxCoins, zeroSw, oneSw, twoSw)
    }
    return int64(maxCoins)
}

func main() {
    // Example 1:
    // Input: lane1 = [1,-2,-10,3], lane2 = [-5,10,0,1]
    // Output: 14
    // Explanation:
    // Mario drives the first mile on lane 1.
    // He then changes to lane 2 and drives for two miles.
    // He changes back to lane 1 for the last mile.
    // Mario collects 1 + 10 + 0 + 3 = 14 coins.
    fmt.Println(maxCoins([]int{1,-2,-10,3}, []int{-5,10,0,1})) // 14
    // Example 2:
    // Input: lane1 = [1,-1,-1,-1], lane2 = [0,3,4,-5]
    // Output: 8
    // Explanation:
    // Mario starts at mile 0 in lane 1 and drives one mile.
    // He then changes to lane 2 and drives for two more miles. He exits the freeway before mile 3.
    // He collects 1 + 3 + 4 = 8 coins.
    fmt.Println(maxCoins([]int{1,-1,-1,-1}, []int{0,3,4,-5})) // 8
    // Example 3:
    // Input: lane1 = [-5,-4,-3], lane2 = [-1,2,3]
    // Output: 5
    // Explanation:
    // Mario enters at mile 1 and immediately switches to lane 2. He stays here the entire way.
    // He collects a total of 2 + 3 = 5 coins.
    fmt.Println(maxCoins([]int{-5,-4,-3}, []int{-1,2,3})) // 5
    // Example 4:
    // Input: lane1 = [-3,-3,-3], lane2 = [9,-2,4]
    // Output: 11
    // Explanation:
    // Mario starts at the beginning of the freeway and immediately switches to lane 2. He stays here the whole way.
    // He collects a total of 9 + (-2) + 4 = 11 coins.
    fmt.Println(maxCoins([]int{-3,-3,-3}, []int{9,-2,4})) // 11
    // Example 5:
    // Input: lane1 = [-10], lane2 = [-2]
    // Output: -2
    // Explanation:
    // Since Mario must ride on the freeway for at least one mile, he rides just one mile in lane 2.
    // He collects a total of -2 coins.
    fmt.Println(maxCoins([]int{-10}, []int{-2})) // -2

    fmt.Println(maxCoins([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 65

    fmt.Println(maxCoins1([]int{1,-2,-10,3}, []int{-5,10,0,1})) // 14
    fmt.Println(maxCoins1([]int{1,-1,-1,-1}, []int{0,3,4,-5})) // 8
    fmt.Println(maxCoins1([]int{-5,-4,-3}, []int{-1,2,3})) // 5
    fmt.Println(maxCoins1([]int{-3,-3,-3}, []int{9,-2,4})) // 11
    fmt.Println(maxCoins1([]int{-10}, []int{-2})) // -2
    fmt.Println(maxCoins1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 65

    fmt.Println(maxCoins2([]int{1,-2,-10,3}, []int{-5,10,0,1})) // 14
    fmt.Println(maxCoins2([]int{1,-1,-1,-1}, []int{0,3,4,-5})) // 8
    fmt.Println(maxCoins2([]int{-5,-4,-3}, []int{-1,2,3})) // 5
    fmt.Println(maxCoins2([]int{-3,-3,-3}, []int{9,-2,4})) // 11
    fmt.Println(maxCoins2([]int{-10}, []int{-2})) // -2
    fmt.Println(maxCoins2([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 65
}