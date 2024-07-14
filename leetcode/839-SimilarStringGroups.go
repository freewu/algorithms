package main

// 839. Similar String Groups
// Two strings, X and Y, are considered similar if either they are identical 
// or we can make them equivalent by swapping at most two letters (in distinct positions) within the string X.

// For example, "tars" and "rats" are similar (swapping at positions 0 and 2), and "rats" and "arts" are similar, 
// but "star" is not similar to "tars", "rats", or "arts".

// Together, these form two connected groups by similarity: {"tars", "rats", "arts"} and {"star"}.  
// Notice that "tars" and "arts" are in the same group even though they are not similar.  
// Formally, each group is such that a word is in the group if and only if it is similar to at least one other word in the group.

// We are given a list strs of strings where every string in strs is an anagram of every other string in strs.
// How many groups are there?

// Example 1:
// Input: strs = ["tars","rats","arts","star"]
// Output: 2

// Example 2:
// Input: strs = ["omv","ovm"]
// Output: 1

// Constraints:
//     1 <= strs.length <= 300
//     1 <= strs[i].length <= 300
//     strs[i] consists of lowercase letters only.
//     All words in strs have the same length and are anagrams of each other.

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