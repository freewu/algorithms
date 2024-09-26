package main

// 1804. Implement Trie II (Prefix Tree)
// A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. 
// There are various applications of this data structure, such as autocomplete and spellchecker.

// Implement the Trie class:
//     Trie() 
//         Initializes the trie object.
//     void insert(String word) 
//         Inserts the string word into the trie.
//     int countWordsEqualTo(String word) 
//         Returns the number of instances of the string word in the trie.
//     int countWordsStartingWith(String prefix) 
//         Returns the number of strings in the trie that have the string prefix as a prefix.
//     void erase(String word) 
//         Erases the string word from the trie.

// Example 1:
// Input
// ["Trie", "insert", "insert", "countWordsEqualTo", "countWordsStartingWith", "erase", "countWordsEqualTo", "countWordsStartingWith", "erase", "countWordsStartingWith"]
// [[], ["apple"], ["apple"], ["apple"], ["app"], ["apple"], ["apple"], ["app"], ["apple"], ["app"]]
// Output
// [null, null, null, 2, 2, null, 1, 1, null, 0]
// Explanation
// Trie trie = new Trie();
// trie.insert("apple");               // Inserts "apple".
// trie.insert("apple");               // Inserts another "apple".
// trie.countWordsEqualTo("apple");    // There are two instances of "apple" so return 2.
// trie.countWordsStartingWith("app"); // "app" is a prefix of "apple" so return 2.
// trie.erase("apple");                // Erases one "apple".
// trie.countWordsEqualTo("apple");    // Now there is only one instance of "apple" so return 1.
// trie.countWordsStartingWith("app"); // return 1
// trie.erase("apple");                // Erases "apple". Now the trie is empty.
// trie.countWordsStartingWith("app"); // return 0

// Constraints:
//     1 <= word.length, prefix.length <= 2000
//     word and prefix consist only of lowercase English letters.
//     At most 3 * 10^4 calls in total will be made to insert, countWordsEqualTo, countWordsStartingWith, and erase.
//     It is guaranteed that for any function call to erase, the string word will exist in the trie.

import "fmt"

type Trie struct {
    children [26]*Trie
    pre int
    end int
}

func Constructor() Trie { 
    return Trie{}
}

func (t *Trie) Insert(s string) {
    o := t
    for _, b := range s {
        b -= 'a'
        if o.children[b] == nil {
            o.children[b] = &Trie{}
        }
        o = o.children[b]
        o.pre++
    }
    o.end++
}

func (t *Trie) CountWordsEqualTo(s string) int {
    o := t
    for _, b := range s {
        o = o.children[b-'a']
        if o == nil {
            return 0
        }
    }
    return o.end
}

func (t *Trie) CountWordsStartingWith(s string) int {
    o := t
    for _, b := range s {
        o = o.children[b-'a']
        if o == nil {
            return 0
        }
    }
    return o.pre
}

func (t *Trie) Erase(s string) {
    o := t
    for _, b := range s {
        o = o.children[b-'a']
        o.pre--
    }
    o.end--
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.CountWordsEqualTo(word);
 * param_3 := obj.CountWordsStartingWith(prefix);
 * obj.Erase(word);
 */

func main() {
    // Trie trie = new Trie();
    obj := Constructor()
    fmt.Println(obj)
    // trie.insert("apple");               // Inserts "apple".
    obj.Insert("apple")
    fmt.Println(obj)
    // trie.insert("apple");               // Inserts another "apple".
    obj.Insert("apple")
    fmt.Println(obj)
    // trie.countWordsEqualTo("apple");    // There are two instances of "apple" so return 2.
    fmt.Println(obj.CountWordsEqualTo("apple")) // 2
    // trie.countWordsStartingWith("app"); // "app" is a prefix of "apple" so return 2.
    fmt.Println(obj.CountWordsEqualTo("app")) // 2
    // trie.erase("apple");                // Erases one "apple".
    obj.Erase("apple")
    fmt.Println(obj)
    // trie.countWordsEqualTo("apple");    // Now there is only one instance of "apple" so return 1.
    fmt.Println(obj.CountWordsEqualTo("apple")) // 1
    // trie.countWordsStartingWith("app"); // return 1
    fmt.Println(obj.CountWordsEqualTo("app")) // 1
    // trie.erase("apple");                // Erases "apple". Now the trie is empty.
    obj.Erase("apple")
    fmt.Println(obj)
    // trie.countWordsStartingWith("app"); // return 0
    fmt.Println(obj.CountWordsEqualTo("apple")) // 0
    // trie.countWordsStartingWith("app"); // return 0
    fmt.Println(obj.CountWordsEqualTo("app")) // 0
}