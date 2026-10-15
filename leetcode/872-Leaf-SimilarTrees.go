package main 

// 872. Leaf-Similar Trees
// Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/16/tree.png" />
// For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).
// Two binary trees are considered leaf-similar if their leaf value sequence is the same.
// Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

// Example 1:
//             3                           3
//           /   \                      /    \
//          5     1                    5      1
//        /  \   /  \                /  \    /  \
//      (6)  2 (9)  (8)            (6)  (7) (4)  2
//          / \                                 / \
//        (7)  (4)                            (9)  (8)
// <img src="https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-1.jpg" />
// Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-2.jpg" />
//         1          1
//        /  \       /  \
//       2    3     3    2
// Input: root1 = [1,2,3], root2 = [1,3,2]
// Output: false
 
// Constraints:
//     The number of nodes in each tree will be in the range [1, 200].
//     Both of the given trees will have values in the range [0, 200].

import "fmt"
import "reflect"

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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    list1, list2 := []int{}, []int{}
    var recursion func(node *TreeNode, leafList *[]int)
    recursion = func(node *TreeNode, leafList *[]int) { // 取出叶子节点列表
        if node == nil {
            return
        }
        if node.Left == nil && node.Right == nil {
            *leafList = append(*leafList, node.Val)
            return
        }
        recursion(node.Left, leafList)
        recursion(node.Right, leafList)
    }
    recursion(root1, &list1)
    recursion(root2, &list2)
    return reflect.DeepEqual(list1, list2)
}

func main() {
    // Example 1:
    //             3                           3
    //           /   \                      /    \
    //          5     1                    5      1
    //        /  \   /  \                /  \    /  \
    //      (6)  2 (9)  (8)            (6)  (7) (4)  2
    //          / \                                 / \
    //        (7)  (4)                            (9)  (8)
    // <img src="https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-1.jpg" />
    // Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
    // Output: true
    tree11 := buildTree([]*int{ptr(3), ptr(5), ptr(1), ptr(6), ptr(2), ptr(9), ptr(8), nil, nil, ptr(7), ptr(4)})
    printTree(tree11)
    tree12 := buildTree([]*int{ptr(3), ptr(5), ptr(1), ptr(6), ptr(7), ptr(4), ptr(2), nil, nil, nil, nil, nil, nil, ptr(9), ptr(8)})
    printTree(tree12)
    fmt.Println(leafSimilar(tree11, tree12)) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-2.jpg" />
    //         1          1
    //        /  \       /  \
    //       2    3     3    2
    // Input: root1 = [1,2,3], root2 = [1,3,2]
    // Output: false
    tree21 := buildTree([]*int{ptr(1), ptr(2), ptr(3)})
    printTree(tree21)
    tree22 := buildTree([]*int{ptr(1), ptr(3), ptr(2)})
    printTree(tree22)
    fmt.Println(leafSimilar(tree21, tree22)) // false
}