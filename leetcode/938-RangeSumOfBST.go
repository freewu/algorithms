package main

// 938. Range Sum of BST
// Given the root node of a binary search tree and two integers low and high, 
// return the sum of values of all nodes with a value in the inclusive range [low, high].

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/05/bst1.jpg" />
// Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
// Output: 32
// Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/05/bst2.jpg" />
// Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
// Output: 23
// Explanation: Nodes 6, 7, and 10 are in the range [6, 10]. 6 + 7 + 10 = 23.
 
// Constraints:
//         The number of nodes in the tree is in the range [1, 2 * 10^4].
//         1 <= Node.val <= 10^5
//         1 <= low <= high <= 10^5
//         All Node.val are unique.

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
// 递归前序遍历
func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0
	preOrder(root, low, high, &res)
	return res
}

func preOrder(root *TreeNode, low, high int, res *int) {
	if root == nil {
		return
	}
    // Val 为 [low,high] 的值累加到 res
	if low <= root.Val && root.Val <= high {
		*res += root.Val
	}
	preOrder(root.Left, low, high, res)
	preOrder(root.Right, low, high, res)
}

// best solution
func rangeSumBST1(root *TreeNode, low int, high int) int {
    if root==nil {
        return 0
    }
    if root.Val < low {
        return rangeSumBST(root.Right,low,high)
    } else if root.Val == low {
        return rangeSumBST(root.Right,low,high) + root.Val
    } else if root.Val < high {
        return rangeSumBST(root.Right,low,high) + root.Val + rangeSumBST(root.Left,low,high)
    } else if root.Val == high {
        return rangeSumBST(root.Left,low,high) + root.Val
    } else {
        return rangeSumBST(root.Left,low,high)
    }
}

 func main() {
    tree1 := &TreeNode {
		10,
        &TreeNode {
			5,
			&TreeNode{3, nil, nil},
			&TreeNode{7, nil, nil},
		},
		&TreeNode {
			15,
			nil,
            &TreeNode{18, nil, nil},
		},
	}
    tree2 := &TreeNode {
		10,
        &TreeNode {
			5,
			&TreeNode{
                3, 
                &TreeNode{1, nil, nil},
                nil,
            },
			&TreeNode{
                7, 
                &TreeNode{6, nil, nil},
                nil,
            },
		},
		&TreeNode {
			15,
			&TreeNode{13, nil, nil},
            &TreeNode{18, nil, nil},
		},
	}

    fmt.Println(rangeSumBST(tree1,7,15)) // 32
    fmt.Println(rangeSumBST(tree2,6,10)) // 23

    fmt.Println(rangeSumBST1(tree1,7,15)) // 32
    fmt.Println(rangeSumBST1(tree2,6,10)) // 23
 }