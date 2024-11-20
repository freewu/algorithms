package main

// 1026. Maximum Difference Between Node and Ancestor
// Given the root of a binary tree, 
// find the maximum value v for which there exist different nodes a and b where v = |a.val - b.val| and a is an ancestor of b.
// A node a is an ancestor of b if either: any child of a is equal to b or any child of a is an ancestor of b.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/09/tmp-tree.jpg" />
// Input: root = [8,3,10,1,6,null,14,null,null,4,7,13]
// Output: 7
// Explanation: We have various ancestor-node differences, some of which are given below :
// |8 - 3| = 5
// |3 - 7| = 4
// |8 - 1| = 7
// |10 - 13| = 3
// Among all possible differences, the maximum value of 7 is obtained by |8 - 1| = 7.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/09/tmp-tree-1.jpg" />
// Input: root = [1,null,2,null,0,3]
// Output: 3
 
// Constraints:
//     The number of nodes in the tree is in the range [2, 5000].
//     0 <= Node.val <= 10^5

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
// dfs
func maxAncestorDiff(root *TreeNode) int {
    var dfs func(*TreeNode, int, int) int
    dfs = func(node *TreeNode, pmax, pmin int) int {
        if node == nil {
            return pmax - pmin
        }
        max := func(x, y int) int { if x > y { return x; }; return y; }
        min := func(x, y int) int { if x < y { return x; }; return y; }
        mx, mi := max(node.Val, pmax), min(node.Val, pmin)
        return max(dfs(node.Left, mx, mi), dfs(node.Right, mx, mi))
    }
    return dfs(root, root.Val, root.Val)
}

func maxAncestorDiff1(root *TreeNode) int {
    var dfs func(root *TreeNode, mi, ma int) int
    dfs = func(root *TreeNode, mi, ma int) int {
        if root == nil {
            return 0
        }
        max := func(x, y int) int { if x > y { return x; }; return y; }
        min := func(x, y int) int { if x < y { return x; }; return y; }
        abs := func(x int) int { if x < 0 { return -x; }; return x; }

        diff := max(abs(root.Val - mi), abs(root.Val - ma))
        mi, ma = min(mi, root.Val), max(ma, root.Val)
        diff = max(diff, dfs(root.Left, mi, ma))
        diff = max(diff, dfs(root.Right, mi, ma))
        return diff
    }
    return dfs(root, root.Val, root.Val)
}

func main() {
    tree1 := &TreeNode {
        8,
        &TreeNode {
            3,
            &TreeNode{1, nil, nil},
            &TreeNode{
                6, 
                &TreeNode{4, nil, nil},
                &TreeNode{7, nil, nil},
            },
        },
        &TreeNode {
            10,
            nil,
            &TreeNode{
                14, 
                &TreeNode{13, nil, nil},
                nil, 
            },
        },
    }
    fmt.Println(maxAncestorDiff(tree1)) // 7

    tree2 := &TreeNode {
        1,
        nil,
        &TreeNode {
            2,
            nil,
            &TreeNode{
                0, 
                &TreeNode{3, nil, nil},
                nil, 
            },
        },
    }
    fmt.Println(maxAncestorDiff(tree2)) // 3

    fmt.Println(maxAncestorDiff1(tree1)) // 7
    fmt.Println(maxAncestorDiff1(tree2)) // 3
}