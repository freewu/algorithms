package main

// 1223. Dice Roll Simulation
// A die simulator generates a random number from 1 to 6 for each roll. 
// You introduced a constraint to the generator such that it cannot roll the number i more than rollMax[i] (1-indexed) consecutive times.

// Given an array of integers rollMax and an integer n, return the number of distinct sequences that can be obtained with exact n rolls. 
// Since the answer may be too large, return it modulo 10^9 + 7.

// Two sequences are considered different if at least one element differs from each other.

// Example 1:
// Input: n = 2, rollMax = [1,1,2,2,2,3]
// Output: 34
// Explanation: There will be 2 rolls of die, if there are no constraints on the die, there are 6 * 6 = 36 possible combinations. In this case, looking at rollMax array, the numbers 1 and 2 appear at most once consecutively, therefore sequences (1,1) and (2,2) cannot occur, so the final answer is 36-2 = 34.

// Example 2:
// Input: n = 2, rollMax = [1,1,1,1,1,1]
// Output: 30

// Example 3:
// Input: n = 3, rollMax = [1,1,1,2,2,3]
// Output: 181

// Constraints:
//     1 <= n <= 5000
//     rollMax.length == 6
//     1 <= rollMax[i] <= 15

import "fmt"

func dieSimulator(n int, rollMax []int) int {
    dp, mod := make([][]int, n+1), 1_000_000_007
    for i := range dp {
        dp[i] = make([]int, 7)
    }
    for i := range dp[1] {
        dp[1][i] = 1
    }
    dp[1][6] = 6 //use dp[i][6] as sum of all results of i indices
    for i := 2; i <= n; i++ {
        sum := 0
        for j := 0; j <= 5; j++ {
            dp[i][j] = dp[i-1][6]
            if i == rollMax[j]+1 {
                dp[i][j] -= 1
            } else if i > rollMax[j]+1 {
                dp[i][j] -= dp[i-rollMax[j]-1][6] - dp[i-rollMax[j]-1][j]
                dp[i][j] = (dp[i][j] + mod) % mod
            }
            sum += dp[i][j]
            sum %= mod
        }
        dp[i][6] = sum
    }
    return dp[n][6]
}

func dieSimulator1(n int, rollMax []int) int {
    dp, mod := make([][]int, n), 1_000_000_007
    for i := range dp {
        dp[i] = make([]int, 6)
        for j := range dp[i] {
            dp[0][j] = 1
        }
    }
    arr := make([]int, n)
    arr[0] = 6
    for i := 1; i < n; i++ {
        t := 0
        for j, v := range rollMax {
            res, pre := arr[i-1], i - v
            if pre > 0 {
                res -= arr[pre-1] - dp[pre-1][j]
            } else if pre == 0 {
                res--
            }
            dp[i][j] = (res % mod + mod) % mod
            t += dp[i][j]
        }
        arr[i] = t % mod
    }
    return arr[n-1]
}

func main() {
    // Example 1:
    // Input: n = 2, rollMax = [1,1,2,2,2,3]
    // Output: 34
    // Explanation: There will be 2 rolls of die, if there are no constraints on the die, there are 6 * 6 = 36 possible combinations. In this case, looking at rollMax array, the numbers 1 and 2 appear at most once consecutively, therefore sequences (1,1) and (2,2) cannot occur, so the final answer is 36-2 = 34.
    fmt.Println(dieSimulator(2, []int{1,1,2,2,2,3})) // 34
    // Example 2:
    // Input: n = 2, rollMax = [1,1,1,1,1,1]
    // Output: 30
    fmt.Println(dieSimulator(2, []int{1,1,1,1,1,1})) // 30
    // Example 3:
    // Input: n = 3, rollMax = [1,1,1,2,2,3]
    // Output: 181
    fmt.Println(dieSimulator(3, []int{1,1,1,2,2,3})) // 181

    fmt.Println(dieSimulator1(2, []int{1,1,2,2,2,3})) // 34
    fmt.Println(dieSimulator1(2, []int{1,1,1,1,1,1})) // 30
    fmt.Println(dieSimulator1(3, []int{1,1,1,2,2,3})) // 181
}