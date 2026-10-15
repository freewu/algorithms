package main

// 285. Inorder Successor in BST
// Given the root of a binary search tree and a node p in it, return the in-order successor of that node in the BST. 
// If the given node has no in-order successor in the tree, return null.
// The successor of a node p is the node with the smallest key greater than p.val.

// Example 1:
//         2
//       /   \
//      1     3
// <img src="https://assets.leetcode.com/uploads/2019/01/23/285_example_1.PNG" />
// Input: root = [2,1,3], p = 1
// Output: 2
// Explanation: 1's in-order successor node is 2. Note that both p and the return value is of TreeNode type.

// Example 2:
//           5
//         /   \
//        3     6
//      /   \
//     2     4
//    /
//   1
// <img src="https://assets.leetcode.com/uploads/2019/01/23/285_example_2.PNG" />
// Input: root = [5,3,6,2,4,null,null,1], p = 6
// Output: null
// Explanation: There is no in-order successor of the current node, so the answer is null.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -10^5 <= Node.val <= 10^5
//     All Nodes will have unique values.

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
// 中序遍历
// 如果节点 p 是最后被访问的节点，则不存在节点 p 的中序后继，返回 nil
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    stack := []*TreeNode{}
    pre, cur := &TreeNode{}, root
    for len(stack) > 0 || cur != nil {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }
        cur = stack[len(stack)-1] // pop
        stack = stack[:len(stack)-1]
        if pre == p {
            return cur
        }
        pre = cur
        cur = cur.Right
    }
    return nil
}

// 二叉搜索树的一个性质是中序遍历序列单调递增，因此二叉搜索树中的节点 ppp 的中序后继满足以下条件：
//     中序后继的节点值大于 p 的节点值；
//     中序后继是节点值大于 p 的节点值的所有节点中节点值最小的一个节点。
// 利用二叉搜索树的性质，可以在不做中序遍历的情况下找到节点 ppp 的中序后继。
func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
    var successor *TreeNode
    if p.Right != nil {
        successor = p.Right
        for successor.Left != nil {
            successor = successor.Left
        }
        return successor
    }
    node := root
    for node != nil {
        if node.Val > p.Val {
            successor = node
            node = node.Left
        } else {
            node = node.Right
        }
    }
    return successor
}

func main() {
    // Example 1:
    //         2
    //       /   \
    //      1     3
    // <img src="https://assets.leetcode.com/uploads/2019/01/23/285_example_1.PNG" />
    // Input: root = [2,1,3], p = 1
    // Output: 2
    // Explanation: 1's in-order successor node is 2. Note that both p and the return value is of TreeNode type.
    tree1 := buildTree([]*int{ptr(2), ptr(1), ptr(3)})
    printTree(tree1)
    p1 := tree1.Left
    t1 := inorderSuccessor(tree1, p1)
    printTree(t1)
    fmt.Println("t1: ", t1) 
    fmt.Println("t1.Val: ", t1.Val) // 2 
    // Example 2:
    //           5
    //         /   \
    //        3     6
    //      /   \
    //     2     4
    //    /
    //   1
    // <img src="https://assets.leetcode.com/uploads/2019/01/23/285_example_2.PNG" />
    // Input: root = [5,3,6,2,4,null,null,1], p = 6
    // Output: null
    // Explanation: There is no in-order successor of the current node, so the answer is null.
    tree2 := buildTree([]*int{ptr(5), ptr(3), ptr(6), ptr(2), ptr(4), nil, nil, ptr(1)})
    printTree(tree2)
    p2 := tree2.Right
    t2 := inorderSuccessor(tree2, p2)
    printTree(t2)
    fmt.Println("t2: ",t2)

    tree11 := buildTree([]*int{ptr(2), ptr(1), ptr(3)})
    printTree(tree11)
    p11 := tree11.Left
    t11 := inorderSuccessor1(tree11, p11)
    printTree(t11)
    fmt.Println("t11: ", t11) 
    fmt.Println("t11.Val: ", t11.Val) // 2 

    tree12 := buildTree([]*int{ptr(5), ptr(3), ptr(6), ptr(2), ptr(4), nil, nil, ptr(1)})
    printTree(tree12)
    p12 := tree12.Right
    t12 := inorderSuccessor1(tree12, p12)
    printTree(t12)
    fmt.Println("t12: ",t12) 
}