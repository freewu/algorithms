package main

// 669. Trim a Binary Search Tree
// Given the root of a binary search tree and the lowest and highest boundaries as low and high, 
// trim the tree so that all its elements lies in [low, high]. 
// Trimming the tree should not change the relative structure of the elements that will remain in the tree (i.e., any node's descendant should remain a descendant). It can be proven that there is a unique answer.

// Return the root of the trimmed binary search tree. Note that the root may change depending on the given bounds.

// Example 1:
//         1                     1
//       /    \      =>           \ 
//      0      2                    2
// <img src="https://assets.leetcode.com/uploads/2020/09/09/trim1.jpg" /> 
// Input: root = [1,0,2], low = 1, high = 2
// Output: [1,null,2]

// Example 2:
//             3                 3
//           /   \              /
//          0     4            2
//           \         =>     /
//            2              1
//           /
//          1
// <img src="https://assets.leetcode.com/uploads/2020/09/09/trim2.jpg" /> 
// Input: root = [3,0,4,null,2,null,null,1], low = 1, high = 3
// Output: [3,2,null,1]
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     0 <= Node.val <= 10^4
//     The value of each node in the tree is unique.
//     root is guaranteed to be a valid binary search tree.
//     0 <= low <= high <= 10^4

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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
    if root == nil {
        return root
    }
    if root.Val >= low && root.Val <= high {
        root.Left  = trimBST(root.Left, low, high) // Trim the left subtree
        root.Right = trimBST(root.Right, low, high)  // Trim the right subtree
    } else if root.Val < low {
        // if the root val is less than low then getting values from left 
        // will be even lower(binary tree rule) so we need to set the root to root.right;
        root = trimBST(root.Right, low, high) 
    } else if root.Val > high {
        // if the root val is greater than high then getting values from right 
        // will be even higher(binary tree rule) so we need to the root to root.left.
        root = trimBST(root.Left, low, high) 
    }
    return root
}

func trimBST1(root *TreeNode, low int, high int) *TreeNode {
    var dfs func(root *TreeNode, low int, high int) *TreeNode
    dfs = func(root *TreeNode, low int, high int) *TreeNode {
        if root == nil {
            return nil
          }
          // Trim the left and right subtrees
          root.Left, root.Right = trimBST1(root.Left, low, high), trimBST1(root.Right, low, high)
          // If the current node's value is less than low, return the right subtree
          if root.Val < low {
            return root.Right
          }
          // If the current node's value is greater than high, return the left subtree
          if root.Val > high {
            return root.Left
          }
          // Otherwise, the current node is within the range
          return root
    }
    return dfs(root,low, high)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/09/trim1.jpg" /> 
    // Input: root = [1,0,2], low = 1, high = 2
    // Output: [1,null,2]
    tree1 := &TreeNode {
        1,
        &TreeNode{0, nil, nil, },
        &TreeNode{2, nil, nil, },
    }
    fmt.Println(trimBST(tree1,1,2))
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/09/09/trim2.jpg" /> 
    // Input: root = [3,0,4,null,2,null,null,1], low = 1, high = 3
    // Output: [3,2,null,1]
    tree2 := &TreeNode {
        3,
        &TreeNode{0, nil, &TreeNode{2, &TreeNode{1, nil, nil, }, nil, }, },
        &TreeNode{4, nil, nil, },
    }
    fmt.Println(trimBST(tree2,1,3))

    fmt.Println(trimBST1(tree1,1,2))
    fmt.Println(trimBST1(tree2,1,3))
}