package main

// 3981. Count Distinct Ways to Form Target from Two Strings
// You are given three strings word1, word2, and target.

// Your task is to count the number of ways to form target by choosing characters from word1 and word2 under the following conditions:
//     1. For each character of target, choose one matching character from either word1 or word2.
//     2. The chosen indices from word1 must be strictly increasing.
//     3. The chosen indices from word2 must be strictly increasing.
//     4. At least one character must be chosen from both word1 and word2.

// Two ways are considered different if, for at least one position in target, the chosen character comes from a different string or a different index.

// Return the number of ways. Since the answer may be very large, return it modulo 109 + 7.

// Example 1:
// Input: word1 = "abc", word2 = "bac", target = "abc"
// Output: 5
// Explanation:
// There are 5 ways to form target:
// word1[0] = 'a', word1[1] = 'b', word2[2] = 'c'
// word1[0] = 'a', word2[0] = 'b', word1[2] = 'c'
// word1[0] = 'a', word2[0] = 'b', word2[2] = 'c'
// word2[1] = 'a', word1[1] = 'b', word1[2] = 'c'
// word2[1] = 'a', word1[1] = 'b', word2[2] = 'c'
// All ways preserve the increasing index order inside each string and choose at least one character from each string.

// Example 2:
// Input: word1 = "cd", word2 = "cd", target = "ccd"
// Output: 4
// Explanation:
// There are 4 ways to form target:
// word1[0] = 'c', word2[0] = 'c', word1[1] = 'd'
// word1[0] = 'c', word2[0] = 'c', word2[1] = 'd'
// word2[0] = 'c', word1[0] = 'c', word1[1] = 'd'
// word2[0] = 'c', word1[0] = 'c', word2[1] = 'd'
// The first two 'c' characters in target must come one from each string. The final 'd' can be chosen from either string.

// Example 3:
// Input: word1 = "xy", word2 = "xy", target = "xyxy"
// Output: 2
// Explanation:
// There are 2 ways to form target:
// word1[0] = 'x', word1[1] = 'y', word2[0] = 'x', word2[1] = 'y'
// word2[0] = 'x', word2[1] = 'y', word1[0] = 'x', word1[1] = 'y'
// Each "xy" part in target comes entirely from one string.

// Example 4:
// Input: word1 = "ab", word2 = "cde", target = "ace"
// Output: 1
// Explanation:
// The only way is to choose word1[0] = 'a', word2[0] = 'c', and word2[2] = 'e'. Thus, the answer is 1.

// Constraints:
//     1 <= word1.length, word2.length, target.length <= 100
//     word1, word2, and target consist of lowercase English letters only.

import "fmt"

func interleaveCharacters(word1, word2, target string) int {
    n, m1, m2, mod := len(target), len(word1), len(word2), 1_000_000_007
    memo := make([][][]int, n)
    for i := range memo {
        memo[i] = make([][]int, m1+1)
        for j := range memo[i] {
            memo[i][j] = make([]int, m2+1)
            for p := range memo[i][j] {
                memo[i][j][p] = -1 << 61 
            }
        }
    }
    var dfs func(int, int, int) int
    dfs = func(i, j, k int) int {
        if j < -1 || k < -1 || j+k+2 < i+1 {
            return 0
        }
        if i < 0 {
            return 1
        }
        p := &memo[i][j+1][k+1]
        if *p != -1 << 61 {
            return *p
        }
        // 不选 word1[j] 或 word2[k]
        res := dfs(i, j-1, k) + dfs(i, j, k-1) - dfs(i, j-1, k-1) // 容斥
        if j >= 0 && word1[j] == target[i] {
            // 选 word1[j]
            res += dfs(i-1, j-1, k) - dfs(i-1, j-1, k-1)
        }
        if k >= 0 && word2[k] == target[i] {
            // 选 word2[k]
            res += dfs(i-1, j, k-1) - dfs(i-1, j-1, k-1)
        }
        res %= mod
        *p = res
        return res
    }
    numDistinct := func(s, t string) int { // 不同的子序列
        n, m := len(s), len(t)
        if n < m {
            return 0
        }
        f := make([]int, m+1)
        f[0] = 1
        for i, x := range s {
            for j := min(i, m-1); j >= max(m-n+i, 0); j-- {
                if byte(x) == t[j] {
                    f[j+1] = (f[j+1] + f[j]) % mod
                }
            }
        }
        return f[m]
    }
    res := dfs(n-1, m1-1, m2-1) - numDistinct(word1, target) - numDistinct(word2, target)
    return (res % mod + mod) % mod // 保证 res 非负
}

func main() {
    // Example 1:
    // Input: word1 = "abc", word2 = "bac", target = "abc"
    // Output: 5
    // Explanation:
    // There are 5 ways to form target:
    // word1[0] = 'a', word1[1] = 'b', word2[2] = 'c'
    // word1[0] = 'a', word2[0] = 'b', word1[2] = 'c'
    // word1[0] = 'a', word2[0] = 'b', word2[2] = 'c'
    // word2[1] = 'a', word1[1] = 'b', word1[2] = 'c'
    // word2[1] = 'a', word1[1] = 'b', word2[2] = 'c'
    // All ways preserve the increasing index order inside each string and choose at least one character from each string.
    fmt.Println(interleaveCharacters("abc", "bac", "abc")) // 5 
    // Example 2:
    // Input: word1 = "cd", word2 = "cd", target = "ccd"
    // Output: 4
    // Explanation:
    // There are 4 ways to form target:
    // word1[0] = 'c', word2[0] = 'c', word1[1] = 'd'
    // word1[0] = 'c', word2[0] = 'c', word2[1] = 'd'
    // word2[0] = 'c', word1[0] = 'c', word1[1] = 'd'
    // word2[0] = 'c', word1[0] = 'c', word2[1] = 'd'
    // The first two 'c' characters in target must come one from each string. The final 'd' can be chosen from either string.
    fmt.Println(interleaveCharacters("cd", "cd", "ccd")) // 4
    // Example 3:
    // Input: word1 = "xy", word2 = "xy", target = "xyxy"
    // Output: 2
    // Explanation:
    // There are 2 ways to form target:
    // word1[0] = 'x', word1[1] = 'y', word2[0] = 'x', word2[1] = 'y'
    // word2[0] = 'x', word2[1] = 'y', word1[0] = 'x', word1[1] = 'y'
    // Each "xy" part in target comes entirely from one string.
    fmt.Println(interleaveCharacters("xy", "xy", "xyxy")) // 2
    // Example 4:
    // Input: word1 = "ab", word2 = "cde", target = "ace"
    // Output: 1
    // Explanation:
    // The only way is to choose word1[0] = 'a', word2[0] = 'c', and word2[2] = 'e'. Thus, the answer is 1.
    fmt.Println(interleaveCharacters("ab", "cde", "ace")) // 1

    fmt.Println(interleaveCharacters("bluefrog", "leetcode", "freewu")) // 0
}