package main

// 2490. Circular Sentence
// A sentence is a list of words that are separated by a single space with no leading or trailing spaces.
//     For example, "Hello World", "HELLO", "hello world hello world" are all sentences.

// Words consist of only uppercase and lowercase English letters. 
// Uppercase and lowercase English letters are considered different.

// A sentence is circular if:
//     The last character of a word is equal to the first character of the next word.
//     The last character of the last word is equal to the first character of the first word.

// For example, "leetcode exercises sound delightful", "eetcode", "leetcode eats soul" are all circular sentences.
// However, "Leetcode is cool", "happy Leetcode", "Leetcode" and "I like Leetcode" are not circular sentences.

// Given a string sentence, return true if it is circular. Otherwise, return false.

// Example 1:
// Input: sentence = "leetcode exercises sound delightful"
// Output: true
// Explanation: The words in sentence are ["leetcode", "exercises", "sound", "delightful"].
// - leetcode's last character is equal to exercises's first character.
// - exercises's last character is equal to sound's first character.
// - sound's last character is equal to delightful's first character.
// - delightful's last character is equal to leetcode's first character.
// The sentence is circular.

// Example 2:
// Input: sentence = "eetcode"
// Output: true
// Explanation: The words in sentence are ["eetcode"].
// - eetcode's last character is equal to eetcode's first character.
// The sentence is circular.

// Example 3:
// Input: sentence = "Leetcode is cool"
// Output: false
// Explanation: The words in sentence are ["Leetcode", "is", "cool"].
// - Leetcode's last character is not equal to is's first character.
// The sentence is not circular.

// Constraints:
//     1 <= sentence.length <= 500
//     sentence consist of only lowercase and uppercase English letters and spaces.
//     The words in sentence are separated by a single space.
//     There are no leading or trailing spaces.

import "fmt"

func isCircularSentence(sentence string) bool {
    n, prev := len(sentence), sentence[0]
    if prev != sentence[n-1] {  return false } // 判断首尾字符是否一致
    for i := 0; i < n; {
        if sentence[i] != ' ' { // 如果下一个是 ' ', prev 会保留 单词的最后一个字符  
            prev = sentence[i]
            i++
            continue
        }
        for sentence[i] == ' ' { i++ } // 如果是 ' ' 要向前走一步 i++ 
        if sentence[i] != prev { // 判读单词开头和上一个单词是否一致
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: sentence = "leetcode exercises sound delightful"
    // Output: true
    // Explanation: The words in sentence are ["leetcode", "exercises", "sound", "delightful"].
    // - leetcode's last character is equal to exercises's first character.
    // - exercises's last character is equal to sound's first character.
    // - sound's last character is equal to delightful's first character.
    // - delightful's last character is equal to leetcode's first character.
    // The sentence is circular.
    fmt.Println(isCircularSentence("leetcode exercises sound delightful")) // true
    // Example 2:
    // Input: sentence = "eetcode"
    // Output: true
    // Explanation: The words in sentence are ["eetcode"].
    // - eetcode's last character is equal to eetcode's first character.
    // The sentence is circular.
    fmt.Println(isCircularSentence("eetcode")) // true
    // Example 3:
    // Input: sentence = "Leetcode is cool"
    // Output: false
    // Explanation: The words in sentence are ["Leetcode", "is", "cool"].
    // - Leetcode's last character is not equal to is's first character.
    // The sentence is not circular.
    fmt.Println(isCircularSentence("Leetcode is cool")) // false
}