package main

// 2378. Choose Edges to Maximize Score in a Tree
// You are given a weighted tree consisting of n nodes numbered from 0 to n - 1.

// The tree is rooted at node 0 and represented with a 2D array edges of size n where edges[i] = [pari, weighti] indicates that node pari is the parent of node i, and the edge between them has a weight equal to weighti. 
// Since the root does not have a parent, you have edges[0] = [-1, -1].

// Choose some edges from the tree such that no two chosen edges are adjacent and the sum of the weights of the chosen edges is maximized.

// Return the maximum sum of the chosen edges.

// Note:
//     You are allowed to not choose any edges in the tree, the sum of weights in this case will be 0.
//     Two edges Edge1 and Edge2 in the tree are adjacent if they have a common node.
//     In other words, they are adjacent if Edge1 connects nodes a and b and Edge2 connects nodes b and c.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/08/16/treedrawio.png" />
// Input: edges = [[-1,-1],[0,5],[0,10],[2,6],[2,4]]
// Output: 11
// Explanation: The above diagram shows the edges that we have to choose colored in red.
// The total score is 5 + 6 = 11.
// It can be shown that no better score can be obtained.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/08/17/treee1293712983719827.png" />
// Input: edges = [[-1,-1],[0,5],[0,-6],[0,7]]
// Output: 7
// Explanation: We choose the edge with weight 7.
// Note that we cannot choose more than one edge because all edges are adjacent to each other.

// Constraints:
//     n == edges.length
//     1 <= n <= 10^5
//     edges[i].length == 2
//     par0 == weight0 == -1
//     0 <= pari <= n - 1 for all i >= 1.
//     pari != i
//     -10^6 <= weighti <= 10^6 for all i >= 1.
//     edges represents a valid tree.

import "fmt"

// 树形dp
// 选择父节点下的一条边e时，e连接的子节点的边都不可选，未选择的子节点的边可选。
// 记忆化搜索即可
func maxScore(edges [][]int) int64 {
    n := len(edges)
    children := make([][]int, n)
    for i, edge := range edges {
        if i ==0{
            continue
        }
        children[edge[0]] = append(children[edge[0]], i)
    }
    memo := map[int]int64{}
    var dp func(i, selected int) int64
    dp = func(i, selected int) int64 {
        if res, ok := memo[i * selected]; ok {
            return res
        }
        res, sum, all, sel := int64(0), int64(0), []int64{}, []int64{}
        for _, child := range children[i] {
            t := dp(child, -1)
            all = append(all, t)
            sel = append(sel, dp(child, 1))
            sum += t
        }
        res = sum
        if selected !=1{
            for j := 0; j < len(all); j++ {
                if sum - all[j]+ sel[j] + int64(edges[children[i][j]][1]) > res {
                    res = sum - all[j]+ sel[j] + int64(edges[children[i][j]][1])
                }
            }
        }
        memo[i * selected] = res
        return res
    }
    return dp(0, -1)
}

func maxScore1(edges [][]int) int64 {
    // 树形dp (树上最大独立集), 同树上打家劫舍
    // 子节点返回两个值, 都是最大子树边权 一个是用到了子节点的  一个是没有用到子节点的
    n := len(edges)
    type Pair struct{ To, Weight int }
    g := make([][]Pair, n)
    for i := 1; i < n; i++ { // 建树
        pa, weight := edges[i][0], edges[i][1]
        g[pa] = append(g[pa], Pair{i, weight})
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(sn int) (use, noUse int) // 返回以node为root的子树的最大边权和. use:用了root noUse:没有用root,
    dfs = func(sn int) (use, noUse int) {
        if len(g[sn]) == 0 { // 叶子节点没有被选择状态
            return -1 << 31, 0
        }
        // 重大bug!! 对于use, 如果use了, 它就占据了父节点,那么其它子节点就不能选边了=>但是只是不能选边了,子节点还是可以用use,noUse中的最大值
        // 本题为最大独立集(而非最大支配集),use代表用了一条到子节点的边(那么其它的节点都是max(use,noUse)),
        // 也就是max(one(child_noUse+边权) + other(child_mx)) = all(child_mx) + max(diff(child_noUse+边权-child_mx)
        diff := 0
        for _, nx := range g[sn] {
            u, nU := dfs(nx.To)
            mxNx := max(u, nU)
            use += mxNx
            diff = max(diff, nx.Weight + nU - mxNx)
            noUse += mxNx
        }
        use += diff
        return use, noUse
    }
    return int64(max(dfs(0)))
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/08/16/treedrawio.png" />
    // Input: edges = [[-1,-1],[0,5],[0,10],[2,6],[2,4]]
    // Output: 11
    // Explanation: The above diagram shows the edges that we have to choose colored in red.
    // The total score is 5 + 6 = 11.
    // It can be shown that no better score can be obtained.
    fmt.Println(maxScore([][]int{{-1,-1},{0,5},{0,10},{2,6},{2,4}})) // 11
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/08/17/treee1293712983719827.png" />
    // Input: edges = [[-1,-1],[0,5],[0,-6],[0,7]]
    // Output: 7
    // Explanation: We choose the edge with weight 7.
    // Note that we cannot choose more than one edge because all edges are adjacent to each other.
    fmt.Println(maxScore([][]int{{-1,-1},{0,5},{0,-6},{0,7}})) // 7

    fmt.Println(maxScore1([][]int{{-1,-1},{0,5},{0,10},{2,6},{2,4}})) // 11
    fmt.Println(maxScore1([][]int{{-1,-1},{0,5},{0,-6},{0,7}})) // 7
}