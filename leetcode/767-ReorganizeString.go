package main

// 767. Reorganize String
// Given a string s, rearrange the characters of s so that any two adjacent characters are not the same.
// Return any possible rearrangement of s or return "" if not possible.

// Example 1:
// Input: s = "aab"
// Output: "aba"

// Example 2:
// Input: s = "aaab"
// Output: ""

// Constraints:
//     1 <= s.length <= 500
//     s consists of lowercase English letters.

import "fmt"
import "container/heap"
import "strings"
import "sort"

type CharFreq struct {
    char  rune
    count int
}

type MaxHeap []CharFreq
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(CharFreq))
}
func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func reorganizeString(s string) string {
    freqMap := make(map[rune]int)
    for _, c := range s {
        freqMap[c]++
    }
    maxHeap := &MaxHeap{}
    heap.Init(maxHeap)
    for c, freq := range freqMap {
        heap.Push(maxHeap, CharFreq{c, freq})
    }
    var res strings.Builder
    for maxHeap.Len() >= 2 {
        charFreq1 := heap.Pop(maxHeap).(CharFreq)
        charFreq2 := heap.Pop(maxHeap).(CharFreq)
        res.WriteRune(charFreq1.char)
        res.WriteRune(charFreq2.char)
        if charFreq1.count > 1 {
            heap.Push(maxHeap, CharFreq{charFreq1.char, charFreq1.count - 1})
        }
        if charFreq2.count > 1 {
            heap.Push(maxHeap, CharFreq{charFreq2.char, charFreq2.count - 1})
        }
    }
    if maxHeap.Len() > 0 {
        lastFreq := heap.Pop(maxHeap).(CharFreq)
        if lastFreq.count > 1 {
            return ""
        }
        res.WriteRune(lastFreq.char)
    }
    return res.String()
}

func reorganizeString1(s string) string {
    type node struct {
        b byte
        cnt int
    }
    res, mp, lst, mx := []byte(s), make(map[byte]int), []node{}, 0
    for i := 0; i < len(s); i++ {
        mp[s[i]]++
        if mp[s[i]] > mx {
            mx = mp[s[i]]
        }
    }
    if len(s) - mx < mx - 1 {
        return ""
    }
    for k, v := range mp {
        n := node{k, v}
        lst = append(lst, n)
    }
    sort.Slice(lst, func (i, j int) bool {
        return lst[i].cnt > lst[j].cnt
    })
    idx := 0
    for _, n := range lst {
        v := n.cnt
        key := n.b
        for v > 0 {
            res[idx] = key
            idx += 2
            if idx >= len(s) {
                idx = 1
            }
            v--
        }
    }
    return string(res)
}

func main() {
    // Example 1:
    // Input: s = "aab"
    // Output: "aba"
    fmt.Println(reorganizeString("aab")) // "aba"
    // Example 2:
    // Input: s = "aaab"
    // Output: ""
    fmt.Println(reorganizeString("aaab")) // ""

    fmt.Println(reorganizeString1("aab")) // "aba"
    fmt.Println(reorganizeString1("aaab")) // ""
}