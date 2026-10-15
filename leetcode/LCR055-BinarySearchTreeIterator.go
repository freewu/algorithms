package main

// LCR 055. 二叉搜索树迭代器
// 实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
//     BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
//     boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
//     int next()将指针向右移动，然后返回指针处的数字。

// 注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。
// 可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 的中序遍历中至少存在一个下一个数字。

// 示例：
// 输入
// inputs = ["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
// inputs = [[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
// 输出
// [null, 3, 7, true, 9, true, 15, true, 20, false]
// 解释
// BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
// bSTIterator.next();    // 返回 3
// bSTIterator.next();    // 返回 7
// bSTIterator.hasNext(); // 返回 True
// bSTIterator.next();    // 返回 9
// bSTIterator.hasNext(); // 返回 True
// bSTIterator.next();    // 返回 15
// bSTIterator.hasNext(); // 返回 True
// bSTIterator.next();    // 返回 20
// bSTIterator.hasNext(); // 返回 False

// 提示：
//     树中节点的数目在范围 [1, 10^5] 内
//     0 <= Node.val <= 10^6
//     最多调用 10^5 次 hasNext 和 next 操作

// 进阶：
//     你可以设计一个满足下述条件的解决方案吗？next() 和 hasNext() 操作均摊时间复杂度为 O(1) ，并使用 O(h) 内存。其中 h 是树的高度。

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
type BSTIterator1 struct {
    stack []*TreeNode
}

func Constructor1(root *TreeNode) BSTIterator1 {
    iterator := BSTIterator1{stack: []*TreeNode{root}}
    iterator.GetLeftPath()
    return iterator
}

func (this *BSTIterator1) GetLeftPath() {
    currentNode := this.stack[len(this.stack) - 1]
    for currentNode.Left != nil {
        this.stack = append(this.stack, currentNode.Left)
        currentNode = currentNode.Left
    }
}

func (this *BSTIterator1) Next() int {
    currentNode := this.stack[len(this.stack) - 1]
    this.stack = this.stack[:len(this.stack) - 1]
    if currentNode.Right != nil {
        this.stack = append(this.stack, currentNode.Right)
        this.GetLeftPath()
    }
    return currentNode.Val
}

func (this *BSTIterator1) HasNext() bool {
    return len(this.stack) != 0
}

type stack struct {
    buffer []*TreeNode
}

func (s *stack) Len() int {
    return len(s.buffer)
}

func (s *stack) Push(node *TreeNode) {
    s.buffer = append(s.buffer, node)
}

func (s *stack) IsEmpty() bool {
    return len(s.buffer) == 0
}

func (s *stack) Peek() *TreeNode {
    if s.IsEmpty() {
        return nil
    }

    idx := len(s.buffer) - 1
    return s.buffer[idx]
}

func (s *stack) Pop() {
    if s.IsEmpty() {
        return
    }

    idx := len(s.buffer) - 1
    s.buffer = s.buffer[:idx]
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
    stack *stack
    curr *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
    stack := &stack {
        buffer: []*TreeNode{},
    }

    return BSTIterator {
        stack: stack,
        curr: root,
    }
}

func (this *BSTIterator) Next() int {
    node := this.curr
    for node != nil {
        this.stack.Push(node)
        node = node.Left
    }

    node = this.stack.Peek()
    this.stack.Pop()

    this.curr = node.Right
    return node.Val
}

func (this *BSTIterator) HasNext() bool {
    return !this.stack.IsEmpty() || this.curr != nil
}

/**
* Your BSTIterator object will be instantiated and called as such:
* obj := Constructor(root);
* param_1 := obj.Next();
* param_2 := obj.HasNext();
*/

func main() {
    tree1 := buildTree([]*int{ptr(7), ptr(3), ptr(15), nil, nil, ptr(9), ptr(20)})
    printTree(tree1)
    // Explanation
    // BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
    obj := Constructor(tree1)
    // bSTIterator.next();    // return 3
    fmt.Println(obj.Next()) // 3
    // bSTIterator.next();    // return 7
    fmt.Println(obj.Next()) // 7
    // bSTIterator.hasNext(); // return True
    fmt.Println(obj.HasNext()) // true
    // bSTIterator.next();    // return 9
    fmt.Println(obj.Next()) // 9
    // bSTIterator.hasNext(); // return True
    fmt.Println(obj.HasNext()) // true
    // bSTIterator.next();    // return 15
    fmt.Println(obj.Next()) // 15
    // bSTIterator.hasNext(); // return True
    fmt.Println(obj.HasNext()) // true
    // bSTIterator.next();    // return 20
    fmt.Println(obj.Next()) // 20
    // bSTIterator.hasNext(); // return False
    fmt.Println(obj.HasNext()) // false

    tree11 := buildTree([]*int{ptr(7), ptr(3), ptr(15), nil, nil, ptr(9), ptr(20)})
    printTree(tree11)
    // Explanation
    // BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
    obj1 := Constructor1(tree11)
    // bSTIterator.next();    // return 3
    fmt.Println(obj1.Next()) // 3
    // bSTIterator.next();    // return 7
    fmt.Println(obj1.Next()) // 7
    // bSTIterator.hasNext(); // return True
    fmt.Println(obj1.HasNext()) // true
    // bSTIterator.next();    // return 9
    fmt.Println(obj1.Next()) // 9
    // bSTIterator.hasNext(); // return True
    fmt.Println(obj1.HasNext()) // true
    // bSTIterator.next();    // return 15
    fmt.Println(obj1.Next()) // 15
    // bSTIterator.hasNext(); // return True
    fmt.Println(obj1.HasNext()) // true
    // bSTIterator.next();    // return 20
    fmt.Println(obj1.Next()) // 20
    // bSTIterator.hasNext(); // return False
    fmt.Println(obj1.HasNext()) // false
}