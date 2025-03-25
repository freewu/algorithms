package main

// 2313. Minimum Flips in Binary Tree to Get Result
// You are given the root of a binary tree with the following properties:
//     Leaf nodes have either the value 0 or 1, representing false and true respectively.
//     Non-leaf nodes have either the value 2, 3, 4, or 5, representing the boolean operations OR, AND, XOR, and NOT, respectively.

// You are also given a boolean result, which is the desired result of the evaluation of the root node.

// The evaluation of a node is as follows:
//     If the node is a leaf node, the evaluation is the value of the node, i.e. true or false.
//     Otherwise, evaluate the node's children and apply the boolean operation of its value with the children's evaluations.

// n one operation, you can flip a leaf node, which causes a false node to become true, and a true node to become false.

// Return the minimum number of operations that need to be performed such that the evaluation of root yields result. 
// It can be shown that there is always a way to achieve result.

// A leaf node is a node that has zero children.

// Note: NOT nodes have either a left child or a right child, but other non-leaf nodes have both a left child and a right child.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/06/20/operationstree.png" />
// Input: root = [3,5,4,2,null,1,1,1,0], result = true
// Output: 2
// Explanation:
// It can be shown that a minimum of 2 nodes have to be flipped to make the root of the tree
// evaluate to true. One way to achieve this is shown in the diagram above.

// Example 2:
// Input: root = [0], result = false
// Output: 0
// Explanation:
// The root of the tree already evaluates to false, so 0 nodes have to be flipped.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     0 <= Node.val <= 5
//     OR, AND, and XOR nodes have 2 children.
//     NOT nodes have 1 child.
//     Leaf nodes have a value of 0 or 1.
//     Non-leaf nodes have a value of 2, 3, 4, or 5.

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
func minimumFlips(root *TreeNode, result bool) int {
    type pair struct {
        node   *TreeNode
        status bool
    }
    dp := make(map[pair]int)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(*TreeNode, bool) int
    dfs = func(node *TreeNode, status bool) int {
        if node == nil { return 0 }
        p := pair{ node, status}
        if _, ok := dp[p]; ok { return dp[p] }

        if node.Val == 1 { // true
            if status {
                dp[p] = 0
            } else {
                dp[p] = 1
            }
            return dp[p]
        }
        if node.Val == 0 { // false
            if status {
                dp[p] = 1
            } else {
                dp[p] = 0
            }
            return dp[p]
        }
        cur := 0
        // 非叶节点的值为 2、3、4、5，分别表示布尔运算 OR, AND, XOR, NOT
        if status {
            switch node.Val {
            case 2: // OR
                cur = min(dfs(node.Left, true), dfs(node.Right, true))
            case 3: // AND
                cur = dfs(node.Left, true) + dfs(node.Right, true)
            case 4: // XOR
                cur = min(dfs(node.Left, false) + dfs(node.Right, true), dfs(node.Left, true)+dfs(node.Right, false))
            case 5: // NOT
                if node.Left != nil {
                    cur = dfs(node.Left, false)
                } else {
                    cur = dfs(node.Right, false)
                }
            }
        } else {
            switch node.Val {
            case 2: // OR
                cur = dfs(node.Left, false) + dfs(node.Right, false)
            case 3: // AND
                cur = min(dfs(node.Left, false), dfs(node.Right, false))
            case 4: // XOR
                cur = min(dfs(node.Left, false)+dfs(node.Right, false), dfs(node.Left, true)+dfs(node.Right, true))
            case 5: // NOT
                if node.Left != nil {
                    cur = dfs(node.Left, true)
                } else {
                    cur = dfs(node.Right, true)
                }
            }
        }
        dp[p] = cur
        return cur
    }
    return dfs(root, result)
}

func minimumFlips1(root *TreeNode, result bool) int {
    f := [2]map[*TreeNode]int{}
    f[0], f[1] = map[*TreeNode]int{}, map[*TreeNode]int{}
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(t *TreeNode, r int) int
    dfs = func(t *TreeNode, r int) int {
        if t == nil { return 0 }
        if t.Left == nil && t.Right == nil { // leaf
            if r == t.Val { return 0
            }
            return 1
        }
        if v, ok := f[r][t]; ok {
            return v
        }
        res := 0
        if t.Val == 2 { // OR
            if r == 0 {
                res = dfs(t.Left, 0) + dfs(t.Right, 0)
            } else {
                l0, l1 := dfs(t.Left, 0), dfs(t.Left, 1)
                r0, r1 := dfs(t.Right, 0), dfs(t.Right, 1)
                res = min(l0+r1, min(l1+r0, l1+r1))
            }
        } else if t.Val == 3 { // AND
            if r == 0 {
                l0, l1 := dfs(t.Left, 0), dfs(t.Left, 1)
                r0, r1 := dfs(t.Right, 0), dfs(t.Right, 1)
                res = min(l0+r1, min(l1+r0, l0+r0))
            } else {
                res = dfs(t.Left, 1)+dfs(t.Right, 1)
            }
        } else if t.Val == 4 { // XOR
            l0, l1 := dfs(t.Left, 0), dfs(t.Left, 1)
            r0, r1 := dfs(t.Right, 0), dfs(t.Right, 1)
            if r == 0 {
                res = min(l0+r0, l1+r1)
            } else {
                res = min(l0+r1, l1+r0)
            }
        } else if t.Val == 5 { // NOT
            if r == 0 {
                res =dfs(t.Left, 1) + dfs(t.Right, 1)
            } else {
                res = dfs(t.Left, 0) + dfs(t.Right, 0)
            }
        }
        f[r][t] = res
        return res
    }
    if result == true {
        return dfs(root, 1)
    }
    return dfs(root, 0)
}

func minimumFlips2(root *TreeNode, result bool) int {
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(root *TreeNode) (int, int)
    dfs = func(root *TreeNode) (int, int) {
        if root.Val < 2 {
            if root.Val == 0 {
                return 1, 0
            } else {
                return 0, 1
            }
        } else if root.Val == 5 {
            var l, r int
            if root.Left != nil {
                l, r = dfs(root.Left)
            } else {
                l, r = dfs(root.Right)
            }
            return r, l
        }
        l1, l2 := dfs(root.Left)
        r1, r2 := dfs(root.Right)
        if root.Val == 2 {
            return min(l1 + r1, min(l1 + r2, l2 + r1)), l2 + r2
        } else if root.Val == 3 {
            return l1 + r1, min(l1 + r2, min(l2 + r1, l2 + r2))
        } else {
            return min(l1 + r2, l2 + r1), min(l1 + r1, l2 + r2)
        }
    }
    a, b := dfs(root)
    if result {
        return a
    }
    return b
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/06/20/operationstree.png" />
    // Input: root = [3,5,4,2,null,1,1,1,0], result = true
    // Output: 2
    // Explanation:
    // It can be shown that a minimum of 2 nodes have to be flipped to make the root of the tree
    // evaluate to true. One way to achieve this is shown in the diagram above.
    tree1 := &TreeNode {
        3,
        &TreeNode { 5, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{0, nil, nil} },  nil, },
        &TreeNode { 4, &TreeNode{1, nil, nil},  &TreeNode{1, nil, nil}, },
    }
    fmt.Println(minimumFlips(tree1, true)) // 2
    // Example 2:
    // Input: root = [0], result = false
    // Output: 0
    // Explanation:
    // The root of the tree already evaluates to false, so 0 nodes have to be flipped.
    tree2 := &TreeNode{0, nil, nil}
    fmt.Println(minimumFlips(tree2, false)) // 0

    fmt.Println(minimumFlips1(tree1, true)) // 2
    fmt.Println(minimumFlips1(tree2, false)) // 0

    fmt.Println(minimumFlips2(tree1, true)) // 2
    fmt.Println(minimumFlips2(tree2, false)) // 0
}