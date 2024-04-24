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

func wordBreak1(s string, wordDict []string) bool {
    dp := make([]bool, len(s)+1) // dp[i] 是s[0:i+1] 是否可以利用字典中出现的单词拼接出s
    dp[0] = true // 初始化
    for i := 1; i <= len(s); i++ { // 递推
        for j := 0; j < len(wordDict); j++ {
            if i >= len(wordDict[j]) && s[i-len(wordDict[j]):i] == wordDict[j] {
                if dp[i] = dp[i-len(wordDict[j])]; dp[i] { // 将重复情况过滤，eg. dogs 在s时候现有s，后有gs。gs会污染
                    break
                } 
            }
        }
    }
    return dp[len(s)]
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

    fmt.Println(wordBreak1("leetcode",[]string{"leet","code"})) // true
    fmt.Println(wordBreak1("applepenapple",[]string{"apple","pen"})) // true
    fmt.Println(wordBreak1("catsandog",[]string{"cats","dog","sand","and","cat"})) // false
}