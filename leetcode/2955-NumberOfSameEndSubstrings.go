package main

// 2955. Number of Same-End Substrings
// You are given a 0-indexed string s, and a 2D array of integers queries, 
// where queries[i] = [li, ri] indicates a substring of s starting from the index li and ending at the index ri (both inclusive), i.e. s[li..ri].

// Return an array ans where ans[i] is the number of same-end substrings of queries[i].

// A 0-indexed string t of length n is called same-end if it has the same character at both of its ends, i.e., t[0] == t[n - 1].

// A substring is a contiguous non-empty sequence of characters within a string.

// Example 1:
// Input: s = "abcaab", queries = [[0,0],[1,4],[2,5],[0,5]]
// Output: [1,5,5,10]
// Explanation: Here is the same-end substrings of each query:
// 1st query: s[0..0] is "a" which has 1 same-end substring: "a".
// 2nd query: s[1..4] is "bcaa" which has 5 same-end substrings: "bcaa", "bcaa", "bcaa", "bcaa", "bcaa".
// 3rd query: s[2..5] is "caab" which has 5 same-end substrings: "caab", "caab", "caab", "caab", "caab".
// 4th query: s[0..5] is "abcaab" which has 10 same-end substrings: "abcaab", "abcaab", "abcaab", "abcaab", "abcaab", "abcaab", "abcaab", "abcaab", "abcaab", "abcaab".

// Example 2:
// Input: s = "abcd", queries = [[0,3]]
// Output: [4]
// Explanation: The only query is s[0..3] which is "abcd". It has 4 same-end substrings: "abcd", "abcd", "abcd", "abcd".

// Constraints:
//     2 <= s.length <= 3 * 10^4
//     s consists only of lowercase English letters.
//     1 <= queries.length <= 3 * 10^4
//     queries[i] = [li, ri]
//     0 <= li <= ri < s.length

import "fmt"

func sameEndSubstringCount(s string, queries [][]int) []int {
    res, n := []int{},len(s)
    count := make([][]int, 26)
    for i := range count {
        count[i] = make([]int, n + 1)
    }
    for j := 1; j <= n; j++ {
        for i := 0; i < 26; i++ {
            count[i][j] = count[i][j - 1]
        }
        count[s[j-1] - 'a'][j]++
    }
    for _, q := range queries {
        l, r := q[0], q[1]
        res = append(res, r - l + 1)
        for i := 0; i < 26; i++ {
            v := count[i][r + 1] - count[i][l]
            res[len(res)-1] += v * (v - 1) / 2
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abcaab", queries = [[0,0],[1,4],[2,5],[0,5]]
    // Output: [1,5,5,10]
    // Explanation: Here is the same-end substrings of each query:
    // 1st query: s[0..0] is "a" which has 1 same-end substring: "a".
    // 2nd query: s[1..4] is "bcaa" which has 5 same-end substrings: "bcaa", "bcaa", "bcaa", "bcaa", "bcaa".
    // 3rd query: s[2..5] is "caab" which has 5 same-end substrings: "caab", "caab", "caab", "caab", "caab".
    // 4th query: s[0..5] is "abcaab" which has 10 same-end substrings: "abcaab", "abcaab", "abcaab", "abcaab", "abcaab", "abcaab", "abcaab", "abcaab", "abcaab", "abcaab".
    fmt.Println(sameEndSubstringCount("abcaab", [][]int{{0,0},{1,4},{2,5},{0,5}})) // [1,5,5,10]
    // Example 2:
    // Input: s = "abcd", queries = [[0,3]]
    // Output: [4]
    // Explanation: The only query is s[0..3] which is "abcd". It has 4 same-end substrings: "abcd", "abcd", "abcd", "abcd".
    fmt.Println(sameEndSubstringCount("abcd", [][]int{{0,3}})) // [4]
}