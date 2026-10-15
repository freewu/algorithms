package main

// 108. Convert Sorted Array to Binary Search Tree
// Given an integer array nums where the elements are sorted in ascending order, 
// convert it to a height-balanced binary search tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/18/btree1.jpg" />
// <img src="https://assets.leetcode.com/uploads/2021/02/18/btree2.jpg" />
//         0             0
//        / \		     /  \
//      -3   9   =>  -10    5
//      /   /          \      \
//     -10  5            -3     9
// Input: nums = [-10,-3,0,5,9]
// Output: [0,-3,9,-10,null,5]
// Explanation: [0,-10,5,null,-3,null,9] is also accepted:

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/02/18/btree.jpg" />
//         3      1
//         /	=>  \
//        1         3
// Input: nums = [1,3]
// Output: [3,1]
// Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.
 
// Constraints:
//     1 <= nums.length <= 10^4
//     -10^4 <= nums[i] <= 10^4
//     nums is sorted in a strictly increasing order.

// 解题思路:
//     把一个有序数组转换成高度平衡的二叉搜索数

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
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    return &TreeNode{
        Val: nums[len(nums)/2], // 取中间值
        Left: sortedArrayToBST(nums[:len(nums)/2]), // 取左半部分 递归
        Right: sortedArrayToBST(nums[len(nums)/2+1:]), // 取右半部分 递归
    }
    }

// best solution
func sortedArrayToBST1(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    if len(nums) == 1 { // 多了这部分 如果只有一个值时直接返回无叶子的节点，减少递归调用
        return &TreeNode{Val: nums[0]}
    }
    mid := len(nums)/2
    root := &TreeNode{
        Val: nums[mid],
        Left: sortedArrayToBST1(nums[:mid]),
        Right: sortedArrayToBST1(nums[mid+1:]),
    }
    return root
}

func sortedArrayToBST2(nums []int) *TreeNode {
    var dfs func(left, right int, nums []int) *TreeNode
    dfs = func(left, right int, nums []int) *TreeNode {
        if left > right {
            return nil
        }
        mid := (left + right) / 2
        root := &TreeNode{Val: nums[mid]}
        root.Left = dfs(left, mid - 1, nums)
        root.Right = dfs(mid + 1, right, nums)
        return root
    }
    return dfs(0, len(nums) - 1, nums)
}

func main() {
    fmt.Printf("sortedArrayToBST([]int{-10,-3,0,5,9}) = %v\n",sortedArrayToBST([]int{-10,-3,0,5,9}))
    fmt.Printf("sortedArrayToBST([]int{13})= %v\n",sortedArrayToBST([]int{1,3}))

    fmt.Printf("sortedArrayToBST1([]int{-10,-3,0,5,9}) = %v\n",sortedArrayToBST1([]int{-10,-3,0,5,9}))
    fmt.Printf("sortedArrayToBST1([]int{13})= %v\n",sortedArrayToBST1([]int{1,3}))

    fmt.Printf("sortedArrayToBST2([]int{-10,-3,0,5,9}) = %v\n",sortedArrayToBST2([]int{-10,-3,0,5,9}))
    fmt.Printf("sortedArrayToBST2([]int{1,3})= %v\n",sortedArrayToBST2([]int{1,3}))
}
