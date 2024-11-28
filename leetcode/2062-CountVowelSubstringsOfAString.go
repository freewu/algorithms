package main

// 2062. Count Vowel Substrings of a String
// A substring is a contiguous (non-empty) sequence of characters within a string.

// A vowel substring is a substring that only consists of vowels ('a', 'e', 'i', 'o', and 'u') and has all five vowels present in it.

// Given a string word, return the number of vowel substrings in word.

// Example 1:
// Input: word = "aeiouu"
// Output: 2
// Explanation: The vowel substrings of word are as follows (underlined):
// - "aeiouu"
// - "aeiouu"

// Example 2:
// Input: word = "unicornarihan"
// Output: 0
// Explanation: Not all 5 vowels are present, so there are no vowel substrings.

// Example 3:
// Input: word = "cuaieuouac"
// Output: 7
// Explanation: The vowel substrings of word are as follows (underlined):
// - "cuaieuouac"
// - "cuaieuouac"
// - "cuaieuouac"
// - "cuaieuouac"
// - "cuaieuouac"
// - "cuaieuouac"
// - "cuaieuouac"

// Constraints:
//     1 <= word.length <= 100
//     word consists of lowercase English letters only.

import "fmt"
import "strings"

// sliding window
func countVowelSubstrings(word string) int {
    res := 0
    for i := 0; i < len(word); i++ {
        check := make(map[byte]int)
        for j := i; j < len(word); j++ {
            if strings.Contains("aeiou", string(word[j])) {
                check[word[j]]++
            } else {
                break
            }
            if len(check) == 5 {
                res++
            }
        }
    }
    return res
}

func countVowelSubstrings1(word string) int {
    arr := []string{}
    for i := 0; i < len(word); i++ {
        j := i
        for j < len(word) && strings.ContainsRune("aeiou", rune(word[j])) { j++ }
        if j > i { 
            arr = append(arr, word[i:j]) 
        }
        i = j
    }
    containsAll := func (count [128]int) bool {
        for _, v := range "aeiou" {
            if count[v] == 0 { return false }
        }
        return true
    }
    res := 0
    for _, s := range arr {
        count := [128]int{}
        l, r := 0, 0
        for ; r < len(s); r++ {
            count[s[r]]++
            for l < r && containsAll(count) {
                count[s[l]]--
                l++
            }
            res += l
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: word = "aeiouu"
    // Output: 2
    // Explanation: The vowel substrings of word are as follows (underlined):
    // - "[aeiou]u"
    // - "[aeiouu]"
    fmt.Println(countVowelSubstrings("aeiouu")) // 2
    // Example 2:
    // Input: word = "unicornarihan"
    // Output: 0
    // Explanation: Not all 5 vowels are present, so there are no vowel substrings.
    fmt.Println(countVowelSubstrings("unicornarihan")) // 0
    // Example 3:
    // Input: word = "cuaieuouac"
    // Output: 7
    // Explanation: The vowel substrings of word are as follows (underlined):
    // - "c[uaieuo]uac"
    // - "c[uaieuou]ac"
    // - "c[uaieuoua]c"
    // - "cu[aieuo]uac"
    // - "cu[aieuou]ac"
    // - "cu[aieuoua]c"
    // - "cua[ieuoua]c"
    fmt.Println(countVowelSubstrings("cuaieuouac")) // 0
    
    fmt.Println(countVowelSubstrings1("aeiouu")) // 2
    fmt.Println(countVowelSubstrings1("unicornarihan")) // 0
    fmt.Println(countVowelSubstrings1("cuaieuouac")) // 0
}