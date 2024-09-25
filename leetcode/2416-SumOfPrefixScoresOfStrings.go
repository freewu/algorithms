package main

// 2416. Sum of Prefix Scores of Strings
// You are given an array words of size n consisting of non-empty strings.

// We define the score of a string word as the number of strings words[i] such that word is a prefix of words[i].
//     For example, if words = ["a", "ab", "abc", "cab"], 
//     then the score of "ab" is 2, since "ab" is a prefix of both "ab" and "abc".

// Return an array answer of size n where answer[i] is the sum of scores of every non-empty prefix of words[i].

// Note that a string is considered as a prefix of itself.

// Example 1:
// Input: words = ["abc","ab","bc","b"]
// Output: [5,4,3,2]
// Explanation: The answer for each string is the following:
// - "abc" has 3 prefixes: "a", "ab", and "abc".
// - There are 2 strings with the prefix "a", 2 strings with the prefix "ab", and 1 string with the prefix "abc".
// The total is answer[0] = 2 + 2 + 1 = 5.
// - "ab" has 2 prefixes: "a" and "ab".
// - There are 2 strings with the prefix "a", and 2 strings with the prefix "ab".
// The total is answer[1] = 2 + 2 = 4.
// - "bc" has 2 prefixes: "b" and "bc".
// - There are 2 strings with the prefix "b", and 1 string with the prefix "bc".
// The total is answer[2] = 2 + 1 = 3.
// - "b" has 1 prefix: "b".
// - There are 2 strings with the prefix "b".
// The total is answer[3] = 2.

// Example 2:
// Input: words = ["abcd"]
// Output: [4]
// Explanation:
// "abcd" has 4 prefixes: "a", "ab", "abc", and "abcd".
// Each prefix has a score of one, so the total is answer[0] = 1 + 1 + 1 + 1 = 4.

// Constraints:
//     1 <= words.length <= 1000
//     1 <= words[i].length <= 1000
//     words[i] consists of lowercase English letters.

import "fmt"
import "slices"

// 前缀树
type TireNode struct {
    count int
    sons [26]*TireNode
}

func (tree *TireNode) insert(s string) {
    curNode := tree
    for i := range s {
        if curNode.sons[s[i] - 'a'] == nil {
            curNode.sons[s[i] - 'a'] = new(TireNode)
        }
        curNode = curNode.sons[s[i] - 'a']
        curNode.count++ 
    }    
}

func (tree *TireNode) find(s string) int {
    curNode := tree
    res := 0
    for i := range s {
        curNode = curNode.sons[s[i] - 'a']
        res += curNode.count
    }    
    return res
}

func sumPrefixScores(words []string) []int {
    tree := new(TireNode)
    for i := range words {
        tree.insert(words[i])
    }   
    res := make([]int, len(words))
    for i := range words {
        res[i] += tree.find(words[i]) 
    }
    return res
}

// trie 树
func sumPrefixScores1(words []string) []int {
    type pair struct {
        i int
        w string
    }
    ps := make([]pair, len(words))
    for i, w := range words {
        ps[i] = pair{i, w}
    }
    slices.SortFunc(ps, func(a, b pair) int { return len(a.w) - len(b.w) })
    t := newTrie()
    for _, p := range ps {
        t.put(p.w)
    }
    res := make([]int, len(words))
    for _, p := range ps {
        res[p.i] = t.find(p.w)
    }
    return res
}

type TrieNode struct {
    children [26]*TrieNode
    cnt int
}

type Trie struct {
    root *TrieNode
}

func newTrie() *Trie {
    return &Trie{ &TrieNode{} }
}

func (t *Trie) put(s string) {
    o := t.root
    for _, b := range s {
        b -= 'a'
        if o.children[b] == nil {
            o.children[b] = &TrieNode{}
        }
        o = o.children[b]
        o.cnt++
    }
}

func (t *Trie) find(s string) int {
    res, o := 0, t.root
    for _, b := range s {
        b -= 'a'
        if o.children[b] == nil {
            return 0
        }
        o = o.children[b]
        res += o.cnt
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["abc","ab","bc","b"]
    // Output: [5,4,3,2]
    // Explanation: The answer for each string is the following:
    // - "abc" has 3 prefixes: "a", "ab", and "abc".
    // - There are 2 strings with the prefix "a", 2 strings with the prefix "ab", and 1 string with the prefix "abc".
    // The total is answer[0] = 2 + 2 + 1 = 5.
    // - "ab" has 2 prefixes: "a" and "ab".
    // - There are 2 strings with the prefix "a", and 2 strings with the prefix "ab".
    // The total is answer[1] = 2 + 2 = 4.
    // - "bc" has 2 prefixes: "b" and "bc".
    // - There are 2 strings with the prefix "b", and 1 string with the prefix "bc".
    // The total is answer[2] = 2 + 1 = 3.
    // - "b" has 1 prefix: "b".
    // - There are 2 strings with the prefix "b".
    // The total is answer[3] = 2.
    fmt.Println(sumPrefixScores([]string{ "abc","ab","bc","b"})) // [5,4,3,2]
    // Example 2:
    // Input: words = ["abcd"]
    // Output: [4]
    // Explanation:
    // "abcd" has 4 prefixes: "a", "ab", "abc", and "abcd".
    // Each prefix has a score of one, so the total is answer[0] = 1 + 1 + 1 + 1 = 4.
    fmt.Println(sumPrefixScores([]string{ "abcd" })) // [4]

    fmt.Println(sumPrefixScores1([]string{ "abc","ab","bc","b"})) // [5,4,3,2]
    fmt.Println(sumPrefixScores1([]string{ "abcd" })) // [4]
}