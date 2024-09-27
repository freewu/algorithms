package main

// 1724. Checking Existence of Edge Length Limited Paths II
// An undirected graph of n nodes is defined by edgeList, 
// where edgeList[i] = [ui, vi, disi] denotes an edge between nodes ui and vi with distance disi. 
// Note that there may be multiple edges between two nodes, and the graph may not be connected.

// Implement the DistanceLimitedPathsExist class:
//     DistanceLimitedPathsExist(int n, int[][] edgeList) 
//         Initializes the class with an undirected graph.
//     boolean query(int p, int q, int limit) 
//         Returns true if there exists a path from p to q such that each edge on the path has a distance strictly less than limit, and otherwise false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/05/messed.png" />
// Input
// ["DistanceLimitedPathsExist", "query", "query", "query", "query"]
// [[6, [[0, 2, 4], [0, 3, 2], [1, 2, 3], [2, 3, 1], [4, 5, 5]]], [2, 3, 2], [1, 3, 3], [2, 0, 3], [0, 5, 6]]
// Output
// [null, true, false, true, false]
// Explanation
// DistanceLimitedPathsExist distanceLimitedPathsExist = new DistanceLimitedPathsExist(6, [[0, 2, 4], [0, 3, 2], [1, 2, 3], [2, 3, 1], [4, 5, 5]]);
// distanceLimitedPathsExist.query(2, 3, 2); // return true. There is an edge from 2 to 3 of distance 1, which is less than 2.
// distanceLimitedPathsExist.query(1, 3, 3); // return false. There is no way to go from 1 to 3 with distances strictly less than 3.
// distanceLimitedPathsExist.query(2, 0, 3); // return true. There is a way to go from 2 to 0 with distance < 3: travel from 2 to 3 to 0.
// distanceLimitedPathsExist.query(0, 5, 6); // return false. There are no paths from 0 to 5.

// Constraints:
//     2 <= n <= 10^4
//     0 <= edgeList.length <= 10^4
//     edgeList[i].length == 3
//     0 <= ui, vi, p, q <= n-1
//     ui != vi
//     p != q
//     1 <= disi, limit <= 10^9
//     At most 10^4 calls will be made to query.

import "fmt"
import "slices"

type Edge struct { // 边类
    X int // from
    Y int // to
    L int // 边长
}

type UnionFind struct {
    parent []int
    len    []int
    count  int
}

func NewUnionFind(n int) UnionFind {
    parent, len := make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        parent[i], len[i] = i, -1
    }
    return UnionFind{ parent, len, 0 }
}

func (u *UnionFind) Union(x, y, l int) {
    xRoot, yRoot := u.Find(x), u.Find(y)
    if(xRoot != yRoot) {
        u.count++
        u.parent[yRoot] = xRoot
        u.len[yRoot] = l
    }
}

func (u *UnionFind) Find(x int) int {
    if(u.parent[x] == x) { return x }
    return u.Find(u.parent[x])
}

func (u *UnionFind) Query(x, limit int) int {
    if(u.len[x] >= limit || u.parent[x] == x) { return x }
    return u.Query(u.parent[x], limit)
}

// kruskal 并查集
type DistanceLimitedPathsExist struct {
    uf    UnionFind
}

func Constructor(n int, edgeList [][]int) DistanceLimitedPathsExist {
    edges := []Edge{}
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // 建立边集
    for i := 0; i < len(edgeList); i++ {
        from, to, l := min(edgeList[i][0], edgeList[i][1]), max(edgeList[i][0], edgeList[i][1]), edgeList[i][2]
        edges = append(edges, Edge{ from, to, l })
    }
    slices.SortFunc(edges, func(e1, e2 Edge) int {
        return e1.L - e2.L
    })
    // 建立并查集
    uf := NewUnionFind(n)
    for i := 0; i < len(edges); i++ {
        uf.Union(edges[i].X, edges[i].Y, edges[i].L)
    }
    return DistanceLimitedPathsExist{ uf }
}

func (this *DistanceLimitedPathsExist) Query(p int, q int, limit int) bool {
    return this.uf.Query(p, limit) == this.uf.Query(q, limit)
}

/**
 * Your DistanceLimitedPathsExist object will be instantiated and called as such:
 * obj := Constructor(n, edgeList);
 * param_1 := obj.Query(p,q,limit);
 */

func main() {
    // DistanceLimitedPathsExist distanceLimitedPathsExist = new DistanceLimitedPathsExist(6, [[0, 2, 4], [0, 3, 2], [1, 2, 3], [2, 3, 1], [4, 5, 5]]);
    obj := Constructor(6, [][]int{{0, 2, 4}, {0, 3, 2}, {1, 2, 3}, {2, 3, 1}, {4, 5, 5}})
    // distanceLimitedPathsExist.query(2, 3, 2); // return true. There is an edge from 2 to 3 of distance 1, which is less than 2.
    fmt.Println(obj.Query(2, 3, 2)) // true
    // distanceLimitedPathsExist.query(1, 3, 3); // return false. There is no way to go from 1 to 3 with distances strictly less than 3.
    fmt.Println(obj.Query(1, 3, 3)) // false
    // distanceLimitedPathsExist.query(2, 0, 3); // return true. There is a way to go from 2 to 0 with distance < 3: travel from 2 to 3 to 0.
    fmt.Println(obj.Query(2, 0, 3)) // true
    // distanceLimitedPathsExist.query(0, 5, 6); // return false. There are no paths from 0 to 5.
    fmt.Println(obj.Query(0, 5, 6)) // false
}