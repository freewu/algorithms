package main

// 897. Increasing Order Search Tree
// Given the root of a binary search tree, 
// rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree, 
// and every node has no left child and only one right child.

// Example 1:
//             5                           1
//           /    \                          \
//          3      6                          2 
//        /   \      \             =>          \
//       2     4      8                          3
//     /             /  \                         \
//    1             7    9                         4
//                                                 [\]
//                                                   8
//                                                     \
//                                                      9
// <img src="https://assets.leetcode.com/uploads/2020/11/17/ex1.jpg" />
// Input: root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
// Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/17/ex2.jpg" />
//         5                   1
//       /   \        =>         \
//      1     7                    5
//                                  \
//                                    7
// Input: root = [5,1,7] 
// Output: [1,null,5,null,7]

// Constraints:
//     The number of nodes in the given tree will be in the range [1, 100].
//     0 <= Node.val <= 1000

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
func increasingBST1(root *TreeNode) *TreeNode {
    nodes := make([]int, 0)
    var inorder func (root *TreeNode)
    inorder = func (root *TreeNode) { // 中序遍历
        if root != nil {
            inorder(root.Left)
            nodes = append(nodes, root.Val)
            inorder(root.Right)
        }
    }
    inorder(root)
    res := new(TreeNode)
    dummy := res
    for _, val := range nodes { // 重建树
        dummy.Right = &TreeNode{ Val: val, Left: nil, Right: nil }
        dummy = dummy.Right
    }
    return res.Right
}


// 递归
func increasingBST(root *TreeNode) *TreeNode {
    if root == nil { return nil }
    left := increasingBST(root.Left) // 先把左右子树拉平
    root.Left = nil // 然后把根节点左子树置空
    right := increasingBST(root.Right) // 再把右子树拉平
    root.Right = right // 连上根节点
    if left == nil { // 左子树为空的话，就不用处理了
        return root
    }
    p := left
    for p != nil && p.Right != nil { // 左子树非空，把根节点和右子树接到左子树末尾
        p = p.Right
    }
    p.Right = root
    return left
}

func main() {
    // Example 1:
    //             5                           1
    //           /    \                          \
    //          3      6                          2 
    //        /   \      \             =>          \
    //       2     4      8                          3
    //     /             /  \                         \
    //    1             7    9                         4
    //                                                 [\]
    //                                                   8
    //                                                     \
    //                                                      9
    // <img src="https://assets.leetcode.com/uploads/2020/11/17/ex1.jpg" />
    // Input: root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
    // Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
    tree1 := buildTree([]*int{ptr(5), ptr(3), ptr(6), ptr(2), ptr(4), nil, ptr(8), ptr(1), nil, nil, nil, ptr(7), ptr(9)})
    printTree(tree1)
    fmt.Println(increasingBST(tree1)) // &{1 <nil> 0xc000094120}
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/11/17/ex2.jpg" />
    //         5                   1
    //       /   \        =>         \
    //      1     7                    5
    //                                  \
    //                                    7
    // Input: root = [5,1,7] 
    // Output: [1,null,5,null,7]
    tree2 := buildTree([]*int{ptr(5), ptr(1), ptr(7)})
    printTree(tree2)
    fmt.Println(increasingBST(tree2)) // &{1 <nil> 0xc000094258}

    tree11 := buildTree([]*int{ptr(5), ptr(3), ptr(6), ptr(2), ptr(4), nil, ptr(8), ptr(1), nil, nil, nil, ptr(7), ptr(9)})
    printTree(tree11)
    fmt.Println(increasingBST1(tree11)) // &{1 <nil> 0xc000094120}
    tree12 := buildTree([]*int{ptr(5), ptr(1), ptr(7)})
    printTree(tree12)
    fmt.Println(increasingBST1(tree12)) // &{1 <nil> 0xc000094258}
}