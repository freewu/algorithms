package main

// 1579. Remove Max Number of Edges to Keep Graph Fully Traversable
// Alice and Bob have an undirected graph of n nodes and three types of edges:
//     Type 1: Can be traversed by Alice only.
//     Type 2: Can be traversed by Bob only.
//     Type 3: Can be traversed by both Alice and Bob.

// Given an array edges where edges[i] = [typei, ui, vi] represents a bidirectional edge of type typei between nodes ui and vi, 
// find the maximum number of edges you can remove so that after removing the edges, 
// the graph can still be fully traversed by both Alice and Bob. 
// The graph is fully traversed by Alice and Bob if starting from any node, they can reach all other nodes.

// Return the maximum number of edges you can remove, or return -1 if Alice and Bob cannot fully traverse the graph.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/08/19/ex1.png" />
// Input: n = 4, edges = [[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]]
// Output: 2
// Explanation: If we remove the 2 edges [1,1,2] and [1,1,3]. The graph will still be fully traversable by Alice and Bob. Removing any additional edge will not make it so. So the maximum number of edges we can remove is 2.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/08/19/ex2.png" />
// Input: n = 4, edges = [[3,1,2],[3,2,3],[1,1,4],[2,1,4]]
// Output: 0
// Explanation: Notice that removing any edge will not make the graph fully traversable by Alice and Bob.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/08/19/ex3.png" />
// Input: n = 4, edges = [[3,2,3],[1,1,2],[2,3,4]]
// Output: -1
// Explanation: In the current graph, Alice cannot reach node 4 from the other nodes. Likewise, Bob cannot reach 1. Therefore it's impossible to make the graph fully traversable.

// Constraints:
//     1 <= n <= 10^5
//     1 <= edges.length <= min(10^5, 3 * n * (n - 1) / 2)
//     edges[i].length == 3
//     1 <= typei <= 3
//     1 <= ui < vi <= n
//     All tuples (typei, ui, vi) are distinct.

import "fmt"

func maxNumEdgesToRemove(n int, edges [][]int) int {
    father1, father2, father3 := make([]int,n+1), make([]int,n+1), make([]int,n+1)  
    for i, _ := range father3 {
        father3[i] = i
    }
    findgrand := func(father []int, u int)int{
        u2 := u
        for u != father[u]{
            u = father[u]
        }
        // 压缩路径  不压缩会超时
        fu := father[u2]
        for u2 != fu{
            father[u2] = u
            u2 = fu
            fu = father[u2]
        }
        return u
    }
    // 不同连通分量，连接；同连通分量，不连，返回false
    connect := func(father []int, u int, v int,count *int) bool{
        u,v = findgrand(father,u), findgrand(father,v)
        if u != v{
            father[v] = u
            *count++   // 添加的边数
            return true
        }
        return false
    }
    // 返回：处理一种类型的边后，图能否完全遍历
    deal := func(father []int,count *int, edgetype int) bool{
        for _, e := range edges{
            if e[0] == edgetype{
                connect(father,e[1],e[2],count)
            }
        }
        return *count >= n - 1
    }
    count3 := 0
    // 先处理type3的边
    if deal(father3, &count3, 3){
        return len(edges) - count3
    }
    copy(father1,father3)
    copy(father2,father3)
    count1, count2 := count3, count3
    if !deal(father1, &count1, 1) {
        return -1
    }
    if !deal(father2,&count2,2) {
        return -1
    }
    return len(edges) - count1 - count2 + count3
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/08/19/ex1.png" />
    // Input: n = 4, edges = [[3,1,2],[3,2,3],[1,1,3],[1,2,4],[1,1,2],[2,3,4]]
    // Output: 2
    // Explanation: If we remove the 2 edges [1,1,2] and [1,1,3]. The graph will still be fully traversable by Alice and Bob. Removing any additional edge will not make it so. So the maximum number of edges we can remove is 2.
    fmt.Println(maxNumEdgesToRemove(4,[][]int{{3,1,2},{3,2,3},{1,1,3},{1,2,4},{1,1,2},{2,3,4}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/08/19/ex2.png" />
    // Input: n = 4, edges = [[3,1,2],[3,2,3],[1,1,4],[2,1,4]]
    // Output: 0
    // Explanation: Notice that removing any edge will not make the graph fully traversable by Alice and Bob.
    fmt.Println(maxNumEdgesToRemove(4,[][]int{{3,1,2},{3,2,3},{1,1,4},{2,1,4}})) // 0
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/08/19/ex3.png" />
    // Input: n = 4, edges = [[3,2,3],[1,1,2],[2,3,4]]
    // Output: -1
    // Explanation: In the current graph, Alice cannot reach node 4 from the other nodes. Likewise, Bob cannot reach 1. Therefore it's impossible to make the graph fully traversable.
    fmt.Println(maxNumEdgesToRemove(4,[][]int{{3,2,3},{1,1,2},{2,3,4}})) // -1
}