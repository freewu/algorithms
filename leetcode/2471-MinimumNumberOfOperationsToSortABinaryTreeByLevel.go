package main

// 2471. Minimum Number of Operations to Sort a Binary Tree by Level
// You are given the root of a binary tree with unique values.

// In one operation, you can choose any two nodes at the same level and swap their values.

// Return the minimum number of operations needed to make the values at each level sorted in a strictly increasing order.

// The level of a node is the number of edges along the path between it and the root node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/09/18/image-20220918174006-2.png" />
//             1
//           /    \
//          4      3
//        /   \   /   \
//       7    6  8     5
//              /     /
//             9     10
// Input: root = [1,4,3,7,6,8,5,null,null,null,null,9,null,10]
// Output: 3
// Explanation:
// - Swap 4 and 3. The 2nd level becomes [3,4].
// - Swap 7 and 5. The 3rd level becomes [5,6,8,7].
// - Swap 8 and 7. The 3rd level becomes [5,6,7,8].
// We used 3 operations so return 3.
// It can be proven that 3 is the minimum number of operations needed.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/09/18/image-20220918174026-3.png" />
//             1
//          /     \
//         3       2
//       /   \   /    \
//      7     6 5      4
// Input: root = [1,3,2,7,6,5,4]
// Output: 3
// Explanation:
// - Swap 3 and 2. The 2nd level becomes [2,3].
// - Swap 7 and 4. The 3rd level becomes [4,6,5,7].
// - Swap 6 and 5. The 3rd level becomes [4,5,6,7].
// We used 3 operations so return 3.
// It can be proven that 3 is the minimum number of operations needed.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2022/09/18/image-20220918174052-4.png" />
//            1
//          /    \
//         2      3
//       /   \    / 
//      4     5  6
// Input: root = [1,2,3,4,5,6]
// Output: 0
// Explanation: Each level is already sorted in increasing order so return 0.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     1 <= Node.val <= 10^5
//     All the values of the tree are unique.

import "fmt"
import "sort"

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
func minimumOperations(root *TreeNode) int {
    if root == nil { return 0 }
    helper := func(nums []int) int {
        sorted := make([]int, len(nums))
        copy(sorted, nums)
        sort.Ints(sorted)
        res, index := 0, make(map[int]int)
        for i, v := range nums {
            index[v] = i
        }
        for i := range sorted {
            if sorted[i] != nums[i] {
                res++
                newPos := index[sorted[i]]
                index[nums[i]] = newPos
                index[nums[newPos]] = i
                nums[i], nums[newPos] = nums[newPos], nums[i]
            }
        }
        return res
    }
    res, queue := 0, []*TreeNode{root}
    for len(queue) > 0 {
        n := len(queue)
        arr := make([]int, 0, n)
        
        for i := 0; i < n; i++ {
            top := queue[0]
            queue = queue[1:] // pop
            arr = append(arr, top.Val)
            if top.Left != nil  { queue = append(queue, top.Left)  }
            if top.Right != nil { queue = append(queue, top.Right) }
        }
        res += helper(arr)
    }
    return res
}

func minimumOperations1(root *TreeNode) int {
    res, queue := 0, []*TreeNode{root}
    for len(queue) > 0 {
        n := len(queue)
        arr := make([]int, n)
        tmp := queue
        queue = nil
        for i, node := range tmp {
            arr[i] = node.Val
            if node.Left != nil  { queue = append(queue, node.Left) }
            if node.Right != nil { queue = append(queue, node.Right) }
        }
        index := make([]int, n)
        for i := range index {
            index[i] = i
        }
        sort.Slice(index, func(i, j int) bool {
            return arr[index[i]] < arr[index[j]]
        })
        res += n
        visited := make([]bool, n)
        for _, v := range index {
            if !visited[v] {
                for ; !visited[v]; v = index[v] {
                    visited[v] = true
                }
                res--
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/09/18/image-20220918174006-2.png" />
    //             1
    //           /    \
    //          4      3
    //        /   \   /   \
    //       7    6  8     5
    //              /     /
    //             9     10
    // Input: root = [1,4,3,7,6,8,5,null,null,null,null,9,null,10]
    // Output: 3
    // Explanation:
    // - Swap 4 and 3. The 2nd level becomes [3,4].
    // - Swap 7 and 5. The 3rd level becomes [5,6,8,7].
    // - Swap 8 and 7. The 3rd level becomes [5,6,7,8].
    // We used 3 operations so return 3.
    // It can be proven that 3 is the minimum number of operations needed.
    tree1 := buildTree([]*int{ptr(1), ptr(4), ptr(3), ptr(7), ptr(6), ptr(8), ptr(5), nil, nil, nil, nil, ptr(9), nil, ptr(10)})
    printTree(tree1)
    fmt.Println(minimumOperations(tree1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/09/18/image-20220918174026-3.png" />
    //             1
    //          /     \
    //         3       2
    //       /   \   /    \
    //      7     6 5      4
    // Input: root = [1,3,2,7,6,5,4]
    // Output: 3
    // Explanation:
    // - Swap 3 and 2. The 2nd level becomes [2,3].
    // - Swap 7 and 4. The 3rd level becomes [4,6,5,7].
    // - Swap 6 and 5. The 3rd level becomes [4,5,6,7].
    // We used 3 operations so return 3.
    // It can be proven that 3 is the minimum number of operations needed.
    tree2 := buildTree([]*int{ptr(1), ptr(3), ptr(2), ptr(7), ptr(6), ptr(5), ptr(4)})
    printTree(tree2)
    fmt.Println(minimumOperations(tree2)) // 3
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2022/09/18/image-20220918174052-4.png" />
    //            1
    //          /    \
    //         2      3
    //       /   \    / 
    //      4     5  6
    // Input: root = [1,2,3,4,5,6]
    // Output: 0
    // Explanation: Each level is already sorted in increasing order so return 0.
    tree3 := buildTree([]*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(5), ptr(6)})
    printTree(tree3)
    fmt.Println(minimumOperations(tree3)) // 0

    fmt.Println(minimumOperations1(tree1)) // 3
    fmt.Println(minimumOperations1(tree2)) // 3
    fmt.Println(minimumOperations1(tree3)) // 0
}