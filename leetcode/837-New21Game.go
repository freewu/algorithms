package main

// 837. New 21 Game
// Alice plays the following game, loosely based on the card game "21".

// Alice starts with 0 points and draws numbers while she has less than k points. 
// During each draw, she gains an integer number of points randomly from the range [1, maxPts], where maxPts is an integer. 
// Each draw is independent and the outcomes have equal probabilities.

// Alice stops drawing numbers when she gets k or more points.

// Return the probability that Alice has n or fewer points.
// Answers within 10^-5 of the actual answer are considered accepted.

// Example 1:
// Input: n = 10, k = 1, maxPts = 10
// Output: 1.00000
// Explanation: Alice gets a single card, then stops.

// Example 2:
// Input: n = 6, k = 1, maxPts = 10
// Output: 0.60000
// Explanation: Alice gets a single card, then stops.
// In 6 out of 10 possibilities, she is at or below 6 points.

// Example 3:
// Input: n = 21, k = 17, maxPts = 10
// Output: 0.73278
 
// Constraints:
//     0 <= k <= n <= 10^4
//     1 <= maxPts <= 10^4

import "fmt"

func new21Game(n int, k int, maxPts int) float64 {
    if maxPts + k - 1 <= n || k == 0 {
        return 1.0
    }
    dp, windowSum, res, maxP := make([]float64, n + 1), 1.0, 0.0, float64(maxPts)
    dp[0] = 1.0
    for i := 1; i <= n; i++ {
        dp[i] = windowSum / maxP
        if i < k {
            windowSum += dp[i]
        } else {
            res += dp[i]
        }
        if i >= maxPts {
            windowSum -= dp[i - maxPts]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 10, k = 1, maxPts = 10
    // Output: 1.00000
    // Explanation: Alice gets a single card, then stops.
    fmt.Println(new21Game(10, 1, 10)) // 1.00000
    // Example 2:
    // Input: n = 6, k = 1, maxPts = 10
    // Output: 0.60000
    // Explanation: Alice gets a single card, then stops.
    // In 6 out of 10 possibilities, she is at or below 6 points.
    fmt.Println(new21Game(6, 1, 10)) // 0.60000
    // Example 3:
    // Input: n = 21, k = 17, maxPts = 10
    // Output: 0.73278
    fmt.Println(new21Game(21, 17, 10)) // 0.73278
}