package main

// LCP 52. 二叉搜索树染色
// 欢迎各位勇者来到力扣城，本次试炼主题为「二叉搜索树染色」。

// 每位勇士面前设有一个二叉搜索树的模型，模型的根节点为 root，树上的各个节点值均不重复。
// 初始时，所有节点均为蓝色。现
// 在按顺序对这棵二叉树进行若干次操作， ops[i] = [type, x, y] 表示第 i 次操作为：
//     1. type 等于 0 时，将节点值范围在 [x, y] 的节点均染蓝
//     2. type 等于 1 时，将节点值范围在 [x, y] 的节点均染红

// 请返回完成所有染色后，该二叉树中红色节点的数量。

// 注意：
//     题目保证对于每个操作的 x、y 值定出现在二叉搜索树节点中

// 示例 1：
// 输入：root = [1,null,2,null,3,null,4,null,5], ops = [[1,2,4],[1,1,3],[0,3,5]]
// 输出：2
// 解释：  第 0 次操作，将值为 2、3、4 的节点染红； 
//        第 1 次操作，将值为 1、2、3 的节点染红； 
//        第 2 次操作，将值为 3、4、5 的节点染蓝； 
//        因此，最终值为 1、2 的节点为红色节点，返回数量 2
//        <img src="https://pic.leetcode-cn.com/1649833948-arSlXd-image.png" />

// 示例 2：
// 输入：root = [4,2,7,1,null,5,null,null,null,null,6] ops = [[0,2,2],[1,1,5],[0,4,5],[1,5,7]]
// 输出：5
// 解释： 第 0 次操作，将值为 2 的节点染蓝； 
//       第 1 次操作，将值为 1、2、4、5 的节点染红； 
//       第 2 次操作，将值为 4、5 的节点染蓝； 
//       第 3 次操作，将值为 5、6、7 的节点染红； 
//       因此，最终值为 1、2、5、6、7 的节点为红色节点，返回数量 5
//       <img src="https://pic.leetcode-cn.com/1649833763-BljEbP-image.png" />

// 提示：
//     1 <= 二叉树节点数量 <= 10^5
//     1 <= ops.length <= 10^5
//     ops[i].length == 3
//     ops[i][0] 仅为 0 or 1
//     0 <= ops[i][1] <= ops[i][2] <= 10^9
//     0 <= 节点值 <= 10^9

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
// dfs
func getNumber(root *TreeNode, ops [][]int) int {
    res, n := 0, len(ops)
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil { return }
        for i := n - 1; i >= 0; i-- {
            if root.Val >= ops[i][1] && root.Val <= ops[i][2] {
                res += ops[i][0]
                break
            }
        }
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)
    return res
}

// bfs
func getNumber1(root *TreeNode, ops [][]int) int {
    res, queue := 0, []*TreeNode{ root }
    for len(queue) > 0 {
        cur := queue[len(queue) - 1]
        queue = queue[:len(queue) - 1]
        if cur.Left != nil {
            queue = append(queue, cur.Left)
        }
        if cur.Right != nil {
            queue = append(queue, cur.Right)
        }
        v := cur.Val
        for i := len(ops) - 1; i >= 0; i-- {
            if ops[i][1] <= v && v <= ops[i][2] {
                if ops[i][0] == 1 {
                    res++
                }
                break
            }
        }
    }
    return res
}

func main() {
    // 示例 1：
    // 输入：root = [1,null,2,null,3,null,4,null,5], ops = [[1,2,4],[1,1,3],[0,3,5]]
    // 输出：2
    // 解释：  第 0 次操作，将值为 2、3、4 的节点染红； 
    //        第 1 次操作，将值为 1、2、3 的节点染红； 
    //        第 2 次操作，将值为 3、4、5 的节点染蓝； 
    //        因此，最终值为 1、2 的节点为红色节点，返回数量 2
    //        <img src="https://pic.leetcode-cn.com/1649833948-arSlXd-image.png" />
    tree1 := &TreeNode {
        1,
        nil,
        &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, &TreeNode{5, nil, nil, }, }, }, },
    }
    fmt.Println(getNumber(tree1, [][]int{{1,2,4},{1,1,3},{0,3,5}})) // 2
    // 示例 2：
    // 输入：root = [4,2,7,1,null,5,null,null,null,null,6] ops = [[0,2,2],[1,1,5],[0,4,5],[1,5,7]]
    // 输出：5
    // 解释： 第 0 次操作，将值为 2 的节点染蓝； 
    //       第 1 次操作，将值为 1、2、4、5 的节点染红； 
    //       第 2 次操作，将值为 4、5 的节点染蓝； 
    //       第 3 次操作，将值为 5、6、7 的节点染红； 
    //       因此，最终值为 1、2、5、6、7 的节点为红色节点，返回数量 5
    //       <img src="https://pic.leetcode-cn.com/1649833763-BljEbP-image.png" />
    tree2 := &TreeNode {
        4,
        &TreeNode{2,  &TreeNode{1, nil, nil, }, nil, },
        &TreeNode{7,  &TreeNode{5, nil,  &TreeNode{6, nil, nil, }, }, nil, },
    }
    fmt.Println(getNumber(tree2, [][]int{{0,2,2},{1,1,5},{0,4,5},{1,5,7}})) // 5

    tree11 := &TreeNode {
        1,
        nil,
        &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, &TreeNode{5, nil, nil, }, }, }, },
    }
    fmt.Println(getNumber1(tree11, [][]int{{1,2,4},{1,1,3},{0,3,5}})) // 2
    tree12 := &TreeNode {
        4,
        &TreeNode{2,  &TreeNode{1, nil, nil, }, nil, },
        &TreeNode{7,  &TreeNode{5, nil,  &TreeNode{6, nil, nil, }, }, nil, },
    }
    fmt.Println(getNumber1(tree12, [][]int{{0,2,2},{1,1,5},{0,4,5},{1,5,7}})) // 5
}