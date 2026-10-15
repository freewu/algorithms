package main

// 998. Maximum Binary Tree II
// A maximum tree is a tree where every node has a value greater than any other value in its subtree.
// You are given the root of a maximum binary tree and an integer val.
//     Just as in the previous problem, the given tree was constructed from a list a (root = Construct(a)) recursively with the following Construct(a) routine:
//     If a is empty, return null.
//     Otherwise, let a[i] be the largest element of a. Create a root node with the value a[i].
//     The left child of root will be Construct([a[0], a[1], ..., a[i - 1]]).
//     The right child of root will be Construct([a[i + 1], a[i + 2], ..., a[a.length - 1]]).
//     Return root.

// Note that we were not given a directly, only a root node root = Construct(a).
// Suppose b is a copy of a with the value val appended to it. It is guaranteed that b has unique values.
// Return Construct(b).

// Example 1:
//          4                      5
//        /   \                   /
//       1     3   =>            4
//             /               /   \
//            2               1     3
//                                  /
//                                 2
// <img src="https://assets.leetcode.com/uploads/2021/08/09/maxtree1.JPG" />
// Input: root = [4,1,3,null,null,2], val = 5
// Output: [5,4,null,1,3,null,null,2]
// Explanation: a = [1,4,2,3], b = [1,4,2,3,5]

// Example 2:
//          5                 5
//        /   \             /   \
//       2     4   =>      2     4
//        \                  \     \
//          1                 1      3
// <img src="https://assets.leetcode.com/uploads/2021/08/09/maxtree21.JPG" />
// Input: root = [5,2,4,null,1], val = 3
// Output: [5,2,4,null,1,null,3]
// Explanation: a = [2,1,5,4], b = [2,1,5,4,3]

// Example 3:
//           5                5
//         /   \            /   \
//        2     3   =>     2     4
//         \                \    /
//          1                1  3
// <img src="https://assets.leetcode.com/uploads/2021/08/09/maxtree3.JPG" />
// Input: root = [5,2,3,null,1], val = 4
// Output: [5,2,4,null,1,3]
// Explanation: a = [2,1,5,3], b = [2,1,5,3,4]

// Constraints:
//     The number of nodes in the tree is in the range [1, 100].
//     1 <= Node.val <= 100
//     All the values of the tree are unique.
//     1 <= val <= 100

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

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
    if val > root.Val {
        root = &TreeNode{ Val: val, Left: root, }
        return root
    }
    if root.Right == nil {
        root.Right = &TreeNode{ Val: val, }
    }else {
        root.Right = insertIntoMaxTree(root.Right, val)
    }
    return root
}

func main() {
    // Example 1:
    //          4                      5
    //        /   \                   /
    //       1     3   =>            4
    //             /               /   \
    //            2               1     3
    //                                  /
    //                                 2
    // <img src="https://assets.leetcode.com/uploads/2021/08/09/maxtree1.JPG" />
    // Input: root = [4,1,3,null,null,2], val = 5
    // Output: [5,4,null,1,3,null,null,2]
    // Explanation: a = [1,4,2,3], b = [1,4,2,3,5]
    tree1 := buildTree([]*int{ptr(4), ptr(1), ptr(3), nil, nil, ptr(2)})
    printTree(tree1)
    fmt.Println(insertIntoMaxTree(tree1, 5)) // 
    // Example 2:
    //          5                 5
    //        /   \             /   \
    //       2     4   =>      2     4
    //        \                  \     \
    //          1                 1      3
    // <img src="https://assets.leetcode.com/uploads/2021/08/09/maxtree21.JPG" />
    // Input: root = [5,2,4,null,1], val = 3
    // Output: [5,2,4,null,1,null,3]
    // Explanation: a = [2,1,5,4], b = [2,1,5,4,3]
    tree2 := buildTree([]*int{ptr(5), ptr(2), ptr(4), nil, ptr(1)})
    printTree(tree2)
    fmt.Println(insertIntoMaxTree(tree2, 3)) // 
    // Example 3:
    //           5                5
    //         /   \            /   \
    //        2     3   =>     2     4
    //         \                \    /
    //          1                1  3
    // <img src="https://assets.leetcode.com/uploads/2021/08/09/maxtree3.JPG" />
    // Input: root = [5,2,3,null,1], val = 4
    // Output: [5,2,4,null,1,3]
    // Explanation: a = [2,1,5,3], b = [2,1,5,3,4]
    tree3 := buildTree([]*int{ptr(5), ptr(2), ptr(3), nil, ptr(1)})
    printTree(tree3)
    fmt.Println(insertIntoMaxTree(tree3, 4)) // 
}