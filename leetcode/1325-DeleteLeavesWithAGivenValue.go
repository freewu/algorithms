package main

// 1325. Delete Leaves With a Given Value
// Given a binary tree root and an integer target, delete all the leaf nodes with value target.
// Note that once you delete a leaf node with value target, 
// if its parent node becomes a leaf node and has the value target, 
// it should also be deleted (you need to continue doing that until you cannot).

// Example 1:
//        1                   1           1
//      /    \               /  \          \
//     2      3      =>    (2)   3   =>     3
//    /     /   \                 \          \
//  (2)    (2)  (4)                4          4
// <img src="https://assets.leetcode.com/uploads/2020/01/09/sample_1_1684.png" />
// Input: root = [1,2,3,2,null,2,4], target = 2
// Output: [1,null,3,null,4]
// Explanation: Leaf nodes in green with value (target = 2) are removed (Picture in left). 
// After removing, new nodes become leaf nodes with value (target = 2) (Picture in center).

// Example 2:
//        1                 1
//      /   \              /
//     3    (3)     =>    3
//    /  \                 \ 
//  (3)   2                  2
// <img src="https://assets.leetcode.com/uploads/2020/01/09/sample_2_1684.png" />
// Input: root = [1,3,3,3,2], target = 3
// Output: [1,3,null,null,2]

// Example 3:
//        1         1         1       1
//       /         /         /
//      2   =>    2   =>   (2)  =>  
//     /         /
//    2        (2)
//   /
// (2)
// <img src="https://assets.leetcode.com/uploads/2020/01/15/sample_3_1684.png" />
// Input: root = [1,2,null,2,null,2], target = 2
// Output: [1]
// Explanation: Leaf nodes in green with value (target = 2) are removed at each step.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 3000].
//     1 <= Node.val, target <= 1000

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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    var dfsDelete func(root, node *TreeNode, l, target int) bool
    dfsDelete = func(root, node *TreeNode, l, target int) bool{
        if node == nil { return false; }
        if node.Left == nil && node.Right == nil && node.Val == target { // 如果是叶子（node.Left == nil && node.Right == nil）且是删除目标 做删除操作
            if l == 0 { root.Left = nil; } else { root.Right = nil; }
            return true
        }
        return dfsDelete(node, node.Left, 0, target) || dfsDelete(node, node.Right, 1, target)
    }
    for {
        if dfsDelete(root, root.Left, 0, target) || dfsDelete(root, root.Right, 1, target) { 
            continue
        } else {
            break
        }
    }
    if root.Left == nil && root.Right == nil { // 处理整棵树都是 target 的情况
        if root.Val == target { return nil; }  // 如果根节点也是要删除的值
    }
    return root
}

func removeLeafNodes1(root *TreeNode, target int) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left = removeLeafNodes1(root.Left, target)
    root.Right = removeLeafNodes1(root.Right, target)
    if root.Val == target && root.Left == nil && root.Right == nil {
        return nil
    }
    return root
}

func main() {
    // Example 1:
    //        1                   1           1
    //      /    \               /  \          \
    //     2      3      =>    (2)   3   =>     3
    //    /     /   \                 \          \
    //  (2)    (2)  (4)                4          4
    // <img src="https://assets.leetcode.com/uploads/2020/01/09/sample_1_1684.png" />
    // Input: root = [1,2,3,2,null,2,4], target = 2
    // Output: [1,null,3,null,4]
    // Explanation: Leaf nodes in green with value (target = 2) are removed (Picture in left). 
    // After removing, new nodes become leaf nodes with value (target = 2) (Picture in center).
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(2), nil, ptr(2), ptr(4)})
    printTree(tree1)
    fmt.Println("tree1.Left: ",tree1.Left)
    t1 := removeLeafNodes(tree1, 2)
    printTree(t1)
    fmt.Println("t1.Left: ",t1.Left)
    // Example 2:
    //        1                 1
    //      /   \              /
    //     3    (3)     =>    3
    //    /  \                 \ 
    //  (3)   2                  2
    // <img src="https://assets.leetcode.com/uploads/2020/01/09/sample_2_1684.png" />
    // Input: root = [1,3,3,3,2], target = 3
    // Output: [1,3,null,null,2]
    tree2 := buildTree([]*int{ptr(1), ptr(3), ptr(3), ptr(3), ptr(3)})
    printTree(tree2)
    fmt.Println("tree2.Left: ", tree2.Left)
    t2 := removeLeafNodes(tree2, 3)
    printTree(t2)
    fmt.Println("t2.Left: ",t2.Left)
    // Example 3:
    //        1         1         1       1
    //       /         /         /
    //      2   =>    2   =>   (2)  =>  
    //     /         /
    //    2        (2)
    //   /
    // (2)
    // <img src="https://assets.leetcode.com/uploads/2020/01/15/sample_3_1684.png" />
    // Input: root = [1,2,null,2,null,2], target = 2
    // Output: [1]
    // Explanation: Leaf nodes in green with value (target = 2) are removed at each step.
    tree3 := buildTree([]*int{ptr(1), ptr(2), nil, ptr(2), nil, ptr(2)})
    printTree(tree3)
    fmt.Println("tree3.Left: ",tree3.Left)
    t3 := removeLeafNodes(tree3, 2)
    printTree(t3)
    fmt.Println("t3.Left: ",t3.Left)


    tree11 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(2), nil, ptr(2), ptr(4)})
    printTree(tree11)
    fmt.Println("tree11.Left: ",tree11.Left)
    t11 := removeLeafNodes(tree11, 2)
    printTree(t11)
    fmt.Println("t11.Left: ",t11.Left)
    tree12 := buildTree([]*int{ptr(1), ptr(3), ptr(3), ptr(3), ptr(3)})
    printTree(tree12)
    fmt.Println("tree12.Left: ", tree12.Left)
    t12 := removeLeafNodes(tree12, 3)
    printTree(t12)
    fmt.Println("t12.Left: ",t12.Left)
    tree13 := buildTree([]*int{ptr(1), ptr(2), nil, ptr(2), nil, ptr(2)})
    printTree(tree13)
    fmt.Println("tree13.Left: ",tree13.Left)
    t13 := removeLeafNodes(tree13, 2)
    printTree(t13)
    fmt.Println("t13.Left: ",t13.Left)
}