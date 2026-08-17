package main

// 4018. Total Sum of Interaction Cost in Tree Groups II
// You are given an integer n and an undirected tree rooted at node 0 with n nodes numbered from 0 to n - 1. 
// The tree is represented by a 2D integer array edges of length n - 1, where edges[i] = [ui, vi] indicates an undirected edge between nodes ui and vi.

// You are also given an integer array group of length n, where group[i] denotes the group label assigned to node i.
//     1. Two nodes u and v belong to the same group if and only if group[u] == group[v].
//     2. The interaction cost between two nodes is the shortest distance between them in the tree.

// Return the sum of interaction costs over all pairs of node indices (u, v) such that 0 <= u < v < n and group[u] == group[v].

// The shortest distance between two nodes is the number of edges on the unique path connecting them in the tree.

// Example 1:
// Input: n = 3, edges = [[0,1],[1,2]], group = [1,1,1]
// Output: 4
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2026/05/04/screenshot-2026-05-05-at-40329am.png" />
// All nodes belong to group 1. The interaction costs between the pairs of nodes are:
// Nodes [0, 1]: 1
// Nodes [1, 2]: 1
// Nodes [0, 2]: 2
// Thus, the total interaction cost is 1 + 1 + 2 = 4.

// Example 2:
// Input: n = 3, edges = [[0,1],[1,2]], group = [3,2,3]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2026/05/04/screenshot-2026-05-05-at-40416am.png" />
// Nodes 0 and 2 belong to group 3. The interaction cost between this pair is 2.
// Node 1 belongs to a different group and forms no valid pair. Therefore, the total interaction cost is 2.

// Example 3:
// Input: n = 4, edges = [[0,1],[0,2],[0,3]], group = [1,1,4,4]
// Output: 3
// Explanation:
// ​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/04/screenshot-2026-05-05-at-40819am.png" />
// Nodes belonging to the same groups and their interaction costs are:
// Group 1: Nodes [0, 1]: 1
// Group 4: Nodes [2, 3]: 2
// Thus, the total interaction cost is 1 + 2 = 3.

// Example 4:
// Input: n = 2, edges = [[0,1]], group = [1,2]
// Output: 0
// Explanation:
// All nodes belong to different groups and there are no valid pairs. Therefore, the total interaction cost is 0.

// Constraints:
//     1 <= n <= 10^5
//     edges.length == n - 1
//     edges[i] = [ui, vi]
//     0 <= ui, vi <= n - 1
//     group.length == n
//     1 <= group[i] <= n
//     The input is generated such that edges represents a valid tree.

import "fmt"
import "math/bits"
import "slices"

// 超出时间限制 992 / 1000
func interactionCosts(n int, edges [][]int, group []int) int64 {
    adj := make([][]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    total := make(map[int]int)
    for _, g := range group {
        total[g]++
    }
    type stackItem struct {
        node    int
        parent  int
        visited bool
    }
    stack := []stackItem{{node: 0, parent: -1, visited: false}}
    subMap := make([]map[int]int, n)
    var res int64 = 0
    for len(stack) > 0 {
        item := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        u, parent, visited := item.node, item.parent, item.visited
        if !visited {
            stack = append(stack, stackItem{u, parent, true})
            for _, v := range adj[u] {
                if v != parent {
                    stack = append(stack, stackItem{v, u, false})
                }
            }
        } else {
            cnt := make(map[int]int)
            cnt[group[u]] = 1
            for _, v := range adj[u] {
                if v == parent {
                    continue
                }
                child := subMap[v]
                // small‑to‑large: always merge smaller into larger
                if len(child) > len(cnt) {
                    cnt, child = child, cnt
                }
                for g, c := range child {
                    cnt[g] += c
                }
                subMap[v] = nil
            }
            subMap[u] = cnt
            if parent != -1 {
                var contrib int64
                for g, c := range cnt {
                    t := total[g]
                    contrib += int64(c) * int64(t-c)
                }
                res += contrib
            }
        }
    }
    return res
}

// 超出时间限制 998 / 1000
func interactionCosts1(n int, edges [][]int, group []int) int64 {
    type Pair struct {
        g int
        c int
    }
    adj := make([][]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    total := make([]int, n+1)
    for _, g := range group {
        total[g]++
    }
    res := int64(0)
    var dfs func(u, parent int) []Pair
    dfs = func(u, parent int) []Pair {
        // 当前节点
        list := []Pair{{g: group[u], c: 1}}
        for _, v := range adj[u] {
            if v == parent {
                continue
            }
            child := dfs(v, u)
            // small‑to‑large：短合并到长，减少拷贝
            if len(child) > len(list) {
                list, child = child, list
            }
            // 把child合并进list
            for _, cp := range child {
                found := false
                for i := range list {
                    if list[i].g == cp.g {
                        list[i].c += cp.c
                        found = true
                        break
                    }
                }
                if !found {
                    list = append(list, cp)
                }
            }
        }
        if parent != -1 {
            var contrib int64
            for _, p := range list {
                contrib += int64(p.c) * int64(total[p.g]-p.c)
            }
            res += contrib
        }
        return list
    }
    dfs(0, -1)
    return res
}

func interactionCosts2(n int, edges [][]int, group []int) int64 {
    g := make([][]int, n)
    for _, e := range edges {
        v, w := e[0], e[1]
        g[v] = append(g[v], w)
        g[w] = append(g[w], v)
    }
    dfn := make([]int, n)
    res, ts := int64(0), 0
    pa := make([][17]int, n)
    dep := make([]int, n)
    var build func(int, int)
    build = func(v, p int) {
        dfn[v] = ts
        ts++
        pa[v][0] = p
        for _, w := range g[v] {
            if w != p {
                dep[w] = dep[v] + 1
                build(w, v)
            }
        }
    }
    build(0, -1)
    mx := bits.Len(uint(n))
    for i := range mx - 1 {
        for v := range pa {
            p := pa[v][i]
            if p != -1 {
                pa[v][i+1] = pa[p][i]
            } else {
                pa[v][i+1] = -1
            }
        }
    }
    uptoDep := func(v, d int) int {
        for k := uint32(dep[v] - d); k > 0; k &= k - 1 {
            v = pa[v][bits.TrailingZeros32(k)]
        }
        return v
    }
    getLCA := func(v, w int) int {
        if dep[v] > dep[w] {
            v, w = w, v
        }
        w = uptoDep(w, dep[v])
        if w == v {
            return v
        }
        for i := mx - 1; i >= 0; i-- {
            pv, pw := pa[v][i], pa[w][i]
            if pv != pw {
                v, w = pv, pw
            }
        }
        return pa[v][0]
    }
    nodesMap := map[int][]int{}
    for i, x := range group {
        nodesMap[x] = append(nodesMap[x], i)
    }
    vt := make([][]int, n)   // 虚树
    isNode := make([]int, n) // 用来区分是关键节点还是 LCA
    for i := range isNode {
        isNode[i] = -1
    }
    addVtEdge := func(v, w int) {
        vt[v] = append(vt[v], w) // 往虚树上添加一条有向边
    }
    const root = 0
    st := []int{root} // 用根节点作为栈底哨兵
    for val, nodes := range nodesMap {
        // 对于相同点权的这一组关键节点 nodes，构建虚树
        slices.SortFunc(nodes, func(a, b int) int { return dfn[a] - dfn[b] })
        vt[root] = vt[root][:0] // 重置虚树
        st = st[:1]
        for _, v := range nodes {
            isNode[v] = val
            if v == root {
                continue
            }
            vt[v] = vt[v][:0]
            lca := getLCA(st[len(st)-1], v) // 路径的拐点（LCA）也加到虚树中
            // 回溯，加边
            for len(st) > 1 && dfn[lca] <= dfn[st[len(st)-2]] {
                addVtEdge(st[len(st)-2], st[len(st)-1])
                st = st[:len(st)-1]
            }
            if lca != st[len(st)-1] { // lca 不在栈中（首次遇到）
                vt[lca] = vt[lca][:0]
                addVtEdge(lca, st[len(st)-1])
                st[len(st)-1] = lca // 加到栈中
            }
            st = append(st, v)
        }
        // 最后的回溯，加边
        for i := 1; i < len(st); i++ {
            addVtEdge(st[i-1], st[i])
        }
        var dfs func(int) int
        dfs = func(v int) (size int) {
            // 如果 isNode[v] != t，那么 v 只是关键节点之间路径上的「拐点」
            if isNode[v] == val {
                size = 1
            }
            for _, w := range vt[v] {
                sz := dfs(w)
                wt := dep[w] - dep[v] // 虚树边权
                // 贡献法
                res += int64(wt) * int64(sz) * int64(len(nodes)-sz)
                size += sz
            }
            return
        }
        rt := root
        if isNode[rt] != val && len(vt[rt]) == 1 {
            // 注意 root 只是一个哨兵，不一定在虚树上，得从真正的根节点开始
            rt = vt[rt][0]
        }
        dfs(rt)
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, edges = [[0,1],[1,2]], group = [1,1,1]
    // Output: 4
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2026/05/04/screenshot-2026-05-05-at-40329am.png" />
    // All nodes belong to group 1. The interaction costs between the pairs of nodes are:
    // Nodes [0, 1]: 1
    // Nodes [1, 2]: 1
    // Nodes [0, 2]: 2
    // Thus, the total interaction cost is 1 + 1 + 2 = 4.
    fmt.Println(interactionCosts(3, [][]int{{0,1},{1,2}}, []int{1,1,1})) // 4
    // Example 2:
    // Input: n = 3, edges = [[0,1],[1,2]], group = [3,2,3]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2026/05/04/screenshot-2026-05-05-at-40416am.png" />
    // Nodes 0 and 2 belong to group 3. The interaction cost between this pair is 2.
    // Node 1 belongs to a different group and forms no valid pair. Therefore, the total interaction cost is 2.
    fmt.Println(interactionCosts(3, [][]int{{0,1},{1,2}}, []int{3,2,3})) // 2
    // Example 3:
    // Input: n = 4, edges = [[0,1],[0,2],[0,3]], group = [1,1,4,4]
    // Output: 3
    // Explanation:
    // ​​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/05/04/screenshot-2026-05-05-at-40819am.png" />
    // Nodes belonging to the same groups and their interaction costs are:
    // Group 1: Nodes [0, 1]: 1
    // Group 4: Nodes [2, 3]: 2
    // Thus, the total interaction cost is 1 + 2 = 3.
    fmt.Println(interactionCosts(4, [][]int{{0,1},{0,2},{0,3}}, []int{1,1,4,4})) // 3
    // Example 4:
    // Input: n = 2, edges = [[0,1]], group = [1,2]
    // Output: 0
    // Explanation:
    // All nodes belong to different groups and there are no valid pairs. Therefore, the total interaction cost is 0.
    fmt.Println(interactionCosts(2, [][]int{{0,1}}, []int{1,2})) // 0

    fmt.Println(interactionCosts(
        21, 
        [][]int{{3,9},{2,4},{3,7},{1,15},{10,13},{6,8},{1,11},{9,18},{8,20},{0,1},{0,2},{0,5},{2,3},{15,16},{7,14},{5,10},{3,6},{7,17},{4,12},{13,19}}, 
        []int{15,19,3,19,6,8,5,2,9,16,9,17,14,21,11,3,7,5,9,19,14}),
    ) // 42
    fmt.Println(interactionCosts(
        31, 
        [][]int{{13,22},{5,25},{1,5},{7,29},{2,4},{1,3},{20,28},{20,23},{21,30},{0,13},{0,1},{5,11},{19,26},{7,8},{9,19},{9,10},{3,24},{1,6},{0,2},{14,17},{18,20},{14,16},{5,7},{2,14},{2,9},{5,12},{10,15},{10,18},{25,27},{7,21}}, 
        []int{6,30,3,8,22,6,21,11,8,2,23,11,7,14,10,17,3,1,5,15,21,21,9,19,6,12,31,15,22,13,13}),
    ) // 52
    fmt.Println(interactionCosts(
        64, 
        [][]int{{50,54},{39,61},{4,34},{0,43},{37,44},{3,21},{8,17},{21,28},{38,47},{10,19},{0,2},{6,11},{15,18},{36,63},{27,30},{34,55},{22,27},{37,52},{4,8},{24,37},{7,14},{30,59},{28,49},{8,24},{42,57},{0,33},{13,16},{3,4},{27,36},{25,48},{20,35},{10,25},{1,39},{13,56},{45,50},{13,29},{0,6},{25,32},{14,22},{4,15},{0,7},{4,26},{1,10},{0,42},{12,13},{21,23},{11,20},{1,5},{36,40},{7,38},{2,3},{24,41},{38,53},{0,1},{2,51},{10,60},{0,9},{8,12},{14,46},{5,62},{35,45},{45,58},{4,31}}, 
        []int{55,42,56,60,26,64,53,17,49,3,25,64,1,38,12,14,14,62,28,2,27,15,24,45,48,25,9,57,26,34,30,64,52,59,37,28,59,49,5,31,38,2,7,34,39,60,51,38,12,40,20,49,29,51,42,23,60,59,37,37,62,59,36,45}),
    ) // 216

    fmt.Println(interactionCosts1(3, [][]int{{0,1},{1,2}}, []int{1,1,1})) // 4
    fmt.Println(interactionCosts1(3, [][]int{{0,1},{1,2}}, []int{3,2,3})) // 2
    fmt.Println(interactionCosts1(4, [][]int{{0,1},{0,2},{0,3}}, []int{1,1,4,4})) // 3
    fmt.Println(interactionCosts1(2, [][]int{{0,1}}, []int{1,2})) // 0
    fmt.Println(interactionCosts1(
        21, 
        [][]int{{3,9},{2,4},{3,7},{1,15},{10,13},{6,8},{1,11},{9,18},{8,20},{0,1},{0,2},{0,5},{2,3},{15,16},{7,14},{5,10},{3,6},{7,17},{4,12},{13,19}}, 
        []int{15,19,3,19,6,8,5,2,9,16,9,17,14,21,11,3,7,5,9,19,14}),
    ) // 42
    fmt.Println(interactionCosts1(
        31, 
        [][]int{{13,22},{5,25},{1,5},{7,29},{2,4},{1,3},{20,28},{20,23},{21,30},{0,13},{0,1},{5,11},{19,26},{7,8},{9,19},{9,10},{3,24},{1,6},{0,2},{14,17},{18,20},{14,16},{5,7},{2,14},{2,9},{5,12},{10,15},{10,18},{25,27},{7,21}}, 
        []int{6,30,3,8,22,6,21,11,8,2,23,11,7,14,10,17,3,1,5,15,21,21,9,19,6,12,31,15,22,13,13}),
    ) // 52
    fmt.Println(interactionCosts1(
        64, 
        [][]int{{50,54},{39,61},{4,34},{0,43},{37,44},{3,21},{8,17},{21,28},{38,47},{10,19},{0,2},{6,11},{15,18},{36,63},{27,30},{34,55},{22,27},{37,52},{4,8},{24,37},{7,14},{30,59},{28,49},{8,24},{42,57},{0,33},{13,16},{3,4},{27,36},{25,48},{20,35},{10,25},{1,39},{13,56},{45,50},{13,29},{0,6},{25,32},{14,22},{4,15},{0,7},{4,26},{1,10},{0,42},{12,13},{21,23},{11,20},{1,5},{36,40},{7,38},{2,3},{24,41},{38,53},{0,1},{2,51},{10,60},{0,9},{8,12},{14,46},{5,62},{35,45},{45,58},{4,31}}, 
        []int{55,42,56,60,26,64,53,17,49,3,25,64,1,38,12,14,14,62,28,2,27,15,24,45,48,25,9,57,26,34,30,64,52,59,37,28,59,49,5,31,38,2,7,34,39,60,51,38,12,40,20,49,29,51,42,23,60,59,37,37,62,59,36,45}),
    ) // 216

    
    fmt.Println(interactionCosts2(3, [][]int{{0,1},{1,2}}, []int{1,1,1})) // 4
    fmt.Println(interactionCosts2(3, [][]int{{0,1},{1,2}}, []int{3,2,3})) // 2
    fmt.Println(interactionCosts2(4, [][]int{{0,1},{0,2},{0,3}}, []int{1,1,4,4})) // 3
    fmt.Println(interactionCosts2(2, [][]int{{0,1}}, []int{1,2})) // 0
    fmt.Println(interactionCosts2(
        21, 
        [][]int{{3,9},{2,4},{3,7},{1,15},{10,13},{6,8},{1,11},{9,18},{8,20},{0,1},{0,2},{0,5},{2,3},{15,16},{7,14},{5,10},{3,6},{7,17},{4,12},{13,19}}, 
        []int{15,19,3,19,6,8,5,2,9,16,9,17,14,21,11,3,7,5,9,19,14}),
    ) // 42
    fmt.Println(interactionCosts2(
        31, 
        [][]int{{13,22},{5,25},{1,5},{7,29},{2,4},{1,3},{20,28},{20,23},{21,30},{0,13},{0,1},{5,11},{19,26},{7,8},{9,19},{9,10},{3,24},{1,6},{0,2},{14,17},{18,20},{14,16},{5,7},{2,14},{2,9},{5,12},{10,15},{10,18},{25,27},{7,21}}, 
        []int{6,30,3,8,22,6,21,11,8,2,23,11,7,14,10,17,3,1,5,15,21,21,9,19,6,12,31,15,22,13,13}),
    ) // 52
    fmt.Println(interactionCosts2(
        64, 
        [][]int{{50,54},{39,61},{4,34},{0,43},{37,44},{3,21},{8,17},{21,28},{38,47},{10,19},{0,2},{6,11},{15,18},{36,63},{27,30},{34,55},{22,27},{37,52},{4,8},{24,37},{7,14},{30,59},{28,49},{8,24},{42,57},{0,33},{13,16},{3,4},{27,36},{25,48},{20,35},{10,25},{1,39},{13,56},{45,50},{13,29},{0,6},{25,32},{14,22},{4,15},{0,7},{4,26},{1,10},{0,42},{12,13},{21,23},{11,20},{1,5},{36,40},{7,38},{2,3},{24,41},{38,53},{0,1},{2,51},{10,60},{0,9},{8,12},{14,46},{5,62},{35,45},{45,58},{4,31}}, 
        []int{55,42,56,60,26,64,53,17,49,3,25,64,1,38,12,14,14,62,28,2,27,15,24,45,48,25,9,57,26,34,30,64,52,59,37,28,59,49,5,31,38,2,7,34,39,60,51,38,12,40,20,49,29,51,42,23,60,59,37,37,62,59,36,45}),
    ) // 216
}