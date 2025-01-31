package main

// 2781. Length of the Longest Valid Substring
// You are given a string word and an array of strings forbidden.

// A string is called valid if none of its substrings are present in forbidden.

// Return the length of the longest valid substring of the string word.

// A substring is a contiguous sequence of characters in a string, possibly empty.

// Example 1:
// Input: word = "cbaaaabc", forbidden = ["aaa","cb"]
// Output: 4
// Explanation: There are 11 valid substrings in word: "c", "b", "a", "ba", "aa", "bc", "baa", "aab", "ab", "abc" and "aabc". The length of the longest valid substring is 4. 
// It can be shown that all other substrings contain either "aaa" or "cb" as a substring. 

// Example 2:
// Input: word = "leetcode", forbidden = ["de","le","e"]
// Output: 4
// Explanation: There are 11 valid substrings in word: "l", "t", "c", "o", "d", "tc", "co", "od", "tco", "cod", and "tcod". The length of the longest valid substring is 4.
// It can be shown that all other substrings contain either "de", "le", or "e" as a substring. 

// Constraints:
//     1 <= word.length <= 10^5
//     word consists only of lowercase English letters.
//     1 <= forbidden.length <= 10^5
//     1 <= forbidden[i].length <= 10
//     forbidden[i] consists only of lowercase English letters.

import "fmt"

func longestValidSubstring(word string, forbidden []string) int {
    mp := make(map[string]bool)
    for _, v := range forbidden {
        mp[v] = true
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, n := 0, len(word)
    for i, j := 0, 0; j < n; j++ {
        for k := j; k > max(j - 10, i - 1); k-- {
            if mp[word[k:j + 1]] {
                i = k + 1
                break
            }
        }
        res = max(res, j - i + 1)
    }
    return res
}

func longestValidSubstring1(word string, forbidden []string) int {
    mp := make(map[string]bool, len(forbidden))
    for _, v := range forbidden {
        mp[v] = true
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, l := 0, 0
    for r := range word {
        for i := r; i >= l && i > r - 10; i-- {
            if mp[word[i:r + 1]] {
                l = i + 1
                break
            }
        }
        res = max(res, r - l + 1)
    }
    return res
}

func main() {
    // Example 1:
    // Input: word = "cbaaaabc", forbidden = ["aaa","cb"]
    // Output: 4
    // Explanation: There are 11 valid substrings in word: "c", "b", "a", "ba", "aa", "bc", "baa", "aab", "ab", "abc" and "aabc". The length of the longest valid substring is 4. 
    // It can be shown that all other substrings contain either "aaa" or "cb" as a substring. 
    fmt.Println(longestValidSubstring("cbaaaabc", []string{"aaa","cb"})) // 4
    // Example 2:
    // Input: word = "leetcode", forbidden = ["de","le","e"]
    // Output: 4
    // Explanation: There are 11 valid substrings in word: "l", "t", "c", "o", "d", "tc", "co", "od", "tco", "cod", and "tcod". The length of the longest valid substring is 4.
    // It can be shown that all other substrings contain either "de", "le", or "e" as a substring. 
    fmt.Println(longestValidSubstring("leetcode", []string{"de","le","e"})) // 4

    fmt.Println(longestValidSubstring("bluefrog", []string{"de","le","e"})) // 4

    fmt.Println(longestValidSubstring1("cbaaaabc", []string{"aaa","cb"})) // 4
    fmt.Println(longestValidSubstring1("leetcode", []string{"de","le","e"})) // 4
    fmt.Println(longestValidSubstring1("bluefrog", []string{"de","le","e"})) // 4
}