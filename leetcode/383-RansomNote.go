package main

// 383. Ransom Note
// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
// Each letter in magazine can only be used once in ransomNote.

// Example 1:
// Input: ransomNote = "a", magazine = "b"
// Output: false

// Example 2:
// Input: ransomNote = "aa", magazine = "ab"
// Output: false

// Example 3:
// Input: ransomNote = "aa", magazine = "aab"
// Output: true

// Constraints:
//     1 <= ransomNote.length, magazine.length <= 10^5
//     ransomNote and magazine consist of lowercase English letters.

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
    if len(ransomNote) > len(magazine) {
        return false
    }
    cnt := make([]int, 26) // ransomNote and magazine consist of lowercase English letters.
    for _, v := range magazine { // 统计出杂志中的字母数
        cnt[v-'a']++
    }
    for _, v := range ransomNote {
        cnt[v-'a']-- // 依次减去赎金信需要用掉的字母
        if cnt[v-'a'] < 0 { // 如果字符不够返回 false
            return false
        }
    }
    return true
}

func main() {
    fmt.Println(canConstruct("a", "b"))      // false
    fmt.Println(canConstruct("aa", "ab"))    // false
    fmt.Println(canConstruct("aa", "aab"))    // true

    fmt.Println(canConstruct("ab", "bac"))   // true
    fmt.Println(canConstruct("cab", "ba"))   // false
    fmt.Println(canConstruct("abc", "abcc")) // true
    fmt.Println(canConstruct("", "abcc"))    // true
    fmt.Println(canConstruct("abc", ""))     // false
}
