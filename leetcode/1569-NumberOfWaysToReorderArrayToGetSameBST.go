package main

// 1569. Number of Ways to Reorder Array to Get Same BST
// Given an array nums that represents a permutation of integers from 1 to n. 
// We are going to construct a binary search tree (BST) by inserting the elements of nums in order into an initially empty BST. 
// Find the number of different ways to reorder nums so that the constructed BST is identical to that formed from the original array nums.
//     For example, given nums = [2,1,3], we will have 2 as the root, 1 as a left child, and 3 as a right child. 
//     The array [2,3,1] also yields the same BST but [3,2,1] yields a different BST.

// Return the number of ways to reorder nums such that the BST formed is identical to the original BST formed from nums.
// Since the answer may be very large, return it modulo 10^9 + 7.

// Example 1:
//      2
//    /   \
//   1     3
// <img src="https://assets.leetcode.com/uploads/2020/08/12/bb.png" />
// Input: nums = [2,1,3]
// Output: 1
// Explanation: We can reorder nums to be [2,3,1] which will yield the same BST. There are no other ways to reorder nums which will yield the same BST.

// Example 2:
//         3
//       /   \
//      1     4
//       \     \ 
//        2      3
// <img src="https://assets.leetcode.com/uploads/2020/08/12/ex1.png" />
// Input: nums = [3,4,5,1,2]
// Output: 5
// Explanation: The following 5 arrays will yield the same BST: 
// [3,1,2,4,5]
// [3,1,4,2,5]
// [3,1,4,5,2]
// [3,4,1,2,5]
// [3,4,1,5,2]

// Example 3:
//     1
//      \ 
//       2 
//        \ 
//         3
// <img src="https://assets.leetcode.com/uploads/2020/08/12/ex4.png" />
// Input: nums = [1,2,3]
// Output: 0
// Explanation: There are no other orderings of nums that will yield the same BST.
 
// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= nums.length
//     All integers in nums are distinct.

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

func numOfWays(nums []int) int {
    binom, mod := make([][]int,len(nums)),1_000_000_007
    for i:=0; i < len(nums); i++ {
        binom[i] = make([]int,i+1)
        binom[i][0], binom[i][i] = 1, 1
        for j := 1; j < i; j++ {
            binom[i][j]= (binom[i-1][j]+binom[i-1][j-1]) % mod
        }
    }
    var count func( nums []int) int
    count = func ( nums []int) int {
        n := len(nums)
        if n <= 2 {
            return 1
        }
        root,left, right := nums[0], []int{}, []int{}
        for i := 1; i < n; i++ {
            if nums[i] < root {
                left = append(left, nums[i])
            } else {
                right = append(right, nums[i])
            }
        }
        lc := count(left) % mod
        rc := count(right) % mod
        return (((lc * rc) % mod) * binom[n-1][len(left)] ) % mod
    }
    return count(nums) - 1
}

func numOfWays1(nums []int) int {
    type node struct {
        l, r  int
        i     int
        left  *node
        right *node
    }    
    n, mod := len(nums), 1_000_000_007
    caches := make([]int, n + 1)
    caches[0] = 1
    for i := 1; i < n + 1; i++ {
        caches[i] = caches[i-1] * i % mod
    }
    cacheDivs := make([]int, n + 1)
    getDiv := func(i int) int {
        if cacheDivs[i] != 0 {
            return cacheDivs[i]
        }
        res, r := 1, caches[i]
        for k := int64(mod - 2); k != 0; k >>= 1 {
            if k & 1 == 1 {
                res = (res * r) % mod
            }
            r = r * r % mod
        }
        cacheDivs[i] = res
        return res
    }
    root := node{ l: 1, r: len(nums), i: nums[0] }

    for i := 1; i < len(nums); i++ {
        cur := &root
        for {
            if cur.i > nums[i] {
                if cur.left != nil {
                    cur = cur.left
                    continue
                }
                cur.left = &node{ l: cur.l, r: cur.i - 1, i: nums[i], }
                break
            } else {
                if cur.right != nil {
                    cur = cur.right
                    continue
                }
                cur.right = &node{ l: cur.i + 1, r: cur.r, i: nums[i], }
                break
            }
        }
    }
    var getRes func(cur *node) int
    getRes = func(cur *node) int {
        if cur == nil {
            return 1
        }
        res := 1
        res = res * getRes(cur.right) % mod
        res = res * getRes(cur.left) % mod
        n, r := cur.r - cur.l, cur.i - cur.l
        res = res * caches[n] % mod * getDiv(n-r) % mod * getDiv(r) % mod
        return res
    }
    return (getRes(&root) - 1 + mod) % mod
}

func main() {
    // Example 1:
    //      2
    //    /   \
    //   1     3
    // <img src="https://assets.leetcode.com/uploads/2020/08/12/bb.png" />
    // Input: nums = [2,1,3]
    // Output: 1
    // Explanation: We can reorder nums to be [2,3,1] which will yield the same BST. There are no other ways to reorder nums which will yield the same BST.
    tree1 := buildTree([]*int{ptr(2), ptr(1), ptr(3)})
    printTree(tree1)
    fmt.Println(tree1)
    fmt.Println(numOfWays([]int{2,1,3})) // 1
    // Example 2:
    //         3
    //       /   \
    //      1     4
    //       \     \ 
    //        2     5
    // <img src="https://assets.leetcode.com/uploads/2020/08/12/ex1.png" />
    // Input: nums = [3,4,5,1,2]
    // Output: 5
    // Explanation: The following 5 arrays will yield the same BST: 
    // [3,1,2,4,5]
    // [3,1,4,2,5]
    // [3,1,4,5,2]
    // [3,4,1,2,5]
    // [3,4,1,5,2]
    tree2 := buildTree([]*int{ptr(3), ptr(1), ptr(4), nil, ptr(2), nil, ptr(5)})
    printTree(tree2)
    fmt.Println(tree2)
    fmt.Println(numOfWays([]int{3,4,5,1,2})) // 5
    // Example 3:
    //     1
    //      \ 
    //       2 
    //        \ 
    //         3
    // <img src="https://assets.leetcode.com/uploads/2020/08/12/ex4.png" />
    // Input: nums = [1,2,3]
    // Output: 0
    // Explanation: There are no other orderings of nums that will yield the same BST.
    tree3 := buildTree([]*int{ptr(1), nil, ptr(2), nil, ptr(1)})
    printTree(tree3)
    fmt.Println(tree3)
    fmt.Println(numOfWays([]int{1,2,3})) // 0

    fmt.Println(numOfWays1([]int{2,1,3})) // 1
    fmt.Println(numOfWays1([]int{3,4,5,1,2})) // 5
    fmt.Println(numOfWays1([]int{1,2,3})) // 0
}