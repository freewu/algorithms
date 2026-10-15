package main

// LCP 52. 二叉搜索树染色
// 欢迎各位勇者来到力扣城，本次试炼主题为「二叉搜索树染色」。

// 每位勇士面前设有一个二叉搜索树的模型，模型的根节点为 root，树上的各个节点值均不重复。
// 初始时，所有节点均为蓝色。现
// 在按顺序对这棵二叉树进行若干次操作， ops[i] = [type, x, y] 表示第 i 次操作为：
//     1. type 等于 0 时，将节点值范围在 [x, y] 的节点均染蓝
//     2. type 等于 1 时，将节点值范围在 [x, y] 的节点均染红

// 请返回完成所有染色后，该二叉树中红色节点的数量。

// 注意：
//     题目保证对于每个操作的 x、y 值定出现在二叉搜索树节点中

// 示例 1：
// 输入：root = [1,null,2,null,3,null,4,null,5], ops = [[1,2,4],[1,1,3],[0,3,5]]
// 输出：2
// 解释：  第 0 次操作，将值为 2、3、4 的节点染红； 
//        第 1 次操作，将值为 1、2、3 的节点染红； 
//        第 2 次操作，将值为 3、4、5 的节点染蓝； 
//        因此，最终值为 1、2 的节点为红色节点，返回数量 2
//        <img src="https://pic.leetcode-cn.com/1649833948-arSlXd-image.png" />

// 示例 2：
// 输入：root = [4,2,7,1,null,5,null,null,null,null,6] ops = [[0,2,2],[1,1,5],[0,4,5],[1,5,7]]
// 输出：5
// 解释： 第 0 次操作，将值为 2 的节点染蓝； 
//       第 1 次操作，将值为 1、2、4、5 的节点染红； 
//       第 2 次操作，将值为 4、5 的节点染蓝； 
//       第 3 次操作，将值为 5、6、7 的节点染红； 
//       因此，最终值为 1、2、5、6、7 的节点为红色节点，返回数量 5
//       <img src="https://pic.leetcode-cn.com/1649833763-BljEbP-image.png" />

// 提示：
//     1 <= 二叉树节点数量 <= 10^5
//     1 <= ops.length <= 10^5
//     ops[i].length == 3
//     ops[i][0] 仅为 0 or 1
//     0 <= ops[i][1] <= ops[i][2] <= 10^9
//     0 <= 节点值 <= 10^9

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
func getNumber(root *TreeNode, ops [][]int) int {
    res, n := 0, len(ops)
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil { return }
        for i := n - 1; i >= 0; i-- {
            if root.Val >= ops[i][1] && root.Val <= ops[i][2] {
                res += ops[i][0]
                break
            }
        }
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)
    return res
}

// bfs
func getNumber1(root *TreeNode, ops [][]int) int {
    res, queue := 0, []*TreeNode{ root }
    for len(queue) > 0 {
        cur := queue[len(queue) - 1]
        queue = queue[:len(queue) - 1]
        if cur.Left != nil {
            queue = append(queue, cur.Left)
        }
        if cur.Right != nil {
            queue = append(queue, cur.Right)
        }
        v := cur.Val
        for i := len(ops) - 1; i >= 0; i-- {
            if ops[i][1] <= v && v <= ops[i][2] {
                if ops[i][0] == 1 {
                    res++
                }
                break
            }
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：root = [1,null,2,null,3,null,4,null,5], ops = [[1,2,4],[1,1,3],[0,3,5]]
    // 输出：2
    // 解释：  第 0 次操作，将值为 2、3、4 的节点染红； 
    //        第 1 次操作，将值为 1、2、3 的节点染红； 
    //        第 2 次操作，将值为 3、4、5 的节点染蓝； 
    //        因此，最终值为 1、2 的节点为红色节点，返回数量 2
    //        <img src="https://pic.leetcode-cn.com/1649833948-arSlXd-image.png" />
    tree1 := buildTree([]*int{ptr(1), nil, ptr(2), nil, ptr(3), nil, ptr(4), nil, ptr(5)})
    printTree(tree1)
    fmt.Println(getNumber(tree1, [][]int{{1,2,4},{1,1,3},{0,3,5}})) // 2
    // 示例 2：
    // 输入：root = [4,2,7,1,null,5,null,null,null,null,6] ops = [[0,2,2],[1,1,5],[0,4,5],[1,5,7]]
    // 输出：5
    // 解释： 第 0 次操作，将值为 2 的节点染蓝； 
    //       第 1 次操作，将值为 1、2、4、5 的节点染红； 
    //       第 2 次操作，将值为 4、5 的节点染蓝； 
    //       第 3 次操作，将值为 5、6、7 的节点染红； 
    //       因此，最终值为 1、2、5、6、7 的节点为红色节点，返回数量 5
    //       <img src="https://pic.leetcode-cn.com/1649833763-BljEbP-image.png" />
    tree2 := buildTree([]*int{ptr(4), ptr(2), ptr(7), ptr(1), nil, ptr(5), nil, nil, nil, nil, ptr(6)})
    printTree(tree2)
    fmt.Println(getNumber(tree2, [][]int{{0,2,2},{1,1,5},{0,4,5},{1,5,7}})) // 5

    tree11 := buildTree([]*int{ptr(1), nil, ptr(2), nil, ptr(3), nil, ptr(4), nil, ptr(5)})
    printTree(tree11)
    fmt.Println(getNumber1(tree11, [][]int{{1,2,4},{1,1,3},{0,3,5}})) // 2
    tree12 := buildTree([]*int{ptr(4), ptr(2), ptr(7), ptr(1), nil, ptr(5), nil, nil, nil, nil, ptr(6)})
    printTree(tree12)
    fmt.Println(getNumber1(tree12, [][]int{{0,2,2},{1,1,5},{0,4,5},{1,5,7}})) // 5
}