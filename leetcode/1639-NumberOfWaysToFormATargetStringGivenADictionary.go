package main

// 1639. Number of Ways to Form a Target String Given a Dictionary
// You are given a list of strings of the same length words and a string target.
// Your task is to form target using the given words under the following rules:
//     target should be formed from left to right.
//     To form the ith character (0-indexed) of target, 
//         you can choose the kth character of the jth string in words if target[i] = words[j][k].
//     Once you use the kth character of the jth string of words, 
//         you can no longer use the xth character of any string in words where x <= k. 
//         In other words, all characters to the left of or at index k become unusuable for every string.
//     Repeat the process until you form the string target.
//     Notice that you can use multiple characters from the same string in words provided the conditions above are met.

// Return the number of ways to form target from words. Since the answer may be too large, return it modulo 10^9 + 7.


// Example 1:
// Input: words = ["acca","bbbb","caca"], target = "aba"
// Output: 6
// Explanation: There are 6 ways to form target.
// "aba" -> index 0 ("acca"), index 1 ("bbbb"), index 3 ("caca")
// "aba" -> index 0 ("acca"), index 2 ("bbbb"), index 3 ("caca")
// "aba" -> index 0 ("acca"), index 1 ("bbbb"), index 3 ("acca")
// "aba" -> index 0 ("acca"), index 2 ("bbbb"), index 3 ("acca")
// "aba" -> index 1 ("caca"), index 2 ("bbbb"), index 3 ("acca")
// "aba" -> index 1 ("caca"), index 2 ("bbbb"), index 3 ("caca")

// Example 2:
// Input: words = ["abba","baab"], target = "bab"
// Output: 4
// Explanation: There are 4 ways to form target.
// "bab" -> index 0 ("baab"), index 1 ("baab"), index 2 ("abba")
// "bab" -> index 0 ("baab"), index 1 ("baab"), index 3 ("baab")
// "bab" -> index 0 ("baab"), index 2 ("baab"), index 3 ("baab")
// "bab" -> index 1 ("abba"), index 2 ("baab"), index 3 ("baab")
 
// Constraints:
//     1 <= words.length <= 1000
//     1 <= words[i].length <= 1000
//     All strings in words have the same length.
//     1 <= target.length <= 1000
//     words[i] and target contain only lowercase English letters.

import "fmt"

func numWays(words []string, target string) int {
    n, m, mod := len(target), len(words[0]), int(1e9 + 7)
    dp, freq := make([][]int, n), make([][26]int, m)
    for _, word := range words {
        for i, ch := range word {
            freq[i][ch-'a']++
        }
    }
    for i := range dp {
        dp[i] = make([]int, m)
    }
    for i := n - 1; i >= 0; i-- {
        for j := m - 1; j >= 0; j-- {
            if i == n-1 {
                if j == m-1 {
                    dp[i][j] = freq[j][target[i]-'a'] 
                } else {
                    dp[i][j] = (dp[i][j+1] + freq[j][target[i]-'a']) % mod 
                }
            } else {
                if j != m-1 {
                    dp[i][j] = (dp[i][j+1] + dp[i+1][j+1]*freq[j][target[i]-'a']) % mod 
                }
            }
        }
    }
    return dp[0][0]
}

func numWays1(words []string, target string) int {
    n, m, mod := len(words[0]), len(target), int(1e9 + 7)
    dp := make([]int, m + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i <= n; i++ {
        pre, cnt := 1, make([]int, 26)
        for _, word := range words {
            cnt[word[i - 1] - 'a']++
        }
        for j := 1; j <= min(m, i); j++ {
            tmp := dp[j]
            dp[j] = (dp[j] + (pre * cnt[target[j - 1] - 'a'])) % mod
            pre = tmp
        }
    }
    return dp[m]
}

func main() {
    // Example 1:
    // Input: words = ["acca","bbbb","caca"], target = "aba"
    // Output: 6
    // Explanation: There are 6 ways to form target.
    // "aba" -> index 0 ("acca"), index 1 ("bbbb"), index 3 ("caca")
    // "aba" -> index 0 ("acca"), index 2 ("bbbb"), index 3 ("caca")
    // "aba" -> index 0 ("acca"), index 1 ("bbbb"), index 3 ("acca")
    // "aba" -> index 0 ("acca"), index 2 ("bbbb"), index 3 ("acca")
    // "aba" -> index 1 ("caca"), index 2 ("bbbb"), index 3 ("acca")
    // "aba" -> index 1 ("caca"), index 2 ("bbbb"), index 3 ("caca")
    fmt.Println(numWays([]string{"acca","bbbb","caca"}, "aba")) // 6
    // Example 2:
    // Input: words = ["abba","baab"], target = "bab"
    // Output: 4
    // Explanation: There are 4 ways to form target.
    // "bab" -> index 0 ("baab"), index 1 ("baab"), index 2 ("abba")
    // "bab" -> index 0 ("baab"), index 1 ("baab"), index 3 ("baab")
    // "bab" -> index 0 ("baab"), index 2 ("baab"), index 3 ("baab")
    // "bab" -> index 1 ("abba"), index 2 ("baab"), index 3 ("baab")
    fmt.Println(numWays([]string{"abba","baab"}, "bab")) // 4

    fmt.Println(numWays1([]string{"acca","bbbb","caca"}, "aba")) // 6
    fmt.Println(numWays1([]string{"abba","baab"}, "bab")) // 4
}