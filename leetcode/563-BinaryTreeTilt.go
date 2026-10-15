package main

// 563. Binary Tree Tilt
// Given the root of a binary tree, return the sum of every tree node's tilt.
// The tilt of a tree node is the absolute difference between the sum of all left subtree node values and all right subtree node values. If a node does not have a left child, then the sum of the left subtree node values is treated as 0. The rule is similar if the node does not have a right child.

// Example 1:
//          1            1
//        /   \   =>   /   \
//       2     3      0     0
// <img src="https://assets.leetcode.com/uploads/2020/10/20/tilt2.jpg" />
// Input: root = [1,2,3]
// Output: 1
// Explanation: 
// Tilt of node 2 : |0-0| = 0 (no children)
// Tilt of node 3 : |0-0| = 0 (no children)
// Tilt of node 1 : |2-3| = 1 (left subtree is just left child, so sum is 2; right subtree is just right child, so sum is 3)
// Sum of every tilt : 0 + 0 + 1 = 1

// Example 2:
//         4                    6
//       /   \                /   \
//      2     9     =>       2     7
//    /   \     \           /  \     \  
//   3     5     7         0    0     0
// <img src="https://assets.leetcode.com/uploads/2020/10/20/tilt2.jpg" />
// Input: root = [4,2,9,3,5,null,7]
// Output: 15
// Explanation: 
// Tilt of node 3 : |0-0| = 0 (no children)
// Tilt of node 5 : |0-0| = 0 (no children)
// Tilt of node 7 : |0-0| = 0 (no children)
// Tilt of node 2 : |3-5| = 2 (left subtree is just left child, so sum is 3; right subtree is just right child, so sum is 5)
// Tilt of node 9 : |0-7| = 7 (no left child, so sum is 0; right subtree is just right child, so sum is 7)
// Tilt of node 4 : |(3+5+2)-(9+7)| = |10-16| = 6 (left subtree values are 3, 5, and 2, which sums to 10; right subtree values are 9 and 7, which sums to 16)
// Sum of every tilt : 0 + 0 + 0 + 2 + 7 + 6 = 15

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/10/20/tilt3.jpg" />
// Input: root = [21,7,14,1,1,2,2,3,3]
// Output: 9

// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     -1000 <= Node.val <= 1000

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
func findTilt(root *TreeNode) int {
    res := 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left, right := dfs(node.Left), dfs(node.Right)
        res += abs(left - right)
        return left + right + node.Val
    }
    dfs(root)
    return res
}

// bfs
func findTilt1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    type Visit struct {
        Node *TreeNode
        Visited bool
    }
    res, subSum, stack := 0, make(map[*TreeNode]int), []*Visit{ &Visit{root, false} }
    subSum[nil] = 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for len(stack) != 0 {
        node, visited := stack[len(stack) - 1].Node, stack[len(stack) - 1].Visited
        stack = stack[:len(stack) - 1]
        if visited {
            res += abs(subSum[node.Left] - subSum[node.Right])
            subSum[node] = node.Val + subSum[node.Left] + subSum[node.Right]
        } else {
            stack = append(stack, &Visit{node, true})
            if node.Left != nil  { stack = append(stack, &Visit{node.Left, false});  }
            if node.Right != nil { stack = append(stack, &Visit{node.Right, false}); }
        }
    }
    return res
}

func main() {
    // Example 1:
    //          1            1
    //        /   \   =>   /   \
    //       2     3      0     0
    // <img src="https://assets.leetcode.com/uploads/2020/10/20/tilt2.jpg" />
    // Input: root = [1,2,3]
    // Output: 1
    // Explanation: 
    // Tilt of node 2 : |0-0| = 0 (no children)
    // Tilt of node 3 : |0-0| = 0 (no children)
    // Tilt of node 1 : |2-3| = 1 (left subtree is just left child, so sum is 2; right subtree is just right child, so sum is 3)
    // Sum of every tilt : 0 + 0 + 1 = 1
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3)})
    printTree(tree1)
    fmt.Println(findTilt(tree1)) // 1
    // Example 2:
    //         4                    6
    //       /   \                /   \
    //      2     9     =>       2     7
    //    /   \     \           /  \     \  
    //   3     5     7         0    0     0
    // <img src="https://assets.leetcode.com/uploads/2020/10/20/tilt2.jpg" />
    // Input: root = [4,2,9,3,5,null,7]
    // Output: 15
    // Explanation: 
    // Tilt of node 3 : |0-0| = 0 (no children)
    // Tilt of node 5 : |0-0| = 0 (no children)
    // Tilt of node 7 : |0-0| = 0 (no children)
    // Tilt of node 2 : |3-5| = 2 (left subtree is just left child, so sum is 3; right subtree is just right child, so sum is 5)
    // Tilt of node 9 : |0-7| = 7 (no left child, so sum is 0; right subtree is just right child, so sum is 7)
    // Tilt of node 4 : |(3+5+2)-(9+7)| = |10-16| = 6 (left subtree values are 3, 5, and 2, which sums to 10; right subtree values are 9 and 7, which sums to 16)
    // Sum of every tilt : 0 + 0 + 0 + 2 + 7 + 6 = 15
    tree2 := buildTree([]*int{ptr(4), ptr(2), ptr(9), ptr(3), ptr(5), nil, ptr(7)})
    printTree(tree2)
    fmt.Println(findTilt(tree2)) // 15
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/10/20/tilt3.jpg" />
    // Input: root = [21,7,14,1,1,2,2,3,3]
    // Output: 9
    tree3 := buildTree([]*int{ptr(21), ptr(7), ptr(14), ptr(1), ptr(1), ptr(2), ptr(2), ptr(3), ptr(3)})
    printTree(tree3)
    fmt.Println(findTilt(tree3)) // 9

    fmt.Println(findTilt1(tree1)) // 1
    fmt.Println(findTilt1(tree2)) // 15
    fmt.Println(findTilt1(tree3)) // 9
}