package main

// 2531. Make Number of Distinct Characters Equal
// You are given two 0-indexed strings word1 and word2.

// A move consists of choosing two indices i and j such that 0 <= i < word1.length and 0 <= j < word2.length and swapping word1[i] with word2[j].

// Return true if it is possible to get the number of distinct characters in word1 and word2 to be equal with exactly one move. Return false otherwise.

// Example 1:
// Input: word1 = "ac", word2 = "b"
// Output: false
// Explanation: Any pair of swaps would yield two distinct characters in the first string, and one in the second string.

// Example 2:
// Input: word1 = "abcc", word2 = "aab"
// Output: true
// Explanation: We swap index 2 of the first string with index 0 of the second string. The resulting strings are word1 = "abac" and word2 = "cab", which both have 3 distinct characters.

// Example 3:
// Input: word1 = "abcde", word2 = "fghij"
// Output: true
// Explanation: Both resulting strings will have 5 distinct characters, regardless of which indices we swap.

// Constraints:
//     1 <= word1.length, word2.length <= 10^5
//     word1 and word2 consist of only lowercase English letters.

import "fmt"

func isItPossible(word1 string, word2 string) bool {
    count1, count2 := make([]int, 26), make([]int, 26)
    for _, v := range word1 { count1[v - 'a']++ }
    for _, v := range word2 { count2[v - 'a']++ }
    for i := range count1 {
        for j := range count2 {
            if count1[i] > 0 && count2[j] > 0 {
                count1[i], count2[j] = count1[i] - 1, count2[j] - 1
                count1[j], count2[i] = count1[j] + 1, count2[i] + 1
                dist := 0
                for k, v := range count1 {
                    if v > 0         { dist++ }
                    if count2[k] > 0 { dist-- }
                }
                if dist == 0 { return true }
                count1[i], count2[j] = count1[i] + 1, count2[j] + 1
                count1[j], count2[i] = count1[j] - 1, count2[i] - 1
            }
        }
    }
    return false
}

func main() {
    // Example 1:
    // Input: word1 = "ac", word2 = "b"
    // Output: false
    // Explanation: Any pair of swaps would yield two distinct characters in the first string, and one in the second string.
    fmt.Println(isItPossible("ac", "b")) // false
    // Example 2:
    // Input: word1 = "abcc", word2 = "aab"
    // Output: true
    // Explanation: We swap index 2 of the first string with index 0 of the second string. The resulting strings are word1 = "abac" and word2 = "cab", which both have 3 distinct characters.
    fmt.Println(isItPossible("abcc", "aab")) // true
    // Example 3:
    // Input: word1 = "abcde", word2 = "fghij"
    // Output: true
    // Explanation: Both resulting strings will have 5 distinct characters, regardless of which indices we swap.
    fmt.Println(isItPossible("abcde", "fghij")) // true

    fmt.Println(isItPossible("bluefrog", "leetcode")) // true
}