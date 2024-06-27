package main

// 685. Redundant Connection II
// In this problem, a rooted tree is a directed graph such that, 
// there is exactly one node (the root) for which all other nodes are descendants of this node, 
// plus every node has exactly one parent, except for the root node which has no parents.

// The given input is a directed graph that started as a rooted tree with n nodes (with distinct values from 1 to n), 
// with one additional directed edge added. The added edge has two different vertices chosen from 1 to n, 
// and was not an edge that already existed.

// The resulting graph is given as a 2D-array of edges. 
// Each element of edges is a pair [ui, vi] that represents a directed edge connecting nodes ui and vi, 
// where ui is a parent of child vi.

// Return an edge that can be removed so that the resulting graph is a rooted tree of n nodes. 
// If there are multiple answers, return the answer that occurs last in the given 2D-array.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/20/graph1.jpg" />
// Input: edges = [[1,2],[1,3],[2,3]]
// Output: [2,3]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/20/graph2.jpg" />
// Input: edges = [[1,2],[2,3],[3,4],[4,1],[1,5]]
// Output: [4,1]

// Constraints:
//     n == edges.length
//     3 <= n <= 1000
//     edges[i].length == 2
//     1 <= ui, vi <= n
//     ui != vi

import "fmt"

func findRedundantDirectedConnection(edges [][]int) []int {
    n := 1010
    father := make([]int,n)
    deg := make([]int,len(edges)+1)
    for _,v := range edges{
        deg[v[1]]++
    }
    twoDeg := make([]int,0)
    for i := len(edges)-1;i>=0;i--{
        if deg[edges[i][1]]==2{
            twoDeg = append(twoDeg,i)
        }
    }
    var find func (x int) int 
    find = func (x int) int {
        if father[x] == x{
            return x
        }
        father[x] = find(father[x])
        return father[x]
    }
    isSame := func (u,v int)bool{
        pu, pv := find(u), find(v)
        if pu == pv {
            return true
        }
        return false
    }
    join := func (u,v int) {
        u, v = find(u), find(v)
        if u == v {
            return
        }
        father[v] = u
    }
    initialize := func(){
        for i := range father {
            father[i] = i
        }
    }
    // 判断有向环图
    removeNode := func(edges [][]int) []int {
        initialize()
        res := []int{}
        for _,v := range edges {
            if isSame(v[0],v[1]) {
                res = v
            }
            join(v[0],v[1])
        }
        return res
    }
    // 判断具有两个入度的节点，在删除一个节点后，是否为树
    // 入度为2的节点有两种情况，删除一个节点后为环图，另一种情况为两个边随便删除一条都可以
    isTree := func (edges [][]int,towNode int)bool{
        initialize()
        for i,v := range edges {
            if i == towNode{
                continue
            }
            if isSame(v[0],v[1]) {
                return false
            }
            join(v[0],v[1])
        }
        return true
    }
    if len(twoDeg) > 0 {
        if isTree(edges,twoDeg[0]){
            return edges[twoDeg[0]]
        }
        return edges[twoDeg[1]]
    }
    return removeNode(edges)
}

func findRedundantDirectedConnection1(edges [][]int) []int {
    n := len(edges)
    father := make([]int, n+1) // 有向图
    init := func() {
        for i := 1; i <= n; i ++ {
            father[i] = i
        }
    }
    var find func(u int) int 
    find = func(u int) int {
        if father[u] != u {
            father[u] = find(father[u])
        }
        return father[u]
    }
    join := func(u, v int) {
        fu := find(u)
        fv := find(v)
        if fu == fv {
            return
        }
        father[fv] = fu
    }
    isSame := func(u, v int) bool {
        fu := find(u)
        fv := find(v)
        return fu == fv
    }
    // 如果存在入度为2的顶点，最后要构成树，那么删除的边一定为该顶点的某一条入边 
    // 统计入度，找到入度为2的点
    indegree := make([]int, n+1)
    targetNode := 0
    for i := 0; i < n; i ++ {
        indegree[edges[i][1]]++
        if indegree[edges[i][1]] == 2{
            targetNode = edges[i][1]
            break
        }
    }
    // 找入度为2的节点所对应的边(2条)，注意要倒序，因为优先返回最后出现在二维数组中的答案
    targetEdges := [][]int{}
    for i := n-1; i >= 0; i -- {
        if edges[i][1] == targetNode {
            targetEdges = append(targetEdges, edges[i])
        }
    }
    // 删一条边之后判断是不是树
    judge := func(u, v int) bool {
        init()
        for i := 0; i < n; i ++ {
            if edges[i][0] == u && edges[i][1] == v {
                continue // 删除
            }
            if isSame(edges[i][0], edges[i][1]) {
                return false // 如果删除之后，还存在一条边没加入之前已经联通，那么这条边就是多余
            }else{
                join(edges[i][0], edges[i][1])
            }
        }
        return true
    }
    // 两条边里删一个，看删哪个可以构成树
    for i := 0; i < len(targetEdges); i ++ {
        if judge(targetEdges[i][0], targetEdges[i][1]) {
            return targetEdges[i]
        }
    }
    // 如果不存在入度为2的顶点，那我们可以直接删除加入会构成环的边（同684. 冗余连接）
    res := []int{}
    init() // 注意重新初始化
    for i := 0; i < n; i ++ {
        if isSame(edges[i][0], edges[i][1]) {
            res = edges[i]
        } else {
            join(edges[i][0], edges[i][1])
        }

    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/12/20/graph1.jpg" />
    // Input: edges = [[1,2],[1,3],[2,3]]
    // Output: [2,3]
    fmt.Println(findRedundantDirectedConnection([][]int{{1,2},{1,3},{2,3}})) // [2,3]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/12/20/graph2.jpg" />
    // Input: edges = [[1,2],[2,3],[3,4],[4,1],[1,5]]
    // Output: [4,1]
    fmt.Println(findRedundantDirectedConnection([][]int{{1,2},{2,3},{3,4},{4,1},{1,5}})) // [4,1]

    fmt.Println(findRedundantDirectedConnection1([][]int{{1,2},{1,3},{2,3}})) // [2,3]
    fmt.Println(findRedundantDirectedConnection1([][]int{{1,2},{2,3},{3,4},{4,1},{1,5}})) // [4,1]
}