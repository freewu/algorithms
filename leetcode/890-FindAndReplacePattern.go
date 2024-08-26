package main

// 890. Find and Replace Pattern
// Given a list of strings words and a string pattern, return a list of words[i] that match pattern. 
// You may return the answer in any order.

// A word matches the pattern if there exists a permutation of letters p so 
// that after replacing every letter x in the pattern with p(x), we get the desired word.

// Recall that a permutation of letters is a bijection from letters to letters: 
// every letter maps to another letter, and no two letters map to the same letter.

// Example 1:
// Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
// Output: ["mee","aqq"]
// Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}. 
// "ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation, since a and b map to the same letter.

// Example 2:
// Input: words = ["a","b","c"], pattern = "a"
// Output: ["a","b","c"]

// Constraints:
//     1 <= pattern.length <= 20
//     1 <= words.length <= 50
//     words[i].length == pattern.length
//     pattern and words[i] are lowercase English letters.

import "fmt"

func findAndReplacePattern(words []string, pattern string) []string {
    res := []string{}
    for _, word := range words {
        flag, wp, pw := true, map[byte]byte{}, map[byte]byte{}
        for i := 0; i < len(word); i++ {
            if ( pw[pattern[i]] != 0 && pw[pattern[i]] != word[i]    ) || 
               ( wp[word[i]]    != 0 && wp[word[i]]    != pattern[i] ) { // 对应位置不一至
                flag = false
                break
            }
            pw[pattern[i]] = word[i]
            wp[word[i]] = pattern[i]
        }
        if flag == true { 
            res = append(res, word) 
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
    // Output: ["mee","aqq"]
    // Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}. 
    // "ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation, since a and b map to the same letter.
    fmt.Println(findAndReplacePattern([]string{"abc","deq","mee","aqq","dkd","ccc"}, "abb")) // ["mee","aqq"]
    // Example 2:
    // Input: words = ["a","b","c"], pattern = "a"
    // Output: ["a","b","c"]
    fmt.Println(findAndReplacePattern([]string{"a","b","c"}, "a")) // ["a","b","c"]
}