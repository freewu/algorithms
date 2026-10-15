package main

// 1586. Binary Search Tree Iterator II
// Implement the BSTIterator class that represents an iterator over the in-order traversal of a binary search tree (BST):
//     BSTIterator(TreeNode root) 
//         Initializes an object of the BSTIterator class. 
//         The root of the BST is given as part of the constructor. 
//         The pointer should be initialized to a non-existent number smaller than any element in the BST.
//     boolean hasNext() 
//         Returns true if there exists a number in the traversal to the right of the pointer, otherwise returns false.
//     int next() 
//         Moves the pointer to the right, then returns the number at the pointer.
//     boolean hasPrev() 
//         Returns true if there exists a number in the traversal to the left of the pointer, otherwise returns false.
//     int prev() 
//         Moves the pointer to the left, then returns the number at the pointer.

// Notice that by initializing the pointer to a non-existent smallest number, 
// the first call to next() will return the smallest element in the BST.

// You may assume that next() and prev() calls will always be valid. 
// That is, there will be at least a next/previous number in the in-order traversal when next()/prev() is called.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/14/untitled-diagram-1.png" />
// Input
// ["BSTIterator", "next", "next", "prev", "next", "hasNext", "next", "next", "next", "hasNext", "hasPrev", "prev", "prev"]
// [[[7, 3, 15, null, null, 9, 20]], [null], [null], [null], [null], [null], [null], [null], [null], [null], [null], [null], [null]]
// Output
// [null, 3, 7, 3, 7, true, 9, 15, 20, false, true, 15, 9]
// Explanation
// // The underlined element is where the pointer currently is.
// BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]); // state is   [3, 7, 9, 15, 20]
// bSTIterator.next(); // state becomes [3, 7, 9, 15, 20], return 3
// bSTIterator.next(); // state becomes [3, 7, 9, 15, 20], return 7
// bSTIterator.prev(); // state becomes [3, 7, 9, 15, 20], return 3
// bSTIterator.next(); // state becomes [3, 7, 9, 15, 20], return 7
// bSTIterator.hasNext(); // return true
// bSTIterator.next(); // state becomes [3, 7, 9, 15, 20], return 9
// bSTIterator.next(); // state becomes [3, 7, 9, 15, 20], return 15
// bSTIterator.next(); // state becomes [3, 7, 9, 15, 20], return 20
// bSTIterator.hasNext(); // return false
// bSTIterator.hasPrev(); // return true
// bSTIterator.prev(); // state becomes [3, 7, 9, 15, 20], return 15
// bSTIterator.prev(); // state becomes [3, 7, 9, 15, 20], return 9

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     0 <= Node.val <= 10^6
//     At most 10^5 calls will be made to hasNext, next, hasPrev, and prev. 

// Follow up: Could you solve the problem without precalculating the values of the tree?

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
type BSTIterator struct {
    arr     []int
    n       int
    pointer int
}

func Constructor(root *TreeNode) BSTIterator {
    iterator := BSTIterator{}
    iterator.inorder(root) // 中序遍历出数据到数组
    iterator.n = len(iterator.arr)
    iterator.pointer = -1
    return iterator
}

func (this *BSTIterator) HasNext() bool {
    return this.pointer < this.n - 1
}


func (this *BSTIterator) Next() int {
    this.pointer++
    min := func (x, y int) int { if x < y { return x; }; return y; }
    return this.arr[min(this.pointer, this.n - 1)]
}


func (this *BSTIterator) HasPrev() bool {
    return this.pointer > 0
}


func (this *BSTIterator) Prev() int {
    this.pointer--
    max := func (x, y int) int { if x > y { return x; }; return y; }
    return this.arr[max(this.pointer, 0)]
}

func (this *BSTIterator) inorder(root *TreeNode) {
    if root == nil { return }
    this.inorder(root.Left)
    this.arr = append(this.arr, root.Val)
    this.inorder(root.Right)
}
 
 
 /**
  * Your BSTIterator object will be instantiated and called as such:
  * obj := Constructor(root);
  * param_1 := obj.HasNext();
  * param_2 := obj.Next();
  * param_3 := obj.HasPrev();
  * param_4 := obj.Prev();
  */
func main() {
    // BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]); // state is   [3, 7, 9, 15, 20]
    tree1 := buildTree([]*int{ptr(7), ptr(3), ptr(15), nil, nil, ptr(9), ptr(20)})
    printTree(tree1)
    obj := Constructor(tree1)
    // bSTIterator.next(); // state becomes [3, 7, 9, 15, 20], return 3
    fmt.Println(obj.Next()) // 3
    // bSTIterator.next(); // state becomes [3, 7, 9, 15, 20], return 7
    fmt.Println(obj.Next()) // 7
    // bSTIterator.prev(); // state becomes [3, 7, 9, 15, 20], return 3
    fmt.Println(obj.Prev()) // 3
    // bSTIterator.next(); // state becomes [3, 7, 9, 15, 20], return 7
    fmt.Println(obj.Next()) // 7
    // bSTIterator.hasNext(); // return true
    fmt.Println(obj.HasNext()) // true
    // bSTIterator.next(); // state becomes [3, 7, 9, 15, 20], return 9
    fmt.Println(obj.Next()) // 9
    // bSTIterator.next(); // state becomes [3, 7, 9, 15, 20], return 15
    fmt.Println(obj.Next()) // 15
    // bSTIterator.next(); // state becomes [3, 7, 9, 15, 20], return 20
    fmt.Println(obj.Next()) // 30
    // bSTIterator.hasNext(); // return false
    fmt.Println(obj.HasNext()) // false
    // bSTIterator.hasPrev(); // return true
    fmt.Println(obj.HasPrev()) // true
    // bSTIterator.prev(); // state becomes [3, 7, 9, 15, 20], return 15
    fmt.Println(obj.Prev()) // 15
    // bSTIterator.prev(); // state becomes [3, 7, 9, 15, 20], return 9
    fmt.Println(obj.Prev()) // 9
}