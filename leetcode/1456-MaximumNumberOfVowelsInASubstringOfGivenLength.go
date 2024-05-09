package main

// 1456. Maximum Number of Vowels in a Substring of Given Length
// Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.
// Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

// Example 1:
// Input: s = "abciiidef", k = 3
// Output: 3
// Explanation: The substring "iii" contains 3 vowel letters.

// Example 2:
// Input: s = "aeiou", k = 2
// Output: 2
// Explanation: Any substring of length 2 contains 2 vowels.

// Example 3:
// Input: s = "leetcode", k = 3
// Output: 2
// Explanation: "lee", "eet" and "ode" contain 2 vowels.
 
// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase English letters.
//     1 <= k <= s.length

import "fmt"

// 超出时间限制 102 / 106 
func maxVowels1(s string, k int) int {
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    isVowel := func(ch byte) bool {
        return 'a' == ch || 'e' == ch || 'i' == ch || 'o' == ch || 'u' == ch 
    }
    for i := 0; i < len(s) - k + 1; i++ {
        j, t := 0, 0
        for j < k {
            if isVowel(s[i + j]) {
                t++
            }
            j++
        }
        res = max(res,t)
    }
    return res
}

func maxVowels(s string, k int) int {
    count, vowels := 0, map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
    for i := 0; i < k; i++ { // 得到第一个循环的元音数量
        if vowels[rune(s[i])] {
            count++
        }
    }
    res := count
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := k; i < len(s); i++ {
        if vowels[rune(s[i])] { // 检查新纳入到窗口的字符为元音则 + 1
            count++
        }
        if vowels[rune(s[i-k])] { // 检查要移出窗口的字符为元音则 - 1
            count--
        }
        res = max(res, count)
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abciiidef", k = 3
    // Output: 3
    // Explanation: The substring "iii" contains 3 vowel letters.
    fmt.Println(maxVowels("abciiidef",3)) // 3
    // Example 2:
    // Input: s = "aeiou", k = 2
    // Output: 2
    // Explanation: Any substring of length 2 contains 2 vowels.
    fmt.Println(maxVowels("aeiou",2)) // 2
    // Example 3:
    // Input: s = "leetcode", k = 3
    // Output: 2
    // Explanation: "lee", "eet" and "ode" contain 2 vowels.
    fmt.Println(maxVowels("leetcode", 3)) // 2

    fmt.Println(maxVowels1("abciiidef",3)) // 3
    fmt.Println(maxVowels1("aeiou",2)) // 2
    fmt.Println(maxVowels1("leetcode", 3)) // 2
}