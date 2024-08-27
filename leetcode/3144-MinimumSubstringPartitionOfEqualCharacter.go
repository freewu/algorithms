package main

// 3144. Minimum Substring Partition of Equal Character 
// Given a string s, you need to partition it into one or more balanced substrings. 
// For example, if s == "ababcc" then ("abab", "c", "c"), ("ab", "abc", "c"), 
// and ("ababcc") are all valid partitions, but ("a", "bab", "cc"), ("aba", "bc", "c"), and ("ab", "abcc") are not. 
// The unbalanced substrings are bolded.

// Return the minimum number of substrings that you can partition s into.

// Note: A balanced string is a string where each character in the string occurs the same number of times.

// Example 1:
// Input: s = "fabccddg"
// Output: 3
// Explanation:
// We can partition the string s into 3 substrings in one of the following ways: ("fab, "ccdd", "g"), or ("fabc", "cd", "dg").

// Example 2:
// Input: s = "abababaccddb"
// Output: 2
// Explanation:
// We can partition the string s into 2 substrings like so: ("abab", "abaccddb").

// Constraints:
//     1 <= s.length <= 1000
//     s consists only of English lowercase letters.

import "fmt"

// func minimumSubstringsInPartition(s string) int {
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     valid := func(arr []int, mx int) bool {
//         for v := range arr {
//             if v != 0 && v != mx {
//                 return false
//             }
//         }
//         return true
//     }
//     var dfs func(s string) int
//     dfs = func(s string) int {
//         n := len(s)
//         dp := make([]int, n)
//         dp[0] = 1
//         for i := 1; i < n; i++ {
//             arr := make([]int, 26)
//             res, mx := n, 0
//             for j := i; j >= 0; j-- {
//                 arr[s[j] - 'a']++
//                 mx = max(mx, arr[s[j] - 'a'])
//                 if valid(arr, mx) {
//                     if j - 1 < 0 {
//                         res = 1
//                     } else {
//                         res = min(res, 1 + dp[j-1])
//                     }
//                 }
//             }
//             dp[i] = res
//         }
//         return dp[n-1]
//     }
//     return dfs(s)
// }

func minimumSubstringsInPartition(s string) int {
    n, inf := len(s), 1 << 31
    dp := make([]int, n + 1)
    for i := range dp {
        dp[i] = inf
    }
    dp[0] = 0
    for i := 1; i <= n; i++ {
        mx := 0
        mp := make(map[byte]int)
        for j := i; j >= 1; j-- {
            mp[s[j - 1]]++
            if mp[s[j - 1]] > mx {
                mx = mp[s[j-1]]
            }
            if mx * len(mp) == (i - j + 1) && dp[j - 1] != inf {
                if dp[i] > dp[j - 1] + 1 {
                    dp[i] = dp[j - 1] + 1
                }
            }
        }
    }
    return dp[n]
}

func minimumSubstringsInPartition1(s string) int {
    n, inf := len(s), 1 << 31
    memo := make([]int, n)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(int) int
    dfs = func(i int) int {
        if i < 0 {
            return 0
        }
        if memo[i] > 0 { // 之前计算过
            return memo[i]
        }
        k, mx, res, cnt := 0, 0, inf, [26]int{}
        for j := i; j >= 0; j-- {
            b := s[j] - 'a'
            if cnt[b] == 0 {
                k++
            }
            cnt[b]++
            mx = max(mx, cnt[b])
            if i-j+1 == k * mx {
                res = min(res, dfs(j-1)+1)
            }
        }
        memo[i] = res // 记忆化
        return res
    }
    return dfs(n - 1)
}

func main() {
    // Example 1:
    // Input: s = "fabccddg"
    // Output: 3
    // Explanation:
    // We can partition the string s into 3 substrings in one of the following ways: ("fab, "ccdd", "g"), or ("fabc", "cd", "dg").
    fmt.Println(minimumSubstringsInPartition("fabccddg")) // 3
    // Example 2:
    // Input: s = "abababaccddb"
    // Output: 2
    // Explanation:
    // We can partition the string s into 2 substrings like so: ("abab", "abaccddb").
    fmt.Println(minimumSubstringsInPartition("abababaccddb")) // 2

    fmt.Println(minimumSubstringsInPartition1("fabccddg")) // 3
    fmt.Println(minimumSubstringsInPartition1("abababaccddb")) // 2
}