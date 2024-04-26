package main

// 173. Binary Search Tree Iterator
// Implement the BSTIterator class that represents an iterator over the in-order traversal of a binary search tree (BST):
//     BSTIterator(TreeNode root) 
//         Initializes an object of the BSTIterator class. 
//         The root of the BST is given as part of the constructor. 
//         The pointer should be initialized to a non-existent number smaller than any element in the BST.
//     boolean hasNext() 
//         Returns true if there exists a number in the traversal to the right of the pointer, 
//         otherwise returns false.
//     int next() 
//         Moves the pointer to the right, then returns the number at the pointer.

// Notice that by initializing the pointer to a non-existent smallest number, 
// the first call to next() will return the smallest element in the BST.

// You may assume that next() calls will always be valid. 
// That is, there will be at least a next number in the in-order traversal when next() is called.

// Example 1:

// Input
// ["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
// [[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
// Output
// [null, 3, 7, true, 9, true, 15, true, 20, false]
// Explanation
// BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
// bSTIterator.next();    // return 3
// bSTIterator.next();    // return 7
// bSTIterator.hasNext(); // return True
// bSTIterator.next();    // return 9
// bSTIterator.hasNext(); // return True
// bSTIterator.next();    // return 15
// bSTIterator.hasNext(); // return True
// bSTIterator.next();    // return 20
// bSTIterator.hasNext(); // return False

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     0 <= Node.val <= 10^6
//     At most 10^5 calls will be made to hasNext, and next.
 
// Follow up:
//     Could you implement next() and hasNext() to run in average O(1) time and use O(h) memory, where h is the height of the tree?

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
// type BSTIterator struct {
//     stack []*TreeNode
// }

// func Constructor(root *TreeNode) BSTIterator {
//     iterator := BSTIterator{stack: []*TreeNode{root}}
//     iterator.GetLeftPath()
//     return iterator
// }

// func (this *BSTIterator) GetLeftPath() {
//     currentNode := this.stack[len(this.stack) - 1]
//     for currentNode.Left != nil {
//         this.stack = append(this.stack, currentNode.Left)
//         currentNode = currentNode.Left
//     }
// }

// func (this *BSTIterator) Next() int {
//     currentNode := this.stack[len(this.stack) - 1]
//     this.stack = this.stack[:len(this.stack) - 1]
//     if currentNode.Right != nil {
//         this.stack = append(this.stack, currentNode.Right)
//         this.GetLeftPath()
//     }
//     return currentNode.Val
// }

// func (this *BSTIterator) HasNext() bool {
//     return len(this.stack) != 0
// }

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
}