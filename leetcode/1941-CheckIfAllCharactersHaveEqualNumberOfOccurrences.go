package main

// 1941. Check if All Characters Have Equal Number of Occurrences
// Given a string s, return true if s is a good string, or false otherwise.

// A string s is good if all the characters that appear in s have the same number of occurrences (i.e., the same frequency).

// Example 1:
// Input: s = "abacbc"
// Output: true
// Explanation: The characters that appear in s are 'a', 'b', and 'c'. All characters occur 2 times in s.

// Example 2:
// Input: s = "aaabb"
// Output: false
// Explanation: The characters that appear in s are 'a' and 'b'.
// 'a' occurs 3 times while 'b' occurs 2 times, which is not the same number of times.

// Constraints:
//     1 <= s.length <= 1000
//     s consists of lowercase English letters.

import "fmt"

func areOccurrencesEqual(s string) bool {
    mp := make(map[rune]int)
    for _, v:= range s {
        mp[v]++
    }
    count := -1
    for _, v := range mp {
        if count == -1 { count = v }
        if v != count { return false } // 数量不一致
    }
    return true
}

func areOccurrencesEqual1(s string) bool {
    mp := make([]int, 26)
    for _, v := range s {
        mp[int(v - 'a')]++
    }
    count := -1
    for _, v := range mp {
        if v == 0 { continue }
        if count == -1 { count = v }
        if v != count { return false } // 数量不一致
    }
    return true
}

func main() {
    // Example 1:
    // Input: s = "abacbc"
    // Output: true
    // Explanation: The characters that appear in s are 'a', 'b', and 'c'. All characters occur 2 times in s.
    fmt.Println(areOccurrencesEqual("abacbc")) // true
    // Example 2:
    // Input: s = "aaabb"
    // Output: false
    // Explanation: The characters that appear in s are 'a' and 'b'.
    // 'a' occurs 3 times while 'b' occurs 2 times, which is not the same number of times.
    fmt.Println(areOccurrencesEqual("aaabb")) // false

    fmt.Println(areOccurrencesEqual1("abacbc")) // true
    fmt.Println(areOccurrencesEqual1("aaabb")) // false
}