package main

// LCP 67. 装饰树
// 力扣嘉年华上的 DIY 手工展位准备了一棵缩小版的二叉装饰树root和灯饰，你需要将灯饰逐一插入装饰树中，要求如下：
//     1. 完成装饰的二叉树根结点与root的根结点值相同
//     2. 若一个节点拥有父节点，则在该节点和他的父节点之间插入一个灯饰（即插入一个值为-1的节点）。具体地：
//         2.1 在一个 父节点 x 与其左子节点 y 之间添加 -1 节点， 节点 -1、节点 y 为各自父节点的左子节点，
//         2.2 在一个 父节点 x 与其右子节点 y 之间添加 -1 节点， 节点 -1、节点 y 为各自父节点的右子节点，

// 现给定二叉树的根节点root，请返回完成装饰后的树的根节点。

// 示例 1：
// 输入：root = [7,5,6]
// 输出：[7,-1,-1,5,null,null,6]
// 解释：如下图所示，
// <img src="https://pic.leetcode-cn.com/1663575757-yRLGaq-image.png"/>

// 示例 2：
// 输入：root = [3,1,7,3,8,null,4]
// 输出：[3,-1,-1,1,null,null,7,-1,-1,null,-1,3,null,null,8,null,4]
// 解释：如下图所示
// <img src="https://pic.leetcode-cn.com/1663577920-sjrAYH-image.png"/>

// 提示：
//     0 <= root.Val <= 1000 root节点数量范围为[1, 10^5]

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
func expandBinaryTree1(root *TreeNode) *TreeNode {
    var dfs func(root, parent *TreeNode, flag bool) 
    dfs = func(root, parent *TreeNode, flag bool) {
        if root == nil { return }
        if parent != nil {
            if flag {
                parent.Left  = &TreeNode{ Val: -1, Left:  root }
            } else {
                parent.Right = &TreeNode{ Val: -1, Right: root }
            }
        }
        dfs(root.Left, root, true)
        dfs(root.Right, root, false)
    }
    dfs(root, nil, true)
    return root
}

func expandBinaryTree(node *TreeNode) *TreeNode {
    if node == nil { return nil }
    l := expandBinaryTree(node.Left)
    if l != nil {
        node.Left = &TreeNode{ Val: -1, Left: l }
    }
    r := expandBinaryTree(node.Right)
    if r != nil {
        node.Right = &TreeNode{ Val:-1, Right: r }
    }
    return node
}

func main() {
    // 示例 1：
    // 输入：root = [7,5,6]
    // 输出：[7,-1,-1,5,null,null,6]
    // 解释：如下图所示，
    // <img src="https://pic.leetcode-cn.com/1663575757-yRLGaq-image.png"/>
    tree1 := buildTree([]*int{ptr(7), ptr(5), ptr(6)})
    printTree(tree1)
    res1 := expandBinaryTree(tree1)
    printTree(res1)
    fmt.Println(res1) // 
    // 示例 2：
    // 输入：root = [3,1,7,3,8,null,4]
    // 输出：[3,-1,-1,1,null,null,7,-1,-1,null,-1,3,null,null,8,null,4]
    // 解释：如下图所示
    // <img src="https://pic.leetcode-cn.com/1663577920-sjrAYH-image.png"/>
    tree2 := buildTree([]*int{ptr(3), ptr(1), ptr(7), ptr(3), ptr(8), nil, ptr(4)})
    printTree(tree2)
    res2 := expandBinaryTree(tree2)
    printTree(res2)
    fmt.Println(res2) // 

    tree11 := buildTree([]*int{ptr(7), ptr(5), ptr(6)})
    printTree(tree11)
    res3 := expandBinaryTree1(tree11)
    printTree(res3)
    fmt.Println(res3) // 
    tree12 := buildTree([]*int{ptr(3), ptr(1), ptr(7), ptr(3), ptr(8), nil, ptr(4)})
    printTree(tree12)
    res4 := expandBinaryTree1(tree12)
    printTree(res4)
    fmt.Println(res4) // 
}