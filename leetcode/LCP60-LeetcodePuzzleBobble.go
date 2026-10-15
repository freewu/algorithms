package main

// LCP 60. 力扣泡泡龙
// 欢迎各位勇者来到力扣城，本次试炼主题为「力扣泡泡龙」。

// 游戏初始状态的泡泡形如二叉树 root，每个节点值对应了该泡泡的分值。
// 勇者最多可以击破一个节点泡泡，要求满足：
//     1. 被击破的节点泡泡 至多 只有一个子节点泡泡
//     2. 当被击破的节点泡泡有子节点泡泡时，则子节点泡泡将取代被击破泡泡的位置

// 注：即整棵子树泡泡上移

// 请问在击破一个节点泡泡操作或无击破操作后，二叉泡泡树的最大「层和」是多少。

// 注意：
//     「层和」为同一高度的所有节点的分值之和

// 示例 1：
// 输入：root = [6,0,3,null,8]
// 输出：11
// 解释：勇者的最佳方案如图所示
// <img src="https://pic.leetcode-cn.com/1648180809-XSWPLu-image.png" />

// 示例 2：
// 输入：root = [5,6,2,4,null,null,1,3,5]
// 输出：9
// 解释：勇者击破 6 节点，此时「层和」最大为 3+5+1 = 9
// <img src="https://pic.leetcode-cn.com/1648180769-TLpYop-image.png" />

// 示例 3：
// 输入：root = [-5,1,7]
// 输出：8
// 解释：勇者不击破节点，「层和」最大为 1+7 = 8

// 提示：
//     2 <= 树中节点个数 <= 10^5
//     -10000 <= 树中节点的值 <= 10000

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
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
func getMaxLayerSum(root *TreeNode) int {
    type LevelInfo struct { PrefixSum, IsOccupied, LeftPtr, RightPtr int }   // 用于存储每层的信息，包括前缀和、是否被占用以及左、右子节点的指针
    type RemoveNode struct { NodeIndex, Level int } // 用于存储需要被移除的节点信息
    levelInfos := make([][]LevelInfo, 0) // 用于存储每一层的信息
    removeNodes := make([]RemoveNode, 0) // 用于存储需要被移除的节点信息
    // collectNodes 执行DFS，并收集树的每一层的信息和需要被移除的节点。
    var collectNodes func(level int, node *TreeNode) int
    collectNodes = func(level int, node *TreeNode) int {
        if node == nil { return 0 }
        // 扩展 levelInfos 列表以适应新层。
        for len(levelInfos) <= level+1 {
            levelInfos = append(levelInfos, []LevelInfo{{ PrefixSum: 0, IsOccupied: -1, LeftPtr: -1, RightPtr: -1 }})
        }
        // 更新当前层的信息。
        lastInfo := levelInfos[level][len(levelInfos[level]) - 1]
        newInfo := LevelInfo{
            PrefixSum:  lastInfo.PrefixSum + node.Val, // 计算前缀和
            IsOccupied: -1,                            // 初始化为未占用
            LeftPtr:    len(levelInfos[level + 1]),      // 设置左子节点的指针
            RightPtr:   -1,                            // 初始化右子节点的指针为-1
        }
        levelInfos[level] = append(levelInfos[level], newInfo)
        node.Val = len(levelInfos[level]) - 1 // 用于识别节点的索引
        // 如果节点有两个子节点，则添加到 removeNodes 列表中。
        if collectNodes(level+1, node.Left) + collectNodes(level+1, node.Right) != 2 {
            removeNodes = append(removeNodes, RemoveNode{NodeIndex: len(levelInfos[level]) - 1, Level: level})
        }
        // 更新右子节点的指针。
        levelInfos[level][len(levelInfos[level])-1].RightPtr = len(levelInfos[level+1]) - 1
        return 1
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    collectNodes(0, root)
    res, treeHeight := -1 << 31, len(levelInfos) - 1 // 获取树的高度
    for level := 0; level < treeHeight; level++ { // 计算没有移除任何节点时的最大层和
        res = max(res, levelInfos[level][len(levelInfos[level]) - 1].PrefixSum)
    }
    for _, removeInfo := range removeNodes { // 遍历 removeNodes 列表，并尝试移除每一个节点以计算新的最大层和
        nodeIndex, startLevel := removeInfo.NodeIndex, removeInfo.Level
        left, right := nodeIndex, nodeIndex
        // 计算被移除节点的值。
        lostVal := levelInfos[startLevel][left].PrefixSum - levelInfos[startLevel][left - 1].PrefixSum
        for level := startLevel; level < treeHeight; level++ {
            if left > right { break }
            leftInfo, rightInfo := &levelInfos[level][left], &levelInfos[level][right]
            if leftInfo.IsOccupied != -1 && leftInfo.IsOccupied == rightInfo.IsOccupied { break } // 如果节点已被标记为占用，则跳过
            leftInfo.IsOccupied, rightInfo.IsOccupied = nodeIndex, nodeIndex // 标记节点为占用
            addVal := 0
            if leftInfo.LeftPtr <= rightInfo.RightPtr { // 计算被移除节点的子节点和。
                addVal = levelInfos[level+1][rightInfo.RightPtr].PrefixSum - levelInfos[level+1][leftInfo.LeftPtr-1].PrefixSum
            }
            newSum := levelInfos[level][len(levelInfos[level])-1].PrefixSum - lostVal + addVal // 计算新的层和
            if newSum != 0 || (right-left + 1 != len(levelInfos[level]) - 1) { // 如果新的层和不为0，或者该层的节点数不等于总节点数，则更新最大层和
                res = max(res, newSum)
            }
            // 更新 left 和 right 指针，以及需要减去的值。
            left, right, lostVal = leftInfo.LeftPtr, rightInfo.RightPtr, addVal
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：root = [6,0,3,null,8]
    // 输出：11
    // 解释：勇者的最佳方案如图所示
    // <img src="https://pic.leetcode-cn.com/1648180809-XSWPLu-image.png" />
    tree1 := buildTree([]*int{ptr(6), ptr(0), ptr(3), nil, ptr(8)})
    printTree(tree1)
    fmt.Println(getMaxLayerSum(tree1)) // 11
    // 示例 2：
    // 输入：root = [5,6,2,4,null,null,1,3,5]
    // 输出：9
    // 解释：勇者击破 6 节点，此时「层和」最大为 3+5+1 = 9
    // <img src="https://pic.leetcode-cn.com/1648180769-TLpYop-image.png" />
    tree2 := buildTree([]*int{ptr(5), ptr(6), ptr(2), ptr(4), nil, nil, ptr(1), ptr(3), ptr(5)})
    printTree(tree2)
    fmt.Println(getMaxLayerSum(tree2)) // 9
    // 示例 3：
    // 输入：root = [-5,1,7]
    // 输出：8
    // 解释：勇者不击破节点，「层和」最大为 1+7 = 8
    tree3 := buildTree([]*int{ptr(-5), ptr(1), ptr(7)})
    printTree(tree3)
    fmt.Println(getMaxLayerSum(tree3)) // 8
}