package main

// 3120. Count the Number of Special Characters I
// You are given a string word. A letter is called special if it appears both in lowercase and uppercase in word.

// Return the number of special letters in word.

// Example 1:
// Input: word = "aaAbcBC"
// Output: 3
// Explanation:
// The special characters in word are 'a', 'b', and 'c'.

// Example 2:
// Input: word = "abc"
// Output: 0
// Explanation:
// No character in word appears in uppercase.

// Example 3:
// Input: word = "abBCab"
// Output: 1
// Explanation:
// The only special character in word is 'b'.

// Constraints:
//     1 <= word.length <= 50
//     word consists of only lowercase and uppercase English letters.

import "fmt"

func numberOfSpecialChars(word string) int {
    res, mp := 0, make(map[byte]int)
    for i := range word {
        mp[word[i]]++
    }
    for i := 0; i < 26; i++ {
        if mp[byte(i + 'a')] > 0 && mp[byte(i + 'A')] > 0 {
            res++
        }
    }
    return res
}

func numberOfSpecialChars1(word string) int {
    res, mp := 0, [52]bool{}
    for i := 0; i < len(word); i++ {
        if word[i] >= 'a' && word[i] <= 'z' {
            mp[word[i]-'a'] = true
        }
        if word[i] >= 'A' && word[i] <= 'Z' {
            mp[word[i]-'A' + 26] = true
        }
    }
    for i := 0; i < 26; i++ {
        if mp[i] && mp[i + 26] {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: word = "aaAbcBC"
    // Output: 3
    // Explanation:
    // The special characters in word are 'a', 'b', and 'c'.
    fmt.Println(numberOfSpecialChars("aaAbcBC")) // 3
    // Example 2:
    // Input: word = "abc"
    // Output: 0
    // Explanation:
    // No character in word appears in uppercase.
    fmt.Println(numberOfSpecialChars("abc")) // 0
    // Example 3:
    // Input: word = "abBCab"
    // Output: 1
    // Explanation:
    // The only special character in word is 'b'.
    fmt.Println(numberOfSpecialChars("abBCab")) // 1

    fmt.Println(numberOfSpecialChars1("aaAbcBC")) // 3
    fmt.Println(numberOfSpecialChars1("abc")) // 0
    fmt.Println(numberOfSpecialChars1("abBCab")) // 1
}