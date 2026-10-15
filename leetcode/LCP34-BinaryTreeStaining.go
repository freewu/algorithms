package main

// LCP 34. 二叉树染色
// 小扣有一个根结点为 root 的二叉树模型，初始所有结点均为白色，可以用蓝色染料给模型结点染色，模型的每个结点有一个 val 价值。
// 小扣出于美观考虑，希望最后二叉树上每个蓝色相连部分的结点个数不能超过 k 个，求所有染成蓝色的结点价值总和最大是多少？

// 示例 1：
// 输入：root = [5,2,3,4], k = 2
// 输出：12
// 解释：结点 5、3、4 染成蓝色，获得最大的价值 5+3+4=12
// <img src="https://pic.leetcode-cn.com/1616126267-BqaCRj-image.png" />

// 示例 2：
// 输入：root = [4,1,3,9,null,null,2], k = 2
// 输出：16
// 解释：结点 4、3、9 染成蓝色，获得最大的价值 4+3+9=16
// <img src="https://pic.leetcode-cn.com/1616126301-gJbhba-image.png" />

// 提示：
//     1 <= k <= 10
//     1 <= val <= 10000
//     1 <= 结点数量 <= 10000

import "fmt"
import "slices"

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
func maxValue(root *TreeNode, k int) int {
    // 每个节点可以有两种处理结果：染色或者不染色
    // 根据题意显然是要用后序遍历，自底向上，由下层给出结果，供上层使用和选择
    // 定义：dp[0] 表示：以某个节点为根节点，根节点不染色的情况
    // 定义：dp[i] 表示：以某个节点为根节点，（根节点染色的情况），蓝色相连节点为i个的最大染色值， 1 <= i <= k， 
    // 所以dp数组的长度应为（k+1）
    // 根据每个节点的染色情况：
    // 不染色 ： dp[0] = 左子节点处理的最大值+右子节点处理的最大值
    // 染色： dp[i] = left_dp[j] + right_dp[i-1-j] + root.Val     (i = [1,k] , j=[0,i）)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(root *TreeNode, k int) []int
    dfs = func(root *TreeNode, k int) []int {
        dp := make([]int, k + 1)
        if root == nil { return dp }
        left_dp, right_dp := dfs(root.Left, k), dfs(root.Right, k)
        dp[0] = slices.Max(left_dp) + slices.Max(right_dp)
        for i := 1; i <= k; i++ {
            for j := 0; j < i; j++ {
                // 一共有 i， 左子树分配 j个， 右子树分配 i - j- 1 个 （根结点一个）
                dp[i] = max(dp[i], (left_dp[j] + right_dp[i - j - 1] + root.Val))
            }
        }
        return dp
    }
    return slices.Max(dfs(root, k))
}

func main() {
    // 示例 1：
    // 输入：root = [5,2,3,4], k = 2
    // 输出：12
    // 解释：结点 5、3、4 染成蓝色，获得最大的价值 5+3+4=12
    // <img src="https://pic.leetcode-cn.com/1616126267-BqaCRj-image.png" />
    tree1 := buildTree([]*int{ptr(5), ptr(2), ptr(3), ptr(4)})
    printTree(tree1)
    fmt.Println(maxValue(tree1, 2)) // 12
    // 示例 2：
    // 输入：root = [4,1,3,9,null,null,2], k = 2
    // 输出：16
    // 解释：结点 4、3、9 染成蓝色，获得最大的价值 4+3+9=16
    // <img src="https://pic.leetcode-cn.com/1616126301-gJbhba-image.png" />
    tree2 := buildTree([]*int{ptr(4), ptr(1), ptr(3), ptr(9), nil, nil, ptr(2)})
    printTree(tree2)
    fmt.Println(maxValue(tree2, 2)) // 16
}