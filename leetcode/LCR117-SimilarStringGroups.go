package main

// LCR 117. 相似字符串组
// 如果交换字符串 X 中的两个不同位置的字母，使得它和字符串 Y 相等，那么称 X 和 Y 两个字符串相似。
// 如果这两个字符串本身是相等的，那它们也是相似的。

// 例如，"tars" 和 "rats" 是相似的 (交换 0 与 2 的位置)； "rats" 和 "arts" 也是相似的，但是 "star" 不与 "tars"，"rats"，或 "arts" 相似。

// 总之，它们通过相似性形成了两个关联组：{"tars", "rats", "arts"} 和 {"star"}。
// 注意，"tars" 和 "arts" 是在同一组中，即使它们并不相似。形式上，对每个组而言，要确定一个单词在组中，只需要这个词和该组中至少一个单词相似。

// 给定一个字符串列表 strs。列表中的每个字符串都是 strs 中其它所有字符串的一个 字母异位词 。
// 请问 strs 中有多少个相似字符串组？

// 字母异位词（anagram），一种把某个字符串的字母的位置（顺序）加以改换所形成的新词。

// 示例 1：
// 输入：strs = ["tars","rats","arts","star"]
// 输出：2

// 示例 2：
// 输入：strs = ["omv","ovm"]
// 输出：1

// 提示：
//     1 <= strs.length <= 300
//     1 <= strs[i].length <= 300
//     strs[i] 只包含小写字母。
//     strs 中的所有单词都具有相同的长度，且是彼此的字母异位词。


import "fmt"

// 并查集
func numSimilarGroups(strs []string) int {
    parent, rank := []int{}, []int{}
    var find func(x int) int
    find = func(x int) int {
        if parent[x] != x { parent[x] = find(parent[x]) }
        return parent[x]
    }
    union := func(x, y int) {
        xset, yset := find(x), find(y)
        if xset == yset {
            return
        } else if rank[xset] < rank[yset] {
            parent[xset] = yset
        } else if rank[xset] > rank[yset] {
            parent[yset] = xset
        } else {
            parent[yset] = xset
            rank[xset]++
        }
    }
    isSimilar := func(a, b string) bool {
        diff := 0
        for i := 0; i < len(a); i++ {
            if a[i] != b[i] {
                diff++
            }
        }
        return diff == 0 || diff == 2
    }
    n, count := len(strs), len(strs)
    parent = make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
    }
    rank = make([]int, n)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if isSimilar(strs[i], strs[j]) && find(i) != find(j) {
                count--
                union(i, j)
            }
        }
    }
    return count
}

func numSimilarGroups1(strs []string) int {
    n := len(strs)    
    set := NewUnionFindSet(n)
    isSimilar := func(s, t string) bool {
        diff := 0
        for i := range s {
            if s[i] != t[i] {
                diff++
                if diff > 2 {
                    return false
                }
            }
        }
        return true
    }
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if set.Find(j) != set.Find(i) {
                if isSimilar(strs[i], strs[j]) {
                    set.Union(j, i)
                }
            }
        }
    }
    m := make(map[int]bool)
    for i := 0; i < n; i++ {
        m[set.Find(i)] = true
    }
    return len(m)
}

type UnionFindSet struct {
    parents []int
}

func NewUnionFindSet(n int) *UnionFindSet {
    p := make([]int, n)
    for i:=0; i < n; i++ {
        p[i] = i
    }
    return &UnionFindSet{ parents: p}
}

func (this *UnionFindSet) Find(a int) int {
    if this.parents[a] != this.parents[this.parents[a]] {
        this.parents[a]  = this.Find(this.parents[a])
    }
    return this.parents[a]
}

func (this *UnionFindSet) Union(a, b int) bool {
    pa, pb := this.Find(a), this.Find(b)
    if pa == pb {
        return false
    }
    this.parents[pa] = pb
    return true
}

func main() {
    // Example 1:
    // Input: strs = ["tars","rats","arts","star"]
    // Output: 2
    fmt.Println(numSimilarGroups([]string{"tars","rats","arts","star"})) // 2
    // Example 2:
    // Input: strs = ["omv","ovm"]
    // Output: 1
    fmt.Println(numSimilarGroups([]string{"omv","ovm"})) // 1

    fmt.Println(numSimilarGroups1([]string{"tars","rats","arts","star"})) // 2
    fmt.Println(numSimilarGroups1([]string{"omv","ovm"})) // 1
}