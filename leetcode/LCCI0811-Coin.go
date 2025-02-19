package main

// 面试题 08.11. Coin LCCI
// Given an infinite number of quarters (25 cents), dimes (10 cents), nickels (5 cents), 
// and pennies (1 cent), write code to calculate the number of ways of representing n cents. 
// (The result may be large, so you should return it modulo 1000000007)

// Example1:
// Input: n = 5
// Output: 2
// Explanation: There are two ways:
// 5=5
// 5=1+1+1+1+1

// Example2:
// Input: n = 10
// Output: 4
// Explanation: There are four ways:
// 10=10
// 10=5+5
// 10=5+1+1+1+1+1
// 10=1+1+1+1+1+1+1+1+1+1

// Notes:
// You can assume:
//     0 <= n <= 1000000

import "fmt"

func waysToChange(n int) int {
    dp := make([]int,n + 1)
    dp[0] = 1
    coins := []int{ 1,5,10,25 }
    for _, coin := range coins {
        for i := coin; i <= n; i++ {
            dp[i] += dp[i - coin]
        }
    }
    return dp[n] % 1_000_000_007
}

func waysToChange1(n int) int {
    dp := make([]int, n + 1)
    dp[0] = 1
    for i := 1; i <= n; i++ {
        dp[i] = dp[i - 1]
    }
    for i := 25; i <= n; i++ {
        dp[i] += dp[i - 25]
    }
    for i := 5; i <= n; i++ {
        dp[i] += dp[i - 5]
    }
    for i := 10; i <= n; i++ {
        dp[i] += dp[i-10]
    }
    return dp[n] % 1_000_000_007
}

func main() {
    // Example1:
    // Input: n = 5
    // Output: 2
    // Explanation: There are two ways:
    // 5=5
    // 5=1+1+1+1+1
    fmt.Println(waysToChange(5)) // 2
    // Example2:
    // Input: n = 10
    // Output: 4
    // Explanation: There are four ways:
    // 10=10
    // 10=5+5
    // 10=5+1+1+1+1+1
    // 10=1+1+1+1+1+1+1+1+1+1
    fmt.Println(waysToChange(10)) // 4

    fmt.Println(waysToChange(0)) // 0
    fmt.Println(waysToChange(1)) // 1
    fmt.Println(waysToChange(2)) // 1
    fmt.Println(waysToChange(3)) // 1
    fmt.Println(waysToChange(8)) // 2
    fmt.Println(waysToChange(64)) // 73
    fmt.Println(waysToChange(999)) // 140430
    fmt.Println(waysToChange(1024)) // 151039
    fmt.Println(waysToChange(999999)) // 332496620
    fmt.Println(waysToChange(1000000)) // 332576607

    fmt.Println(waysToChange1(5)) // 2
    fmt.Println(waysToChange1(10)) // 4
    fmt.Println(waysToChange1(0)) // 0
    fmt.Println(waysToChange1(1)) // 1
    fmt.Println(waysToChange1(2)) // 1
    fmt.Println(waysToChange1(3)) // 1
    fmt.Println(waysToChange1(8)) // 2
    fmt.Println(waysToChange1(64)) // 73
    fmt.Println(waysToChange1(999)) // 140430
    fmt.Println(waysToChange1(1024)) // 151039
    fmt.Println(waysToChange1(999999)) // 332496620
    fmt.Println(waysToChange1(1000000)) // 332576607
}