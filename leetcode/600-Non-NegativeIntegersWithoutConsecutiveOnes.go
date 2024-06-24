package main

// 600. Non-negative Integers without Consecutive Ones
// Given a positive integer n, return the number of the integers in the range [0, n] 
// whose binary representations do not contain consecutive ones.

// Example 1:
// Input: n = 5
// Output: 5
// Explanation:
// Here are the non-negative integers <= 5 with their corresponding binary representations:
// 0 : 0
// 1 : 1
// 2 : 10
// 3 : 11
// 4 : 100
// 5 : 101
// Among them, only integer 3 disobeys the rule (two consecutive ones) and the other 5 satisfy the rule. 

// Example 2:
// Input: n = 1
// Output: 2

// Example 3:
// Input: n = 2
// Output: 3
 
// Constraints:
//     1 <= n <= 10^9

import "fmt"

func findIntegers(n int) int {
    memo := make([]int, 32)
    memo[0], memo[1] = 1, 2
    for i := 2; i < 32; i++ {
        memo[i] = memo[i-1] + memo[i-2]
    }
    res, k, preBit := 0, 30, 0
    for k >= 0 {
        if (n & (1 << k)) > 0 {
            res += memo[k]
            if preBit > 0 {
                return res
            }
            preBit = 1
        } else {
            preBit = 0
        }
        k--
    }
    return res + 1
}

func findIntegers1(n int) int {
    res, dp := 0, [31]int{1, 1}
    for i := 2; i < 31; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    for i, pre := 29, 0; i >= 0; i-- {
        v := 1 << i
        if n & v > 0 {
            res += dp[i+1]
            if pre == 1 {
                break
            }
            pre = 1
        } else {
            pre = 0
        }
        if i == 0 {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 5
    // Output: 5
    // Explanation:
    // Here are the non-negative integers <= 5 with their corresponding binary representations:
    // 0 : 0
    // 1 : 1
    // 2 : 10
    // 3 : 11
    // 4 : 100
    // 5 : 101
    // Among them, only integer 3 disobeys the rule (two consecutive ones) and the other 5 satisfy the rule. 
    fmt.Println(findIntegers(5)) // 5 
    // Example 2:
    // Input: n = 1
    // Output: 2
    fmt.Println(findIntegers(1)) // 2 
    // Example 3:
    // Input: n = 2
    // Output: 3
    fmt.Println(findIntegers(2)) // 3 

    fmt.Println(findIntegers1(5)) // 5 
    fmt.Println(findIntegers1(1)) // 2 
    fmt.Println(findIntegers1(2)) // 3
}