package main

// 1246. Palindrome Removal
// You are given an integer array arr.

// In one move, you can select a palindromic subarray arr[i], arr[i + 1], ..., arr[j] where i <= j, and remove that subarray from the given array. 
// Note that after removing a subarray, the elements on the left and on the right of that subarray move to fill the gap left by the removal.

// Return the minimum number of moves needed to remove all numbers from the array.

// Example 1:
// Input: arr = [1,2]
// Output: 2

// Example 2:
// Input: arr = [1,3,4,1,5]
// Output: 3
// Explanation: Remove [4] then remove [1,3,1] then remove [5].

// Constraints:
//     1 <= arr.length <= 100
//     1 <= arr[i] <= 20

import "fmt"

func minimumMoves(arr []int) int {
    dp, inf := make([][]int, len(arr)), 1 << 31 // dp[i][j] 表示删除arr[i: j + 1]的最少次数
    for i := range dp {
        dp[i] = make([]int, len(arr))
        for j := range dp[i] {
            if j > i {
                dp[i][j] = inf
            }
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    dp[0][0] = 1
    for j := 1; j < len(arr); j ++ {
        dp[j][j] = 1
        for i := j - 1; i >= 0; i -- {
            if arr[i] == arr[j] {
                if j - i >= 2 {
                    dp[i][j] = min(dp[i + 1][j - 1], dp[i][j])
                } else {
                    dp[i][j] = 1
                }
            } 
            for k := i; k < j; k ++ {
                dp[i][j] = min(dp[i][k] + dp[k + 1][j], dp[i][j])
            }
        }
    }
    return dp[0][len(arr) - 1]
}

func main() {
    // Example 1:
    // Input: arr = [1,2]
    // Output: 2
    fmt.Println(minimumMoves([]int{1,2})) // 2
    // Example 2:
    // Input: arr = [1,3,4,1,5]
    // Output: 3
    // Explanation: Remove [4] then remove [1,3,1] then remove [5].
    fmt.Println(minimumMoves([]int{1,3,4,1,5})) // 3
}