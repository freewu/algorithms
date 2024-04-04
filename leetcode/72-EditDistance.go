package main

// 72. Edit Distance
// Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
// You have the following three operations permitted on a word:
//     Insert a character
//     Delete a character
//     Replace a character
 
// Example 1:
// Input: word1 = "horse", word2 = "ros"
// Output: 3
// Explanation: 
// horse -> rorse (replace 'h' with 'r')
// rorse -> rose (remove 'r')
// rose -> ros (remove 'e')

// Example 2:
// Input: word1 = "intention", word2 = "execution"
// Output: 5
// Explanation: 
// intention -> inention (remove 't')
// inention -> enention (replace 'i' with 'e')
// enention -> exention (replace 'n' with 'x')
// exention -> exection (replace 'n' with 'c')
// exection -> execution (insert 'u')
 
// Constraints:
//     0 <= word1.length, word2.length <= 500
//     word1 and word2 consist of lowercase English letters.

import "fmt"

// dp 
func minDistance(word1, word2 string) int {
    if len(word1) == 0 || len(word2) == 0 {
        return len(word1) + len(word2)
    }
    ed, last := make([]int, len(word2)), make([]int, len(word2))
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := range word1 {
        for j := range word2 {
            if word1[i] == word2[j] {
                switch {
                case i == 0, j == 0:
                    ed[j] = i + j
                default:
                    ed[j] = last[j-1]
                }
            } else {
                switch {
                case i == 0 && j == 0:
                    ed[j] = 1
                case i == 0:
                    ed[j] = ed[j-1] + 1
                case j == 0:
                    ed[j] = last[j] + 1
                default:
                    ed[j] = min(min(last[j], last[j-1]), ed[j-1]) + 1
                }
            }
        }
        ed, last = last, ed
    }
    return last[len(last) - 1]
}

func main() {
    // horse -> rorse (replace 'h' with 'r')
    // rorse -> rose (remove 'r')
    // rose -> ros (remove 'e')
    fmt.Println(minDistance("horse","ros")) // 3
    // intention -> inention (remove 't')
    // inention -> enention (replace 'i' with 'e')
    // enention -> exention (replace 'n' with 'x')
    // exention -> exection (replace 'n' with 'c')
    // exection -> execution (insert 'u')
    fmt.Println(minDistance("intention","execution")) // 5
}