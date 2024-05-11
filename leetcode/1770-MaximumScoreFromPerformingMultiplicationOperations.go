package main

// 1770. Maximum Score from Performing Multiplication Operations
// You are given two 0-indexed integer arrays nums and multipliers of size n and m respectively, where n >= m.

// You begin with a score of 0. You want to perform exactly m operations. 
// On the ith operation (0-indexed) you will:
//     Choose one integer x from either the start or the end of the array nums.
//     Add multipliers[i] * x to your score.
//     Note that multipliers[0] corresponds to the first operation, multipliers[1] to the second operation, and so on.
//     Remove x from nums.

// Return the maximum score after performing m operations.

// Example 1:
// Input: nums = [1,2,3], multipliers = [3,2,1]
// Output: 14
// Explanation: An optimal solution is as follows:
// - Choose from the end, [1,2,3], adding 3 * 3 = 9 to the score.
// - Choose from the end, [1,2], adding 2 * 2 = 4 to the score.
// - Choose from the end, [1], adding 1 * 1 = 1 to the score.
// The total score is 9 + 4 + 1 = 14.

// Example 2:
// Input: nums = [-5,-3,-3,-2,7,1], multipliers = [-10,-5,3,4,6]
// Output: 102
// Explanation: An optimal solution is as follows:
// - Choose from the start, [-5,-3,-3,-2,7,1], adding -5 * -10 = 50 to the score.
// - Choose from the start, [-3,-3,-2,7,1], adding -3 * -5 = 15 to the score.
// - Choose from the start, [-3,-2,7,1], adding -3 * 3 = -9 to the score.
// - Choose from the end, [-2,7,1], adding 1 * 4 = 4 to the score.
// - Choose from the end, [-2,7], adding 7 * 6 = 42 to the score. 
// The total score is 50 + 15 - 9 + 4 + 42 = 102.
 
// Constraints:
//     n == nums.length
//     m == multipliers.length
//     1 <= m <= 300
//     m <= n <= 10^5
//     -1000 <= nums[i], multipliers[i] <= 1000

import "fmt"

func maximumScore(nums []int, multipliers []int) int {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    n, m := len(nums), len(multipliers)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, m + 1)
    }
    for i := m - 1; i >= 0; i-- {
        for j := i; j >= 0; j-- {
            index := n - 1 + j - i
            l := nums[j] * multipliers[i] + dp[i+1][j+1]
            r := nums[index] * multipliers[i] + dp[i+1][j]
            dp[i][j] = max(l, r)
        }
    }
    return dp[0][0]
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3], multipliers = [3,2,1]
    // Output: 14
    // Explanation: An optimal solution is as follows:
    // - Choose from the end, [1,2,3], adding 3 * 3 = 9 to the score.
    // - Choose from the end, [1,2], adding 2 * 2 = 4 to the score.
    // - Choose from the end, [1], adding 1 * 1 = 1 to the score.
    // The total score is 9 + 4 + 1 = 14.
    fmt.Println(maximumScore([]int{1,2,3},[]int{3,2,1})) // 14
    // Example 2:
    // Input: nums = [-5,-3,-3,-2,7,1], multipliers = [-10,-5,3,4,6]
    // Output: 102
    // Explanation: An optimal solution is as follows:
    // - Choose from the start, [-5,-3,-3,-2,7,1], adding -5 * -10 = 50 to the score.
    // - Choose from the start, [-3,-3,-2,7,1], adding -3 * -5 = 15 to the score.
    // - Choose from the start, [-3,-2,7,1], adding -3 * 3 = -9 to the score.
    // - Choose from the end, [-2,7,1], adding 1 * 4 = 4 to the score.
    // - Choose from the end, [-2,7], adding 7 * 6 = 42 to the score. 
    // The total score is 50 + 15 - 9 + 4 + 42 = 102.
    fmt.Println(maximumScore([]int{-5,-3,-3,-2,7,1},[]int{-10,-5,3,4,6})) // 102
}