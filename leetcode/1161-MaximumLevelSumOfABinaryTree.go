package main

// 1161. Maximum Level Sum of a Binary Tree
// Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.
// Return the smallest level x such that the sum of all the values of nodes at level x is maximal.

// Example 1:
//         1
//       /   \
//      7     0
//    /  \
//   7    -8
// <img src="https://assets.leetcode.com/uploads/2019/05/03/capture.JPG" />
// Input: root = [1,7,0,7,-8,null,null]
// Output: 2
// Explanation: 
// Level 1 sum = 1.
// Level 2 sum = 7 + 0 = 7.
// Level 3 sum = 7 + -8 = -1.
// So we return the level with the maximum sum which is level 2.

// Example 2:
// Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
// Output: 2
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -10^5 <= Node.val <= 10^5

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
// bfs
func maxLevelSum(root *TreeNode) int {
    res, index, level, queue := -1 << 32 - 1, -1, 0, []*TreeNode{root}
    for len(queue) > 0 {
        sum, l := 0, len(queue) // 只取本层的 l 
        for i := 0; i < l; i++ {
            node := queue[0] // 出队列
            queue = queue[1:]
            sum += node.Val // 累加每层的值
            if node.Left != nil { // 如果存在左节点，加入到队列中
                queue = append(queue, node.Left)
            }   
            if node.Right != nil { // 如果存在右节点，加入到队列中
                queue = append(queue, node.Right)
            }
        }
        level++
        if sum > res { // 如果是大的则
            res, index = sum, level 
        }
    }
    return index
}

// dfs
func maxLevelSum1(root *TreeNode) int {
    levelSums := []int{}
    var dfs func(node *TreeNode, level int)
    dfs = func(node *TreeNode, level int) {
        if node == nil { return }
        if len(levelSums) == level {
            levelSums = append(levelSums, node.Val)
        } else {
            levelSums[level] += node.Val
        }
        dfs(node.Left, level + 1)
        dfs(node.Right, level + 1)
    }
    dfs(root, 0)
    maxSum, index := levelSums[0], 0
    for i := 1; i < len(levelSums); i++ {
        if levelSums[i] > maxSum {
            maxSum = levelSums[i]
            index = i 
        }
    }
    return index + 1
}

func main() {
    // Example 1:
    //         1
    //       /   \
    //      7     0
    //    /  \
    //   7    -8
    // <img src="https://assets.leetcode.com/uploads/2019/05/03/capture.JPG" />
    // Input: root = [1,7,0,7,-8,null,null]
    // Output: 2
    // Explanation: 
    // Level 1 sum = 1.
    // Level 2 sum = 7 + 0 = 7.
    // Level 3 sum = 7 + -8 = -1.
    // So we return the level with the maximum sum which is level 2.
    tree1 := &TreeNode{
        1, 
        &TreeNode{7, &TreeNode{7, nil, nil}, &TreeNode{-8, nil, nil}, },
        &TreeNode{0, nil, nil},
    }
    fmt.Println(maxLevelSum(tree1)) // 2
    // Example 2:
    // Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
    // Output: 2
    tree2 := &TreeNode{
        989, 
        nil,
        &TreeNode{10250, &TreeNode{98693, nil, nil}, &TreeNode{-89388, nil, &TreeNode{-32127, nil, nil}, }, },
    }
    fmt.Println(maxLevelSum(tree2)) // 2

    fmt.Println(maxLevelSum1(tree1)) // 2
    fmt.Println(maxLevelSum1(tree2)) // 2
}