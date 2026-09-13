package main

// 4050. Minimum Days to Score Exactly N Points
// You are given an integer n representing a target score.

// Your score starts at 0, and each day you either earn points or skip.

// Points are earned during a streak. 
// On the first day of a streak you earn 1 point, on the second day 2 points, on the third day 3 points, and so on. 
// Skipping a day earns nothing and resets the streak, so the next time you earn points, you start from 1 again.

// Return the minimum number of days, including any skipped days, needed to reach a score of exactly n.

// Example 1:
// Input: n = 2
// Output: 3
// Explanation:​​​​​​​
// Day 1: earn 1 point. Score is 1.
// Day 2: skip, which resets the streak. Earning here would add 2 points and take the score past n = 2.
// Day 3: the streak has reset, so earning gives 1 point. Score is exactly n = 2 in 3 days.

// Example 2:
// Input: n = 9
// Output: 6
// Explanation:​​​​​​​
// Days 1 to 3: earn 1, 2, and 3 points. Score is 1 + 2 + 3 = 6.
// Day 4: skip, which resets the streak.
// Days 5 and 6: earn 1 and 2 points. Score is exactly 6 + 1 + 2 = 9 in 6 days.

// Example 3:
// Input: n = 12
// Output: 7
// Explanation:​​​​​​​
// Days 1 to 3: earn 1, 2, and 3 points. Score is 1 + 2 + 3 = 6.
// Day 4: skip, which resets the streak.
// Days 5 to 7: earn 1, 2, and 3 points. Score is exactly 6 + 1 + 2 + 3 = 12 in 7 days.

// Constraints:
//     1 <= n <= 10^5

import "fmt"

func minDays(n int) int {
    dp, score, inf := make([]int, n+1), make([]int, 0), 1 << 61
    for i := range dp {
        dp[i] = inf
    }
    dp[0] = -1
    for k := 1; ; k++ {
        s := k * (k + 1) / 2
        if s > n {
            break
        }
        score = append(score, s)
    }
    for i := 1; i <= n; i++ {
        for k := 1; k <= len(score); k++ {
            s := score[k-1]
            if s > i {
                break
            }
            if dp[i-s] != inf {
                if dp[i-s]+k+1 < dp[i] {
                    dp[i] = dp[i-s] + k + 1
                }
            }
        }
    }
    return dp[n]
}

func main() {
    // Example 1:
    // Input: n = 2
    // Output: 3
    // Explanation:​​​​​​​
    // Day 1: earn 1 point. Score is 1.
    // Day 2: skip, which resets the streak. Earning here would add 2 points and take the score past n = 2.
    // Day 3: the streak has reset, so earning gives 1 point. Score is exactly n = 2 in 3 days.
    fmt.Println(minDays(2)) // 3
    // Example 2:
    // Input: n = 9
    // Output: 6
    // Explanation:​​​​​​​
    // Days 1 to 3: earn 1, 2, and 3 points. Score is 1 + 2 + 3 = 6.
    // Day 4: skip, which resets the streak.
    // Days 5 and 6: earn 1 and 2 points. Score is exactly 6 + 1 + 2 = 9 in 6 days.
    fmt.Println(minDays(9)) // 6
    // Example 3:
    // Input: n = 12
    // Output: 7
    // Explanation:​​​​​​​
    // Days 1 to 3: earn 1, 2, and 3 points. Score is 1 + 2 + 3 = 6.
    // Day 4: skip, which resets the streak.
    // Days 5 to 7: earn 1, 2, and 3 points. Score is exactly 6 + 1 + 2 + 3 = 12 in 7 days.
    fmt.Println(minDays(12)) // 7

    fmt.Println(minDays(1)) // 1
    fmt.Println(minDays(3)) // 2
    fmt.Println(minDays(8)) // 7
    fmt.Println(minDays(64)) // 16
    fmt.Println(minDays(99)) // 19
    fmt.Println(minDays(100)) // 20
    fmt.Println(minDays(101)) // 18
    fmt.Println(minDays(1024)) // 356
    fmt.Println(minDays(99_999)) // 480
    fmt.Println(minDays(100_000)) // 481
}