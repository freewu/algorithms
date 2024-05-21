package main

// 450. Delete Node in a BST
// Given a root node reference of a BST and a key, delete the node with the given key in the BST. 
// Return the root node reference (possibly updated) of the BST.

// Basically, the deletion can be divided into two stages:
//     Search for a node to remove.
//     If the node is found, delete the node.
 
// Example 1:
//         5                 5
//       /   \              /   \
//      3     6     =>     4     6 
//     /  \    \          /       \
//    2    4    8        2         8
// <img src="https://assets.leetcode.com/uploads/2020/09/04/del_node_1.jpg" />
// Input: root = [5,3,6,2,4,null,7], key = 3
// Output: [5,4,6,2,null,null,7]
// Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
// One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
// Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.

// Example 2:
// Input: root = [5,3,6,2,4,null,7], key = 0
// Output: [5,3,6,2,4,null,7]
// Explanation: The tree does not contain a node with value = 0.

// Example 3:
// Input: root = [], key = 0
// Output: []

// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     -10^5 <= Node.val <= 10^5
//     Each node has a unique value.
//     root is a valid binary search tree.
//     -10^5 <= key <= 10^5

// Follow up: Could you solve it with time complexity O(height of tree)?

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
func deleteNode(root *TreeNode, key int) *TreeNode {
    var dfs func(*TreeNode) *TreeNode
    dfs = func(node *TreeNode) *TreeNode {
        if node == nil { return nil }
        if node.Val == key {
            if node.Left == nil { return node.Right }
            if node.Right == nil { return node.Left }
            curr := node.Right
            for curr.Left != nil { 
                curr = curr.Left
            }
            curr.Left = node.Left // 移除节点
            return node.Right
        }
        if key > node.Val { // 二分找到要删除的节点
            node.Right = dfs(node.Right)
        } else {
            node.Left = dfs(node.Left)
        }
        return node
    }
    return dfs(root)
}

func deleteNode1(root *TreeNode, key int) *TreeNode {
    var dfs func (cur **TreeNode,key int)
    dfs = func (cur **TreeNode,key int) {
        if (*cur) == nil { return }
        if key < (*cur).Val {
            dfs(&((*cur).Left),key)
        } else if key == (*cur).Val {
            if (*cur).Left == nil && (*cur).Right == nil {
                (*cur) = nil
            } else if (*cur).Left == nil && (*cur).Right != nil {
                (*cur) = (*cur).Right
            } else if (*cur).Left != nil && (*cur).Right == nil {
                (*cur) = (*cur).Left
            } else {
                tmp := (*cur).Right
                for tmp.Left != nil{
                    tmp = tmp.Left
                }
                (*cur).Val = tmp.Val
                dfs(&((*cur).Right),tmp.Val)
            }
        } else {
            dfs(&((*cur).Right),key)
        }
    }
    dfs(&root,key)
    return root
}

func main() {
    // Example 1:
    //         5                 5
    //       /   \              /   \
    //      3     6     =>     4     6 
    //     /  \    \          /       \
    //    2    4    7        2         7
    // <img src="https://assets.leetcode.com/uploads/2020/09/04/del_node_1.jpg" />
    // Input: root = [5,3,6,2,4,null,7], key = 3
    // Output: [5,4,6,2,null,null,7]
    // Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
    // One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
    // Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.
    tree1 := &TreeNode {
        5,
        &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}, },
        &TreeNode{6, nil, &TreeNode{7, nil, nil}, },
    }
    fmt.Println("tree1.Left: ", tree1.Left) 
    t1 := deleteNode(tree1,3)
    fmt.Println("t1: ", t1) // t1:  &{5 0xc000008078 0xc000008090}
    fmt.Println("t1.Left: ", t1.Left) 
    // Example 2:
    // Input: root = [5,3,6,2,4,null,7], key = 0
    // Output: [5,3,6,2,4,null,7]
    // Explanation: The tree does not contain a node with value = 0.
    tree2 := &TreeNode {
        5,
        &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}, },
        &TreeNode{6, nil, &TreeNode{7, nil, nil}, },
    }
    fmt.Println("tree2: ", tree2) 
    t2 := deleteNode(tree2, 0)
    fmt.Println("t2: ", t2) 
    // Example 3:
    // Input: root = [], key = 0
    // Output: []
    t3 := deleteNode(nil, 0)
    fmt.Println("t3: ", t3) 


    tree11 := &TreeNode {
        5,
        &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}, },
        &TreeNode{6, nil, &TreeNode{7, nil, nil}, },
    }
    fmt.Println("tree11.Left: ", tree11.Left) 
    t11 := deleteNode1(tree11,3)
    fmt.Println("t11: ", t11) // t11:  &{5 0xc000008078 0xc000008090}
    fmt.Println("t11.Left: ", t11.Left) 
    tree12 := &TreeNode {
        5,
        &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}, },
        &TreeNode{6, nil, &TreeNode{7, nil, nil}, },
    }
    fmt.Println("tree12: ", tree12) 
    t12 := deleteNode1(tree12, 0)
    fmt.Println("t12: ", t12) 
    t13 := deleteNode1(nil, 0)
    fmt.Println("t13: ", t13) 
}