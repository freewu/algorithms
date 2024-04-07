package main

// 199. Binary Tree Right Side View
// Given the root of a binary tree, imagine yourself standing on the right side of it, 
// return the values of the nodes you can see ordered from top to bottom.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/14/tree.jpg" />
// Input: root = [1,2,3,null,5,null,4]
// Output: [1,3,4]

// Example 2:
// Input: root = [1,null,3]
// Output: [1,3]

// Example 3:
// Input: root = []
// Output: []

// Constraints:
//     The number of nodes in the tree is in the range [0, 100].
//     -100 <= Node.val <= 100

import "fmt"

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
func rightSideView(root *TreeNode) []int {
    if root == nil{
        return []int{}
    }
    res := []int{}
    queue := []*TreeNode{root}
    for len(queue) > 0{
        // its gonna append rightmost value to result array
        res = append(res, queue[len(queue)-1].Val)
        // there we need to delete nodes we have and append
        // their children left TO right that is important cuz 
        // right mos value have to be at the end of the queue
        l := len(queue)
        for i := 0; i < l; i++ {
            cur := queue[0]
            queue = queue[1:]
            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
    }
    return res
}

// dfs
func rightSideView1(root *TreeNode) []int {
    res := []int{}
    var dfs func(root *TreeNode, level int)
    dfs = func(root *TreeNode, level int) {
        if root == nil {
            return
        }
        if len(res) <= level {
            res = append(res, root.Val)
        }
        dfs(root.Right, level + 1)
        dfs(root.Left, level + 1)
    }
    dfs(root, 0)
    return res
}

func main() {
    tree1 := &TreeNode {
        1,
        &TreeNode { 
            2, 
            nil,
            &TreeNode{5, nil, nil},
        },
        &TreeNode {
            3,
            nil,
            &TreeNode{4, nil, nil},
        },
    }
    fmt.Println(rightSideView(tree1)) // [1,3,4]
    tree2 := &TreeNode {
        1,
        nil,
        &TreeNode{3, nil, nil},
    }
    fmt.Println(rightSideView(tree2)) // [1,3]

    tree3 := &TreeNode {}
    fmt.Println(rightSideView(tree3)) // []

    fmt.Println(rightSideView1(tree1)) // [1,3,4]
    fmt.Println(rightSideView1(tree2)) // [1,3]
    fmt.Println(rightSideView1(tree3)) // []
}