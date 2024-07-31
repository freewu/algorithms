package main

// 2707. Extra Characters in a String
// You are given a 0-indexed string s and a dictionary of words dictionary. 
// You have to break s into one or more non-overlapping substrings such that each substring is present in dictionary. 
// There may be some extra characters in s which are not present in any of the substrings.

// Return the minimum number of extra characters left over if you break up s optimally.

// Example 1:
// Input: s = "leetscode", dictionary = ["leet","code","leetcode"]
// Output: 1
// Explanation: We can break s in two substrings: "leet" from index 0 to 3 and "code" from index 5 to 8. There is only 1 unused character (at index 4), so we return 1.

// Example 2:
// Input: s = "sayhelloworld", dictionary = ["hello","world"]
// Output: 3
// Explanation: We can break s in two substrings: "hello" from index 3 to 7 and "world" from index 8 to 12. The characters at indices 0, 1, 2 are not used in any substring and thus are considered as extra characters. Hence, we return 3.

// Constraints:
//     1 <= s.length <= 50
//     1 <= dictionary.length <= 50
//     1 <= dictionary[i].length <= 50
//     dictionary[i] and s consists of only lowercase English letters
//     dictionary contains distinct words

import "fmt"

// Memoization
func minExtraChar(s string, dictionary []string) int {
    dictSet := make(map[string]bool)
    for _, w := range dictionary {
        dictSet[w] = true
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    memo := make(map[string]int)
    var minExtraCharRecursive func(s string) int
    minExtraCharRecursive = func(s string) int {
        if dictSet[s] { return 0 }
        if v, ok := memo[s]; ok {
            return v
        }
        minChars := len(s)
        for i := range s {
            for j := i; j < len(s); j++ {
                if dictSet[s[i:j+1]] {
                    remaining := minExtraCharRecursive(s[j+1:])
                    minChars = min(minChars, i + remaining)
                }
            }
        }
        memo[s] = minChars
        return minChars
    }
    return minExtraCharRecursive(s)
}

// Tabulation
func minExtraChar1(s string, dictionary []string) int {
    n := len(s)
    dictSet, dp := make(map[string]bool), make([]int, n + 1)
    for _, d := range dictionary {
        dictSet[d] = true
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < n + 1; i++ {
        dp[i] = dp[i-1] + 1
        for j := 0; j < i; j++ {
            slen := i - j
            if dictSet[s[j:i]] {
                slen = 0
            }
            dp[i] = min(dp[i], dp[j] + slen)
        }
    }
    return dp[n]
}

func minExtraChar2(s string, dictionary []string) int {
    n := len(s)
    dp := make([]int, n + 1)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < n; i++ {
        dp[i + 1] = i + 1
        for _, w := range dictionary {
            if i + 1 - len(w) >= 0 && s[i + 1 - len(w) :i + 1] == w {
                dp[i + 1] = min(dp[i + 1], dp[i + 1 - len(w)])
            }
        }
        for j := i - 1; j >= 0; j-- {
            dp[i + 1] = min(dp[i + 1], dp[j + 1] + i - j)
        }
    } 
    return dp[n]
}

func main() {
    // Example 1:
    // Input: s = "leetscode", dictionary = ["leet","code","leetcode"]
    // Output: 1
    // Explanation: We can break s in two substrings: "leet" from index 0 to 3 and "code" from index 5 to 8. There is only 1 unused character (at index 4), so we return 1.
    fmt.Println(minExtraChar("leetscode",[]string{"leet","code","leetcode"})) // 1
    // Example 2:
    // Input: s = "sayhelloworld", dictionary = ["hello","world"]
    // Output: 3
    // Explanation: We can break s in two substrings: "hello" from index 3 to 7 and "world" from index 8 to 12. The characters at indices 0, 1, 2 are not used in any substring and thus are considered as extra characters. Hence, we return 3.
    fmt.Println(minExtraChar("sayhelloworld",[]string{"hello","world"})) // 3

    fmt.Println(minExtraChar1("leetscode",[]string{"leet","code","leetcode"})) // 1
    fmt.Println(minExtraChar1("sayhelloworld",[]string{"hello","world"})) // 3

    fmt.Println(minExtraChar2("leetscode",[]string{"leet","code","leetcode"})) // 1
    fmt.Println(minExtraChar2("sayhelloworld",[]string{"hello","world"})) // 3
}