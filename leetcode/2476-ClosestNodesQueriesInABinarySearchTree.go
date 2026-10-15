package main

// 2476. Closest Nodes Queries in a Binary Search Tree
// You are given the root of a binary search tree and an array queries of size n consisting of positive integers.
// Find a 2D array answer of size n where answer[i] = [mini, maxi]:
//     mini is the largest value in the tree that is smaller than or equal to queries[i]. 
//     If a such value does not exist, add -1 instead.
//     maxi is the smallest value in the tree that is greater than or equal to queries[i]. 
//     If a such value does not exist, add -1 instead.

// Return the array answer.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/09/28/bstreeedrawioo.png" />
// Input: root = [6,2,13,1,4,9,15,null,null,null,null,null,null,14], queries = [2,5,16]
// Output: [[2,2],[4,6],[15,-1]]
// Explanation: We answer the queries in the following way:
// - The largest number that is smaller or equal than 2 in the tree is 2, and the smallest number that is greater or equal than 2 is still 2. So the answer for the first query is [2,2].
// - The largest number that is smaller or equal than 5 in the tree is 4, and the smallest number that is greater or equal than 5 is 6. So the answer for the second query is [4,6].
// - The largest number that is smaller or equal than 16 in the tree is 15, and the smallest number that is greater or equal than 16 does not exist. So the answer for the third query is [15,-1].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/09/28/bstttreee.png" />
// Input: root = [4,null,9], queries = [3]
// Output: [[-1,4]]
// Explanation: The largest number that is smaller or equal to 3 in the tree does not exist, and the smallest number that is greater or equal to 3 is 4. So the answer for the query is [-1,4].
 
// Constraints:
//         The number of nodes in the tree is in the range [2, 10^5].
//         1 <= Node.val <= 10^6
//         n == queries.length
//         1 <= n <= 10^5
//         1 <= queries[i] <= 10^6

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

// Time Limit Exceeded
func closestNodes(root *TreeNode, queries []int) [][]int {
    res := [][]int{}
    for _ , n := range queries {
        // answer[i] = [mini, maxi] ：
        // mini 是树中小于等于 queries[i] 的 最大值 。如果不存在这样的值，则使用 -1 代替。
        // maxi 是树中大于等于 queries[i] 的 最小值 。如果不存在这样的值，则使用 -1 代替。
        curr := []int{-1,-1}
        traverseTree(root, n, &curr)
        res = append(res, curr)
    }
    return res
}

func traverseTree(root *TreeNode, n int, res *[]int) {
    if root != nil {
        if root.Val < n { // 树中小于等于 queries[i] 的 最大值
            (*res)[0] = root.Val
            traverseTree(root.Right, n, res)
        } else if root.Val > n { // 树中大于等于 queries[i] 的 最小值
            (*res)[1] = root.Val
            traverseTree(root.Left, n, res)
        } else {
            (*res)[0] = n
            (*res)[1] = n
        }
    }
}

// 由于该二叉搜索树并不是平衡的，则最坏情况该树可能形成一条链，直接在树上查找可能存在超时。
// 我们可以保存树中所有的节点值，并将其排序，每次给定查询值 val 时，利用二分查找直接在树中找到大于等于 val 的最小值与小于等于 val 的最小值
func closestNodes1(root *TreeNode, queries []int) [][]int {
    arr := []int{}
    var dfs func(*TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        arr = append(arr, root.Val)
        dfs(root.Right)
    }
    
    dfs(root)
    res := make([][]int, len(queries))
    for i, val := range queries {
        maxVal, minVal := -1, -1
        index := sort.SearchInts(arr, val)
        if index < len(arr) {
            maxVal = arr[index]
            if arr[index] == val {
                minVal = arr[index]
                res[i] = []int{minVal, maxVal}
                continue
            }
        }
        if index != 0 {
            minVal = arr[index - 1]
        }
        res[i] = []int{minVal, maxVal}
    }
    return res
}

// best solution
func closestNodes2(root *TreeNode, queries []int) [][]int {
    a := []int{}
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        a = append(a, node.Val)
        dfs(node.Right)
    }
    dfs(root)

    res := make([][]int, len(queries))
    for i, q := range queries {
        mn, mx := -1, -1
        j, ok := slices.BinarySearch(a, q)
        if j < len(a) {
            mx = a[j]
        }
        if !ok { // a[j]>q, a[j-1]<q
            j--
        }
        if j >= 0 {
            mn = a[j]
        }
        res[i] = []int{mn, mx}
    }
    return res
}