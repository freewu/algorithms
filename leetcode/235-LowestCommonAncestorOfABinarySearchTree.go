package main

// 235. Lowest Common Ancestor of a Binary Search Tree
// Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.
// According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png" />
// Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
// Output: 6
// Explanation: The LCA of nodes 2 and 8 is 6.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png" />
// Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
// Output: 2
// Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

// Example 3:
// Input: root = [2,1], p = 2, q = 1
// Output: 2
 
// Constraints:
//         The number of nodes in the tree is in the range [2, 10^5].
//         -10^9 <= Node.val <= 10^9
//         All Node.val are unique.
//         p != q
//         p and q will exist in the BST.

// 公共祖先
// 对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
// 满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。

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
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if p == nil || q == nil || root == nil {
		return nil
	}
    // 满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

// 迭代
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    var curr *TreeNode = root
    for curr != nil {
        if p.Val < curr.Val && q.Val < curr.Val {
            curr = curr.Left
            continue
        }
        if p.Val > curr.Val && q.Val > curr.Val {
            curr = curr.Right
            continue
        }
        return curr
    }
    return curr
}

// best solution
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	for root != nil {
        if root.Val < p.Val && root.Val < q.Val {
            root = root.Right
        } else if root.Val > p.Val && root.Val > q.Val {
            root = root.Left
        } else {
            return root
        }
    }
    return nil
}

func main() {
    tree1 := &TreeNode {
		6,
		&TreeNode {
			2,
			&TreeNode{0, nil, nil},
			&TreeNode{
                4, 
                &TreeNode{3, nil, nil},
                &TreeNode{5, nil, nil},
            },
		},
        &TreeNode{
            8, 
            &TreeNode{7, nil, nil},
            &TreeNode{9, nil, nil},
        },
	}
	tree3 := &TreeNode {
		2,
		&TreeNode {
			1,
			nil,
			nil,
		},
		nil,
	}
    
    fmt.Println(lowestCommonAncestor(tree1,&TreeNode{2, nil, nil},&TreeNode{8, nil, nil})) // 6
    fmt.Println(lowestCommonAncestor(tree1,&TreeNode{2, nil, nil},&TreeNode{4, nil, nil})) // 2
    fmt.Println(lowestCommonAncestor(tree3,&TreeNode{2, nil, nil},&TreeNode{1, nil, nil})) // 2

    fmt.Println(lowestCommonAncestor1(tree1,&TreeNode{2, nil, nil},&TreeNode{8, nil, nil})) // 6
    fmt.Println(lowestCommonAncestor1(tree1,&TreeNode{2, nil, nil},&TreeNode{4, nil, nil})) // 2
    fmt.Println(lowestCommonAncestor1(tree3,&TreeNode{2, nil, nil},&TreeNode{1, nil, nil})) // 2

    fmt.Println(lowestCommonAncestor2(tree1,&TreeNode{2, nil, nil},&TreeNode{8, nil, nil})) // 6
    fmt.Println(lowestCommonAncestor2(tree1,&TreeNode{2, nil, nil},&TreeNode{4, nil, nil})) // 2
    fmt.Println(lowestCommonAncestor2(tree3,&TreeNode{2, nil, nil},&TreeNode{1, nil, nil})) // 2
}