package main

// 916. Word Subsets
// You are given two string arrays words1 and words2.

// A string b is a subset of string a if every letter in b occurs in a including multiplicity.
//     For example, "wrr" is a subset of "warrior" but is not a subset of "world".

// A string a from words1 is universal if for every string b in words2, b is a subset of a.

// Return an array of all the universal strings in words1. You may return the answer in any order.

// Example 1:
// Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","o"]
// Output: ["facebook","google","leetcode"]

// Example 2:
// Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["l","e"]
// Output: ["apple","google","leetcode"]

// Constraints:
//     1 <= words1.length, words2.length <= 10^4
//     1 <= words1[i].length, words2[i].length <= 10
//     words1[i] and words2[i] consist only of lowercase English letters.
//     All the strings of words1 are unique.

import "fmt"

func wordSubsets(word1[]string, word2 []string) []string {
    m := [26]byte{}
    getChars := func(s string) [26]byte {
        r := [26]byte{}
        for i := 0; i < len(s); i++ {
            r[s[i]-'a']++
        }
        return r
    }
    merge := func(a *[26]byte, b [26]byte) {
        for i := 0; i < 26; i++ {
            if b[i] > a[i] {
                a[i] = b[i]
            }
        }
    }
    contains := func(a, b [26]byte) bool {
        for i := 0; i < 26; i++ {
            if a[i] < b[i] {
                return false
            }
        }
        return true
    }
    for _, s := range word2 {
        ch := getChars(s)
        merge(&m, ch)
    }
    res := []string{}
    for _, s := range word1 {
        ch := getChars(s)
        if contains(ch, m) {
            res = append(res, s)
        }
    }
    return res
}

func wordSubsets1(words1 []string, words2 []string) []string {
    calc := func (word string, a []int) {
        for i := range a {
            a[i] = 0
        }
        for _, c := range word {
            a[c-'a']++
        }
    }
    a1, a2 := make([]int, 26), make([]int, 26)
    for _, word := range words2 {
        calc(word, a2)
        for i := range a1 {
            a1[i] = max(a1[i], a2[i])
        }
    }
    res := []string{}
    for _, word := range words1 {
        calc(word, a2)
        flag := true
        for i := range a1 {
            if a1[i] > a2[i] {
                flag = false
                break
            }
        }
        if flag {
            res = append(res, word)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","o"]
    // Output: ["facebook","google","leetcode"]
    fmt.Println(wordSubsets([]string{"amazon","apple","facebook","google","leetcode"}, []string{"e","o"})) // ["facebook","google","leetcode"]
    // Example 2:
    // Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["l","e"]
    // Output: ["apple","google","leetcode"]
    fmt.Println(wordSubsets([]string{"amazon","apple","facebook","google","leetcode"}, []string{"l","e"})) // "apple","google","leetcode"
    
    fmt.Println(wordSubsets1([]string{"amazon","apple","facebook","google","leetcode"}, []string{"e","o"})) // ["facebook","google","leetcode"]
    fmt.Println(wordSubsets1([]string{"amazon","apple","facebook","google","leetcode"}, []string{"l","e"})) // "apple","google","leetcode"
}