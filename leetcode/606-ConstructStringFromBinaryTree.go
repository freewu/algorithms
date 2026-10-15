package main

// 606. Construct String from Binary Tree
// Given the root node of a binary tree, 
// your task is to create a string representation of the tree following a specific set of formatting rules. 
// The representation should be based on a preorder traversal of the binary tree and must adhere to the following guidelines:
//     1. Node Representation: Each node in the tree should be represented by its integer value.
//     2. Parentheses for Children: If a node has at least one child (either left or right), 
//        its children should be represented inside parentheses. Specifically:
//         2.1 If a node has a left child, the value of the left child should be enclosed in parentheses immediately following the node's value.
//         2.2 If a node has a right child, the value of the right child should also be enclosed in parentheses. 
//             The parentheses for the right child should follow those of the left child.
//     3. Omitting Empty Parentheses: Any empty parentheses pairs (i.e., ()) should be omitted from the final string representation of the tree, with one specific exception: when a node has a right child but no left child. 
//        In such cases, you must include an empty pair of parentheses to indicate the absence of the left child. 
//        This ensures that the one-to-one mapping between the string representation and the original binary tree structure is maintained.

// In summary, empty parentheses pairs should be omitted when a node has only a left child or no children. 
// However, when a node has a right child but no left child, an empty pair of parentheses must precede the representation of the right child to reflect the tree's structure accurately.

// Example 1:
//         1
//       /   \
//      2     3
//     /
//    4
// <img src="https://assets.leetcode.com/uploads/2021/05/03/cons1-tree.jpg" />
// Input: root = [1,2,3,4]
// Output: "1(2(4))(3)"
// Explanation: Originally, it needs to be "1(2(4)())(3()())", but you need to omit all the empty parenthesis pairs. And it will be "1(2(4))(3)".

// Example 2:
//         1
//       /   \
//      2     3
//       \ 
//        4
// <img src="https://assets.leetcode.com/uploads/2021/05/03/cons2-tree.jpg" />
// Input: root = [1,2,3,null,4]
// Output: "1(2()(4))(3)"
// Explanation: Almost the same as the first example, except the () after 2 is necessary to indicate the absence of a left child for 2 and the presence of a right child.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -1000 <= Node.val <= 1000

import "fmt"
import "strconv"
import "strings"

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

// Recursion
func tree2str(t *TreeNode) string {
    var dfs func (node *TreeNode) string
    dfs = func (node *TreeNode) string {
        if node == nil {
            return ""
        }
        if node.Left != nil && node.Right != nil {
            return fmt.Sprintf("%d(%s)(%s)", node.Val, dfs(node.Left), dfs(node.Right))
        } else if node.Left == nil && node.Right == nil {
            return strconv.Itoa(node.Val)
        } else if node.Left != nil {
            return fmt.Sprintf("%d(%s)", node.Val, dfs(node.Left))
        } else {
            return fmt.Sprintf("%d()(%s)", node.Val, dfs(node.Right))
        }
    }
    return dfs(t)
}

func tree2str1(root *TreeNode) string {
    var dfs func(root *TreeNode) string
    dfs = func(root *TreeNode) string {
        switch {
        case root == nil:
            return ""
        case root.Left == nil && root.Right == nil:
            return strconv.Itoa(root.Val)
        case root.Right == nil:
            return fmt.Sprintf("%d(%s)", root.Val, dfs(root.Left))
        default:
            return fmt.Sprintf("%d(%s)(%s)", root.Val, dfs(root.Left), dfs(root.Right))
        }
    }
    return dfs(root)
}

func tree2str2(root *TreeNode) string {
    var dfs func(root *TreeNode) string
    dfs = func(root *TreeNode) string {
        r := &strings.Builder{}
        r.WriteString(strconv.Itoa(root.Val))
        if root.Left == nil && root.Right == nil {
            return r.String()
        }
        if root.Left != nil{
            r.WriteByte('(')
            r.WriteString(dfs(root.Left))
            r.WriteByte(')')
        } else {
            r.WriteString("()")
        }
        if root.Right != nil{
            r.WriteByte('(')
            r.WriteString(dfs(root.Right))
            r.WriteByte(')')
        }
        return r.String()
    }
    return dfs(root)
}

func main() {
    // Example 1:
    //         1
    //       /   \
    //      2     3
    //     /
    //    4
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/cons1-tree.jpg" />
    // Input: root = [1,2,3,4]
    // Output: "1(2(4))(3)"
    // Explanation: Originally, it needs to be "1(2(4)())(3()())", but you need to omit all the empty parenthesis pairs. And it will be "1(2(4))(3)".
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(5)})
    printTree(tree1)
    fmt.Println(tree2str(tree1)) // "1(2(4))(3)"
    // Example 2:
    //         1
    //       /   \
    //      2     3
    //       \ 
    //        4
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/cons2-tree.jpg" />
    // Input: root = [1,2,3,null,4]
    // Output: "1(2()(4))(3)"
    // Explanation: Almost the same as the first example, except the () after 2 is necessary to indicate the absence of a left child for 2 and the presence of a right child.
    tree2 := buildTree([]*int{ptr(1), ptr(2), ptr(3), nil, ptr(5)})
    printTree(tree2)
    fmt.Println(tree2str(tree2)) // "1(2()(4))(3)"

    fmt.Println(tree2str1(tree1)) // "1(2(4))(3)"
    fmt.Println(tree2str1(tree2)) // "1(2()(4))(3)"

    fmt.Println(tree2str2(tree1)) // "1(2(4))(3)"
    fmt.Println(tree2str2(tree2)) // "1(2()(4))(3)"
}