package main

// 3253. Construct String with Minimum Cost (Easy)
// You are given a string target, an array of strings words, and an integer array costs, both arrays of the same length.

// Imagine an empty string s.

// You can perform the following operation any number of times (including zero):
//     Choose an index i in the range [0, words.length - 1].
//     Append words[i] to s.
//     The cost of operation is costs[i].

// Return the minimum cost to make s equal to target. If it's not possible, return -1.

// Example 1:
// Input: target = "abcdef", words = ["abdef","abc","d","def","ef"], costs = [100,1,1,10,5]
// Output: 7
// Explanation:
// The minimum cost can be achieved by performing the following operations:
// Select index 1 and append "abc" to s at a cost of 1, resulting in s = "abc".
// Select index 2 and append "d" to s at a cost of 1, resulting in s = "abcd".
// Select index 4 and append "ef" to s at a cost of 5, resulting in s = "abcdef".

// Example 2:
// Input: target = "aaaa", words = ["z","zz","zzz"], costs = [1,10,100]
// Output: -1
// Explanation:
// It is impossible to make s equal to target, so we return -1.

// Constraints:
//     1 <= target.length <= 2000
//     1 <= words.length == costs.length <= 50
//     1 <= words[i].length <= target.length
//     target and words[i] consist only of lowercase English letters.
//     1 <= costs[i] <= 10^5

import "fmt"

type Trie struct {
    children [26]*Trie
    cost     int
}

func NewTrie() *Trie {
    return &Trie{ cost: 1 << 31 }
}

func (t *Trie) insert(word string, cost int) {
    node := t
    for _, v := range word {
        index := v - 'a'
        if node.children[index] == nil {
            node.children[index] = NewTrie()
        }
        node = node.children[index]
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    node.cost = min(node.cost, cost)
}

func minimumCost(target string, words []string, costs []int) int {
    trie := NewTrie()
    for i, word := range words {
        trie.insert(word, costs[i])
    }
    n, inf := len(target), 1 << 31
    facts := make([]int, n)
    var dfs func(int) int
    dfs = func(i int) int {
        if i >= n { return 0 }
        if facts[i] != 0 { return facts[i] }
        facts[i] = inf
        node := trie
        for j := i; j < n; j++ {
            index := target[j] - 'a'
            if node.children[index] == nil { return facts[i] }
            node = node.children[index]
            facts[i] = min(facts[i], node.cost + dfs(j + 1))
        }
        return facts[i]
    }
    if res := dfs(0); res < inf {
        return res
    }
    return -1
}

func main() {
    // Example 1:
    // Input: target = "abcdef", words = ["abdef","abc","d","def","ef"], costs = [100,1,1,10,5]
    // Output: 7
    // Explanation:
    // The minimum cost can be achieved by performing the following operations:
    // Select index 1 and append "abc" to s at a cost of 1, resulting in s = "abc".
    // Select index 2 and append "d" to s at a cost of 1, resulting in s = "abcd".
    // Select index 4 and append "ef" to s at a cost of 5, resulting in s = "abcdef".
    fmt.Println(minimumCost("abcdef", []string{"abdef","abc","d","def","ef"}, []int{100,1,1,10,5})) // 7
    // Example 2:
    // Input: target = "aaaa", words = ["z","zz","zzz"], costs = [1,10,100]
    // Output: -1
    // Explanation:
    // It is impossible to make s equal to target, so we return -1.
    fmt.Println(minimumCost("aaaa", []string{"z","zz","zzz"}, []int{1,10,100})) // -1
}