package main

// 1602. Find Nearest Right Node in Binary Tree
// Given the root of a binary tree and a node u in the tree, 
// return the nearest node on the same level that is to the right of u, 
// or return null if u is the rightmost node in its level.

// Example 1:
//              1
//           /     \
//          2       3
//           \     /  \
//           (4)  5     6
// <img src="https://assets.leetcode.com/uploads/2020/09/24/p3.png" />
// Input: root = [1,2,3,null,4,5,6], u = 4
// Output: 5
// Explanation: The nearest node on the same level to the right of node 4 is node 5.

// Example 2:
//         3
//           \
//            4
//           /
//         (2)
// <img src="https://assets.leetcode.com/uploads/2020/09/23/p2.png" />
// Input: root = [3,null,4,2], u = 2
// Output: null
// Explanation: There are no nodes to the right of 2.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     1 <= Node.val <= 10^5
//     All values in the tree are distinct.
//     u is a node in the binary tree rooted at root.

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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
    // iter to find the target u
    height_stack := []int{0}
    iter_stack := []*TreeNode{ root }
    for len(iter_stack) > 0 {
        cur := iter_stack[0]
        height := height_stack[0]
        if cur.Val == u.Val {
            if len(height_stack) == 1 {
                return nil
            }
            next_h := height_stack[1]
            if height == next_h {
                return iter_stack[1]
            }
            break
        }
        iter_stack = iter_stack[1:]
        height_stack = height_stack[1:]
        if cur.Left != nil {
            iter_stack = append(iter_stack, cur.Left)
            height_stack = append(height_stack, height+1)
        }
        if cur.Right != nil {
            iter_stack = append(iter_stack, cur.Right)
            height_stack = append(height_stack, height+1)
        }
    }
    return nil
}

func findNearestRightNode1(root *TreeNode, u *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        levelSize := len(queue)
        for i := 0; i < levelSize; i++ {
            curr := queue[0]
            queue = queue[1:]
            if curr.Val == u.Val {
                if i == levelSize - 1 {
                    return nil
                }
                return queue[0]
            }
            if curr.Left != nil {
                queue = append(queue, curr.Left)
            }
            if curr.Right != nil {
                queue = append(queue, curr.Right)
            }
        }
    }
    return nil 
}

func main() {
    // Example 1:
    //              1
    //           /     \
    //          2       3
    //           \     /  \
    //           (4)  5     6
    // <img src="https://assets.leetcode.com/uploads/2020/09/24/p3.png" />
    // Input: root = [1,2,3,null,4,5,6], u = 4
    // Output: 5
    // Explanation: The nearest node on the same level to the right of node 4 is node 5.
    tree1 := &TreeNode { 
        1, 
        &TreeNode { 2, nil,                        &TreeNode { 4, nil, nil, }, }, 
        &TreeNode { 3, &TreeNode { 5, nil, nil, }, &TreeNode { 6, nil, nil, }, },
    }
    fmt.Println(findNearestRightNode(tree1,  &TreeNode { 4, nil, nil, })) // &{5 <nil> <nil>}
    // Example 2:
    //         3
    //           \
    //            4
    //           /
    //         (2)
    // <img src="https://assets.leetcode.com/uploads/2020/09/23/p2.png" />
    // Input: root = [3,null,4,2], u = 2
    // Output: null
    // Explanation: There are no nodes to the right of 2.
    tree2 := &TreeNode { 
        3, 
        nil,
        &TreeNode { 4, &TreeNode { 2, nil, nil, }, nil, },
    }
    fmt.Println(findNearestRightNode(tree2,  &TreeNode { 2, nil, nil, })) // nil

    fmt.Println(findNearestRightNode1(tree1,  &TreeNode { 4, nil, nil, })) // &{5 <nil> <nil>}
    fmt.Println(findNearestRightNode1(tree2,  &TreeNode { 2, nil, nil, })) // nil
}