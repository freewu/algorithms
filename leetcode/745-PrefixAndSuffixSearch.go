package main

// 745. Prefix and Suffix Search
// Design a special dictionary that searches the words in it by a prefix and a suffix.
// Implement the WordFilter class:
//     WordFilter(string[] words) 
//         Initializes the object with the words in the dictionary.
//     f(string pref, string suff) 
//         Returns the index of the word in the dictionary, which has the prefix pref and the suffix suff. 
//         If there is more than one valid index, return the largest of them. 
//         If there is no such word in the dictionary, return -1.

// Example 1:
// Input
// ["WordFilter", "f"]
// [[["apple"]], ["a", "e"]]
// Output
// [null, 0]
// Explanation
// WordFilter wordFilter = new WordFilter(["apple"]);
// wordFilter.f("a", "e"); // return 0, because the word at index 0 has prefix = "a" and suffix = "e".

// Constraints:
//     1 <= words.length <= 10^4
//     1 <= words[i].length <= 7
//     1 <= pref.length, suff.length <= 7
//     words[i], pref and suff consist of lowercase English letters only.
//     At most 10^4 calls will be made to the function f.

import "fmt"

type TrieNode struct {
    Children []*TrieNode
    Weight   int
}

func NewTrieNode() *TrieNode {
    return &TrieNode{
        Children: make([]*TrieNode, 27),
        Weight:   0,
    }
}

type WordFilter struct {
    root *TrieNode
}

func Constructor(words []string) WordFilter {
    rt := NewTrieNode()
    for i, _ := range words {
        w := "{" + words[i]
        insert(rt, w, i)
        for j := 0; j < len(w); j++ {
            insert(rt, w[j+1:]+w, i)
        }
    }
    return WordFilter{rt}
}

func insert(root *TrieNode, word string, weight int) {
    cur := root
    for _, c := range word {
        k := c - 'a'
        if cur.Children[k] == nil {
            cur.Children[k] = NewTrieNode()
        }
        cur = cur.Children[k]
        cur.Weight = weight
    }
}

func (this *WordFilter) F(prefix string, suffix string) int {
    cur, str := this.root, suffix+"{"+prefix
    for _, c := range str {
        if cur.Children[c-'a'] == nil {
            return -1
        }
        cur = cur.Children[c-'a']
    }
    return cur.Weight
}


type pair struct {
    idx int
    str string
}

type WordFilter1 struct {
    pre, suf map[string][]pair
}

func Constructor1(words []string) WordFilter1 {
    pre, suf := make(map[string][]pair), make(map[string][]pair)
    for j, s := range words {
        for i := 1; i <= len(s); i++ {
            pre[s[:i]] = append(pre[s[:i]], pair{j, s})
        }
        for i := 0; i < len(s); i++ {
            suf[s[i:]] = append(suf[s[i:]], pair{j, s})
        }
    }
    return WordFilter1{pre, suf}
}

func (wf *WordFilter1) F(pref string, suff string) int {
    preMatches, sufMatches := wf.pre[pref], wf.suf[suff]
    i, j := len(preMatches)-1, len(sufMatches)-1

    for i >= 0 && j >= 0 {
        if preMatches[i].idx == sufMatches[j].idx {
            return preMatches[i].idx
        }
        if preMatches[i].idx > sufMatches[j].idx {
            i--
        } else {
            j--
        }
    }
    return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */

func main() {
    // WordFilter wordFilter = new WordFilter(["apple"]);
    obj := Constructor([]string{"apple"})
    fmt.Println(obj)
    // wordFilter.f("a", "e"); // return 0, because the word at index 0 has prefix = "a" and suffix = "e".
    fmt.Println(obj.F("a","e")) // 0

    fmt.Println(obj.F("a","a")) // -1


    obj1 := Constructor1([]string{"apple"})
    fmt.Println(obj1)
    // wordFilter.f("a", "e"); // return 0, because the word at index 0 has prefix = "a" and suffix = "e".
    fmt.Println(obj1.F("a","e")) // 0
    fmt.Println(obj1.F("a","a")) // -1
}