package main

// LCP 10. 二叉树任务调度
// 任务调度优化是计算机性能优化的关键任务之一。
// 在任务众多时，不同的调度策略可能会得到不同的总体执行时间，因此寻求一个最优的调度方案是非常有必要的。

// 通常任务之间是存在依赖关系的，即对于某个任务，你需要先完成他的前导任务（如果非空），才能开始执行该任务。
// 我们保证任务的依赖关系是一棵二叉树，其中 root 为根任务，root.left 和 root.right 为他的两个前导任务（可能为空），
// root.val 为其自身的执行时间。

// 在一个 CPU 核执行某个任务时，我们可以在任何时刻暂停当前任务的执行，并保留当前执行进度。
// 在下次继续执行该任务时，会从之前停留的进度开始继续执行。
// 暂停的时间可以不是整数。

// 现在，系统有两个 CPU 核，即我们可以同时执行两个任务，但是同一个任务不能同时在两个核上执行。
// 给定这颗任务树，请求出所有任务执行完毕的最小时间。

// 示例 1：
//         47
//       /    \
//     74     31
// <img src="https://pic.leetcode-cn.com/3522fbf8ce4ebb20b79019124eb9870109fdfe97fe9da99f6c20c07ceb1c60b3-image.png" />
// 输入：root = [47, 74, 31]
// 输出：121
// 解释：根节点的左右节点可以并行执行31分钟，剩下的43+47分钟只能串行执行，因此总体执行时间是121分钟。

// 示例 2：
//             15
//            /
//          12
//         /
//        24
//       /   \
//      27   26
// <img src="https://pic.leetcode-cn.com/13accf172ee4a660d241e25901595d55b759380b090890a17e6e7bd51a143e3f-image.png" />
// 输入：root = [15, 21, null, 24, null, 27, 26]
// 输出：87

// 示例 3：
//         1
//       /   \
//      3     2
//          /    \
//         4      4
// <img src="https://pic.leetcode-cn.com/bef743a12591aafb9047dd95d335b8083dfa66e8fdedc63f50fd406b4a9d163a-image.png" />
// 输入：root = [1,3,2,null,null,4,4]
// 输出：7.5

// 限制：
//     1 <= 节点数量 <= 1000
//     1 <= 单节点执行时间 <= 1000

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
func minimalExecTime(root *TreeNode) float64 {
    var dfs func(node *TreeNode) (mx, par float64)
    dfs = func(node *TreeNode) (mx, par float64) {
        if node == nil {
            return 0.0, 0.0
        }
        mx = float64(node.Val)
        lmx, lpar := dfs(node.Left)
        rmx, rpar := dfs(node.Right)
        mx += (lmx + rmx)
        if lmx < rmx {
            lmx, rmx = rmx, lmx
            lpar, rpar = rpar, lpar
        }
        if lmx - rmx <= 2 * lpar {
            par = (lmx + rmx) / 2
        } else {
            par = lpar + rmx
        }
        return mx, par
    }
    mx, par := dfs(root)
    return mx - par
}

func main() {
    // 示例 1：
    //         47
    //       /    \
    //     74     31
    // <img src="https://pic.leetcode-cn.com/3522fbf8ce4ebb20b79019124eb9870109fdfe97fe9da99f6c20c07ceb1c60b3-image.png" />
    // 输入：root = [47, 74, 31]
    // 输出：121
    // 解释：根节点的左右节点可以并行执行31分钟，剩下的43+47分钟只能串行执行，因此总体执行时间是121分钟。
    tree1 := buildTree([]*int{ptr(47), ptr(74), ptr(31)})
    printTree(tree1)
    fmt.Println(minimalExecTime(tree1)) // 121  (47 + 74)
    // 示例 2：
    //             15
    //            /
    //          12
    //         /
    //        24
    //       /   \
    //      27   26
    // <img src="https://pic.leetcode-cn.com/13accf172ee4a660d241e25901595d55b759380b090890a17e6e7bd51a143e3f-image.png" />
    // 输入：root = [15, 21, null, 24, null, 27, 26]
    // 输出：87
    tree2 := buildTree([]*int{ptr(15), ptr(21), nil, ptr(24), nil, ptr(27), ptr(26)})
    printTree(tree2)
    fmt.Println(minimalExecTime(tree2)) // 87 (15 + 12 + 24 + 27)
    // 示例 3：
    //         1
    //       /   \
    //      3     2
    //          /    \
    //         4      4
    // <img src="https://pic.leetcode-cn.com/bef743a12591aafb9047dd95d335b8083dfa66e8fdedc63f50fd406b4a9d163a-image.png" />
    // 输入：root = [1,3,2,null,null,4,4]
    // 输出：7.5
    tree3 := buildTree([]*int{ptr(1), ptr(3), ptr(2), nil, nil, ptr(4), ptr(4)})
    printTree(tree3)
    fmt.Println(minimalExecTime(tree3)) // 7.5
}