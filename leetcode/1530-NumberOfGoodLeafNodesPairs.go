package main

// 1530. Number of Good Leaf Nodes Pairs
// You are given the root of a binary tree and an integer distance. 
// A pair of two different leaf nodes of a binary tree is said to be good if the length of the shortest path between them is less than or equal to distance.

// Return the number of good leaf node pairs in the tree.

// Example 1:
//         1
//       /   \
//      2     3
//       \ 
//         4
// <img src="https://assets.leetcode.com/uploads/2020/07/09/e1.jpg" />
// Input: root = [1,2,3,null,4], distance = 3
// Output: 1
// Explanation: The leaf nodes of the tree are 3 and 4 and the length of the shortest path between them is 3. This is the only good pair.

// Example 2:
//         1
//       /   \
//      2     3
//    /  \   /  \
//   4    5 6    7
// <img src="https://assets.leetcode.com/uploads/2020/07/09/e2.jpg" />
// Input: root = [1,2,3,4,5,6,7], distance = 3
// Output: 2
// Explanation: The good pairs are [4,5] and [6,7] with shortest path = 2. The pair [4,6] is not good because the length of ther shortest path between them is 4.

// Example 3:
// Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2], distance = 3
// Output: 1
// Explanation: The only good pair is [2,5].

// Constraints:
//     The number of nodes in the tree is in the range [1, 2^10].
//     1 <= Node.val <= 100
//     1 <= distance <= 10

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
func countPairs(root *TreeNode, distance int) int {
    res := 0
    var dfs func(*TreeNode) []int
    dfs = func(node *TreeNode) []int {
        if node == nil { return []int{} }
        if node.Left == nil && node.Right == nil { return []int{0} }
        arr := make([]int, 0)
        left := dfs(node.Left)
        for _, l := range left {
            if l++; l <= distance {
                arr = append(arr, l)
            }
        }
        right := dfs(node.Right)
        for _, r := range right {
            if r++; r <= distance {
                arr = append(arr, r)
            }
        }
        for _, l := range left {
            for _, r := range right {
                if l+r+2 <= distance { // a与left b与right+2
                    res++
                }
            }
        }
        return arr
    }
    dfs(root)
    return res
}

func countPairs1(root *TreeNode, distance int) int {
    var dfs func(*TreeNode) ([]int, int)
    dfs = func (node *TreeNode) ([]int, int) {
        if node == nil {
            return []int{}, 0
        }
        if node.Left == nil && node.Right == nil {
            return []int{1}, 0
        }
        lLeaves, lCount := dfs(node.Left)
        rLeaves, rCount := dfs(node.Right)
        count := lCount + rCount
        for _, left := range lLeaves {
            for _, right := range rLeaves {
                if left + right <= distance {
                    count++
                }
            }
        }
        leaves := make([]int, 0, len(lLeaves) + len(rLeaves))
        for _, leaf := range lLeaves {
            if leaf < 9 {
                leaves = append(leaves, leaf + 1)
            }
        }
        for _, leaf := range rLeaves {
            if leaf < 9 {
                leaves = append(leaves, leaf + 1)
            }
        }
        return leaves, count
    }
    _, res := dfs(root)
    return res
}
 
func main() {
    // Example 1:
    //         1
    //       /   \
    //      2     3
    //       \ 
    //         4
    // <img src="https://assets.leetcode.com/uploads/2020/07/09/e1.jpg" />
    // Input: root = [1,2,3,null,4], distance = 3
    // Output: 1
    // Explanation: The leaf nodes of the tree are 3 and 4 and the length of the shortest path between them is 3. This is the only good pair.
    tree1 := buildTree([]*int{ptr(1), ptr(2), ptr(3), nil, ptr(4)})
    printTree(tree1)
    fmt.Println(countPairs(tree1, 3)) // 1
    // Example 2:
    //         1
    //       /   \
    //      2     3
    //    /  \   /  \
    //   4    5 6    7
    // <img src="https://assets.leetcode.com/uploads/2020/07/09/e2.jpg" />
    // Input: root = [1,2,3,4,5,6,7], distance = 3
    // Output: 2
    // Explanation: The good pairs are [4,5] and [6,7] with shortest path = 2. The pair [4,6] is not good because the length of ther shortest path between them is 4.
    tree2 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(5), ptr(6), ptr(7)})
    printTree(tree2)
    fmt.Println(countPairs(tree2, 3)) // 2
    // Example 3:
    // Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2], distance = 3
    // Output: 1
    // Explanation: The only good pair is [2,5].
    tree3 := buildTree([]*int{ptr(7), ptr(1), ptr(4), ptr(6), nil, ptr(5), ptr(3), nil, nil, nil, nil, nil, ptr(2)})
    printTree(tree3)
    fmt.Println(countPairs(tree3, 3)) // 1


    tree11 := buildTree([]*int{ptr(1), ptr(2), ptr(3), nil, ptr(4)})
    printTree(tree11)
    fmt.Println(countPairs1(tree11, 3)) // 1
    tree12 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(5), ptr(6), ptr(7)})
    printTree(tree12)
    fmt.Println(countPairs1(tree12, 3)) // 2
    // Example 3:
    // Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2], distance = 3
    // Output: 1
    // Explanation: The only good pair is [2,5].
    tree13 := buildTree([]*int{ptr(7), ptr(1), ptr(4), ptr(6), nil, ptr(5), ptr(3), nil, nil, nil, nil, nil, ptr(2)})
    printTree(tree13)
    fmt.Println(countPairs1(tree13, 3)) // 1
}