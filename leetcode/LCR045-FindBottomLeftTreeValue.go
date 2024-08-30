package main

// LCR 045. 找树左下角的值
// 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
// 假设二叉树中至少有一个节点。

// 示例 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/14/tree1.jpg" />
// 输入: root = [2,1,3]
// 输出: 1

// 示例 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/14/tree2.jpg" />
// 输入: [1,2,3,4,null,5,6,null,null,7]
// 输出: 7
 
// 提示:
//     二叉树的节点个数的范围是 [1,10^4]
//     -2^31 <= Node.val <= 2^31 - 1 

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
    var dfs func(root *TreeNode, curHeight int) 
    dfs = func (root *TreeNode, curHeight int) {
        if curHeight > maxHeight && root.Left == nil && root.Right == nil { // 到达最深叶子
            maxHeight = curHeight
            res = root.Val
        }
        if root.Left != nil {
            dfs(root.Left, curHeight + 1)
        }
        if root.Right != nil {
            dfs(root.Right, curHeight + 1)
        }
    }
    dfs(root, 0)
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