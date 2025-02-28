package main

// 1092. Shortest Common Supersequence
// Given two strings str1 and str2, return the shortest string that has both str1 and str2 as subsequences. 
// If there are multiple valid strings, return any of them.

// A string s is a subsequence of string t if deleting some number of characters from t (possibly 0) results in the string s.

// Example 1:
// Input: str1 = "abac", str2 = "cab"
// Output: "cabac"
// Explanation: 
// str1 = "abac" is a subsequence of "cabac" because we can delete the first "c".
// str2 = "cab" is a subsequence of "cabac" because we can delete the last "ac".
// The answer provided is the shortest such string that satisfies these properties.

// Example 2:
// Input: str1 = "aaaaaaaa", str2 = "aaaaaaaa"
// Output: "aaaaaaaa"

// Constraints:
//     1 <= str1.length, str2.length <= 1000
//     str1 and str2 consist of lowercase English letters.

import "fmt"

func shortestCommonSupersequence(s1 string, s2 string) string {
    res, m, n := "", len(s1), len(s2)
    dp := make([][]int,m + 1)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < m + 1; i++ {
        dp[i] = make([]int, n+1)
        for j := 0; j < n+1; j++ {
            if i == 0 || j == 0 {
                dp[i][j] = 0
                continue
            }
            if s1[i-1] == s2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i][j-1], dp[i-1][j])
            }
        }
    }
    if dp[m][n] == 0 {
        return s1+s2
    } else if dp[m][n] == m {
        return s2
    } else if dp[m][n] == n {
        return s1
    }
    ind, tmp := dp[m][n], dp[m][n]
    a, b := make([]int, ind), make([]int, ind)
    i, j := m,n
    for i > 0 && j > 0 {
        if s1[i-1] == s2[j-1] {
            a[tmp-1], b[tmp-1] = i-1, j-1
            i--
            j--
            tmp--
        } else if dp[i-1][j] > dp[i][j-1] {
            i--
        } else {
            j--
        }
    }
    i = 0
    for i < ind {
        if i == 0 {
            res += s1[:a[i]] +  s2[:b[i]] +  string(s1[a[i]])
        } else {
            res += s1[a[i-1]+1:a[i]] + s2[b[i-1]+1:b[i]] + string(s1[a[i]])
        }
        i++
    }
    if a[ind-1] + 1 < m {
        res += s1[a[ind-1]+1:]
    }
    if b[ind-1] + 1 < n {
        res += s2[b[ind-1]+1:]
    }
    return res
}

func shortestCommonSupersequence1(str1 string, str2 string) string {
    // Make sure str1 is the longest
    if len(str1) < len(str2) {
        str1, str2 = str2, str1
        // return shortestCommonSupersequence(str2, str1)
    }
    dp := make([][]int, len(str1) + 1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(str2) + 1)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := len(str1) - 1; i >= 0; i-- {
        for j := len(str2) - 1; j >= 0; j-- {
            if str1[i] == str2[j] {
                dp[i][j] = 1 + dp[i+1][j+1]
            } else {
                dp[i][j] = max(dp[i][j+1], dp[i+1][j])
            }
        }
    }
    res := []byte{} // var res strings.Builder
    i, j := 0, 0
    for i < len(str1) && j < len(str2) {
        if str1[i] == str2[j] {
            // Characters are the same; copy character to result and advance both strings.
            res = append(res, str1[i]) // res.WriteByte(str1[i])
            i++
            j++
        } else if dp[i][j+1] > dp[i+1][j] {
            // Picking str2[j], because the str1[i:] and str2[j+1:] have longest subsequence
            res = append(res, str2[j]) // res.WriteByte(str2[j])
            j++
        } else {
            // if equal or less, we want to advance str1 first as it is bigger
            res = append(res, str1[i]) // res.WriteByte(str1[i])
            i++
        }
    }
    // Simply copy the remaining characters from both strings
    res = append(res, str1[i:]...) // res.WriteString(str1[i:])
    res = append(res, str2[j:]...)  // res.WriteString(str2[j:])
    return string(res) // return res.String()
}

func main() {
    // Example 1:
    // Input: str1 = "abac", str2 = "cab"
    // Output: "cabac"
    // Explanation: 
    // str1 = "abac" is a subsequence of "cabac" because we can delete the first "c".
    // str2 = "cab" is a subsequence of "cabac" because we can delete the last "ac".
    // The answer provided is the shortest such string that satisfies these properties.
    fmt.Println(shortestCommonSupersequence("abac","cab")) // "cabac"
    // Example 2:
    // Input: str1 = "aaaaaaaa", str2 = "aaaaaaaa"
    // Output: "aaaaaaaa"
    fmt.Println(shortestCommonSupersequence("aaaaaaaa","aaaaaaaa")) // "aaaaaaaa"

    fmt.Println(shortestCommonSupersequence("bluefrog","leetcode")) // "bluefretcogde"

    fmt.Println(shortestCommonSupersequence1("abac","cab")) // "cabac"
    fmt.Println(shortestCommonSupersequence1("aaaaaaaa","aaaaaaaa")) // "aaaaaaaa"
    fmt.Println(shortestCommonSupersequence1("bluefrog","leetcode")) // "bluefretcogde"
}