package main

// 310. Minimum Height Trees
// A tree is an undirected graph in which any two vertices are connected by exactly one path. 
// In other words, any connected graph without simple cycles is a tree.
// Given a tree of n nodes labelled from 0 to n - 1, and an array of n - 1 edges where edges[i] = [ai, bi] indicates that there is an undirected edge between the two nodes ai and bi in the tree, you can choose any node of the tree as the root. When you select a node x as the root, the result tree has height h. Among all possible rooted trees, those with minimum height (i.e. min(h))  are called minimum height trees (MHTs).
// Return a list of all MHTs' root labels. You can return the answer in any order.
// The height of a rooted tree is the number of edges on the longest downward path between the root and a leaf.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/01/e1.jpg" />
// Input: n = 4, edges = [[1,0],[1,2],[1,3]]
// Output: [1]
// Explanation: As shown, the height of the tree is 1 when the root is the node with label 1 which is the only MHT.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/09/01/e2.jpg" />
// Input: n = 6, edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
// Output: [3,4]

// Constraints:
//     1 <= n <= 2 * 10^4
//     edges.length == n - 1
//     0 <= ai, bi < n
//     ai != bi
//     All the pairs (ai, bi) are distinct.
//     The given input is guaranteed to be a tree and there will be no repeated edges.

import "fmt"

// bfs
func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {
        return []int{ 0 }
    }
    queue := []int{}
    //graph := make(map[int][]int)
    graph := make([][]int, n)
    degree := make([]int, n)
    for _, e := range edges {
        na, nb := e[0], e[1] 
        graph[na] = append(graph[na], nb)
        graph[nb] = append(graph[nb], na)
        degree[na]++
        degree[nb]++ 
    }
    for i := range degree {
        if degree[i] == 1 {
            queue = append(queue, i)
        }
    }
    // topological sort
    for n > 2 {
        size := len(queue)
        n -= size
        for size > 0 {
            node := queue[0]
            queue = queue[1:]
            for _, v := range graph[node] {
                degree[v]--
                if degree[v] == 1 {
                    queue = append(queue, v)
                }
            }
            size--
        }
    }
    return queue
}

func findMinHeightTrees1(n int, edges [][]int) []int {
    if n == 1 {
        return []int{0}
    }
    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }
    parents := make([]int, n)
    bfs := func(start int) (x int) {
        vis := make([]bool, n)
        vis[start] = true
        q := []int{start}
        for len(q) > 0 {
            x, q = q[0], q[1:]
            for _, y := range g[x] {
                if !vis[y] {
                    vis[y] = true
                    parents[y] = x
                    q = append(q, y)
                }
            }
        }
        return
    }
    x := bfs(0) // 找到与节点 0 最远的节点 x
    y := bfs(x) // 找到与节点 x 最远的节点 y
    path := []int{}
    parents[x] = -1
    for y != -1 {
        path = append(path, y)
        y = parents[y]
    }
    m := len(path)
    if m%2 == 0 {
        return []int{path[m/2-1], path[m/2]}
    }
    return []int{path[m/2]}
}

func main() {
	fmt.Println(findMinHeightTrees(4,[][]int{ []int{1,0}, []int{1,2}, []int{1,3} })) // [1]
    fmt.Println(findMinHeightTrees(6,[][]int{ []int{3,0}, []int{3,1}, []int{3,2}, []int{3,4}, []int{5,4} })) // [3,4]

    fmt.Println(findMinHeightTrees1(4,[][]int{ []int{1,0}, []int{1,2}, []int{1,3} })) // [1]
    fmt.Println(findMinHeightTrees1(6,[][]int{ []int{3,0}, []int{3,1}, []int{3,2}, []int{3,4}, []int{5,4} })) // [3,4]

}