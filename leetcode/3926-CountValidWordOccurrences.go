package main

// 3926. Count Valid Word Occurrences
// You are given an array of strings chunks. The strings are concatenated in order to form a single string s.

// You are also given an array of strings queries.

// A word is defined as a substring of s that:
//     1. consists of lowercase English letters ('a' to 'z'),
//     2. may include hyphens ('-') only if each hyphen is surrounded by lowercase English letters, and
//     3. is not part of a longer substring that also satisfies the above conditions.

// Any character that is not a lowercase English letter or a valid hyphen acts as a separator.

// Return an integer array ans such that ans[i] is the number of occurrences of queries[i] as a word in s.

// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Input: chunks = ["hello wor","ld hello"], queries = ["hello","world","wor"]
// Output: [2,1,0]
// Explanation:
// Concatenating all strings in chunks gives s = "hello world hello".
// The valid words in s are "hello" which appears twice and "world" which appears once.
// Thus, the ans = [2, 1, 0].

// Example 2:
// Input: chunks = ["a--b a-","-c"], queries = ["a","b","c"]
// Output: [2,1,1]
// Explanation:
// Concatenating all strings in chunks gives s = "a--b a--c".
// The valid words in s are "a" which appears twice, "b" which appears once, and "c" which appears once.
// Thus, the ans = [2, 1, 1].

// Example 3:
// Input: chunks = ["hello"], queries = ["hello","ell"]
// Output: [1,0]
// Explanation:
// The valid word in s is "hello" which appears once.
// Thus, the ans = [1, 0].
 
// Constraints:
//     1 <= chunks.length <= 10^5
//     1 <= chunks[i].length <= 10^5​​​​​​​
//     chunks[i] may consist of lowercase English letters, spaces, and hyphens.
//     The total length of all strings in chunks does not exceed 10^5
//     1 <= queries.length <= 10^5
//     1 <= queries[i].length <= 10^5​​​​​​​
//     queries[i] is a valid word
//     The total length of all strings in queries does not exceed 10^5

import "fmt"
import "strings"

func countWordOccurrences(chunks []string, queries []string) []int {
    var s strings.Builder
    for _, chunk := range chunks {
        s.WriteString(chunk)
    }
    s.WriteRune(' ')
    parseStrings := func(s string) map[string]int {
        res := map[string]int{}
        start := -1
        for i,c := range s {
            if c == '-' {
                if s[i + 1] < 50 && start != -1 {
                    res[s[start: i]] ++
                    start = -1
                }
            } else if c >50 {
                if start == -1 {
                    start = i
                }
            } else {
                if start == -1 {
                    continue
                }
                res[s[start: i]] ++
                start = -1
            }
        }
        if start == -1 {
            return res
        }
        res[s[start: len(s)]]++
        return res
    }
    stringFreq := parseStrings(s.String())
    res := make([]int, len(queries))
    for i, query := range queries {
        res[i] = stringFreq[query]
    }
    return res
}

func countWordOccurrences1(chunks []string, queries []string) []int {
    var builder strings.Builder
    for _, chunk := range chunks {
        builder.WriteString(chunk)
    }
    s := builder.String()
    freq := make(map[string]int)
    i, n := 0, len(s)
    isLowerCase := func(ch byte) bool { return ch >= 'a' && ch <= 'z' }
    for i < n {
        if !isLowerCase(s[i]) {
            i++
            continue
        }
        start := i
        i++
        for i < n {
            if isLowerCase(s[i]) {
                i++
                continue
            }
            if s[i] == '-' &&
                i > 0 &&
                i+1 < n &&
                isLowerCase(s[i-1]) &&
                isLowerCase(s[i+1]) {
                i++
                continue
            }
            break
        }
        freq[s[start:i]]++
    }
    res := make([]int, len(queries))
    for i, q := range queries {
        res[i] = freq[q]
    }
    return res
}

func main() {
    // Example 1:
    // Input: chunks = ["hello wor","ld hello"], queries = ["hello","world","wor"]
    // Output: [2,1,0]
    // Explanation:
    // Concatenating all strings in chunks gives s = "hello world hello".
    // The valid words in s are "hello" which appears twice and "world" which appears once.
    // Thus, the ans = [2, 1, 0].
    fmt.Println(countWordOccurrences([]string{"hello wor","ld hello"}, []string{"hello","world","wor"})) // [2,1,0]
    // Example 2:
    // Input: chunks = ["a--b a-","-c"], queries = ["a","b","c"]
    // Output: [2,1,1]
    // Explanation:
    // Concatenating all strings in chunks gives s = "a--b a--c".
    // The valid words in s are "a" which appears twice, "b" which appears once, and "c" which appears once.
    // Thus, the ans = [2, 1, 1].
    fmt.Println(countWordOccurrences([]string{"a--b a-","-c"}, []string{"a","b","c"})) // [2,1,1]
    // Example 3:
    // Input: chunks = ["hello"], queries = ["hello","ell"]
    // Output: [1,0]
    // Explanation:
    // The valid word in s is "hello" which appears once.
    // Thus, the ans = [1, 0].
    fmt.Println(countWordOccurrences([]string{"hello"}, []string{"hello","ell"})) // [1,0]

    fmt.Println(countWordOccurrences1([]string{"hello wor","ld hello"}, []string{"hello","world","wor"})) // [2,1,0]
    fmt.Println(countWordOccurrences1([]string{"a--b a-","-c"}, []string{"a","b","c"})) // [2,1,1]
    fmt.Println(countWordOccurrences1([]string{"hello"}, []string{"hello","ell"})) // [1,0]
}