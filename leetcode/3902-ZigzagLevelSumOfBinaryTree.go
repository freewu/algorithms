package main

// 3902. Zigzag Level Sum of Binary Tree
// You are given the root of a binary tree.

// Traverse the tree level by level using a zigzag pattern:
//     1. At odd-numbered levels (1-indexed), traverse nodes from left to right.
//     2. At even-numbered levels, traverse nodes from right to left.

// While traversing a level in the specified direction, process nodes in order and stop immediately before the first node that violates the condition:
//     1. At odd levels: the node does not have a left child.
//     2. At even levels: the node does not have a right child.

// Only the nodes processed before this stopping condition contribute to the level sum.

// Return an integer array ans where ans[i] is the sum of the node values that are processed at level i + 1.

// Example 1:
// Input: root = [5,2,8,1,null,9,6]
// Output: [5,8,0]
// Explanation:
//                 5 
//               /    \
//              2      8
//             /      /  \
//            1      9    6
// ​​​​​​<img src="https://assets.leetcode.com/uploads/2026/04/12/screenshot-2026-04-13-at-22054am.png" />
// At level 1, nodes are processed left to right. Node 5 is included, thus ans[0] = 5.
// At level 2, nodes are processed right to left. Node 8 is included, but node 2 lacks a right child, so processing stops, thus ans[1] = 8.
// At level 3, nodes are processed left to right. The first node 1 lacks a left child, so no nodes are included, and ans[2] = 0.
// Thus, ans = [5, 8, 0].

// Example 2:
// Input: root = [1,2,3,4,5,null,7]
// Output: [1,5,0]
// Explanation:
//                 1 
//               /    \
//              2      3
//             /  \      \
//            4     5     7
// <img src="https://assets.leetcode.com/uploads/2026/04/12/screenshot-2026-04-13-at-22232am.png" />
// At level 1, nodes are processed left to right. Node 1 is included, thus ans[0] = 1.
// At level 2, nodes are processed right to left. Nodes 3 and 2 are included since both have right children, thus ans[1] = 3 + 2 = 5.
// At level 3, nodes are processed left to right. The first node 4 lacks a left child, so no nodes are included, and ans[2] = 0.
// Thus, ans = [1, 5, 0].

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     -10^5 <= Node.val <= 10^5

import "fmt"

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
func zigzagLevelSum(root *TreeNode) []int64 {
    res := []int64{}
    if root == nil { return res } // 处理根节点为空的情况
    queue := []*TreeNode{root}
    level := 1
    for len(queue) > 0 {
        var childNodes []*TreeNode
        violated := false
        // 追加当前层的和，初始为0
        res = append(res, 0)
        // 逆序遍历 curNodes[::-1]
        for i := len(queue) - 1; i >= 0; i-- {
            node := queue[i]
            if level & 1 != 0 {  // level & 1 等价于 奇数层
                // 交换左右节点
                node.Left, node.Right = node.Right, node.Left
            }
            if node.Right == nil {
                violated = true
            }
            if !violated { // 未违反则累加值
                res[len(res)-1] += int64(node.Val)
            }
            if node.Right != nil {
                childNodes = append(childNodes, node.Right)
            }
            if node.Left != nil {
                childNodes = append(childNodes, node.Left)
            }
        }
        level++
        queue = childNodes
    }
    return res
}

func main() {
    // Example 1:
    // Input: root = [5,2,8,1,null,9,6]
    // Output: [5,8,0]
    // Explanation:
    //                 5 
    //               /    \
    //              2      8
    //             /      /  \
    //            1      9    6
    // ​​​​​​<img src="https://assets.leetcode.com/uploads/2026/04/12/screenshot-2026-04-13-at-22054am.png" />
    // At level 1, nodes are processed left to right. Node 5 is included, thus ans[0] = 5.
    // At level 2, nodes are processed right to left. Node 8 is included, but node 2 lacks a right child, so processing stops, thus ans[1] = 8.
    // At level 3, nodes are processed left to right. The first node 1 lacks a left child, so no nodes are included, and ans[2] = 0.
    // Thus, ans = [5, 8, 0].
    tree1 := &TreeNode{
        Val: 5, 
        Left:  &TreeNode{ 
            Val: 2,
            Left: &TreeNode{ Val: 1},
        }, 
        Right: &TreeNode{ 
            Val: 8,
            Left: &TreeNode{ Val: 9},
            Right: &TreeNode{ Val: 6},
        },
    }
    fmt.Println(zigzagLevelSum(tree1)) // [5,8,0]
    // Example 2:
    // Input: root = [1,2,3,4,5,null,7]
    // Output: [1,5,0]
    // Explanation:
    //                 1 
    //               /    \
    //              2      3
    //             /  \      \
    //            4     5     7
    // <img src="https://assets.leetcode.com/uploads/2026/04/12/screenshot-2026-04-13-at-22232am.png" />
    // At level 1, nodes are processed left to right. Node 1 is included, thus ans[0] = 1.
    // At level 2, nodes are processed right to left. Nodes 3 and 2 are included since both have right children, thus ans[1] = 3 + 2 = 5.
    // At level 3, nodes are processed left to right. The first node 4 lacks a left child, so no nodes are included, and ans[2] = 0.
    // Thus, ans = [1, 5, 0].
    tree2 := &TreeNode{
        Val: 1, 
        Left:  &TreeNode{ 
            Val: 2,
            Left: &TreeNode{ Val: 4},
            Right: &TreeNode{ Val: 5},
        }, 
        Right: &TreeNode{ 
            Val: 3,
            Right: &TreeNode{ Val: 7},
        },
    }
    fmt.Println(zigzagLevelSum(tree2)) // [1,5,0]
}

