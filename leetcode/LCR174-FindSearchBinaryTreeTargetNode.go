package main

// LCR 174. 寻找二叉搜索树中的目标节点
// 某公司组织架构以二叉搜索树形式记录，节点值为处于该职位的员工编号。请返回第 cnt 大的员工编号。

// 示例 1：
// <img src="https://pic.leetcode.cn/1695101634-kzHKZW-image.png" />
// 输入：root = [7, 3, 9, 1, 5], cnt = 2
//        7
//       / \
//      3   9
//     / \
//    1   5
// 输出：7

// 示例 2：
// <img src="https://pic.leetcode.cn/1695101636-ESZtLa-image.png" />
// 输入: root = [10, 5, 15, 2, 7, null, 20, 1, null, 6, 8], cnt = 4
//        10
//       / \
//      5   15
//     / \    \
//    2   7    20
//   /   / \ 
//  1   6   8
// 输出: 8

// 提示：
//     1 ≤ cnt ≤ 二叉搜索树元素个数

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
func findTargetNode(root *TreeNode, cnt int) int {
    stack := []*TreeNode{}
    // 正常中序遍历为 左中右（可以得到第cnt【小】的数值）
    // 本题变化为：中序遍历（右中左）（可以得到第cnt【大】的数值）
    for root != nil || len(stack) > 0 {
        if root != nil {
            stack = append(stack, root)
            root = root.Right // 正常这里是 root = root.Left
        } else {
            root = stack[len(stack)-1] // pop
            stack = stack[:len(stack)-1]
            cnt--
            if cnt == 0 {
                return root.Val
            }
            root = root.Left
        }
    }
    return 0
}

func findTargetNode1(root *TreeNode, cnt int) int {
    if root == nil {
        return 0
    }
    res, count := 0, 0
    var inorder func(root *TreeNode)
    inorder = func(root *TreeNode) {
        if root == nil {
            return
        }
        inorder(root.Right) // right 大
        count++
        if count == cnt {
            res = root.Val
            return
        }
        inorder(root.Left)
    }
    inorder(root)
    return res
}

func main() {
    // 示例 1：
    // <img src="https://pic.leetcode.cn/1695101634-kzHKZW-image.png" />
    // 输入：root = [7, 3, 9, 1, 5], cnt = 2
    //        7
    //       / \
    //      3   9
    //     / \
    //    1   5
    // 输出：7
    tree1 := &TreeNode {
        3,
        &TreeNode{3, &TreeNode{1, nil, nil, }, &TreeNode{5, nil, nil, }, },
        &TreeNode{9, nil, nil, },
    }
    fmt.Println(findTargetNode(tree1, 2)) // 7
    // 示例 2：
    // <img src="https://pic.leetcode.cn/1695101636-ESZtLa-image.png" />
    // 输入: root = [10, 5, 15, 2, 7, null, 20, 1, null, 6, 8], cnt = 4
    //        10
    //       / \
    //      5   15
    //     / \    \
    //    2   7    20
    //   /   / \ 
    //  1   6   8
    // 输出: 8
    tree2 := &TreeNode {
        10,
        &TreeNode{5,  &TreeNode{2, &TreeNode{1, nil, nil, }, nil, }, &TreeNode{7, &TreeNode{6, nil, nil, }, &TreeNode{8, nil, nil, }, }, },
        &TreeNode{15, nil, &TreeNode{20, nil, nil, }, },
    }
    fmt.Println(findTargetNode(tree2, 4)) // 8

    tree11 := &TreeNode {
        3,
        &TreeNode{3, &TreeNode{1, nil, nil, }, &TreeNode{5, nil, nil, }, },
        &TreeNode{9, nil, nil, },
    }
    fmt.Println(findTargetNode1(tree11, 2)) // 7
    tree12 := &TreeNode {
        10,
        &TreeNode{5,  &TreeNode{2, &TreeNode{1, nil, nil, }, nil, }, &TreeNode{7, &TreeNode{6, nil, nil, }, &TreeNode{8, nil, nil, }, }, },
        &TreeNode{15, nil, &TreeNode{20, nil, nil, }, },
    }
    fmt.Println(findTargetNode1(tree12, 4)) // 8
}