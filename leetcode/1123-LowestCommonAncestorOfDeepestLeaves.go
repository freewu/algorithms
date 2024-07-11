package main

// 1123. Lowest Common Ancestor of Deepest Leaves
// Given the root of a binary tree, return the lowest common ancestor of its deepest leaves.
// Recall that:
//     The node of a binary tree is a leaf if and only if it has no children
//     The depth of the root of the tree is 0. if the depth of a node is d, the depth of each of its children is d + 1.
//     The lowest common ancestor of a set S of nodes, is the node A with the largest depth such that every node in S is in the subtree with root A.

// Example 1:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/01/sketch1.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4]
// Output: [2,7,4]
// Explanation: We return the node with value 2, colored in yellow in the diagram.
// The nodes coloured in blue are the deepest leaf-nodes of the tree.
// Note that nodes 6, 0, and 8 are also leaf nodes, but the depth of them is 2, but the depth of nodes 7 and 4 is 3.

// Example 2:
// Input: root = [1]
// Output: [1]
// Explanation: The root is the deepest node in the tree, and it's the lca of itself.

// Example 3:
// Input: root = [0,1,3,null,2]
// Output: [2]
// Explanation: The deepest leaf node in the tree is 2, the lca of one node is itself.

// Constraints:
//     The number of nodes in the tree will be in the range [1, 1000].
//     0 <= Node.val <= 1000
//     The values of the nodes in the tree are unique.

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
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    var traverse func(root *TreeNode) (*TreeNode, int) 
    traverse = func (root *TreeNode) (*TreeNode, int) {
        if root == nil { return nil, 0 }
        p1, n1 := traverse(root.Left)
        p2, n2 := traverse(root.Right)
        if p1 == nil && p2 == nil { return root, 1 }
        if p1 == nil { return p2, n2 + 1 } // 右边最深
        if p2 == nil { return p1, n1 + 1 } // 左边最深

        if n1 > n2 {
            return p1, n1 + 1
        } else if n1 < n2 {
            return p2, n2 + 1
        } else {
            return root, n1 + 1
        }
    }
    res, _ := traverse(root)
    return res
}

func lcaDeepestLeaves1(root *TreeNode) *TreeNode {
    type Result struct {
        node  *TreeNode // record the recent common ancestor node
        depth int       // record the maximum depth of the binary tree with node as the root
    }
    var maxDepth func(root *TreeNode) Result 
    maxDepth = func(root *TreeNode) Result {
        if root == nil { return Result{nil, 0} }
        left := maxDepth(root.Left)
        right := maxDepth(root.Right)
        if left.depth == right.depth { // 当左右子树的最大深度相同时，这个根节点是新的最近公共祖先
            return Result{root, left.depth + 1}
        }
        res := left
        if right.depth > left.depth { // 左右子树的深度不同，则最近公共祖先在 depth 较大的一边
            res = right
        }
        res.depth++
        return res
    }
    res := maxDepth(root)
    return res.node
}

func main() {
    // Example 1:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/01/sketch1.png" />
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4]
    // Output: [2,7,4]
    // Explanation: We return the node with value 2, colored in yellow in the diagram.
    // The nodes coloured in blue are the deepest nodes of the tree.
    // Notice that nodes 5, 3 and 2 contain the deepest nodes in the tree but node 2 is the smallest subtree among them, so we return it.
    tree1 := &TreeNode {
        3,
        &TreeNode{5, &TreeNode{6, nil, nil, }, &TreeNode{2, &TreeNode{7, nil, nil, }, &TreeNode{4, nil, nil, }, }, },
        &TreeNode{1, &TreeNode{0, nil, nil, }, &TreeNode{8, nil,                      nil,                      }, },
    }
    fmt.Println(lcaDeepestLeaves(tree1)) // &{2 0xc0000080a8 0xc0000080c0}
    // Example 2:
    // Input: root = [1]
    // Output: [1]
    // Explanation: The root is the deepest node in the tree.
    tree2 := &TreeNode{1, nil, nil, }
    fmt.Println(lcaDeepestLeaves(tree2)) // &{1 <nil> <nil>}
    // Example 3:
    // Input: root = [0,1,3,null,2]
    // Output: [2]
    // Explanation: The deepest node in the tree is 2, the valid subtrees are the subtrees of nodes 2, 1 and 0 but the subtree of node 2 is the smallest.
    tree3 := &TreeNode{
                0, 
                &TreeNode{1, nil, &TreeNode{2, nil, nil, }, }, 
                &TreeNode{3, nil, nil, }, 
            }
    fmt.Println(lcaDeepestLeaves(tree3)) // &{2 <nil> <nil>}

    tree11 := &TreeNode {
        3,
        &TreeNode{5, &TreeNode{6, nil, nil, }, &TreeNode{2, &TreeNode{7, nil, nil, }, &TreeNode{4, nil, nil, }, }, },
        &TreeNode{1, &TreeNode{0, nil, nil, }, &TreeNode{8, nil,                      nil,                      }, },
    }
    fmt.Println(lcaDeepestLeaves1(tree11)) // &{2 0xc0000080a8 0xc0000080c0}
    tree12 := &TreeNode{1, nil, nil, }
    fmt.Println(lcaDeepestLeaves1(tree12)) // &{1 <nil> <nil>}
    tree13 := &TreeNode{
                0, 
                &TreeNode{1, nil, &TreeNode{2, nil, nil, }, }, 
                &TreeNode{3, nil, nil, }, 
            }
    fmt.Println(lcaDeepestLeaves1(tree13)) // &{2 <nil> <nil>}
}