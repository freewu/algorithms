package main

// 1420. Build Array Where You Can Find The Maximum Exactly K Comparisons
// You are given three integers n, m and k. 
// Consider the following algorithm to find the maximum element of an array of positive integers:
// <img src="https://assets.leetcode.com/uploads/2020/04/02/e.png" />

// You should build the array arr which has the following properties:
//     arr has exactly n integers.
//     1 <= arr[i] <= m where (0 <= i < n).
//     After applying the mentioned algorithm to arr, the value search_cost is equal to k.

// Return the number of ways to build the array arr under the mentioned conditions. 
// As the answer may grow large, the answer must be computed modulo 10^9 + 7.

// Example 1:
// Input: n = 2, m = 3, k = 1
// Output: 6
// Explanation: The possible arrays are [1, 1], [2, 1], [2, 2], [3, 1], [3, 2] [3, 3]

// Example 2:
// Input: n = 5, m = 2, k = 3
// Output: 0
// Explanation: There are no possible arrays that satisfy the mentioned conditions.

// Example 3:
// Input: n = 9, m = 1, k = 1
// Output: 1
// Explanation: The only possible array is [1, 1, 1, 1, 1, 1, 1, 1, 1]

// Constraints:
//     1 <= n <= 50
//     1 <= m <= 100
//     0 <= k <= n

import "fmt"

func numOfArrays(n int, m int, k int) int {
    //var dp[51][51][101] int
    dp := make([][][]int, 51)
    for i := range dp {
        dp[i] = make([][]int, 51)
        for j := range dp[i] {
            dp[i][j] = make([]int, 101)
            for v := range dp[i][j] {
                dp[i][j][v] = -1 // fill -1
            }
        }
    }
    var solve func(idx, lis, largest int) int
    solve = func(idx, lis, largest int) int {
        if idx == n {
            if lis == k { return 1 }
            return 0
        }
        if dp[idx][lis][largest] != -1 { return dp[idx][lis][largest] }
        res := 0
        for i := 1; i <= m; i++ {
            if i > largest {
                res += solve(idx+1, lis+1, i)
            } else {
                res += solve(idx+1, lis, largest)
            }
        }
        res = res % 1_000_000_007
        dp[idx][lis][largest] = res
        return res
    }
    return solve(0, 0, 0)
}

// 有可能入参的三个数刚好就是答案需要的三个数
func numOfArrays1(n int, m int, k int) int {
    mod := 1_000_000_007
    if k == 0 { return 0 }
    dp := make([][][]int, n + 1)
    for i := range dp {
        dp[i] = make([][]int, m + 1)
        for j := range dp[i] {
            dp[i][j] = make([]int, k + 1)
        }
    }
    // 边界条件 长度为1的数组搜索代价是1
    // 第n个数，最大是m，搜索次数是k
    // 看第n个数是否改变搜索次数
    for i := 0; i <= m; i++ {
        dp[1][i][1] = 1
    }
    for i := 2; i <= n; i++ { // 从第二个开始搜索
        for j := 1; j <= k && j <= i; j++ { //搜不超过数组长度
            for h := 1; h <= m; h++ {
                // 1-h 自由选择
                dp[i][h][j] = (dp[i-1][h][j] * h) % mod
                for hh := 1; hh < h; hh++ {
                    dp[i][h][j] += dp[i-1][hh][j-1]
                    dp[i][h][j] = dp[i][h][j] % mod
                }
            }
        }
    }
    res := 0
    for i := 1; i <= m; i++ {
        res += dp[n][i][k]
        res = res % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 2, m = 3, k = 1
    // Output: 6
    // Explanation: The possible arrays are [1, 1], [2, 1], [2, 2], [3, 1], [3, 2] [3, 3]
    fmt.Println(numOfArrays(2,3,1)) // 6
    // Example 2:
    // Input: n = 5, m = 2, k = 3
    // Output: 0
    // Explanation: There are no possible arrays that satisfy the mentioned conditions.
    fmt.Println(numOfArrays(5,2,3)) // 0
    // Example 3:
    // Input: n = 9, m = 1, k = 1
    // Output: 1
    // Explanation: The only possible array is [1, 1, 1, 1, 1, 1, 1, 1, 1]
    fmt.Println(numOfArrays(9,1,1)) // 1

    fmt.Println(numOfArrays(50,100,50)) // 538992043
    fmt.Println(numOfArrays(50,100,1)) // 435886633

    fmt.Println(numOfArrays1(2,3,1)) // 6
    fmt.Println(numOfArrays1(5,2,3)) // 0
    fmt.Println(numOfArrays1(9,1,1)) // 1
    fmt.Println(numOfArrays1(50,100,50)) // 538992043
    fmt.Println(numOfArrays1(50,100,1)) // 435886633
}