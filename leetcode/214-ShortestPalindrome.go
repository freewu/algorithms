package main

// 214. Shortest Palindrome
// You are given a string s. You can convert s to a palindrome by adding characters in front of it.
// Return the shortest palindrome you can find by performing this transformation.

// Example 1:
// Input: s = "aacecaaa"
// Output: "aaacecaaa"

// Example 2:
// Input: s = "abcd"
// Output: "dcbabcd"

// Constraints:
//     0 <= s.length <= 5 * 10^4
//     s consists of lowercase English letters only.

import "fmt"
import "strings"

func shortestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    i, j, r := 0, len(s)-1, len(s)-1
    reverse := func(s string) string {
        res := ""
        for _,ch := range s {
            res = string(ch) + res
        }
        return res
    }
    for j > 0  {
        i, j = 0, r
        for i < j {
            if s[i]== s[j] {
                i++
                j--
            } else {
               break 
            }
        }
        if i >=j  {
            return reverse(s[r+1:]) + s
        }
        r--
    }
    return reverse(s[1:]) + s
}

func shortestPalindrome1(s string) string {
    binary, mod := 26, 1638477777777
    mul, hash, reverseHash, index, maxIndex := 1, 0, 0, 0, 0
    for index < len(s) {
        hash = (hash * binary + int(s[index])) % mod
        reverseHash = (reverseHash + int(s[index]) * mul) % mod
        if hash == reverseHash {
            maxIndex = index
        }
        index ++
        mul = (mul * binary) % mod
    }
    var sb strings.Builder
    for i := len(s) - 1; i >= maxIndex + 1; i--{
        sb.WriteByte(s[i])
    }
    sb.WriteString(s)
    return sb.String()
}

func main() {
    // Example 1:
    // Input: s = "aacecaaa"
    // Output: "aaacecaaa"
    fmt.Println(shortestPalindrome("aaacecaaa")) // aaacecaaa
    // Example 2:
    // Input: s = "abcd"
    // Output: "dcbabcd"
    fmt.Println(shortestPalindrome("abcd")) // dcbabcd

    fmt.Println(shortestPalindrome1("aaacecaaa")) // aaacecaaa
    fmt.Println(shortestPalindrome1("abcd")) // dcbabcd
}