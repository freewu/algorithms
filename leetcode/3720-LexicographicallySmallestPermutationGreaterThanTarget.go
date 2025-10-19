package main

// 3720. Lexicographically Smallest Permutation Greater Than Target
// You are given two strings s and target, both having length n, consisting of lowercase English letters.

// Return the lexicographically smallest permutation of s that is strictly greater than target. 
// If no permutation of s is lexicographically strictly greater than target, return an empty string.

// A string a is lexicographically strictly greater than a string b (of the same length) if in the first position where a and b differ, 
// string a has a letter that appears later in the alphabet than the corresponding letter in b.

// A permutation is a rearrangement of all the characters of a string.

// Example 1:
// Input: s = "abc", target = "bba"
// Output: "bca"
// Explanation:
// The permutations of s (in lexicographical order) are "abc", "acb", "bac", "bca", "cab", and "cba".
// The lexicographically smallest permutation that is strictly greater than target is "bca".

// Example 2:
// Input: s = "leet", target = "code"
// Output: "eelt"
// Explanation:
// The permutations of s (in lexicographical order) are "eelt", "eetl", "elet", "elte", "etel", "etle", "leet", "lete", "ltee", "teel", "tele", and "tlee".
// The lexicographically smallest permutation that is strictly greater than target is "eelt".

// Example 3:
// Input: s = "baba", target = "bbaa"
// Output: ""
// Explanation:
// The permutations of s (in lexicographical order) are "aabb", "abab", "abba", "baab", "baba", and "bbaa".
// None of them is lexicographically strictly greater than target. Therefore, the answer is "".
 
// Constraints:
//     1 <= s.length == target.length <= 300
//     s and target consist of only lowercase English letters.

import "fmt"
import "strings"

func lexGreaterPermutation(s, target string) string {
    res, left := []byte(target),make([]int, 26)
    for i, b := range s {
        left[b-'a']++
        left[target[i]-'a']-- // 消耗 s 中的一个字母 target[i]
    }
    for i := len(s) - 1; i >= 0; i-- {
        flag := false
        b := target[i] - 'a'
        left[b]++ // 撤销消耗
        for _, c := range left {
            if c < 0 { // [0,i-1] 无法做到全部一样
                flag = true
                break
            }
        }
        if flag { continue }
        // target[i] 增大到 j
        for j := b + 1; j < 26; j++ {
            if left[j] == 0 { continue }
            left[j]--
            res[i] = 'a' + j
            res = res[:i+1]
            for k, c := range left {
                ch := string('a' + byte(k))
                res = append(res, strings.Repeat(ch, c)...)
            }
            return string(res)
        }
        // 增大失败，继续枚举
    }
    return ""
}

func main() {
    // Example 1:
    // Input: s = "abc", target = "bba"
    // Output: "bca"
    // Explanation:
    // The permutations of s (in lexicographical order) are "abc", "acb", "bac", "bca", "cab", and "cba".
    // The lexicographically smallest permutation that is strictly greater than target is "bca".
    fmt.Println(lexGreaterPermutation("abc", "bba")) // bca
    // Example 2:
    // Input: s = "leet", target = "code"
    // Output: "eelt"
    // Explanation:
    // The permutations of s (in lexicographical order) are "eelt", "eetl", "elet", "elte", "etel", "etle", "leet", "lete", "ltee", "teel", "tele", and "tlee".
    // The lexicographically smallest permutation that is strictly greater than target is "eelt".
    fmt.Println(lexGreaterPermutation("leet", "code")) // eelt  
    // Example 3:
    // Input: s = "baba", target = "bbaa"
    // Output: ""
    // Explanation:
    // The permutations of s (in lexicographical order) are "aabb", "abab", "abba", "baab", "baba", and "bbaa".
    // None of them is lexicographically strictly greater than target. Therefore, the answer is "".
    fmt.Println(lexGreaterPermutation("baba", "bbaa")) // ""

    fmt.Println(lexGreaterPermutation("bluefrog", "leetcode")) // lefbgoru
    fmt.Println(lexGreaterPermutation("leetcode", "bluefrog")) // cdeeelot
}