package main

// 1408. String Matching in an Array
// Given an array of string words, return all strings in words that is a substring of another word. 
// You can return the answer in any order.

// A substring is a contiguous sequence of characters within a string

// Example 1:
// Input: words = ["mass","as","hero","superhero"]
// Output: ["as","hero"]
// Explanation: "as" is substring of "mass" and "hero" is substring of "superhero".
// ["hero","as"] is also a valid answer.

// Example 2:
// Input: words = ["leetcode","et","code"]
// Output: ["et","code"]
// Explanation: "et", "code" are substring of "leetcode".

// Example 3:
// Input: words = ["blue","green","bu"]
// Output: []
// Explanation: No string of words is substring of another string.

// Constraints:
//     1 <= words.length <= 100
//     1 <= words[i].length <= 30
//     words[i] contains only lowercase English letters.
//     All the strings of words are unique.

import "fmt"
import "sort"
import "strings"

func stringMatching(words []string) []string {
    mp, res := make(map[string]bool), []string{}
    for _, v1 := range words {
        for _, v2 := range words {
            if v1 == v2 || len(v2) > len(v1) { continue } // 字符串一样或者长度更多就没有了必要判断了
            if _, ok := mp[v2]; !ok && strings.Contains(v1, v2) {
                mp[v2] = true
                res = append(res, v2)
            }
        }
    }
    return res
}

func stringMatching1(words []string) []string {
    sort.Slice(words, func(i, j int) bool { // 从小到大排序
        return len(words[i]) < len(words[j])
    })
    isSubStr := func(s, t string) bool {
        if len(s) < len(t) { return false }
        for i := 0; i < len(s); i++ {
            if s[i] == t[0] {
                p := 1
                for k := i + 1; p < len(t) && k < len(s); {
                    if s[k] != t[p] { break }
                    k++
                    p++
                }
                if p == len(t) {
                    return true
                }
            }
        }
        return false
    }
    res, mp := []string{}, map[string]struct{}{}
    for i := 1; i < len(words); i++ {
        for j := i - 1; j >= 0; j-- {
            if isSubStr(words[i], words[j]) {
                mp[words[j]] = struct{}{}
            }
        }
    }
    for k := range mp {
        res = append(res, k)
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["mass","as","hero","superhero"]
    // Output: ["as","hero"]
    // Explanation: "as" is substring of "mass" and "hero" is substring of "superhero".
    // ["hero","as"] is also a valid answer.
    fmt.Println(stringMatching([]string{"mass","as","hero","superhero"})) // ["as","hero"]
    // Example 2:
    // Input: words = ["leetcode","et","code"]
    // Output: ["et","code"]
    // Explanation: "et", "code" are substring of "leetcode".
    fmt.Println(stringMatching([]string{"leetcode","et","code"})) // ["et","code"]
    // Example 3:
    // Input: words = ["blue","green","bu"]
    // Output: []
    // Explanation: No string of words is substring of another string.
    fmt.Println(stringMatching([]string{"blue","green","bu"})) // []

    fmt.Println(stringMatching1([]string{"mass","as","hero","superhero"})) // ["as","hero"]
    fmt.Println(stringMatching1([]string{"leetcode","et","code"})) // ["et","code"]
    fmt.Println(stringMatching1([]string{"blue","green","bu"})) // []
}