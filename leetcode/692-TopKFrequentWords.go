package main

// 692. Top K Frequent Words
// Given an array of strings words and an integer k, return the k most frequent strings.
// Return the answer sorted by the frequency from highest to lowest. 
// Sort the words with the same frequency by their lexicographical order.

// Example 1:
// Input: words = ["i","love","leetcode","i","love","coding"], k = 2
// Output: ["i","love"]
// Explanation: "i" and "love" are the two most frequent words.
// Note that "i" comes before "love" due to a lower alphabetical order.

// Example 2:
// Input: words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
// Output: ["the","is","sunny","day"]
// Explanation: "the", "is", "sunny" and "day" are the four most frequent words, with the number of occurrence being 4, 3, 2 and 1 respectively.

// Constraints:
//     1 <= words.length <= 500
//     1 <= words[i].length <= 10
//     words[i] consists of lowercase English letters.
//     k is in the range [1, The number of unique words[i]]
 
// Follow-up: Could you solve it in O(n log(k)) time and O(n) extra space?

import "fmt"
import "sort"
import "container/heap"

func topKFrequent(words []string, k int) []string {
    freq, uniq := make(map[string]int), []string{}
    for _, w := range words {
        if _, ok := freq[w]; !ok { // 第一次出现加到数组中用于排序
            uniq = append(uniq, w)
        }
        freq[w]++ // 单词出现频次
    }
    sort.SliceStable(uniq, func(i, j int) bool {
        if freq[uniq[i]] > freq[uniq[j]] {
            return true
        }
        return freq[uniq[i]] == freq[uniq[j]] && uniq[i] < uniq[j]
     })
    return uniq[:k]
}

type Item struct{
    word string
    cnt int
}

type ItemHeap []Item

func (w ItemHeap) Len() int {
    return len(w)
}

func (w ItemHeap) Less(i,j int) bool {
    if w[i].cnt != w[j].cnt{
        return w[i].cnt > w[j].cnt
    }
    return w[i].word < w[j].word
}

func (w ItemHeap) Swap(i, j int) {
    w[i], w[j] = w[j], w[i]
}

func (w *ItemHeap) Pop() interface{} {
    x := (*w)[len(*w)-1]
    *w = (*w)[:len(*w)-1]
    return x
}

func (w *ItemHeap) Push(val interface{}) {
    *w = append(*w, val.(Item))
}

func topKFrequent1(words []string, k int) []string {
    mp := map[string]int{}
    for i := range words{
        mp[words[i]]++
    }
    h := &ItemHeap{}
    for key, value := range mp {
        heap.Push(h, Item{key, value})
    }
    res := make([]string, k)
    for i := 0; i < k; i++{
        res[i] = heap.Pop(h).(Item).word
    }
    return res
}

func main() {
    // Example 1:
    // Input: words = ["i","love","leetcode","i","love","coding"], k = 2
    // Output: ["i","love"]
    // Explanation: "i" and "love" are the two most frequent words.
    // Note that "i" comes before "love" due to a lower alphabetical order.
    fmt.Println(topKFrequent([]string{"i","love","leetcode","i","love","coding"}, 2)) // ["i","love"]
    // Example 2:
    // Input: words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
    // Output: ["the","is","sunny","day"]
    // Explanation: "the", "is", "sunny" and "day" are the four most frequent words, with the number of occurrence being 4, 3, 2 and 1 respectively.
    fmt.Println(topKFrequent([]string{"the","day","is","sunny","the","the","the","sunny","is","is"}, 4)) // ["the","is","sunny","day"]
    
    fmt.Println(topKFrequent1([]string{"i","love","leetcode","i","love","coding"}, 2)) // ["i","love"]
    fmt.Println(topKFrequent1([]string{"the","day","is","sunny","the","the","the","sunny","is","is"}, 4)) // ["the","is","sunny","day"]
}