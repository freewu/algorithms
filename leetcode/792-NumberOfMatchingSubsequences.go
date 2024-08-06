package main

// 792. Number of Matching Subsequences
// Given a string s and an array of strings words, return the number of words[i] that is a subsequence of s.
// A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
// For example, "ace" is a subsequence of "abcde".

// Example 1:
// Input: s = "abcde", words = ["a","bb","acd","ace"]
// Output: 3
// Explanation: There are three strings in words that are a subsequence of s: "a", "acd", "ace".

// Example 2:
// Input: s = "dsahjpjauf", words = ["ahjpjau","ja","ahbwzgqnuk","tnmlanowax"]
// Output: 2

// Constraints:
//     1 <= s.length <= 5 * 10^4
//     1 <= words.length <= 5000
//     1 <= words[i].length <= 50
//     s and words[i] consist of only lowercase English letters.

import "fmt"

func numMatchingSubseq(s string, words []string) int {
    res, n := 0, len(s)
    for _, word := range words {
        l := len(word)
        if l > n {
            continue
        } else if l == n || l == 0 {
            if word == s { 
                res++ 
            }
        } else {
            i := 0
            for j := 0; j < n; j++ {
                if s[j] == word[i] { 
                    i++ 
                }
                if i == l { 
                    res++
                    break 
                }
            }
        }
    }
    return res
}

func numMatchingSubseq1(s string, words []string) int {
    type pair struct{ i,j int }
    ps := [26][]pair{} // positions,pairs
    for i, w :=range words {
        ps[w[0]-'a'] = append(ps[w[0]-'a'], pair{i,0})
    }
    res := 0
    for _, c := range s {
        q := ps[c-'a'] //queue
        ps[c-'a'] = ps[c-'a'][:0] // 速度比ps[c-'a']=nil要快
        for _, p := range q { // pair
            p.j++
            if p.j == len(words[p.i]) {
                res++
            } else {
                w := words[p.i][p.j]-'a'
                ps[w] = append(ps[w],p)
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abcde", words = ["a","bb","acd","ace"]
    // Output: 3
    // Explanation: There are three strings in words that are a subsequence of s: "a", "acd", "ace".
    fmt.Println(numMatchingSubseq("abcde", []string{"a","bb","acd","ace"})) // 3
    // Example 2:
    // Input: s = "dsahjpjauf", words = ["ahjpjau","ja","ahbwzgqnuk","tnmlanowax"]
    // Output: 2
    fmt.Println(numMatchingSubseq("dsahjpjauf", []string{"ahjpjau","ja","ahbwzgqnuk","tnmlanowax"})) // 2

    fmt.Println(numMatchingSubseq1("abcde", []string{"a","bb","acd","ace"})) // 3
    fmt.Println(numMatchingSubseq1("dsahjpjauf", []string{"ahjpjau","ja","ahbwzgqnuk","tnmlanowax"})) // 2
}