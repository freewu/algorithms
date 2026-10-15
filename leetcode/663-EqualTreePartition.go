package main

// 663. Equal Tree Partition
// Given the root of a binary tree, 
// return true if you can partition the tree into two trees with equal sums of values 
// after removing exactly one edge on the original tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/03/split1-tree.jpg" />
// Input: root = [5,10,10,null,null,2,3]
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/05/03/split2-tree.jpg" />
// Input: root = [1,2,10,null,null,2,20]
// Output: false
// Explanation: You cannot split the tree into two trees with equal sums after removing exactly one edge on the tree.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -10^5 <= Node.val <= 10^5

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

func checkEqualTree(root *TreeNode) bool {
    m := map[int]int{} // 求 root 的和的时候，就可以把中间的和记录下来。所以解决方案就是找到是否有节点和为n/2的。
    var cal func(*TreeNode) int
    cal = func(n *TreeNode) int {
        if nil == n {
            return 0
        }
        v := n.Val + cal(n.Left) + cal(n.Right)
        m[v]++
        return v
    }
    s := cal(root)
    if 0 == s {
        return m[0] > 1
    }
    return s&1 == 0 && m[s >> 1] != 0 // 两颗树相等 -> 一个树的和是 n/2，
}

func checkEqualTree1(root *TreeNode) bool {
    var dfs func(root *TreeNode)int
    mp := make(map[int]bool)
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        v := dfs(node.Left) + dfs(node.Right) + node.Val
        mp[v] = true
        return v
    }
    sum := dfs(root.Left) + dfs(root.Right) + root.Val
    if sum % 2 != 0 {
        return false
    }
    _, ok := mp[sum / 2] // 两颗树相等 -> 一个树的和是 n/2，
    return ok
}

func checkEqualTree2(root *TreeNode) bool {
    // 深度遍历，统计每个子树的和
    // 将所有的和记录到一个 map 中，最终如果存在子树的和等于总和的一半，则返回 true
    mp := make(map[int]int)
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return 0 }
        sum := node.Val + dfs(node.Left) + dfs(node.Right)
        mp[sum]++ // 将所有的和记录到一个 map 中
        return sum
    }
    sum := dfs(root)
    mp[sum]--
    if sum % 2 != 0 { return false }
    return mp[sum / 2] > 0 // 如果存在子树的和等于总和的一半，则返回 true
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/split1-tree.jpg" />
    // Input: root = [5,10,10,null,null,2,3]
    // Output: true
    tree1 := buildTree([]*int{ptr(5), ptr(10), ptr(10), nil, nil, ptr(2), ptr(3)})
    printTree(tree1)
    fmt.Println(checkEqualTree(tree1)) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/split2-tree.jpg" />
    // Input: root = [1,2,10,null,null,2,20]
    // Output: false
    // Explanation: You cannot split the tree into two trees with equal sums after removing exactly one edge on the tree.
    tree2 := buildTree([]*int{ptr(1), ptr(2), ptr(10), nil, nil, ptr(2), ptr(20)})
    printTree(tree2)
    fmt.Println(checkEqualTree(tree2)) // false

    fmt.Println(checkEqualTree1(tree1)) // true
    fmt.Println(checkEqualTree1(tree2)) // false

    fmt.Println(checkEqualTree2(tree1)) // true
    fmt.Println(checkEqualTree2(tree2)) // false
}