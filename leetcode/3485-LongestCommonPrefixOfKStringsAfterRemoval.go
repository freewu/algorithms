package main

// 3485. Longest Common Prefix of K Strings After Removal
// You are given an array of strings words and an integer k.

// For each index i in the range [0, words.length - 1], find the length of the longest common prefix among any k strings (selected at distinct indices) from the remaining array after removing the ith element.

// Return an array answer, where answer[i] is the answer for ith element. 
// If removing the ith element leaves the array with fewer than k strings, answer[i] is 0.

// Example 1:
// Input: words = ["jump","run","run","jump","run"], k = 2
// Output: [3,4,4,3,4]
// Explanation:
// Removing index 0 ("jump"):
// words becomes: ["run", "run", "jump", "run"]. "run" occurs 3 times. Choosing any two gives the longest common prefix "run" (length 3).
// Removing index 1 ("run"):
// words becomes: ["jump", "run", "jump", "run"]. "jump" occurs twice. Choosing these two gives the longest common prefix "jump" (length 4).
// Removing index 2 ("run"):
// words becomes: ["jump", "run", "jump", "run"]. "jump" occurs twice. Choosing these two gives the longest common prefix "jump" (length 4).
// Removing index 3 ("jump"):
// words becomes: ["jump", "run", "run", "run"]. "run" occurs 3 times. Choosing any two gives the longest common prefix "run" (length 3).
// Removing index 4 ("run"):
// words becomes: ["jump", "run", "run", "jump"]. "jump" occurs twice. Choosing these two gives the longest common prefix "jump" (length 4).

// Example 2:
// Input: words = ["dog","racer","car"], k = 2
// Output: [0,0,0]
// Explanation:
// Removing any index results in an answer of 0.

// Constraints:
//     1 <= k <= words.length <= 10^5
//     1 <= words[i].length <= 10^4
//     words[i] consists of lowercase English letters.
//     The sum of words[i].length is smaller than or equal 10^5.

import "fmt"
import "sort"
import "strings"
import "slices"
import "cmp"

type Node struct {
    Childs[26] *Node
    Freq int
}

type Trie struct {
    Root *Node
    K int
    PrefixesFreqMap map[string]int 
    PrefixesArray []string // <- in golang we cannot order the map directly, so we need to use this axiliary array.
    // In C++ is not necessary, you can use std::map with a custom comparator.
}

func NewTrie(words []string, k int) (* Trie){
    trie := Trie{K: k, Root: &Node{}}
    trie.PrefixesFreqMap = make(map[string]int)
    for _,word := range words {
        trie.AddWord(word)
    }
    sort.Slice(trie.PrefixesArray, func(i, j int) bool{
        if len(trie.PrefixesArray[i]) != len(trie.PrefixesArray[j]) {
            return len(trie.PrefixesArray[i]) > len(trie.PrefixesArray[j])
        }
        return trie.PrefixesArray[i] > trie.PrefixesArray[j]
    })
    return &trie
}

func (t *Trie) AddWord(word string) {
    it := t.Root
    for i := 0 ; i < len(word); i++ {
        index := word[i] - 'a'
        if it.Childs[index] == nil {
            it.Childs[index] = &Node{}
        }
        it.Childs[index].Freq++
        if it.Childs[index].Freq >= t.K {
            prefix := word[0:i + 1]
            if _, ok := t.PrefixesFreqMap[prefix]; !ok{
                t.PrefixesArray = append(t.PrefixesArray, prefix)
            }
            t.PrefixesFreqMap[prefix] = it.Childs[index].Freq
        }
        it = it.Childs[index]
    }
}

func longestCommonPrefix(words []string, k int) []int {
    trie := NewTrie(words,k)
    res := make([]int,len(words))
    for i, word :=  range words {
        for _, prefix := range trie.PrefixesArray {
            isPrefix := strings.HasPrefix(word, prefix)
            if (isPrefix &&  trie.PrefixesFreqMap[prefix] > k) || (!isPrefix && trie.PrefixesFreqMap[prefix] >= k) {
                res[i] = len(prefix)
                break
            }
        }
    }
    return res
}

func longestCommonPrefix1(words []string, k int) []int {
    n := len(words)
    if k >= n { // 移除一个字符串后，剩余字符串少于 k 个
        return make([]int, n)
    }
    idx := make([]int, n)
    for i := range idx {
        idx[i] = i
    }
    calcLCP := func (s, t string) int { // 计算 s 和 t 的最长公共前缀（LCP）长度
        n := min(len(s), len(t))
        for i := 0; i < n; i++ {
            if s[i] != t[i] {
                return i
            }
        }
        return n
    }
    slices.SortFunc(idx, func(i, j int) int { 
        return cmp.Compare(words[i], words[j]) 
    })
    // 计算最大 LCP 长度和次大 LCP 长度，同时记录最大 LCP 来自哪里
    mx, mx2, mxI := -1, -1, -1
    for i := 0; i < n - k + 1; i++ {
        // 排序后，[i, i+k-1] 的 LCP 等于两端点的 LCP
        lcp := calcLCP(words[idx[i]], words[idx[i+k-1]])
        if lcp >= mx {
            mx, mx2, mxI = lcp, mx, i
        } else if lcp > mx2 {
            mx2 = lcp
        }
    }
    res := make([]int, n)
    for i := range res {
        res[i] = mx // 先初始化成最大 LCP 长度
    }
    // 移除下标在 [mxI, mxI+k-1] 中的字符串，会导致最大 LCP 变成次大 LCP
    for _, i := range idx[mxI : mxI+k] {
        res[i] = mx2 // 改成次大 LCP 长度
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["jump","run","run","jump","run"], k = 2
    // Output: [3,4,4,3,4]
    // Explanation:
    // Removing index 0 ("jump"):
    // words becomes: ["run", "run", "jump", "run"]. "run" occurs 3 times. Choosing any two gives the longest common prefix "run" (length 3).
    // Removing index 1 ("run"):
    // words becomes: ["jump", "run", "jump", "run"]. "jump" occurs twice. Choosing these two gives the longest common prefix "jump" (length 4).
    // Removing index 2 ("run"):
    // words becomes: ["jump", "run", "jump", "run"]. "jump" occurs twice. Choosing these two gives the longest common prefix "jump" (length 4).
    // Removing index 3 ("jump"):
    // words becomes: ["jump", "run", "run", "run"]. "run" occurs 3 times. Choosing any two gives the longest common prefix "run" (length 3).
    // Removing index 4 ("run"):
    // words becomes: ["jump", "run", "run", "jump"]. "jump" occurs twice. Choosing these two gives the longest common prefix "jump" (length 4).
    fmt.Println(longestCommonPrefix([]string{"jump","run","run","jump","run"}, 2)) // [3,4,4,3,4]
    // Example 2:
    // Input: words = ["dog","racer","car"], k = 2
    // Output: [0,0,0]
    // Explanation:
    // Removing any index results in an answer of 0.
    fmt.Println(longestCommonPrefix([]string{"dog","racer","car"}, 2)) // [0,0,0]

    fmt.Println(longestCommonPrefix([]string{"bluefrog","leetcode"}, 2)) // [0,0]

    fmt.Println(longestCommonPrefix1([]string{"jump","run","run","jump","run"}, 2)) // [3,4,4,3,4]
    fmt.Println(longestCommonPrefix1([]string{"dog","racer","car"}, 2)) // [0,0,0]
    fmt.Println(longestCommonPrefix1([]string{"bluefrog","leetcode"}, 2)) // [0,0]
}