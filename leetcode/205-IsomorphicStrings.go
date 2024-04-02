package main

// 205. Isomorphic Strings
// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving the order of characters. 
// No two characters may map to the same character, but a character may map to itself.

// Example 1:
// Input: s = "egg", t = "add"
// Output: true

// Example 2:
// Input: s = "foo", t = "bar"
// Output: false

// Example 3:
// Input: s = "paper", t = "title"
// Output: true
 
// Constraints:
//     1 <= s.length <= 5 * 10^4
//     t.length == s.length
//     s and t consist of any valid ascii character.

import "fmt"

// 暴力破解法 O(n^2)
func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    for i := 0; i < len(s); i++ {
        for j := i + 1; j < len(s); j++ {
            // 找到 s 相同的字符的位置 再判断 t 相同位置是否一致
            if (s[i] == s[j] && t[i] != t[j]) || (t[i] == t[j] && s[i] != s[j]) {
                return false
            }
        }
    }
    return true
}

// O(n)
func isIsomorphic1(s string, t string) bool {
    m1, m2 := make(map[rune]int), make(map[rune]int)
    code_s := make([]int, len(s))
    start := 0
    for i, c := range s {
        // 第一次出现的词累加编号
        if _, ok := m1[c]; !ok {
            m1[c] = start
            start++
        }
        // 使用 code_s 来保存 s 的构词模式
        code_s[i] = m1[c]
    }
    start = 0
    for i, c := range t {
        if _, ok := m2[c]; !ok {
            m2[c] = start
            start++
        }
        if m2[c] != code_s[i] {
            return false
        }
    }
    return true
}

func isIsomorphic2(s string, t string) bool {
    s2t := [128]byte{}
    t2s := [128]byte{}
    for i := 0; i < len(s); i++ {
        if s2t[s[i]] != t2s[t[i]] {
            return false
        }
        s2t[s[i]], t2s[t[i]] = byte(i+1), byte(i+1)
    }
    return true
}

func main() {
    fmt.Println(isIsomorphic("egg","add")) // true
    fmt.Println(isIsomorphic("foo","bar")) // false
    fmt.Println(isIsomorphic("paper","title")) // true
    fmt.Println(isIsomorphic("tomato","potato")) // false
    fmt.Println(isIsomorphic("babc","baba")) // false

    fmt.Println(isIsomorphic1("egg","add")) // true
    fmt.Println(isIsomorphic1("foo","bar")) // false
    fmt.Println(isIsomorphic1("paper","title")) // true
    fmt.Println(isIsomorphic1("tomato","potato")) // false
    fmt.Println(isIsomorphic1("babc","baba")) // false

    fmt.Println(isIsomorphic1("egg","add")) // true
    fmt.Println(isIsomorphic1("foo","bar")) // false
    fmt.Println(isIsomorphic1("paper","title")) // true
    fmt.Println(isIsomorphic1("tomato","potato")) // false
    fmt.Println(isIsomorphic1("babc","baba")) // false
}