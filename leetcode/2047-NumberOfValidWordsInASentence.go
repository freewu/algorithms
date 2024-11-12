package main

// 2047. Number of Valid Words in a Sentence
// A sentence consists of lowercase letters ('a' to 'z'), digits ('0' to '9'), hyphens ('-'), punctuation marks ('!', '.', and ','), and spaces (' ') only. 
// Each sentence can be broken down into one or more tokens separated by one or more spaces ' '.

// A token is a valid word if all three of the following are true:
//     1. It only contains lowercase letters, hyphens, and/or punctuation (no digits).
//     2. There is at most one hyphen '-'. 
//        If present, it must be surrounded by lowercase characters ("a-b" is valid, but "-ab" and "ab-" are not valid).
//     3. There is at most one punctuation mark. 
//        If present, it must be at the end of the token ("ab,", "cd!", and "." are valid, but "a!b" and "c.," are not valid).

// Examples of valid words include "a-b.", "afad", "ba-c", "a!", and "!".

// Given a string sentence, return the number of valid words in sentence.

// Example 1:
// Input: sentence = "cat and  dog"
// Output: 3
// Explanation: The valid words in the sentence are "cat", "and", and "dog".

// Example 2:
// Input: sentence = "!this  1-s b8d!"
// Output: 0
// Explanation: There are no valid words in the sentence.
// "!this" is invalid because it starts with a punctuation mark.
// "1-s" and "b8d" are invalid because they contain digits.

// Example 3:
// Input: sentence = "alice and  bob are playing stone-game10"
// Output: 5
// Explanation: The valid words in the sentence are "alice", "and", "bob", "are", and "playing".
// "stone-game10" is invalid because it contains digits.

// Constraints:
//     1 <= sentence.length <= 1000
//     sentence only contains lowercase English letters, digits, ' ', '-', '!', '.', and ','.
//     There will be at least 1 token.

import "fmt"
import "strings"

func countValidWords(sentence string) int {
    res, words := 0, strings.Fields(sentence)
    check := func(s string) bool {
        if s[0] == '-' || s[len(s) - 1] == '-' { return false } // - 开头 or 结尾
        punc := map[rune]bool{ '!': true, '.': true, ',': true, }
        count := 0  
        for k, v := range s {
            if v == '-' {
                count++
                next := rune(s[k + 1])
                if count > 1 || punc[next] { return false } // - 只能出现在字符之间
            }
            if v >= '0' && v <= '9' { return false }
            if punc[v] && k < len(s) - 1 { return false }
        }
        return true
    }
    for _, word := range words {
        if check(word) { res++ }
    }
    return res
}

func main() {
    // Example 1:
    // Input: sentence = "cat and  dog"
    // Output: 3
    // Explanation: The valid words in the sentence are "cat", "and", and "dog".
    fmt.Println(countValidWords("cat and  dog")) // 3
    // Example 2:
    // Input: sentence = "!this  1-s b8d!"
    // Output: 0
    // Explanation: There are no valid words in the sentence.
    // "!this" is invalid because it starts with a punctuation mark.
    // "1-s" and "b8d" are invalid because they contain digits.
    fmt.Println(countValidWords("!this  1-s b8d!")) // 0
    // Example 3:
    // Input: sentence = "alice and  bob are playing stone-game10"
    // Output: 5
    // Explanation: The valid words in the sentence are "alice", "and", "bob", "are", and "playing".
    // "stone-game10" is invalid because it contains digits.
    fmt.Println(countValidWords("alice and  bob are playing stone-game10")) // 5
}