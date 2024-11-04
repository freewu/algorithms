package main

// 3163. String Compression III
// Given a string word, compress it using the following algorithm:
//     1. Begin with an empty string comp. While word is not empty, use the following operation:
//         1.1 Remove a maximum length prefix of word made of a single character c repeating at most 9 times.
//         2.2 Append the length of the prefix followed by c to comp.

// Return the string comp.

// Example 1:
// Input: word = "abcde"
// Output: "1a1b1c1d1e"
// Explanation:
// Initially, comp = "". Apply the operation 5 times, choosing "a", "b", "c", "d", and "e" as the prefix in each operation.
// For each prefix, append "1" followed by the character to comp.

// Example 2:
// Input: word = "aaaaaaaaaaaaaabb"
// Output: "9a5a2b"
// Explanation:
// Initially, comp = "". Apply the operation 3 times, choosing "aaaaaaaaa", "aaaaa", and "bb" as the prefix in each operation.
// For prefix "aaaaaaaaa", append "9" followed by "a" to comp.
// For prefix "aaaaa", append "5" followed by "a" to comp.
// For prefix "bb", append "2" followed by "b" to comp.

// Constraints:
//     1 <= word.length <= 2 * 10^5
//     word consists only of lowercase English letters.

import "fmt"

func compressedString(word string) string {
    res, start, count := []byte{}, word[0], 1
    for i := 1; i < len(word); i++ {
        if count == 9 || word[i] != start {
            res = append(res, byte(count + '0'))
            res = append(res, start)
            count = 1
            start = word[i]
        } else {
            count++
        }
    }
    // 处理结尾
    res = append(res, byte(count + '0'))
    res = append(res, start)
    return string(res)
}

func main() {
    // Example 1:
    // Input: word = "abcde"
    // Output: "1a1b1c1d1e"
    // Explanation:
    // Initially, comp = "". Apply the operation 5 times, choosing "a", "b", "c", "d", and "e" as the prefix in each operation.
    // For each prefix, append "1" followed by the character to comp.
    fmt.Println(compressedString("abcde")) // 1a1b1c1d1e
    // Example 2:
    // Input: word = "aaaaaaaaaaaaaabb"
    // Output: "9a5a2b"
    // Explanation:
    // Initially, comp = "". Apply the operation 3 times, choosing "aaaaaaaaa", "aaaaa", and "bb" as the prefix in each operation.
    // For prefix "aaaaaaaaa", append "9" followed by "a" to comp.
    // For prefix "aaaaa", append "5" followed by "a" to comp.
    // For prefix "bb", append "2" followed by "b" to comp.
    fmt.Println(compressedString("aaaaaaaaaaaaaabb")) // "9a5a2b"
}