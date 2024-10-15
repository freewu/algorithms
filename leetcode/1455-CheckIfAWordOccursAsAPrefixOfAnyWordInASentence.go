package main

// 1455. Check If a Word Occurs As a Prefix of Any Word in a Sentence
// Given a sentence that consists of some words separated by a single space, 
// and a searchWord, check if searchWord is a prefix of any word in sentence.

// Return the index of the word in sentence (1-indexed) where searchWord is a prefix of this word. 
// If searchWord is a prefix of more than one word, return the index of the first word (minimum index). 
// If there is no such word return -1.

// A prefix of a string s is any leading contiguous substring of s.

// Example 1:
// Input: sentence = "i love eating burger", searchWord = "burg"
// Output: 4
// Explanation: "burg" is prefix of "burger" which is the 4th word in the sentence.

// Example 2:
// Input: sentence = "this problem is an easy problem", searchWord = "pro"
// Output: 2
// Explanation: "pro" is prefix of "problem" which is the 2nd and the 6th word in the sentence, but we return 2 as it's the minimal index.

// Example 3:
// Input: sentence = "i am tired", searchWord = "you"
// Output: -1
// Explanation: "you" is not a prefix of any word in the sentence.

// Constraints:
//     1 <= sentence.length <= 100
//     1 <= searchWord.length <= 10
//     sentence consists of lowercase English letters and spaces.
//     searchWord consists of lowercase English letters.

import "fmt"
import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
    for i, v := range strings.Split(sentence, " ") { 
        if strings.HasPrefix(v, searchWord) { 
            return i + 1 
        } 
    } 
    return -1
}

func isPrefixOfWord1(sentence string, searchWord string) int {
    sentence, searchWord = " " + sentence, " " + searchWord
    index := 0
    for i, j := 0, 0; i < len(sentence); i++ {
        if sentence[i] == ' ' { // 多一个单词
            index++ 
        }
        if sentence[i] == searchWord[j] {
            j++
        } else {
            j = 0
        }
        if j == len(searchWord) { 
            return index 
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: sentence = "i love eating burger", searchWord = "burg"
    // Output: 4
    // Explanation: "burg" is prefix of "burger" which is the 4th word in the sentence.
    fmt.Println(isPrefixOfWord("i love eating burger", "burg")) // 4
    // Example 2:
    // Input: sentence = "this problem is an easy problem", searchWord = "pro"
    // Output: 2
    // Explanation: "pro" is prefix of "problem" which is the 2nd and the 6th word in the sentence, but we return 2 as it's the minimal index.
    fmt.Println(isPrefixOfWord("this problem is an easy problem", "pro")) // 2
    // Example 3:
    // Input: sentence = "i am tired", searchWord = "you"
    // Output: -1
    // Explanation: "you" is not a prefix of any word in the sentence.
    fmt.Println(isPrefixOfWord("i am tired", "you")) // -1
    fmt.Println(isPrefixOfWord("b bu bur burg burger", "burg")) // 4

    fmt.Println(isPrefixOfWord1("i love eating burger", "burg")) // 4
    fmt.Println(isPrefixOfWord1("this problem is an easy problem", "pro")) // 2
    fmt.Println(isPrefixOfWord1("i am tired", "you")) // -1
    fmt.Println(isPrefixOfWord1("b bu bur burg burger", "burg")) // 4
}