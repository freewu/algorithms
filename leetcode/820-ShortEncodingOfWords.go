package main

// 820. Short Encoding of Words
// A valid encoding of an array of words is any reference string s and array of indices indices such that:
//     words.length == indices.length
//     The reference string s ends with the '#' character.
//     For each index indices[i], the substring of s starting from indices[i] and up to (but not including) the next '#' character is equal to words[i].

// Given an array of words, return the length of the shortest reference string s possible of any valid encoding of words.

// Example 1:
// Input: words = ["time", "me", "bell"]
// Output: 10
// Explanation: A valid encoding would be s = "time#bell#" and indices = [0, 2, 5].
// words[0] = "time", the substring of s starting from indices[0] = 0 to the next '#' is underlined in "time#bell#"
// words[1] = "me", the substring of s starting from indices[1] = 2 to the next '#' is underlined in "time#bell#"
// words[2] = "bell", the substring of s starting from indices[2] = 5 to the next '#' is underlined in "time#bell#"

// Example 2:
// Input: words = ["t"]
// Output: 2
// Explanation: A valid encoding would be s = "t#" and indices = [0].

// Constraints:
//     1 <= words.length <= 2000
//     1 <= words[i].length <= 7
//     words[i] consists of only lowercase letters.

import "fmt"

func minimumLengthEncoding(words []string) int {
    res, set1, set2 := 0, make(map[string]bool), make(map[string]bool)
    for _, word := range words {
        set1[word], set2[word] = true, true
    }
    for word := range set1 {
        for i := 1; i < len(word); i++ { // 删除子单词
            delete(set2, word[i:])
        }
    }
    for word := range set2 { // 统计剩余的单词 + # 长度
        res += len(word) + 1
    }
    return res
}

func minimumLengthEncoding1(words []string) int {
    res, mp := 0, make(map[string]bool)
    for _, word := range words{
        mp[word] = true
    }
    for word := range mp {
        for i := 1; i < len(word);i++{
            delete(mp, word[i:])
        }
    }
    for word := range mp {
        res += len(word) + 1
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["time", "me", "bell"]
    // Output: 10
    // Explanation: A valid encoding would be s = "time#bell#" and indices = [0, 2, 5].
    // words[0] = "time", the substring of s starting from indices[0] = 0 to the next '#' is underlined in "time#bell#"
    // words[1] = "me", the substring of s starting from indices[1] = 2 to the next '#' is underlined in "time#bell#"
    // words[2] = "bell", the substring of s starting from indices[2] = 5 to the next '#' is underlined in "time#bell#"
    fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"})) // 10
    // Example 2:
    // Input: words = ["t"]
    // Output: 2
    // Explanation: A valid encoding would be s = "t#" and indices = [0].
    fmt.Println(minimumLengthEncoding([]string{"t"})) // t

    fmt.Println(minimumLengthEncoding1([]string{"time", "me", "bell"})) // 10
    fmt.Println(minimumLengthEncoding1([]string{"t"})) // t
}