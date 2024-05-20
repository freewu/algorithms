package main

// 808. Soup Servings
// There are two types of soup: type A and type B. Initially, we have n ml of each type of soup. 
// There are four kinds of operations:
//     Serve 100 ml of soup A and 0 ml of soup B,
//     Serve 75 ml of soup A and 25 ml of soup B,
//     Serve 50 ml of soup A and 50 ml of soup B, and
//     Serve 25 ml of soup A and 75 ml of soup B.

// When we serve some soup, we give it to someone, and we no longer have it. 
// Each turn, we will choose from the four operations with an equal probability 0.25. 
// If the remaining volume of soup is not enough to complete the operation, we will serve as much as possible.
//  We stop once we no longer have some quantity of both types of soup.

// Note that we do not have an operation where all 100 ml's of soup B are used first.

// Return the probability that soup A will be empty first, 
// plus half the probability that A and B become empty at the same time. Answers within 10^-5 of the actual answer will be accepted.

// Example 1:
// Input: n = 50
// Output: 0.62500
// Explanation: If we choose the first two operations, A will become empty first.
// For the third operation, A and B will become empty at the same time.
// For the fourth operation, B will become empty first.
// So the total probability of A becoming empty first plus half the probability that A and B become empty at the same time, is 0.25 * (1 + 1 + 0.5 + 0) = 0.625.

// Example 2:
// Input: n = 100
// Output: 0.71875

// Constraints:
//     0 <= n <= 10^9

import "fmt"
import "strconv"

func soupServings(n int) float64 {
    m := (n + 24) / 25 // use units of 25 to measure and rounded up
    if m > 200 { return 1.0 } // above 200 units will get 1.00000 when rounded
    dp := make(map[string]float64)
    var prob func(a, b int, dp map[string]float64) float64
    prob = func(a, b int, dp map[string]float64) float64 {
        if a <= 0 && b <= 0 { return 0.5 } // half prob both empty at the same time
        if a <= 0 { return 1.0 } // a is definitely empty first
        if b <= 0 { return 0.0 } // a is definitely not empty first
        key := strconv.Itoa(a) + "," + strconv.Itoa(b)
        if _, ok := dp[key]; ok { return dp[key] }
        val := 0.25 * ( prob(a - 4, b, dp) +  prob(a - 3, b - 1, dp) +  prob(a - 2, b - 2, dp) +   prob(a - 1, b - 3, dp))
        dp[key] = val
        return val
    }
    return prob(m, m, dp)
}

func soupServings1(n int) float64 {
    n = (n + 24) / 25
    if n >= 179 {
        return 1
    }
    dp := make([][]float64, n + 1)
    for i := range dp {
        dp[i] = make([]float64, n + 1)
    }
    var dfs func(int, int) float64
    dfs = func(a, b int) float64 {
        if a <= 0 && b <= 0 { return 0.5 }
        if a <= 0 { return 1 }
        if b <= 0 { return 0 }
        if dp[a][b] > 0 {
            return dp[a][b]
        }
        dp[a][b] = (dfs(a - 4, b) + dfs(a - 3, b-1) + dfs(a-2, b-2) + dfs(a-1, b-3)) / 4
        return dp[a][b]
    }
    return dfs(n, n)
}

func main() {
    // Example 1:
    // Input: n = 50
    // Output: 0.62500
    // Explanation: If we choose the first two operations, A will become empty first.
    // For the third operation, A and B will become empty at the same time.
    // For the fourth operation, B will become empty first.
    // So the total probability of A becoming empty first plus half the probability that A and B become empty at the same time, is 0.25 * (1 + 1 + 0.5 + 0) = 0.625.
    fmt.Println(soupServings(50)) // 0.62500
    // Example 2:
    // Input: n = 100
    // Output: 0.71875
    fmt.Println(soupServings(100)) // 0.71875

    fmt.Println(soupServings1(50)) // 0.62500
    fmt.Println(soupServings1(100)) // 0.71875
}