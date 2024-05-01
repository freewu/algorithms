package main 

// 140. Word Break II
// Given a string s and a dictionary of strings wordDict, add spaces in s to construct a sentence where each word is a valid dictionary word. 
// Return all such possible sentences in any order.

// Note that the same word in the dictionary may be reused multiple times in the segmentation.

// Example 1:
// Input: s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
// Output: ["cats and dog","cat sand dog"]

// Example 2:
// Input: s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
// Output: ["pine apple pen apple","pineapple pen apple","pine applepen apple"]
// Explanation: Note that you are allowed to reuse a dictionary word.

// Example 3:
// Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
// Output: []
 
// Constraints:
//     1 <= s.length <= 20
//     1 <= wordDict.length <= 1000
//     1 <= wordDict[i].length <= 10
//     s and wordDict[i] consist of only lowercase English letters.
//     All the strings of wordDict are unique.
//     Input is generated in a way that the length of the answer doesn't exceed 10^5.

import "fmt"

func wordBreak(s string, wordDict []string) []string {
    var dfs func (s string, cur int, wordDict []string) []string
    dfs = func (s string, cur int, wordDict []string) []string {
        res := []string{}
        for _, word := range wordDict {
            if len(word) + cur < len(s) && s[cur : cur + len(word)] == word {
                arr := dfs(s, cur + len(word), wordDict)
                for _, a := range arr {
                    res = append(res, word + " " + a)
                }
            } else if len(word) + cur == len(s) && s[cur : cur + len(word)] == word {
                res = append(res, word)
            }
        }
        return res
    }
    return dfs(s, 0, wordDict)
}

func main() {
    // Example 1:
    // Input: s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
    // Output: ["cats and dog","cat sand dog"]
    fmt.Println(wordBreak("catsanddog",[]string{"cat","cats","and","sand","dog"})) // ["cats and dog","cat sand dog"]
    // Example 2:
    // Input: s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
    // Output: ["pine apple pen apple","pineapple pen apple","pine applepen apple"]
    // Explanation: Note that you are allowed to reuse a dictionary word.
    fmt.Println(wordBreak("pineapplepenapple",[]string{"apple","pen","applepen","pine","pineapple"})) // ["pine apple pen apple","pineapple pen apple","pine applepen apple"]
    // Example 3:
    // Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
    // Output: []
    fmt.Println(wordBreak("catsandog",[]string{"cats","dog","sand","and","cat"})) // []
}