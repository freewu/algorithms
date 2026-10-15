package main

// LCR 153. 二叉树中和为目标值的路径
// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
// 叶子节点 是指没有子节点的节点。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsumii1.jpg" />
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：[[5,4,11,2],[5,8,4,5]]

// 示例 2：
// <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg" />
// 输入：root = [1,2,3], targetSum = 5
// 输出：[]

// 示例 3：
// 输入：root = [1,2], targetSum = 0
// 输出：[]

// 提示：
//     树中节点总数在范围 [0, 5000] 内
//     -1000 <= Node.val <= 1000
//     -1000 <= targetSum <= 1000

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
// dfs
func pathTarget(root *TreeNode, target int) [][]int {
    var dfs func(node *TreeNode, res []int, targetSum, currentSum int) [][]int 
    dfs = func(node *TreeNode, res []int, targetSum, currentSum int) [][]int {
        if node == nil {
            return [][]int{}
        }
        res = append(res, node.Val)
        currentSum += node.Val
        if node.Left == nil && node.Right == nil && currentSum == targetSum {
            return [][]int{res}
        }
        // Without copying it to another recursion It would give a pointer of the list
        newGroupOfNodes := make([]int, len(res))
        copy(newGroupOfNodes, res)
        return append(dfs(node.Left, newGroupOfNodes, targetSum, currentSum), dfs(node.Right, newGroupOfNodes, targetSum, currentSum)...)
    }
    return dfs(root, []int{}, target, 0)
}

// bfs
func pathTarget1(root *TreeNode, target int) [][]int {
    if root == nil {
        return nil
    }
    res := [][]int{}
    nodes, sums, paths := []*TreeNode{ root }, []int{ target - root.Val }, [][]int{{ root.Val }}
    for len(nodes) > 0 {
        curr, sum, path := nodes[len(nodes)-1], sums[len(sums)-1], paths[len(paths)-1]
        nodes, sums, paths = nodes[:len(nodes)-1], sums[:len(sums)-1], paths[:len(paths)-1]
        if curr.Left == nil && curr.Right == nil && sum == 0 { // 满足要求的路径
            res = append(res, append([]int{}, path...))
        }
        if curr.Right != nil {
            nodes = append(nodes, curr.Right)
            sums = append(sums, sum-curr.Right.Val)
            paths = append(paths, append(append([]int{}, path...), curr.Right.Val))
        }
        if curr.Left != nil {
            nodes = append(nodes, curr.Left)
            sums = append(sums, sum-curr.Left.Val)
            paths = append(paths, append(append([]int{}, path...), curr.Left.Val))
        }
    }
    return res
}

func main() {
    // Example 1:
    //            （5）
    //            /   \
    //          (4)   (8)
    //          /     /  \
    //        (11)   13   (4)
    //        /  \        /  \
    //       7   (2)     (5)  1
    // <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsumii1.jpg" />
    // Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
    // Output: [[5,4,11,2],[5,8,4,5]]
    // Explanation: There are two paths whose sum equals targetSum:
    // 5 + 4 + 11 + 2 = 22
    // 5 + 8 + 4 + 5 = 22
    tree1 := &TreeNode{
        5, 
        &TreeNode{
            4, 
            &TreeNode{11, &TreeNode{7, nil, nil }, &TreeNode{2, nil, nil }, }, 
            nil,
        }, 
        &TreeNode{
            8, 
            &TreeNode{13, nil, nil }, 
            &TreeNode{4, &TreeNode{5, nil, nil }, &TreeNode{1, nil, nil } },
        }, 
    }
    fmt.Println(pathTarget(tree1, 22)) // [[5,4,11,2],[5,8,4,5]]
    // Example 2:
    //      1
    //    /   \
    //   2     3
    // <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg" />
    // Input: root = [1,2,3], targetSum = 5
    // Output: []
    tree2 := &TreeNode{
        1, 
        &TreeNode{2, nil, nil }, 
        &TreeNode{3, nil, nil }, 
    }
    fmt.Println(pathTarget(tree2, 5)) // []
    // Example 3:
    //      1
    //    /  
    //   2     
    // Input: root = [1,2], targetSum = 0
    // Output: []
    tree3 := &TreeNode{
        1, 
        &TreeNode{2, nil, nil }, 
        nil, 
    }
    fmt.Println(pathTarget(tree3,0)) // []

    fmt.Println(pathTarget1(tree1, 22)) // [[5,4,11,2],[5,8,4,5]]
    fmt.Println(pathTarget1(tree2, 5)) // []
    fmt.Println(pathTarget1(tree3, 0)) // []
}