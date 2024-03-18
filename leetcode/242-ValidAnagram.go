package main

// 242. Valid Anagram
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:
// Input: s = "anagram", t = "nagaram"
// Output: true

// Example 2:
// Input: s = "rat", t = "car"
// Output: false
 
// Constraints:
//     1 <= s.length, t.length <= 5 * 10^4
//     s and t consist of lowercase English letters.
 
// Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

import "fmt"

// 支持 utf8
func isAnagram(s string, t string) bool {
    if len(t) != len(s) {
        return false
    }
    // 逐字写到 map中
    ms := make(map[rune]int)
    mt := make(map[rune]int)
    for _, v := range s {
        ms[v]++
    }
    for _, v := range t {
        mt[v]++
    }
    // 数量出现不一,返回 false
    for k,v := range ms {
        if v != mt[k] {
            return false
        }
    }
    return true
}

// best solution
func isAnagram1(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    cnt := make([]int, 26)
    for i := range s {
        cnt[s[i] - 'a'] ++
        cnt[t[i] - 'a'] --
    }
    for i := range cnt {
        if cnt[i] != 0 {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println(isAnagram("anagram","nagaram")) // true
    fmt.Println(isAnagram("rat","car")) // false

    fmt.Println(isAnagram1("anagram","nagaram")) // true
    fmt.Println(isAnagram1("rat","car")) // false
}