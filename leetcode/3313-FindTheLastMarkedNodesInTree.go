package main

// 3313. Find the Last Marked Nodes in Tree
// There exists an undirected tree with n nodes numbered 0 to n - 1. 
// You are given a 2D integer array edges of length n - 1, 
// where edges[i] = [ui, vi] indicates that there is an edge between nodes ui and vi in the tree.

// Initially, all nodes are unmarked. 
// After every second, you mark all unmarked nodes which have at least one marked node adjacent to them.

// Return an array nodes where nodes[i] is the last node to get marked in the tree, 
// if you mark node i at time t = 0. If nodes[i] has multiple answers for any node i, 
// you can choose any one answer.

// Example 1:
// Input: edges = [[0,1],[0,2]]
// Output: [2,2,1]
// Explanation:
//         0
//      /     \
//     1       2
// <img src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122236.png" />
// For i = 0, the nodes are marked in the sequence: [0] -> [0,1,2]. Either 1 or 2 can be the answer.
// For i = 1, the nodes are marked in the sequence: [1] -> [0,1] -> [0,1,2]. Node 2 is marked last.
// For i = 2, the nodes are marked in the sequence: [2] -> [0,2] -> [0,1,2]. Node 1 is marked last.

// Example 2:
// Input: edges = [[0,1]]
// Output: [1,0]
// Explanation:
//         1
//       /
//      0
// <img src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122249.png">
// For i = 0, the nodes are marked in the sequence: [0] -> [0,1].
// For i = 1, the nodes are marked in the sequence: [1] -> [0,1].

// Example 3:
// Input: edges = [[0,1],[0,2],[2,3],[2,4]]
// Output: [3,3,1,1,1]
// Explanation:
//         0
//      /     \
//     1       2
//            /  \
//           3    4
// <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-2024-06-03-210550.png" />
// For i = 0, the nodes are marked in the sequence: [0] -> [0,1,2] -> [0,1,2,3,4].
// For i = 1, the nodes are marked in the sequence: [1] -> [0,1] -> [0,1,2] -> [0,1,2,3,4].
// For i = 2, the nodes are marked in the sequence: [2] -> [0,2,3,4] -> [0,1,2,3,4].
// For i = 3, the nodes are marked in the sequence: [3] -> [2,3] -> [0,2,3,4] -> [0,1,2,3,4].
// For i = 4, the nodes are marked in the sequence: [4] -> [2,4] -> [0,2,3,4] -> [0,1,2,3,4].

// Constraints:
//     2 <= n <= 10^5
//     edges.length == n - 1
//     edges[i].length == 2
//     0 <= edges[i][0], edges[i][1] <= n - 1
//     The input is generated such that edges represents a valid tree.

import "fmt"

func lastMarkedNodes(edges [][]int) []int {
    n := len(edges) + 1
    dc, mc, res := make([][]int, n), make([][4]int, n), make([]int, n)
    for _, v := range edges {
        dc[v[0]] = append(dc[v[0]], v[1])
        dc[v[1]] = append(dc[v[1]], v[0])
    }
    for i := 0; i < n; i++ {
        res[i] = -1
        mc[i] = [4]int{ i, 0, -1, -1}
    }
    var dfs1 func(x, far int) (int,int,int,int) 
    dfs1 = func(x, far int) (int,int,int,int) {
        a, b, c, d := mc[x][0], mc[x][1], mc[x][2], mc[x][3]
        for i := range dc[x] {
            if i != far {
                p, q, _, _ := dfs1(i,x)
                q++
                if q > b {
                    a,b,c,d = p,q,a,b
                } else if q > d {
                    c,d = p,q
                }
            }
        }
        mc[x] = [4]int{a,b,c,d}
        return a, b, c, d
    }
    dfs1(0,-1)
    var dfs2 func(x,far int)
    dfs2 = func(x,far int) {
        a, b, c, d := mc[x][0], mc[x][1], mc[x][2], mc[x][3]
        p, q := 0, 0
        if far >= 0 {
            p, q = mc[far][0], mc[far][1]
            if p == a {
                p, q = mc[far][2], mc[far][3]
            }
            q++
        }
        if q > b {
            a,b,c,d = p,q,a,b
        } else if q > d {
            c, d = p, q
        }
        res[x], mc[x] = a, [4]int{a,b,c,d}
        for i := range dc[x] {
            if i != far {
                dfs2(i, x)
            }
        }
    }
    dfs2(0,-1)
    return res
}

// class Solution:
//     def lastMarkedNodes(self, edges: List[List[int]]) -> List[int]:
//         n = len(edges) + 1
//         dc, mc, res = [[] for _ in range(n)], [[i,0,-1,-1] for i in range(n)], [-1] * n
//         for i,j in edges:
//             dc[i].append(j)
//             dc[j].append(i)
//         def dfs(x,far):
//             a,b,c,d = mc[x]
//             for i in dc[x]:
//                 if i != far:
//                     p,q,r,t = dfs(i,x)
//                     q += 1
//                     if q > b: a,b,c,d = p,q,a,b
//                     elif q > d: c,d = p,q
//             mc[x] = [a,b,c,d]
//             return mc[x]
//         dfs(0,-1)
//         def dfs2(x,far):
//             a,b,c,d = mc[x]
//             if far >= 0:
//                 p,q = mc[far][0], mc[far][1]
//                 if p == a: p,q = mc[far][2], mc[far][3]
//                 q += 1
//             else: p,q = 0,0
//             if q > b: a,b,c,d = p,q,a,b
//             elif q > d: c,d = p,q
//             res[x], mc[x] = a, [a,b,c,d]
//             for i in dc[x]:
//                 if i != far: dfs2(i,x)
//         dfs2(0,-1)
//         return res

// func lastMarkedNodes(edges [][]int) []int {
//     n, dia, pos := len(edges) + 1, 0, [2]int{-1, -1} // dia = 0 直径 pos = [-1, -1] # 直径端点
//     g := make([][]int, n)
//     for _, v := range edges {
//         g[v[0]] = append(g[v[0]], v[1])
//         g[v[1]] = append(g[v[1]], v[0])
//     }
//     var dfs1 func(x, far int) (int,int) // 求树的直径
//     dfs1 = func(x, far int) (int, int) {
//         m := [2]int{ 0, x }
//         for y in g[x] {
//             if y != far {
//                 t0, t1 := dfs1(y, x)
//                 if m[0] + t0 > dia {
//                     dia = m[0] + t0
//                     pos = [2]int{ m[1], t1 }
//                 }
//                 if t0 > m[0] {
//                     m = [2]int{ t0, t1 }
//                 }
//             }
//         }
//         return m[0] + 1, m[1]
//     }
//     dfs1(0, -1)
//     res := make([]int, n)
//     for i := 0; i < n; i++ {
//         res[i] = -1
//     }
//     var dijkstra func(x int) ([2]int)
//     dijkstra = func(x int) ([2]int) {
//         queue := [][2]int{}
//         queue = append(queue, [2]int{0, x})

//     }
    
//     //         def dijkstra(x):
// //             q = []
// //             q.append([0, x])
// //             dist = [inf] * n
// //             dist[x] = 0
// //             while q:
// //                 cost, x = heappop(q)
// //                 if dist[x] < cost:  # 多终点，需要把q里面的数据清空
// //                     continue
// //                 for y in g[x]:
// //                     if 1 + cost < dist[y]:
// //                         dist[y] = cost + 1
// //                         heappush(q, [1 + cost, y])
// //             return dist
// //         dist0 = dijkstra(res[0])#用dijkstra分别求出两个端点到其他点的距离
// //         dist1 = dijkstra(res[1])
// //         for i in range(n):#距离较大的端点就是答案
// //             if dist0[i] > dist1[i]:
// //                 ans[i] = res[0]
// //             else:
// //                 ans[i] = res[1]
// //         return ans
// }

// class Solution:
//     def lastMarkedNodes(self, edges: List[List[int]]) -> List[int]:
//         n = len(edges) + 1
//         g = [[] for _ in range(n)]
//         for x, y in edges:
//             g[x].append(y)
//             g[y].append(x)
//         self.dia = 0#直径
//         res = [-1, -1] # 直径端点
//         def dfs(x, fa):#求树的直径
//             m = [0, x]
//             for y in g[x]:
//                 if y != fa:
//                     tmp = dfs(y, x)
//                     if m[0] + tmp[0] > self.dia:
//                         self.dia = m[0] + tmp[0]
//                         res[0] = m[1]
//                         res[1] = tmp[1]
//                     if tmp[0] > m[0]:
//                         m = tmp
//             return [m[0] + 1, m[1]]
//         dfs(0, -1)
//         ans = [-1] * n
//         def dijkstra(x):
//             q = []
//             q.append([0, x])
//             dist = [inf] * n
//             dist[x] = 0
//             while q:
//                 cost, x = heappop(q)
//                 if dist[x] < cost:  # 多终点，需要把q里面的数据清空
//                     continue
//                 for y in g[x]:
//                     if 1 + cost < dist[y]:
//                         dist[y] = cost + 1
//                         heappush(q, [1 + cost, y])
//             return dist
//         dist0 = dijkstra(res[0])#用dijkstra分别求出两个端点到其他点的距离
//         dist1 = dijkstra(res[1])
//         for i in range(n):#距离较大的端点就是答案
//             if dist0[i] > dist1[i]:
//                 ans[i] = res[0]
//             else:
//                 ans[i] = res[1]
//         return ans



func main() {
    // Example 1:
    // Input: edges = [[0,1],[0,2]]
    // Output: [2,2,1]
    // Explanation:
    //         0
    //      /     \
    //     1       2
    // <img src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122236.png" />
    // For i = 0, the nodes are marked in the sequence: [0] -> [0,1,2]. Either 1 or 2 can be the answer.
    // For i = 1, the nodes are marked in the sequence: [1] -> [0,1] -> [0,1,2]. Node 2 is marked last.
    // For i = 2, the nodes are marked in the sequence: [2] -> [0,2] -> [0,1,2]. Node 1 is marked last.
    fmt.Println(lastMarkedNodes([][]int{{0,1},{0,2}})) // [2,2,1]
    // Example 2:
    // Input: edges = [[0,1]]
    // Output: [1,0]
    // Explanation:
    //         1
    //       /
    //      0
    // <img src="https://assets.leetcode.com/uploads/2024/06/01/screenshot-2024-06-02-122249.png">
    // For i = 0, the nodes are marked in the sequence: [0] -> [0,1].
    // For i = 1, the nodes are marked in the sequence: [1] -> [0,1].
    fmt.Println(lastMarkedNodes([][]int{{0,1}})) // [1,0]
    // Example 3:
    // Input: edges = [[0,1],[0,2],[2,3],[2,4]]
    // Output: [3,3,1,1,1]
    // Explanation:
    //         0
    //      /     \
    //     1       2
    //            /  \
    //           3    4
    // <img src="https://assets.leetcode.com/uploads/2024/06/03/screenshot-2024-06-03-210550.png" />
    // For i = 0, the nodes are marked in the sequence: [0] -> [0,1,2] -> [0,1,2,3,4].
    // For i = 1, the nodes are marked in the sequence: [1] -> [0,1] -> [0,1,2] -> [0,1,2,3,4].
    // For i = 2, the nodes are marked in the sequence: [2] -> [0,2,3,4] -> [0,1,2,3,4].
    // For i = 3, the nodes are marked in the sequence: [3] -> [2,3] -> [0,2,3,4] -> [0,1,2,3,4].
    // For i = 4, the nodes are marked in the sequence: [4] -> [2,4] -> [0,2,3,4] -> [0,1,2,3,4].
    fmt.Println(lastMarkedNodes([][]int{{0,1},{0,2},{2,3},{2,4}})) // [3,3,1,1,1]
}