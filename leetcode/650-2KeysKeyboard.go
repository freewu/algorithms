package main

// 650. 2 Keys Keyboard
// There is only one character 'A' on the screen of a notepad. 
// You can perform one of two operations on this notepad for each step:
//     Copy All: You can copy all the characters present on the screen (a partial copy is not allowed).
//     Paste: You can paste the characters which are copied last time.

// Given an integer n, return the minimum number of operations to get the character 'A' exactly n times on the screen.

// Example 1:
// Input: n = 3
// Output: 3
// Explanation: Initially, we have one character 'A'.
// In step 1, we use Copy All operation.
// In step 2, we use Paste operation to get 'AA'.
// In step 3, we use Paste operation to get 'AAA'.

// Example 2:
// Input: n = 1
// Output: 0

// Constraints:
//     1 <= n <= 1000

import "fmt"
import "math"

// 递归
func minSteps(n int) int {
    if n == 1 {
        return 0
    }
    res := n
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if n % i == 0 {
            // actually mn = minSteps(i) + 1 + n/1 - 1
            mn := minSteps(i) + n / i
            if mn < res {
                res = mn
            }
            mn = minSteps(n/i) + i
            if mn < res {
                res = mn
            }
        }
    }
    return res
}

func minSteps1(n int) int {
    h := int(math.Sqrt(float64(n)))
    dp := make([]int, n + 1)
    dp[0], dp[1] = 0, 0
    for i := 2; i <= n; i++ {
        dp[i] = i
        for j := 2; j <= h && j < i; j++ {
            if i%j == 0 {
                dp[i] = dp[j] + dp[i/j]
                break
            }
        }
    }
    return dp[n]
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: 3
    // Explanation: Initially, we have one character 'A'.
    // In step 1, we use Copy All operation.
    // In step 2, we use Paste operation to get 'AA'.
    // In step 3, we use Paste operation to get 'AAA'.
    fmt.Println(minSteps(3)) // 3
    // Example 2:
    // Input: n = 1
    // Output: 0
    fmt.Println(minSteps(1)) // 0
    fmt.Println(minSteps(10)) // 7
    fmt.Println(minSteps(100)) // 14
    fmt.Println(minSteps(1000)) // 21

    fmt.Println(minSteps1(3)) // 3
    fmt.Println(minSteps1(1)) // 0
    fmt.Println(minSteps1(10)) // 7
    fmt.Println(minSteps1(100)) // 14
    fmt.Println(minSteps1(1000)) // 21
}