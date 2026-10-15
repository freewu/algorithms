package main

// 572. Subtree of Another Tree
// Given the roots of two binary trees root and subRoot, 
// return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

// A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. 
// The tree tree could also be considered as a subtree of itself.

// Example 1:
//         root                  subRoot
//           3
//         /   \                  4
//       [4]    5                /  \
//      /  \                    1    2
//    [1]  [2]
// <img src="https://assets.leetcode.com/uploads/2021/04/28/subtree1-tree.jpg" />
// Input: root = [3,4,5,1,2], subRoot = [4,1,2]
// Output: true

// Example 2:
//        root               subRoot
//          3
//        /   \                 4
//       4     5               /  \
//      /  \                  1    2
//     1    2
//          /
//         0
// <img src="https://assets.leetcode.com/uploads/2021/04/28/subtree2-tree.jpg" />
// Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
// Output: false
 
// Constraints:
//     The number of nodes in the root tree is in the range [1, 2000].
//     The number of nodes in the subRoot tree is in the range [1, 1000].
//     -10^4 <= root.val <= 10^4
//     -10^4 <= subRoot.val <= 10^4

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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return subRoot == nil
    }
    var isSameTree func(p *TreeNode, q *TreeNode) bool
    isSameTree = func(p *TreeNode, q *TreeNode) bool {
        if p == nil && q == nil {
            return true
        }
        if p == nil || q == nil || p.Val != q.Val {
            return false
        }
        return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
    }
    if isSameTree(root, subRoot) {
        return true
    }
    return isSubtree(root.Right, subRoot) || isSubtree(root.Left, subRoot)
}

func isSubtree1(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }
    var same func (p, q *TreeNode) bool
    same = func (p, q *TreeNode) bool {
        if p == nil && q == nil {
            return true
        } else if p == nil || q == nil || p.Val != q.Val {
            return false
        }
        return same(p.Left, q.Left) && same(p.Right, q.Right)
    }
    return same(root, subRoot) || isSubtree1(root.Left, subRoot) || isSubtree1(root.Right, subRoot)
}

func main() {
    // Example 1:
    //         root                  subRoot
    //           3
    //         /   \                  4
    //       [4]    5                /  \
    //      /  \                    1    2
    //    [1]  [2]
    // <img src="https://assets.leetcode.com/uploads/2021/04/28/subtree1-tree.jpg" />
    // Input: root = [3,4,5,1,2], subRoot = [4,1,2]
    // Output: true
    root1 := buildTree([]*int{ptr(3), ptr(4), ptr(5), ptr(1), ptr(2)})
    printTree(root1)
    sub1 := buildTree([]*int{ptr(4), ptr(1), ptr(2)})
    printTree(sub1)
    fmt.Println(isSubtree(root1,sub1)) // true
    // Example 2:
    //        root               subRoot
    //          3
    //        /   \                 4
    //       4     5               /  \
    //      /  \                  1    2
    //     1    2
    //          /
    //         0
    // <img src="https://assets.leetcode.com/uploads/2021/04/28/subtree2-tree.jpg" />
    // Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
    // Output: false
    root2 := buildTree([]*int{ptr(3), ptr(4), ptr(5), ptr(1), ptr(2), nil, nil, nil, nil, ptr(0)})
    printTree(root2)
    sub2 := buildTree([]*int{ptr(4), ptr(1), ptr(2)})
    printTree(sub2)
    fmt.Println(isSubtree(root2,sub2)) // false

    fmt.Println(isSubtree1(root1,sub1)) // true
    fmt.Println(isSubtree1(root2,sub2)) // false
}