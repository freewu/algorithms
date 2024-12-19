package main

// 2223. Sum of Scores of Built Strings
// You are building a string s of length n one character at a time, prepending each new character to the front of the string. 
// The strings are labeled from 1 to n, where the string with length i is labeled si.
//     For example, for s = "abaca", s1 == "a", s2 == "ca", s3 == "aca", etc.

// The score of si is the length of the longest common prefix between si and sn (Note that s == sn).

// Given the final string s, return the sum of the score of every si.

// Example 1:
// Input: s = "babab"
// Output: 9
// Explanation:
// For s1 == "b", the longest common prefix is "b" which has a score of 1.
// For s2 == "ab", there is no common prefix so the score is 0.
// For s3 == "bab", the longest common prefix is "bab" which has a score of 3.
// For s4 == "abab", there is no common prefix so the score is 0.
// For s5 == "babab", the longest common prefix is "babab" which has a score of 5.
// The sum of the scores is 1 + 0 + 3 + 0 + 5 = 9, so we return 9.

// Example 2:
// Input: s = "azbazbzaz"
// Output: 14
// Explanation: 
// For s2 == "az", the longest common prefix is "az" which has a score of 2.
// For s6 == "azbzaz", the longest common prefix is "azb" which has a score of 3.
// For s9 == "azbazbzaz", the longest common prefix is "azbazbzaz" which has a score of 9.
// For all other si, the score is 0.
// The sum of the scores is 2 + 3 + 9 = 14, so we return 14.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase English letters.

import "fmt"

func sumScores(s string) int64 {
    res, n, count := 0, len(s), make(map[int]int)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i, l, r := 1, 0, 0; i < n; i++ {
        if i <= r {
            count[i] = min((r - i + 1), count[i - l])
        }
        for i + count[i] < n && s[count[i]] == s[i + count[i]] {
            count[i]++
        }
        if i + count[i] - 1 > r {
            l, r = i, i + count[i] - 1
        }
    }
    for _, v := range count {
        res += v
    }
    return int64(n + res)
}

func sumScores1(s string) int64 {
    sum, l, r, n := 0, 0, 0, len(s)
    count := make([]int, n)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < n; i ++ {
        if i < r {
            count[i] = min(r - i, count[i - l])
        }
        for i + count[i] < n && s[i + count[i]] == s[count[i]] {
            count[i] ++
        }
        if i + count[i] > r {
            r, l = i + count[i], i
        }
        sum += count[i]
    }
    return int64(sum + n)
}

func main() {
    // Example 1:
    // Input: s = "babab"
    // Output: 9
    // Explanation:
    // For s1 == "b", the longest common prefix is "b" which has a score of 1.
    // For s2 == "ab", there is no common prefix so the score is 0.
    // For s3 == "bab", the longest common prefix is "bab" which has a score of 3.
    // For s4 == "abab", there is no common prefix so the score is 0.
    // For s5 == "babab", the longest common prefix is "babab" which has a score of 5.
    // The sum of the scores is 1 + 0 + 3 + 0 + 5 = 9, so we return 9.
    fmt.Println(sumScores("babab")) // 9
    // Example 2:
    // Input: s = "azbazbzaz"
    // Output: 14
    // Explanation: 
    // For s2 == "az", the longest common prefix is "az" which has a score of 2.
    // For s6 == "azbzaz", the longest common prefix is "azb" which has a score of 3.
    // For s9 == "azbazbzaz", the longest common prefix is "azbazbzaz" which has a score of 9.
    // For all other si, the score is 0.
    // The sum of the scores is 2 + 3 + 9 = 14, so we return 14.
    fmt.Println(sumScores("azbazbzaz")) // 14

    fmt.Println(sumScores1("babab")) // 9
    fmt.Println(sumScores1("azbazbzaz")) // 14
}