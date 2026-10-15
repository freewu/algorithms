package main

// 1932. Merge BSTs to Create Single BST
// You are given n BST (binary search tree) root nodes for n separate BSTs stored in an array trees (0-indexed). 
// Each BST in trees has at most 3 nodes, and no two roots have the same value. 
// In one operation, you can:
//     Select two distinct indices i and j such that the value stored at one of the leaves of trees[i] is equal to the root value of trees[j].
//     Replace the leaf node in trees[i] with trees[j].
//     Remove trees[j] from trees.

// Return the root of the resulting BST if it is possible to form a valid BST after performing n - 1 operations, 
// or null if it is impossible to create a valid BST.

// A BST (binary search tree) is a binary tree where each node satisfies the following property:
//     Every node in the node's left subtree has a value strictly less than the node's value.
//     Every node in the node's right subtree has a value strictly greater than the node's value.

// A leaf is a node that has no children.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/08/d1.png" />
// Input: trees = [[2,1],[3,2,5],[5,4]]
// Output: [3,2,5,1,null,4]
// Explanation:
// In the first operation, pick i=1 and j=0, and merge trees[0] into trees[1].
// Delete trees[0], so trees = [[3,2,5,1],[5,4]].
// <img src="https://assets.leetcode.com/uploads/2021/06/24/diagram.png">
// In the second operation, pick i=0 and j=1, and merge trees[1] into trees[0].
// Delete trees[1], so trees = [[3,2,5,1,null,4]].
// <img src="https://assets.leetcode.com/uploads/2021/06/24/diagram-2.png" />
// The resulting tree, shown above, is a valid BST, so return its root.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/08/d2.png" />
// Input: trees = [[5,3,8],[3,2,6]]
// Output: []
// Explanation:
// Pick i=0 and j=1 and merge trees[1] into trees[0].
// Delete trees[1], so trees = [[5,3,8,2,6]].
// <img src="https://assets.leetcode.com/uploads/2021/06/24/diagram-3.png" />
// The resulting tree is shown above. This is the only valid operation that can be performed, but the resulting tree is not a valid BST, so return null.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/06/08/d3.png" />
// Input: trees = [[5,4],[3]]
// Output: []
// Explanation: It is impossible to perform any operations.

// Constraints:
//     n == trees.length
//     1 <= n <= 5 * 10^4
//     The number of nodes in each tree is in the range [1, 3].
//     Each node in the input may have children but no grandchildren.
//     No two roots of trees have the same value.
//     All the trees in the input are valid BSTs.
//     1 <= TreeNode.val <= 5 * 10^4.

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
func canMerge(trees []*TreeNode) *TreeNode {
    // Collect the leaves
    leaves, nodeMap := make(map[int]struct{}), make(map[int]*TreeNode)
    for _, tree := range trees {
        nodeMap[tree.Val] = tree
        if tree.Left != nil {
            leaves[tree.Left.Val] = struct{}{}
        }
        if tree.Right != nil {
            leaves[tree.Right.Val] = struct{}{}
        }
    }
    // Decide the root of the resulting tree
    var res *TreeNode
    for _, tree := range trees {
        if _, exists := leaves[tree.Val]; !exists {
            res = tree
            break
        }
    }
    if res == nil { return nil }
    var traverse func(root *TreeNode, nodeMap map[int]*TreeNode, mn, mx int) bool
    traverse = func(root *TreeNode, nodeMap map[int]*TreeNode, mn, mx int) bool {
        if root == nil {  return true }
        if root.Val <= mn || root.Val >= mx { return false }
        if root.Left == nil && root.Right == nil {
            if next, exists := nodeMap[root.Val]; exists && root != next {
                root.Left = next.Left
                root.Right = next.Right
                delete(nodeMap, root.Val)
            }
        }
        return traverse(root.Left, nodeMap, mn, root.Val) && traverse(root.Right, nodeMap, root.Val, mx)
    }
    if traverse(res, nodeMap, -1 << 63, 1 << 63-1) && len(nodeMap) == 1 {
        return res
    }
    return nil
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/08/d1.png" />
    // Input: trees = [[2,1],[3,2,5],[5,4]]
    // Output: [3,2,5,1,null,4]
    // Explanation:
    // In the first operation, pick i=1 and j=0, and merge trees[0] into trees[1].
    // Delete trees[0], so trees = [[3,2,5,1],[5,4]].
    // <img src="https://assets.leetcode.com/uploads/2021/06/24/diagram.png">
    // In the second operation, pick i=0 and j=1, and merge trees[1] into trees[0].
    // Delete trees[1], so trees = [[3,2,5,1,null,4]].
    // <img src="https://assets.leetcode.com/uploads/2021/06/24/diagram-2.png" />
    // The resulting tree, shown above, is a valid BST, so return its root.
    arr1 := []*TreeNode{
        buildTree([]*int{ptr(2), ptr(1)}),
        buildTree([]*int{ptr(3), ptr(2), ptr(5)}),
        buildTree([]*int{ptr(5), ptr(4)}),
    }
    printTree(arr1[0])
    printTree(arr1[1])
    printTree(arr1[2])
    fmt.Println(canMerge(arr1)) // [3,2,5,1,null,4] &{3 0xc000008090 0xc0000080a8}
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/08/d2.png" />
    // Input: trees = [[5,3,8],[3,2,6]]
    // Output: []
    // Explanation:
    // Pick i=0 and j=1 and merge trees[1] into trees[0].
    // Delete trees[1], so trees = [[5,3,8,2,6]].
    // <img src="https://assets.leetcode.com/uploads/2021/06/24/diagram-3.png" />
    // The resulting tree is shown above. This is the only valid operation that can be performed, but the resulting tree is not a valid BST, so return null.
    arr2 := []*TreeNode{
        buildTree([]*int{ptr(5), ptr(3), ptr(8)}),
        buildTree([]*int{ptr(3), ptr(2), ptr(6)}),
    }
    printTree(arr2[0])
    printTree(arr2[1])
    fmt.Println(canMerge(arr2)) // [] nil
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/06/08/d3.png" />
    // Input: trees = [[5,4],[3]]
    // Output: []
    // Explanation: It is impossible to perform any operations.
    arr3 := []*TreeNode{
        buildTree([]*int{ptr(5), ptr(4)}),
        buildTree([]*int{ptr(3)}),
    }
    printTree(arr3[0])
    printTree(arr3[1])
    fmt.Println(canMerge(arr3)) // [] nil
}