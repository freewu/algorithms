package main

// 1897. Redistribute Characters to Make All Strings Equal
// You are given an array of strings words (0-indexed).

// In one operation, pick two distinct indices i and j, where words[i] is a non-empty string, and move any character from words[i] to any position in words[j].

// Return true if you can make every string in words equal using any number of operations, and false otherwise.

// Example 1:
// Input: words = ["abc","aabc","bc"]
// Output: true
// Explanation: Move the first 'a' in words[1] to the front of words[2],
// to make words[1] = "abc" and words[2] = "abc".
// All the strings are now equal to "abc", so return true.

// Example 2:
// Input: words = ["ab","a"]
// Output: false
// Explanation: It is impossible to make all the strings equal using the operation.

// Constraints:
//     1 <= words.length <= 100
//     1 <= words[i].length <= 100
//     words[i] consists of lowercase English letters.

import "fmt"

func makeEqual(words []string) bool {
    mp := make([]int, 26)
    for _, word := range words {// 统计字符出现频率
        for i := 0; i < len(word); i++ {
            mp[int(word[i] - 'a')]++
        }
    }
    for i := 0; i < 26; i++ {
        if mp[i] % len(words) != 0 { return false } // 不能被单词数整除
    }
    return true
}

func main() {
    // Example 1:
    // Input: words = ["abc","aabc","bc"]
    // Output: true
    // Explanation: Move the first 'a' in words[1] to the front of words[2],
    // to make words[1] = "abc" and words[2] = "abc".
    // All the strings are now equal to "abc", so return true.
    fmt.Println(makeEqual([]string{"abc","aabc","bc"})) // true
    // Example 2:
    // Input: words = ["ab","a"]
    // Output: false
    // Explanation: It is impossible to make all the strings equal using the operation.
    fmt.Println(makeEqual([]string{"ab","a"})) // false
}