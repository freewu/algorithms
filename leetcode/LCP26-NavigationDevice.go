package main

// LCP 26. 导航装置
// 小扣参加的秋日市集景区共有 N 个景点，景点编号为 1~N。
// 景点内设有 N−1 条双向道路，使所有景点形成了一个二叉树结构，根结点记为 root，景点编号即为节点值。

// 由于秋日市集景区的结构特殊，游客很容易迷路，主办方决定在景区的若干个景点设置导航装置，按照所在景点编号升序排列后定义装置编号为 1 ~ M。
// 导航装置向游客发送数据，数据内容为列表 [游客与装置 1 的相对距离,游客与装置 2 的相对距离,...,游客与装置 M 的相对距离]。
// 由于游客根据导航装置发送的信息来确认位置，因此主办方需保证游客在每个景点接收的数据信息皆不相同。
// 请返回主办方最少需要设置多少个导航装置。

// 示例 1：
// 输入：root = [1,2,null,3,4]
// 输出：2
// 解释：在景点 1、3 或景点 1、4 或景点 3、4 设置导航装置。
// <img src="https://pic.leetcode-cn.com/1597996812-tqrgwu-image.png" />

// 示例 2：
// 输入：root = [1,2,3,4]
// 输出：1
// 解释：在景点 3、4 设置导航装置皆可。
// <img src="https://pic.leetcode-cn.com/1597996826-EUQRyz-image.png" />

// 提示：
//     2 <= N <= 50000
//     二叉树的非空节点值为 1~N 的一个排列。

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
func navigation(root *TreeNode) int {
    // 每个三叉点（该节点既有不为空的左右子树，又有父亲节点）的三条支路中，至少两条支路上有导航装置（
    res := 0
    if root == nil { return 0 }
    // 0 代表该点及孩子节点都不会放导航装置
    // 1 代表左右孩子都没有导航装置时，必须选一个孩子支路放一个导航装置，另一个暂时不放，以1状态标记交上面节点去处理
    // 2 表示左右孩子支路都有导航装置，该节点的父节点支路可以放也可以不放导航装置
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return 0 }
        left, right := dfs(node.Left), dfs(node.Right)
        if node.Left != nil && node.Right != nil { // 这里只要左右不为空，即为三叉点
            if left == 0 && right == 0 { // 左右子树都没放导航，那么必须选一条支路放一个导航，另一条支路暂时不放，返回状态1
                res++
                return 1
            }
            if left == 0 || right == 0 { // 一条支路有导航，另一条支路没有导航，继续暂时不放，返回状态1，要不要加交给上面节点来判断处理；
                return 1
            }
            return 2  // 左右孩子支路都有导航，那么就返回状态2
        } else if node.Left == nil { // 左孩子为空，该节点状态等于右孩子
            return right
        } else if node.Right == nil { // 右孩子为空，该节点状态等于左孩子
            return left
        }
        return 0
    }
    left, right := dfs(root.Left), dfs(root.Right)
    //下面这行代码是多种状态综合处理的表现形式，具体如下：
    //左右孩子状态集(左右交换就不写了)：
    //[0,0]：左右都没放导航(包含子树为空的情况)，那么要必须放一个导航，即res+1，而且总共放一个导航；
    //[0,1]：一个孩子没放导航，再和根节点放一起，视为最近的三叉点的一条支路。由于另一个孩子状态为1，表示该三叉点有两条支路未放导航，所以还要再增加一个导航，那么res+1；
    //[0,2]：一个孩子没放导航，再和根节点放一起，视为最近的三叉点的一条支路。另一个孩子状态为2，表示该三叉点有两条支路已放导航，所以不用再增加导航，返回res即可；
    //[1,1]：其中一个孩子和根节点放一起，视为一条支路，那么就表示该合并后的支路上有导航（只要有就可以了），所以该三叉点的两条支路有导航，返回res即可；
    //[1,2]：前面同上，该三叉点的三条支路都有导航，返回res即可；
    //[2,2]：直接同上。
    //综合上面分类情况，左右返回的状态值相加left+right>=2时，返回res即可，其他情况都要再加一个导航。
    if left + right >= 2 { return res }
    return res + 1
}

func main() {
    // 示例 1：
    // 输入：root = [1,2,null,3,4]
    // 输出：2
    // 解释：在景点 1、3 或景点 1、4 或景点 3、4 设置导航装置。
    // <img src="https://pic.leetcode-cn.com/1597996812-tqrgwu-image.png" />
    tree1 := buildTree([]*int{ptr(1), ptr(2), nil, ptr(3), ptr(4)})
    printTree(tree1)
    fmt.Println(navigation(tree1)) // 2
    // 示例 2：
    // 输入：root = [1,2,3,4]
    // 输出：1
    // 解释：在景点 3、4 设置导航装置皆可。
    // <img src="https://pic.leetcode-cn.com/1597996826-EUQRyz-image.png" />
    tree2 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4)})
    printTree(tree2)
    fmt.Println(navigation(tree2)) // 1
}