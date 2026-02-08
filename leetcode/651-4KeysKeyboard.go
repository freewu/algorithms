package main 

// 651. 4 Keys Keyboard
// Imagine you have a special keyboard with the following keys:
//     A: Print one 'A' on the screen.
//     Ctrl-A: Select the whole screen.
//     Ctrl-C: Copy selection to buffer.
//     Ctrl-V: Print buffer on screen appending it after what has already been printed.

// Given an integer n, return the maximum number of 'A' you can print on the screen with at most n presses on the keys. 

// Example 1:
// Input: n = 3
// Output: 3
// Explanation: We can at most get 3 A's on screen by pressing the following key sequence:
// A, A, A

// Example 2:
// Input: n = 7
// Output: 9
// Explanation: We can at most get 9 A's on screen by pressing following key sequence:
// A, A, A, Ctrl A, Ctrl C, Ctrl V, Ctrl V
 
// Constraints:
//     1 <= n <= 50

import "fmt"

// dp
func maxA(n int) int {
    dp := make([]int,n+1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        dp[i] = dp[i-1] + 1
        for j := i - 3; j >= 0; j-- {
            // 第i次可以是前面第j次的值经过ca,cc,cv得到，j每往前一位都可以多cv一次，取所有j经过i-3-j+1次cv的值的最大
            dp[i] = max(dp[j] * (i-3-j+1) + dp[j], dp[i])
            dp[i] = max(dp[i], 2 * dp[j] + i - 3 -j)
        }
    }
    return dp[n]
}

func main() {
    // Example 1:
    // Input: n = 3
    // Output: 3
    // Explanation: We can at most get 3 A's on screen by pressing the following key sequence:
    // A, A, A
    fmt.Println(maxA(3)) // 3

    // Example 2:
    // Input: n = 7
    // Output: 9
    // Explanation: We can at most get 9 A's on screen by pressing following key sequence:
    // A, A, A, Ctrl A, Ctrl C, Ctrl V, Ctrl V
    fmt.Println(maxA(7)) // 9

    fmt.Println(maxA(1)) // 1
    fmt.Println(maxA(2)) // 2
    fmt.Println(maxA(4)) // 4
    fmt.Println(maxA(8)) // 12
    fmt.Println(maxA(49)) // 1048576
    fmt.Println(maxA(50)) // 1327104
}