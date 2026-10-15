package main

// 545. Boundary of Binary Tree
// The boundary of a binary tree is the concatenation of 
// the root, the left boundary, the leaves ordered from left-to-right, and the reverse order of the right boundary.

// The left boundary is the set of nodes defined by the following:
//     The root node's left child is in the left boundary. If the root does not have a left child, then the left boundary is empty.
//     If a node in the left boundary and has a left child, then the left child is in the left boundary.
//     If a node is in the left boundary, has no left child, but has a right child, then the right child is in the left boundary.
//     The leftmost leaf is not in the left boundary.

// The right boundary is similar to the left boundary, except it is the right side of the root's right subtree. 
// Again, the leaf is not part of the right boundary, and the right boundary is empty if the root does not have a right child.

// The leaves are nodes that do not have any children. For this problem, the root is not a leaf.
// Given the root of a binary tree, return the values of its boundary.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/11/boundary1.jpg" />
// Input: root = [1,null,2,3,4]
// Output: [1,3,4,2]
// Explanation:
// - The left boundary is empty because the root does not have a left child.
// - The right boundary follows the path starting from the root's right child 2 -> 4.
//   4 is a leaf, so the right boundary is [2].
// - The leaves from left to right are [3,4].
// Concatenating everything results in [1] + [] + [3,4] + [2] = [1,3,4,2].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/11/boundary2.jpg" />
// Input: root = [1,2,3,4,5,6,null,null,null,7,8,9,10]
// Output: [1,2,4,7,8,9,10,6,3]
// Explanation:
// - The left boundary follows the path starting from the root's left child 2 -> 4.
//   4 is a leaf, so the left boundary is [2].
// - The right boundary follows the path starting from the root's right child 3 -> 6 -> 10.
//   10 is a leaf, so the right boundary is [3,6], and in reverse order is [6,3].
// - The leaves from left to right are [4,7,8,9,10].
// Concatenating everything results in [1] + [2] + [4,7,8,9,10] + [6,3] = [1,2,4,7,8,9,10,6,3].

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -1000 <= Node.val <= 1000


import "fmt"

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
func boundaryOfBinaryTree(root *TreeNode) []int {
    //按题意要求，将节点分为几个类别：
    //根节点的左子为左边节点，左边节点的左子（或右子）不是叶子的话就也是左边节点；
    //同理右边界也类似。
    //然后先序遍历，遍历到的顺序就是根->左边->左边叶子->普通节点叶子->右边->右树叶子（可参考示例图）
    //把这几个部分分别存最后根据顺序拼接即可（注意部分结果需要逆序）
    leftB, rightB, normal := 1,2,3
    res, lLeaves, rBound, rLeaves :=  []int{}, []int{}, []int{}, []int{} //有可能是左叶子,右边界,和右叶子
    var traversal func(node *TreeNode, boundaryType int)
    traversal = func(node *TreeNode, boundaryType int) {
        if node == nil {
            return 
        }
        switch boundaryType {
            case leftB:
                if !(node.Left == nil && node.Right == nil) {
                    res = append(res, node.Val)
                    if node.Left!=nil {
                        traversal(node.Left, leftB)
                        traversal(node.Right, normal)
                    } else {
                        traversal(node.Right, leftB)
                    }
                } else {
                    lLeaves = append(lLeaves, node.Val)
                }
            case rightB: 
                if !(node.Left == nil && node.Right == nil) {
                    rBound = append(rBound, node.Val)
                    if node.Right!=nil {
                        traversal(node.Left, normal)
                        traversal(node.Right, rightB)
                    } else {
                        traversal(node.Left, rightB)
                    }
                } else {
                    rLeaves = append(rLeaves, node.Val)
                }
            case normal: 
                if node.Left == nil && node.Right == nil {
                    lLeaves = append(lLeaves, node.Val)
                } else {
                    traversal(node.Left, normal)
                    traversal(node.Right, normal)
                }
        }
    }
    res = append(res, root.Val) // 添加根节点
    if root.Left != nil {
        traversal(root.Left, leftB)
    }
    if root.Right != nil {
        traversal(root.Right, rightB)
    }
    for i:=0; i < len(lLeaves); i++ { // 添加 lLeaves
        res = append(res, lLeaves[i])
    }
    for i:=len(rLeaves)-1; i>=0; i-- { // 添加 rLeaves
        res = append(res, rLeaves[i])
    }
    for i := len(rBound) - 1; i >= 0; i-- { // 添加右边界
        res = append(res, rBound[i])
    }
    return res
}

func boundaryOfBinaryTree1(root *TreeNode) []int {
    if root == nil {
        return []int{0}
    }
    isLeafNode := func(node *TreeNode) bool {
        return node.Left == nil && node.Right == nil
    }
    if isLeafNode(root) {
        return []int{ root.Val }
    }
    res := []int{ root.Val }
    left, right :=make([]int,0),make([]int,0)
    var getLeftBoundary func(root *TreeNode,res []int) []int
    getLeftBoundary = func(root *TreeNode,res []int) []int {
        if root == nil{ 
            return res
        }
        if !isLeafNode(root) {
            res = append(res,root.Val)
        }
        if root.Left != nil {
            return getLeftBoundary(root.Left,res)
        } else if root.Right != nil {
            return getLeftBoundary(root.Right,res)
        }
        return res
    }
    var getRightBoundary func(root *TreeNode,res []int) []int
    getRightBoundary = func(root *TreeNode,res []int) []int {
        if root == nil {
            return res
        }
        if !isLeafNode(root) {
            res = append(res,root.Val)
        }
        if root.Right != nil {
            return getRightBoundary(root.Right,res)
        } else if root.Left != nil {
            return getRightBoundary(root.Left,res)
        }
        return res
    }
    getLeafNode := func(root *TreeNode) []int {
        res := make([]int,0)
        var dfs func(* TreeNode)
        dfs = func(node *TreeNode) {
            if node == nil {
                return
            }
            dfs(node.Left)
            if isLeafNode(node) {
                res = append(res,node.Val)
            }
            dfs(node.Right)
       }
       dfs(root)
       return res
    }
    reverse := func(src []int ) []int {
        res := make([]int,0)
        for i := len(src)-1; i >= 0; i-- {
            res = append(res, src[i])
        }
        return res
    }
    left = getLeftBoundary(root.Left,left)
    leaf := getLeafNode(root)
    right = getRightBoundary(root.Right, right)
    right = reverse(right)
    res = append(res,left...)
    res = append(res,leaf...)
    res = append(res,right...)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/11/11/boundary1.jpg" />
    // Input: root = [1,null,2,3,4]
    // Output: [1,3,4,2]
    // Explanation:
    // - The left boundary is empty because the root does not have a left child.
    // - The right boundary follows the path starting from the root's right child 2 -> 4.
    //   4 is a leaf, so the right boundary is [2].
    // - The leaves from left to right are [3,4].
    // Concatenating everything results in [1] + [] + [3,4] + [2] = [1,3,4,2].
    tree1 := buildTree([]*int{ptr(1), nil, ptr(2), ptr(3), ptr(4)})
    printTree(tree1)
    fmt.Println(boundaryOfBinaryTree(tree1)) // [1,3,4,2]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/11/11/boundary2.jpg" />
    // Input: root = [1,2,3,4,5,6,null,null,null,7,8,9,10]
    // Output: [1,2,4,7,8,9,10,6,3]
    // Explanation:
    // - The left boundary follows the path starting from the root's left child 2 -> 4.
    //   4 is a leaf, so the left boundary is [2].
    // - The right boundary follows the path starting from the root's right child 3 -> 6 -> 10.
    //   10 is a leaf, so the right boundary is [3,6], and in reverse order is [6,3].
    // - The leaves from left to right are [4,7,8,9,10].
    // Concatenating everything results in [1] + [2] + [4,7,8,9,10] + [6,3] = [1,2,4,7,8,9,10,6,3].
    tree2 := buildTree([]*int{ptr(1), ptr(2), ptr(4), ptr(7), ptr(8), ptr(9), ptr(10), ptr(6), ptr(3)})
    printTree(tree2)
    fmt.Println(boundaryOfBinaryTree(tree2)) // [1,2,4,7,8,9,10,6,3]

    fmt.Println(boundaryOfBinaryTree1(tree1)) // [1,3,4,2]
    fmt.Println(boundaryOfBinaryTree1(tree2)) // [1,2,4,7,8,9,10,6,3]
}