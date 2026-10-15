package main

// 654. Maximum Binary Tree
// You are given an integer array nums with no duplicates. 
// A maximum binary tree can be built recursively from nums using the following algorithm:
//     Create a root node whose value is the maximum value in nums.
//     Recursively build the left subtree on the subarray prefix to the left of the maximum value.
//     Recursively build the right subtree on the subarray suffix to the right of the maximum value.

// Return the maximum binary tree built from nums.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/24/tree1.jpg" />
// Input: nums = [3,2,1,6,0,5]
// Output: [6,3,5,null,2,0,null,null,1]
// Explanation: The recursive calls are as follow:
// - The largest value in [3,2,1,6,0,5] is 6. Left prefix is [3,2,1] and right suffix is [0,5].
//     - The largest value in [3,2,1] is 3. Left prefix is [] and right suffix is [2,1].
//         - Empty array, so no child.
//         - The largest value in [2,1] is 2. Left prefix is [] and right suffix is [1].
//             - Empty array, so no child.
//             - Only one element, so child is a node with value 1.
//     - The largest value in [0,5] is 5. Left prefix is [0] and right suffix is [].
//         - Only one element, so child is a node with value 0.
//         - Empty array, so no child.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/24/tree2.jpg" />
// Input: nums = [3,2,1]
// Output: [3,null,2,null,1]

// Constraints:
//     1 <= nums.length <= 1000
//     0 <= nums[i] <= 1000
//     All integers in nums are unique.

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

func constructMaximumBinaryTree(nums []int) *TreeNode {
    // 找到左边最大的数作为左子树的根节点
    // 找到最大的数作为根节点
    // 以此类推
    if len(nums) == 0 {
        return nil
    }
    var root TreeNode
    mx, index := nums[0], 0
    for i, v := range nums { // 找到最大的数作为根节点
        if mx < v {
            mx = v
            index = i
        }
    }
    root.Val = mx
    root.Left = constructMaximumBinaryTree(nums[:index]) // 找到左边最大的数作为左子树的根节点
    root.Right = constructMaximumBinaryTree(nums[index + 1:])
    return &root
}

func constructMaximumBinaryTree1(nums []int) *TreeNode {
    var build func(left, right int) *TreeNode
    build = func(left, right int) *TreeNode {
        if right < left {
            return nil
        }
        mx, index := -1, 0
        for i := left; i <= right; i++ {
            if nums[i] > mx {
                mx = nums[i]
                index = i
            }
        }
        root := &TreeNode{ Val: nums[index], }
        root.Left = build(left, index - 1)
        root.Right = build(index + 1, right)
        return root
    }
    return build(0, len(nums) - 1)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/12/24/tree1.jpg" />
    // Input: nums = [3,2,1,6,0,5]
    // Output: [6,3,5,null,2,0,null,null,1]
    // Explanation: The recursive calls are as follow:
    // - The largest value in [3,2,1,6,0,5] is 6. Left prefix is [3,2,1] and right suffix is [0,5].
    //     - The largest value in [3,2,1] is 3. Left prefix is [] and right suffix is [2,1].
    //         - Empty array, so no child.
    //         - The largest value in [2,1] is 2. Left prefix is [] and right suffix is [1].
    //             - Empty array, so no child.
    //             - Only one element, so child is a node with value 1.
    //     - The largest value in [0,5] is 5. Left prefix is [0] and right suffix is [].
    //         - Only one element, so child is a node with value 0.
    //         - Empty array, so no child.
    tree1 := constructMaximumBinaryTree([]int{3,2,1,6,0,5})
    fmt.Println(tree1) // &{6 0xc000008060 0xc0000080a8}
    printTree(tree1)
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/12/24/tree2.jpg" />
    // Input: nums = [3,2,1]
    // Output: [3,null,2,null,1]
    tree2 := constructMaximumBinaryTree([]int{3,2,1})
    fmt.Println(tree2) // &{3 <nil> 0xc000008108}
    printTree(tree2)

    tree11 := constructMaximumBinaryTree1([]int{3,2,1,6,0,5})
    fmt.Println(tree11) // &{6 0xc000008060 0xc0000080a8}
    printTree(tree11)

    tree12 := constructMaximumBinaryTree1([]int{3,2,1})
    fmt.Println(tree12) // &{3 <nil> 0xc000008108}
    printTree(tree12)
}