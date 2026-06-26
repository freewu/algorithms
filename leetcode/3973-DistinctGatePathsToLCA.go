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

const MOD = 1_000_000_007
const LOG = 16 // 覆盖2^16=65536 > 2e4

// 2*2状态转移矩阵：mat[sFrom][sTo] = 从状态sFrom到sTo的方案数
type Matrix [2][2]int

// 单位矩阵（无移动，状态不变）
var I = Matrix{
    {1, 0},
    {0, 1},
}

// 矩阵乘法：A * B = 先走A路径，再走B路径
func mul(a, b Matrix) Matrix {
    res := Matrix{}
    res[0][0] = (a[0][0]*b[0][0] + a[0][1]*b[1][0]) % MOD
    res[0][1] = (a[0][0]*b[0][1] + a[0][1]*b[1][1]) % MOD
    res[1][0] = (a[1][0]*b[0][0] + a[1][1]*b[1][0]) % MOD
    res[1][1] = (a[1][0]*b[0][1] + a[1][1]*b[1][1]) % MOD
    return res
}

func distinctPaths(n int, parent []int, gates [][]int, queries [][]int) int {
    // 1. 构建每个节点的转移矩阵M[u]：经过u节点时的状态变化
    res, M := 0, make([]Matrix, n)
    for u := 0; u < n; u++ {
        r, b, w := gates[u][0], gates[u][1], gates[u][2]
        mat := Matrix{}
        // 初始状态=红色(1) → 经过u后转移到的状态
        mat[1][1] = r % MOD // 红门：保持红
        mat[1][0] = w % MOD // 白门：变蓝
        // 初始状态=蓝色(0) → 经过u后转移到的状态
        mat[0][0] = b % MOD // 蓝门：保持蓝
        mat[0][1] = w % MOD // 白门：变红
        M[u] = mat
    }
    // 2. 建树 + BFS计算节点深度
    adj := make([][]int, n)
    for i := 1; i < n; i++ {
        p := parent[i]
        adj[p] = append(adj[p], i)
    }
    depth := make([]int, n)
    queue := []int{0}
    for len(queue) > 0 {
        u := queue[0]
        queue = queue[1:]
        for _, v := range adj[u] {
            depth[v] = depth[u] + 1
            queue = append(queue, v)
        }
    }
    // 3. 预处理倍增表
    // up[k][u]：u向上跳2^k步到达的节点
    // matJump[k][u]：u向上跳2^k步的总转移矩阵（先走u的1步，再跳2^(k-1)步）
    up := make([][]int, LOG)
    matJump := make([][]Matrix, LOG)
    for k := range up {
        up[k] = make([]int, n)
        matJump[k] = make([]Matrix, n)
    }
    // 初始化k=0：跳1步（2^0=1）
    for u := 0; u < n; u++ {
        if parent[u] == -1 { // 根节点无父节点，跳1步仍为自身
            up[0][u] = u
            matJump[0][u] = I
        } else {
            up[0][u] = parent[u]
            matJump[0][u] = M[u] // 跳1步的矩阵=当前节点的转移矩阵
        }
    }
    // 填充倍增表（k≥1）
    for k := 1; k < LOG; k++ {
        for u := 0; u < n; u++ {
            mid := up[k-1][u]
            up[k][u] = up[k-1][mid]
            // 乘法顺序：先走2^(k-1)步，再走2^(k-1)步 → matJump[k-1][u] * matJump[k-1][mid]
            matJump[k][u] = mul(matJump[k-1][u], matJump[k-1][mid])
        }
    }
    // LCA：求x和y的最近公共祖先
    getLCA := func(x, y int) int {
        if depth[x] < depth[y] {
            x, y = y, x
        }
        // x上跳至与y同深度
        for k := LOG - 1; k >= 0; k-- {
            if depth[x]-(1<<k) >= depth[y] {
                x = up[k][x]
            }
        }
        if x == y {
            return x
        }
        // 同步上跳至LCA子节点
        for k := LOG - 1; k >= 0; k-- {
            if up[k][x] != up[k][y] {
                x = up[k][x]
                y = up[k][y]
            }
        }
        return up[0][x]
    }
    // getUpMatrix：计算x向上走到祖先anc的**总转移矩阵**（正向路径：x→parent(x)→…→anc）
    getUpMatrix := func(x, anc int) Matrix {
        res := I
        cur := x
        for k := LOG - 1; k >= 0; k-- {
            if depth[up[k][cur]] >= depth[anc] {
                res = mul(res, matJump[k][cur]) // 乘法顺序：先已有路径，再跳k步
                cur = up[k][cur]
            }
        }
        return res
    }
    // getDownMatrix：计算祖先anc向下走到x的**总转移矩阵**（正向路径：anc→…→x）
    getDownMatrix := func(anc, x int) Matrix {
        res := I
        cur := x
        // 反向收集跳步，再反向相乘（保证向下路径的矩阵顺序正确）
        var steps []Matrix
        for k := LOG - 1; k >= 0; k-- {
            if depth[up[k][cur]] >= depth[anc] {
                steps = append(steps, matJump[k][cur])
                cur = up[k][cur]
            }
        }
        // 反向遍历steps，相乘得到向下路径的矩阵
        for i := len(steps) - 1; i >= 0; i-- {
            res = mul(res, steps[i])
        }
        return res
    }
    // 处理所有查询，计算异或结果
    for _, q := range queries {
        a, sa, b := q[0], q[1],  q[2] // a：起点，sa：初始状态 b：终点
        lcaNode := getLCA(a, b)
        // 1. 计算a→LCA的向上转移矩阵
        upMat := getUpMatrix(a, lcaNode)
        // 2. 计算LCA→b的向下转移矩阵
        downMat := getDownMatrix(lcaNode, b)
        // 3. 总转移矩阵：先走upMat，再走downMat
        totalMat := mul(upMat, downMat)
        // 总方案数：初始状态sa经过总矩阵后的所有可能状态之和
        ways := (totalMat[sa][0] + totalMat[sa][1]) % MOD
        res ^= ways
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
}