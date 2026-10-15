package main

// 863. All Nodes Distance K in Binary Tree
// Given the root of a binary tree, the value of a target node target, and an integer k, 
// return an array of the values of all nodes that have a distance k from the target node.

// You can return the answer in any order.

// Example 1:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/28/sketch0.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
// Output: [7,4,1]
// Explanation: The nodes that are a distance 2 from the target node (with value 5) have values 7, 4, and 1.

// Example 2:
// Input: root = [1], target = 1, k = 3
// Output: []
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 500].
//     0 <= Node.val <= 500
//     All the values Node.val are unique.
//     target is the value of one of the nodes in the tree.
//     0 <= k <= 1000

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
// 并查集
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    res, mp := []int{}, make(map[*TreeNode]int)
    var find func(root, target *TreeNode)
    find = func(root, target *TreeNode) {
        if root == nil { return }
        if root == target {
            mp[root] = 0
            return
        }
        find(root.Left, target)
        if val, ok := mp[root.Left]; ok {
            mp[root] = val + 1
            return
        }
        find(root.Right, target)
        if val, ok := mp[root.Right]; ok {
            mp[root] = val + 1
            return
        }
    }
    var search func(root *TreeNode, distance int)
    search = func(root *TreeNode, distance int) {
        if root == nil { return }
        if val, ok := mp[root]; ok {
            distance = val
        }
        if distance == k {
            res = append(res, root.Val)
        }
        search(root.Left, distance+1)
        search(root.Right, distance+1)
    }
    find(root, target)
    search(root, 0)
    return res
}

// dfs
func distanceK1(root *TreeNode, target *TreeNode, k int) []int {
    res := []int{}
    var dfs func(node *TreeNode, tmp int)
    dfs = func(node *TreeNode, tmp int) {
        if node == nil { return }
        if tmp == k {
            res = append(res, node.Val)
            return
        }
        dfs(node.Left, tmp+1)
        dfs(node.Right, tmp+1)
    }
    var dfs2 func(node *TreeNode) int 
    dfs2 = func(node *TreeNode) int {
        if node == nil { return -1 }
        if node == target {
            dfs(node, 0)
            return 1
        } else {
            l, r := dfs2(node.Left), dfs2(node.Right)
            if l > 0 {
                if l == k {
                    res = append(res, node.Val)
                }
                dfs(node.Right, l+1)
                return l+1
            }
            if r > 0 {
                if r == k {
                    res = append(res, node.Val)
                }
                dfs(node.Left, r+1)
                return r+1
            }
        }
        return -1
    }
    dfs2(root)
    return res
}

func main() {
    // Example 1:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/06/28/sketch0.png" />
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
    // Output: [7,4,1]
    // Explanation: The nodes that are a distance 2 from the target node (with value 5) have values 7, 4, and 1.
    tree1 := buildTree([]*int{ptr(3), ptr(5), ptr(1), ptr(6), ptr(2), ptr(0), ptr(8), nil, nil, ptr(7), ptr(4)})
    printTree(tree1)
    fmt.Println(distanceK(tree1,&TreeNode{5, nil, nil, }, 2)) // [7,4,1]
    // Example 2:
    // Input: root = [1], target = 1, k = 3
    // Output: []
    tree2 := buildTree([]*int{ptr(1)})
    printTree(tree2)
    fmt.Println(distanceK(tree2, &TreeNode{1, nil, nil, }, 3)) // []

    tree11 := buildTree([]*int{ptr(3), ptr(5), ptr(1), ptr(6), ptr(2), ptr(0), ptr(8), nil, nil, ptr(7), ptr(4)})
    printTree(tree11)
    fmt.Println(distanceK1(tree11, &TreeNode{5, nil, nil, }, 2)) // [7,4,1]
    tree12 := buildTree([]*int{ptr(1)})
    printTree(tree12)
    fmt.Println(distanceK1(tree12, &TreeNode{1, nil, nil, }, 3)) // []
}