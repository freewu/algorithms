package main

// 358. Rearrange String k Distance Apart
// Given a string s and an integer k, rearrange s such that the same characters are at least distance k from each other. 
// If it is not possible to rearrange the string, return an empty string "".

// Example 1:
// Input: s = "aabbcc", k = 3
// Output: "abcabc"
// Explanation: The same letters are at least a distance of 3 from each other.

// Example 2:
// Input: s = "aaabc", k = 3
// Output: ""
// Explanation: It is not possible to rearrange the string.

// Example 3:
// Input: s = "aaadbbcc", k = 2
// Output: "abacabcd"
// Explanation: The same letters are at least a distance of 2 from each other.

// Constraints:
//     1 <= s.length <= 3 * 10^5
//     s consists of only lowercase English letters.
//     0 <= k <= s.length

import "fmt"
import "container/heap"

// 最大堆
func rearrangeString(s string, k int) string {
    if k== 0 {
        return s
    }
    m := [26]uint16{}
    for i := range s { // 统计字符出现的次数
        m[s[i]-'a']++
    }
    h := MaxHeap{}
    for c,cnt := range m { // 维护大顶堆
        if cnt > 0 {
            heap.Push(&h, Pair{ c:byte( c + 'a'), cnt: cnt })
        }
    }
    if h.Top().cnt > uint16((len(s)+1/k)) {
        return ""
    }
    res, list := make([]byte, 0, len(s)), make([]Pair, k) // 缓存数组
    for h.Len() > 0 { // 每次取前k个
        if h.Len() < k && h.Top().cnt > 1 { // 如果长度不够k，则不能完成任务
            return ""
        }
        pos := 0
        for i := 0; i < k && h.Len() > 0; i++ {
            p := heap.Pop(&h).(Pair)
            res = append(res, p.c)
            p.cnt--
            if p.cnt > 0 {
                list[pos] = p
                pos++
            }
        }
        for j := 0; j < pos; j++ {
            heap.Push(&h, list[j])
        }
    }
    return string(res)
}

type Pair struct{
    c byte
    cnt uint16
}

type MaxHeap []Pair

func (this MaxHeap) Len() int{
    return len(this)
}

func (this MaxHeap) Less(i,j int) bool{
    return this[i].cnt>this[j].cnt || this[i].cnt==this[j].cnt && this[i].c<this[j].c
}

func (this MaxHeap) Swap(i,j int){
    this[i],this[j] = this[j],this[i]
}

func (this *MaxHeap) Push(i interface{}){
    *this = append(*this, i.(Pair))
}

func (this *MaxHeap) Pop() interface{}{
    last := this.Len()-1
    v := (*this)[last]
    *this = (*this)[:last]
    return v
}

func (this MaxHeap) Top() Pair {
    return this[0]
}


func main() {
    // Example 1:
    // Input: s = "aabbcc", k = 3
    // Output: "abcabc"
    // Explanation: The same letters are at least a distance of 3 from each other.
    fmt.Println(rearrangeString("aabbcc", 3)) // abcabc
    // Example 2:
    // Input: s = "aaabc", k = 3
    // Output: ""
    // Explanation: It is not possible to rearrange the string.
    fmt.Println(rearrangeString("aaabc", 3)) // ""
    // Example 3:
    // Input: s = "aaadbbcc", k = 2
    // Output: "abacabcd"
    // Explanation: The same letters are at least a distance of 2 from each other.
    fmt.Println(rearrangeString("aaadbbcc", 2)) // "abacabcd"
}