package main

// 3563. Lexicographically Smallest String After Adjacent Removals
// You are given a string s consisting of lowercase English letters.

// You can perform the following operation any number of times (including zero):
//     1. Remove any pair of adjacent characters in the string that are consecutive in the alphabet, in either order (e.g., 'a' and 'b', or 'b' and 'a').
//     2. Shift the remaining characters to the left to fill the gap.

// Return the lexicographically smallest string that can be obtained after performing the operations optimally.

// Note: Consider the alphabet as circular, thus 'a' and 'z' are consecutive.

// Example 1:
// Input: s = "abc"
// Output: "a"
// Explanation:
// Remove "bc" from the string, leaving "a" as the remaining string.
// No further operations are possible. Thus, the lexicographically smallest string after all possible removals is "a".

// Example 2:
// Input: s = "bcda"
// Output: ""
// Explanation:
// ​​​​​​​Remove "cd" from the string, leaving "ba" as the remaining string.
// Remove "ba" from the string, leaving "" as the remaining string.
// No further operations are possible. Thus, the lexicographically smallest string after all possible removals is "".

// Example 3:
// Input: s = "zdce"
// Output: "zdce"
// Explanation:
// Remove "dc" from the string, leaving "ze" as the remaining string.
// No further operations are possible on "ze".
// However, since "zdce" is lexicographically smaller than "ze", the smallest string after all possible removals is "zdce".

// Constraints:
//     1 <= s.length <= 250
//     s consists only of lowercase English letters.

import "fmt"

func lexicographicallySmallestString(s string) string {
    dp := make([][]string, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = make([]string, len(s))
        for j := 0; j < len(s); j++ {
            dp[i][j] = "-"
        }
    }
    valid := func(x, y byte) bool {
        a, b := int(x), int(y)
        return a - b == 1 || a - b == -1 || a - b == 25 || a - b == -25
    } 
    min := func (x, y string) string { if x < y { return x; }; return y; }
    var dfs func(prev, index int) string
    dfs = func(prev, index int) (res string) {
        if prev > index { return "" }
        if index - prev == 0 { return string(s[index]) }
        if !(len(dp[prev][index]) == 1 && dp[prev][index][0] == '-') { return dp[prev][index] }
        res = s[prev:index + 1]
        for i := prev; i <= index; i++ {
            if valid(s[i], s[prev]) {
                if dfs(prev+1, i-1) == "" {
                     res = min(res, dfs(i + 1, index))   
                }
            }
        }
        res = min(res, s[prev:prev + 1] + dfs(prev + 1, index))
        dp[prev][index] = res
        return res
    }
    return dfs(0, len(s) - 1)
}

func lexicographicallySmallestString1(s string) string {
    n := len(s)
    can := make([][]bool, n)
    for i := range can {
        can[i] = make([]bool, n)
    }
    isConsecutive := func(a, b byte) bool {
        diff := int(a) - int(b)
        if diff < 0 { diff = -diff }
        return diff == 1 || diff == 25
    }
    for i := 1; i <= n; i++ {
        for l := 0; l + i - 1 < n; l++ {
            r := l + i - 1
            if i & 1 == 1 { continue }
            for m := l + 1; m <= r; m++ {
                if isConsecutive(s[l], s[m]) {
                    if ((m == l+1) || can[l+1][m-1]) && ((m == r) || can[m+1][r]) {
                        can[l][r] = true
                        break
                    }
                }
            }
        }
    }
    dp := make([]string, n + 1)
    dp[n] = ""
    for i := n - 1; i >= 0; i-- {
        dp[i] = string(s[i]) + dp[i+1]
        for j := i + 1; j < n; j++ {
            if isConsecutive(s[i], s[j]) && ((j == i+1) || can[i+1][j-1]) {
                if dp[j+1] < dp[i] {
                    dp[i] = dp[j+1]
                }
            }
        }
    }
    return dp[0]
}

func lexicographicallySmallestString2(s string) (ans string) {
    n := len(s)
    canBeEmpty := make([][]bool, n)
    for i := range canBeEmpty {
        canBeEmpty[i] = make([]bool, n)
    }
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    isConsecutive := func(x, y byte) bool {
        diff := abs(int(x) - int(y))
        return diff == 1 || diff == 25
    }
    for i := n - 2; i >= 0; i-- {
        canBeEmpty[i+1][i] = true // 空串
        for j := i + 1; j < n; j += 2 {
            // 性质 2
            if isConsecutive(s[i], s[j]) && canBeEmpty[i+1][j-1] {
                canBeEmpty[i][j] = true
                continue
            }
            // 性质 3
            for k := i + 1; k < j-1; k += 2 {
                if canBeEmpty[i][k] && canBeEmpty[k+1][j] {
                    canBeEmpty[i][j] = true
                    break
                }
            }
        }
    }
    f := make([]string, n + 1)
    for i := n - 1; i >= 0; i-- {
        // 包含 s[i]
        res := string(s[i]) + f[i+1]
        // 不包含 s[i]，注意 s[i] 不能单独消除，必须和其他字符一起消除
        for j := i + 1; j < n; j += 2 {
            if canBeEmpty[i][j] { // 消除 s[i] 到 s[j]
                res = min(res, f[j+1])
            }
        }
        f[i] = res
    }
    return f[0]
}

func main() {
    // Example 1:
    // Input: s = "abc"
    // Output: "a"
    // Explanation:
    // Remove "bc" from the string, leaving "a" as the remaining string.
    // No further operations are possible. Thus, the lexicographically smallest string after all possible removals is "a".
    fmt.Println(lexicographicallySmallestString("abc")) // "a"
    // Example 2:
    // Input: s = "bcda"
    // Output: ""
    // Explanation:
    // ​​​​​​​Remove "cd" from the string, leaving "ba" as the remaining string.
    // Remove "ba" from the string, leaving "" as the remaining string.
    // No further operations are possible. Thus, the lexicographically smallest string after all possible removals is "".
    fmt.Println(lexicographicallySmallestString("bcda")) // ""
    // Example 3:
    // Input: s = "zdce"
    // Output: "zdce"
    // Explanation:
    // Remove "dc" from the string, leaving "ze" as the remaining string.
    // No further operations are possible on "ze".
    // However, since "zdce" is lexicographically smaller than "ze", the smallest string after all possible removals is "zdce".
    fmt.Println(lexicographicallySmallestString("zdce")) // "zdce"

    fmt.Println(lexicographicallySmallestString("bluefrog")) // "bluefrog"
    fmt.Println(lexicographicallySmallestString("leetcode")) // "leetco"

    fmt.Println(lexicographicallySmallestString1("abc")) // "a"
    fmt.Println(lexicographicallySmallestString1("bcda")) // ""
    fmt.Println(lexicographicallySmallestString1("zdce")) // "zdce"
    fmt.Println(lexicographicallySmallestString1("bluefrog")) // "bluefrog"
    fmt.Println(lexicographicallySmallestString1("leetcode")) // "leetco"

    fmt.Println(lexicographicallySmallestString2("abc")) // "a"
    fmt.Println(lexicographicallySmallestString2("bcda")) // ""
    fmt.Println(lexicographicallySmallestString2("zdce")) // "zdce"
    fmt.Println(lexicographicallySmallestString2("bluefrog")) // "bluefrog"
    fmt.Println(lexicographicallySmallestString2("leetcode")) // "leetco"
}