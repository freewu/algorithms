package main

// 2641. Cousins in Binary Tree II
// Given the root of a binary tree, replace the value of each node in the tree with the sum of all its cousins' values.
// Two nodes of a binary tree are cousins if they have the same depth with different parents.
// Return the root of the modified tree.
// Note that the depth of a node is the number of edges in the path from the root node to it.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2023/01/11/example11.png" />
// Input: root = [5,4,9,1,10,null,7]
// Output: [0,0,0,7,7,null,11]
// Explanation: 
// 		The diagram above shows the initial binary tree and the binary tree after changing the value of each node.
// 		- Node with value 5 does not have any cousins so its sum is 0.
// 		- Node with value 4 does not have any cousins so its sum is 0.
// 		- Node with value 9 does not have any cousins so its sum is 0.
// 		- Node with value 1 has a cousin with value 7 so its sum is 7.
// 		- Node with value 10 has a cousin with value 7 so its sum is 7.
// 		- Node with value 7 has cousins with values 1 and 10 so its sum is 11.
	
// Example 2:
// Input: root = [3,1,2]
// Output: [0,0,0]
// Explanation: 
// 		The diagram above shows the initial binary tree and the binary tree after changing the value of each node.
// 		- Node with value 3 does not have any cousins so its sum is 0.
// 		- Node with value 1 does not have any cousins so its sum is 0.
// 		- Node with value 2 does not have any cousins so its sum is 0.
 
// Constraints:
// 		The number of nodes in the tree is in the range [1, 10^5].
// 		1 <= Node.val <= 10^4

// 如果两个节点在树中有相同的深度且它们的父节点不同，那么它们互为 堂兄弟 

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
    children := []*TreeNode{root}
    root.Val = 0

    // Each iteration children list contains all elements with the same depth
    for len(children) > 0 {
        parents := children
        children = []*TreeNode{}

        // Find sum of all children in this depth
        sum := 0
        for _, parent := range parents {
            if parent.Left != nil { 
				sum += parent.Left.Val 
			}
            if parent.Right != nil { 
				sum += parent.Right.Val 
			} 
        }

        // Update children values:
        // Calculated sum minus the current childer values (if they exist)
        // Also add all children to the list for the next iteration (depth+1)
        for _, parent := range parents {
            reduce := 0
            if parent.Left != nil { 
				reduce += parent.Left.Val 
			}
            if parent.Right != nil { 
				reduce += parent.Right.Val 
			}
            
            if parent.Left != nil {
                parent.Left.Val = sum - reduce
                children = append(children, parent.Left)
            }
            if parent.Right != nil {
                parent.Right.Val = sum - reduce
                children = append(children, parent.Right)
            }
        }
    }
    return root
}