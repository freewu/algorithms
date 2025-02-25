package main

// 面试题 17.15. Longest Word LCCI
// Given a list of words, write a program to find the longest word made of other words in the list. 
// If there are more than one answer, return the one that has smallest lexicographic order. 
// If no answer, return an empty string.

// Example:
// Input:  ["cat","banana","dog","nana","walk","walker","dogwalker"]
// Output:  "dogwalker"
// Explanation:  "dogwalker" can be made of "dog" and "walker".

// Note:
//     0 <= len(words) <= 200
//     1 <= len(words[i]) <= 100

import "fmt"
import "sort"

func longestWord(words []string) string {
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })
    res, suffix := "", [26][]int{}
    check := func(s, sub string, i int) bool {
        if i < len(sub) - 1 { return false }
        for j := len(sub) - 1; j >= 0; j-- {
            if sub[j] != s[i] { return false }
            i--
        }
        return true
    }
    var dfs func(word string, start int) bool
    dfs = func(word string, start int) bool {
        if start == -1 { return true }
        for _, key := range suffix[word[start] - 'a'] {
            if check(word, words[key], start) && dfs(word, start - len(words[key])) {
                return true
            }
        }
        return false
    }
    for i, word := range words {
        start := len(word) - 1
        if start < 0 { continue }
        index := word[start] - 'a'
        if dfs(word, start) && len(word) >= len(res) {
            if len(word) > len(res) {
                res = word
            } else if word < res {
                res = word
            }
        }
        suffix[index] = append(suffix[index], i)
    }
    return res
}

func main() {
    // Example:
    // Input:  ["cat","banana","dog","nana","walk","walker","dogwalker"]
    // Output:  "dogwalker"
    // Explanation:  "dogwalker" can be made of "dog" and "walker".
    fmt.Println(longestWord([]string{"cat","banana","dog","nana","walk","walker","dogwalker"})) // dogwalker

    fmt.Println(longestWord([]string{"blue","leet","frog","leetcode","bluefrog","code","abc"})) // bluefrog
}