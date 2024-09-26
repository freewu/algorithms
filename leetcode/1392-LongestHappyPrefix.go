package main

// 1392. Longest Happy Prefix
// A string is called a happy prefix if is a non-empty prefix which is also a suffix (excluding itself).

// Given a string s, return the longest happy prefix of s. Return an empty string "" if no such prefix exists.

// Example 1:
// Input: s = "level"
// Output: "l"
// Explanation: s contains 4 prefix excluding itself ("l", "le", "lev", "leve"), and suffix ("l", "el", "vel", "evel"). The largest prefix which is also suffix is given by "l".

// Example 2:
// Input: s = "ababab"
// Output: "abab"
// Explanation: "abab" is the largest prefix which is also suffix. They can overlap in the original string.

// Constraints:
//     1 <= s.length <= 10^5
//     s contains only lowercase English letters.

import "fmt"

// KMP algorithm
func longestPrefix(s string) string {
    i, j, n := 0, 1, len(s)
    lps := make([]int, n)
    lps[0] = 0
    
    for j < n {
        if s[i] == s[j] { // When the character matches, we want to update both the index and store the updated i in lps array.
            i++
            lps[j] = i
            j++
        } else if i == 0 { // Special case: when i=0, we can't decrease i, base case!
            lps[j] = 0
            j++
        } else { // If it doesnt matches, Backtrack (start matching from the index of previous match) 
            i = lps[i-1]
        }
    }
    return s[0:lps[n-1]]  // lps of last index will tell the max index of longest suffix which is also prefix
}

func longestPrefix1(s string) string {
    s += "#"
    n := len(s)
    next := make([]int, n)
    next[0], next[1] = -1, 0
    for i, j := 2, 0; i < n; {
        if s[i-1] == s[j] {
            j++
            next[i] = j
            i++
        } else if j > 0 {
            j = next[j]
        } else {
            next[i] = 0
            i++
        }
    }
    return s[:next[n-1]]
}

func main() {
    // Example 1:
    // Input: s = "level"
    // Output: "l"
    // Explanation: s contains 4 prefix excluding itself ("l", "le", "lev", "leve"), and suffix ("l", "el", "vel", "evel"). The largest prefix which is also suffix is given by "l".
    fmt.Println(longestPrefix("level")) // "l"
    // Example 2:
    // Input: s = "ababab"
    // Output: "abab"
    // Explanation: "abab" is the largest prefix which is also suffix. They can overlap in the original string.
    fmt.Println(longestPrefix("ababab")) // "abab"

    fmt.Println(longestPrefix1("level")) // "l"
    fmt.Println(longestPrefix1("ababab")) // "abab"
}