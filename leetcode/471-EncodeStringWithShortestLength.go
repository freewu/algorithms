package main

// 471. Encode String with Shortest Length
// Given a string s, encode the string such that its encoded length is the shortest.
// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. k should be a positive integer.
// If an encoding process does not make the string shorter, then do not encode it. 
// If there are several solutions, return any of them.

// Example 1:
// Input: s = "aaa"
// Output: "aaa"
// Explanation: There is no way to encode it such that it is shorter than the input string, so we do not encode it.

// Example 2:
// Input: s = "aaaaa"
// Output: "5[a]"
// Explanation: "5[a]" is shorter than "aaaaa" by 1 character.

// Example 3:
// Input: s = "aaaaaaaaaa"
// Output: "10[a]"
// Explanation: "a9[a]" or "9[a]a" are also valid solutions, both of them have the same length = 5, which is the same as "10[a]".
 
// Constraints:
//     1 <= s.length <= 150
//     s consists of only lowercase English letters.

import "fmt"
import "strings"
import "strconv"

func encode(s string) string {
    n := len(s)
    dp := make([][]string, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]string, n)
        dp[i][i] = string(s[i])
    }
    findRepSubSize := func(s string) int { return strings.Index(s[1:]+s, s) + 1; }
    var numLen func(n int) int
    numLen = func(n int) int {
        if n < 10 { return 1; }
        return 1 + numLen(n / 10)
    }
    for l := 2; l <= n; l++ {
        for i := 0; i+l <= n; i++ {
            j := i + l - 1
            dp[i][j] = s[i : j+1]
            subSize := findRepSubSize(s[i : j+1])
            if subSize < l && numLen(l/subSize)+2+len(dp[i][i+subSize-1]) < len(dp[i][j]) {
                dp[i][j] = strconv.Itoa(l/subSize) + "[" + dp[i][i+subSize-1] + "]"
            }
            // fmt.Println(dp[i][j], i, j)
            for k := i; k < j; k++ {
                if len(dp[i][k])+len(dp[k+1][j]) < len(dp[i][j]) {
                    dp[i][j] = dp[i][k] + dp[k+1][j]
                }
            }
        }
    }
    return dp[0][n-1]
}

func main() {
    // Example 1:
    // Input: s = "aaa"
    // Output: "aaa"
    // Explanation: There is no way to encode it such that it is shorter than the input string, so we do not encode it.
    fmt.Println(encode("aaa")) // aaa
    // Example 2:
    // Input: s = "aaaaa"
    // Output: "5[a]"
    // Explanation: "5[a]" is shorter than "aaaaa" by 1 character.
    fmt.Println(encode("aaaaa")) // 5[a]
    // Example 3:
    // Input: s = "aaaaaaaaaa"
    // Output: "10[a]"
    // Explanation: "a9[a]" or "9[a]a" are also valid solutions, both of them have the same length = 5, which is the same as "10[a]".
    fmt.Println(encode("aaaaaaaaaa")) // 10[a]
}