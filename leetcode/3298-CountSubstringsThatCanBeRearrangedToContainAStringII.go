package main

// 3298. Count Substrings That Can Be Rearranged to Contain a String II
// You are given two strings word1 and word2.

// A string x is called valid if x can be rearranged to have word2 as a prefix.

// Return the total number of valid substrings of word1.

// Note that the memory limits in this problem are smaller than usual, so you must implement a solution with a linear runtime complexity.

// Example 1:
// Input: word1 = "bcca", word2 = "abc"
// Output: 1
// Explanation:
// The only valid substring is "bcca" which can be rearranged to "abcc" having "abc" as a prefix.

// Example 2:
// Input: word1 = "abcabc", word2 = "abc"
// Output: 10
// Explanation:
// All the substrings except substrings of size 1 and size 2 are valid.

// Example 3:
// Input: word1 = "abcabc", word2 = "aaabc"
// Output: 0

// Constraints:
//     1 <= word1.length <= 10^6
//     1 <= word2.length <= 10^4
//     word1 and word2 consist only of lowercase English letters.

import "fmt"

func validSubstringCount(word1 string, word2 string) int64 {
    mp1, mp2 := make([]int, 26), make([]int, 26)
    for _, v := range word2 {
        mp2[v-'a']++
    }
    res, left, n1, k := 0, 0, len(word1), len(word2)
    for right := 0; right < n1; right++ {
        if mp2[word1[right]-'a'] > 0 {
            if mp1[word1[right]-'a'] < mp2[word1[right]-'a'] {
                k--
            }
        }
        mp1[word1[right]-'a']++
        for k == 0 {
            res += n1 - right
            mp1[word1[left]-'a']--
            if mp2[word1[left]-'a'] > 0 && mp1[word1[left]-'a'] < mp2[word1[left]-'a'] {
                k++
            }
            left++
        }
    }
    return int64(res)
}

func validSubstringCount1(s, t string) int64 {
    if len(s) < len(t) { return 0 }
    mp := [26]int{} // t 的字母出现次数与 s 的字母出现次数之差
    for _, v := range t {
        mp[v - 'a']++
    }
    kind := 0 // 统计窗口内有多少种字母出现
    for _, v := range mp {
        if v > 0 {
            kind++
        }
    }
    res, left := 0, 0
    for _, v := range s {
        mp[v - 'a']--
        if mp[v-'a'] == 0 { // 窗口内 v 的出现次数和 t 一样
            kind--
        }
        for kind == 0 { // 窗口符合要求
            if mp[s[left]-'a'] == 0 {
                // s[left] 移出窗口之前，检查出现次数，
                // 如果窗口内 s[left] 的出现次数和 t 一样，
                // 那么 s[left] 移出窗口后，窗口内 s[left] 的出现次数比 t 的少
                kind++
            }
            mp[s[left]-'a']++
            left++
        }
        res += left
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: word1 = "bcca", word2 = "abc"
    // Output: 1
    // Explanation:
    // The only valid substring is "bcca" which can be rearranged to "abcc" having "abc" as a prefix.
    fmt.Println(validSubstringCount("bcca", "abc")) // 1
    // Example 2:
    // Input: word1 = "abcabc", word2 = "abc"
    // Output: 10
    // Explanation:
    // All the substrings except substrings of size 1 and size 2 are valid.
    fmt.Println(validSubstringCount("abcabc", "abc")) // 10
    // Example 3:
    // Input: word1 = "abcabc", word2 = "aaabc"
    // Output: 0
    fmt.Println(validSubstringCount("abcabc", "aaabc")) // 0

    fmt.Println(validSubstringCount1("bcca", "abc")) // 1
    fmt.Println(validSubstringCount1("abcabc", "abc")) // 10
    fmt.Println(validSubstringCount1("abcabc", "aaabc")) // 0
}