package main

// 1553. Minimum Number of Days to Eat N Oranges
// There are n oranges in the kitchen and you decided to eat some of these oranges every day as follows:
//     Eat one orange.
//     If the number of remaining oranges n is divisible by 2 then you can eat n / 2 oranges.
//     If the number of remaining oranges n is divisible by 3 then you can eat 2 * (n / 3) oranges.

// You can only choose one of the actions per day.
// Given the integer n, return the minimum number of days to eat n oranges.

// Example 1:
// Input: n = 10
// Output: 4
// Explanation: You have 10 oranges.
// Day 1: Eat 1 orange,  10 - 1 = 9.  
// Day 2: Eat 6 oranges, 9 - 2*(9/3) = 9 - 6 = 3. (Since 9 is divisible by 3)
// Day 3: Eat 2 oranges, 3 - 2*(3/3) = 3 - 2 = 1. 
// Day 4: Eat the last orange  1 - 1  = 0.
// You need at least 4 days to eat the 10 oranges.

// Example 2:
// Input: n = 6
// Output: 3
// Explanation: You have 6 oranges.
// Day 1: Eat 3 oranges, 6 - 6/2 = 6 - 3 = 3. (Since 6 is divisible by 2).
// Day 2: Eat 2 oranges, 3 - 2*(3/3) = 3 - 2 = 1. (Since 3 is divisible by 3)
// Day 3: Eat the last orange  1 - 1  = 0.
// You need at least 3 days to eat the 6 oranges.
 
// Constraints:
//     1 <= n <= 2 * 10^9

import "fmt"

// func minDays(n int) int {
//     days := 0
//     for n > 0 {
//         days++
//         if n % 3 == 0 { // 如果剩余橘子数 n 能被 3 整除，那么你可以吃掉 2*(n/3) 个橘子。
//             n -= 2*(n/3)
//             continue
//         }
//         if n %2 == 0 { // 如果剩余橘子数 n 能被 2 整除，那么你可以吃掉 n/2 个橘子
//             n -= n / 2
//             continue
//         }
//         n-- // 吃掉一个橘子。
//     }
//     return days
// }

// dfs
func minDays(n int) int {
    dp := map[int]int{
        0: 0,
        1: 1,
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(n int) int
    dfs = func(n int) int {
        if v, f := dp[n]; f {
            return v
        }
        one := 1 + (n % 2) + dfs(n/2) // 如果剩余橘子数 n 能被 2 整除，那么你可以吃掉 n/2 个橘子
        two := 1 + (n % 3) + dfs(n/3) // 如果剩余橘子数 n 能被 3 整除，那么你可以吃掉 2*(n/3) 个橘子
        dp[n] = min(one, two) // 取最优的
        return dp[n]
    }
    return dfs(n)
}

func minDays1(n int) int {
    dp := map[int]int{}
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(i int) int
    dfs = func(i int) int {
        if i <= 1 {
            return i
        }
        if dp[i] > 0 {
            return dp[i]
        }
        dp[i] = min(dfs(i/2) + i % 2, dfs(i/3) + i % 3) + 1
        return dp[i]
    }
    return dfs(n)
}

func main() {
    // Example 1:
    // Input: n = 10
    // Output: 4
    // Explanation: You have 10 oranges.
    // Day 1: Eat 1 orange,  10 - 1 = 9.  
    // Day 2: Eat 6 oranges, 9 - 2*(9/3) = 9 - 6 = 3. (Since 9 is divisible by 3)
    // Day 3: Eat 2 oranges, 3 - 2*(3/3) = 3 - 2 = 1. 
    // Day 4: Eat the last orange  1 - 1  = 0.
    // You need at least 4 days to eat the 10 oranges.
    fmt.Println(minDays(10)) // 4
    // Example 2:
    // Input: n = 6
    // Output: 3
    // Explanation: You have 6 oranges.
    // Day 1: Eat 3 oranges, 6 - 6/2 = 6 - 3 = 3. (Since 6 is divisible by 2).
    // Day 2: Eat 2 oranges, 3 - 2*(3/3) = 3 - 2 = 1. (Since 3 is divisible by 3)
    // Day 3: Eat the last orange  1 - 1  = 0.
    // You need at least 3 days to eat the 6 oranges.  
    fmt.Println(minDays(6)) // 3

    fmt.Println(minDays1(10)) // 4
    fmt.Println(minDays1(6)) // 3
}