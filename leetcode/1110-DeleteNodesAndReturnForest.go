package main

// 1110. Delete Nodes And Return Forest
// Given the root of a binary tree, each node in the tree has a distinct value.
// After deleting all nodes with a value in to_delete, we are left with a forest (a disjoint union of trees).
// Return the roots of the trees in the remaining forest. You may return the result in any order.

// Example 1:
//             1
//           /   \
//          2     3
//        /   \  /  \
//       4    5 6    7
// <img src="https://assets.leetcode.com/uploads/2019/07/01/screen-shot-2019-07-01-at-53836-pm.png" />
// Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
// Output: [[1,2,null,4],[6],[7]]

// Example 2:
// Input: root = [1,2,4,null,3], to_delete = [3]
// Output: [[1,2,4]]

// Constraints:
//     The number of nodes in the given tree is at most 1000.
//     Each node has a distinct value between 1 and 1000.
//     to_delete.length <= 1000
//     to_delete contains distinct values between 1 and 1000.

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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    deleteNodes := make(map[int]bool, len(to_delete))
    for _, d := range to_delete {
        deleteNodes[d] = true
    }
    res, queue := []*TreeNode{}, []*TreeNode{root}
    var dfs func(root *TreeNode) *TreeNode
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        dfs = func(root *TreeNode) *TreeNode {
            if root == nil {
                return nil
            }
            if deleteNodes[root.Val] {
                if root.Left != nil  { queue = append(queue, root.Left)  }
                if root.Right != nil { queue = append(queue, root.Right) }
                delete(deleteNodes,root.Val)
                return nil
            }
            root.Left,root.Right = dfs(root.Left), dfs(root.Right)
            return root
        }
        node = dfs(node)
        if node != nil {
            res = append(res, node)
        }
    }
    return res
}

func main() {
    // Example 1:
    //             1
    //           /   \
    //          2     3
    //        /   \  /  \
    //       4    5 6    7
    // <img src="https://assets.leetcode.com/uploads/2019/07/01/screen-shot-2019-07-01-at-53836-pm.png" />
    // Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
    // Output: [[1,2,null,4],[6],[7]]
    tree1 := &TreeNode{
        1, 
        &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}, },
        &TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}, },
    }
    fmt.Println(delNodes(tree1,[]int{3,5}))
    // Example 2:
    // Input: root = [1,2,4,null,3], to_delete = [3]
    // Output: [[1,2,4]]
    tree2 := &TreeNode{
        1, 
        &TreeNode{2, nil, &TreeNode{3, nil, nil}, },
        &TreeNode{4, nil, nil },
    }
    fmt.Println(delNodes(tree2,[]int{3}))
}