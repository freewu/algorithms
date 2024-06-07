package main

// 648. Replace Words
// In English, we have a concept called root, which can be followed by some other word to form another longer word - let's call this word derivative. 
// For example, when the root "help" is followed by the word "ful", we can form a derivative "helpful".

// Given a dictionary consisting of many roots and a sentence consisting of words separated by spaces, 
// replace all the derivatives in the sentence with the root forming it. 
// If a derivative can be replaced by more than one root, replace it with the root that has the shortest length.

// Return the sentence after the replacement.

// Example 1:
// Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
// Output: "the cat was rat by the bat"

// Example 2:
// Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
// Output: "a a b c"

// Constraints:
//     1 <= dictionary.length <= 1000
//     1 <= dictionary[i].length <= 100
//     dictionary[i] consists of only lower-case letters.
//     1 <= sentence.length <= 10^6
//     sentence consists of only lower-case letters and spaces.
//     The number of words in sentence is in the range [1, 1000]
//     The length of each word in sentence is in the range [1, 1000]
//     Every two consecutive words in sentence will be separated by exactly one space.
//     sentence does not have leading or trailing spaces.

import "fmt"
import "strings"

type TrieNode struct {
    childrens [26]*TrieNode
    isWordEnd bool
}

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{
        root: &TrieNode{},
    }
}

func (this *Trie) insert(word string) {
    n, current := len(word), this.root
    for i := 0; i < n; i++ {
        index := word[i] - 'a'
        if current.childrens[index] == nil {
            current.childrens[index] = &TrieNode{}
        }
        current = current.childrens[index]
    }
    current.isWordEnd = true
}

func (this *Trie) find(word string) (bool, int) {
    n, current := len(word), this.root
    for i := 0; i < n; i++ {
        index := word[i] - 'a'
        if current.childrens[index] == nil {
            return false, -1
        }
        current = current.childrens[index]
        if current.isWordEnd {
            return true, i
        }
    }
    return false, -1
}

func replaceWords(dictionary []string, sentence string) string {
    t := NewTrie()
    for _, word := range dictionary {
        t.insert(word)
    }
    words := strings.Split(sentence, " ")
    for i, word := range words {
        found, index := t.find(word)
        if found {
            words[i] = word[:index+1]
        }
    }
    return strings.Join(words, " ")
}

func main() {
    // Example 1:
    // Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
    // Output: "the cat was rat by the bat"
    fmt.Println(replaceWords([]string{"cat","bat","rat"},"the cattle was rattled by the battery")) // "the cat was rat by the bat"
    // Example 2:
    // Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
    // Output: "a a b c"
    fmt.Println(replaceWords([]string{"a","b","c"}, "aadsfasf absbs bbab cadsfafs")) // "a a b c"
}