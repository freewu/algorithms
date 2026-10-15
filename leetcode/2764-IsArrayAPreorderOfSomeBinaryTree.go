package main

// 2764. Is Array a Preorder of Some ‌Binary Tree
// Given a 0-indexed integer 2D array nodes, 
// your task is to determine if the given array represents the preorder traversal of some binary tree.

// For each index i, nodes[i] = [id, parentId], 
// where id is the id of the node at the index i and parentId is the id of its parent in the tree 
// (if the node has no parent, then parentId == -1).

// Return true if the given array represents the preorder traversal of some tree, and false otherwise.

// Note: the preorder traversal of a tree is a recursive way to traverse a tree in which we first visit the current node, 
// then we do the preorder traversal for the left child, and finally, we do it for the right child.

// Example 1:
// Input: nodes = [[0,-1],[1,0],[2,0],[3,2],[4,2]]
// Output: true
// Explanation: The given nodes make the tree in the picture below.
// We can show that this is the preorder traversal of the tree, first we visit node 0, then we do the preorder traversal of the right child which is [1], then we do the preorder traversal of the left child which is [2,3,4].
// <img src="https://assets.leetcode.com/uploads/2023/07/04/1.png" />
//         0
//       /   \
//      1     2
//          /    \
//         3      4

// Example 2:
// Input: nodes = [[0,-1],[1,0],[2,0],[3,1],[4,1]]
// Output: false
// Explanation: The given nodes make the tree in the picture below.
// For the preorder traversal, first we visit node 0, then we do the preorder traversal of the right child which is [1,3,4], but we can see that in the given order, 2 comes between 1 and 3, so, it's not the preorder traversal of the tree.
// <img src="https://assets.leetcode.com/uploads/2023/07/04/2.png" />
//         0
//       /   \
//      1     2
//     /  \
//    3    4

// Constraints:
//     1 <= nodes.length <= 10^5
//     nodes[i].length == 2
//     0 <= nodes[i][0] <= 10^5
//     -1 <= nodes[i][1] <= 10^5
//     The input is generated such that nodes make a binary tree.

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

func isPreorder(nodes [][]int) bool {
    mp := make(map[int]*TreeNode)
    for _, v := range nodes {
        id := v[0]
        mp[id] = &TreeNode{ Val: id}
    }
    // step1 还原二叉树
    var root *TreeNode
    for _, v := range nodes {
        id, parentId := v[0], v[1]
        if parentId == -1 {
            root = mp[id]
        } else {
            parent := mp[parentId]
            if parent.Left == nil {
                parent.Left = mp[id]
                continue
            }
            if parent.Right == nil {
                parent.Right = mp[id]
                continue
            }
        }
    }
    // step2 得出其先序遍历
    preOrderNums := []int{}
    var preOrder func(root *TreeNode)
    preOrder = func(root *TreeNode) {
        if root == nil { return }
        preOrderNums = append(preOrderNums, root.Val)
        preOrder(root.Left)
        preOrder(root.Right)
    }
    preOrder(root)
    // 与 nodes 中的遍历进行比较
    for i, v := range nodes {
        if preOrderNums[i] != v[0] {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: nodes = [[0,-1],[1,0],[2,0],[3,2],[4,2]]
    // Output: true
    // Explanation: The given nodes make the tree in the picture below.
    // We can show that this is the preorder traversal of the tree, first we visit node 0, then we do the preorder traversal of the right child which is [1], then we do the preorder traversal of the left child which is [2,3,4].
    // <img src="https://assets.leetcode.com/uploads/2023/07/04/1.png" />
    //         0
    //       /   \
    //      1     2
    //          /    \
    //         3      4
    fmt.Println(isPreorder([][]int{{0,-1},{1,0},{2,0},{3,2},{4,2}})) // true
    // Example 2:
    // Input: nodes = [[0,-1],[1,0],[2,0],[3,1],[4,1]]
    // Output: false
    // Explanation: The given nodes make the tree in the picture below.
    // For the preorder traversal, first we visit node 0, then we do the preorder traversal of the right child which is [1,3,4], but we can see that in the given order, 2 comes between 1 and 3, so, it's not the preorder traversal of the tree.
    // <img src="https://assets.leetcode.com/uploads/2023/07/04/2.png" />
    //         0
    //       /   \
    //      1     2
    //     /  \
    //    3    4
    fmt.Println(isPreorder([][]int{{0,-1},{1,0},{2,0},{3,1},{4,1}})) // false
}