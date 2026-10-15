package main

// 2458. Height of Binary Tree After Subtree Removal Queries
// You are given the root of a binary tree with n nodes. 
// Each node is assigned a unique value from 1 to n. 
// You are also given an array queries of size m.

// You have to perform m independent queries on the tree where in the ith query you do the following:
//     Remove the subtree rooted at the node with the value queries[i] from the tree. 
//     It is guaranteed that queries[i] will not be equal to the value of the root.

// Return an array answer of size m where answer[i] is the height of the tree after performing the ith query.

// Note:
//     The queries are independent, so the tree returns to its initial state after each query.
//     The height of a tree is the number of edges in the longest simple path from the root to some node in the tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/09/07/binaryytreeedrawio-1.png" />
// Input: root = [1,3,4,2,null,6,5,null,null,null,null,null,7], queries = [4]
// Output: [2]
// Explanation: The diagram above shows the tree after removing the subtree rooted at node with value 4.
// The height of the tree is 2 (The path 1 -> 3 -> 2).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/09/07/binaryytreeedrawio-2.png" />
// Input: root = [5,8,9,2,1,3,7,4,6], queries = [3,2,4,8]
// Output: [3,2,3,2]
// Explanation: We have the following queries:
// - Removing the subtree rooted at node with value 3. The height of the tree becomes 3 (The path 5 -> 8 -> 2 -> 4).
// - Removing the subtree rooted at node with value 2. The height of the tree becomes 2 (The path 5 -> 8 -> 1).
// - Removing the subtree rooted at node with value 4. The height of the tree becomes 3 (The path 5 -> 8 -> 2 -> 6).
// - Removing the subtree rooted at node with value 8. The height of the tree becomes 2 (The path 5 -> 9 -> 3).

// Constraints:
//     The number of nodes in the tree is n.
//     2 <= n <= 10^5
//     1 <= Node.val <= n
//     All the values in the tree are unique.
//     m == queries.length
//     1 <= m <= min(n, 10^4)
//     1 <= queries[i] <= n
//     queries[i] != root.val

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
func treeQueries(root *TreeNode, queries []int) []int {
    depthMap, heightMap, h1Map, h2Map := make(map[int]int),make(map[int]int),make(map[int]int), make(map[int]int)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var calculateHeightDepth func(root *TreeNode, depth int) int 
    calculateHeightDepth = func(root *TreeNode, depth int) int {
        if root == nil {
            return 0
        }
        depthMap[root.Val] = depth
        leftHeight, rightHeight := calculateHeightDepth(root.Left, depth + 1), calculateHeightDepth(root.Right, depth + 1)
        heightMap[root.Val] = max(leftHeight, rightHeight) + 1
        if h1Map[depth] < heightMap[root.Val] {
            h2Map[depth] = h1Map[depth]
            h1Map[depth] = heightMap[root.Val]
        } else if h2Map[depth] < heightMap[root.Val] {
            h2Map[depth] = heightMap[root.Val]
        }
        return heightMap[root.Val]
    }
    calculateHeightDepth(root, 0)
    res := []int{}
    for _, query := range queries {
        depth := depthMap[query]
        if h1Map[depth] == heightMap[query] {
            res = append(res, depth + h2Map[depth]-1)
        } else {
            res = append(res, depth + h1Map[depth]-1)
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/09/07/binaryytreeedrawio-1.png" />
    // Input: root = [1,3,4,2,null,6,5,null,null,null,null,null,7], queries = [4]
    // Output: [2]
    // Explanation: The diagram above shows the tree after removing the subtree rooted at node with value 4.
    // The height of the tree is 2 (The path 1 -> 3 -> 2).
    tree1 := buildTree([]*int{ptr(1), ptr(3), ptr(4), ptr(2), nil, ptr(6), ptr(5), nil, nil, nil, nil, nil, ptr(7)})
    printTree(tree1)
    fmt.Println(treeQueries(tree1,[]int{4})) // [2]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/09/07/binaryytreeedrawio-2.png" />
    // Input: root = [5,8,9,2,1,3,7,4,6], queries = [3,2,4,8]
    // Output: [3,2,3,2]
    // Explanation: We have the following queries:
    // - Removing the subtree rooted at node with value 3. The height of the tree becomes 3 (The path 5 -> 8 -> 2 -> 4).
    // - Removing the subtree rooted at node with value 2. The height of the tree becomes 2 (The path 5 -> 8 -> 1).
    // - Removing the subtree rooted at node with value 4. The height of the tree becomes 3 (The path 5 -> 8 -> 2 -> 6).
    // - Removing the subtree rooted at node with value 8. The height of the tree becomes 2 (The path 5 -> 9 -> 3).
    tree2 := buildTree([]*int{ptr(5), ptr(8), ptr(9), ptr(2), ptr(1), ptr(3), ptr(7), ptr(4), ptr(6)})
    printTree(tree2)
    fmt.Println(treeQueries(tree2,[]int{3,2,4,8})) // [3,2,3,2]
}