package main

// 1048. Longest String Chain
// You are given an array of words where each word consists of lowercase English letters.
// wordA is a predecessor of wordB if and only if we can insert exactly one letter anywhere in wordA without changing the order of the other characters to make it equal to wordB.
//     For example, "abc" is a predecessor of "abac", while "cba" is not a predecessor of "bcad".

// A word chain is a sequence of words [word1, word2, ..., wordk] with k >= 1, where word1 is a predecessor of word2, word2 is a predecessor of word3, and so on. 
// A single word is trivially a word chain with k == 1.
// Return the length of the longest possible word chain with words chosen from the given list of words.

// Example 1:
// Input: words = ["a","b","ba","bca","bda","bdca"]
// Output: 4
// Explanation: One of the longest word chains is ["a","ba","bda","bdca"].

// Example 2:
// Input: words = ["xbc","pcxbcf","xb","cxbc","pcxbc"]
// Output: 5
// Explanation: All the words can be put in a word chain ["xb", "xbc", "cxbc", "pcxbc", "pcxbcf"].

// Example 3:
// Input: words = ["abcd","dbqca"]
// Output: 1
// Explanation: The trivial word chain ["abcd"] is one of the longest word chains.
// ["abcd","dbqca"] is not a valid word chain because the ordering of the letters is changed.
 
// Constraints:
//     1 <= words.length <= 1000
//     1 <= words[i].length <= 16
//     words[i] only consists of lowercase English letters.

import "fmt"
import "sort"
import "slices"

func longestStrChain(words []string) int {
    sort.Slice(words, func(i, j int) bool { // 按长度从小到大排序
        return len(words[i]) < len(words[j])
    })
    dp, res := make(map[string]int), 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, word := range words {
        dp[word] = 1
        for i := 0; i < len(word); i++ {
            prev_word := word[:i] + word[i+1:]
            if val, exists := dp[prev_word]; exists {
                dp[word] = max(dp[word], val + 1)
            }
        }
        res = max(res, dp[word])
    }
    return res
}

func longestStrChain1(words []string) int {
    sort.Slice(words,func(i,j int) bool { return len(words[i]) < len(words[j])})
    max := func (x, y int) int { if x > y { return x; }; return y; }
    dp, res := map[string]int{}, 0
    for _, s := range words {
        t := 0
        for i := range s {
            t = max(t, dp[s[:i] + s[i+1:]])
        }
        dp[s] = t + 1
        res = max(res, t + 1)
    }
    return res
}


func longestStrChain2(words []string) int {
    type Node struct {
        str     string
        maxPath int
    }
    NewNode := func (str string) *Node {
        return &Node{str, 1}
    }
    lenToWords, maxLen := make(map[int][]*Node, len(words)), 0
    for _, word := range words {
        lenToWords[len(word)] = append(lenToWords[len(word)], NewNode(word))
        if len(word) > maxLen {
            maxLen = len(word)
        }
    }
    lens, maxPath := make([]int, 0, len(lenToWords)), 1
    for l := range lenToWords {
        lens = append(lens, l)
    }
    slices.Sort(lens)
    isSuccessor := func (prev, next string) bool {
        i := 0
        for i < len(prev) && prev[i] == next[i] {
            i += 1
        }
        for i < len(prev) && prev[i] == next[i+1] {
            i += 1
        }
        return i == len(prev)
    }
    for _, prevLen := range lens {
        for _, prev := range lenToWords[prevLen] {
            for _, next := range lenToWords[prevLen+1] {
                if isSuccessor(prev.str, next.str) && next.maxPath < prev.maxPath+1 {
                    next.maxPath = prev.maxPath + 1
                    if next.maxPath > maxPath {
                        maxPath = next.maxPath
                    }
                }
            }
        }
    }
    return maxPath
}

func main() {
    // Example 1:
    // Input: words = ["a","b","ba","bca","bda","bdca"]
    // Output: 4
    // Explanation: One of the longest word chains is ["a","ba","bda","bdca"].
    fmt.Println(longestStrChain([]string{"a","b","ba","bca","bda","bdca"})) // 4
    // Example 2:
    // Input: words = ["xbc","pcxbcf","xb","cxbc","pcxbc"]
    // Output: 5
    // Explanation: All the words can be put in a word chain ["xb", "xbc", "cxbc", "pcxbc", "pcxbcf"].
    fmt.Println(longestStrChain([]string{"xbc","pcxbcf","xb","cxbc","pcxbc"})) // 5
    // Example 3:
    // Input: words = ["abcd","dbqca"]
    // Output: 1
    // Explanation: The trivial word chain ["abcd"] is one of the longest word chains.
    // ["abcd","dbqca"] is not a valid word chain because the ordering of the letters is changed.
    fmt.Println(longestStrChain([]string{"abcd","dbqca"})) // 1

    fmt.Println(longestStrChain1([]string{"a","b","ba","bca","bda","bdca"})) // 4
    fmt.Println(longestStrChain1([]string{"xbc","pcxbcf","xb","cxbc","pcxbc"})) // 5
    fmt.Println(longestStrChain1([]string{"abcd","dbqca"})) // 1

    fmt.Println(longestStrChain2([]string{"a","b","ba","bca","bda","bdca"})) // 4
    fmt.Println(longestStrChain2([]string{"xbc","pcxbcf","xb","cxbc","pcxbc"})) // 5
    fmt.Println(longestStrChain2([]string{"abcd","dbqca"})) // 1
}