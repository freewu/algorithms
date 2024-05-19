package main

// 431. Encode N-ary Tree to Binary Tree
// Design an algorithm to encode an N-ary tree into a binary tree and decode the binary tree to get the original N-ary tree. 
// An N-ary tree is a rooted tree in which each node has no more than N children. 
// Similarly, a binary tree is a rooted tree in which each node has no more than 2 children. 
// There is no restriction on how your encode/decode algorithm should work. 
// You just need to ensure that an N-ary tree can be encoded to a binary tree and this binary tree can be decoded to the original N-nary tree structure.

// Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value (See following example).
// For example, you may encode the following 3-ary tree to a binary tree in this way:
// <img src="https://assets.leetcode.com/uploads/2018/10/12/narytreebinarytreeexample.png" />

// Input: root = [1,null,3,2,4,null,5,6]
// Note that the above is just an example which might or might not work. 
// You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

// Example 1:
// Input: root = [1,null,3,2,4,null,5,6]
// Output: [1,null,3,2,4,null,5,6]

// Example 2:
// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]

// Example 3:
// Input: root = []
// Output: []

// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     0 <= Node.val <= 10^4
//     The height of the n-ary tree is less than or equal to 1000
//     Do not use class member/global/static variables to store states. Your encode and decode algorithms should be stateless.

import "fmt"

type Node struct {
    Val int
    Children []*Node
}

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Codec struct {
}

func Constructor() *Codec {
    return &Codec{}
}

func (this *Codec) encode(root *Node) *TreeNode {
    if root == nil {
        return nil
    }
    tree := &TreeNode{ Val: root.Val }
    if 0 < len(root.Children) {
        tree.Left = this.encode(root.Children[0])
        cur := tree.Left
        for _, c := range root.Children[1:] {
            cur.Right = this.encode(c)
            cur = cur.Right
        }
    }
    return tree
}

func (this *Codec) decode(root *TreeNode) *Node {
    if root == nil {
        return nil
    }
    node := &Node{Val: root.Val}
    cur := root.Left
    for cur != nil {
        node.Children = append(node.Children, this.decode(cur))
        cur = cur.Right
    }
    return node
}
 
 /**
  * Your Codec object will be instantiated and called as such:
  * obj := Constructor();
  * bst := obj.encode(root);
  * ans := obj.decode(bst);
  */

func main() {
// Example 1:
// Input: root = [1,null,3,2,4,null,5,6]
// Output: [1,null,3,2,4,null,5,6]

// Example 2:
// Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// Output: [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]

// Example 3:
// Input: root = []
// Output: []
}
