package main

// 472. Concatenated Words
// Given an array of strings words (without duplicates), return all the concatenated words in the given list of words.
// A concatenated word is defined as a string that is comprised entirely of at least two shorter words (not necessarily distinct) in the given array.

// Example 1:
// Input: words = ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]
// Output: ["catsdogcats","dogcatsdog","ratcatdogcat"]
// Explanation: "catsdogcats" can be concatenated by "cats", "dog" and "cats"; 
// "dogcatsdog" can be concatenated by "dog", "cats" and "dog"; 
// "ratcatdogcat" can be concatenated by "rat", "cat", "dog" and "cat".

// Example 2:
// Input: words = ["cat","dog","catdog"]
// Output: ["catdog"]

// Constraints:
//     1 <= words.length <= 10^4
//     1 <= words[i].length <= 30
//     words[i] consists of only lowercase English letters.
//     All the strings of words are unique.
//     1 <= sum(words[i].length) <= 10^5

import "fmt"

// dp
func findAllConcatenatedWordsInADict(words []string) []string {
    res, dict := []string{}, make(map[string]bool)
    for _, word := range words {
        dict[word] = true
    }
    for _, word := range words {
        wl := len(word)
        dp := make([]bool, wl + 1)
        dp[0] = true
        for end := 1; end <= wl; end++ {
            start := 0
            if end == wl {
                start = 1 // Don't search entire word, it will be found
            }
            for ; start < end && !dp[end]; start++ { // Break when we find dp[end] true, word found
                if dp[start] && dict[word[start:end]] {
                    dp[end] = true
                }
            }
            if dp[wl] {
                res = append(res, word)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]
    // Output: ["catsdogcats","dogcatsdog","ratcatdogcat"]
    // Explanation: "catsdogcats" can be concatenated by "cats", "dog" and "cats"; 
    // "dogcatsdog" can be concatenated by "dog", "cats" and "dog"; 
    // "ratcatdogcat" can be concatenated by "rat", "cat", "dog" and "cat".
    fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"})) // ["catsdogcats","dogcatsdog","ratcatdogcat"]
    // Example 2:
    // Input: words = ["cat","dog","catdog"]
    // Output: ["catdog"]
    fmt.Println(findAllConcatenatedWordsInADict([]string{"cat","dog","catdog"})) // ["catdog"]
}