package main

// 1602. Find Nearest Right Node in Binary Tree
// Given the root of a binary tree and a node u in the tree, 
// return the nearest node on the same level that is to the right of u, 
// or return null if u is the rightmost node in its level.

// Example 1:
//              1
//           /     \
//          2       3
//           \     /  \
//           (4)  5     6
// <img src="https://assets.leetcode.com/uploads/2020/09/24/p3.png" />
// Input: root = [1,2,3,null,4,5,6], u = 4
// Output: 5
// Explanation: The nearest node on the same level to the right of node 4 is node 5.

// Example 2:
//         3
//           \
//            4
//           /
//         (2)
// <img src="https://assets.leetcode.com/uploads/2020/09/23/p2.png" />
// Input: root = [3,null,4,2], u = 2
// Output: null
// Explanation: There are no nodes to the right of 2.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     1 <= Node.val <= 10^5
//     All values in the tree are distinct.
//     u is a node in the binary tree rooted at root.

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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
    // iter to find the target u
    height_stack := []int{0}
    iter_stack := []*TreeNode{ root }
    for len(iter_stack) > 0 {
        cur := iter_stack[0]
        height := height_stack[0]
        if cur.Val == u.Val {
            if len(height_stack) == 1 {
                return nil
            }
            next_h := height_stack[1]
            if height == next_h {
                return iter_stack[1]
            }
            break
        }
        iter_stack = iter_stack[1:]
        height_stack = height_stack[1:]
        if cur.Left != nil {
            iter_stack = append(iter_stack, cur.Left)
            height_stack = append(height_stack, height+1)
        }
        if cur.Right != nil {
            iter_stack = append(iter_stack, cur.Right)
            height_stack = append(height_stack, height+1)
        }
    }
    return nil
}

func findNearestRightNode1(root *TreeNode, u *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        levelSize := len(queue)
        for i := 0; i < levelSize; i++ {
            curr := queue[0]
            queue = queue[1:]
            if curr.Val == u.Val {
                if i == levelSize - 1 {
                    return nil
                }
                return queue[0]
            }
            if curr.Left != nil {
                queue = append(queue, curr.Left)
            }
            if curr.Right != nil {
                queue = append(queue, curr.Right)
            }
        }
    }
    return nil 
}

func main() {
    // Example 1:
    //              1
    //           /     \
    //          2       3
    //           \     /  \
    //           (4)  5     6
    // <img src="https://assets.leetcode.com/uploads/2020/09/24/p3.png" />
    // Input: root = [1,2,3,null,4,5,6], u = 4
    // Output: 5
    // Explanation: The nearest node on the same level to the right of node 4 is node 5.
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3), nil, ptr(4), ptr(5), ptr(6)})
    printTree(tree1)
    fmt.Println(findNearestRightNode(tree1,  &TreeNode { 4, nil, nil, })) // &{5 <nil> <nil>}
    // Example 2:
    //         3
    //           \
    //            4
    //           /
    //         (2)
    // <img src="https://assets.leetcode.com/uploads/2020/09/23/p2.png" />
    // Input: root = [3,null,4,2], u = 2
    // Output: null
    // Explanation: There are no nodes to the right of 2.
    tree2 := buildTree([]*int{ptr(3), nil, ptr(4), ptr(2)})
    printTree(tree2)
    fmt.Println(findNearestRightNode(tree2,  &TreeNode { 2, nil, nil, })) // nil

    fmt.Println(findNearestRightNode1(tree1,  &TreeNode { 4, nil, nil, })) // &{5 <nil> <nil>}
    fmt.Println(findNearestRightNode1(tree2,  &TreeNode { 2, nil, nil, })) // nil
}