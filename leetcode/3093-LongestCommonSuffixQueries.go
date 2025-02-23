package main

// 3093. Longest Common Suffix Queries
// You are given two arrays of strings wordsContainer and wordsQuery.

// For each wordsQuery[i], you need to find a string from wordsContainer that has the longest common suffix with wordsQuery[i]. 
// If there are two or more strings in wordsContainer that share the longest common suffix, find the string that is the smallest in length. 
// If there are two or more such strings that have the same smallest length, find the one that occurred earlier in wordsContainer.

// Return an array of integers ans, where ans[i] is the index of the string in wordsContainer that has the longest common suffix with wordsQuery[i].

// Example 1:
// Input: wordsContainer = ["abcd","bcd","xbcd"], wordsQuery = ["cd","bcd","xyz"]
// Output: [1,1,1]
// Explanation:
// Let's look at each wordsQuery[i] separately:
// For wordsQuery[0] = "cd", strings from wordsContainer that share the longest common suffix "cd" are at indices 0, 1, and 2. Among these, the answer is the string at index 1 because it has the shortest length of 3.
// For wordsQuery[1] = "bcd", strings from wordsContainer that share the longest common suffix "bcd" are at indices 0, 1, and 2. Among these, the answer is the string at index 1 because it has the shortest length of 3.
// For wordsQuery[2] = "xyz", there is no string from wordsContainer that shares a common suffix. Hence the longest common suffix is "", that is shared with strings at index 0, 1, and 2. Among these, the answer is the string at index 1 because it has the shortest length of 3.

// Example 2:
// Input: wordsContainer = ["abcdefgh","poiuygh","ghghgh"], wordsQuery = ["gh","acbfgh","acbfegh"]
// Output: [2,0,2]
// Explanation:
// Let's look at each wordsQuery[i] separately:
// For wordsQuery[0] = "gh", strings from wordsContainer that share the longest common suffix "gh" are at indices 0, 1, and 2. Among these, the answer is the string at index 2 because it has the shortest length of 6.
// For wordsQuery[1] = "acbfgh", only the string at index 0 shares the longest common suffix "fgh". Hence it is the answer, even though the string at index 2 is shorter.
// For wordsQuery[2] = "acbfegh", strings from wordsContainer that share the longest common suffix "gh" are at indices 0, 1, and 2. Among these, the answer is the string at index 2 because it has the shortest length of 6.

// Constraints:
//     1 <= wordsContainer.length, wordsQuery.length <= 10^4
//     1 <= wordsContainer[i].length <= 5 * 10^3
//     1 <= wordsQuery[i].length <= 5 * 10^3
//     wordsContainer[i] consists only of lowercase English letters.
//     wordsQuery[i] consists only of lowercase English letters.
//     Sum of wordsContainer[i].length is at most 5 * 10^5.
//     Sum of wordsQuery[i].length is at most 5 * 10^5.

import "fmt"

type Trie struct {
    ch    []*Trie
    index   int // stores min string index till current ch
}

func NewTrie(index int) *Trie {
    return &Trie{ ch:  make([]*Trie, 26), index: index }
}

func (t *Trie) Insert(str string, index int, ws []string) {
    temp := t
    for i := len(str) - 1; i >= 0; i-- {
        k := int(str[i] - 'a')
        if temp.ch[k] == nil {
            temp.ch[k] = NewTrie(index)
        }
        curr := temp.ch[k].index
        if len(ws[curr]) > len(str) {
            temp.ch[k].index = index
        }
        temp = temp.ch[k]
    }
}

func (t *Trie) Find(str string) int {
    temp := t
    for i := len(str) - 1; i >= 0; i-- {
        k := int(str[i] - 'a')
        if temp.ch[k] == nil { break }
        temp = temp.ch[k]
    }
    return temp.index
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
    t := NewTrie(-1)
    res, mn := make([]int, len(wordsQuery)), 0
    for i, str := range wordsContainer {
        t.Insert(str, i, wordsContainer)
        if len(wordsContainer[mn]) > len(str) {
            mn = i
        }
    }
    for i, str := range wordsQuery {
        res[i] = t.Find(str)
        if res[i] == -1 {
            res[i] = mn
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: wordsContainer = ["abcd","bcd","xbcd"], wordsQuery = ["cd","bcd","xyz"]
    // Output: [1,1,1]
    // Explanation:
    // Let's look at each wordsQuery[i] separately:
    // For wordsQuery[0] = "cd", strings from wordsContainer that share the longest common suffix "cd" are at indices 0, 1, and 2. Among these, the answer is the string at index 1 because it has the shortest length of 3.
    // For wordsQuery[1] = "bcd", strings from wordsContainer that share the longest common suffix "bcd" are at indices 0, 1, and 2. Among these, the answer is the string at index 1 because it has the shortest length of 3.
    // For wordsQuery[2] = "xyz", there is no string from wordsContainer that shares a common suffix. Hence the longest common suffix is "", that is shared with strings at index 0, 1, and 2. Among these, the answer is the string at index 1 because it has the shortest length of 3.
    fmt.Println(stringIndices([]string{"abcd","bcd","xbcd"}, []string{"cd","bcd","xyz"})) // [1,1,1]
    // Example 2:
    // Input: wordsContainer = ["abcdefgh","poiuygh","ghghgh"], wordsQuery = ["gh","acbfgh","acbfegh"]
    // Output: [2,0,2]
    // Explanation:
    // Let's look at each wordsQuery[i] separately:
    // For wordsQuery[0] = "gh", strings from wordsContainer that share the longest common suffix "gh" are at indices 0, 1, and 2. Among these, the answer is the string at index 2 because it has the shortest length of 6.
    // For wordsQuery[1] = "acbfgh", only the string at index 0 shares the longest common suffix "fgh". Hence it is the answer, even though the string at index 2 is shorter.
    // For wordsQuery[2] = "acbfegh", strings from wordsContainer that share the longest common suffix "gh" are at indices 0, 1, and 2. Among these, the answer is the string at index 2 because it has the shortest length of 6.
    fmt.Println(stringIndices([]string{"abcdefgh","poiuygh","ghghgh"}, []string{"gh","acbfgh","acbfegh"})) // [2,0,2]
}