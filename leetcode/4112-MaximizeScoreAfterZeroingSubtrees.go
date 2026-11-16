package main

// 4112. Maximize Score After Zeroing Subtrees
// You are given an integer array parent of length n representing a rooted tree with nodes labeled from 0 to n - 1, where parent[i] denotes the parent of node i. 
// The tree is rooted at node 0, so parent[0] = -1.

// You are also given two integer arrays op and score, both of length n, where op[i] determines the operation allowed at node i, and score[i] represents its score.

// You need to process each non-leaf node i according to the following rules:
//     1. If op[i] = 0, you can choose at most one child of node i and set the scores of all nodes in that child's subtree to 0.
//     2. If op[i] = 1, you must choose exactly one child of node i and set the scores of all nodes in that child's subtree to 0.

// Each non-leaf node must be processed exactly once. 
// You may process the nodes in any order.

// Return the maximum possible sum of the scores after performing the operations.

// A leaf is a node with no children.

// The subtree rooted at a node consists of that node and all of its descendants.

// Example 1:
// <img alt="" src="https://assets.leetcode.com/uploads/2026/10/10/ex1_ok.gif" style="width: 533px; height: 300px;">​​​​​​​
// Input: parent = [-1,0,0,0,1,2,3,3], op = [1,1,0,0,1,1,0,1], score = [5,6,1,2,3,2,3,-3]
// Output: 16
// Explanation:
// One optimal way is to process the nodes in the following order:
//     Nodes 4, 5, 6, and 7: They are leaf nodes, so no operations are performed on them.
//     Node 1: Since op[1] = 1, we must select one of its children. Choose node 4, setting the scores of all nodes in its subtree (node 4) to 0.
//     Node 2: Since op[2] = 0, we may choose a child or do nothing. For this node, we do nothing.
//     Node 3: Since op[3] = 0, we may choose a child or do nothing. Choose node 7, setting the scores of all nodes in its subtree (node 7) to 0.
//     Node 0: Since op[0] = 1, we must select one of its children. Choose node 2, setting the scores of all nodes in its subtree (nodes 2 and 5) to 0.
// After performing these operations, the scores are [5, 6, 0, 2, 0, 0, 3, 0], giving a total sum of 16.
// No valid sequence of operations can produce a larger sum. Therefore, the answer is 16.

// Example 2:
// <img alt="" src="https://assets.leetcode.com/uploads/2026/10/07/ex2_fix.png" style="width: 220px; height: 180px;">​​​​​​​
// Input: parent = [-1,0,0], op = [0,0,0], score = [1,2,3]
// Output: 6
// Explanation:
//     We only need to process node 0, since node 1 and node 2 are leaves.
//     Since op[0] = 0, we can choose to do nothing.
//     Therefore, all scores are preserved, giving the maximum possible sum of 6.

// Example 3:
// <img alt="" src="https://assets.leetcode.com/uploads/2026/10/07/ex3.png" style="width: 139px; height: 60px;">​​​​​​​
// Input: parent = [-1], op = [1], score = [-10]
// Output: -10
// Explanation:
//     The only node in the tree is a leaf, so we can't process any node.
//     Therefore, the maximum possible score sum is -10.

// Constraints:
//     1 <= n <= 10^{5}
//     n == parent.length == op.length == score.length
//     parent[0] == -1
//     0 <= parent[i] <= n - 1 for all i in [1, n - 1]
//     0 <= op[i] <= 1
//     -10^{4} <= score[i] <= 10^{4}
//     The input is generated such that the array parent represents a valid tree rooted at node 0.

import "fmt"

func maxScore(parent []int, op []int, score []int) int {
    n := len(parent)
    graph := make([][]int, n)
    for i := 1; i < n; i++ {
        p := parent[i]
        graph[p] = append(graph[p], i)
    }
    // 返回子树 x 的最大分数和
    var dfs func(int) int
    dfs = func(x int) int {
        sum := score[x]
        if graph[x] == nil { // x 是叶子
            return sum
        }
        mn := 1 << 61
        for _, y := range graph[x] {
            s := dfs(y)     // 子树 y 的最大分数和
            sum += s        // 累加分数和
            mn = min(mn, s) // 维护最小的分数和
        }
        if mn < 0 || op[x] > 0 {
            sum -= mn // 去掉最小的分数和
        }
        return sum
    }
    return dfs(0)
}

func maxScore1(parent []int, op []int, score []int) int {
    n := len(parent)
    c, m, q := make([]int, n), make([]int,n), make([]int, 0, n)
    for i := 1; i < n; i++ {
        c[parent[i]]++
    }
    for i := range m {
        m[i] = 1 << 61
        if c[i] == 0 {
            q = append(q, i)
        }
    }
    for k := 0; k < len(q); k++ {
        u := q[k]
        f := score[u]
        if m[u] < 1 << 61 {
            f -= min(m[u], op[u] << 61)
        }
        if u == 0 {
            return f
        }
        p := parent[u]
        score[p] += f
        m[p] = min(m[p], f)
        c[p]--
        if c[p] == 0 {
            q = append(q, p)
        }
    }
    return 0
}

func main() {
    // Example 1:
    // <img alt="" src="https://assets.leetcode.com/uploads/2026/10/10/ex1_ok.gif" style="width: 533px; height: 300px;">​​​​​​​
    // Input: parent = [-1,0,0,0,1,2,3,3], op = [1,1,0,0,1,1,0,1], score = [5,6,1,2,3,2,3,-3]
    // Output: 16
    // Explanation:
    // One optimal way is to process the nodes in the following order:
    //     Nodes 4, 5, 6, and 7: They are leaf nodes, so no operations are performed on them.
    //     Node 1: Since op[1] = 1, we must select one of its children. Choose node 4, setting the scores of all nodes in its subtree (node 4) to 0.
    //     Node 2: Since op[2] = 0, we may choose a child or do nothing. For this node, we do nothing.
    //     Node 3: Since op[3] = 0, we may choose a child or do nothing. Choose node 7, setting the scores of all nodes in its subtree (node 7) to 0.
    //     Node 0: Since op[0] = 1, we must select one of its children. Choose node 2, setting the scores of all nodes in its subtree (nodes 2 and 5) to 0.
    // After performing these operations, the scores are [5, 6, 0, 2, 0, 0, 3, 0], giving a total sum of 16.
    // No valid sequence of operations can. Therefore, the. Therefore, the answer is 16.
    fmt.Println(maxScore([]int{-1,0,0,0,1,2,3,3}, []int{1,1,0,0,1,1,0,1}, []int{5,6,1,2,3,2,3,-3})) // 16
    // Example 2:
    // <img alt="" src="https://assets.leetcode.com/uploads/2026/10/07/ex2_fix.png" style="width: 220px; height: 180px;">​​​​​​​
    // Input: parent = [-1,0,0], op = [0,0,0], score = [1,2,3]
    // Output: 6
    // Explanation:
    //     We only need to process node 0, since node 1 and node 2 are leaves.
    //     Since op[0] = 0, we can choose to do nothing.
    //     Therefore, all scores are preserved, giving the maximum possible sum of 6.
    fmt.Println(maxScore([]int{-1,0,0}, []int{0,0,0}, []int{1,2,3})) // 6
    // Example 3:
    // <img alt="" src="https://assets.leetcode.com/uploads/2026/10/07/ex3.png" style="width: 139px; height: 60px;">​​​​​​​
    // Input: parent = [-1], op = [1], score = [-10]
    // Output: -10
    // Explanation:
    //     The only node in the tree is a leaf, so we can't process any node.
    //     Therefore, the maximum possible score sum is -10.'
    fmt.Println(maxScore([]int{-1}, []int{1}, []int{-10})) // -10

    fmt.Println(maxScore1([]int{-1,0,0,0,1,2,3,3}, []int{1,1,0,0,1,1,0,1}, []int{5,6,1,2,3,2,3,-3})) // 16
    fmt.Println(maxScore1([]int{-1,0,0}, []int{0,0,0}, []int{1,2,3})) // 6
    fmt.Println(maxScore1([]int{-1}, []int{1}, []int{-10})) // -10
}