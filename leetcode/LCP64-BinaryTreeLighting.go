package main

// LCP 64. 二叉树灯饰
// 「力扣嘉年华」的中心广场放置了一个巨型的二叉树形状的装饰树。每个节点上均有一盏灯和三个开关。
// 节点值为 0 表示灯处于「关闭」状态，节点值为 1 表示灯处于「开启」状态。
// 每个节点上的三个开关各自功能如下：
//     1. 开关 1：切换当前节点的灯的状态；
//     2. 开关 2：切换 以当前节点为根 的子树中，所有节点上的灯的状态，；
//     3. 开关 3：切换 当前节点及其左右子节点（若存在的话） 上的灯的状态；

// 给定该装饰的初始状态 root，请返回最少需要操作多少次开关，可以关闭所有节点的灯。

// 示例 1：
// 输入：root = [1,1,0,null,null,null,1]
// 输出：2
// 解释：以下是最佳的方案之一，如图所示
// <img src="https://pic.leetcode-cn.com/1629357030-GSbzpY-b71b95bf405e3b223e00b2820a062ba4.gif" />

// 示例 2：
// 输入：root = [1,1,1,1,null,null,1]
// 输出：1
// 解释：以下是最佳的方案，如图所示
// <img src="https://pic.leetcode-cn.com/1629356950-HZsKZC-a4091b6448a0089b4d9e8f0390ff9ac6.gif" />

// 示例 3：
// 输入：root = [0,null,0]
// 输出：0
// 解释：无需操作开关，当前所有节点上的灯均已关闭

// 提示：
//     1 <= 节点个数 <= 10^5
//     0 <= Node.val <= 1

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
func closeLampInTree(root *TreeNode) int {
    type Tuple struct {
        node *TreeNode
        switch2, switch3 bool
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    memo := make(map[Tuple]int) // 记忆化搜索
    var dfs func(*TreeNode, bool, bool) int
    dfs = func(node *TreeNode, switch2, switch3 bool) int {
        if node == nil {
            return 0
        }
        p := Tuple{node, switch2, switch3}
        if res, ok := memo[p]; ok { // 之前计算过
            return res
        }
        if node.Val == 1 == (switch2 == switch3) { // 当前节点为开灯
            res1 := dfs(node.Left, switch2, false) + dfs(node.Right, switch2, false) + 1
            res2 := dfs(node.Left, !switch2, false) + dfs(node.Right, !switch2, false) + 1
            res3 := dfs(node.Left, switch2, true) + dfs(node.Right, switch2, true) + 1
            r123 := dfs(node.Left, !switch2, true) + dfs(node.Right, !switch2, true) + 3
            memo[p] = min(min(res1, res2), min(res3, r123))
        } else { // 当前节点为关灯
            res0 := dfs(node.Left, switch2, false) + dfs(node.Right, switch2, false)
            res12 := dfs(node.Left, !switch2, false) + dfs(node.Right, !switch2, false) + 2
            res13 := dfs(node.Left, switch2, true) + dfs(node.Right, switch2, true) + 2
            res23 := dfs(node.Left, !switch2, true) + dfs(node.Right, !switch2, true) + 2
            memo[p] = min(min(res0, res12), min(res13, res23))
        }
        return memo[p]
    }
    return dfs(root, false, false)
}

func closeLampInTree1(root *TreeNode) int {
    var dfs func(*TreeNode)[]int
    dfs = func(node *TreeNode)[]int {
        if node == nil { return []int{ 0, 0, 0, 0 } }
        on, off, l, r := node.Val, node.Val ^ 1, dfs(node.Left), dfs(node.Right)
        a := min(min(l[0] + r[0] + on,     l[1] + r[1] + off + 1), min(l[2] + r[2] + off + 1, l[3] + r[3] + on + 2)) //全灭
        b := min(min(l[0] + r[0] + on + 1, l[1] + r[1] + off),     min(l[2] + r[2] + off + 2, l[3] + r[3] + on + 1)) //全亮
        c := min(min(l[0] + r[0] + off,    l[1] + r[1] + on + 1),  min(l[2] + r[2] + on + 1,  l[3] + r[3] + off + 2)) //根亮
        d := min(min(l[0] + r[0] + off + 1,l[1] + r[1] + on),      min(l[2] + r[2] + on + 2,  l[3] + r[3] + off + 1)) //根灭
        return []int{ a, b, c, d }
    }
    return dfs(root)[0]
}

func main() {
    // 示例 1：
    // 输入：root = [1,1,0,null,null,null,1]
    // 输出：2
    // 解释：以下是最佳的方案之一，如图所示
    // <img src="https://pic.leetcode-cn.com/1629357030-GSbzpY-b71b95bf405e3b223e00b2820a062ba4.gif" />
    tree1 := buildTree([]*int{ptr(1), ptr(1), ptr(0), nil, nil, nil, ptr(1)})
    printTree(tree1)
    fmt.Println(closeLampInTree(tree1)) // 2
    // 示例 2：
    // 输入：root = [1,1,1,1,null,null,1]
    // 输出：1
    // 解释：以下是最佳的方案，如图所示
    // <img src="https://pic.leetcode-cn.com/1629356950-HZsKZC-a4091b6448a0089b4d9e8f0390ff9ac6.gif" />
    tree2 := buildTree([]*int{ptr(1), ptr(1), ptr(1), ptr(1), nil, nil, ptr(1)})
    printTree(tree2)
    fmt.Println(closeLampInTree(tree2)) // 1
    // 示例 3：
    // 输入：root = [0,null,0]
    // 输出：0
    // 解释：无需操作开关，当前所有节点上的灯均已关闭
    tree3 := buildTree([]*int{ptr(0), nil, ptr(0)})
    printTree(tree3)
    fmt.Println(closeLampInTree(tree3)) // 0

    tree11 := buildTree([]*int{ptr(1), ptr(1), ptr(0), nil, nil, nil, ptr(1)})
    printTree(tree11)
    fmt.Println(closeLampInTree1(tree11)) // 2
    tree12 := buildTree([]*int{ptr(1), ptr(1), ptr(1), ptr(1), nil, nil, ptr(1)})
    printTree(tree12)
    fmt.Println(closeLampInTree1(tree12)) // 1
    tree13 := buildTree([]*int{ptr(0), nil, ptr(0)})
    printTree(tree13)
    fmt.Println(closeLampInTree1(tree13)) // 0
}