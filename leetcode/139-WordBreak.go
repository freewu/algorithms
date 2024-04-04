package main

// 139. Word Break
// Given a string s and a dictionary of strings wordDict, 
// return true if s can be segmented into a space-separated sequence of one or more dictionary words.
// Note that the same word in the dictionary may be reused multiple times in the segmentation.
 
// Example 1:
// Input: s = "leetcode", wordDict = ["leet","code"]
// Output: true
// Explanation: Return true because "leetcode" can be segmented as "leet code".

// Example 2:
// Input: s = "applepenapple", wordDict = ["apple","pen"]
// Output: true
// Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
// Note that you are allowed to reuse a dictionary word.

// Example 3:
// Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
// Output: false
 
// Constraints:
//     1 <= s.length <= 300
//     1 <= wordDict.length <= 1000
//     1 <= wordDict[i].length <= 20
//     s and wordDict[i] consist of only lowercase English letters.
//     All the strings of wordDict are unique.

import "fmt"

// func wordBreak(s string, wordDict []string) bool {
//     // 把单词保存到map
//     m := make(map[string]int, len(wordDict)) 
//     sl, ll := -1, -1 // 最短的词&最长的词
//     min := func (x, y int) int { if x < y { return x; }; return y; }
//     max := func (x, y int) int { if x > y { return x; }; return y; }
//     for _,v := range wordDict {
//         m[v] = 1
//         // 第一次
//         if sl == -1 {
//             sl, ll = len(v), len(v)
//             continue
//         }
//         sl = min(sl, len(v))
//         ll = max(ll, len(v))
//     }
//     for i := 0; i < len(s); i++ {

//     }
//     return true
// }

// dp
func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s)+1)
    dp[0] = true
    dict := make(map[string]struct{})
    for _, val := range wordDict {
        dict[val] = struct{}{}
    }
    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            if dp[j] {
                substr := s[j:i]
                if _, ok := dict[substr]; ok {
                    // fmt.Println(substr)
                    dp[i] = true
                    break
                }
            }
        }
    }
    //fmt.Println(dp)
    return dp[len(dp)-1]
}

func main() {
    // Explanation: Return true because "leetcode" can be segmented as "leet code".
    fmt.Println(wordBreak("leetcode",[]string{"leet","code"})) // true

    // Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
    // Note that you are allowed to reuse a dictionary word.
    fmt.Println(wordBreak("applepenapple",[]string{"apple","pen"})) // true

    // Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
    // Output: false
    fmt.Println(wordBreak("catsandog",[]string{"cats","dog","sand","and","cat"})) // false
}