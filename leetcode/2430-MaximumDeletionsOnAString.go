package main

// 2430. Maximum Deletions on a String
// You are given a string s consisting of only lowercase English letters. 
// In one operation, you can:
//     1. Delete the entire string s, or
//     2. Delete the first i letters of s if the first i letters of s are equal to the following i letters in s, for any i in the range 1 <= i <= s.length / 2.

// For example, if s = "ababc", then in one operation, you could delete the first two letters of s to get "abc", 
// since the first two letters of s and the following two letters of s are both equal to "ab".

// Return the maximum number of operations needed to delete all of s.

// Example 1:
// Input: s = "abcabcdabc"
// Output: 2
// Explanation:
// - Delete the first 3 letters ("abc") since the next 3 letters are equal. Now, s = "abcdabc".
// - Delete all the letters.
// We used 2 operations so return 2. It can be proven that 2 is the maximum number of operations needed.
// Note that in the second operation we cannot delete "abc" again because the next occurrence of "abc" does not happen in the next 3 letters.

// Example 2:
// Input: s = "aaabaab"
// Output: 4
// Explanation:
// - Delete the first letter ("a") since the next letter is equal. Now, s = "aabaab".
// - Delete the first 3 letters ("aab") since the next 3 letters are equal. Now, s = "aab".
// - Delete the first letter ("a") since the next letter is equal. Now, s = "ab".
// - Delete all the letters.
// We used 4 operations so return 4. It can be proven that 4 is the maximum number of operations needed.

// Example 3:
// Input: s = "aaaaa"
// Output: 5
// Explanation: In each operation, we can delete the first letter of s.

// Constraints:
//     1 <= s.length <= 4000
//     s consists only of lowercase English letters.

import "fmt"

func deleteString(s string) int {
    n, mp := len(s), make(map[int]int)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(i int) int
    dfs = func(i int) int {
        if i >= n { return 0 }
        if mp[i] > 0 {  return mp[i] }
        mp[i] = 1 // !!! base condition, 1 deletion for every substring at least
        for j := 1; j <= (n - i) / 2; j++ {
            pattern := s[i:i + j]
            target := s[i + j:i + j + len(pattern)]
            if pattern == target {
                mp[i] = max(mp[i], 1 + dfs(i + j))
            }
        }
        return mp[i]
    }
    return dfs(0)
}

func deleteString1(s string) int {
    n, p := len(s), 13337
    P, L := make([]int, n + 1), make([]int, n + 1)
    P[0], L[0] = 1, 1
    for i := 1; i <= n; i++ {
        P[i] = P[i-1] * p
        L[i] = L[i-1] * p + int(s[i-1]-'a'+1)
    }
    calc := func(f []int, p []int, l, r int) int { return f[r] - f[l-1] * p[r-l+1]; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    f := make([]int, n + 1)
    for i := n; i > 0; i-- {
        f[i] = 1
        for j := 1; j < 1+(n-i+1)/2; j++ {
            if calc(L, P, i, i+j-1) == calc(L, P, i+j, i+j*2-1) {
                f[i] = max(f[i], f[i+j]+1)
            }
        }
    }
    return f[1]
}

func deleteString2(s string) int {
    n := len(s)
    g := make([][]int, n + 1)
    for i := range g {
        g[i] = make([]int, n+1)
    }
    for i := n - 1; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            if s[i] == s[j] {
                g[i][j] = g[i+1][j+1] + 1
            }
        }
    }
    f := make([]int, n)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(int) int
    dfs = func(i int) int {
        if i == n { return 0 }
        if f[i] > 0 { return f[i] }
        f[i] = 1
        for j := 1; j <= (n-i)/2; j++ {
            if g[i][i+j] >= j {
                f[i] = max(f[i], dfs(i+j)+1)
            }
        }
        return f[i]
    }
    return dfs(0)
}

func main() {
    // Example 1:
    // Input: s = "abcabcdabc"
    // Output: 2
    // Explanation:
    // - Delete the first 3 letters ("abc") since the next 3 letters are equal. Now, s = "abcdabc".
    // - Delete all the letters.
    // We used 2 operations so return 2. It can be proven that 2 is the maximum number of operations needed.
    // Note that in the second operation we cannot delete "abc" again because the next occurrence of "abc" does not happen in the next 3 letters.
    fmt.Println(deleteString("abcabcdabc")) // 2
    // Example 2:
    // Input: s = "aaabaab"
    // Output: 4
    // Explanation:
    // - Delete the first letter ("a") since the next letter is equal. Now, s = "aabaab".
    // - Delete the first 3 letters ("aab") since the next 3 letters are equal. Now, s = "aab".
    // - Delete the first letter ("a") since the next letter is equal. Now, s = "ab".
    // - Delete all the letters.
    // We used 4 operations so return 4. It can be proven that 4 is the maximum number of operations needed.
    fmt.Println(deleteString("aaabaab")) // 4
    // Example 3:
    // Input: s = "aaaaa"
    // Output: 5
    // Explanation: In each operation, we can delete the first letter of s.
    fmt.Println(deleteString("aaaaa")) // 5

    fmt.Println(deleteString("bluefrog")) // 1
    fmt.Println(deleteString("leetcode")) // 1

    fmt.Println(deleteString1("abcabcdabc")) // 2
    fmt.Println(deleteString1("aaabaab")) // 4
    fmt.Println(deleteString1("aaaaa")) // 5
    fmt.Println(deleteString1("bluefrog")) // 1
    fmt.Println(deleteString1("leetcode")) // 1

    fmt.Println(deleteString2("abcabcdabc")) // 2
    fmt.Println(deleteString2("aaabaab")) // 4
    fmt.Println(deleteString2("aaaaa")) // 5
    fmt.Println(deleteString2("bluefrog")) // 1
    fmt.Println(deleteString2("leetcode")) // 1
}