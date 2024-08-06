package main

// LCR 062. 实现 Trie (前缀树)
// Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。
// 这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

// 请你实现 Trie 类：
//     Trie() 初始化前缀树对象。
//     void insert(String word) 向前缀树中插入字符串 word 。
//     boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
//     boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

// 示例：
// 输入
// inputs = ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
// inputs = [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
// 输出
// [null, null, true, false, true, null, true]
// 解释
// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple");   // 返回 True
// trie.search("app");     // 返回 False
// trie.startsWith("app"); // 返回 True
// trie.insert("app");
// trie.search("app");     // 返回 True
 
// 提示：
//     1 <= word.length, prefix.length <= 2000
//     word 和 prefix 仅由小写英文字母组成
//     insert、search 和 startsWith 调用次数 总计 不超过 3 * 10^4 次

import "fmt"

// use map
type Trie struct {
    isWord bool
    nodes map[rune]*Trie
}

func Constructor() Trie {
    return Trie { nodes: map[rune]*Trie{}, }
}

func (this *Trie) Insert(word string)  {
    curr := this
    for _, r := range word {
        if _, ok := curr.nodes[r]; !ok {
            curr.nodes[r] = &Trie{ nodes: map[rune]*Trie{} }
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