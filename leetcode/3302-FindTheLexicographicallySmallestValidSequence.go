package main

// 3302. Find the Lexicographically Smallest Valid Sequence
// You are given two strings word1 and word2.

// A string x is called almost equal to y if you can change at most one character in x to make it identical to y.

// A sequence of indices seq is called valid if:
//     The indices are sorted in ascending order.
//     Concatenating the characters at these indices in word1 in the same order results in a string that is almost equal to word2.

// Return an array of size word2.length representing the lexicographically smallest valid sequence of indices. 
// If no such sequence of indices exists, return an empty array.

// Note that the answer must represent the lexicographically smallest array, 
// not the corresponding string formed by those indices.

// Example 1:
// Input: word1 = "vbcca", word2 = "abc"
// Output: [0,1,2]
// Explanation:
// The lexicographically smallest valid sequence of indices is [0, 1, 2]:
// Change word1[0] to 'a'.
// word1[1] is already 'b'.
// word1[2] is already 'c'.

// Example 2:
// Input: word1 = "bacdc", word2 = "abc"
// Output: [1,2,4]
// Explanation:
// The lexicographically smallest valid sequence of indices is [1, 2, 4]:
// word1[1] is already 'a'.
// Change word1[2] to 'b'.
// word1[4] is already 'c'.

// Example 3:
// Input: word1 = "aaaaaa", word2 = "aaabc"
// Output: []
// Explanation:
// There is no valid sequence of indices.

// Example 4:
// Input: word1 = "abc", word2 = "ab"
// Output: [0,1]

// Constraints:
//     1 <= word2.length < word1.length <= 3 * 10^5
//     word1 and word2 consist only of lowercase English letters.

import "fmt"

func validSequence(word1 string, word2 string) []int {
    n1, n2 := len(word1), len(word2)
    pref := make([]int, n1)
    // right to left
    j := n2 - 1
    for i := n1 - 1; i >= 0; i-- {
        if i < n1-1 {
            pref[i] = pref[i+1]
        }
        if j >= 0 && word1[i] == word2[j] {
            pref[i]++
            j--
        }
    }
    // left to right
    res := make([]int, n2)
    match, i, j := 0, 0, 0
    for i < n1 && j < n2 {
        if word1[i] == word2[j] {
            res[j] = i
            j++
            match++
        } else if i < n1-1 && pref[i+1] >= n2-match-1 {
            res[j] = i
            j++
            i++
            for j < n2 && i < n1 {
                if word1[i] == word2[j] {
                    res[j] = i
                    j++
                }
                i++
            }
            return res
        }
        i++
    }
    if match == n2 { return res }
    return []int{}
}

func validSequence1(word1 string, word2 string) []int {
    n1, n2 := len(word1), len(word2)
    dp := make([]int, n1 + 1) // word1 匹配word2的后缀长度
    for i, j := n1 - 1, n2 - 1; i >= 0; i-- {
        if j >= 0 && word1[i] == word2[j] {
            j--
        }
        dp[i] = n2 - j - 1
    }
    res := make([]int, n2)
    i, j, change := 0, 0, false
    for i < n1 && j < n2 {
        if word1[i] == word2[j] { // 匹配的就尽早添加
            res[j] = i
            j++
        } else {
            if !change && (i == n1 - 1 || j + 1 + dp[i + 1] >= n2) { // 当前字符不匹配时，如何判断这一位能不能替换：
                // 看这一位后面是否能包含word2剩余的全部字符（因为这一位替换后，后面就不能再有替换了，必须完全匹配）。
                // 预先计算word1的每个位置的后缀包含的word2后缀长度就可以方便完成这个判断。
                res[j] = i
                j++
                change = true
            }
        }
        i++
    }
    if j < n2 { return []int{} }
    return res
}

func main() {
    // Example 1:
    // Input: word1 = "vbcca", word2 = "abc"
    // Output: [0,1,2]
    // Explanation:
    // The lexicographically smallest valid sequence of indices is [0, 1, 2]:
    // Change word1[0] to 'a'.
    // word1[1] is already 'b'.
    // word1[2] is already 'c'.
    fmt.Println(validSequence("vbcca", "abc")) // [0,1,2]
    // Example 2:
    // Input: word1 = "bacdc", word2 = "abc"
    // Output: [1,2,4]
    // Explanation:
    // The lexicographically smallest valid sequence of indices is [1, 2, 4]:
    // word1[1] is already 'a'.
    // Change word1[2] to 'b'.
    // word1[4] is already 'c'.
    fmt.Println(validSequence("bacdc", "abc")) // [1,2,4]
    // Example 3:
    // Input: word1 = "aaaaaa", word2 = "aaabc"
    // Output: []
    // Explanation:
    // There is no valid sequence of indices.
    fmt.Println(validSequence("aaaaaa", "aaabc")) // []
    // Example 4:
    // Input: word1 = "abc", word2 = "ab"
    // Output: [0,1]
    fmt.Println(validSequence("abc", "ab")) // [0,1]

    fmt.Println(validSequence1("vbcca", "abc")) // [0,1,2]
    fmt.Println(validSequence1("bacdc", "abc")) // [1,2,4]
    fmt.Println(validSequence1("aaaaaa", "aaabc")) // []
    fmt.Println(validSequence1("abc", "ab")) // [0,1]
}