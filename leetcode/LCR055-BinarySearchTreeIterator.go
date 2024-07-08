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
    tree1 := &TreeNode {
        7,
        &TreeNode{3, nil, nil},
        &TreeNode{15, &TreeNode{9, nil, nil}, &TreeNode{20, nil, nil}, },
    }
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

    tree11 := &TreeNode {
        7,
        &TreeNode{3, nil, nil},
        &TreeNode{15, &TreeNode{9, nil, nil}, &TreeNode{20, nil, nil}, },
    }
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