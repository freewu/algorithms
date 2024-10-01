package main

// 1405. Longest Happy String
// A string s is called happy if it satisfies the following conditions:
//     s only contains the letters 'a', 'b', and 'c'.
//     s does not contain any of "aaa", "bbb", or "ccc" as a substring.
//     s contains at most a occurrences of the letter 'a'.
//     s contains at most b occurrences of the letter 'b'.
//     s contains at most c occurrences of the letter 'c'.

// Given three integers a, b, and c, return the longest possible happy string. 
// If there are multiple longest happy strings, return any of them. 
// If there is no such string, return the empty string "".

// A substring is a contiguous sequence of characters within a string.

// Example 1:
// Input: a = 1, b = 1, c = 7
// Output: "ccaccbcc"
// Explanation: "ccbccacc" would also be a correct answer.

// Example 2:
// Input: a = 7, b = 1, c = 0
// Output: "aabaa"
// Explanation: It is the only correct answer in this case.

// Constraints:
//     0 <= a, b, c <= 100
//     a + b + c > 0

import "fmt"
import "container/heap"

type Data struct {
    Count  int
    Symbol rune
}

type IntHeap []Data
func (h IntHeap)  Len() int           { return len(h) }
func (h IntHeap)  Less(i, j int) bool { return h[i].Count > h[j].Count }
func (h IntHeap)  Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(Data)) }
func (h *IntHeap) Pop() interface{}   { 
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func longestDiverseString(a int, b int, c int) string {
    hp := IntHeap{}
    if a > 0 { hp = append(hp, Data{Count: a, Symbol: 'a'}) }
    if b > 0 { hp = append(hp, Data{Count: b, Symbol: 'b'}) }
    if c > 0 { hp = append(hp, Data{Count: c, Symbol: 'c'}) }
    heap.Init(&hp)
    res, last, prev := "", 'd', 'd'
    for len(hp) > 0 {
        cur := heap.Pop(&hp).(Data)
        if cur.Symbol == last && prev == last {
            if len(hp) > 0 {
                newCur := heap.Pop(&hp).(Data)
                res += string(newCur.Symbol)
                newCur.Count--
                if newCur.Count > 0 {
                    heap.Push(&hp, Data{Count: newCur.Count, Symbol: newCur.Symbol})
                }
                heap.Push(&hp, Data{Count: cur.Count, Symbol: cur.Symbol})
                prev = last
                last = newCur.Symbol
            }
        } else {
            res += string(cur.Symbol)
            cur.Count--
            if cur.Count > 0 {
                heap.Push(&hp, Data{Count: cur.Count, Symbol: cur.Symbol})
            }
            prev = last
            last = cur.Symbol
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: a = 1, b = 1, c = 7
    // Output: "ccaccbcc"
    // Explanation: "ccbccacc" would also be a correct answer.
    fmt.Println(longestDiverseString(1,1,7)) // "ccaccbcc"
    // Example 2:
    // Input: a = 7, b = 1, c = 0
    // Output: "aabaa"
    // Explanation: It is the only correct answer in this case.
    fmt.Println(longestDiverseString(7,1,0)) // "aabaa"
}