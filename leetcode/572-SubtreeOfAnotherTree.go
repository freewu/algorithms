package main

// 572. Subtree of Another Tree
// Given the roots of two binary trees root and subRoot, 
// return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

// A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. 
// The tree tree could also be considered as a subtree of itself.

// Example 1:
//         root                  subRoot
//           3
//         /   \                  4
//       [4]    5                /  \
//      /  \                    1    2
//    [1]  [2]
// <img src="https://assets.leetcode.com/uploads/2021/04/28/subtree1-tree.jpg" />
// Input: root = [3,4,5,1,2], subRoot = [4,1,2]
// Output: true

// Example 2:
//        root               subRoot
//          3
//        /   \                 4
//       4     5               /  \
//      /  \                  1    2
//     1    2
//          /
//         0
// <img src="https://assets.leetcode.com/uploads/2021/04/28/subtree2-tree.jpg" />
// Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
// Output: false
 
// Constraints:
//     The number of nodes in the root tree is in the range [1, 2000].
//     The number of nodes in the subRoot tree is in the range [1, 1000].
//     -10^4 <= root.val <= 10^4
//     -10^4 <= subRoot.val <= 10^4

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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return subRoot == nil
    }
    var isSameTree func(p *TreeNode, q *TreeNode) bool
    isSameTree = func(p *TreeNode, q *TreeNode) bool {
        if p == nil && q == nil {
            return true
        }
        if p == nil || q == nil || p.Val != q.Val {
            return false
        }
        return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
    }
    if isSameTree(root, subRoot) {
        return true
    }
    return isSubtree(root.Right, subRoot) || isSubtree(root.Left, subRoot)
}

func isSubtree1(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }
    var same func (p, q *TreeNode) bool
    same = func (p, q *TreeNode) bool {
        if p == nil && q == nil {
            return true
        } else if p == nil || q == nil || p.Val != q.Val {
            return false
        }
        return same(p.Left, q.Left) && same(p.Right, q.Right)
    }
    return same(root, subRoot) || isSubtree1(root.Left, subRoot) || isSubtree1(root.Right, subRoot)
}

func main() {
    // Example 1:
    //         root                  subRoot
    //           3
    //         /   \                  4
    //       [4]    5                /  \
    //      /  \                    1    2
    //    [1]  [2]
    // <img src="https://assets.leetcode.com/uploads/2021/04/28/subtree1-tree.jpg" />
    // Input: root = [3,4,5,1,2], subRoot = [4,1,2]
    // Output: true
    root1 := &TreeNode {
        3,
        &TreeNode{
            4, 
            &TreeNode{1, nil, nil},
            &TreeNode{2, nil, nil},
        },
        &TreeNode{5, nil, nil},
    }
    sub1 := &TreeNode{
        4, 
        &TreeNode{1, nil, nil},
        &TreeNode{2, nil, nil},
    }
    fmt.Println(isSubtree(root1,sub1)) // true
    // Example 2:
    //        root               subRoot
    //          3
    //        /   \                 4
    //       4     5               /  \
    //      /  \                  1    2
    //     1    2
    //          /
    //         0
    // <img src="https://assets.leetcode.com/uploads/2021/04/28/subtree2-tree.jpg" />
    // Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
    // Output: false
    root2 := &TreeNode {
        3,
        &TreeNode{
            4, 
            &TreeNode{1, nil, nil},
            &TreeNode{2, &TreeNode{0, nil, nil}, nil},
        },
        &TreeNode{5, nil, nil},
    }
    sub2 := &TreeNode{
        4, 
        &TreeNode{1, nil, nil},
        &TreeNode{2, nil, nil},
    }
    fmt.Println(isSubtree(root2,sub2)) // false

    fmt.Println(isSubtree1(root1,sub1)) // true
    fmt.Println(isSubtree1(root2,sub2)) // false
}