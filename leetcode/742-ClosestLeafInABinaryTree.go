package main

// 742. Closest Leaf in a Binary Tree
// Given the root of a binary tree where every node has a unique value and a target integer k, 
// return the value of the nearest leaf node to the target k in the tree.

// Nearest to a leaf means the least number of edges traveled on the binary tree to reach any leaf of the tree. 
// Also, a node is called a leaf if it has no children.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/13/closest1-tree.jpg" />
// Input: root = [1,3,2], k = 1
// Output: 2
// Explanation: Either 2 or 3 is the nearest leaf node to the target of 1.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/13/closest2-tree.jpg" />
// Input: root = [1], k = 1
// Output: 1
// Explanation: The nearest leaf node is the root node itself.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/06/13/closest3-tree.jpg" />
// Input: root = [1,2,3,4,null,null,null,5,null,6], k = 2
// Output: 3
// Explanation: The leaf node with value 3 (and not the leaf node with value 6) is nearest to the node with value 2.

// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     1 <= Node.val <= 1000
//     All the values of the tree are unique.
//     There exist some node in the tree where Node.val == k.

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
func findClosestLeaf(root *TreeNode, k int) int {
    res, deep, pre := 0, 1001, []*TreeNode{}
    var findK func(*TreeNode, int)
    findK = func(root *TreeNode, k int) {
        var pLast *TreeNode
        nd := root
        for nd != nil {
            pre = append(pre, nd)
            nd = nd.Left
        }
        for len(pre) > 0 {
            nd = pre[len(pre)-1]
            if nd.Right == nil || nd.Right == pLast {
                if nd.Val == k { return }
                pLast = nd
                pre = pre[:len(pre)-1]
                
            } else {
                nd = nd.Right
                for nd != nil {
                    pre = append(pre, nd)
                    nd = nd.Left
                }
                
            }
        }       
    }
    // find k node and its ancestor
    findK(root, k)
    var findTarget func(*TreeNode, int)
    findTarget = func(nd *TreeNode, level int) {
        if nd == nil { return }
        if nd.Left == nil && nd.Right == nil && level < deep {
            res = nd.Val
            deep = level
        }
        findTarget(nd.Left, level+1)
        findTarget(nd.Right, level+1)
    }
    for i := len(pre)-1; i >= 0; i-- {
        findTarget(pre[i], len(pre)-1-i)
    }
    return res
}


func findClosestLeaf1(root *TreeNode, k int) int {
    parentMap := make(map[*TreeNode]*TreeNode) // 找到目标节点并构建所需的子节点到父节点的映射
    // 遍历函数，找到值为 k 的那个节点，同时记录子节点到父节点的映射
    var traverse func(root *TreeNode, k int, parent *TreeNode, parentMap map[*TreeNode]*TreeNode) *TreeNode
    traverse = func (root *TreeNode, k int, parent *TreeNode, parentMap map[*TreeNode]*TreeNode) *TreeNode {
        if root == nil { return nil  }
        parentMap[root] = parent // 记录子节点到父节点的映射
        if root.Val == k { return root }
        left := traverse(root.Left, k, root, parentMap)
        if left != nil { return left }
        return traverse(root.Right, k, root, parentMap)
    }
    var bfs func(target *TreeNode, parentMap map[*TreeNode]*TreeNode) int 
    bfs = func(target *TreeNode, parentMap map[*TreeNode]*TreeNode) int {
        queue, visited := []*TreeNode{target}, make(map[*TreeNode]bool)
        visited[target] = true
        for len(queue) > 0 {
            size := len(queue)
            for i := 0; i < size; i++ {
                cur := queue[0]
                queue = queue[1:]
                if cur.Left == nil && cur.Right == nil { // 首次到达的叶子结点就是最近的叶子结点
                    return cur.Val
                }
                if cur.Left != nil && !visited[cur.Left] {
                    queue = append(queue, cur.Left)
                    visited[cur.Left] = true
                }
                if cur.Right != nil && !visited[cur.Right] {
                    queue = append(queue, cur.Right)
                    visited[cur.Right] = true
                }
                parentNode := parentMap[cur]
                if parentNode != nil && !visited[parentNode] {
                    queue = append(queue, parentNode)
                    visited[parentNode] = true
                }
            }
        }
        return -1
    }
    target := traverse(root, k, nil, parentMap) 
    return bfs(target, parentMap)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/13/closest1-tree.jpg" />
    // Input: root = [1,3,2], k = 1
    // Output: 2
    // Explanation: Either 2 or 3 is the nearest leaf node to the target of 1.
    tree1 := buildTree([]*int{ptr(1), ptr(3), ptr(2)})
    printTree(tree1)
    fmt.Println(findClosestLeaf(tree1,1)) // 2 | 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/13/closest2-tree.jpg" />
    // Input: root = [1], k = 1
    // Output: 1
    // Explanation: The nearest leaf node is the root node itself.
    tree2 := buildTree([]*int{ptr(1)})
    printTree(tree2)
    fmt.Println(findClosestLeaf(tree2,1)) // 1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/06/13/closest3-tree.jpg" />
    // Input: root = [1,2,3,4,null,null,null,5,null,6], k = 2
    // Output: 3
    // Explanation: The leaf node with value 3 (and not the leaf node with value 6) is nearest to the node with value 2.
    tree3 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), nil, nil, nil, ptr(5), nil, ptr(6)})
    printTree(tree3)
    fmt.Println(findClosestLeaf(tree3,2)) // 3

    tree11 := buildTree([]*int{ptr(1), ptr(3), ptr(2)})
    printTree(tree11)
    fmt.Println(findClosestLeaf1(tree11,1)) // 2
    tree12 := buildTree([]*int{ptr(1)})
    printTree(tree12)
    fmt.Println(findClosestLeaf1(tree12,1)) // 1
    tree13 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), nil, nil, nil, ptr(5), nil, ptr(6)})
    printTree(tree13)
    fmt.Println(findClosestLeaf1(tree13,2)) // 3
}