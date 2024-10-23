package main

// 3297. Count Substrings That Can Be Rearranged to Contain a String I
// You are given two strings word1 and word2.

// A string x is called valid if x can be rearranged to have word2 as a prefix.

// Return the total number of valid substrings of word1.

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
//     1 <= word1.length <= 10^5
//     1 <= word2.length <= 10^4
//     word1 and word2 consist only of lowercase English letters.

import "fmt"

// Sliding Window
func validSubstringCount(word1 string, word2 string) int64 {
    mp, count := make([]int, 26), 0 // 出现的字符种类
    for _, ch := range word2 {
        if mp[ch - 'a'] == 0 {
            count++
        }
        mp[ch - 'a']++
    }
    res, j, n := 0, 0, len(word1)
    for i := 0; i < n; i++ {
        k := word1[i] - 'a'
        mp[k]--
        if mp[k] == 0 {
            count--
        }
        for count == 0 {
            res += n - i
            p := word1[j] - 'a'
            mp[p]++
            if mp[p] == 1 {
                count++
            }
            j++
        }
    }
    return int64(res)
}

func validSubstringCount1(word1 string, word2 string) int64 {
    res, left, kind, mp := 0, 0, 0, make([]int, 26)
    for _, v := range word2 {
        mp[v - 'a']++
    }
    for _, v := range mp {
        if v > 0 {
            kind++
        }
    }
    for _, v := range word1 {
        mp[v - 'a']--
        if mp[v - 'a'] == 0 {
            kind--
        }
        for kind == 0 {
            if mp[word1[left]-'a'] == 0 {
                kind++
            }
            mp[word1[left]-'a']++
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