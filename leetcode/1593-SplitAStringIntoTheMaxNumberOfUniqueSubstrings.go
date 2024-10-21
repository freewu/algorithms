package main

// 1593. Split a String Into the Max Number of Unique Substrings
// Given a string s, return the maximum number of unique substrings that the given string can be split into.

// You can split string s into any list of non-empty substrings, 
// where the concatenation of the substrings forms the original string.
// However, you must split the substrings such that all of them are unique.

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: s = "ababccc"
// Output: 5
// Explanation: One way to split maximally is ['a', 'b', 'ab', 'c', 'cc']. Splitting like ['a', 'b', 'a', 'b', 'c', 'cc'] is not valid as you have 'a' and 'b' multiple times.

// Example 2:
// Input: s = "aba"
// Output: 2
// Explanation: One way to split maximally is ['a', 'ba'].

// Example 3:
// Input: s = "aa"
// Output: 1
// Explanation: It is impossible to split the string any further.
 
// Constraints:
//     1 <= s.length <= 16
//     s contains only lower case English letters.

import "fmt"

// Backtracking
func maxUniqueSplit(s string) int {
    res, n := 0, len(s)
    memo := make(map[string]bool)
    var backtracking func(pos, curr int) 
    backtracking = func(pos, curr int) {
        if pos >= n {
            if curr > res {
                res = curr
            }
            return
        }
        for i := pos + 1; i <= n; i++ {
            ns := s[pos:i]
            if v, ok := memo[ns]; ok && v { continue }
            memo[ns] = true
            backtracking(i, curr + 1)
            memo[ns] = false
        }
    }
    backtracking(0, 0)
    return res
}

func main() {
    // Example 1:
    // Input: s = "ababccc"
    // Output: 5
    // Explanation: One way to split maximally is ['a', 'b', 'ab', 'c', 'cc']. Splitting like ['a', 'b', 'a', 'b', 'c', 'cc'] is not valid as you have 'a' and 'b' multiple times.
    fmt.Println(maxUniqueSplit("ababccc")) // 5 ['a', 'b', 'a', 'b', 'c', 'cc']
    // Example 2:
    // Input: s = "aba"
    // Output: 2
    // Explanation: One way to split maximally is ['a', 'ba'].
    fmt.Println(maxUniqueSplit("aba")) // 2 ['a', 'ba']
    // Example 3:
    // Input: s = "aa"
    // Output: 1
    // Explanation: It is impossible to split the string any further.
    fmt.Println(maxUniqueSplit("aa")) // 1 ['aa']
}