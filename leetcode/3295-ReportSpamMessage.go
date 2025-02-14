package main

// 3295. Report Spam Message
// You are given an array of strings message and an array of strings bannedWords.

// An array of words is considered spam if there are at least two words in it that exactly match any word in bannedWords.

// Return true if the array message is spam, and false otherwise.

// Example 1:
// Input: message = ["hello","world","leetcode"], bannedWords = ["world","hello"]
// Output: true
// Explanation:
// The words "hello" and "world" from the message array both appear in the bannedWords array.

// Example 2:
// Input: message = ["hello","programming","fun"], bannedWords = ["world","programming","leetcode"]
// Output: false
// Explanation:
// Only one word from the message array ("programming") appears in the bannedWords array.

// Constraints:
//     1 <= message.length, bannedWords.length <= 10^5
//     1 <= message[i].length, bannedWords[i].length <= 15
//     message[i] and bannedWords[i] consist only of lowercase English letters.

import "fmt"

func reportSpam(message []string, bannedWords []string) bool {
    mp := make(map[string]int)
    for _, word := range message {
        mp[word]++
    }
    count := 0
    for _, word := range bannedWords {
        if mp[word] > 0 {
            count += mp[word]
            mp[word] = 0
        }
    }
    return count >= 2
}

func reportSpam1(message []string, bannedWords []string) bool {
    mp := make(map[string]bool, len(bannedWords))
    for _, word := range bannedWords {
        mp[word] = true
    }
    seen := false
    for _, word := range message {
        if mp[word] {
            if seen { return true } // 出现两次
            seen = true // 出现了一次
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: message = ["hello","world","leetcode"], bannedWords = ["world","hello"]
    // Output: true
    // Explanation:
    // The words "hello" and "world" from the message array both appear in the bannedWords array.
    fmt.Println(reportSpam([]string{"hello","world","leetcode"}, []string{"world","hello"})) // true
    // Example 2:
    // Input: message = ["hello","programming","fun"], bannedWords = ["world","programming","leetcode"]
    // Output: false
    // Explanation:
    // Only one word from the message array ("programming") appears in the bannedWords array.
    fmt.Println(reportSpam([]string{"hello","programming","fun"}, []string{"world","programming","leetcode"})) // false

    fmt.Println(reportSpam([]string{"bluefrog","leetcode"}, []string{"world","programming","leetcode"})) // false

    fmt.Println(reportSpam1([]string{"hello","world","leetcode"}, []string{"world","hello"})) // true
    fmt.Println(reportSpam1([]string{"hello","programming","fun"}, []string{"world","programming","leetcode"})) // false
    fmt.Println(reportSpam1([]string{"bluefrog","leetcode"}, []string{"world","programming","leetcode"})) // false
}