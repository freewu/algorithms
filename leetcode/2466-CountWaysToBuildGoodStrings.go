package main

// 2466. Count Ways To Build Good Strings
// Given the integers zero, one, low, and high, we can construct a string by starting with an empty string, 
// and then at each step perform either of the following:
//     Append the character '0' zero times.
//     Append the character '1' one times.

// This can be performed any number of times.
// A good string is a string constructed by the above process having a length between low and high (inclusive).
// Return the number of different good strings that can be constructed satisfying these properties.
// Since the answer can be large, return it modulo 109 + 7.

// Example 1:
// Input: low = 3, high = 3, zero = 1, one = 1
// Output: 8
// Explanation: 
// One possible valid good string is "011". 
// It can be constructed as follows: "" -> "0" -> "01" -> "011". 
// All binary strings from "000" to "111" are good strings in this example.

// Example 2:
// Input: low = 2, high = 3, zero = 1, one = 2
// Output: 5
// Explanation: The good strings are "00", "11", "000", "110", and "011".
 
// Constraints:
//     1 <= low <= high <= 10^5
//     1 <= zero, one <= low

import "fmt"

func countGoodStrings(low int, high int, zero int, one int) (res int) {
    dp := make([]int,high + 1)
    dp[0] = 1
    for cur := 0; cur < high + 1; cur++ {
        cur0, cur1 := cur + zero, cur + one
        if cur0 <= high  {
            dp[cur0] += dp[cur] 
            dp[cur0] %= 1000000007
        }
        if cur1 <= high {
            dp[cur1] += dp[cur]
            dp[cur1] %= 1000000007
        }
        if cur >= low {
            res += dp[cur]
        }
    }
    return res  % 1000000007
}

func countGoodStrings1(low, high, zero, one int) int {
    res, mod := 0, 1_000_000_007
    dp := make([]int, high + 1) // dp[i] 表示构造长为 i 的字符串的方案数
    dp[0] = 1                 // 构造空串的方案数为 1
    for i := 1; i <= high; i++ {
        if i >= one {
            dp[i] = (dp[i] + dp[i-one]) % mod
        }
        if i >= zero {
            dp[i] = (dp[i] + dp[i-zero]) % mod
        }
        if i >= low {
            res = (res + dp[i]) % mod
        }
    }
    return res 
}

func main() {
    // One possible valid good string is "011". 
    // It can be constructed as follows: "" -> "0" -> "01" -> "011". 
    // All binary strings from "000" to "111" are good strings in this example.
    fmt.Println(countGoodStrings(3,3,1,1)) // 8
    // Explanation: The good strings are "00", "11", "000", "110", and "011".
    fmt.Println(countGoodStrings(2,3,1,2)) // 5

    fmt.Println(countGoodStrings1(3,3,1,1)) // 8
    fmt.Println(countGoodStrings1(2,3,1,2)) // 5
}