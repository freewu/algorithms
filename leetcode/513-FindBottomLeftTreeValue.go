package main

// 513. Find Bottom Left Tree Value
// Given the root of a binary tree, return the leftmost value in the last row of the tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/14/tree1.jpg" />
// Input: root = [2,1,3]
// Output: 1

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/14/tree2.jpg" />
// Input: root = [1,2,3,4,null,5,6,null,null,7]
// Output: 7

// Constraints:
//         The number of nodes in the tree is in the range [1, 10^4].
//         -2^31 <= Node.val <= 2^31 - 1

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
// DFS
func findBottomLeftValue(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res, maxHeight := 0, -1
    var dfs func(root *TreeNode, curHeight int, res, maxHeight *int) 
    dfs = func (root *TreeNode, curHeight int, res, maxHeight *int) {
        if curHeight > *maxHeight && root.Left == nil && root.Right == nil { // 到达最深叶子
            *maxHeight = curHeight
            *res = root.Val
        }
        if root.Left != nil {
            dfs(root.Left, curHeight + 1, res, maxHeight)
        }
        if root.Right != nil {
            dfs(root.Right, curHeight + 1, res, maxHeight)
        }
    }
    dfs(root, 0, &res, &maxHeight)
    return res
}

// BFS
func findBottomLeftValue1(root *TreeNode) int {
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        next := []*TreeNode{}
        for _, node := range queue {
            if node.Left != nil {
                next = append(next, node.Left)
            }
            if node.Right != nil {
                next = append(next, node.Right)
            }
        }
        if len(next) == 0 {
            return queue[0].Val
        }
        queue = next
    }
    return 0
}

func main() {
    fmt.Println(findBottomLeftValue(
        &TreeNode {
            1,
            &TreeNode{1, nil, nil},
            &TreeNode{3, nil, nil},
        },
    )) // 1

    fmt.Println(findBottomLeftValue(
        &TreeNode {
            1,
            &TreeNode {
                2,
                &TreeNode{4, nil, nil},
                nil,
            },
            &TreeNode {
                3,
                &TreeNode{
                    5, 
                    &TreeNode{7, nil, nil}, 
                    nil,
                },
                &TreeNode{6, nil, nil},
            },
        },
    )) // 7

    fmt.Println(findBottomLeftValue(
        &TreeNode {
            1,
            &TreeNode {
                2,
                &TreeNode{4, nil, nil},
                nil,
            },
            &TreeNode {
                3,
                &TreeNode{
                    5, 
                    nil,
                    &TreeNode{7, nil, nil}, 
                },
                &TreeNode{6, nil, nil},
            },
        },
    )) // 7

    fmt.Println(findBottomLeftValue1(
        &TreeNode {
            1,
            &TreeNode{1, nil, nil},
            &TreeNode{3, nil, nil},
        },
    )) // 1

    fmt.Println(findBottomLeftValue1(
        &TreeNode {
            1,
            &TreeNode {
                2,
                &TreeNode{4, nil, nil},
                nil,
            },
            &TreeNode {
                3,
                &TreeNode{
                    5, 
                    &TreeNode{7, nil, nil}, 
                    nil,
                },
                &TreeNode{6, nil, nil},
            },
        },
    )) // 7

    fmt.Println(findBottomLeftValue1(
        &TreeNode {
            1,
            &TreeNode {
                2,
                &TreeNode{4, nil, nil},
                nil,
            },
            &TreeNode {
                3,
                &TreeNode{
                    5, 
                    nil,
                    &TreeNode{7, nil, nil}, 
                },
                &TreeNode{6, nil, nil},
            },
        },
    )) // 7
}