package main

// 1119. Remove Vowels from a String
// Given a string s, remove the vowels 'a', 'e', 'i', 'o', and 'u' from it, and return the new string.

// Example 1:
// Input: s = "leetcodeisacommunityforcoders"
// Output: "ltcdscmmntyfrcdrs"

// Example 2:
// Input: s = "aeiou"
// Output: ""

// Constraints:
//     1 <= s.length <= 1000
//     s consists of only lowercase English letters.

import "fmt"

func removeVowels(s string) string {
    res := []byte{}
    isVowel := func(c byte) bool { return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' }
    for i := 0; i < len(s); i++ {
        if !isVowel(s[i]) { res = append(res, s[i]) }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "leetcodeisacommunityforcoders"
    // Output: "ltcdscmmntyfrcdrs"
    fmt.Println(removeVowels("leetcodeisacommunityforcoders")) // "ltcdscmmntyfrcdrs"
    // Example 2:
    // Input: s = "aeiou"
    // Output: ""
    fmt.Println(removeVowels("aeiou")) // ""
}