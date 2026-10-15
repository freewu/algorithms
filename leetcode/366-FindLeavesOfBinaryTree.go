package main 

// 366. Find Leaves of Binary Tree
// Given the root of a binary tree, collect a tree's nodes as if you were doing this:
//     Collect all the leaf nodes.
//     Remove all the leaf nodes.
//     Repeat until the tree is empty.

// Example 1:
//         1                 1           (1)
//       /   \              /
//      2    (3)    =>    (2)      => 
//     /  \
//   (4)  (5)
//    [4, 5, 3]             [2]           [1]
// <img src="https://assets.leetcode.com/uploads/2021/03/16/remleaves-tree.jpg" />
// Input: root = [1,2,3,4,5]
// Output: [[4,5,3],[2],[1]]
// Explanation:
// [[3,5,4],[2],[1]] and [[3,4,5],[2],[1]] are also considered correct answers since per each level it does not matter the order on which elements are returned.

// Example 2:
// Input: root = [1]
// Output: [[1]]

// Constraints:
//     The number of nodes in the tree is in the range [1, 100].
//     -100 <= Node.val <= 100


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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findLeaves(root *TreeNode) [][]int {
    // 最先被删除叶子的节点也是下一轮最先被删除的叶子候选者（为什么说是候选，因为一个节点可以有一个叶子孩子和一个非叶子孩子，那么它的叶子孩子被删除后自己并不会立即变成叶子。但是只有叶子孩子的节点肯定是下一批要删除的。
    // 而叶子的删除顺序其实就是先序/中序遍历时遇到叶子的顺序。
    // 所以可以维护一个子节点指向父节点的映射，这样，先遍历一遍即可得到当前要删除的叶子，并且可以在删除后立即修改并检查父亲节点的状态，从而确定下一批要删除的叶子。
    res := [][]int{}
    if root == nil { return res }
    type Ref struct { // 为方便得知当前叶节点是其父亲的哪个孩子，这里定义一个结构体
        parent *TreeNode // 父节点
        leftOrRight bool // 是父节点的左子还是右子
    }
    isLeaves := func(node *TreeNode) bool { // 判断是否是叶子节点
        return node.Left == nil && node.Right == nil 
    }
    leaves,toParent := []*TreeNode{}, map[*TreeNode]Ref{} // 维护一个指向父亲的 map
    var traversal func(node *TreeNode)
    traversal = func(node *TreeNode) { 
        if isLeaves(node) { // 是叶子节点,把节点加入到数组,跳出递归
            leaves = append(leaves, node)
            return 
        }
        if node.Left != nil {
            toParent[node.Left] = Ref{ parent: node, leftOrRight: true}
            traversal(node.Left)
        }
        if node.Right != nil {
            toParent[node.Right] = Ref{ parent: node, leftOrRight: false}
            traversal(node.Right)
        }
    }
    traversal(root) // 1. 遍历，构造toParent映射和第一批叶子
    for len(leaves) > 0 { // 删除leaves直到为空
        nextLeaves := []*TreeNode{} // 下一批叶子
        curVals := make([]int, len(leaves)) //本批次叶子值
        for i, leaf := range leaves {
            curVals[i] = leaf.Val
            ref, ok := toParent[leaf]
            if !ok { // 找不到父节点说明是根了，处理完成
                break
            }
            if ref.leftOrRight { // 删除的是左子叶
                ref.parent.Left = nil
            } else {
                ref.parent.Right = nil
            }
            if isLeaves(ref.parent) {// 父节点变成叶子了
                nextLeaves = append(nextLeaves, ref.parent)
            }
        }
        res = append(res, curVals)
        leaves = nextLeaves
    }
    return res
}

func findLeaves1(root *TreeNode) [][]int {
    res := [][]int{}
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(cur *TreeNode) int
    dfs = func(cur *TreeNode) int {
        if cur == nil {
            return 0
        }
        left, right := dfs(cur.Left), dfs(cur.Right)
        level := max(left, right) + 1
        if level - 1 == len(res) {
            res = append(res, []int{ cur.Val })
        } else {
            res[level-1] = append(res[level-1], cur.Val)
        }
        return level
    }
    dfs(root)
    return res
}

func main() {
    // Example 1:
    //         1                 1           (1)
    //       /   \              /
    //      2    (3)    =>    (2)      => 
    //     /  \
    //   (4)  (5)
    //    [4, 5, 3]             [2]           [1]
    // <img src="https://assets.leetcode.com/uploads/2021/03/16/remleaves-tree.jpg" />
    // Input: root = [1,2,3,4,5]
    // Output: [[4,5,3],[2],[1]]
    // Explanation:
    // [[3,5,4],[2],[1]] and [[3,4,5],[2],[1]] are also considered correct answers since per each level it does not matter the order on which elements are returned.
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(4), ptr(5)})
    printTree(tree1)
    fmt.Println(findLeaves(tree1)) // [[4,5,3],[2],[1]]
    // Example 2:
    // Input: root = [1]
    // Output: [[1]]
    tree2 := buildTree([]*int{ptr(1)})
    printTree(tree2)
    fmt.Println(findLeaves(tree2)) // [[1]]

    tree11 := buildTree([]*int{ptr(1), ptr(2), ptr(4), ptr(5)})
    printTree(tree11)
    fmt.Println(findLeaves(tree11)) // [[4,5,3],[2],[1]]
    tree12 := buildTree([]*int{ptr(1)})
    printTree(tree12)
    fmt.Println(findLeaves(tree12)) // [[1]]
}