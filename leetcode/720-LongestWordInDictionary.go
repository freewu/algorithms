package main

// 720. Longest Word in Dictionary
// Given an array of strings words representing an English Dictionary, 
// return the longest word in words that can be built one character at a time by other words in words.

// If there is more than one possible answer, return the longest word with the smallest lexicographical order. 
// If there is no answer, return the empty string.

// Note that the word should be built from left to right with each additional character being added to the end of a previous word. 

// Example 1:
// Input: words = ["w","wo","wor","worl","world"]
// Output: "world"
// Explanation: The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".

// Example 2:
// Input: words = ["a","banana","app","appl","ap","apply","apple"]
// Output: "apple"
// Explanation: Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".

// Constraints:
//     1 <= words.length <= 1000
//     1 <= words[i].length <= 30
//     words[i] consists of lowercase English letters.

import "fmt"
import "sort"

func longestWord(words []string) string {
    sort.Slice(words, func(i int, j int) bool {
        if len(words[i]) == len(words[j]) {
            return words[i] > words[j]
        }
        return len(words[i]) < len(words[j])
    })
    res, set := "", make(map[string]struct{})
    set[""] = struct{}{}
    for _, s := range words {
        if _, ok := set[s[:len(s) - 1]]; ok {
            set[s] = struct{}{}
            res = s
        }
    }
    return res
}

func longestWord1(words []string) string {
    root := &Trie{}
    for _, word := range words {
        root.Insert(word)
    }
    longest := ""
    for _, word := range words {
        if root.Search(word) && // 能找到
           (len(word) > len(longest) || len(word) == len(longest) && word < longest) { // 字典序靠前
            longest = word
        }
    }
   return longest
}

type Trie struct {
    isEnd bool 
    children [26]*Trie 
}

func(this *Trie) Insert(s string) {
    cur := this 
    for i := range s {
        idx := s[i] - 'a'
        if cur.children[idx] == nil {
            cur.children[idx] = &Trie{}
        }
        cur = cur.children[idx]
    }
    cur.isEnd = true
}

func(this *Trie) Search(word string) bool {
    cur := this 
    for i := range word {
        idx := word[i] - 'a'
        if cur.children[idx] == nil || !cur.children[idx].isEnd {
            return false 
        }
        cur = cur.children[idx]
    }
    return true 
}

func main() {
    // Example 1:
    // Input: words = ["w","wo","wor","worl","world"]
    // Output: "world"
    // Explanation: The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".
    fmt.Println(longestWord([]string{"w","wo","wor","worl","world"})) // "world"
    // Example 2:
    // Input: words = ["a","banana","app","appl","ap","apply","apple"]
    // Output: "apple"
    // Explanation: Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".
    fmt.Println(longestWord([]string{"a","banana","app","appl","ap","apply","apple"})) // "apple"

    fmt.Println(longestWord1([]string{"w","wo","wor","worl","world"})) // "world"
    fmt.Println(longestWord1([]string{"a","banana","app","appl","ap","apply","apple"})) // "apple"
}