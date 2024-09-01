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
    tree1 := &TreeNode {
        7,
        &TreeNode{3, nil, nil},
        &TreeNode{15,  &TreeNode{9, nil, nil},  &TreeNode{20, nil, nil}, },
    }
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