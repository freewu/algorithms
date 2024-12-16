package main

// 3291. Minimum Number of Valid Strings to Form Target I
// You are given an array of strings words and a string target.

// A string x is called valid if x is a prefix of any string in words.

// Return the minimum number of valid strings that can be concatenated to form target. 
// If it is not possible to form target, return -1.

// Example 1:
// Input: words = ["abc","aaaaa","bcdef"], target = "aabcdabc"
// Output: 3
// Explanation:
// The target string can be formed by concatenating:
// Prefix of length 2 of words[1], i.e. "aa".
// Prefix of length 3 of words[2], i.e. "bcd".
// Prefix of length 3 of words[0], i.e. "abc".

// Example 2:
// Input: words = ["abababab","ab"], target = "ababaababa"
// Output: 2
// Explanation:
// The target string can be formed by concatenating:
// Prefix of length 5 of words[0], i.e. "ababa".
// Prefix of length 5 of words[0], i.e. "ababa".

// Example 3:
// Input: words = ["abcdef"], target = "xyz"
// Output: -1

// Constraints:
//     1 <= words.length <= 100
//     1 <= words[i].length <= 5 * 10^3
//     The input is generated such that sum(words[i].length) <= 10^5.
//     words[i] consists only of lowercase English letters.
//     1 <= target.length <= 5 * 10^3
//     target consists only of lowercase English letters.

import "fmt"

type Trie struct {
    children [26]*Trie
}

func (t *Trie) insert(word string) {
    node := t
    for _, c := range word {
        idx := c - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &Trie{}
        }
        node = node.children[idx]
    }
}

func minValidStrings(words []string, target string) int {
    n, inf  := len(target), 1 << 31
    trie := &Trie{}
    for _, w := range words {
        trie.insert(w)
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    dp := make([]int, n)
    var dfs func(int) int
    dfs = func(i int) int {
        if i >= n { return 0 }
        if dp[i] != 0 { return dp[i] }
        node := trie
        dp[i] = inf
        for j := i; j < n; j++ {
            k := int(target[j] - 'a')
            if node.children[k] == nil { break }
            dp[i] = min(dp[i], 1 + dfs(j + 1))
            node = node.children[k]
        }
        return dp[i]
    }
    if res := dfs(0); res < inf { return res }
    return -1
}

func minValidStrings1(words []string, target string) int {
    helper := func(word, target string) []int {
        s := word + "#" + target
        n := len(s)
        pi := make([]int, n)
        for i := 1; i < n; i++ {
            j := pi[i - 1]
            for j > 0 && s[i] != s[j] {
                j = pi[j - 1]
            }
            if s[i] == s[j] {
                j++
            }
            pi[i] = j
        }
        return pi
    }
    n, inf := len(target), 1 << 31
    back := make([]int, n)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, word := range words {
        pi, m := helper(word, target), len(word)
        for i := 0; i < n; i++ {
            back[i] = max(back[i], pi[m + 1 + i])
        }
    }
    dp := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        dp[i] = inf
    }
    for i := 0; i < n; i++ {
        dp[i + 1] = dp[i + 1 - back[i]] + 1
        if dp[i + 1] > n {
            return -1
        }
    }
    return dp[n]
}

func main() {
    // Example 1:
    // Input: words = ["abc","aaaaa","bcdef"], target = "aabcdabc"
    // Output: 3
    // Explanation:
    // The target string can be formed by concatenating:
    // Prefix of length 2 of words[1], i.e. "aa".
    // Prefix of length 3 of words[2], i.e. "bcd".
    // Prefix of length 3 of words[0], i.e. "abc".
    fmt.Println(minValidStrings([]string{"abc","aaaaa","bcdef"}, "aabcdabc")) // 3
    // Example 2:
    // Input: words = ["abababab","ab"], target = "ababaababa"
    // Output: 2
    // Explanation:
    // The target string can be formed by concatenating:
    // Prefix of length 5 of words[0], i.e. "ababa".
    // Prefix of length 5 of words[0], i.e. "ababa".
    fmt.Println(minValidStrings([]string{"abababab","ab"}, "ababaababa")) // 2
    // Example 3:
    // Input: words = ["abcdef"], target = "xyz"
    // Output: -1
    fmt.Println(minValidStrings([]string{"abcdef"}, "xyz")) // -1

    fmt.Println(minValidStrings1([]string{"abc","aaaaa","bcdef"}, "aabcdabc")) // 3
    fmt.Println(minValidStrings1([]string{"abababab","ab"}, "ababaababa")) // 2
    fmt.Println(minValidStrings1([]string{"abcdef"}, "xyz")) // -1
}