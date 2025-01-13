package main

// 2585. Number of Ways to Earn Points
// There is a test that has n types of questions. 
// You are given an integer target and a 0-indexed 2D integer array types where types[i] = [counti, marksi] indicates that there are counti questions of the ith type, 
// and each one of them is worth marksi points.

// Return the number of ways you can earn exactly target points in the exam. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// Note that questions of the same type are indistinguishable.

// For example, if there are 3 questions of the same type, 
// then solving the 1st and 2nd questions is the same as solving the 1st and 3rd questions, 
// or the 2nd and 3rd questions.

// Example 1:
// Input: target = 6, types = [[6,1],[3,2],[2,3]]
// Output: 7
// Explanation: You can earn 6 points in one of the seven ways:
// - Solve 6 questions of the 0th type: 1 + 1 + 1 + 1 + 1 + 1 = 6
// - Solve 4 questions of the 0th type and 1 question of the 1st type: 1 + 1 + 1 + 1 + 2 = 6
// - Solve 2 questions of the 0th type and 2 questions of the 1st type: 1 + 1 + 2 + 2 = 6
// - Solve 3 questions of the 0th type and 1 question of the 2nd type: 1 + 1 + 1 + 3 = 6
// - Solve 1 question of the 0th type, 1 question of the 1st type and 1 question of the 2nd type: 1 + 2 + 3 = 6
// - Solve 3 questions of the 1st type: 2 + 2 + 2 = 6
// - Solve 2 questions of the 2nd type: 3 + 3 = 6

// Example 2:
// Input: target = 5, types = [[50,1],[50,2],[50,5]]
// Output: 4
// Explanation: You can earn 5 points in one of the four ways:
// - Solve 5 questions of the 0th type: 1 + 1 + 1 + 1 + 1 = 5
// - Solve 3 questions of the 0th type and 1 question of the 1st type: 1 + 1 + 1 + 2 = 5
// - Solve 1 questions of the 0th type and 2 questions of the 1st type: 1 + 2 + 2 = 5
// - Solve 1 question of the 2nd type: 5

// Example 3:
// Input: target = 18, types = [[6,1],[3,2],[2,3]]
// Output: 1
// Explanation: You can only earn 18 points by answering all questions.

// Constraints:
//     1 <= target <= 1000
//     n == types.length
//     1 <= n <= 50
//     types[i].length == 2
//     1 <= counti, marksi <= 50

import "fmt"

func waysToReachTarget(target int, types [][]int) int {
    dp := make([][]int, 2)
    dp[0], dp[1] = make([]int, target + 1), make([]int, target + 1)
    dp[0][0], dp[1][0] = 1, 1
    for i := 0; i < len(types); i++ {
        for j := 1; j <= target; j++ {
            dp[(i + 1) % 2][j] = dp[i % 2][j] 
            for k := 1; k <= types[i][0] && j >= k * types[i][1]; k++ {
                dp[(i + 1) % 2][j] = (dp[(i + 1) % 2][j] + dp[i % 2][j - k * types[i][1]]) % 1_000_000_007
            }
        }
    }
    return dp[len(types) % 2][target]
}

func waysToReachTarget1(target int, types [][]int) int {
    dp := make([]int, target + 1)
    dp[0] = 1
    for i := 0; i < len(types); i++{
        count1, count2 := types[i][0], types[i][1]
        for j := target; j > 0; j-- {
            for k := 1; k <= count1 && k * count2 <= j; k++ {
                dp[j] += dp[j - k * count2]
            }
            dp[j] %= 1_000_000_007
        }
    }
    return dp[target]
}

func main() {
    // Example 1:
    // Input: target = 6, types = [[6,1],[3,2],[2,3]]
    // Output: 7
    // Explanation: You can earn 6 points in one of the seven ways:
    // - Solve 6 questions of the 0th type: 1 + 1 + 1 + 1 + 1 + 1 = 6
    // - Solve 4 questions of the 0th type and 1 question of the 1st type: 1 + 1 + 1 + 1 + 2 = 6
    // - Solve 2 questions of the 0th type and 2 questions of the 1st type: 1 + 1 + 2 + 2 = 6
    // - Solve 3 questions of the 0th type and 1 question of the 2nd type: 1 + 1 + 1 + 3 = 6
    // - Solve 1 question of the 0th type, 1 question of the 1st type and 1 question of the 2nd type: 1 + 2 + 3 = 6
    // - Solve 3 questions of the 1st type: 2 + 2 + 2 = 6
    // - Solve 2 questions of the 2nd type: 3 + 3 = 6
    fmt.Println(waysToReachTarget(6, [][]int{{6,1},{3,2},{2,3}})) // 7
    // Example 2:
    // Input: target = 5, types = [[50,1],[50,2],[50,5]]
    // Output: 4
    // Explanation: You can earn 5 points in one of the four ways:
    // - Solve 5 questions of the 0th type: 1 + 1 + 1 + 1 + 1 = 5
    // - Solve 3 questions of the 0th type and 1 question of the 1st type: 1 + 1 + 1 + 2 = 5
    // - Solve 1 questions of the 0th type and 2 questions of the 1st type: 1 + 2 + 2 = 5
    // - Solve 1 question of the 2nd type: 5
    fmt.Println(waysToReachTarget(5, [][]int{{50,1},{50,2},{50,5}})) // 4
    // Example 3:
    // Input: target = 18, types = [[6,1],[3,2],[2,3]]
    // Output: 1
    // Explanation: You can only earn 18 points by answering all questions.
    fmt.Println(waysToReachTarget(1, [][]int{{6,1},{3,2},{2,3}})) // 1

    fmt.Println(waysToReachTarget1(6, [][]int{{6,1},{3,2},{2,3}})) // 7
    fmt.Println(waysToReachTarget1(5, [][]int{{50,1},{50,2},{50,5}})) // 4
    fmt.Println(waysToReachTarget1(1, [][]int{{6,1},{3,2},{2,3}})) // 1
}