package main

// 面试题 01.01. Is Unique LCCI
// Implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures?

// Example 1:
// Input: s = "leetcode"
// Output: false

// Example 2:
// Input: s = "abc"
// Output: true

// Note:
//      0 <= len(s) <= 100

import "fmt"


// map
func isUnique(str string) bool {
    m := make(map[byte]int)
    for i := 0; i < len(str); i += 1 {
        m[str[i]] += 1
        // 如果大于 1 说明有重复出现
        if m[str[i]] > 1 {
            return false
        }
    }
    return true
}

// array
func isUnique1(astr string) bool {
    cnt := [26]byte{}
    for _, ch := range astr {
        if cnt[ch - 'a'] == 1 {
            return false
        }
        cnt[ch - 'a'] ++
    }
    return true
}

func main() {
    fmt.Println(isUnique("leetcode")) // false
    fmt.Println(isUnique("abc")) // true

    fmt.Println(isUnique1("leetcode")) // false
    fmt.Println(isUnique1("abc")) // true
}