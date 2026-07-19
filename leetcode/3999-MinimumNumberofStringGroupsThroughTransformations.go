package main

// 3999. Minimum Number of String Groups Through Transformations
// You are given an array of strings words.

// Define a transformation on a string s as follows:
//     1.  Let E be the subsequence of characters at even indices of s.
//     2.  Let O be the subsequence of characters at odd indices of s.
//     3.  Independently cyclically shift E and O by any number of positions to the right, possibly zero.
//     4.  Reconstruct the string by placing the shifted E characters back into even indices and the shifted O characters back into odd indices.

// Two strings are equivalent if one can be transformed into the other by a single transformation.

// Partition words into the minimum number of groups such that:
//     1. Every string belongs to exactly one group.
//     2. Every pair of strings in the same group are equivalent.

// Return an integer denoting the minimum number of groups.

// Example 1:
// Input: words = ["ntgwz","zwntg"]
// Output: 1
// Explanation:
// For "ntgwz", the even-index subsequence is "ngz" and the odd-index subsequence is "tw".
// Shift "ngz" right by 1 position to obtain "zng", and shift "tw" right by 1 position to obtain "wt".
// After reconstructing the string, we obtain "zwntg".
// Therefore, both strings are equivalent and belong to the same group.

// Example 2:
// Input: words = ["abc","cab","bac","acb","bca","cba"]
// Output: 3
// Explanation:
// The strings can be partitioned into the following groups:
// ["abc","cba"]
// ["cab","bac"]
// ["acb","bca"]

// Example 3:
// Input: words = ["leet","abb","bab","deed","edde","code","bba"]
// Output: 5
// Explanation:
// The strings can be partitioned into the following groups:
// ["abb","bba"]
// ["deed","edde"]
// ["leet"]
// ["bab"]
// ["code"]
// ​​​​​​​​​​​​​​All pairs of strings in each group are equivalent.

// Constraints:
//     1 <= words.length <= 10^5
//     1 <= words[i].length <= 5 * 10^5
//     The sum of words[i].length does not exceed 5 * 10^5.
//     words[i] consist of lowercase English letters.

import "fmt"

func minimumGroups(words []string) int {
    smallestRotation := func(s string) string { // rotation of s using Booth's algorithm.
        n := len(s)
        if n <= 1 {
            return s
        }
        doubled := s + s
        i, j, offset := 0, 1, 0
        for i < n && j < n && offset < n {
            left := doubled[i+offset]
            right := doubled[j+offset]
            if left == right {
                offset++
                continue
            }
            if left > right {
                i = i + offset + 1
                if i <= j {
                    i = j + 1
                }
            } else {
                j = j + offset + 1
                if j <= i {
                    j = i + 1
                }
            }
            offset = 0
        }
        start := i
        if j < start {
            start = j
        }
        return doubled[start : start + n]  // can go for the return already. 
    }
    set := make(map[string]struct{})
    for _, word := range words {
        even := make([]byte, 0, (len(word)+1) / 2)
        odd := make([]byte, 0, len(word) / 2)
        for i := 0; i < len(word); i++ {
            if i%2 == 0 {
                even = append(even, word[i])
            } else {
                odd = append(odd, word[i])
            }
        }
        key := smallestRotation(string(even)) + "#" + smallestRotation(string(odd))
        set[key] = struct{}{}
    }
    return len(set)
}

func main() {
    // Example 1:
    // Input: words = ["ntgwz","zwntg"]
    // Output: 1
    // Explanation:
    // For "ntgwz", the even-index subsequence is "ngz" and the odd-index subsequence is "tw".
    // Shift "ngz" right by 1 position to obtain "zng", and shift "tw" right by 1 position to obtain "wt".
    // After reconstructing the string, we obtain "zwntg".
    // Therefore, both strings are equivalent and belong to the same group.
    fmt.Println(minimumGroups([]string{"ntgwz","zwntg"})) // 1
    // Example 2:
    // Input: words = ["abc","cab","bac","acb","bca","cba"]
    // Output: 3
    // Explanation:
    // The strings can be partitioned into the following groups:
    // ["abc","cba"]
    // ["cab","bac"]
    // ["acb","bca"]
    fmt.Println(minimumGroups([]string{"abc","cab","bac","acb","bca","cba"})) // 3
    // Example 3:
    // Input: words = ["leet","abb","bab","deed","edde","code","bba"]
    // Output: 5
    // Explanation:
    // The strings can be partitioned into the following groups:
    // ["abb","bba"]
    // ["deed","edde"]
    // ["leet"]
    // ["bab"]
    // ["code"]
    // ​​​​​​​​​​​​​​All pairs of strings in each group are equivalent.
    fmt.Println(minimumGroups([]string{"leet","abb","bab","deed","edde","code","bba"})) // 5

    fmt.Println(minimumGroups([]string{"bluefrog","leetcode", "freewu"})) // 3
}