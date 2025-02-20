package main

// 面试题 17.25. Word Rectangle LCCI
// Given a list of millions of words, design an algorithm to create the largest possible rectangle of letters such that every row forms a word (reading left to right) and every column forms a word (reading top to bottom). 
// The words need not be chosen consecutively from the list but all rows must be the same length and all columns must be the same height.

// If there are more than one answer, return any one of them. 
// A word can be used more than once.

// Example 1:
// Input: ["this", "real", "hard", "trh", "hea", "iar", "sld"]
// Output:["this", "real", "hard"]

// Example 2:
// Input: ["aa"]
// Output: ["aa","aa"]

// Notes:
//     words.length <= 1000
//     words[i].length <= 100
//     It's guaranteed that all the words are randomly generated.

import "fmt"
import "sort"

func maxRectangle(words []string) []string {
    type Trie struct {
        next [26]*Trie
        end bool
    }
    insert := func(root *Trie, str string) {
        next := root
        for i := range str {
            key := str[i] - 'a'
            if next.next[key] == nil {
                next.next[key] = &Trie{}
            }
            next = next.next[key]
        }
        next.end = true
    }
    root := &Trie{}
    group := make(map[int][]int)
    for i, word := range words {
        insert(root, word)
        group[len(word)] = append(group[len(word)], i)
    }
    lengths := make([]int, len(group))
    {
        i := 0
        for _, v := range group {
            lengths[i] = len(words[v[0]])
            i++
        }
        sort.Slice(lengths, func(i, j int) bool {
            return lengths[i] > lengths[j]
        })
    }
    type Pair struct {
        nodes   []*Trie
        indexes []int
    }
    resIndexes := []int{}
    s := 0
    // 从最大长度开始
    for _, l := range lengths {
        if l * l <= s { break }
        start := &Pair{nodes: make([]*Trie, l)}
        for i := 0; i < l; i++ {
            start.nodes[i] = root
        }
        queue := []*Pair{ start }
        for len(queue) > 0 {
            q := queue[0]
            queue = queue[1:]
            allEnd := true
            for _, idx := range group[l] {
                word, flag := words[idx], false
                var nodes []*Trie
                for col := range word {
                    node := q.nodes[col].next[word[col]-'a']
                    if node == nil {
                        flag = true
                        break
                    }
                    nodes = append(nodes, node)
                    allEnd = allEnd && node.end
                }
                if flag { continue }
                indexes := make([]int, len(q.indexes)+1)
                copy(indexes, q.indexes)
                indexes[len(indexes)-1] = idx
                queue = append(queue, &Pair{nodes, indexes})
                if !allEnd { continue }
                // 计算面积
                if _s := (len(q.indexes) + 1) * l; _s > s {
                    s = _s
                    resIndexes = indexes
                }
            }
        }
    }
    res := make([]string, len(resIndexes))
    for i, idx := range resIndexes {
        res[i] = words[idx]
    }
    return res
}

func main() {
    // Example 1:
    // Input: ["this", "real", "hard", "trh", "hea", "iar", "sld"]
    // Output: ["this", "real", "hard"]
    fmt.Println(maxRectangle([]string{"this", "real", "hard", "trh", "hea", "iar", "sld"})) // ["this", "real", "hard"]
    // Example 2:
    // Input: ["aa"]
    // Output: ["aa","aa"]
    fmt.Println(maxRectangle([]string{"aa"})) // ["aa","aa"]

    fmt.Println(maxRectangle([]string{"bluefrog"})) // []
    fmt.Println(maxRectangle([]string{"leetcode"})) // []
}