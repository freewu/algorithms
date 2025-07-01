package main

// 3330. Find the Original Typed String I
// Alice is attempting to type a specific string on her computer. 
// However, she tends to be clumsy and may press a key for too long, resulting in a character being typed multiple times.

// Although Alice tried to focus on her typing, she is aware that she may still have done this at most once.

// You are given a string word, which represents the final output displayed on Alice's screen.

// Return the total number of possible original strings that Alice might have intended to type.

// Example 1:
// Input: word = "abbcccc"
// Output: 5
// Explanation:
// The possible strings are: "abbcccc", "abbccc", "abbcc", "abbc", and "abcccc".

// Example 2:
// Input: word = "abcd"
// Output: 1
// Explanation:
// The only possible string is "abcd".

// Example 3:
// Input: word = "aaaa"
// Output: 4

// Constraints:
//     1 <= word.length <= 100
//     word consists only of lowercase English letters.

import "fmt"

func possibleStringCount(word string) int {
    n := len(word)
    res := n
    for i := 1; i < n; i++ {
        if word[i] != word[i - 1] { res-- }
    }
    return res
}

func possibleStringCount1(word string) int {
    res, n := 1, len(word)
    for i := 1; i < n; i++ {
        if word[i - 1] == word[i] {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: word = "abbcccc"
    // Output: 5
    // Explanation:
    // The possible strings are: "abbcccc", "abbccc", "abbcc", "abbc", and "abcccc".
    fmt.Println(possibleStringCount("abbcccc")) // 5
    // Example 2:
    // Input: word = "abcd"
    // Output: 1
    // Explanation:
    // The only possible string is "abcd".
    fmt.Println(possibleStringCount("abcd")) // 1
    // Example 3:
    // Input: word = "aaaa"
    // Output: 4
    fmt.Println(possibleStringCount("aaaa")) // 4

    fmt.Println(possibleStringCount("bluefrog")) // 1
    fmt.Println(possibleStringCount("leetcode")) // 2

    fmt.Println(possibleStringCount1("abbcccc")) // 5
    fmt.Println(possibleStringCount1("abcd")) // 1
    fmt.Println(possibleStringCount1("aaaa")) // 4
    fmt.Println(possibleStringCount1("bluefrog")) // 1
    fmt.Println(possibleStringCount1("leetcode")) // 2
}