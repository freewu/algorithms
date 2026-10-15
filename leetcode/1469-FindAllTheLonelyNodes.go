package main

// 1469. Find All The Lonely Nodes
// In a binary tree, a lonely node is a node that is the only child of its parent node. 
// The root of the tree is not lonely because it does not have a parent node.

// Given the root of a binary tree, return an array containing the values of all lonely nodes in the tree. 
// Return the list in any order.

// Example 1:
//             1
//           /   \
//          2     3
//           \ 
//             (4)
// <img src="https://assets.leetcode.com/uploads/2020/06/03/e1.png">
// Input: root = [1,2,3,null,4]
// Output: [4]
// Explanation: Light blue node is the only lonely node.
// Node 1 is the root and is not lonely.
// Nodes 2 and 3 have the same parent and are not lonely.

// Example 2:
//             7
//           /    \
//          1      4
//         /      /  \
//       (6)     5    3
//                      \
//                      (2) 
// <img src="https://assets.leetcode.com/uploads/2020/06/03/e2.png">
// Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2]
// Output: [6,2]
// Explanation: Light blue nodes are lonely nodes.
// Please remember that order doesn't matter, [2,6] is also an acceptable answer.

// Example 3:
//             11
//           /     \
//          99     88
//         /         \
//       (77)        (66)
//       /             \
//     (55)            (44)
//     /                 \
//   (33)                (22)
// <img src="https://assets.leetcode.com/uploads/2020/06/03/tree.png">
// Input: root = [11,99,88,77,null,null,66,55,null,null,44,33,null,null,22]
// Output: [77,55,33,66,44,22]
// Explanation: Nodes 99 and 88 share the same parent. Node 11 is the root.
// All other nodes are lonely.

// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     1 <= Node.val <= 10^6

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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func getLonelyNodes(root *TreeNode) []int {
    res := []int{}
    if root == nil || root.Left == nil && root.Right == nil {
        return res
    }
    if root.Left == nil { // 如果没有左子节点,把右边子节点的值写入返回数组
        res = append(res, root.Right.Val)
    } else if root.Right == nil { // 如果没有右子节点,把左边子节点的值写入返回数组
        res = append(res, root.Left.Val)
    }
    res = append(res, getLonelyNodes(root.Left)...)
    res = append(res, getLonelyNodes(root.Right)...)
    return res
}

func getLonelyNodes1(root *TreeNode) []int {
    res := []int{}
    var traverse func(root *TreeNode)
    traverse = func(root *TreeNode) {
        if root == nil {
            return
        }
        if root.Left != nil && root.Right == nil {
            res = append(res, root.Left.Val)
        }
        if root.Left == nil && root.Right != nil {
            res = append(res, root.Right.Val)
        }
        traverse(root.Left)
        traverse(root.Right)
    }
    traverse(root)
    return res
}

func main() {
    // Example 1:
    //             1
    //           /   \
    //          2     3
    //           \ 
    //            (4)
    // <img src="https://assets.leetcode.com/uploads/2020/06/03/e1.png">
    // Input: root = [1,2,3,null,4]
    // Output: [4]
    // Explanation: Light blue node is the only lonely node.
    // Node 1 is the root and is not lonely.
    // Nodes 2 and 3 have the same parent and are not lonely.
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3), nil, ptr(4)})
    printTree(tree1)
    fmt.Println(getLonelyNodes(tree1)) // [4]
    // Example 2:
    //             7
    //           /    \
    //          1      4
    //         /      /  \
    //       (6)     5    3
    //                      \
    //                      (2) 
    // <img src="https://assets.leetcode.com/uploads/2020/06/03/e2.png">
    // Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2]
    // Output: [6,2]
    // Explanation: Light blue nodes are lonely nodes.
    // Please remember that order doesn't matter, [2,6] is also an acceptable answer.
    tree2 := buildTree([]*int{ptr(7), ptr(1), ptr(4), ptr(6), nil, ptr(5), ptr(3), nil, nil, nil, nil, nil, ptr(2)})
    printTree(tree2)
    fmt.Println(getLonelyNodes(tree2)) // [6,2]
    // Example 3:
    //             11
    //           /     \
    //          99     88
    //         /         \
    //       (77)        (66)
    //       /             \
    //     (55)            (44)
    //     /                 \
    //   (33)                (22)
    // <img src="https://assets.leetcode.com/uploads/2020/06/03/tree.png">
    // Input: root = [11,99,88,77,null,null,66,55,null,null,44,33,null,null,22]
    // Output: [77,55,33,66,44,22]
    // Explanation: Nodes 99 and 88 share the same parent. Node 11 is the root.
    // All other nodes are lonely.
    tree3 := buildTree([]*int{ptr(11), ptr(99), ptr(88), ptr(77), nil, nil, ptr(66), ptr(55), nil, nil, ptr(44), ptr(33), nil, nil, ptr(22)})
    printTree(tree3)
    fmt.Println(getLonelyNodes(tree3)) // [77,55,33,66,44,22]

    fmt.Println(getLonelyNodes1(tree1)) // [4]
    fmt.Println(getLonelyNodes1(tree2)) // [6,2]
    fmt.Println(getLonelyNodes1(tree3)) // [77,55,33,66,44,22]
}