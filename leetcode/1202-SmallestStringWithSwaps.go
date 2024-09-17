package main

// 1202. Smallest String With Swaps
// You are given a string s, and an array of pairs of indices in the string pairs where pairs[i] = [a, b] indicates 2 indices(0-indexed) of the string.

// You can swap the characters at any pair of indices in the given pairs any number of times.

// Return the lexicographically smallest string that s can be changed to after using the swaps.

// Example 1:
// Input: s = "dcab", pairs = [[0,3],[1,2]]
// Output: "bacd"
// Explaination: 
// Swap s[0] and s[3], s = "bcad"
// Swap s[1] and s[2], s = "bacd"

// Example 2:
// Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
// Output: "abcd"
// Explaination: 
// Swap s[0] and s[3], s = "bcad"
// Swap s[0] and s[2], s = "acbd"
// Swap s[1] and s[2], s = "abcd"

// Example 3:
// Input: s = "cba", pairs = [[0,1],[1,2]]
// Output: "abc"
// Explaination: 
// Swap s[0] and s[1], s = "bca"
// Swap s[1] and s[2], s = "bac"
// Swap s[0] and s[1], s = "abc"

// Constraints:
//     1 <= s.length <= 10^5
//     0 <= pairs.length <= 10^5
//     0 <= pairs[i][0], pairs[i][1] < s.length
//     s only contains lower case English letters.

import "fmt"
import "sort"

// union find
func smallestStringWithSwaps(s string, pairs [][]int) string {
    uf := NewUnionFind(len(s))
    for _, pair := range pairs {
        uf.Union(pair[0], pair[1])
    }
    groups := make(map[int][]byte)
    for i := 0; i < len(s); i++ {
        groups[uf.Find(i)] = append(groups[uf.Find(i)], s[i])
    }
    for _, group := range groups {
        sort.Slice(group, func(i, j int) bool {
            return group[i] < group[j]
        })
    }
    res := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        res[i] = groups[uf.Find(i)][0]
        groups[uf.Find(i)] = groups[uf.Find(i)][1:]
    }
    return string(res)
}

type UnionFind struct {
    parent []int
    count  int
}

func NewUnionFind(n int) *UnionFind {
    parent := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
    }
    return &UnionFind{parent, n}
}

func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])
    }
    return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
    rootX, rootY := uf.Find(x), uf.Find(y)
    if rootX == rootY {
        return
    }
    uf.parent[rootX] = rootY
    uf.count--
}

func smallestStringWithSwaps1(s string, pairs [][]int) string {
    init := func(n int) []int {
        res := make([]int,n)
        for i ,_ := range res { res[i] = i }
        return res
    }
    parent := init(len(s))
    var find func(i int) int
    find = func(i int) int {
        if parent[i] != i {
            parent[i] = find(parent[i])
        }
        return parent[i]
    }
    union := func (i,j int) {
        pi, pj := find(i),find(j)
        if pi == pj { return }
        if pi > pj { pi, pj = pj, pi }
        parent[pj] = pi
    }
    for _, v := range pairs {
        union(v[0], v[1])
    }
    group := make(map[int][]byte)
    for i := 0; i < len(s); i++ {
        pi := find(i)
        group[pi] = append(group[pi],s[i])
    }
    for k, v := range group {
        sort.Slice(v, func(i,j int)bool{
            return v[i] < v[j]
        })
        group[k] = v
    }
    bytes := []byte(s)
    for i := 0; i < len(s); i++ {
        pi := find(i)
        bytes[i], group[pi] = group[pi][0], group[pi][1:]
    }
    return string(bytes)
}

// dfs
func smallestStringWithSwaps2(s string, pairs [][]int) string {
    adjList := map[int][]int{}
    for _, p := range pairs {
        adjList[p[0]] = append(adjList[p[0]], p[1])
        adjList[p[1]] = append(adjList[p[1]], p[0])
    }

    var dfs func(i int, adjList map[int][]int, adjIdx *[]int, seen []bool)
    dfs = func(i int, adjList map[int][]int, adjIdx *[]int, seen []bool) {
        *adjIdx = append(*adjIdx, i)
        seen[i] = true
        for _, n := range adjList[i] {
            if !seen[n] { dfs(n, adjList, adjIdx, seen) }
        }
    }
 
    smallStr, seen := make([]byte, len(s)), make([]bool, len(s))
    for i := range s {
        if seen[i] { continue }
        adjIdx := []int{}
        dfs(i, adjList, &adjIdx, seen)
       
        adjCh := []byte{}
        for _, chIdx := range adjIdx {
            adjCh = append(adjCh, s[chIdx])
        }
        sort.Slice(adjCh, func(i,j int) bool {return adjCh[i] < adjCh[j]})
        sort.Slice(adjIdx, func(i,j int) bool {return adjIdx[i] < adjIdx[j]})
        for j, orgIdx := range adjIdx {
            smallStr[orgIdx] = adjCh[j]
        }
    }
    return string(smallStr)
}

func main() {
    // Example 1:
    // Input: s = "dcab", pairs = [[0,3],[1,2]]
    // Output: "bacd"
    // Explaination: 
    // Swap s[0] and s[3], s = "bcad"
    // Swap s[1] and s[2], s = "bacd"
    fmt.Println(smallestStringWithSwaps("dcab",[][]int{{0,3},{1,2}})) // "bacd"
    // Example 2:
    // Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
    // Output: "abcd"
    // Explaination: 
    // Swap s[0] and s[3], s = "bcad"
    // Swap s[0] and s[2], s = "acbd"
    // Swap s[1] and s[2], s = "abcd"
    fmt.Println(smallestStringWithSwaps("dcab",[][]int{{0,3},{1,2},{0,2}})) // "abcd"
    // Example 3:
    // Input: s = "cba", pairs = [[0,1],[1,2]]
    // Output: "abc"
    // Explaination: 
    // Swap s[0] and s[1], s = "bca"
    // Swap s[1] and s[2], s = "bac"
    // Swap s[0] and s[1], s = "abc"
    fmt.Println(smallestStringWithSwaps("cba",[][]int{{0,1},{1,2}})) // "abc"

    fmt.Println(smallestStringWithSwaps1("dcab",[][]int{{0,3},{1,2}})) // "bacd"
    fmt.Println(smallestStringWithSwaps1("dcab",[][]int{{0,3},{1,2},{0,2}})) // "abcd"
    fmt.Println(smallestStringWithSwaps1("cba",[][]int{{0,1},{1,2}})) // "abc"

    fmt.Println(smallestStringWithSwaps2("dcab",[][]int{{0,3},{1,2}})) // "bacd"
    fmt.Println(smallestStringWithSwaps2("dcab",[][]int{{0,3},{1,2},{0,2}})) // "abcd"
    fmt.Println(smallestStringWithSwaps2("cba",[][]int{{0,1},{1,2}})) // "abc"
}