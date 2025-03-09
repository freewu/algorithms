package main

// 面试题 04.10. Check SubTree LCCI
// T1 and T2 are two very large binary trees. Create an algorithm to determine if T2 is a subtree of T1.

// A tree T2 is a subtree of T1 if there exists a node n in T1 such that the subtree of n is identical to T2. That is, if you cut off the tree at node n, the two trees would be identical.

// Note: This problem is slightly different from the original problem.

// Example1:
// Input: t1 = [1, 2, 3], t2 = [2]
// Output: true

// Example2:
// Input: t1 = [1, null, 2, 4], t2 = [3, 2]
// Output: false

// Note:
//     The node numbers of both tree are in [0, 20000].

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
func checkSubTree(root, subRoot *TreeNode) bool {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var getHeight func(root *TreeNode) int // 二叉树的最大深度
    getHeight = func(root *TreeNode) int {
        if root == nil { return 0 }
        leftH, rightH := getHeight(root.Left), getHeight(root.Right)
        return max(leftH, rightH) + 1
    }
    var isSameTree func(p, q *TreeNode) bool
    isSameTree = func(p, q *TreeNode) bool { // 是否相同的树
        if p == nil || q == nil { return p == q } // 必须都是 nil
        return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
    }
    high := getHeight(subRoot)
    var dfs func(*TreeNode) (int, bool) // 返回 node 的高度，以及是否找到了 subRoot
    dfs = func(node *TreeNode) (int, bool) {
        if node == nil { return 0, false }
        leftH, leftFound := dfs(node.Left)
        rightH, rightFound := dfs(node.Right)
        if leftFound || rightFound { return 0, true }
        nodeH := max(leftH, rightH) + 1
        return nodeH, (nodeH == high && isSameTree(node, subRoot))
    }
    _, found := dfs(root)
    return found
}

func main() {
    // Example1:
    // Input: t1 = [1, 2, 3], t2 = [2]
    // Output: true
    tree11 := &TreeNode {
        1,
        &TreeNode{2, nil, nil, },
        &TreeNode{3, nil, nil, },
    }
    tree12 := &TreeNode{2, nil, nil, }
    fmt.Println(checkSubTree(tree11, tree12)) // true
    // Example2:
    // Input: t1 = [1, null, 2, 4], t2 = [3, 2]
    // Output: false
    tree21 := &TreeNode {
        1,
        nil,
        &TreeNode{2, &TreeNode{4, nil, nil, }, nil, },
    }
    tree22 := &TreeNode {
        3,
        &TreeNode{2, nil, nil, },
        nil,
    }
    fmt.Println(checkSubTree(tree21, tree22)) // false
}