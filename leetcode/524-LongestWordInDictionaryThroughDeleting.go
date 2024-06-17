package main

// 524. Longest Word in Dictionary through Deleting
// Given a string s and a string array dictionary, 
// return the longest string in the dictionary that can be formed by deleting some of the given string characters. 
// If there is more than one possible result, return the longest word with the smallest lexicographical order. 
// If there is no possible result, return the empty string.

// Example 1:
// Input: s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
// Output: "apple"

// Example 2:
// Input: s = "abpcplea", dictionary = ["a","b","c"]
// Output: "a"

// Constraints:
//     1 <= s.length <= 1000
//     1 <= dictionary.length <= 1000
//     1 <= dictionary[i].length <= 1000
//     s and dictionary[i] consist of lowercase English letters.

import "fmt"
import "sort"

func findLongestWord(s string, dictionary []string) string {
    res := ""
    isSubSequence := func(p string, s string) bool {
        i := 0
        for j := 0; j < len(s) && i < len(p); j++ {
            if p[i] == s[j] {
                i++
            }
        }
        return i == len(p)
    }
    for _, v := range dictionary {
        if isSubSequence(v, s) {
            if len(v) > len(res) || len(v) == len(res) && v < res {
                res = v
            }
        }
    }
    return res
}

func findLongestWord1(s string, dictionary []string) string {
    res := ""
    sort.Slice(dictionary, func(i int, j int) bool {
        return dictionary[i] < dictionary[j]
    })
    for _, v := range dictionary {
        index := 0
        for i := 0; i < len(s) && index < len(v); i++ {
            if s[i] == v[index] {
                index++
            }
        }
        if index == len(v) && len(v) > len(res) {
            res = v
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
    // Output: "apple"
    fmt.Println(findLongestWord("abpcplea",[]string{"ale","apple","monkey","plea"})) // "apple"
    // Example 2:
    // Input: s = "abpcplea", dictionary = ["a","b","c"]
    // Output: "a"
    fmt.Println(findLongestWord("abpcplea",[]string{"a","b","c"})) // "a"

    fmt.Println(findLongestWord1("abpcplea",[]string{"ale","apple","monkey","plea"})) // "apple"
    fmt.Println(findLongestWord1("abpcplea",[]string{"a","b","c"})) // "a"
}