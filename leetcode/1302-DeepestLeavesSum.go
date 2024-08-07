package main

// 1302. Deepest Leaves Sum
// Given the root of a binary tree, return the sum of values of its deepest leaves.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/07/31/1483_ex1.png" />
// Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
// Output: 15

// Example 2:
// Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
// Output: 19

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     1 <= Node.val <= 100

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
// bfs
func deepestLeavesSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res, queue := 0, []*TreeNode{}
    queue = append(queue, root)
    for len(queue) > 0 {
        size := len(queue)
        res = 0
        for ; size > 0; size-- {
            node := queue[0]
            queue = queue[1:]
            if node.Left == nil && node.Right == nil {
                res += node.Val
            } else {
                if node.Left  != nil { queue = append(queue, node.Left) }
                if node.Right != nil { queue = append(queue, node.Right) }
            }
        }
    }
    return res
}

// dfs
func deepestLeavesSum1(root *TreeNode) int {
    h, res := 0,0
    var dfs func(root *TreeNode , x int) *TreeNode
    dfs = func(root *TreeNode , x int) *TreeNode {
        if(root == nil) {
            return nil
        }
        if(x > h) {
            h = x
            res = 0
        }
        left, right := dfs(root.Left ,x + 1), dfs(root.Right,x + 1)
        if( x == h && left == nil && right == nil) {
            res += root.Val
        }
        return root
    }
    dfs(root, 1)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/07/31/1483_ex1.png" />
    // Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
    // Output: 15
    tree1 := &TreeNode{
        1, 
        &TreeNode{2, &TreeNode{4, &TreeNode{7, nil, nil}, nil}, &TreeNode{5, nil, nil}},
        &TreeNode{3, nil, &TreeNode{6, nil, &TreeNode{8, nil, nil}, }, },
    }
    fmt.Println(deepestLeavesSum(tree1)) // 15
    // Example 2:
    // Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
    // Output: 19
    tree2 := &TreeNode{
        6, 
        &TreeNode{7, &TreeNode{2, &TreeNode{9, nil, nil}, nil}, &TreeNode{7, &TreeNode{1, nil, nil}, &TreeNode{4, nil, nil}}, },
        &TreeNode{8, &TreeNode{1, nil, nil}, &TreeNode{3, nil, &TreeNode{5, nil, nil}, }, },
    }
    fmt.Println(deepestLeavesSum(tree2)) // 19

    tree11 := &TreeNode{
        1, 
        &TreeNode{2, &TreeNode{4, &TreeNode{7, nil, nil}, nil}, &TreeNode{5, nil, nil}},
        &TreeNode{3, nil, &TreeNode{6, nil, &TreeNode{8, nil, nil}, }, },
    }
    fmt.Println(deepestLeavesSum1(tree11)) // 15
    tree12 := &TreeNode{
        6, 
        &TreeNode{7, &TreeNode{2, &TreeNode{9, nil, nil}, nil}, &TreeNode{7, &TreeNode{1, nil, nil}, &TreeNode{4, nil, nil}}, },
        &TreeNode{8, &TreeNode{1, nil, nil}, &TreeNode{3, nil, &TreeNode{5, nil, nil}, }, },
    }
    fmt.Println(deepestLeavesSum1(tree12)) // 19
}
