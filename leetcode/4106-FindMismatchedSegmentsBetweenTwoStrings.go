package main

// 4106. Find Mismatched Segments Between Two Strings
// You are given two strings s1 and s2 of the same length.

// A mismatched segment is a maximal contiguous range of indices where the characters of s1 and s2 differ at every index.

// Return a 2D array containing [sub1, sub2] for each mismatched segment, where sub1 and sub2 are the corresponding substrings from s1 and s2, respectively. 
// Return the segments in increasing order of their starting indices. 
// If there are no mismatched segments, return an empty array.

// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Input: s1 = "earth", s2 = "empty"
// Output: [["ar","mp"],["h","y"]]
// Explanation:
//     The characters at indices 1 and 2 differ, forming the mismatched segment ["ar","mp"].
//     The characters at index 4 differ, forming the mismatched segment ["h","y"].

// Example 2:
// Input: s1 = "action", s2 = "nation"
// Output: [["ac","na"]]
// Explanation:
// The characters at indices 0 and 1 differ, forming the mismatched segment ["ac","na"]. 
// All remaining characters are equal in both strings.

// Example 3:
// Input: s1 = "hot", s2 = "hot"
// Output: []
// Explanation:
// The two strings are identical, so there are no mismatched segments.

// Constraints:
//     1 <= s1.length == s2.length <= 100
//     s1 and s2 consist of lowercase English letters.

import "fmt"

func findMismatchedSegments(s1, s2 string) [][]string {
    res := [][]string{}
    start, n := 0, len(s1)
    for i := range n {
        if s1[i] == s2[i] {
            start = i + 1 // 记录段的起点
        } else if i == n-1 || s1[i+1] == s2[i+1] { // i+1 相同，所以 i 是这一段的末尾
            res = append(res, []string{s1[start : i+1], s2[start : i+1]})
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s1 = "earth", s2 = "empty"
    // Output: [["ar","mp"],["h","y"]]
    // Explanation:
    //     The characters at indices 1 and 2 differ, forming the mismatched segment ["ar","mp"].
    //     The characters at index 4 differ, forming the mismatched segment ["h","y"].
    fmt.Println(findMismatchedSegments("earth", "empty")) // [["ar","mp"],["h","y"]]
    // Example 2:
    // Input: s1 = "action", s2 = "nation"
    // Output: [["ac","na"]]
    // Explanation:
    // The characters at indices 0 and 1 differ, forming the mismatched segment ["ac","na"]. 
    // All remaining characters are equal in both strings.
    fmt.Println(findMismatchedSegments("action", "nation")) // [["ac","na"]]
    // Example 3:
    // Input: s1 = "hot", s2 = "hot"
    // Output: []
    // Explanation:
    // The two strings are identical, so there are no mismatched segments.
    fmt.Println(findMismatchedSegments("hot", "hot")) // []

    fmt.Println(findMismatchedSegments("bluefrog", "leetcode")) // [[bluefrog leetcode]]
    fmt.Println(findMismatchedSegments("leetcode", "bluefrog")) // [[leetcode bluefrog]]
}