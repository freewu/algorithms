package main 

// 102. Binary Tree Level Order Traversal
// Given the root of a binary tree, 
// return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg" />
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]

// Example 2:
// Input: root = [1]
// Output: [[1]]

// Example 3:
// Input: root = []
// Output: []

// Constraints:
//     The number of nodes in the tree is in the range [0, 2000].
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
// BFS
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    // 准备一个队列 queue 先把 root 放里面
    queue := []*TreeNode{root}
    res := make([][]int, 0)
    // 有多个层就会循环多少次
    for len(queue) > 0 {
        l := len(queue)
        tmp := make([]int, 0, l)
        for i := 0; i < l; i++ { // i < l 这里是关键
            // 左枝不为空 放到队列 queue 中
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            // 右枝不为空 放到队列 queue 中
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
            // 把值加入到 tmp 中
            tmp = append(tmp, queue[i].Val)
        }
        fmt.Println("queue: ",queue)
        fmt.Println("tmp: ",tmp)
        // 取出队
        queue = queue[l:] // 这里是关键
        res = append(res, tmp)
    }
    return res
}

// DFS
func levelOrder1(root *TreeNode) [][]int {
    var res [][]int
    var dfsLevel func(node *TreeNode, level int)
    dfsLevel = func(node *TreeNode, level int) {
        if node == nil {
            return
        }
        // 第一次进入需创建新的一行
        if len(res) == level {
            res = append(res, []int{node.Val})
        } else {
            res[level] = append(res[level], node.Val)
        }
        dfsLevel(node.Left, level+1)
        dfsLevel(node.Right, level+1)
    }
    dfsLevel(root, 0)
    return res
}

func main() {
    tree1 := &TreeNode {
        3,
        &TreeNode{9, nil, nil},
        &TreeNode {
            20,
            &TreeNode{15, nil, nil},
            &TreeNode{7, nil, nil},
        },
    }
    tree3 := &TreeNode {
        1,
        nil,
        nil,
    }
    fmt.Println(levelOrder(tree1)) // [[3],[9,20],[15,7]]
    fmt.Println(levelOrder(nil)) // []
    fmt.Println(levelOrder(tree3)) // [1]

    fmt.Println(levelOrder1(tree1)) // [[3],[9,20],[15,7]]
    fmt.Println(levelOrder1(nil)) // []
    fmt.Println(levelOrder1(tree3)) // [1]
}