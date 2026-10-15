package main

// 968. Binary Tree Cameras
// You are given the root of a binary tree. 
// We install cameras on the tree nodes where each camera at a node can monitor its parent, itself, and its immediate children.
// Return the minimum number of cameras needed to monitor all nodes of the tree.

// Example 1:
//         (0)
//         /
//       (c)
//      /   \   
//    (0)   (0)
// <img src="https://assets.leetcode.com/uploads/2018/12/29/bst_cameras_01.png" />
// Input: root = [0,0,null,0,0]
// Output: 1
// Explanation: One camera is enough to monitor all nodes if placed as shown.

// Example 2:
//             (0)
//             /
//           (c)
//           /
//         (0)
//         /
//       (c)
//         \ 
//          (0)
// <img src="https://assets.leetcode.com/uploads/2018/12/29/bst_cameras_02.png" />
// Input: root = [0,0,null,0,null,0,null,null,0]
// Output: 2
// Explanation: At least two cameras are needed to monitor all nodes of the tree. The above image shows one of the valid configurations of camera placement.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     Node.val == 0

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
func minCameraCover(root *TreeNode) int {
    var dfs func(root *TreeNode) (int, int)
    dfs = func(root *TreeNode) (int, int) {
        if root == nil {
            return 0, 1
        }
        c1, l1 := dfs(root.Left)
        c2, l2 := dfs(root.Right)
        if l1 == 2 || l2 == 2 { // child and itself do not has camera, need camera on parent
            return c1 + c2 + 1, 0
        }
        if l1 == 0 || l2 == 0 { // with camera, not need camera, affect to parent
            return c1 + c2, 1
        }
        return c1 + c2, 2 // child has camera, not need camera, not affect to parent
    }
    c, l := dfs(root)
    if l == 2 {
        c++
    }
    return c
}

func minCameraCover1(root *TreeNode) int {
    cameras := 0
    var dfs func(root *TreeNode, cameras *int) int
    dfs = func(root *TreeNode, cameras *int) int {
        if root == nil {
            return 0
        }
        r := dfs(root.Left, cameras) | dfs(root.Right, cameras)
        *cameras += r & 1
        return int("1202"[r] - '0')
    }
    r := dfs(root, &cameras) & 1
    return cameras + r
}


func main() {
    // Example 1:
    //         (0)
    //         /
    //       (c)
    //      /   \   
    //    (0)   (0)
    // <img src="https://assets.leetcode.com/uploads/2018/12/29/bst_cameras_01.png" />
    // Input: root = [0,0,null,0,0]
    // Output: 1
    // Explanation: One camera is enough to monitor all nodes if placed as shown.
    tree1 := buildTree([]*int{ptr(0), ptr(0), nil, ptr(0), ptr(0)})
    printTree(tree1)
    fmt.Println(minCameraCover(tree1)) // 1
    // Example 2:
    //             (0)
    //             /
    //           (c)
    //           /
    //         (0)
    //         /
    //       (c)
    //         \ 
    //          (0)
    // <img src="https://assets.leetcode.com/uploads/2018/12/29/bst_cameras_02.png" />
    // Input: root = [0,0,null,0,null,0,null,null,0]
    // Output: 2
    // Explanation: At least two cameras are needed to monitor all nodes of the tree. The above image shows one of the valid configurations of camera placement.
    tree2 := buildTree([]*int{ptr(0), ptr(0), nil, ptr(0), nil, ptr(0), nil, nil, ptr(0)})
    printTree(tree2)
    fmt.Println(minCameraCover(tree2)) // 2

    fmt.Println(minCameraCover1(tree1)) // 1
    fmt.Println(minCameraCover1(tree2)) // 2
}