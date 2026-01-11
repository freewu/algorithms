package main

// 3805. Count Caesar Cipher Pairs
// You are given an array words of n strings. 
// Each string has length m and contains only lowercase English letters.

// Two strings s and t are similar if we can apply the following operation any number of times (possibly zero times) so that s and t become equal.
//     1. Choose either s or t.
//     2. Replace every letter in the chosen string with the next letter in the alphabet cyclically. The next letter after 'z' is 'a'.

// Count the number of pairs of indices (i, j) such that:
//     1. i < j
//     2. words[i] and words[j] are similar.
//     3. Return an integer denoting the number of such pairs.

// Example 1:
// Input: words = ["fusion","layout"]
// Output: 1
// Explanation:
// words[0] = "fusion" and words[1] = "layout" are similar because we can apply the operation to "fusion" 6 times. 
// The string "fusion" changes as follows.
// "fusion"
// "gvtjpo"
// "hwukqp"
// "ixvlrq"
// "jywmsr"
// "kzxnts"
// "layout"

// Example 2:
// Input: words = ["ab","aa","za","aa"]
// Output: 2
// Explanation:
// words[0] = "ab" and words[2] = "za" are similar. words[1] = "aa" and words[3] = "aa" are similar.

// Constraints:
//     1 <= n == words.length <= 10^5
//     1 <= m == words[i].length <= 10^5
//     1 <= n * m <= 10^5
//     words[i] consists only of lowercase English letters.

import "fmt"
import "strings"

func countPairs(words []string) int64 {
    var sb strings.Builder
    count := make(map[string]int)
    res := 0
    for _, w := range words {
        sb.Reset()
        off := int(w[0] - 'a')
        for i := range len(w) {
            b := byte('a' + (int(w[i] - 'a') + 26 - off) % 26)
            sb.WriteByte(b)
        }
        res += count[sb.String()]
        count[sb.String()]++
    }
    return int64(res)
}

func countPairs1(words []string) int64 {
    res, count := 0, make(map[string]int)
    for _, s := range words {
        t := []byte(s)
        base := t[0]
        for i := range t {
            t[i] = (t[i] - base + 26) % 26 // 保证结果在 [0, 25] 中
        }
        s = string(t)
        res += count[s]
        count[s]++
    }
    return int64(res)   
}

func main() {
    // Example 1:
    // Input: words = ["fusion","layout"]
    // Output: 1
    // Explanation:
    // words[0] = "fusion" and words[1] = "layout" are similar because we can apply the operation to "fusion" 6 times. 
    // The string "fusion" changes as follows.
    // "fusion"
    // "gvtjpo"
    // "hwukqp"
    // "ixvlrq"
    // "jywmsr"
    // "kzxnts"
    // "layout"
    fmt.Println(countPairs([]string{"fusion","layout"})) // 1
    // Example 2:
    // Input: words = ["ab","aa","za","aa"]
    // Output: 2
    // Explanation:
    // words[0] = "ab" and words[2] = "za" are similar. words[1] = "aa" and words[3] = "aa" are similar.
    fmt.Println(countPairs([]string{"ab","aa","za","aa"})) // 2

    fmt.Println(countPairs([]string{"bluefrog","leetcode"})) // 0
    fmt.Println(countPairs([]string{"leetcode","bluefrog"})) // 0

    fmt.Println(countPairs1([]string{"fusion","layout"})) // 1
    fmt.Println(countPairs1([]string{"ab","aa","za","aa"})) // 2
    fmt.Println(countPairs1([]string{"bluefrog","leetcode"})) // 0
    fmt.Println(countPairs1([]string{"leetcode","bluefrog"})) // 0
}