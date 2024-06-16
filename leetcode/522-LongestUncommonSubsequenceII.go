package main

// 522. Longest Uncommon Subsequence II
// Given an array of strings strs, return the length of the longest uncommon subsequence between them. 
// If the longest uncommon subsequence does not exist, return -1.

// An uncommon subsequence between an array of strings is a string that is a subsequence of one string but not the others.
// A subsequence of a string s is a string that can be obtained after deleting any number of characters from s.
//     For example, "abc" is a subsequence of "aebdc" because you can delete the underlined characters in "aebdc" to get "abc". 
//     Other subsequences of "aebdc" include "aebdc", "aeb", and "" (empty string).

// Example 1:
// Input: strs = ["aba","cdc","eae"]
// Output: 3

// Example 2:
// Input: strs = ["aaa","aaa","aa"]
// Output: -1
 
// Constraints:
//     2 <= strs.length <= 50
//     1 <= strs[i].length <= 10
//     strs[i] consists of lowercase English letters.

import "fmt"
import "sort"

func findLUSlength(strs []string) int {
    m, nodes := make(map[string]int), []string{}
    for _, s := range strs { // 统计字符串出现次数
        m[s]++
    }
    for s, c := range m { // 取只出现了一次的字符串
        if c == 1 {
            nodes = append(nodes, s)
        }
    }
    sort.Slice(nodes, func(i, j int) bool { // 由长到短排序
        return len(nodes[i]) > len(nodes[j])
    })
    subseq := func (n, s string) bool {
        if len(n) >= len(s) {
            return false
        }
        ns := 0
        for _, c := range n {
            for ns < len(s) && byte(c) != s[ns] {
                ns++
            } 
            if ns >= len(s) {
                return false
            } else {
                ns++
            }
        } 
        return true
    }
    for _, n := range nodes {
        sub := true
        for _, s := range strs {
            if s == n {
                continue
            }
            if subseq(n, s) {
                sub = false
                break
            }
        }
        if sub {
            return len(n)
        }
    }
    return -1
}

func main() {
    // Example 1:
    // Input: strs = ["aba","cdc","eae"]
    // Output: 3
    fmt.Println(findLUSlength([]string{"aba","cdc","eae"})) // 3
    // Example 2:
    // Input: strs = ["aaa","aaa","aa"]
    // Output: -1
    fmt.Println(findLUSlength([]string{"aaa","aaa","aa"})) // -1
}