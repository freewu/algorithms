package main

// 3820. Pythagorean Distance Nodes in a Tree
// You are given an integer n and an undirected tree with n nodes numbered from 0 to n - 1. 
// The tree is represented by a 2D array edges of length n - 1, where edges[i] = [ui, vi] indicates an undirected edge between ui and vi.

// You are also given three distinct target nodes x, y, and z.

// For any node u in the tree:
//     1. Let dx be the distance from u to node x
//     2. Let dy be the distance from u to node y
//     3. Let dz be the distance from u to node z

// The node u is called special if the three distances form a Pythagorean Triplet.

// Return an integer denoting the number of special nodes in the tree.

// A Pythagorean triplet consists of three integers a, b, and c which, when sorted in ascending order, satisfy a2 + b2 = c2.

// The distance between two nodes in a tree is the number of edges on the unique path between them.

// Example 1:
// Input: n = 4, edges = [[0,1],[0,2],[0,3]], x = 1, y = 2, z = 3
// Output: 3
// Explanation:
// For each node, we compute its distances to nodes x = 1, y = 2, and z = 3.
// Node 0 has distances 1, 1, and 1. After sorting, the distances are 1, 1, and 1, which do not satisfy the Pythagorean condition.
// Node 1 has distances 0, 2, and 2. After sorting, the distances are 0, 2, and 2. Since 02 + 22 = 22, node 1 is special.
// Node 2 has distances 2, 0, and 2. After sorting, the distances are 0, 2, and 2. Since 02 + 22 = 22, node 2 is special.
// Node 3 has distances 2, 2, and 0. After sorting, the distances are 0, 2, and 2. This also satisfies the Pythagorean condition.
// Therefore, nodes 1, 2, and 3 are special, and the answer is 3.

// Example 2:
// Input: n = 4, edges = [[0,1],[1,2],[2,3]], x = 0, y = 3, z = 2
// Output: 0
// Explanation:
// For each node, we compute its distances to nodes x = 0, y = 3, and z = 2.
// Node 0 has distances 0, 3, and 2. After sorting, the distances are 0, 2, and 3, which do not satisfy the Pythagorean condition.
// Node 1 has distances 1, 2, and 1. After sorting, the distances are 1, 1, and 2, which do not satisfy the Pythagorean condition.
// Node 2 has distances 2, 1, and 0. After sorting, the distances are 0, 1, and 2, which do not satisfy the Pythagorean condition.
// Node 3 has distances 3, 0, and 1. After sorting, the distances are 0, 1, and 3, which do not satisfy the Pythagorean condition.
// No node satisfies the Pythagorean condition. Therefore, the answer is 0.

// Example 3:
// Input: n = 4, edges = [[0,1],[1,2],[1,3]], x = 1, y = 3, z = 0
// Output: 1
// Explanation:
// For each node, we compute its distances to nodes x = 1, y = 3, and z = 0.
// Node 0 has distances 1, 2, and 0. After sorting, the distances are 0, 1, and 2, which do not satisfy the Pythagorean condition.
// Node 1 has distances 0, 1, and 1. After sorting, the distances are 0, 1, and 1. Since 02 + 12 = 12, node 1 is special.
// Node 2 has distances 1, 2, and 2. After sorting, the distances are 1, 2, and 2, which do not satisfy the Pythagorean condition.
// Node 3 has distances 1, 0, and 2. After sorting, the distances are 0, 1, and 2, which do not satisfy the Pythagorean condition.
// Therefore, the answer is 1.

// Constraints:
//     4 <= n <= 10^5
//     edges.length == n - 1
//     edges[i] = [ui, vi]
//     0 <= ui, vi, x, y, z <= n - 1
//     x, y, and z are pairwise distinct.
//     The input is generated such that edges represent a valid tree.

import "fmt"
import "sort"

func specialNodes(n int, edges [][]int, x int, y int, z int) int {
    res, graph := 0, make([][]int, n)
    for _, e := range edges {
        v, w := e[0], e[1]
        graph[v] = append(graph[v], w)
        graph[w] = append(graph[w], v)
    }
    calc := func(start int) []int {
        dis := make([]int, n)
        var dfs func(int, int)
        dfs = func(v, fa int) {
            for _, w := range graph[v] {
                if w != fa {
                    dis[w] = dis[v] + 1
                    dfs(w, v)
                }
            }
        }
        dfs(start, -1)
        return dis
    }
    dx, dy, dz := calc(x), calc(y), calc(z)
    for i := range n {
        arr := []int{dx[i], dy[i], dz[i]}
        sort.Ints(arr)
        if arr[0] * arr[0] + arr[1] * arr[1] == arr[2] * arr[2] {
            res++
        }
    }
    return res 
}

func specialNodes1(n int, edges [][]int, x int, y int, z int) int {
    res, m := 0, len(edges)
    graph := make([][]int, n)
    for i := 0; i < m; i++ {
        u := edges[i][0]
        v := edges[i][1]
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
    }
    //从x，y，z出发？
    var bfs func(int) []int
    bfs = func(i int) []int {
        queue := []int{}
        queue = append(queue, i)
        count := make([]int, n)
        for j := 0; j < len(queue); j++ {
            a := queue[j]
            for _, value := range graph[a] {
                if count[value] == 0&&value!=i{
                    count[value] = count[a] + 1
                    queue = append(queue, value)
                }
            }
        }
        return count
    }
    cnt_x, cnt_y, cnt_z := bfs(x), bfs(y), bfs(z)
    for i := 0; i < n; i++ {
        a := cnt_x[i]
        b := cnt_y[i]
        c := cnt_z[i]
        if a * a + b * b == c * c || c * c + b * b == a * a || a * a + c * c == b * b {
            res++
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, edges = [[0,1],[0,2],[0,3]], x = 1, y = 2, z = 3
    // Output: 3
    // Explanation:
    // For each node, we compute its distances to nodes x = 1, y = 2, and z = 3.
    // Node 0 has distances 1, 1, and 1. After sorting, the distances are 1, 1, and 1, which do not satisfy the Pythagorean condition.
    // Node 1 has distances 0, 2, and 2. After sorting, the distances are 0, 2, and 2. Since 02 + 22 = 22, node 1 is special.
    // Node 2 has distances 2, 0, and 2. After sorting, the distances are 0, 2, and 2. Since 02 + 22 = 22, node 2 is special.
    // Node 3 has distances 2, 2, and 0. After sorting, the distances are 0, 2, and 2. This also satisfies the Pythagorean condition.
    // Therefore, nodes 1, 2, and 3 are special, and the answer is 3.
    fmt.Println(specialNodes(4, [][]int{{0,1},{0,2},{0,3}}, 1, 2, 3)) // 3
    // Example 2:
    // Input: n = 4, edges = [[0,1],[1,2],[2,3]], x = 0, y = 3, z = 2
    // Output: 0
    // Explanation:
    // For each node, we compute its distances to nodes x = 0, y = 3, and z = 2.
    // Node 0 has distances 0, 3, and 2. After sorting, the distances are 0, 2, and 3, which do not satisfy the Pythagorean condition.
    // Node 1 has distances 1, 2, and 1. After sorting, the distances are 1, 1, and 2, which do not satisfy the Pythagorean condition.
    // Node 2 has distances 2, 1, and 0. After sorting, the distances are 0, 1, and 2, which do not satisfy the Pythagorean condition.
    // Node 3 has distances 3, 0, and 1. After sorting, the distances are 0, 1, and 3, which do not satisfy the Pythagorean condition.
    // No node satisfies the Pythagorean condition. Therefore, the answer is 0.
    fmt.Println(specialNodes(4, [][]int{{0,1},{1,2},{2,3}}, 0, 3, 2)) // 0
    // Example 3:
    // Input: n = 4, edges = [[0,1],[1,2],[1,3]], x = 1, y = 3, z = 0
    // Output: 1
    // Explanation:
    // For each node, we compute its distances to nodes x = 1, y = 3, and z = 0.
    // Node 0 has distances 1, 2, and 0. After sorting, the distances are 0, 1, and 2, which do not satisfy the Pythagorean condition.
    // Node 1 has distances 0, 1, and 1. After sorting, the distances are 0, 1, and 1. Since 02 + 12 = 12, node 1 is special.
    // Node 2 has distances 1, 2, and 2. After sorting, the distances are 1, 2, and 2, which do not satisfy the Pythagorean condition.
    // Node 3 has distances 1, 0, and 2. After sorting, the distances are 0, 1, and 2, which do not satisfy the Pythagorean condition.
    // Therefore, the answer is 1.
    fmt.Println(specialNodes(4, [][]int{{0,1},{0,2},{0,3}}, 1, 3, 0)) // 1

    fmt.Println(specialNodes1(4, [][]int{{0,1},{0,2},{0,3}}, 1, 2, 3)) // 3
    fmt.Println(specialNodes1(4, [][]int{{0,1},{1,2},{2,3}}, 0, 3, 2)) // 0
    fmt.Println(specialNodes1(4, [][]int{{0,1},{0,2},{0,3}}, 1, 3, 0)) // 1
}