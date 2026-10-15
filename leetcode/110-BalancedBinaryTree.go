package main

// 110. Balanced Binary Tree
// Given a binary tree, determine if it is height-balanced.

// Example 1:
//         3
//        /  \
//       9   20
//          /  \
//         15   7
// <img src="https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg" />
// Input: root = [3,9,20,null,null,15,7]
// Output: true

// Example 2:
//             1 
//           /   \
//          2     2
//        /   \
//       3     3
//     /   \
//    4     4
// <img src="https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg" />
// Input: root = [1,2,2,3,3,null,null,4,4]
// Output: false

// Example 3:
// <img src="" />
// Input: root = []
// Output: true
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 5000].
//     -10^4 <= Node.val <= 10^4

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
// 递归
func isBalanced(root *TreeNode) bool {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // iterative function to measure the maximum depth of a tree
    var maxDepth func(root *TreeNode) int 
    maxDepth = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        return max(maxDepth(root.Left), maxDepth(root.Right))+1
    }
    var helper func(root *TreeNode) bool
    helper = func(root *TreeNode) bool {
        if root == nil {
            return true
        }
        // getting maximum depth possible from both left and right children
        // check if the difference is not more than 1, if it is then the tree is unbalanced and we can stop our check
        if abs(maxDepth(root.Left) -  maxDepth(root.Right)) > 1 {
            return false
        }
        // if both children are balanced then this tree is also balanced
        return helper(root.Left) && helper(root.Right)
    }
    return helper(root)
}

func isBalanced1(root *TreeNode) bool {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return 0  }
        left := dfs(node.Left)
        if left == -1 { return -1 }
        right := dfs(node.Right)
        if right == -1 || abs(left - right) > 1 { return -1 }
        return max(left, right) + 1
    }
    return dfs(root) != -1
}

func main() {
    // Example 1:
    //         3
    //        /  \
    //       9   20
    //          /  \
    //         15   7
    // <img src="https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg" />
    // Input: root = [3,9,20,null,null,15,7]
    // Output: true
    tree1 := &TreeNode {
        3,
        &TreeNode{9, nil, nil},
        &TreeNode {
            20,
            &TreeNode{15, nil, nil},
            &TreeNode{7, nil, nil},
        },
    }
    fmt.Println(isBalanced(tree1)) // true
    // Example 2:
    //             1 
    //           /   \
    //          2     2
    //        /   \
    //       3     3
    //     /   \
    //    4     4
    // <img src="https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg" />
    // Input: root = [1,2,2,3,3,null,null,4,4]
    // Output: false
    tree2 := &TreeNode {
        1,
        &TreeNode {
            2,
            &TreeNode{
                3, 
                &TreeNode{4, nil, nil }, 
                &TreeNode{4, nil, nil }, 
            },
            &TreeNode{3, nil, nil },
        },
        &TreeNode{2, nil, nil},
    }
    fmt.Println(isBalanced(tree2)) // false
    // Example 3:
    // <img src="" />
    // Input: root = []
    // Output: true
    fmt.Println(isBalanced(nil)) // false

    fmt.Println(isBalanced1(tree1)) // true
    fmt.Println(isBalanced1(tree2)) // false
    fmt.Println(isBalanced1(nil)) // false
}