package main

// 1032. Stream of Characters
// Design an algorithm that accepts a stream of characters and checks if a suffix of these characters is a string of a given array of strings words.

// For example, if words = ["abc", "xyz"] and the stream added the four characters (one by one) 'a', 'x', 'y', and 'z', 
// your algorithm should detect that the suffix "xyz" of the characters "axyz" matches "xyz" from words.

// Implement the StreamChecker class:
//     StreamChecker(String[] words) 
//         Initializes the object with the strings array words.
//     boolean query(char letter) 
//         Accepts a new character from the stream and returns true if any non-empty suffix from the stream forms a word that is in words.

// Example 1:
// Input
// ["StreamChecker", "query", "query", "query", "query", "query", "query", "query", "query", "query", "query", "query", "query"]
// [[["cd", "f", "kl"]], ["a"], ["b"], ["c"], ["d"], ["e"], ["f"], ["g"], ["h"], ["i"], ["j"], ["k"], ["l"]]
// Output
// [null, false, false, false, true, false, true, false, false, false, false, false, true]
// Explanation
// StreamChecker streamChecker = new StreamChecker(["cd", "f", "kl"]);
// streamChecker.query("a"); // return False
// streamChecker.query("b"); // return False
// streamChecker.query("c"); // return False
// streamChecker.query("d"); // return True, because 'cd' is in the wordlist
// streamChecker.query("e"); // return False
// streamChecker.query("f"); // return True, because 'f' is in the wordlist
// streamChecker.query("g"); // return False
// streamChecker.query("h"); // return False
// streamChecker.query("i"); // return False
// streamChecker.query("j"); // return False
// streamChecker.query("k"); // return False
// streamChecker.query("l"); // return True, because 'kl' is in the wordlist

// Constraints:
//     1 <= words.length <= 2000
//     1 <= words[i].length <= 200
//     words[i] consists of lowercase English letters.
//     letter is a lowercase English letter.
//     At most 4 * 10^4 calls will be made to query.

import "fmt"

type StreamChecker struct {
    arr []byte
    t *Trie
}

func Constructor(words []string) StreamChecker {
    t := Trie{ nodes: map[byte]*Trie{} }
    for _, w := range words {
        reverseInsert(&t, w)
    }
    return StreamChecker{[]byte{}, &t }
}

func (s *StreamChecker) Query(letter byte) bool {
    s.arr = append(s.arr, letter)
    return reverseSearch(s.t, s.arr)
}

type Trie struct {
    nodes map[byte]*Trie
    end bool
}

func reverseInsert(t *Trie, w string) {
    cur := t
    for i := len(w)-1; i >= 0; i-- {
        char := w[i]
        if cur.nodes[char] == nil {
            cur.nodes[char] = &Trie{ nodes: map[byte]*Trie{} }
        }
        cur = cur.nodes[char]
    }
    cur.end = true
}

func reverseSearch(t *Trie, w []byte) bool {
    cur := t
    for i := len(w) - 1; i >= 0; i-- {
        next, ok := cur.nodes[w[i]]
        if ok && cur.nodes[w[i]].end {
            return true
        } else if !ok {
            return false
        }
        cur = next
    }
    return false
}


type Trie1 struct {
    children [26]*Trie1
    isEnd    bool
}

func (this *Trie1) Insert(word string) {
    node := this
    for i := len(word) - 1; i >= 0; i-- {
        idx := word[i] - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &Trie1{}
        }
        node = node.children[idx]
    }
    node.isEnd = true
}

func (this *Trie1) Search(word []byte) bool {
    node := this
    for i, j := len(word)-1, 0; i >= 0 && j < 201; i, j = i-1, j+1 {
        idx := word[i] - 'a'
        if node.children[idx] == nil {
            return false
        }
        node = node.children[idx]
        if node.isEnd {
            return true
        }
    }
    return false
}

type StreamChecker1 struct {
    trie Trie1
    s    []byte
}

func Constructor1(words []string) StreamChecker1 {
    trie := Trie1{}
    for _, w := range words {
        trie.Insert(w)
    }
    return StreamChecker1{trie, []byte{}}
}

func (this *StreamChecker1) Query(letter byte) bool {
    this.s = append(this.s, letter)
    return this.trie.Search(this.s)
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */

func main() {
    // StreamChecker streamChecker = new StreamChecker(["cd", "f", "kl"]);
    obj := Constructor([]string{"cd", "f", "kl"})
    fmt.Println(obj)
    // streamChecker.query("a"); // return False
    fmt.Println(obj.Query('a')) // false
    // streamChecker.query("b"); // return False
    fmt.Println(obj.Query('b')) // false
    // streamChecker.query("c"); // return False
    fmt.Println(obj.Query('c')) // false
    // streamChecker.query("d"); // return True, because 'cd' is in the wordlist
    fmt.Println(obj.Query('d')) // true
    // streamChecker.query("e"); // return False
    fmt.Println(obj.Query('e')) // false
    // streamChecker.query("f"); // return True, because 'f' is in the wordlist
    fmt.Println(obj.Query('f')) // true
    // streamChecker.query("g"); // return False
    fmt.Println(obj.Query('g')) // false
    // streamChecker.query("h"); // return False
    fmt.Println(obj.Query('h')) // false
    // streamChecker.query("i"); // return False
    fmt.Println(obj.Query('i')) // false
    // streamChecker.query("j"); // return False
    fmt.Println(obj.Query('j')) // false
    // streamChecker.query("k"); // return False
    fmt.Println(obj.Query('k')) // false
    // streamChecker.query("l"); // return True, because 'kl' is in the wordlist
    fmt.Println(obj.Query('l')) // true

    // StreamChecker streamChecker = new StreamChecker(["cd", "f", "kl"]);
    obj1 := Constructor1([]string{"cd", "f", "kl"})
    fmt.Println(obj1)
    // streamChecker.query("a"); // return False
    fmt.Println(obj1.Query('a')) // false
    // streamChecker.query("b"); // return False
    fmt.Println(obj1.Query('b')) // false
    // streamChecker.query("c"); // return False
    fmt.Println(obj.Query('c')) // false
    // streamChecker.query("d"); // return True, because 'cd' is in the wordlist
    fmt.Println(obj1.Query('d')) // true
    // streamChecker.query("e"); // return False
    fmt.Println(obj1.Query('e')) // false
    // streamChecker.query("f"); // return True, because 'f' is in the wordlist
    fmt.Println(obj1.Query('f')) // true
    // streamChecker.query("g"); // return False
    fmt.Println(obj1.Query('g')) // false
    // streamChecker.query("h"); // return False
    fmt.Println(obj1.Query('h')) // false
    // streamChecker.query("i"); // return False
    fmt.Println(obj1.Query('i')) // false
    // streamChecker.query("j"); // return False
    fmt.Println(obj1.Query('j')) // false
    // streamChecker.query("k"); // return False
    fmt.Println(obj1.Query('k')) // false
    // streamChecker.query("l"); // return True, because 'kl' is in the wordlist
    fmt.Println(obj1.Query('l')) // true
}