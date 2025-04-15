package main

// 3517. Smallest Palindromic Rearrangement I
// You are given a palindromic string s.

// Return the lexicographically smallest palindromic permutation of s.

// Example 1:
// Input: s = "z"
// Output: "z"
// Explanation:
// A string of only one character is already the lexicographically smallest palindrome.

// Example 2:
// Input: s = "babab"
// Output: "abbba"
// Explanation:
// Rearranging "babab" → "abbba" gives the smallest lexicographic palindrome.

// Example 3:
// Input: s = "daccad"
// Output: "acddca"
// Explanation:
// Rearranging "daccad" → "acddca" gives the smallest lexicographic palindrome.

// Constraints:
//     1 <= s.length <= 10^5
//     s consists of lowercase English letters.
//     s is guaranteed to be palindromic.

import "fmt"

func smallestPalindrome(s string) string {
    n := len(s)
    freq := make([]int, 26)
    for _, v := range s {
        freq[v- 'a']++
    }
    oddCount, oddChar := 0, byte('a')
    for i, c := range freq { // 得到单字符
        if c % 2 != 0 {
            oddCount++
            oddChar = byte(i + 'a')
        }
    }
    if oddCount > 1 { return "" } // 回文中单字符只能有一个(在中间)
    res := make([]byte, n)
    l, r := 0, n - 1
    for i := 0; i < 26; i++ {
        if freq[i] == 0 { continue }
        c := byte(i + 'a')
        for freq[i] >= 2 {
            res[l], res[r] = c, c
            l++
            r--
            freq[i] -= 2
        }
    }
    if oddCount == 1 {
        res[l] = oddChar
    }
    return string(res)
}

func smallestPalindrome1(s string) string {
    // 1. 统计字符出现次数
    freq := make([]int, 26)
    for _, ch := range s {
        freq[ch-'a']++
    }
    // 2. 构建左半部分和中间字符
    left := make([]byte, 0, len(s)/2)
    var mid byte
    for i := 0; i < 26; i++ {
        // 将出现次数的一半放到左半部分
        for j := 0; j < freq[i]/2; j++ {
            left = append(left, byte('a'+i))
        }
        // 如果是奇数频次，将它作为中间字符（仅可能有一个）
        if freq[i]%2 == 1 {
            mid = byte('a' + i)
        }
    }
    // 3. 构建右半部分（为左半部分的逆序）
    right := make([]byte, len(left))
    for i := 0; i < len(left); i++ {
        right[i] = left[len(left)-1-i]
    }
    // 4. 组合结果
    if mid != 0 {
        return string(left) + string(mid) + string(right)
    }
    return string(left) + string(right)
}

func main() {
    // Example 1:
    // Input: s = "z"
    // Output: "z"
    // Explanation:
    // A string of only one character is already the lexicographically smallest palindrome.
    fmt.Println(smallestPalindrome("z")) // z
    // Example 2:
    // Input: s = "babab"
    // Output: "abbba"
    // Explanation:
    // Rearranging "babab" → "abbba" gives the smallest lexicographic palindrome.
    fmt.Println(smallestPalindrome("babab")) // abbba
    // Example 3:
    // Input: s = "daccad"
    // Output: "acddca"
    // Explanation:
    // Rearranging "daccad" → "acddca" gives the smallest lexicographic palindrome.
    fmt.Println(smallestPalindrome("daccad")) // acddca

    fmt.Println(smallestPalindrome("bluefrogbluefrog")) // befgloruurolgfeb
    fmt.Println(smallestPalindrome("leetcodeleetcode")) // cdeeelottoleeedc

    fmt.Println(smallestPalindrome1("z")) // z
    fmt.Println(smallestPalindrome1("babab")) // abbba
    fmt.Println(smallestPalindrome1("daccad")) // acddca
    fmt.Println(smallestPalindrome1("bluefrogbluefrog")) // befgloruurolgfeb
    fmt.Println(smallestPalindrome1("leetcodeleetcode")) // cdeeelottoleeedc
}