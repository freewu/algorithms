package main

// LCR 118. 冗余连接
// 树可以看成是一个连通且 无环 的 无向 图。

// 给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。
// 添加的边的两个顶点包含在 1 到 n 中间，且这条附加的边不属于树中已存在的边。
// 图的信息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。

// 请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的边。

// 示例 1：
// <img src="https://pic.leetcode-cn.com/1626676174-hOEVUL-image.png" />
// 输入: edges = [[1,2],[1,3],[2,3]]
// 输出: [2,3]

// 示例 2：
// <img src="https://pic.leetcode-cn.com/1626676179-kGxcmu-image.png" />
// 输入: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
// 输出: [1,4]

// 提示:
//     n == edges.length
//     3 <= n <= 1000
//     edges[i].length == 2
//     1 <= ai < bi <= edges.length
//     ai != bi
//     edges 中无重复元素
//     给定的图是连通的 

import "fmt"

// 并查集
func findRedundantConnection(edges [][]int) []int {
    // union find algorithm
    n := len(edges)
    parent, rank := make([]int, n+1), make([]int, n+1)
    find := func(x int) int {
        par := parent[x]
        for par != parent[par] {
            parent[par] = parent[parent[par]]
            par = parent[par]
        }
        return par
    }
    union := func(x int, y int) {
        parentX, parentY := find(x), find(y)
        if parentX == parentY {
            return
        }
        if rank[parentX] >= rank[parentY] {
            parent[parentY] = x
            rank[parentX]++
        } else {
            parent[parentX] = parentY
            rank[parentY]++
        }
    }
    for i := 1; i <= n; i++ {
        parent[i] = i
        rank[i] = 1
    }
    for _, edge := range edges {
        x, y := edge[0], edge[1]
        if find(x) == find(y) {
            return edge
        }
        union(x, y)
    }
    return []int{}
}

func findRedundantConnection1(edges [][]int) []int {
    nums := make([]int , len(edges)+1)
    for i := range nums {
        nums[i] = i
    }
    for _ , v := range edges {
        a, b := v[0], v[1]
        for nums[a] != a { a = nums[a] }
        for nums[b] != b { b = nums[b] }
        if a == b { return v }
        nums[a] = b
    }
    return []int{}
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/02/reduntant1-1-graph.jpg" />
    // Input: edges = [[1,2],[1,3],[2,3]]
    // Output: [2,3]
    fmt.Println(findRedundantConnection([][]int{{1,2},{1,3},{2,3}})) // [2,3]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/05/02/reduntant1-2-graph.jpg" />
    // Input: edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
    // Output: [1,4]
    fmt.Println(findRedundantConnection([][]int{{1,2},{2,3},{3,4},{1,4},{1,5}})) // [1,4]

    fmt.Println(findRedundantConnection1([][]int{{1,2},{1,3},{2,3}})) // [2,3]
    fmt.Println(findRedundantConnection1([][]int{{1,2},{2,3},{3,4},{1,4},{1,5}})) // [1,4]
}