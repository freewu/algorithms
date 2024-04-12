package main

// 107. Binary Tree Level Order Traversal II
// Given the root of a binary tree, 
// return the bottom-up level order traversal of its nodes' values. (i.e., from left to right, level by level from leaf to root).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg" /> 
// Input: root = [3,9,20,null,null,15,7]
// Output: [[15,7],[9,20],[3]]

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
// bfs
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    queue, res := []*TreeNode{root}, make([][]int, 0) // 准备一个队列 queue 先把 root 放里面  
    for len(queue) > 0 { // 有多个层就会循环多少次
        l := len(queue)
        tmp := make([]int, 0, l)
        for i := 0; i < l; i++ { // i < l 这里是关键
            if queue[i].Left != nil { // 左枝不为空 放到队列 queue 中
                queue = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil { // 右枝不为空 放到队列 queue 中
                queue = append(queue, queue[i].Right)
            }
            tmp = append(tmp, queue[i].Val) // 把值加入到 tmp 中
        }
        queue = queue[l:] // 取出队 这里是关键
        res = append(res, tmp)
    }
    // 反转数组
    reversalList := func (l [][]int) (res [][]int) {
        for i := len(l) - 1; i >= 0; i-- {
            res = append(res, l[i])
        }
        return res
    }
    return reversalList(res)
}

// dfs
func levelOrderBottom1(root *TreeNode) [][]int {
    var res [][]int
    var dfs func(node *TreeNode, level int)
    dfs = func(node *TreeNode, level int) {
        if node == nil {
            return
        }
        if len(res) == level { // 第一次进入需创建新的一行
            res = append(res, []int{node.Val})
        } else {
            res[level] = append(res[level], node.Val)
        }
        dfs(node.Left, level+1)
        dfs(node.Right, level+1)
    }
    dfs(root, 0)
    // 反转数组
    reversalList := func (l [][]int) (res [][]int) {
        for i := len(l) - 1; i >= 0; i-- {
            res = append(res, l[i])
        }
        return res
    }
    return reversalList(res)
}

func levelOrderBottom2(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    queue := []*TreeNode{}
    queue = append(queue, root)
    for len(queue) > 0 {
        level, size := []int{}, len(queue)  //  len(queue) 很重要 新加入的 queue 需要下次 for 再消费
        for i := 0; i < size; i++ {
            node := queue[0] // 队列取出 Dequeue
            queue = queue[1:]
            level = append(level, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        res = append(res, level)
    }
    // 首尾交换顺序
    for i, j := 0, len(res)-1; i < j; {
        res[i], res[j] = res[j], res[i]
        i++
        j--
    }
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
    fmt.Println(levelOrderBottom(tree1)) // [[15,7],[9,20],[3]]
    fmt.Println(levelOrderBottom(nil)) // []
    fmt.Println(levelOrderBottom(tree3)) // [1]

    fmt.Println(levelOrderBottom1(tree1)) // [[15,7],[9,20],[3]]
    fmt.Println(levelOrderBottom1(nil)) // []
    fmt.Println(levelOrderBottom1(tree3)) // [1]

    fmt.Println(levelOrderBottom2(tree1)) // [[15,7],[9,20],[3]]
    fmt.Println(levelOrderBottom2(nil)) // []
    fmt.Println(levelOrderBottom2(tree3)) // [1]
}