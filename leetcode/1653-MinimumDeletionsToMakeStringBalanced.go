package main

// 1653. Minimum Deletions to Make String Balanced
// You are given a string s consisting only of characters 'a' and 'b'​​​​.

// You can delete any number of characters in s to make s balanced. 
// s is balanced if there is no pair of indices (i,j) such that i < j and s[i] = 'b' and s[j]= 'a'.

// Return the minimum number of deletions needed to make s balanced.

// Example 1:
// Input: s = "aababbab"
// Output: 2
// Explanation: You can either:
// Delete the characters at 0-indexed positions 2 and 6 ("aababbab" -> "aaabbb"), or
// Delete the characters at 0-indexed positions 3 and 6 ("aababbab" -> "aabbbb").

// Example 2:
// Input: s = "bbaaaaabb"
// Output: 2
// Explanation: The only solution is to delete the first two characters.

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is 'a' or 'b'​​.

import "fmt"
import "sort"
import "strings"

// 转化成最长上升子序列问题
func minimumDeletions(s string) int {
    inc := []rune{}
    for _, ch := range s {
        t := ch - 'a'
        if len(inc) < 1 {
            inc = append(inc, t)
        } else {
            pos := sort.Search(len(inc), func(i int) bool {
                return inc[i] > t
            })
            if pos == len(inc) {
                inc = append(inc, t)
            } else {
                inc[pos] = t
            }
        }
    }
    return len(s) - len(inc)
}

func minimumDeletions1(s string) int {
    del := strings.Count(s,"a")
    res := del
    for _, c := range s {
        del += int((c - 'a') * 2-1)
        if del < res {
            res = del
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "aababbab"
    // Output: 2
    // Explanation: You can either:
    // Delete the characters at 0-indexed positions 2 and 6 ("aababbab" -> "aaabbb"), or
    // Delete the characters at 0-indexed positions 3 and 6 ("aababbab" -> "aabbbb").
    fmt.Println(minimumDeletions("aababbab")) // 2
    // Example 2:
    // Input: s = "bbaaaaabb"
    // Output: 2
    // Explanation: The only solution is to delete the first two characters.
    fmt.Println(minimumDeletions("bbaaaaabb")) // 2

    fmt.Println(minimumDeletions("bbaabaaabb")) // 3
    fmt.Println(minimumDeletions("aaaaaaaaaa")) // 0
    fmt.Println(minimumDeletions("bbbbbbbbbb")) // 0
    fmt.Println(minimumDeletions("ababababab")) // 4
    fmt.Println(minimumDeletions("bababababa")) // 5
    fmt.Println(minimumDeletions("aaaaabbbbb")) // 0
    fmt.Println(minimumDeletions("bbbbbaaaaa")) // 5

    fmt.Println(minimumDeletions1("aababbab")) // 2
    fmt.Println(minimumDeletions1("bbaaaaabb")) // 2
    fmt.Println(minimumDeletions1("bbaabaaabb")) // 3
    fmt.Println(minimumDeletions1("bbaabaaabb")) // 3
    fmt.Println(minimumDeletions1("aaaaaaaaaa")) // 0
    fmt.Println(minimumDeletions1("bbbbbbbbbb")) // 0
    fmt.Println(minimumDeletions1("ababababab")) // 4
    fmt.Println(minimumDeletions1("bababababa")) // 5
    fmt.Println(minimumDeletions1("aaaaabbbbb")) // 0
    fmt.Println(minimumDeletions1("bbbbbaaaaa")) // 5
}