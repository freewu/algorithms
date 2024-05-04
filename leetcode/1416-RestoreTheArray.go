package main

// 1416. Restore The Array
// A program was supposed to print an array of integers. 
// The program forgot to print whitespaces and the array is printed as a string of digits s 
// and all we know is that all integers in the array were in the range [1, k] 
// and there are no leading zeros in the array.

// Given the string s and the integer k, return the number of the possible arrays that can be printed as s using the mentioned program. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: s = "1000", k = 10000
// Output: 1
// Explanation: The only possible array is [1000]

// Example 2:
// Input: s = "1000", k = 10
// Output: 0
// Explanation: There cannot be an array that was printed this way and has all integer >= 1 and <= 10.

// Example 3:
// Input: s = "1317", k = 2000
// Output: 8
// Explanation: Possible arrays are [1317],[131,7],[13,17],[1,317],[13,1,7],[1,31,7],[1,3,17],[1,3,1,7]

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of only digits and does not contain leading zeros.
//     1 <= k <= 10^9

import "fmt"

// dp
func numberOfArrays(s string, k int) int {
    n, mod := len(s), int(1e9 + 7)
    dp := make([]int, n + 1)
    dp[n] = 1
    for i := n - 1; i >= 0; i-- {
        val := int(s[i] - '0')
        for j := i + 1; val > 0 && val <= k && j <= n; j++ {
            dp[i] = (dp[i] + dp[j]) % mod
            if j < n {val = val * 10 + int(s[j] - '0')}
        }
    }
    return dp[0]
}

func numberOfArrays1(s string, k int) int {
    n, mod := len(s), int(1e9 + 7)
    dp := make([]int, n + 1)
    dp[n] = 1
    for i := n - 1; i >= 0; i-- {
        if s[i] == '0' {
            dp[i] = 0
            continue
        }
        num := 0
        for j := i; j < n; j++ {
            num = num*10 + int(s[j]-'0')
            if num > k {
                break
            }
            dp[i] = (dp[i] + dp[j+1]) % mod
        }
    }
    return dp[0]
}

func main() {
    // Example 1:
    // Input: s = "1000", k = 10000
    // Output: 1
    // Explanation: The only possible array is [1000]
    fmt.Println(numberOfArrays("1000", 10000)) // 1
    // Example 2: 
    // Input: s = "1000", k = 10
    // Output: 0
    // Explanation: There cannot be an array that was printed this way and has all integer >= 1 and <= 10.
    fmt.Println(numberOfArrays("1000", 10)) // 0
    // Example 3:
    // Input: s = "1317", k = 2000
    // Output: 8
    // Explanation: Possible arrays are [1317],[131,7],[13,17],[1,317],[13,1,7],[1,31,7],[1,3,17],[1,3,1,7]
    fmt.Println(numberOfArrays("1317", 2000)) // 8


    fmt.Println(numberOfArrays1("1000", 10000)) // 1
    fmt.Println(numberOfArrays1("1000", 10)) // 0
    fmt.Println(numberOfArrays1("1317", 2000)) // 8
}