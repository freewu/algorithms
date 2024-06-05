package main

// 1668. Maximum Repeating Substring
// For a string sequence, a string word is k-repeating if word concatenated k times is a substring of sequence. 
// The word's maximum k-repeating value is the highest value k where word is k-repeating in sequence. 
// If word is not a substring of sequence, word's maximum k-repeating value is 0.

// Given strings sequence and word, return the maximum k-repeating value of word in sequence.

// Example 1:
// Input: sequence = "ababc", word = "ab"
// Output: 2
// Explanation: "abab" is a substring in "ababc".

// Example 2:
// Input: sequence = "ababc", word = "ba"
// Output: 1
// Explanation: "ba" is a substring in "ababc". "baba" is not a substring in "ababc".

// Example 3:
// Input: sequence = "ababc", word = "ac"
// Output: 0
// Explanation: "ac" is not a substring in "ababc". 
 
// Constraints:
//     1 <= sequence.length <= 100
//     1 <= word.length <= 100
//     sequence and word contains only lowercase English letters

import "fmt"

func maxRepeating(sequence string, word string) int {
    res, count, n := 0, 0, len(word)
    for i := 0; i <= len(sequence) - n; i++ {
        if sequence[i:i+n] == word {
            i = i + n - 1
            count++
            if res < count {
                res = count
            }
            continue
        }
        i -= count * len(word)
        count = 0
    }
    return res
}

func main() {
    // Example 1:
    // Input: sequence = "ababc", word = "ab"
    // Output: 2
    // Explanation: "abab" is a substring in "ababc".
    fmt.Println(maxRepeating("ababc", "ab")) // 2
    // Example 2:
    // Input: sequence = "ababc", word = "ba"
    // Output: 1
    // Explanation: "ba" is a substring in "ababc". "baba" is not a substring in "ababc".
    fmt.Println(maxRepeating("ababc", "ba")) // 1
    // Example 3:
    // Input: sequence = "ababc", word = "ac"
    // Output: 0
    // Explanation: "ac" is not a substring in "ababc". 
    fmt.Println(maxRepeating("ababc", "ac")) // 0
}