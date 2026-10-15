package main

// 235. Lowest Common Ancestor of a Binary Search Tree
// Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.
// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png" />
// Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
// Output: 6
// Explanation: The LCA of nodes 2 and 8 is 6.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png" />
// Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
// Output: 2
// Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

// Example 3:
// Input: root = [2,1], p = 2, q = 1
// Output: 2
 
// Constraints:
//         The number of nodes in the tree is in the range [2, 10^5].
//         -10^9 <= Node.val <= 10^9
//         All Node.val are unique.
//         p != q
//         p and q will exist in the BST.

// 公共祖先
// 对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
// 满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。

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
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if p == nil || q == nil || root == nil {
        return nil
    }
    // 满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）
    if p.Val < root.Val && q.Val < root.Val {
        return lowestCommonAncestor(root.Left, p, q)
    }
    if p.Val > root.Val && q.Val > root.Val {
        return lowestCommonAncestor(root.Right, p, q)
    }
    return root
}

// 迭代
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    var curr *TreeNode = root
    for curr != nil {
        if p.Val < curr.Val && q.Val < curr.Val {
            curr = curr.Left
            continue
        }
        if p.Val > curr.Val && q.Val > curr.Val {
            curr = curr.Right
            continue
        }
        return curr
    }
    return curr
}

// best solution
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	for root != nil {
        if root.Val < p.Val && root.Val < q.Val {
            root = root.Right
        } else if root.Val > p.Val && root.Val > q.Val {
            root = root.Left
        } else {
            return root
        }
    }
    return nil
}

func main() {
    tree1 := buildTree([]*int{ptr(6), ptr(2), ptr(8), ptr(0), ptr(4), ptr(7), ptr(9), nil, nil, ptr(3), ptr(5)})
    printTree(tree1)
    tree3 := buildTree([]*int{ptr(2), ptr(1)})
    printTree(tree3)
    fmt.Println(lowestCommonAncestor(tree1,&TreeNode{2, nil, nil},&TreeNode{8, nil, nil})) // 6
    fmt.Println(lowestCommonAncestor(tree1,&TreeNode{2, nil, nil},&TreeNode{4, nil, nil})) // 2
    fmt.Println(lowestCommonAncestor(tree3,&TreeNode{2, nil, nil},&TreeNode{1, nil, nil})) // 2

    fmt.Println(lowestCommonAncestor1(tree1,&TreeNode{2, nil, nil},&TreeNode{8, nil, nil})) // 6
    fmt.Println(lowestCommonAncestor1(tree1,&TreeNode{2, nil, nil},&TreeNode{4, nil, nil})) // 2
    fmt.Println(lowestCommonAncestor1(tree3,&TreeNode{2, nil, nil},&TreeNode{1, nil, nil})) // 2

    fmt.Println(lowestCommonAncestor2(tree1,&TreeNode{2, nil, nil},&TreeNode{8, nil, nil})) // 6
    fmt.Println(lowestCommonAncestor2(tree1,&TreeNode{2, nil, nil},&TreeNode{4, nil, nil})) // 2
    fmt.Println(lowestCommonAncestor2(tree3,&TreeNode{2, nil, nil},&TreeNode{1, nil, nil})) // 2
}