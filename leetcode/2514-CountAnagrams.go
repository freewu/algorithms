package main

// 2514. Count Anagrams
// You are given a string s containing one or more words. 
// Every consecutive pair of words is separated by a single space ' '.

// A string t is an anagram of string s if the ith word of t is a permutation of the ith word of s.
//     For example, "acb dfe" is an anagram of "abc def", but "def cab" and "adc bef" are not.

// Return the number of distinct anagrams of s. 
// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: s = "too hot"
// Output: 18
// Explanation: Some of the anagrams of the given string are "too hot", "oot hot", "oto toh", "too toh", and "too oht".

// Example 2:
// Input: s = "aa"
// Output: 1
// Explanation: There is only one anagram possible for the given string.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase English letters and spaces ' '.
//     There is single space between consecutive words.

import "fmt"
import "strings"

func countAnagrams(s string) int {
    res, n, mod := 1, len(s), 1_000_000_007
    facts := make([]int, n + 1)
    facts[0] = 1
    for i := 1; i <= n; i++ {
        facts[i] = (facts[i - 1] * i) % mod
    }
    var pow func(v, p int) int
    pow = func(v, p int) int {
        if p == 1 { return v }
        if p % 2 == 0 {
            t := pow(v, p / 2)
            return (t * t) % mod
        } else {
            return (pow(v, p - 1) * v) % mod
        }
    }
    countAnagrams := func(word string) int {
        count := [26]int{}
        for _, v := range word {
            count[v - 'a']++
        }
        n := len(word)
        res := facts[n]
        for _, v := range count {
            res = (res * pow(facts[v], mod - 2)) % mod
        }
        return res
    }
    for _, word := range strings.Fields(s) {
        res = (res * countAnagrams(word)) % mod
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "too hot"
    // Output: 18
    // Explanation: Some of the anagrams of the given string are "too hot", "oot hot", "oto toh", "too toh", and "too oht".
    fmt.Println(countAnagrams("too hot")) // 18
    // Example 2:
    // Input: s = "aa"
    // Output: 1
    // Explanation: There is only one anagram possible for the given string.
    fmt.Println(countAnagrams("aa")) // 1

    fmt.Println(countAnagrams("bluefrog")) // 40320
    fmt.Println(countAnagrams("leetcode")) // 6720
    fmt.Println(countAnagrams("leetcode bluefrog")) // 270950400
    fmt.Println(countAnagrams("leetcodebluefrog")) // 945726481
    fmt.Println(countAnagrams("bluefrog leetcode")) // 270950400
    fmt.Println(countAnagrams("bluefrogleetcode")) // 945726481
}