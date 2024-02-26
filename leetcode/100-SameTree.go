package main

// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg" />
// Input: p = [1,2,3], q = [1,2,3]
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/20/ex2.jpg" />
// Input: p = [1,2], q = [1,null,2]
// Output: false

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/12/20/ex3.jpg" />
// Input: p = [1,2,1], q = [1,1,2]
// Output: false
 
// Constraints:
//         The number of nodes in both trees is in the range [0, 100].
//         -10^4 <= Node.val <= 10^4

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTreeNode(t *TreeNode) {
	if nil == t {
		return
	}

	fmt.Println()
}

func makeNodeList(nums []int) *TreeNode {
	var n = &TreeNode{-1, nil, nil}
	// var b = &ListNode{-1, n}
	// for i := 0; i < len(nums); i++ {
	// 	n.Next = &ListNode{nums[i], nil}
	// 	n = n.Next
	// }
	// return b.Next.Next
	return n
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归判断
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
    } else if p == nil && q == nil {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isSameTree(
        &TreeNode {
            1,
            &TreeNode{2, nil, nil},
            &TreeNode{3, nil, nil},
        },
        &TreeNode {
            1,
            &TreeNode{2, nil, nil},
            &TreeNode{3, nil, nil},
        },
    )) // true

    fmt.Println(isSameTree(
        &TreeNode {
            1,
            &TreeNode{2, nil, nil},
            nil,
        },
        &TreeNode {
            1,
            nil,
            &TreeNode{2, nil, nil},
        },
    )) // false

    fmt.Println(isSameTree(
        &TreeNode {
            1,
            &TreeNode{2, nil, nil},
            &TreeNode{1, nil, nil},
        },
        &TreeNode {
            1,
            &TreeNode{1, nil, nil},
            &TreeNode{2, nil, nil},
        },
    )) // false
}
