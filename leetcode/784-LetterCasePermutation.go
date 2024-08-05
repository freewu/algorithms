package main

// 784. Letter Case Permutation
// Given a string s, you can transform every letter individually to be lowercase or uppercase to create another string.
// Return a list of all possible strings we could create. Return the output in any order.

// Example 1:
// Input: s = "a1b2"
// Output: ["a1b2","a1B2","A1b2","A1B2"]

// Example 2:
// Input: s = "3z4"
// Output: ["3z4","3Z4"]

// Constraints:
//     1 <= s.length <= 12
//     s consists of lowercase English letters, uppercase English letters, and digits.

import "fmt"
import "unicode"

// backtrack
func letterCasePermutation(s string) []string {
    res, runes := []string{}, []rune(s)
    var dfs func(int)
    dfs = func(pos int) {
        if pos == len(runes) {
            res = append(res, string(runes))
            return
        }
        v := runes[pos]
        if unicode.ToUpper(v) != v { // lower => upper
            runes[pos] = unicode.ToUpper(v)
            dfs(pos+1)
            runes[pos] = v
        }
        if unicode.ToLower(v) != v { // upper => loer
            runes[pos] = unicode.ToLower(v)
            dfs(pos+1)
            runes[pos] = v
        }
        dfs(pos + 1)
    }
    dfs(0)
    return res
}

func main() {
    // Example 1:
    // Input: s = "a1b2"
    // Output: ["a1b2","a1B2","A1b2","A1B2"]
    fmt.Println(letterCasePermutation("a1b2")) // ["a1b2","a1B2","A1b2","A1B2"]
    // Example 2:
    // Input: s = "3z4"
    // Output: ["3z4","3Z4"]
    fmt.Println(letterCasePermutation("3z4")) // ["3z4","3Z4"]
}