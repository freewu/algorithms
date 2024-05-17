package main

// 642. Design Search Autocomplete System
// Design a search autocomplete system for a search engine. 
// Users may input a sentence (at least one word and end with a special character '#').

// You are given a string array sentences and an integer array times both of length n where sentences[i] is a previously typed sentence and times[i] is the corresponding number of times the sentence was typed. 
// For each input character except '#', return the top 3 historical hot sentences that have the same prefix as the part of the sentence already typed.

// Here are the specific rules:
//     The hot degree for a sentence is defined as the number of times a user typed the exactly same sentence before.
//     The returned top 3 hot sentences should be sorted by hot degree (The first is the hottest one). If several sentences have the same hot degree, use ASCII-code order (smaller one appears first).
//     If less than 3 hot sentences exist, return as many as you can.
//     When the input is a special character, it means the sentence ends, and in this case, you need to return an empty list.

// Implement the AutocompleteSystem class:
//     AutocompleteSystem(String[] sentences, int[] times) 
//         Initializes the object with the sentences and times arrays.
//     List<String> input(char c) 
//         This indicates that the user typed the character c.
//         Returns an empty array [] if c == '#' and stores the inputted sentence in the system.
//         Returns the top 3 historical hot sentences that have the same prefix as the part of the sentence already typed. 
//         If there are fewer than 3 matches, return them all.

// Example 1:
// Input
// ["AutocompleteSystem", "input", "input", "input", "input"]
// [[["i love you", "island", "iroman", "i love leetcode"], [5, 3, 2, 2]], ["i"], [" "], ["a"], ["#"]]
// Output
// [null, ["i love you", "island", "i love leetcode"], ["i love you", "i love leetcode"], [], []]
// Explanation
// AutocompleteSystem obj = new AutocompleteSystem(["i love you", "island", "iroman", "i love leetcode"], [5, 3, 2, 2]);
// obj.input("i"); // return ["i love you", "island", "i love leetcode"]. There are four sentences that have prefix "i". Among them, "ironman" and "i love leetcode" have same hot degree. Since ' ' has ASCII code 32 and 'r' has ASCII code 114, "i love leetcode" should be in front of "ironman". Also we only need to output top 3 hot sentences, so "ironman" will be ignored.
// obj.input(" "); // return ["i love you", "i love leetcode"]. There are only two sentences that have prefix "i ".
// obj.input("a"); // return []. There are no sentences that have prefix "i a".
// obj.input("#"); // return []. The user finished the input, the sentence "i a" should be saved as a historical sentence in system. And the following input will be counted as a new search.

// Constraints:
//     n == sentences.length
//     n == times.length
//     1 <= n <= 100
//     1 <= sentences[i].length <= 100
//     1 <= times[i] <= 50
//     c is a lowercase English letter, a hash '#', or space ' '.
//     Each tested sentence will be a sequence of characters c that end with the character '#'.
//     Each tested sentence will have a length in the range [1, 200].
//     The words in each input sentence are separated by single spaces.
//     At most 5000 calls will be made to input.

import "fmt"
// import "strings"
// import "sort"

// func Constructor(sentences []string, times []int) AutocompleteSystem {
//     as := AutocompleteSystem{
//         root: &Node{
//             ss: map[string]*Ele{
//                 "": &Ele{},
//             },
//             nexts: map[byte]*Node{},
//         },
//     }
//     as.pointer = as.root
//     for i := 0; i < len(sentences); i++ {
//         as.Add(sentences[i], times[i])
//     }
//     return as
// }

// type AutocompleteSystem struct { // 我们的 Trie
//     root *Node  // 根节点
//     process []byte // 当前输入中的单词，已经处理的字符
//     pointer *Node  // 当前已经处理到的 node
// }

// type Node struct { // Trie 节点
//     prefix string  // 节点对应的前缀
//     ss map[string]*Ele  // 节点对应的后续句子的个数
//     nexts map[byte]*Node  // 节点的后续节点
// }

// type Ele struct { // 为方便排序引入的句子数量统计结构
//     val string
//     cnt int
// }

// func (as *AutocompleteSystem) Add(s string, count int) {
//     node := as.root
//     idx := 0
//     for idx < len(s) {
//         _, has := node.ss[s[idx:]]
//         if has {
//             node.ss[s[idx:]].cnt += count
//         } else {
//             node.ss[s[idx:]] = &Ele{
//                 val: s[idx:],
//                 cnt: count,
//             }
//         }
//         nextNode, has := node.nexts[s[idx]]
//         if !has {
//             nextNode = &Node{
//                 prefix: s[:idx] + string(s[idx]),
//                 ss: map[string]*Ele{},
//                 nexts: map[byte]*Node{},
//             }
//             node.nexts[s[idx]] = nextNode
//         }
//         node = nextNode
//         idx += 1
//     }
//     _, has := node.ss[s[idx:]]
//     if has {
//         node.ss[s[idx:]].cnt += count
//     } else {
//         node.ss[s[idx:]] = &Ele{
//             val: s[idx:],
//             cnt: count,
//         }
//     }
// }

// func (as *AutocompleteSystem) Input(c byte) []string {
//     if c == '#' { // 如果结束
//         as.Add(string(as.process), 1)
//         as.process = as.process[:0]
//         as.pointer = as.root
//         return []string{}
//     }
//     // 如果不是 #，应当收集字符
//     as.process = append(as.process, c) // 收集字符

//     // 如果没有结束
//     // 当前指针为 nil,说明之前已经没有匹配的内容了，直接返回
//     node := as.pointer // 当前 node 所在位置
//     if node == nil {
//         return []string{}
//     }
//     if node.nexts[c] == nil { // 当前位置没有对于 c 的匹配
//         as.pointer = nil // 将 pointer 置 nil
//         return []string{}
//     }
//     node = node.nexts[c] // 当前有匹配
//     toSort := []*Ele{}
//     for _, v := range node.ss { // 找出前三个
//         toSort = append(toSort, v)
//     }
//     sort.Slice(toSort, func(i, j int) bool {
//         // 如果数量相等
//         if toSort[i].cnt == toSort[j].cnt {
//             // 比较两个字符串
//             compareRet := strings.Compare(toSort[i].val, toSort[j].val)
//             return compareRet <= 0
//         }
//         return toSort[i].cnt > toSort[j].cnt
//     })
//     res := []string{}
//     if len(toSort) >= 3 {
//         toSort = toSort[:3]
//     }
//     for i := 0; i < len(toSort); i++ {
//         res = append(res, node.prefix + toSort[i].val)
//     }
//     as.pointer = node
//     return res
// }

import "sort"
import "strings"

type AutocompleteSystem struct {
    input string
    root *trie
}

type trie struct {
    next [27]*trie
    end bool
    cnt int
    content string
}

func (this *trie) add (sentence string, cnt int) {
    cur := this
    for i := range sentence {
        var v int
        if sentence[i] == ' ' {
            v = 26
        }else {
            v = int(sentence[i]-'a')
        }
        if cur.next[v] == nil {
            cur.next[v] = &trie{
                next: [27]*trie{},
            }
        }
        cur = cur.next[v]
    }
    cur.end, cur.cnt, cur.content = true, cur.cnt+cnt, sentence
}

func (this *trie) find(sentence string) (*trie, bool) {
    cur := this
    for i := range sentence {
        var v int
        if sentence[i] == ' ' {
            v = 26
        }else {
            v = int(sentence[i]-'a')
        }
        if cur.next[v] == nil {
            return nil, false
        }else {
            cur = cur.next[v]
        }
    }
    return cur, true
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
    root := &trie{
        next: [27]*trie{},
    }
    for i := range sentences {
        root.add(sentences[i], times[i])
    }
    return AutocompleteSystem{
        input: "",
        root: root,
    }
}


func (this *AutocompleteSystem) Input(c byte) []string {
    if c == '#' {
        this.root.add(this.input, 1)
        this.input = ""
        return []string{}
    }
    this.input += string(c)
    start, ok := this.root.find(this.input)
    if !ok {return []string{}}
    var dfs func(cur *trie)
    
    arr := make([]*trie, 0)
    dfs = func(cur *trie) {
        if cur.end {
            arr = append(arr, cur)
        }
        for i := range cur.next {
            if cur.next[i] != nil {
                dfs(cur.next[i])
            }
        }
    }
    dfs(start)
    sort.Slice(arr, func(i, j int) bool {
        if arr[i].cnt != arr[j].cnt {
            return arr[i].cnt > arr[j].cnt
        }
        n := strings.Compare(arr[i].content, arr[j].content)
        return n < 0
    })
    res := make([]string, 0)
    for _, str := range arr {
        res = append(res, str.content)
    }
    if len(res) < 3 {
        return res
    }else {
        return res[:3]
    }
}

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */

func main() {
    // AutocompleteSystem obj = new AutocompleteSystem(["i love you", "island", "iroman", "i love leetcode"], [5, 3, 2, 2]);
    obj := Constructor([]string{"i love you", "island", "iroman", "i love leetcode"},[]int{5, 3, 2, 2})
    // obj.input("i"); // return ["i love you", "island", "i love leetcode"]. There are four sentences that have prefix "i". Among them, "ironman" and "i love leetcode" have same hot degree. Since ' ' has ASCII code 32 and 'r' has ASCII code 114, "i love leetcode" should be in front of "ironman". Also we only need to output top 3 hot sentences, so "ironman" will be ignored.
    fmt.Println(obj.Input('i')) // ["i love you", "island", "i love leetcode"]
    // obj.input(" "); // return ["i love you", "i love leetcode"]. There are only two sentences that have prefix "i ".
    fmt.Println(obj.Input(' ')) // ["i love you", "i love leetcode"]
    // obj.input("a"); // return []. There are no sentences that have prefix "i a".
    fmt.Println(obj.Input('a')) // []
    // obj.input("#"); // return []. The user finished the input, the sentence "i a" should be saved as a historical sentence in system. And the following input will be counted as a new search.
    fmt.Println(obj.Input('#')) // []
}