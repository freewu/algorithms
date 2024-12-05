package main

// 2787. Ways to Express an Integer as Sum of Powers
// Given two positive integers n and x.

// Return the number of ways n can be expressed as the sum of the xth power of unique positive integers, 
// in other words, the number of sets of unique integers [n1, n2, ..., nk] where n = n1x + n2x + ... + nkx.

// Since the result can be very large, return it modulo 109 + 7.

// For example, if n = 160 and x = 3, one way to express n is n = 23 + 33 + 53.

// Example 1:
// Input: n = 10, x = 2
// Output: 1
// Explanation: We can express n as the following: n = 32 + 12 = 10.
// It can be shown that it is the only way to express 10 as the sum of the 2nd power of unique integers.

// Example 2:
// Input: n = 4, x = 1
// Output: 2
// Explanation: We can express n in the following ways:
// - n = 41 = 4.
// - n = 31 + 11 = 4.

// Constraints:
//     1 <= n <= 300
//     1 <= x <= 5

import "fmt"
import "math"

func numberOfWays(n int, x int) int {
    nums, k, mod := []int{}, 1, 1_000_000_007
    pow := func (a, b int) int {
        res := 1
        for ; b > 0; b >>= 1 {
            if (b & 1) == 1 {
                res = (res * a) % mod
            }
            a = ( a * a ) % mod
        }
        return res
    }
    for true {
        curr := pow(k, x)
        if curr > n { break }
        nums = append(nums, curr)
        k++
    }
    dp := make([]int, n + 1)
    dp[0] = 1
    for _ , v := range nums {
        for j := n ; j >= v ; j-- {
            if v <= j {
                dp[j] = (dp[j - v] + dp[j]) % mod
            }
        }
    }
    return dp[n]
}

func numberOfWays1(n int, x int) int {
    mx := 1
    pow := func (x, i int) int { return int(math.Pow(float64(x), float64(i))) }
    for ; pow(mx, x) <= n; mx++ {} // 确认最大值
    dp := make([]int,n + 1)
    dp[0]=1
    for i := 1; i <= mx; i++ {
        v := pow(i,x)
        for k := n; k >= v; k-- {
            dp[k] += dp[k-v] // 因为当前可以选择num，num为不能重复的幂等，那么dp[k]+=dp[k-num]
        }
    }
    return dp[n] % 1_000_000_007
}

func main() {
    // Example 1:
    // Input: n = 10, x = 2
    // Output: 1
    // Explanation: We can express n as the following: n = 32 + 12 = 10.
    // It can be shown that it is the only way to express 10 as the sum of the 2nd power of unique integers.
    fmt.Println(numberOfWays(10, 2)) // 1
    // Example 2:
    // Input: n = 4, x = 1
    // Output: 2
    // Explanation: We can express n in the following ways:
    // - n = 41 = 4.
    // - n = 31 + 11 = 4.
    fmt.Println(numberOfWays(4, 1)) // 2

    fmt.Println(numberOfWays1(10, 2)) // 1
    fmt.Println(numberOfWays1(4, 1)) // 2
}