package main

// 3456. Find Special Substring of Length K
// You are given a string s and an integer k.

// Determine if there exists a substring of length exactly k in s that satisfies the following conditions:
//     1. The substring consists of only one distinct character (e.g., "aaa" or "bbb").
//     2. If there is a character immediately before the substring, 
//        it must be different from the character in the substring.
//     3. If there is a character immediately after the substring, 
//        it must also be different from the character in the substring.

// Return true if such a substring exists. Otherwise, return false.

// Example 1:
// Input: s = "aaabaaa", k = 3
// Output: true
// Explanation:
// The substring s[4..6] == "aaa" satisfies the conditions.
// It has a length of 3.
// All characters are the same.
// The character before "aaa" is 'b', which is different from 'a'.
// There is no character after "aaa".

// Example 2:
// Input: s = "abc", k = 2
// Output: false
// Explanation:
// There is no substring of length 2 that consists of one distinct character and satisfies the conditions.

// Constraints:
//     1 <= k <= s.length <= 100
//     s consists of lowercase English letters only.

import "fmt"

func hasSpecialSubstring(s string, k int) bool {
    mp := make(map[byte]int, 26)
    k--
    for i := 0; i < len(s); i++ {
        mp[s[i]]++
        start, end := byte('#'), byte('$')
        if i > k {
            start = s[i-k-1]
            mp[start]--
            if mp[start] == 0 {
                delete(mp, start)
            }
        }
        if i + 1 < len(s) {
            end = s[i+1]
        }
        if i >= k {
            if len(mp) == 1 {
                flag := true
                if start != '#' && start == s[i] {
                    flag = false
                }
                if end != '$' && end == s[i] {
                    flag = false
                }
                if flag {
                    return true
                }
            }
        }
    }
    return false
}

func hasSpecialSubstring1(s string, k int) bool {
    n := len(s)
    if n < k { return false } // 如果字符串长度小于 k，无法形成长度为 k 的子字符串
    for i := 0; i <= n - k; i++ { // 遍历所有可能的子字符串起始位置
        char := s[i]
        // 快速检查：如果当前位置已经是一个连续序列的中间，可以跳过
        // 因为我们已经在之前的迭代中检查过这个序列
        if i > 0 && s[i-1] == char { continue }
        // 检查从i开始的k个字符是否相同
        pos, same := i + k,  true
        for j := i + 1; j < pos; j++ {
            if s[j] != char {
                same = false
                break
            }
        }
        if !same { continue }
        if pos < n && s[pos] == char { continue }// 检查子字符串后面的字符（如果存在）
        return true // 所有条件都满足
    }
    return false
}

func hasSpecialSubstring2(s string, k int) bool {
    arr, count := []byte(s), 1
    for i := 1; i < len(arr); i++ {
        if arr[i] != arr[i-1] {
            if count == k {
                return true
            }
            count = 0
        }
        count++
    }
    return count == k
}

func main() {
    // Example 1:
    // Input: s = "aaabaaa", k = 3
    // Output: true
    // Explanation:
    // The substring s[4..6] == "aaa" satisfies the conditions.
    // It has a length of 3.
    // All characters are the same.
    // The character before "aaa" is 'b', which is different from 'a'.
    // There is no character after "aaa".
    fmt.Println(hasSpecialSubstring("aaabaaa", 3)) // true
    // Example 2:
    // Input: s = "abc", k = 2
    // Output: false
    // Explanation:
    // There is no substring of length 2 that consists of one distinct character and satisfies the conditions.
    fmt.Println(hasSpecialSubstring("abc", 2)) // false

    fmt.Println(hasSpecialSubstring("bluefrog", 2)) // false
    fmt.Println(hasSpecialSubstring("leetcode", 2)) // true

    fmt.Println(hasSpecialSubstring1("aaabaaa", 3)) // true
    fmt.Println(hasSpecialSubstring1("abc", 2)) // false
    fmt.Println(hasSpecialSubstring1("bluefrog", 2)) // false
    fmt.Println(hasSpecialSubstring1("leetcode", 2)) // true

    fmt.Println(hasSpecialSubstring2("aaabaaa", 3)) // true
    fmt.Println(hasSpecialSubstring2("abc", 2)) // false
    fmt.Println(hasSpecialSubstring2("bluefrog", 2)) // false
    fmt.Println(hasSpecialSubstring2("leetcode", 2)) // true
}