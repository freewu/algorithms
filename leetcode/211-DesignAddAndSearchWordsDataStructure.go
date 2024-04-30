package main

// 211. Design Add and Search Words Data Structure
// Design a data structure that supports adding new words and finding if a string matches any previously added string.
// Implement the WordDictionary class:
//     WordDictionary() Initializes the object.
//     void addWord(word) Adds word to the data structure, it can be matched later.
//     bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.
    
// Example:
// Input
// ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
// [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
// Output
// [null,null,null,null,false,true,true,true]
// Explanation
// WordDictionary wordDictionary = new WordDictionary();
// wordDictionary.addWord("bad");
// wordDictionary.addWord("dad");
// wordDictionary.addWord("mad");
// wordDictionary.search("pad"); // return False
// wordDictionary.search("bad"); // return True
// wordDictionary.search(".ad"); // return True
// wordDictionary.search("b.."); // return True
 
// Constraints:
//     1 <= word.length <= 25
//     word in addWord consists of lowercase English letters.
//     word in search consist of '.' or lowercase English letters.
//     There will be at most 2 dots in word for search queries.
//     At most 10^4 calls will be made to addWord and search.

import "fmt"

type Node struct {
    children []uint32
    end bool
}

type Trie struct { nodes []Node }

func (t *Trie) getChild(n uint32, k byte) uint32 {
    if t.nodes[n].children == nil { return 0 }
    return t.nodes[n].children[k]
}

func (t *Trie) addChild(n uint32, k byte) uint32 {
    if t.nodes[n].children == nil { 
        t.nodes[n].children = make([]uint32, 26) 
    }
    t.nodes[n].children[k] = uint32(len(t.nodes))
    t.nodes = append(t.nodes, Node{})
    return t.nodes[n].children[k]
}

func (t *Trie) search(n uint32, word string) bool {
    if len(word) == 0 { return t.nodes[n].end }
    if word[0] == '.' {
        for _, m := range t.nodes[n].children {
        if m > 0 && t.search(m, word[1:]) { return true }
        }
        return false
    } else {
        child := t.getChild(n, word[0] - 'a')
        if child == 0 { return false }
        return t.search(child, word[1:])
    }
}

type WordDictionary struct { tries []*Trie }

func Constructor() WordDictionary {
    tries := make([]*Trie, 26)
    for i := range tries { 
        tries[i] = &Trie{make([]Node, 1)}
    }
    return WordDictionary{tries}
}

func (this *WordDictionary) AddWord(word string) {
    trie, n := this.tries[len(word)], uint32(0)
    for i := range word {
        ch := word[i] - 'a'
        if child := trie.getChild(n, ch); child == 0 {
            n = trie.addChild(n, ch)
        } else {
            n = child
        }
    }
    trie.nodes[n].end = true
}

func (this *WordDictionary) Search(word string) bool {
    return this.tries[len(word)].search(0, word)  
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
    // WordDictionary wordDictionary = new WordDictionary();
    obj := Constructor()
    // wordDictionary.addWord("bad");
    obj.AddWord("bad")
    fmt.Println(obj)
    // wordDictionary.addWord("dad");
    obj.AddWord("dad")
    fmt.Println(obj)
    // wordDictionary.addWord("mad");
    obj.AddWord("mad")
    fmt.Println(obj)
    // wordDictionary.search("pad"); // return False
    fmt.Println(obj.Search("pad")) // false
    // wordDictionary.search("bad"); // return True
    fmt.Println(obj.Search("bad")) // true
    // wordDictionary.search(".ad"); // return True
    fmt.Println(obj.Search(".ad")) // true
    // wordDictionary.search("b.."); // return True
    fmt.Println(obj.Search("b..")) // true
}