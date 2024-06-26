package main

// 676. Implement Magic Dictionary
// Design a data structure that is initialized with a list of different words. 
// Provided a string, you should determine if you can change exactly one character in this string to match any word in the data structure.

// Implement the MagicDictionary class:
//     MagicDictionary() Initializes the object.
//     void buildDict(String[] dictionary) Sets the data structure with an array of distinct strings dictionary.
//     bool search(String searchWord) Returns true if you can change exactly one character in searchWord to match any string in the data structure, otherwise returns false.

// Example 1:
// Input
// ["MagicDictionary", "buildDict", "search", "search", "search", "search"]
// [[], [["hello", "leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]
// Output
// [null, null, false, true, false, false]
// Explanation
// MagicDictionary magicDictionary = new MagicDictionary();
// magicDictionary.buildDict(["hello", "leetcode"]);
// magicDictionary.search("hello"); // return False
// magicDictionary.search("hhllo"); // We can change the second 'h' to 'e' to match "hello" so we return True
// magicDictionary.search("hell"); // return False
// magicDictionary.search("leetcoded"); // return False

// Constraints:
//     1 <= dictionary.length <= 100
//     1 <= dictionary[i].length <= 100
//     dictionary[i] consists of only lower-case English letters.
//     All the strings in dictionary are distinct.
//     1 <= searchWord.length <= 100
//     searchWord consists of only lower-case English letters.
//     buildDict will be called only once before search.
//     At most 100 calls will be made to search.

import "fmt"

type MagicDictionary struct {
    children [26]*MagicDictionary
    isEnd    bool
}

func Constructor() MagicDictionary {
    return MagicDictionary{}
}

func (m *MagicDictionary) BuildDict(dictionary []string) {
    for _, word := range dictionary {
        m.insert(word)
    }
}

func (m *MagicDictionary) insert(word string) {
    cur := m
    for _, ch := range word {
        if cur.children[ch-'a'] == nil {
            cur.children[ch-'a'] = &MagicDictionary{}
        }
        cur = cur.children[ch-'a']
    }
    cur.isEnd = true
}

func (m *MagicDictionary) Search(searchWord string) bool {
	return dfs(m, searchWord, 0, 1)
}

func dfs(r *MagicDictionary, w string, i int, limit int) bool {
    // base case
    if limit < 0 {
        return false
    }
    if i == len(w) {
        return r.isEnd && limit == 0
    }
    ch := w[i] - 'a' 
    for c, t := range r.children { // iterate current node's all children
        if t == nil {
            continue
        }
        if c == int(ch) && dfs(t, w, i+1, limit) { // c == ch, represent don't need change.
            return true
        }
        if c != int(ch) && dfs(t, w, i+1, limit-1) { // c != ch, represent consume one chance for change
            return true
        }
    }
    return false
}

// type MagicDictionary struct {
//     dicts []string
// }

// func Constructor() MagicDictionary {
//     return MagicDictionary{dicts: []string{}}
// }

// func (this *MagicDictionary) BuildDict(dictionary []string) {
//     this.dicts = dictionary
// }

// func (this *MagicDictionary) Search(searchWord string) bool {
// NEXT:
//     for _, word := range this.dicts {
//         if len(word) != len(searchWord) {
//             continue
//         }
//         hasDiff := false
//         for i := 0; i < len(word); i++ {
//             if word[i] != searchWord[i] {
//                 if hasDiff {
//                     continue NEXT
//                 }
//                 hasDiff = true
//             }
//         }
//         if hasDiff {
//             return true
//         }
//     }
//     return false
// }

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

func main() {
    // MagicDictionary magicDictionary = new MagicDictionary();
    obj := Constructor()
    fmt.Println(obj)
    // magicDictionary.buildDict(["hello", "leetcode"]);
    obj.BuildDict([]string{"hello", "leetcode"})
    fmt.Println(obj)
    // magicDictionary.search("hello"); // return False
    fmt.Println(obj.Search("hello")) // false
    // magicDictionary.search("hhllo"); // We can change the second 'h' to 'e' to match "hello" so we return True
    fmt.Println(obj.Search("hhllo")) // true
    // magicDictionary.search("hell"); // return False
    fmt.Println(obj.Search("hell")) // false
    // magicDictionary.search("leetcoded"); // return False
    fmt.Println(obj.Search("leetcoded")) // false
}