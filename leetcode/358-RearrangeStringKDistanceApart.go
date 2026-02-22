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
import "cmp"
import "slices"

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

func rearrangeString1(s string, k int) string {
    // 1. 先对字符分组，判断是否存在合法解
    // 2. 如果存在合法解，则优先使用出现次数较多的字符，并让字符的间隔为 k
    // 3. 字符的数量超过 n/k+1，或者有超过 n%k 个字符的数量超过 n/k，则不存在合法解
    if k == 0 { return s }
    n := len(s)
    count := [26]int{}
    for _, v := range s { // 统计字符出现次数
        count[v - 'a']++
    }
    // 检查
    limit, countEqLimit := n/k+1, 0
    for _, v := range count {
        if v > limit {
            return ""
        }
        if v == limit { // 统计出现次数等于 limit 的字符数量
            countEqLimit++
        }
    }
    if countEqLimit > n % k {
        return ""
    }
    // 构造结果
    type Item struct {
        ch   byte
        freq int
    }
    items := make([]Item, 0, 26)
    for i, freq := range count {
        if freq > 0 {
            items = append(items, Item{
                ch:   byte(i) + 'a',
                freq: freq,
            })
        }
    }
    slices.SortFunc(items, func(item1, item2 Item) int {
        return -cmp.Compare(item1.freq, item2.freq)
    })
    res := make([]byte, n)
    // if n%items[0].freq == 0 -> len=n/items[0]
    // else (n-1)%(items[0].freq-1) == 0 -> len=(n-1)/(items[0].freq-1)
    // else -> len=n/items[0].freq+1
    var l int
    if n % items[0].freq == 0 {
        l = n / items[0].freq
    } else {
        l = (n - 1) / (items[0].freq - 1)
    }
    iter := make([]int, max(k, l)) // 记录每个位置的迭代器
    remain := 0
    for remain := range iter {
        iter[remain] = remain
    }
    // freq 为 n/len(iter)的必须要从头开始填充，剩下的必须轮转填充
    for i := range items {
        ch, freq := items[i].ch, items[i].freq
        if freq == n / len(iter) && iter[remain] != remain {
            remain = (remain + 1) % len(iter)
        }
        for j := 0; j < freq; j++ {
            for iter[remain] >= n {
                remain = (remain + 1) % len(iter)
            }
            res[iter[remain]] = ch
            iter[remain] += len(iter)
        }
    }
    return string(res)
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

    fmt.Println(rearrangeString("bluefrog", 2)) // "befgloru"
    fmt.Println(rearrangeString("leetcode", 2)) // "ecedelot"
    fmt.Println(rearrangeString("freewu", 2)) // "eferuw"

    fmt.Println(rearrangeString1("aabbcc", 3)) // abcabc
    fmt.Println(rearrangeString1("aaabc", 3)) // ""
    fmt.Println(rearrangeString1("aaadbbcc", 2)) // "abacabcd"
    fmt.Println(rearrangeString1("bluefrog", 2)) // "befgloru"
    fmt.Println(rearrangeString1("leetcode", 2)) // "ecedelot"
    fmt.Println(rearrangeString1("freewu", 2)) // "eferuw"
}