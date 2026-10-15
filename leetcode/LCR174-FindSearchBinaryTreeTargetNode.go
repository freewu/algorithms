package main

// LCR 174. 寻找二叉搜索树中的目标节点
// 某公司组织架构以二叉搜索树形式记录，节点值为处于该职位的员工编号。请返回第 cnt 大的员工编号。

// 示例 1：
// <img src="https://pic.leetcode.cn/1695101634-kzHKZW-image.png" />
// 输入：root = [7, 3, 9, 1, 5], cnt = 2
//        7
//       / \
//      3   9
//     / \
//    1   5
// 输出：7

// 示例 2：
// <img src="https://pic.leetcode.cn/1695101636-ESZtLa-image.png" />
// 输入: root = [10, 5, 15, 2, 7, null, 20, 1, null, 6, 8], cnt = 4
//        10
//       / \
//      5   15
//     / \    \
//    2   7    20
//   /   / \ 
//  1   6   8
// 输出: 8

// 提示：
//     1 ≤ cnt ≤ 二叉搜索树元素个数

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
func findTargetNode(root *TreeNode, cnt int) int {
    stack := []*TreeNode{}
    // 正常中序遍历为 左中右（可以得到第cnt【小】的数值）
    // 本题变化为：中序遍历（右中左）（可以得到第cnt【大】的数值）
    for root != nil || len(stack) > 0 {
        if root != nil {
            stack = append(stack, root)
            root = root.Right // 正常这里是 root = root.Left
        } else {
            root = stack[len(stack)-1] // pop
            stack = stack[:len(stack)-1]
            cnt--
            if cnt == 0 {
                return root.Val
            }
            root = root.Left
        }
    }
    return 0
}

func findTargetNode1(root *TreeNode, cnt int) int {
    if root == nil {
        return 0
    }
    res, count := 0, 0
    var inorder func(root *TreeNode)
    inorder = func(root *TreeNode) {
        if root == nil {
            return
        }
        inorder(root.Right) // right 大
        count++
        if count == cnt {
            res = root.Val
            return
        }
        inorder(root.Left)
    }
    inorder(root)
    return res
}

func main() {
    // 示例 1：
    // <img src="https://pic.leetcode.cn/1695101634-kzHKZW-image.png" />
    // 输入：root = [7, 3, 9, 1, 5], cnt = 2
    //        7
    //       / \
    //      3   9
    //     / \
    //    1   5
    // 输出：7
    tree1 := &TreeNode {
        3,
        &TreeNode{3, &TreeNode{1, nil, nil, }, &TreeNode{5, nil, nil, }, },
        &TreeNode{9, nil, nil, },
    }
    fmt.Println(findTargetNode(tree1, 2)) // 7
    // 示例 2：
    // <img src="https://pic.leetcode.cn/1695101636-ESZtLa-image.png" />
    // 输入: root = [10, 5, 15, 2, 7, null, 20, 1, null, 6, 8], cnt = 4
    //        10
    //       / \
    //      5   15
    //     / \    \
    //    2   7    20
    //   /   / \ 
    //  1   6   8
    // 输出: 8
    tree2 := &TreeNode {
        10,
        &TreeNode{5,  &TreeNode{2, &TreeNode{1, nil, nil, }, nil, }, &TreeNode{7, &TreeNode{6, nil, nil, }, &TreeNode{8, nil, nil, }, }, },
        &TreeNode{15, nil, &TreeNode{20, nil, nil, }, },
    }
    fmt.Println(findTargetNode(tree2, 4)) // 8

    tree11 := &TreeNode {
        3,
        &TreeNode{3, &TreeNode{1, nil, nil, }, &TreeNode{5, nil, nil, }, },
        &TreeNode{9, nil, nil, },
    }
    fmt.Println(findTargetNode1(tree11, 2)) // 7
    tree12 := &TreeNode {
        10,
        &TreeNode{5,  &TreeNode{2, &TreeNode{1, nil, nil, }, nil, }, &TreeNode{7, &TreeNode{6, nil, nil, }, &TreeNode{8, nil, nil, }, }, },
        &TreeNode{15, nil, &TreeNode{20, nil, nil, }, },
    }
    fmt.Println(findTargetNode1(tree12, 4)) // 8
}