package main

// 2385. Amount of Time for Binary Tree to Be Infected
// You are given the root of a binary tree with unique values, and an integer start.
// At minute 0, an infection starts from the node with value start.

// Each minute, a node becomes infected if:
//     The node is currently uninfected.
//     The node is adjacent to an infected node.

// Return the number of minutes needed for the entire tree to be infected.

// Example 1:
//                      1
//                    /   \
//                   5    [3]
//                    \   / \
//                     4 10  6
//                    / \
//                   9   2
// <img src="https://assets.leetcode.com/uploads/2022/06/25/image-20220625231744-1.png" />
// Input: root = [1,5,3,null,4,10,6,9,2], start = 3
// Output: 4
// Explanation: The following nodes are infected during:
// - Minute 0: Node 3
// - Minute 1: Nodes 1, 10 and 6
// - Minute 2: Node 5
// - Minute 3: Node 4
// - Minute 4: Nodes 9 and 2
// It takes 4 minutes for the whole tree to be infected so we return 4.

// Example 2:
// Input: root = [1], start = 1
// Output: 0
// Explanation: At minute 0, the only node in the tree is infected so we return 0.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     1 <= Node.val <= 10^5
//     Each node has a unique value.
//     A node with a value of start exists in the tree.

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// == 二叉树构建 & 打印输出 start ==================================================================
func ptr(v int) *int {
    return &v
}

// buildTree 按层序数组（LeetCode 格式，nil 表示 null）创建二叉树
func buildTree(nums []*int) *TreeNode {
    if len(nums) == 0 || nums[0] == nil {
        return nil
    }
    root := &TreeNode{Val: *nums[0]}
    queue := []*TreeNode{root}
    i := 1
    for len(queue) > 0 && i < len(nums) {
        node := queue[0]
        queue = queue[1:]
        if nums[i] != nil {
            node.Left = &TreeNode{Val: *nums[i]}
            queue = append(queue, node.Left)
        }
        i++
        if i < len(nums) && nums[i] != nil {
            node.Right = &TreeNode{Val: *nums[i]}
            queue = append(queue, node.Right)
        }
        i++
    }
    return root
}

// printTree 自顶向下打印二叉树，父子之间用 / \ 连线
// 第 r 层（根为第 1 层）有 2^(r-1) 个位置，每个节点占 block 列，
// 因此节点中心列 = (2j+1) * 2^(h-r) * block / 2（j 为该层从左往右的序号）
func printTree(root *TreeNode) {
    if root == nil {
        fmt.Println("(empty tree)")
        return
    }
    // itoa 手写整数转字符串，替代 strconv.Itoa
    itoa := func(v int) string {
        if v == 0 {
            return "0"
        }
        neg := v < 0
        u := uint(v)
        if neg {
            u = -u // 补码取绝对值，v 为最小整数时同样安全
        }
        buf := [24]byte{}
        i := len(buf)
        for u > 0 {
            i--
            buf[i] = byte('0' + u%10)
            u /= 10
        }
        if neg {
            i--
            buf[i] = '-'
        }
        return string(buf[i:])
    }
    // blankLine 生成 n 个空格的画布行（替代 strings.Repeat）
    blankLine := func(n int) []byte {
        line := make([]byte, n)
        for i := range line {
            line[i] = ' '
        }
        return line
    }
    // treeHeight 返回二叉树的层数（空树为 0，根节点为第 1 层）
    var treeHeight func(root *TreeNode) int
    treeHeight = func (root *TreeNode) int {
        if root == nil {
            return 0
        }
        l, r := treeHeight(root.Left), treeHeight(root.Right)
        if l < r {
            l = r
        }
        return l + 1
    }
    h := treeHeight(root)
    // block：每个节点占用的列宽，取最长值的宽度 + 1，并保证为偶数使中心落在整数列
    block := 2
    var scan func(node *TreeNode)
    scan = func(node *TreeNode) {
        if node == nil {
            return
        }
        if w := len(itoa(node.Val)) + 1; w > block {
            block = w
        }
        scan(node.Left)
        scan(node.Right)
    }
    scan(root)
    if block%2 != 0 {
        block++
    }

    // 画布：值占 1 行、连线占 1 行，共 2h-1 行
    width := (1 << h) * block
    rows := make([][]byte, 2*h-1)
    for i := range rows {
        rows[i] = blankLine(width)
    }
    put := func(row, col int, s string) {
        start := col - len(s)/2
        for i := 0; i < len(s); i++ {
            if p := start + i; p >= 0 && p < width {
                rows[row][p] = s[i]
            }
        }
    }
    var draw func(node *TreeNode, r, j int)
    draw = func(node *TreeNode, r, j int) {
        if node == nil {
            return
        }
        col := (2*j + 1) * (1 << (h - r)) * block / 2 // 当前节点中心列
        put(2*(r-1), col, itoa(node.Val))
        if r == h {
            return
        }
        d := (1 << (h - r - 1)) * block / 2 // 父子之间的列距
        off := (d + 1) / 2                  // 连线相对父节点中心的偏移（取行间中点）
        if node.Left != nil {
            put(2*r-1, col-off, "/")
            draw(node.Left, r+1, 2*j)
        }
        if node.Right != nil {
            put(2*r-1, col+off, "\\")
            draw(node.Right, r+1, 2*j+1)
        }
    }
    draw(root, 1, 0)

    for _, line := range rows {
        end := len(line)
        for end > 0 && line[end-1] == ' ' {
            end--
        }
        fmt.Println(string(line[:end]))
    }
}
// == 二叉树构建 & 打印输出 end ====================================================================

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// dfs
func amountOfTime(root *TreeNode, start int) int {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(node *TreeNode, start int) (int, int)
    dfs = func(node *TreeNode, start int) (int, int) {
        if node == nil { return 0, 0 }
        rDay, rTotal := dfs(node.Right, start)
        lDay, lTotal := dfs(node.Left, start)
        switch {
            case node.Val == start:
                return 1, max(rTotal, lTotal)
            case rDay > 0:
                return rDay + 1, max(rDay + lTotal, rTotal)
            case lDay > 0:
                return lDay + 1, max(lDay + rTotal, lTotal)
            default:
                return -1, max(lTotal, rTotal) + 1
        }
    }
    _, res := dfs(root, start)
    return res
}

// bfs
func amountOfTime1(root *TreeNode, start int) int {
    graph := make(map[int][]int)
    var buildGraph func (node *TreeNode, parent int, graph map[int][]int) 
    buildGraph = func (node *TreeNode, parent int, graph map[int][]int) {
        if node == nil { return }
        if parent != -1 {
            graph[node.Val] = append(graph[node.Val], parent)
            graph[parent] = append(graph[parent], node.Val)
        }
        buildGraph(node.Left, node.Val, graph)
        buildGraph(node.Right, node.Val, graph)
    }
    bfs := func (graph map[int][]int, start int) int {
        visited := make(map[int]bool)
        queue := []int{ start }
        mins := -1
        for len(queue) > 0 {
            n := len(queue)
            for i := 0; i < n; i++ {
                node := queue[0]
                queue = queue[1:]
                if visited[node] {
                    continue
                }
                visited[node] = true
                for _, adjacent := range graph[node] {
                    if !visited[adjacent] {
                        queue = append(queue, adjacent)
                    }
                }
            }
            mins++
        }
        return mins
    }
    buildGraph(root, -1, graph)
    return bfs(graph, start)
}


func main() {
    // Explanation: The following nodes are infected during:
    // - Minute 0: Node 3
    // - Minute 1: Nodes 1, 10 and 6
    // - Minute 2: Node 5
    // - Minute 3: Node 4
    // - Minute 4: Nodes 9 and 2
    // It takes 4 minutes for the whole tree to be infected so we return 4.
    tree1 := buildTree([]*int{ptr(1), ptr(5), ptr(3), nil, ptr(4), ptr(10), ptr(6), ptr(9), ptr(2)})
    printTree(tree1)
    fmt.Println(amountOfTime(tree1, 3)) // 4
    // Explanation: At minute 0, the only node in the tree is infected so we return 0.
    tree2 := buildTree([]*int{ptr(1)})
    printTree(tree2)
    fmt.Println(amountOfTime(tree2, 1)) // 0

    fmt.Println(amountOfTime1(tree1, 3)) // 4
    fmt.Println(amountOfTime1(tree2, 1)) // 0
}