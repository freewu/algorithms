package main

// 426. Convert Binary Search Tree to Sorted Doubly Linked List
// Convert a Binary Search Tree to a sorted Circular Doubly-Linked List in place.

// You can think of the left and right pointers as synonymous to the predecessor and successor pointers in a doubly-linked list. 
// For a circular doubly linked list, the predecessor of the first element is the last element, 
// and the successor of the last element is the first element.

// We want to do the transformation in place. 
// After the transformation, the left pointer of the tree node should point to its predecessor, 
// and the right pointer should point to its successor. 
// You should return the pointer to the smallest element of the linked list.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/10/12/bstdlloriginalbst.png" />
// Input: root = [4,2,5,1,3]
// <img src="https://assets.leetcode.com/uploads/2018/10/12/bstdllreturndll.png" />
// Output: [1,2,3,4,5]
// <img src="https://assets.leetcode.com/uploads/2018/10/12/bstdllreturnbst.png" />
// Explanation: The figure below shows the transformed BST. The solid line indicates the successor relationship, while the dashed line means the predecessor relationship.

// Example 2:
// Input: root = [2,1,3]
// Output: [1,2,3]

// Constraints:
//     The number of nodes in the tree is in the range [0, 2000].
//     -1000 <= Node.val <= 1000
//     All the values of the tree are unique.

import "fmt"

// Definition for a Node.
type Node struct {
    Val int
    Left *Node
    Right *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */
func treeToDoublyList(root *Node) *Node {
    //中序遍历
    // 把前后连接，每次返回前一个节点，并连接
    if root == nil {
        return nil
    }
    var inorder  func(root *Node) (left,right *Node)
    inorder = func(root *Node) (left,right *Node) {
        if root == nil {
            return nil, nil
        }
        l1,r1 := inorder(root.Left)
        l2,r2 := inorder(root.Right)
        if root.Left == nil && root.Right == nil {
            return root, root
        } else if root.Left == nil && root.Right != nil {
            root.Right = l2
            l2.Left = root
            return root, r2
        } else if root.Left != nil && root.Right == nil {
            root.Left = r1
            r1.Right = root
            return l1, root
        } else {
            root.Left = r1
            r1.Right = root
            root.Right = l2
            l2.Left = root
            return l1,  r2
        }
    }
    left, right := inorder(root)
    left.Left = right
    right.Right = left
    return left
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/10/12/bstdlloriginalbst.png" />
    // Input: root = [4,2,5,1,3]
    // <img src="https://assets.leetcode.com/uploads/2018/10/12/bstdllreturndll.png" />
    // Output: [1,2,3,4,5]
    // <img src="https://assets.leetcode.com/uploads/2018/10/12/bstdllreturnbst.png" />
    // Explanation: The figure below shows the transformed BST. The solid line indicates the successor relationship, while the dashed line means the predecessor relationship.
    tree1 := &Node {
        4,
        &Node{2, &Node{1, nil, nil}, &Node{3, nil, nil}, },
        &Node{5, nil, nil},
    }
    t1 := treeToDoublyList(tree1)
    fmt.Println("t1 ", t1)
    fmt.Println("t1.Left ", t1.Left)
    fmt.Println("t1.Right ", t1.Right)
    // Example 2:
    // Input: root = [2,1,3]
    // Output: [1,2,3]
    tree2 := &Node {
        2,
        &Node{1, nil, nil},
        &Node{3, nil, nil},
    }
    t2 := treeToDoublyList(tree2)
    fmt.Println("t2 ", t2)
    fmt.Println("t2.Left ", t2.Left)
    fmt.Println("t2.Right ", t2.Right)
}