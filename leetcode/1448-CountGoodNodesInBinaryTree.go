package main

// 1448. Count Good Nodes in Binary Tree
// Given a binary tree root, a node X in the tree is named good 
// if in the path from root to X there are no nodes with a value greater than X.
// Return the number of good nodes in the binary tree.

// Example 1:
//         (3)
//        /   \
//       1    (4)
//      /     /  \
//    (3)    1    (5)
// <img src="https://assets.leetcode.com/uploads/2020/04/02/test_sample_1.png" />
// Input: root = [3,1,4,3,null,1,5]
// Output: 4
// Explanation: Nodes in blue are good.
// Root Node (3) is always a good node.
// Node 4 -> (3,4) is the maximum value in the path starting from the root.
// Node 5 -> (3,4,5) is the maximum value in the path
// Node 3 -> (3,1,3) is the maximum value in the path.

// Example 2:
//        (3)
//        /
//       (3)
//      /  \
//    (4)   2
// <img src="https://assets.leetcode.com/uploads/2020/04/02/test_sample_2.png" />
// Input: root = [3,3,null,4,2]
// Output: 3
// Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.

// Example 3:
// Input: root = [1]
// Output: 1
// Explanation: Root is considered as good.

// Constraints:
//     The number of nodes in the binary tree is in the range [1, 10^5].
//     Each node's value is between [-10^4, 10^4].

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
func goodNodes(root *TreeNode) int {
    count := 0 // Initialize a variable to count good nodes.
    var dfs func(node *TreeNode, max int) // Define a helper function for the recursive traversal.
    dfs = func(node *TreeNode, max int) { // The dfs function performs a depth-first search traversal of the binary tree.
        if node == nil { return; } // Base case: If the node is nil, return without processing.
        if node.Val >= max {
            max = node.Val   // Update the maximum value encountered so far.
            count++          // Increment the count to indicate a good node.
        }
        // Recursively traverse the left and right subtrees with the updated max value.
        dfs(node.Left, max)
        dfs(node.Right, max)
    }
    dfs(root, root.Val) // Start the traversal with the root node and its value.
    return count // Return the total count of good nodes in the binary tree.
}

func goodNodes1(root *TreeNode) int {
    value := root.Val
    var dfs func(node *TreeNode, maxValue int) int
    dfs = func(node *TreeNode, maxValue int) int {
        if node == nil { return 0; }
        if node.Val >= maxValue {
            return 1 + dfs(node.Left, node.Val) + dfs(node.Right, node.Val)
        } else {
            return dfs(node.Left, maxValue) + dfs(node.Right, maxValue)
        }
    }
    return 1 + dfs(root.Left, value) + dfs(root.Right, value)
}

func main() {
    // Example 1:
    //         (3)
    //        /   \
    //       1    (4)
    //      /     /  \
    //    (3)    1    (5)
    // <img src="https://assets.leetcode.com/uploads/2020/04/02/test_sample_1.png" />
    // Input: root = [3,1,4,3,null,1,5]
    // Output: 4
    // Explanation: Nodes in blue are good.
    // Root Node (3) is always a good node.
    // Node 4 -> (3,4) is the maximum value in the path starting from the root.
    // Node 5 -> (3,4,5) is the maximum value in the path
    // Node 3 -> (3,1,3) is the maximum value in the path.
    tree1 := buildTree([]*int{ptr(3), ptr(1), ptr(4), ptr(3), nil, ptr(1), ptr(5)})
    printTree(tree1)
    fmt.Println(goodNodes(tree1)) // 4
    // Example 2:
    //         (3)
    //         /
    //       (3)
    //      /  \
    //    (4)   2
    // <img src="https://assets.leetcode.com/uploads/2020/04/02/test_sample_2.png" />
    // Input: root = [3,3,null,4,2]
    // Output: 3
    // Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.
    tree2 := buildTree([]*int{ptr(3), ptr(3), nil, ptr(4), ptr(2)})
    printTree(tree2)
    fmt.Println(goodNodes(tree2)) // 3
    // Example 3:
    // Input: root = [1]
    // Output: 1
    // Explanation: Root is considered as good.
    tree3 := buildTree([]*int{ptr(1)})
    printTree(tree3)
    fmt.Println(goodNodes(tree3)) // 1

    fmt.Println(goodNodes1(tree1)) // 4
    fmt.Println(goodNodes1(tree2)) // 3
    fmt.Println(goodNodes1(tree3)) // 1
}