package main

// 2265. Count Nodes Equal to Average of Subtree
// Given the root of a binary tree, 
// return the number of nodes where the value of the node is equal to the average of the values in its subtree.

// Note:
//     The average of n elements is the sum of the n elements divided by n and rounded down to the nearest integer.
//     A subtree of root is a tree consisting of root and all of its descendants.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/15/image-20220315203925-1.png" />
// Input: root = [4,8,5,0,1,null,6]
// Output: 5
// Explanation: 
// For the node with value 4: The average of its subtree is (4 + 8 + 5 + 0 + 1 + 6) / 6 = 24 / 6 = 4.
// For the node with value 5: The average of its subtree is (5 + 6) / 2 = 11 / 2 = 5.
// For the node with value 0: The average of its subtree is 0 / 1 = 0.
// For the node with value 1: The average of its subtree is 1 / 1 = 1.
// For the node with value 6: The average of its subtree is 6 / 1 = 6.

// Example 2:
// Input: root = [1]
// Output: 1
// Explanation: For the node with value 1: The average of its subtree is 1 / 1 = 1.

// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
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
func averageOfSubtree(root *TreeNode) int {
    res := 0
    var dfs func(node *TreeNode) (int, int)
    dfs = func(node *TreeNode) (int, int) {
        count, sum := 1, node.Val
        if nil != node.Left {
            lcount, lsum := dfs(node.Left)
            count += lcount
            sum += lsum
        }
        if nil != node.Right {
            rcount, rsum := dfs(node.Right)
            count += rcount
            sum += rsum
        }
        if node.Val == sum / count {
            res++
        }
        return count, sum
    }
    dfs(root)
    return res
}

func averageOfSubtree1(root *TreeNode) int {
    res := 0
    var dfs func(node *TreeNode) (int, int)
    dfs = func(node *TreeNode) (int, int) {
        if node == nil { return 0, 0 }
        lcount, lsum := dfs(node.Left)
        rcount, rsum := dfs(node.Right)
        count, sum := lcount + rcount + 1, lsum + rsum + node.Val
        if sum / count == node.Val {
            res++
        }
        return count, sum
    }
    dfs(root)
    return res
}

func averageOfSubtree2(root *TreeNode) int {
    res := 0
    var dfs func(node *TreeNode) (int, int)
    dfs = func(node *TreeNode) (int, int) {
        if node.Right == nil && node.Left == nil {
            res++
            return node.Val, 1
        }
        leftSum, rightSum, leftChildren, rightChildren := 0, 0, 0, 0
        if node.Left != nil {
            leftSum, leftChildren = dfs(node.Left)
        }
        if node.Right != nil {
            rightSum, rightChildren = dfs(node.Right)
        }
        if (leftSum+rightSum+node.Val)/(leftChildren+rightChildren+1) == node.Val {
            res++
        }
        return leftSum + rightSum + node.Val, leftChildren + rightChildren + 1
    }
    dfs(root)
    return res  
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/15/image-20220315203925-1.png" />
    // Input: root = [4,8,5,0,1,null,6]
    // Output: 5
    // Explanation: 
    // For the node with value 4: The average of its subtree is (4 + 8 + 5 + 0 + 1 + 6) / 6 = 24 / 6 = 4.
    // For the node with value 5: The average of its subtree is (5 + 6) / 2 = 11 / 2 = 5.
    // For the node with value 0: The average of its subtree is 0 / 1 = 0.
    // For the node with value 1: The average of its subtree is 1 / 1 = 1.
    // For the node with value 6: The average of its subtree is 6 / 1 = 6.
    tree1 := buildTree([]*int{ptr(4), ptr(8), ptr(5), ptr(0), ptr(1), nil, ptr(6)})
    printTree(tree1)
    fmt.Println(averageOfSubtree(tree1)) // 5
    // Example 2:
    // Input: root = [1]
    // Output: 1
    // Explanation: For the node with value 1: The average of its subtree is 1 / 1 = 1.
    tree2 := buildTree([]*int{ptr(1)})
    printTree(tree2)
    fmt.Println(averageOfSubtree(tree2)) // 1

    fmt.Println(averageOfSubtree1(tree1)) // 5
    fmt.Println(averageOfSubtree1(tree2)) // 1

    fmt.Println(averageOfSubtree2(tree1)) // 5
    fmt.Println(averageOfSubtree2(tree2)) // 1
}