package main

// LCR 047. 二叉树剪枝
// 给定一个二叉树 根节点 root ，树的每个节点的值要么是 0，要么是 1。
// 请剪除该二叉树中所有节点的值为 0 的子树。

// 节点 node 的子树为 node 本身，以及所有 node 的后代。

// 示例 1:
// 输入: [1,null,0,0,1]
// 输出: [1,null,0,null,1] 
// 解释: 
// 只有红色节点满足条件“所有不包含 1 的子树”。
// 右图为返回的答案。
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/06/1028_2.png" />

// 示例 2:
// 输入: [1,0,1,0,0,0,1]
// 输出: [1,null,1,null,1]
// 解释: 
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/06/1028_1.png" />

// 示例 3:
// 输入: [1,1,0,1,1,0,1,0]
// 输出: [1,1,0,1,1,null,1]
// 解释: 
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/05/1028.png" />

// 提示:
//     二叉树的节点个数的范围是 [1,200]
//     二叉树节点的值只会是 0 或 1

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
func pruneTree1(root *TreeNode) *TreeNode {
    var dfs func(node *TreeNode) bool
    dfs = func(node *TreeNode) bool {
        if node == nil {
            return false
        }
        left := dfs(node.Left)
        if !left {
            node.Left = nil
        }
        right := dfs(node.Right) 
        if !right {
            node.Right = nil
        }
        return node.Val == 1 || left || right
    }
    if !dfs(root) {
        return nil
    }
    return root
}

func pruneTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left, root.Right = pruneTree(root.Left), pruneTree(root.Right)
    if root.Val == 0 && root.Left == nil && root.Right == nil { // 后序遍历位置，判断自己是否是值为 0 的叶子节点
        return nil
    }
    return root // 如果不是，正常返回
}

func main() {
    // Example 1:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/06/1028_2.png" />
    // Input: root = [1,null,0,0,1]
    // Output: [1,null,0,null,1]
    // Explanation: 
    // Only the red nodes satisfy the property "every subtree not containing a 1".
    // The diagram on the right represents the answer.
    tree1 := buildTree([]*int{ptr(1), nil, ptr(0), ptr(0), ptr(1)})
    printTree(tree1)
    res1 := pruneTree(tree1)
    printTree(res1)
    fmt.Println(res1) 
    // Example 2:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/06/1028_1.png" />
    // Input: root = [1,0,1,0,0,0,1]
    // Output: [1,null,1,null,1]
    tree2 := buildTree([]*int{ptr(1), ptr(0), ptr(1), ptr(0), ptr(0), ptr(0), ptr(1)})
    printTree(tree2)
    res2 := pruneTree(tree2)
    printTree(res2)
    fmt.Println(res2) 
    // Example 3:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/05/1028.png" />
    // Input: root = [1,1,0,1,1,0,1,0]
    // Output: [1,1,0,1,1,null,1]
    tree3 := buildTree([]*int{ptr(1), ptr(1), ptr(0), ptr(1), ptr(1), ptr(0), ptr(1), ptr(0)})
    printTree(tree3)
    res3 := pruneTree(tree3)
    printTree(res3)
    fmt.Println(res3) 

    tree11 := buildTree([]*int{ptr(1), nil, ptr(0), ptr(0), ptr(1)})
    printTree(tree11)
    res4 := pruneTree1(tree11)
    printTree(res4)
    fmt.Println(res4) 
    tree12 := buildTree([]*int{ptr(1), ptr(0), ptr(1), ptr(0), ptr(0), ptr(0), ptr(1)})
    printTree(tree12)
    res5 := pruneTree1(tree12)
    printTree(res5)
    fmt.Println(res5) 
    tree13 := buildTree([]*int{ptr(1), ptr(1), ptr(0), ptr(1), ptr(1), ptr(0), ptr(1), ptr(0)})
    printTree(tree13)
    res6 := pruneTree1(tree13)
    printTree(res6)
    fmt.Println(res6) 
}