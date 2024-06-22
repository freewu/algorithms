package main

// 563. Binary Tree Tilt
// Given the root of a binary tree, return the sum of every tree node's tilt.
// The tilt of a tree node is the absolute difference between the sum of all left subtree node values and all right subtree node values. If a node does not have a left child, then the sum of the left subtree node values is treated as 0. The rule is similar if the node does not have a right child.

// Example 1:
//          1            1
//        /   \   =>   /   \
//       2     3      0     0
// <img src="https://assets.leetcode.com/uploads/2020/10/20/tilt2.jpg" />
// Input: root = [1,2,3]
// Output: 1
// Explanation: 
// Tilt of node 2 : |0-0| = 0 (no children)
// Tilt of node 3 : |0-0| = 0 (no children)
// Tilt of node 1 : |2-3| = 1 (left subtree is just left child, so sum is 2; right subtree is just right child, so sum is 3)
// Sum of every tilt : 0 + 0 + 1 = 1

// Example 2:
//         4                    6
//       /   \                /   \
//      2     9     =>       2     7
//    /   \     \           /  \     \  
//   3     5     7         0    0     0
// <img src="https://assets.leetcode.com/uploads/2020/10/20/tilt2.jpg" />
// Input: root = [4,2,9,3,5,null,7]
// Output: 15
// Explanation: 
// Tilt of node 3 : |0-0| = 0 (no children)
// Tilt of node 5 : |0-0| = 0 (no children)
// Tilt of node 7 : |0-0| = 0 (no children)
// Tilt of node 2 : |3-5| = 2 (left subtree is just left child, so sum is 3; right subtree is just right child, so sum is 5)
// Tilt of node 9 : |0-7| = 7 (no left child, so sum is 0; right subtree is just right child, so sum is 7)
// Tilt of node 4 : |(3+5+2)-(9+7)| = |10-16| = 6 (left subtree values are 3, 5, and 2, which sums to 10; right subtree values are 9 and 7, which sums to 16)
// Sum of every tilt : 0 + 0 + 0 + 2 + 7 + 6 = 15

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2020/10/20/tilt3.jpg" />
// Input: root = [21,7,14,1,1,2,2,3,3]
// Output: 9

// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     -1000 <= Node.val <= 1000

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
func findTilt(root *TreeNode) int {
    res := 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left, right := dfs(node.Left), dfs(node.Right)
        res += abs(left - right)
        return left + right + node.Val
    }
    dfs(root)
    return res
}

// bfs
func findTilt1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    type Visit struct {
        Node *TreeNode
        Visited bool
    }
    res, subSum, stack := 0, make(map[*TreeNode]int), []*Visit{ &Visit{root, false} }
    subSum[nil] = 0
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    for len(stack) != 0 {
        node, visited := stack[len(stack) - 1].Node, stack[len(stack) - 1].Visited
        stack = stack[:len(stack) - 1]
        if visited {
            res += abs(subSum[node.Left] - subSum[node.Right])
            subSum[node] = node.Val + subSum[node.Left] + subSum[node.Right]
        } else {
            stack = append(stack, &Visit{node, true})
            if node.Left != nil  { stack = append(stack, &Visit{node.Left, false});  }
            if node.Right != nil { stack = append(stack, &Visit{node.Right, false}); }
        }
    }
    return res
}

func main() {
    // Example 1:
    //          1            1
    //        /   \   =>   /   \
    //       2     3      0     0
    // <img src="https://assets.leetcode.com/uploads/2020/10/20/tilt2.jpg" />
    // Input: root = [1,2,3]
    // Output: 1
    // Explanation: 
    // Tilt of node 2 : |0-0| = 0 (no children)
    // Tilt of node 3 : |0-0| = 0 (no children)
    // Tilt of node 1 : |2-3| = 1 (left subtree is just left child, so sum is 2; right subtree is just right child, so sum is 3)
    // Sum of every tilt : 0 + 0 + 1 = 1
    tree1 := &TreeNode {
        1,
        &TreeNode{2, nil, nil, },
        &TreeNode{3, nil, nil, },
    }
    fmt.Println(findTilt(tree1)) // 1
    // Example 2:
    //         4                    6
    //       /   \                /   \
    //      2     9     =>       2     7
    //    /   \     \           /  \     \  
    //   3     5     7         0    0     0
    // <img src="https://assets.leetcode.com/uploads/2020/10/20/tilt2.jpg" />
    // Input: root = [4,2,9,3,5,null,7]
    // Output: 15
    // Explanation: 
    // Tilt of node 3 : |0-0| = 0 (no children)
    // Tilt of node 5 : |0-0| = 0 (no children)
    // Tilt of node 7 : |0-0| = 0 (no children)
    // Tilt of node 2 : |3-5| = 2 (left subtree is just left child, so sum is 3; right subtree is just right child, so sum is 5)
    // Tilt of node 9 : |0-7| = 7 (no left child, so sum is 0; right subtree is just right child, so sum is 7)
    // Tilt of node 4 : |(3+5+2)-(9+7)| = |10-16| = 6 (left subtree values are 3, 5, and 2, which sums to 10; right subtree values are 9 and 7, which sums to 16)
    // Sum of every tilt : 0 + 0 + 0 + 2 + 7 + 6 = 15
    tree2 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{3, nil, nil, }, &TreeNode{5, nil, nil, }, },
        &TreeNode{9, nil,                      &TreeNode{7, nil, nil, }, },
    }
    fmt.Println(findTilt(tree2)) // 15
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2020/10/20/tilt3.jpg" />
    // Input: root = [21,7,14,1,1,2,2,3,3]
    // Output: 9
    tree3 := &TreeNode {
        21,
        &TreeNode{7,  &TreeNode{1, &TreeNode{3, nil, nil, }, &TreeNode{3, nil, nil, }, }, &TreeNode{1, nil, nil, }, },
        &TreeNode{14, &TreeNode{2, nil,                      nil,                      }, &TreeNode{2, nil, nil, }, },
    }
    fmt.Println(findTilt(tree3)) // 9

    fmt.Println(findTilt1(tree1)) // 1
    fmt.Println(findTilt1(tree2)) // 15
    fmt.Println(findTilt1(tree3)) // 9
}