package main

// 106. Construct Binary Tree from Inorder and Postorder Traversal
// Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.

// Constraints:
// 		1 <= inorder.length <= 3000
// 		postorder.length == inorder.length
// 		-3000 <= inorder[i], postorder[i] <= 3000
// 		inorder and postorder consist of unique values.
// 		Each value of postorder also appears in inorder.
// 		inorder is guaranteed to be the inorder traversal of the tree.
// 		postorder is guaranteed to be the postorder traversal of the tree.

// Example 1:
//     3
//    / \
//   9  20
//     /  \
//    15   7
// 	Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 	Output: [3,9,20,null,null,15,7]

// Example 2:
// 	Input: inorder = [-1], postorder = [-1]
// 	Output: [-1]

// 解题思路:
// 	根据一棵树的中序遍历与后序遍历构造二叉树

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// == 二叉树构建 & 打印输出 start ==================================================================
// func ptr(v int) *int {
//     return &v
// }

// // buildTree 按层序数组（LeetCode 格式，nil 表示 null）创建二叉树
// func buildTree(nums []*int) *TreeNode {
//     if len(nums) == 0 || nums[0] == nil {
//         return nil
//     }
//     root := &TreeNode{Val: *nums[0]}
//     queue := []*TreeNode{root}
//     i := 1
//     for len(queue) > 0 && i < len(nums) {
//         node := queue[0]
//         queue = queue[1:]
//         if nums[i] != nil {
//             node.Left = &TreeNode{Val: *nums[i]}
//             queue = append(queue, node.Left)
//         }
//         i++
//         if i < len(nums) && nums[i] != nil {
//             node.Right = &TreeNode{Val: *nums[i]}
//             queue = append(queue, node.Right)
//         }
//         i++
//     }
//     return root
// }

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

// 直接传入需要的 slice 范围作为输入, 可以避免申请对应 inorder 索引的内存,
func buildTree(inorder []int, postorder []int) *TreeNode {
    n := len(postorder)
    if n == 0 {
        return nil
    }
    root := &TreeNode{Val: postorder[n-1]}
    postorder = postorder[:n-1]
    for pos, node := range inorder {
        if node == root.Val {
            root.Left = buildTree(inorder[:pos], postorder[:len(inorder[:pos])])
            root.Right = buildTree(inorder[pos+1:], postorder[len(inorder[:pos]):])
        }
    }
    return root
}

// 
func buildTree1(inorder []int, postorder []int) *TreeNode {
    inPos := make(map[int]int)
    for i := 0; i < len(inorder); i++ {
        inPos[inorder[i]] = i
    }
    var dfs func(post []int, postStart int, postEnd int, inStart int, inPos map[int]int) *TreeNode
    dfs = func(post []int, postStart int, postEnd int, inStart int, inPos map[int]int) *TreeNode {
        if postStart > postEnd {
            return nil
        }
        root := &TreeNode{Val: post[postEnd]}
        rootIdx := inPos[post[postEnd]]
        leftLen := rootIdx - inStart
        root.Left = dfs(post, postStart, postStart+leftLen-1, inStart, inPos)
        root.Right = dfs(post, postStart+leftLen, postEnd-1, rootIdx+1, inPos)
        return root
    }
    return dfs(postorder, 0, len(postorder)-1, 0, inPos)
}


// best solution
func buildTree2(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil
    }
    root := &TreeNode{postorder[len(postorder)-1], nil, nil}
    i := 0
    for ; i < len(inorder); i++ {
        if inorder[i] == postorder[len(postorder)-1] {
            break
        }
    }
    root.Left = buildTree2(inorder[0:i], postorder[0:i])
    root.Right = buildTree2(inorder[i+1:], postorder[i:len(postorder)-1])
    return root
}

func main() {
    // Example 1:
    //     3
    //    / \
    //   9  20
    //     /  \
    //    15   7
    // 	Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
    // 	Output: [3,9,20,null,null,15,7]
    tree1 := buildTree([]int{9,3,15,20,7},[]int{9,15,7,20,3})
    fmt.Printf("buildTree([]int{9,3,15,20,7},[]int{9,15,7,20,3}) = %v\n",tree1) // [3,9,20,null,null,15,7]
    printTree(tree1)
    // Example 2:
    // 	Input: inorder = [-1], postorder = [-1]
    // 	Output: [-1]
    tree2 := buildTree([]int{-1},[]int{-1})
    fmt.Printf("buildTree([]int{-1},[]int{-1}) = %v\n",tree2) // [-1]
    printTree(tree2)

    tree11 := buildTree1([]int{9,3,15,20,7},[]int{9,15,7,20,3})
    fmt.Printf("buildTree1([]int{9,3,15,20,7},[]int{9,15,7,20,3}) = %v\n",tree11) // [3,9,20,null,null,15,7]
    printTree(tree11)
    tree12 := buildTree1([]int{-1},[]int{-1})
    fmt.Printf("buildTree1([]int{-1},[]int{-1}) = %v\n",tree12) // [-1]
    printTree(tree12)

    tree21 := buildTree2([]int{9,3,15,20,7},[]int{9,15,7,20,3})
    fmt.Printf("buildTree2([]int{9,3,15,20,7},[]int{9,15,7,20,3}) = %v\n",tree21) // [3,9,20,null,null,15,7]
    printTree(tree21)
    tree22 := buildTree2([]int{-1},[]int{-1})
    fmt.Printf("buildTree2([]int{-1},[]int{-1}) = %v\n",tree22) // [-1]
    printTree(tree22)
}
