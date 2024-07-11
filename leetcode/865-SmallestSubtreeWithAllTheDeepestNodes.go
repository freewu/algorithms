package main

// 865. Smallest Subtree with all the Deepest Nodes
// Given the root of a binary tree, the depth of each node is the shortest distance to the root.
// Return the smallest subtree such that it contains all the deepest nodes in the original tree.
// A node is called the deepest if it has the largest depth possible among any node in the entire tree.
// The subtree of a node is a tree consisting of that node, plus the set of all descendants of that node.

// Example 1:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/01/sketch1.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4]
// Output: [2,7,4]
// Explanation: We return the node with value 2, colored in yellow in the diagram.
// The nodes coloured in blue are the deepest nodes of the tree.
// Notice that nodes 5, 3 and 2 contain the deepest nodes in the tree but node 2 is the smallest subtree among them, so we return it.

// Example 2:
// Input: root = [1]
// Output: [1]
// Explanation: The root is the deepest node in the tree.

// Example 3:
// Input: root = [0,1,3,null,2]
// Output: [2]
// Explanation: The deepest node in the tree is 2, the valid subtrees are the subtrees of nodes 2, 1 and 0 but the subtree of node 2 is the smallest.

// Constraints:
//     The number of nodes in the tree will be in the range [1, 500].
//     0 <= Node.val <= 500
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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
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

func subtreeWithAllDeepest1(root *TreeNode) *TreeNode {
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

// zh: 定义：输入一棵二叉树，返回该二叉树的最大深度以及最深叶子节点的最近公共祖先节点
// en: definition: given a binary tree, return the maximum depth of the tree and the recent common ancestor node of the deepest leaf nodes


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
    fmt.Println(subtreeWithAllDeepest(tree1)) // &{2 0xc0000080a8 0xc0000080c0}
    // Example 2:
    // Input: root = [1]
    // Output: [1]
    // Explanation: The root is the deepest node in the tree.
    tree2 := &TreeNode{1, nil, nil, }
    fmt.Println(subtreeWithAllDeepest(tree2)) // &{1 <nil> <nil>}
    // Example 3:
    // Input: root = [0,1,3,null,2]
    // Output: [2]
    // Explanation: The deepest node in the tree is 2, the valid subtrees are the subtrees of nodes 2, 1 and 0 but the subtree of node 2 is the smallest.
    tree3 := &TreeNode{
                0, 
                &TreeNode{1, nil, &TreeNode{2, nil, nil, }, }, 
                &TreeNode{3, nil, nil, }, 
            }
    fmt.Println(subtreeWithAllDeepest(tree3)) // &{2 <nil> <nil>}

    tree11 := &TreeNode {
        3,
        &TreeNode{5, &TreeNode{6, nil, nil, }, &TreeNode{2, &TreeNode{7, nil, nil, }, &TreeNode{4, nil, nil, }, }, },
        &TreeNode{1, &TreeNode{0, nil, nil, }, &TreeNode{8, nil,                      nil,                      }, },
    }
    fmt.Println(subtreeWithAllDeepest1(tree11)) // &{2 0xc0000080a8 0xc0000080c0}
    tree12 := &TreeNode{1, nil, nil, }
    fmt.Println(subtreeWithAllDeepest1(tree12)) // &{1 <nil> <nil>}
    tree13 := &TreeNode{
                0, 
                &TreeNode{1, nil, &TreeNode{2, nil, nil, }, }, 
                &TreeNode{3, nil, nil, }, 
            }
    fmt.Println(subtreeWithAllDeepest1(tree13)) // &{2 <nil> <nil>}
}