package main

// 1531. String Compression II
// Run-length encoding is a string compression method that works by replacing consecutive identical characters (repeated 2 or more times) 
// with the concatenation of the character and the number marking the count of the characters (length of the run). 
// For example, to compress the string "aabccc" we replace "aa" by "a2" and replace "ccc" by "c3". 
// Thus the compressed string becomes "a2bc3".

// Notice that in this problem, we are not adding '1' after single characters.

// Given a string s and an integer k. 
// You need to delete at most k characters from s such that the run-length encoded version of s has minimum length.

// Find the minimum length of the run-length encoded version of s after deleting at most k characters.

// Example 1:
// Input: s = "aaabcccd", k = 2
// Output: 4
// Explanation: Compressing s without deleting anything will give us "a3bc3d" of length 6. Deleting any of the characters 'a' or 'c' would at most decrease the length of the compressed string to 5, for instance delete 2 'a' then we will have s = "abcccd" which compressed is abc3d. Therefore, the optimal way is to delete 'b' and 'd', then the compressed version of s will be "a3c3" of length 4.

// Example 2:
// Input: s = "aabbaa", k = 2
// Output: 2
// Explanation: If we delete both 'b' characters, the resulting compressed string would be "a4" of length 2.

// Example 3:
// Input: s = "aaaaaaaaaaa", k = 0
// Output: 3
// Explanation: Since k is zero, we cannot delete anything. The compressed string is "a11" of length 3.

// Constraints:
//     1 <= s.length <= 100
//     0 <= k <= s.length
//     s contains only lowercase English letters.

import "fmt"

// // Wrong Answer 143 / 144 
// func getLengthOfOptimalCompression(s string, k int) int {
//     m, n := len(s), k + 1
//     dp := make([][]int, m)
//     for i := 0; i < m; i++ {
//         dp[i] = make([]int, n)
//         for j := 0; j < n; j++ {
//             dp[i][j] = -1
//         }
//     }
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     getLen := func(most int) int {
//         if most == 1 { return 0 }
//         if most < 10 { return 1 }
//         if most < 99 { return 2 }
//         return 3
//     }
//     var solve func(currIdx, rest int) int
//     solve = func(currIdx, rest int) int {
//         if currIdx == len(s) || len(s) - currIdx <= rest { return 0 }
//         if dp[currIdx][rest] != -1 { return dp[currIdx][rest] }
//         free := make([]int, 26)
//         most, res := 0, 1 << 31
//         for i := currIdx; i < len(s); i++ {
//             idx := s[i] - 'a'
//             free[idx]++
//             if most < free[idx] { most = free[idx] }
//             if rest >= i - currIdx + 1 - most {
//                 res = min(res, getLen(most) + 1 + solve(i + 1, rest - ( i - currIdx + 1 - most)))
//             }
//         }
//         dp[currIdx][rest] = res
//         return res
//     }
//     return solve(0, k)
// }


func getLengthOfOptimalCompression(s string, k int) int {
    n, val := len(s), 1_000_000_007
    dp := make([][]int, n+1)
    dp[0] = make([]int, k+1)
    for i := 1; i < n+1; i += 1 {
        dp[i] = make([]int, k+1)
        for j := 0; j < k+1; j += 1 {
            dp[i][j] = val
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    calcLen := func(count int) int {
        if count == 0 || count == 1 { return count }
        if count < 10 { return 2 }
        if count < 100 { return 3 }
        return 4
    }
    for i := 1; i < n+1; i += 1 {
        for j := 0; j < k+1; j += 1 {
            if j > 0 {
                dp[i][j] = dp[i-1][j-1]
            }
            rm, cnt := 0, 0
            for pe := i; pe >= 1; pe -= 1 {
                if s[pe-1] == s[i-1] {
                    cnt += 1
                } else {
                    rm += 1
                    if rm > j { break }
                }
                a := dp[i][j]
                b := dp[pe-1][j-rm] + calcLen(cnt)
                dp[i][j] = min(a, b)
            }
        }
    }
    return dp[n][k]
}

// dfs
func getLengthOfOptimalCompression1(s string, k int) int {
    n := len(s)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, k+1)
        for j := range dp[i] {
            dp[i][j] = 101
        }
    }
    var dfs func(i, k int) int
    dfs = func(i, k int) int {
        if i + k >= n { return 0 }
        if k < 0 { return 101 }
        if dp[i][k] != 101 { return dp[i][k] }
        res := dfs(i+1, k-1)
        diff, same, l := 0, 0, 0
        for j := i; j < n; j++ {
            if k - diff < 0 { break }
            if s[i] == s[j] {
                same++
                if same <= 2 || same == 10 || same == 100 { l++ }
            } else {
                diff++
            }
            res = min( res, l + dfs(j + 1, k - diff))
        }
        dp[i][k] = res
        return res
    }
    return dfs(0, k)
}

func main() {
    // Example 1:
    // Input: s = "aaabcccd", k = 2
    // Output: 4
    // Explanation: Compressing s without deleting anything will give us "a3bc3d" of length 6. Deleting any of the characters 'a' or 'c' would at most decrease the length of the compressed string to 5, for instance delete 2 'a' then we will have s = "abcccd" which compressed is abc3d. Therefore, the optimal way is to delete 'b' and 'd', then the compressed version of s will be "a3c3" of length 4.
    fmt.Println(getLengthOfOptimalCompression("aaabcccd", 2)) // 4
    // Example 2:
    // Input: s = "aabbaa", k = 2
    // Output: 2
    // Explanation: If we delete both 'b' characters, the resulting compressed string would be "a4" of length 2.
    fmt.Println(getLengthOfOptimalCompression("aabbaa", 2)) // 2
    // Example 3:
    // Input: s = "aaaaaaaaaaa", k = 0
    // Output: 3
    // Explanation: Since k is zero, we cannot delete anything. The compressed string is "a11" of length 3.
    fmt.Println(getLengthOfOptimalCompression("aaaaaaaaaaa", 0)) // 3

    fmt.Println(getLengthOfOptimalCompression("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 1)) // 3

    fmt.Println(getLengthOfOptimalCompression1("aaabcccd", 2)) // 4
    fmt.Println(getLengthOfOptimalCompression1("aabbaa", 2)) // 2
    fmt.Println(getLengthOfOptimalCompression1("aaaaaaaaaaa", 0)) // 3
    fmt.Println(getLengthOfOptimalCompression1("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 1)) // 3
}