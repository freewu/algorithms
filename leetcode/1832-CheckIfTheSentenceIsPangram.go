package main

// 1832. Check if the Sentence Is Pangram
// A pangram is a sentence where every letter of the English alphabet appears at least once.

// Given a string sentence containing only lowercase English letters, 
// return true if sentence is a pangram, or false otherwise.

// Example 1:
// Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
// Output: true
// Explanation: sentence contains at least one of every letter of the English alphabet.

// Example 2:
// Input: sentence = "leetcode"
// Output: false

// Constraints:
//     1 <= sentence.length <= 1000
//     sentence consists of lowercase English letters.

import "fmt"

func checkIfPangram(sentence string) bool {
    mp := make([]int, 26)
    for _, v := range sentence {
        mp[int(v - 'a')]++
    }
    for _, v := range mp {
        if v == 0 { return false } // 该字母未出现
    }
    return true
}

func main() {
    // Example 1:
    // Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
    // Output: true
    // Explanation: sentence contains at least one of every letter of the English alphabet.
    fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog")) // true
    // Example 2:
    // Input: sentence = "leetcode"
    // Output: false
    fmt.Println(checkIfPangram("leetcode")) // false
}