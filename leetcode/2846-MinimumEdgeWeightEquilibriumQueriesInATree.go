package main

// 2846. Minimum Edge Weight Equilibrium Queries in a Tree
// There is an undirected tree with n nodes labeled from 0 to n - 1. 
// You are given the integer n and a 2D integer array edges of length n - 1, where edges[i] = [ui, vi, wi] indicates that there is an edge between nodes ui and vi with weight wi in the tree.
// You are also given a 2D integer array queries of length m, where queries[i] = [ai, bi]. 
// For each query, find the minimum number of operations required to make the weight of every edge on the path from ai to bi equal. 
// In one operation, you can choose any edge of the tree and change its weight to any value.

// Note that:
//     Queries are independent of each other, meaning that the tree returns to its initial state on each new query.
//     The path from ai to bi is a sequence of distinct nodes starting with node ai and ending with node bi such that every two adjacent nodes in the sequence share an edge in the tree.

// Return an array answer of length m where answer[i] is the answer to the ith query.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/08/11/graph-6-1.png" />
// Input: n = 7, edges = [[0,1,1],[1,2,1],[2,3,1],[3,4,2],[4,5,2],[5,6,2]], queries = [[0,3],[3,6],[2,6],[0,6]]
// Output: [0,0,1,3]
// Explanation: In the first query, all the edges in the path from 0 to 3 have a weight of 1. Hence, the answer is 0.
// In the second query, all the edges in the path from 3 to 6 have a weight of 2. Hence, the answer is 0.
// In the third query, we change the weight of edge [2,3] to 2. After this operation, all the edges in the path from 2 to 6 have a weight of 2. Hence, the answer is 1.
// In the fourth query, we change the weights of edges [0,1], [1,2] and [2,3] to 2. After these operations, all the edges in the path from 0 to 6 have a weight of 2. Hence, the answer is 3.
// For each queries[i], it can be shown that answer[i] is the minimum number of operations needed to equalize all the edge weights in the path from ai to bi.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2023/08/11/graph-9-1.png" />
// Input: n = 8, edges = [[1,2,6],[1,3,4],[2,4,6],[2,5,3],[3,6,6],[3,0,8],[7,0,2]], queries = [[4,6],[0,4],[6,5],[7,4]]
// Output: [1,2,2,3]
// Explanation: In the first query, we change the weight of edge [1,3] to 6. After this operation, all the edges in the path from 4 to 6 have a weight of 6. Hence, the answer is 1.
// In the second query, we change the weight of edges [0,3] and [3,1] to 6. After these operations, all the edges in the path from 0 to 4 have a weight of 6. Hence, the answer is 2.
// In the third query, we change the weight of edges [1,3] and [5,2] to 6. After these operations, all the edges in the path from 6 to 5 have a weight of 6. Hence, the answer is 2.
// In the fourth query, we change the weights of edges [0,7], [0,3] and [1,3] to 6. After these operations, all the edges in the path from 7 to 4 have a weight of 6. Hence, the answer is 3.
// For each queries[i], it can be shown that answer[i] is the minimum number of operations needed to equalize all the edge weights in the path from ai to bi.
 
// Constraints:
//         1 <= n <= 10^4
//         edges.length == n - 1
//         edges[i].length == 3
//         0 <= ui, vi < n
//         1 <= wi <= 26
//         The input is generated such that edges represents a valid tree.
//         1 <= queries.length == m <= 2 * 10^4
//         queries[i].length == 2
//         0 <= ai, bi < n

import "fmt"
import "slices"
//import "bits"

// func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
//     // 权重的最大值
//     const W = 26
//     var find func (uf []int, i int) int 
//     find = func (uf []int, i int) int {
//         if uf[i] == i {
//             return i
//         }
//         uf[i] = find(uf, uf[i])
//         return uf[i]
//     }

//     m := len(queries)
//     neighbors := make([]map[int]int, n)
//     for i := 0; i < n; i++ {
//         neighbors[i] = map[int]int{}
//     }
//     for _, edge := range edges {
//         neighbors[edge[0]][edge[1]] = edge[2]
//         neighbors[edge[1]][edge[0]] = edge[2]
//     }
//     queryArr := make([][][2]int, n)
//     for i := 0; i < m; i++ {
//         queryArr[queries[i][0]] = append(queryArr[queries[i][0]], [2]int{queries[i][1], i})
//         queryArr[queries[i][1]] = append(queryArr[queries[i][1]], [2]int{queries[i][0], i})
//     }

//     count := make([][]int, n)
//     for i := 0; i < n; i++ {
//         count[i] = make([]int, W + 1)
//     }
//     visited, uf, lca := make([]int, n), make([]int, n), make([]int, n)
//     var tarjan func(int, int)
//     // 最近公共祖先节点的求解可以采用 Tarjan 算法
//     // https://oi.wiki/graph/lca/
//     tarjan = func(node, parent int) {
//         if parent != -1 {
//             copy(count[node], count[parent])
//             count[node][neighbors[node][parent]]++
//         }
//         uf[node] = node
//         for child, _ := range neighbors[node] {
//             if child == parent {
//                 continue
//             }
//             tarjan(child, node)
//             uf[child] = node
//         }
//         for _, query := range queryArr[node] {
//             node1, index := query[0], query[1]
//             if node != node1 && visited[node1] == 0 {
//                 continue
//             }
//             lca[index] = find(uf, node1)
//         }
//         visited[node] = 1
//     }
//     tarjan(0, -1)
//     res := make([]int, m)
//     for i := 0; i < m; i++ {
//         totalCount, maxCount := 0, 0
//         for j := 1; j <= W; j++ {
//             t := count[queries[i][0]][j] + count[queries[i][1]][j] - 2 * count[lca[i]][j]
//             maxCount, totalCount = max(maxCount, t), totalCount + t
//         }
//         res[i] = totalCount - maxCount
//     }
//     return res
// }

func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]-1
		g[v] = append(g[v], edge{w, wt})
		g[w] = append(g[w], edge{v, wt})
	}
	
	const mx = 14 // 2^14 > 10^4
	type pair struct {
		p   int
		cnt [26]int
	}
	pa := make([][mx]pair, n)
	depth := make([]int, n)
	var build func(int, int, int)
	build = func(v, p, d int) {
		pa[v][0].p = p
		depth[v] = d
		for _, e := range g[v] {
			if w := e.to; w != p {
				pa[w][0].cnt[e.wt] = 1
				build(w, v, d+1)
			}
		}
	}
	build(0, -1, 0)

	// 倍增模板
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p.p != -1 {
				pp := pa[p.p][i]
				pa[v][i+1].p = pp.p
				for j := 0; j < 26; j++ {
					pa[v][i+1].cnt[j] = p.cnt[j] + pp.cnt[j]
				}
			} else {
				pa[v][i+1].p = -1
			}
		}
	}

	// 计算 LCA 模板（这里返回最小操作次数）
	// https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solution/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw/
	f := func(v, w int) int {
		pathLen := depth[v] + depth[w] // 最后减去 depth[lca] * 2
		cnt := [26]int{}
		if depth[v] > depth[w] {
			v, w = w, v
		}
		for i := 0; i < mx; i++ {
			if (depth[w]-depth[v])>>i&1 > 0 {
				p := pa[w][i]
				for j := 0; j < 26; j++ {
					cnt[j] += p.cnt[j]
				}
				w = p.p
			}
		}
		if w != v {
			for i := mx - 1; i >= 0; i-- {
				if pv, pw := pa[v][i], pa[w][i]; pv.p != pw.p {
					for j := 0; j < 26; j++ {
						cnt[j] += pv.cnt[j] + pw.cnt[j]
					}
					v, w = pv.p, pw.p
				}
			}
			for j := 0; j < 26; j++ {
				cnt[j] += pa[v][0].cnt[j] + pa[w][0].cnt[j]
			}
			v = pa[v][0].p
		}
		// 现在 v 是 LCA
		return pathLen - depth[v] * 2 - slices.Max(cnt[:])
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = f(q[0], q[1])
	}
	return ans
}

// best solution
const N int=int(1e4)+10
const M int=N*2
var (
    f [N][26]int
    fa [N][15]int
    h [N]int
    w [M]int
    e [M]int
    ne [M]int
    q [N]int
    depth [N]int
    idx int
)

func add(a,b,c int){
    e[idx]=b
    w[idx]=c
    ne[idx]=h[a]
    h[a]=idx
    idx++
}

func bfs(root int){
    hh,tt:=0,0
    depth[0],depth[root]=0,1
    q[0]=root
    for hh<=tt {
        t:=q[hh]
        hh++
        for i:=h[t];i!=-1;i=ne[i] {
            j:=e[i]
            if depth[j]>depth[t]+1 {
                depth[j]=depth[t]+1
                tt++
                q[tt]=j
                fa[j][0]=t;
                for k:=1;k<=14;k++ {
                    fa[j][k]=fa[fa[j][k-1]][k-1]
                }
            }
        }
    }
}

func dp(u,father int){
    for i:=h[u];i!=-1;i=ne[i] {
        j:=e[i]
        if j==father {
            continue
        }
        for k:=0;k<26;k++ {
            f[j][k]=f[u][k]
        }
        f[j][w[i]]++
        dp(j,u)
    }
}

func lca(a,b int)int {
    if depth[a]<depth[b] {
        a,b=b,a
    }
    for k:=14;k>=0;k-- {
        if depth[fa[a][k]]>=depth[b] {
            a=fa[a][k]
        }
    }

    if a==b {
        return a
    }
    for k:=14;k>=0;k-- {
        if fa[a][k]!=fa[b][k] {
            a=fa[a][k]
            b=fa[b][k]
        }
    }
    return fa[a][0]
}
func minOperationsQueries1(n int, edges [][]int, queries [][]int) []int {
    for i:=0;i<n+5;i++ {
        h[i]=-1
    }
    idx=0
    for _,t:=range edges {
        u,v,w:=t[0]+1,t[1]+1,t[2]-1
        add(u,v,w)
        add(v,u,w)
    }

    for i:=0;i<n+5;i++ {
        depth[i]=int(2e9)
    }
    bfs(1)
    dp(1,-1)

    ans:=[]int{}
    for _,que:=range queries {
        u,v:=que[0]+1,que[1]+1
        p:=lca(u,v)
        d:=depth[u]+depth[v]-depth[p]*2
        res:=d
        for k:=0;k<26;k++ {
            res=min(res,d-(f[u][k]+f[v][k]-f[p][k]*2))
        }
        ans=append(ans,res)
    }
    return ans
}

// func minOperationsQueries2(n int, edges [][]int, queries [][]int) []int {
// 	m := bits.Len(uint(n))
// 	g := make([][][2]int, n)
// 	f := make([][]int, n)
// 	for i := range f {
// 		f[i] = make([]int, m)
// 	}
// 	p := make([]int, n)
// 	cnt := make([][26]int, n)
// 	cnt[0] = [26]int{}
// 	depth := make([]int, n)
// 	for _, e := range edges {
// 		u, v, w := e[0], e[1], e[2]-1
// 		g[u] = append(g[u], [2]int{v, w})
// 		g[v] = append(g[v], [2]int{u, w})
// 	}
// 	q := []int{0}
// 	for len(q) > 0 {
// 		i := q[0]
// 		q = q[1:]
// 		f[i][0] = p[i]
// 		for j := 1; j < m; j++ {
// 			f[i][j] = f[f[i][j-1]][j-1]
// 		}
// 		for _, nxt := range g[i] {
// 			j, w := nxt[0], nxt[1]
// 			if j != p[i] {
// 				p[j] = i
// 				cnt[j] = [26]int{}
// 				for k := 0; k < 26; k++ {
// 					cnt[j][k] = cnt[i][k]
// 				}
// 				cnt[j][w]++
// 				depth[j] = depth[i] + 1
// 				q = append(q, j)
// 			}
// 		}
// 	}
// 	ans := make([]int, len(queries))
// 	for i, qq := range queries {
// 		u, v := qq[0], qq[1]
// 		x, y := u, v
// 		if depth[x] < depth[y] {
// 			x, y = y, x
// 		}
// 		for j := m - 1; j >= 0; j-- {
// 			if depth[x]-depth[y] >= (1 << j) {
// 				x = f[x][j]
// 			}
// 		}
// 		for j := m - 1; j >= 0; j-- {
// 			if f[x][j] != f[y][j] {
// 				x, y = f[x][j], f[y][j]
// 			}
// 		}
// 		if x != y {
// 			x = p[x]
// 		}
// 		mx := 0
// 		for j := 0; j < 26; j++ {
// 			mx = max(mx, cnt[u][j]+cnt[v][j]-2*cnt[x][j])
// 		}
// 		ans[i] = depth[u] + depth[v] - 2*depth[x] - mx
// 	}
// 	return ans
// }

func main() {
    fmt.Println(minOperationsQueries(
        7, 
        [][]int{[]int{0,1,1},[]int{1,2,1},[]int{2,3,1},[]int{3,4,2},[]int{4,5,2},[]int{5,6,2}},
        [][]int{[]int{0,3},[]int{3,6},[]int{2,6},[]int{0,6}},
    )); // [0,0,1,3]

    fmt.Println(minOperationsQueries(
        8,
        [][]int{[]int{1,2,6},[]int{1,3,4},[]int{2,4,6},[]int{2,5,3},[]int{3,6,6},[]int{3,0,8},[]int{7,0,2}},
        [][]int{[]int{4,6},[]int{0,4},[]int{6,5},[]int{7,4}},
    )); // [1,2,2,3]

    fmt.Println(minOperationsQueries1(
        7, 
        [][]int{[]int{0,1,1},[]int{1,2,1},[]int{2,3,1},[]int{3,4,2},[]int{4,5,2},[]int{5,6,2}},
        [][]int{[]int{0,3},[]int{3,6},[]int{2,6},[]int{0,6}},
    )); // [0,0,1,3]

    fmt.Println(minOperationsQueries1(
        8,
        [][]int{[]int{1,2,6},[]int{1,3,4},[]int{2,4,6},[]int{2,5,3},[]int{3,6,6},[]int{3,0,8},[]int{7,0,2}},
        [][]int{[]int{4,6},[]int{0,4},[]int{6,5},[]int{7,4}},
    )); // [1,2,2,3]

    // fmt.Println(minOperationsQueries2(
    //     7, 
    //     [][]int{[]int{0,1,1},[]int{1,2,1},[]int{2,3,1},[]int{3,4,2},[]int{4,5,2},[]int{5,6,2}},
    //     [][]int{[]int{0,3},[]int{3,6},[]int{2,6},[]int{0,6}},
    // )); // [0,0,1,3]

    // fmt.Println(minOperationsQueries2(
    //     8,
    //     [][]int{[]int{1,2,6},[]int{1,3,4},[]int{2,4,6},[]int{2,5,3},[]int{3,6,6},[]int{3,0,8},[]int{7,0,2}},
    //     [][]int{[]int{4,6},[]int{0,4},[]int{6,5},[]int{7,4}},
    // )); // [1,2,2,3]
}