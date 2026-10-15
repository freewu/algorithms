package main

// 623. Add One Row to Tree
// Given the root of a binary tree and two integers val and depth, add a row of nodes with value val at the given depth depth.
// Note that the root node is at depth 1.

// The adding rule is:
//     Given the integer depth, for each not null tree node cur at the depth depth - 1, create two tree nodes with value val as cur's left subtree root and right subtree root.
//     cur's original left subtree should be the left subtree of the new left subtree root.
//     cur's original right subtree should be the right subtree of the new right subtree root.
//     If depth == 1 that means there is no depth depth - 1 at all, then create a tree node with value val as the new root of the whole original tree, and the original tree is the new root's left subtree.
    
// Example 1:
//         4                     4
//        /  \                  /  \
//       2    6    ==>        [1]  [1]
//     /  \  /                /      \
//    3    1 5               2        6
//                          / \      /
//                         3   1     5
// <img src="https://assets.leetcode.com/uploads/2021/03/15/addrow-tree.jpg" />
// Input: root = [4,2,6,3,1,5], val = 1, depth = 2
// Output: [4,1,1,2,null,null,6,3,1,5]

// Example 2:
//         4                     4
//        /                     /   
//       2         ==>         2     
//     /  \                   /  \     
//    3    1                [1]  [1]     
//                          /      \      
//                         3        1      
// <img src="https://assets.leetcode.com/uploads/2021/03/11/add2-tree.jpg" />
// Input: root = [4,2,null,3,1], val = 1, depth = 3
// Output: [4,2,null,1,1,3,null,null,1]
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     The depth of the tree is in the range [1, 10^4].
//     -100 <= Node.val <= 100
//     -10^5 <= val <= 10^5
//     1 <= depth <= the depth of tree + 1

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
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if root == nil {
        return nil
    }
    // 如果 depth == 1 意味着 depth - 1 根本没有深度，
    // 那么创建一个树节点，值 val 作为整个原始树的新根，而原始树就是新根的左子树。
    if depth == 1 { // rule 4
        newRoot := TreeNode{val, root, nil}
        return &newRoot
    }
    // 给定整数 depth，对于深度为 depth - 1 的每个非空树节点 cur ，
    // 创建两个值为 val 的树节点作为 cur 的左子树根和右子树根
    if depth == 2 { // rule 1 
        newLeft := TreeNode{val, root.Left, nil}
        newRight := TreeNode{val, nil, root.Right}
        root.Left = &newLeft
        root.Right = &newRight
        return root
    }
    root.Left = addOneRow(root.Left, val, depth - 1) // rule 2   cur 原来的左子树应该是新的左子树根的左子树
    root.Right = addOneRow(root.Right, val, depth - 1) // rule 3 cur 原来的右子树应该是新的右子树根的右子树。
    return root
}

// bfs
func addOneRow1(root *TreeNode, val int, depth int) *TreeNode {
    // 如果 depth == 1 意味着 depth - 1 根本没有深度，
    // 那么创建一个树节点，值 val 作为整个原始树的新根，而原始树就是新根的左子树。
    if depth == 1 {
        return &TreeNode{
            Val: val,
            Left: root,
        }
    }
    queue := []*TreeNode{root}
    currDepth := 1
    for len(queue) > 0 {
        size := len(queue)
        // 给定整数 depth，对于深度为 depth - 1 的每个非空树节点 cur ，
        // 创建两个值为 val 的树节点作为 cur 的左子树根和右子树根
        // 作为depth-1 的子节点
        if currDepth == depth - 1 {
            for i := 0; i < size; i++ {
                currNode := queue[i]
                newLeft := currNode.Left
                newRight := currNode.Right
                currNode.Left = &TreeNode{ Val: val, Left: newLeft,}
                currNode.Right = &TreeNode{ Val: val, Right: newRight,}
            }
            break
        }
        for i := 0; i < size; i++ {
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left) // rule 2   cur 原来的左子树应该是新的左子树根的左子树
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right) // rule 3 cur 原来的右子树应该是新的右子树根的右子树。
            }
        }
        // 移除整行
        queue = queue[size:]
        currDepth++
    }
    return root
}

func main() {
    tree1 := buildTree([]*int{ptr(4), ptr(2), ptr(6), ptr(3), ptr(1), ptr(5)})
    printTree(tree1)
    fmt.Println(addOneRow(tree1,1,2)) // [4,1,1,2,null,null,6,3,1,5]
    tree2 := buildTree([]*int{ptr(4), ptr(2), nil, ptr(3), ptr(1)})
    printTree(tree2)
    fmt.Println(addOneRow(tree2,1,3)) // [4,2,null,1,1,3,null,null,1]

    tree11 := buildTree([]*int{ptr(4), ptr(2), ptr(6), ptr(3), ptr(1), ptr(5)})
    printTree(tree11)
    fmt.Println(addOneRow1(tree11,1,2)) // [4,1,1,2,null,null,6,3,1,5]
    tree12 := buildTree([]*int{ptr(4), ptr(2), nil, ptr(3), ptr(1)})
    printTree(tree12)
    fmt.Println(addOneRow1(tree12,1,3)) // [4,2,null,1,1,3,null,null,1]
}