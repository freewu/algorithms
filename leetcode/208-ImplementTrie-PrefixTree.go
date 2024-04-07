package main

// 208. Implement Trie (Prefix Tree)
// A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.
// Implement the Trie class:
//     Trie() Initializes the trie object.
//     void insert(String word) Inserts the string word into the trie.
//     boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
//     boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
    
// Example 1:
// Input
// ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
// [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
// Output
// [null, null, true, false, true, null, true]
// Explanation
// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple");   // return True
// trie.search("app");     // return False
// trie.startsWith("app"); // return True
// trie.insert("app");
// trie.search("app");     // return True
 
// Constraints:
//     1 <= word.length, prefix.length <= 2000
//     word and prefix consist only of lowercase English letters.
//     At most 3 * 10^4 calls in total will be made to insert, search, and startsWith.

import "fmt"

// use map
type Trie struct {
    isWord bool
    nodes map[rune]*Trie
}

func Constructor() Trie {
    return Trie {
        nodes: map[rune]*Trie{},
    }
}

func (this *Trie) Insert(word string)  {
    curr := this
    for _, r := range word {
        if _, ok := curr.nodes[r]; !ok {
            curr.nodes[r] = &Trie{nodes: map[rune]*Trie{}}
        }
        curr = curr.nodes[r]
    }
    curr.isWord = true
}

func (this *Trie) Search(word string) bool {
    curr := this
    for _, r := range word {
        if _, ok := curr.nodes[r]; !ok {
            return false
        }
        curr = curr.nodes[r]
    }
    return curr.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
    curr := this
    for _, r := range prefix {
        if _, ok := curr.nodes[r]; !ok {
            return false
        }
        curr = curr.nodes[r]
    }
    return true
}

// use array
type Trie1 struct {
    children [26]*Trie1
    isWord bool
}

func Constructor1() Trie1 {
    return Trie1 {
        children: [26]*Trie1{},
        isWord: false,
    }
}

func (this *Trie1) find(word string) *Trie1 {
    for i := range word {
        ch := word[i]
        if this.children[int(ch - 'a')] == nil {
            return nil   
        }
        this = this.children[int(ch - 'a')]
    }
    return this
}

func (this *Trie1) Insert(word string)  {
    for i := range word {
        ch := word[i]
        if this.children[int(ch - 'a')] == nil {
            this.children[int(ch - 'a')] = &Trie1{}
        }
        this = this.children[int(ch - 'a')]
    }
    this.isWord = true
}

func (this *Trie1) Search(word string) bool {
    return this.find(word) != nil && this.find(word).isWord
}

func (this *Trie1) StartsWith(prefix string) bool {
    return this.find(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
    // Trie trie = new Trie();
    obj := Constructor()
    // trie.insert("apple");
    obj.Insert("apple")
    fmt.Println(obj)
    // trie.search("apple");   // return True
    fmt.Println(obj.Search("apple")) // true
    // trie.search("app");     // return False
    fmt.Println(obj.Search("app")) // false
    // trie.startsWith("app"); // return True
    fmt.Println(obj.StartsWith("app")) // True
    // trie.insert("app");
    obj.Insert("app")
    fmt.Println(obj)
    // trie.search("app");     // return True
    fmt.Println(obj.Search("app")) // true

    // Trie trie = new Trie();
    obj1 := Constructor1()
    // trie.insert("apple");
    obj1.Insert("apple")
    fmt.Println(obj1)
    // trie.search("apple");   // return True
    fmt.Println(obj1.Search("apple")) // true
    // trie.search("app");     // return False
    fmt.Println(obj1.Search("app")) // false
    // trie.startsWith("app"); // return True
    fmt.Println(obj1.StartsWith("app")) // True
    // trie.insert("app");
    obj1.Insert("app")
    fmt.Println(obj1)
    // trie.search("app");     // return True
    fmt.Println(obj1.Search("app")) // true
}