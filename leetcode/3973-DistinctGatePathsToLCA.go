package main

// 3973. Distinct Gate Paths to LCA
// You are given an undirected tree rooted at node 0 with n nodes numbered from 0 to n - 1, 
// represented by an array parent where parent[i] is the parent of node i.

// Each node i has three types of gates, given in a 2D array gates where gates[i] = [redi, bluei, whitei] which represents the number of red, blue, and white gates at node i.
//     1. Red gate: usable only with a red card.
//     2. Blue gate: usable only with a blue card.
//     3. White gate: usable with either card, but flips the card color when used.

// Alice and Bob start at given nodes with either a red or blue card (1 = red, 0 = blue). 
// They must independently move upward to their lowest common ancestor (LCA).

// At each node, a person may move to their parent only if they can use at least one gate at that node with their current card. 
// White gates may be used any number of times to flip the card color.

// Movement rules (one move = from u to parent[u]):
//     1. Movement is only upward toward the root.
//     2. At node u, pick exactly one specific gate instance. Identical gates are treated as separate and counted individually.
//     3. If holding a red card: use a red gate to remain red, or a white gate to change to blue.
//     4. If holding a blue card: use a blue gate to remain blue, or a white gate to change to red.
//     5. If no usable gate exists at u, the sequence ends.

// You are also given a 2D array queries where queries[i] = [aNodei, aCardi, bNodei, bCardi]:
//     1. aNodei, aCardi: Alice's starting node and card.
//     2. bNodei, bCardi: Bob's starting node and card.
    
// For each query, count the number of distinct valid ways modulo 10^9 + 7 for both to reach their LCA.

// After computing the result for all queries, return the bitwise XOR of those values.

// Note:
//     Two ways are distinct if the set of gates used differs for either Alice or Bob.
//     If any person is already at the LCA, then the number of ways for them is 1.

// The lowest common ancestor (LCA) is defined between two nodes a and b as the lowest node in a tree that has both a and b as descendants (where a node is allowed to be a descendant of itself).

// Example 1:
// Input: n = 3, parent = [-1,0,0], gates = [[1,0,1],[0,1,1],[1,1,0]], queries = [[1,0,2,0],[1,1,2,0],[1,0,2,1]]
// Output: 1
// Explanation:
// i | Alice[Node, Card] |	Bob[Node, Card]	| LCA | Alice Path | Bob Path  | Alice Ways	                    | Bob Ways             | Total Ways
// 0 | [1, 0]: Blue      | [2, 0]: Blue     | 0	  | 1 → 0	   | 2 → 0	   | 2 (1 Blue + 1 White at node 1) | 1 (1 Blue at node 2) | 2 × 1 = 2
// 1 | [1, 1]: Red       | [2, 0]: Blue     | 0	  | 1 → 0	   | 2 → 0	   | 1 (1 White at node 1)	        | 1 (1 Blue at node 2) | 1 × 1 = 1
// 2 | [1, 0]: Blue	     | [2, 1]: Red      | 0	  | 1 → 0	   | 2 → 0	   | 2 (1 Blue + 1 White at node 1)	| 1 (1 Red at node 2)  | 2 × 1 = 2
// Thus, the XOR of all values: 2 XOR 1 XOR 2 = 1.

// Example 2:
// Input: n = 3, parent = [-1,0,1], gates = [[0,1,2],[1,0,1],[0,0,3]], queries = [[2,0,1,0],[2,1,0,0],[1,1,2,1]]
// Output: 3
// Explanation:
// i	| Alice[Node, Card] | Bob[Node, Card]	| LCA | Alice Path | Bob Path    | Alice Ways	                    | Bob Ways              | Total Ways
// 0	| [2, 0]: Blue	    | [1, 0]: Blue	    | 1	  | 2 → 1	   | 1	         | 3 (3 White at node 2)	        | 1 (no move)	        | 3 × 1 = 3
// 1	| [2, 1]: Red       | [0, 0]: Blue	    | 0	  | 2 → 1 → 0  | 0	         | 3 (3 White at node 2) ×          | 1 (no move)	        | 3 × 1 = 3
//      |                   |                   |     |            |             | 1 (1 White at node 1) = 3	    |                       |
// 2	| [1, 1]: Red       | [2, 1]: Red	    | 1	  | 1	       | 2 → 1       | 1 (no move)                      | 3 (3 White at node 2)	| 1 × 3 = 3
// Thus, the XOR of all values: 3 XOR 3 XOR 3 = 3.

// Constraints:​​​​​​​
//     2 <= n <= 2 * 10^4
//     n == parent.length == gates.length
//     parent[0] == -1
//     0 <= parent[i] < n for i in [1, n - 1]
//     gates[i] == [redi, bluei, whitei]
//     0 <= redi, bluei, whitei <= 10
//     1 <= queries.length <= 2 * 10^4
//     queries[i] = [aNodei, aCardi, bNodei, bCardi]
//     0 <= aNodei, bNodei <= n - 1
//     0 <= aCardi, bCardi <= 1
//     The input is generated such that the array parent represents a valid tree.

import "fmt"
import "math/bits"

const MOD = 1_000_000_007

// Mat 存储四元组 (bb, br, rb, rr)
type Mat struct {
    bb, br, rb, rr int
}

type LcaBinaryLifting struct {
    depth []int
    pa    [][]int
    val   [][]Mat
    m     int
}

// mul 改造为独立普通函数（移除结构体接收者）
func mul(x, y Mat) Mat {
    xbb, xbr, xrb, xrr := x.bb, x.br, x.rb, x.rr
    ybb, ybr, yrb, yrr := y.bb, y.br, y.rb, y.rr
    return Mat{
        (xbb*ybb + xbr*yrb) % MOD,
        (xbb*ybr + xbr*yrr) % MOD,
        (xrb*ybb + xrr*yrb) % MOD,
        (xrb*ybr + xrr*yrr) % MOD,
    }
}

func NewLcaBinaryLifting(parent []int, gates [][]int) *LcaBinaryLifting {
    n := len(parent)
    m := bits.Len(uint(n))
    // 建图
    g := make([][]int, n)
    for i := 1; i < n; i++ {
        f := parent[i]
        g[f] = append(g[f], i)
    }
    depth := make([]int, n)
    // pa 初始化 -1
    pa := make([][]int, m)
    val := make([][]Mat, m)
    for i := 0; i < m; i++ {
        pa[i] = make([]int, n)
        val[i] = make([]Mat, n)
        for j := 0; j < n; j++ {
            pa[i][j] = -1
            // 单位矩阵 (1,0,0,1)
            val[i][j] = Mat{1, 0, 0, 1}
        }
    }
    // 填充0层val
    for i := 0; i < n; i++ {
        r, b, w := gates[i][0], gates[i][1], gates[i][2]
        val[0][i] = Mat{b, w, w, r}
    }
    // dfs
    var dfs func(x, fa int)
    dfs = func(x, fa int) {
        pa[0][x] = fa
        for _, y := range g[x] {
            depth[y] = depth[x] + 1
            dfs(y, x)
        }
    }
    dfs(0, -1)
    // 倍增预处理（调用独立函数mul，无l.前缀）
    for i := 0; i < m-1; i++ {
        for x := 0; x < n; x++ {
            p := pa[i][x]
            if p != -1 {
                pa[i+1][x] = pa[i][p]
                val[i+1][x] = mul(val[i][p], val[i][x])
            }
        }
    }
    return &LcaBinaryLifting{
        depth: depth,
        pa:    pa,
        val:   val,
        m:     m,
    }
}

// getKthAncestor 向上跳k步祖先
func (l *LcaBinaryLifting) getKthAncestor(node, k int) int {
    for i := 0; i < bits.Len(uint(k)); i++ {
        if (k >> i) & 1 == 1 {
            node = l.pa[i][node]
            if node < 0 {
                return -1
            }
        }
    }
    return node
}

// getLca 求x,y的lca
func (l *LcaBinaryLifting) getLca(x, y int) int {
    if l.depth[x] > l.depth[y] {
        x, y = y, x
    }
    // y上跳到同深度
    diff := l.depth[y] - l.depth[x]
    y = l.getKthAncestor(y, diff)
    if x == y {
        return x
    }
    // 从高层往下跳
    for i := l.m - 1; i >= 0; i-- {
        px := l.pa[i][x]
        py := l.pa[i][y]
        if px != py {
            x, y = px, py
        }
    }
    return l.pa[0][x]
}

// getWays 从x往上走到a，card=0返回bb+rb mod MOD，card=1返回br+rr mod MOD
func (l *LcaBinaryLifting) getWays(x, a, card int) int {
    // 初始单位矩阵
    cur := Mat{1, 0, 0, 1}
    d := l.depth[x] - l.depth[a]
    for i := 0; i < bits.Len(uint(d)); i++ {
        if (d >> i) & 1 == 1 {
            // 调用独立函数mul，无l.前缀
            cur = mul(l.val[i][x], cur)
            x = l.pa[i][x]
        }
    }
    if card == 0 {
        return (cur.bb + cur.rb) % MOD
    } else {
        return (cur.br + cur.rr) % MOD
    }
}

func distinctPaths(n int, parent []int, gates [][]int, queries [][]int) int {
    res, lca := 0, NewLcaBinaryLifting(parent, gates)
    for _, q := range queries {
        a, ac := q[0], q[1]
        b, bc := q[2], q[3]
        c := lca.getLca(a, b)
        wa := lca.getWays(a, c, ac)
        wb := lca.getWays(b, c, bc)
        mulVal := (wa * wb) % MOD
        res ^= mulVal
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, parent = [-1,0,0], gates = [[1,0,1],[0,1,1],[1,1,0]], queries = [[1,0,2,0],[1,1,2,0],[1,0,2,1]]
    // Output: 1
    // Explanation:
    // i | Alice[Node, Card] |	Bob[Node, Card]	| LCA | Alice Path | Bob Path  | Alice Ways	                    | Bob Ways             | Total Ways
    // 0 | [1, 0]: Blue      | [2, 0]: Blue     | 0	  | 1 → 0	   | 2 → 0	   | 2 (1 Blue + 1 White at node 1) | 1 (1 Blue at node 2) | 2 × 1 = 2
    // 1 | [1, 1]: Red       | [2, 0]: Blue     | 0	  | 1 → 0	   | 2 → 0	   | 1 (1 White at node 1)	        | 1 (1 Blue at node 2) | 1 × 1 = 1
    // 2 | [1, 0]: Blue	     | [2, 1]: Red      | 0	  | 1 → 0	   | 2 → 0	   | 2 (1 Blue + 1 White at node 1)	| 1 (1 Red at node 2)  | 2 × 1 = 2
    // Thus, the XOR of all values: 2 XOR 1 XOR 2 = 1.
    fmt.Println(distinctPaths(3, []int{-1,0,0}, [][]int{{1,0,1},{0,1,1},{1,1,0}}, [][]int{{1,0,2,0},{1,1,2,0},{1,0,2,1}})) // 1
    // Example 2:
    // Input: n = 3, parent = [-1,0,1], gates = [[0,1,2],[1,0,1],[0,0,3]], queries = [[2,0,1,0],[2,1,0,0],[1,1,2,1]]
    // Output: 3
    // Explanation:
    // i	| Alice[Node, Card] | Bob[Node, Card]	| LCA | Alice Path | Bob Path    | Alice Ways	                    | Bob Ways              | Total Ways
    // 0	| [2, 0]: Blue	    | [1, 0]: Blue	    | 1	  | 2 → 1	   | 1	         | 3 (3 White at node 2)	        | 1 (no move)	        | 3 × 1 = 3
    // 1	| [2, 1]: Red       | [0, 0]: Blue	    | 0	  | 2 → 1 → 0  | 0	         | 3 (3 White at node 2) ×          | 1 (no move)	        | 3 × 1 = 3
    //      |                   |                   |     |            |             | 1 (1 White at node 1) = 3	    |                       |
    // 2	| [1, 1]: Red       | [2, 1]: Red	    | 1	  | 1	       | 2 → 1       | 1 (no move)                      | 3 (3 White at node 2)	| 1 × 3 = 3
    // Thus, the XOR of all values: 3 XOR 3 XOR 3 = 3.
    fmt.Println(distinctPaths(3, []int{-1,0,1}, [][]int{{0,1,2},{1,0,1},{0,0,3}}, [][]int{{2,0,1,0},{2,1,0,0},{1,1,2,1}})) // 3

    fmt.Println(distinctPaths(2, []int{-1,0}, [][]int{{5,7,4},{5,3,10}}, [][]int{{0,1,0,0},{0,1,1,0}})) // 12 
}