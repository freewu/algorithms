package main 

// 3856. Trim Trailing Vowels
// You are given a string s that consists of lowercase English letters.

// Return the string obtained by removing all trailing vowels from s.

// The vowels consist of the characters 'a', 'e', 'i', 'o', and 'u'.

// Example 1:
// Input: s = "idea"
// Output: "id"
// Explanation:
// Removing "idea", we obtain the string "id".

// Example 2:
// Input: s = "day"
// Output: "day"
// Explanation:
// There are no trailing vowels in the string "day".

// Example 3:
// Input: s = "aeiou"
// Output: ""
// Explanation:
// Removing "aeiou", we obtain the string "".

// Constraints:
//     1 <= s.length <= 100
//     s consists of only lowercase English letters.

import "fmt"

func trimTrailingVowels(s string) string {
    n := len(s)
    i := n - 1
    isVowel := func(c byte) bool { return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' }
    for i >= 0 && isVowel(s[i]) {
        i--
    }
    return s[:i+1]
}

func main() {
    // Example 1:
    // Input: s = "idea"
    // Output: "id"
    // Explanation:
    // Removing "idea", we obtain the string "id".
    fmt.Println(trimTrailingVowels("idea")) // "id"
    // Example 2:
    // Input: s = "day"
    // Output: "day"
    // Explanation:
    // There are no trailing vowels in the string "day".
    fmt.Println(trimTrailingVowels("day")) // "day"
    // Example 3:
    // Input: s = "aeiou"
    // Output: ""
    // Explanation:
    // Removing "aeiou", we obtain the string "".
    fmt.Println(trimTrailingVowels("aeiou")) // ""

    fmt.Println(trimTrailingVowels("bluefrog")) // "bluefrog"
    fmt.Println(trimTrailingVowels("leetcode")) // "leetcod"
    fmt.Println(trimTrailingVowels("freewu")) // "freew"
}