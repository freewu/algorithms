package main

// 3081. Replace Question Marks in String to Minimize Its Value
// You are given a string s. s[i] is either a lowercase English letter or '?'.

// For a string t having length m containing only lowercase English letters, 
// we define the function cost(i) for an index i as the number of characters equal to t[i] that appeared before it, 
// i.e. in the range [0, i - 1].

// The value of t is the sum of cost(i) for all indices i.

// For example, for the string t = "aab":
//     1. cost(0) = 0
//     2. cost(1) = 1
//     3. cost(2) = 0
//     4. Hence, the value of "aab" is 0 + 1 + 0 = 1.

// Your task is to replace all occurrences of '?' in s with any lowercase English letter so that the value of s is minimized.

// Return a string denoting the modified string with replaced occurrences of '?'. 
// If there are multiple strings resulting in the minimum value, return the lexicographically smallest one.

// Example 1:
// Input:  s = "???" 
// Output:  "abc" 
// Explanation: In this example, we can replace the occurrences of '?' to make s equal to "abc".
// For "abc", cost(0) = 0, cost(1) = 0, and cost(2) = 0.
// The value of "abc" is 0.
// Some other modifications of s that have a value of 0 are "cba", "abz", and, "hey".
// Among all of them, we choose the lexicographically smallest.

// Example 2:
// Input: s = "a?a?"
// Output: "abac"
// Explanation: In this example, the occurrences of '?' can be replaced to make s equal to "abac".
// For "abac", cost(0) = 0, cost(1) = 0, cost(2) = 1, and cost(3) = 0.
// The value of "abac" is 1.

// Constraints:
//     1 <= s.length <= 10^5
//     s[i] is either a lowercase English letter or '?'.

import "fmt"
import "container/heap"
import "sort"

type Pair struct {
    Count int
    Char rune
}

type MinHeap []Pair
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { 
    if h[i].Count == h[j].Count { return h[i].Char < h[j].Char }
    return h[i].Count < h[j].Count 
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *MinHeap) Pop() interface{} {
    tmp, n := *h, len(*h)
    res := tmp[n - 1]
    *h = tmp[0 : n - 1]
    return res
}

func minimizeStringValue(s string) string {
    mp := make(map[rune]int)
    for _, v := range s {
        if v != '?' {
            mp[v]++
        }
    }
    pq := make(MinHeap, 0)
    heap.Init(&pq)
    for i := 0; i < 26; i++ {
        ch := rune(int('a') + i)
        heap.Push(&pq, Pair{mp[ch], ch})
    }
    memo := make([]int, 26)
    for _, v := range s {
        if v == '?' {
            it := heap.Pop(&pq).(Pair)
            ch := it.Char
            mp[ch]++
            heap.Push(&pq, Pair{mp[ch], ch})
            memo[int(ch) - int('a')]++
        }
    }
    cur, res := 0, make([]rune, len(s))
    for i, v := range s {
        for cur < 26 && memo[cur] == 0 {
            cur++
        }
        if v == '?' {
            res[i] = rune(int('a') + cur)
            memo[cur]--
        } else {
            res[i] = rune(v)
        }
    }
    return string(res)
}

func minimizeStringValue1(s string) string {
    count, n, mp, bytes, arr := 0, len(s), make([]int, 26), []byte(s), []byte{}
    for i := 0; i < n; i++ {
        if bytes[i] == '?' {
            count++
            continue
        }
        mp[bytes[i]-'a']++
    }
    findMinNum := func (arr []int) int {
        res := 0
        for i, v := range arr {
            if v < arr[res] { res = i }
        }
        return res
    }
    for j := 0; j < count; j++ {
        index := findMinNum(mp)
        mp[index]++
        arr = append(arr, byte('a' + index))
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    i := 0
    for i < n {
        if bytes[i] != '?' {
            i++
            continue
        }
        bytes[i] = arr[0]
        arr = arr[1:]
    }
    return string(bytes)
}

func main() {
    // Example 1:
    // Input:  s = "???" 
    // Output:  "abc" 
    // Explanation: In this example, we can replace the occurrences of '?' to make s equal to "abc".
    // For "abc", cost(0) = 0, cost(1) = 0, and cost(2) = 0.
    // The value of "abc" is 0.
    // Some other modifications of s that have a value of 0 are "cba", "abz", and, "hey".
    // Among all of them, we choose the lexicographically smallest.
    fmt.Println(minimizeStringValue("???")) // "abc"
    // Example 2:
    // Input: s = "a?a?"
    // Output: "abac"
    // Explanation: In this example, the occurrences of '?' can be replaced to make s equal to "abac".
    // For "abac", cost(0) = 0, cost(1) = 0, cost(2) = 1, and cost(3) = 0.
    // The value of "abac" is 1.
    fmt.Println(minimizeStringValue("a?a?")) // "abac"

    fmt.Println(minimizeStringValue1("???")) // "abc"
    fmt.Println(minimizeStringValue1("a?a?")) // "abac"
}