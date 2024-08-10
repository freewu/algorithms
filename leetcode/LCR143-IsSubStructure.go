package main

// LCR 143. 子结构判断
// 给定两棵二叉树 tree1 和 tree2，判断 tree2 是否以 tree1 的某个节点为根的子树具有 相同的结构和节点值 。
// 注意，空树 不会是以 tree1 的某个节点为根的子树具有 相同的结构和节点值 。

// 示例 1：
// <img src="https://pic.leetcode.cn/1694684670-vwyIgY-two_tree.png" />
// 输入：tree1 = [1,7,5], tree2 = [6,1]
// 输出：false
// 解释：tree2 与 tree1 的一个子树没有相同的结构和节点值。

// 示例 2：
// <img src="https://pic.leetcode.cn/1694685602-myWXCv-two_tree_2.png" />
// 输入：tree1 = [3,6,7,1,8], tree2 = [6,1]
// 输出：true
// 解释：tree2 与 tree1 的一个子树拥有相同的结构和节点值。即 6 - > 1。

// 提示：
//     0 <= 节点个数 <= 10000

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
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if B == nil || A == nil {
        return false
    }
    return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func recur(a, b *TreeNode) bool {
    if b == nil { return true } // b 可以有
    if a == nil { return false } // a 必须有
    if a.Val == b.Val { return recur(a.Left, b.Left) && recur(a.Right, b.Right) }
    return false
}

func main() {
    // 示例 1：
    // <img src="https://pic.leetcode.cn/1694684670-vwyIgY-two_tree.png" />
    // 输入：tree1 = [1,7,5], tree2 = [6,1]
    // 输出：false
    // 解释：tree2 与 tree1 的一个子树没有相同的结构和节点值。
    tree11 := &TreeNode{
        1, 
        &TreeNode{7, nil, nil},
        &TreeNode{5, nil, nil},
    }
    tree12 := &TreeNode{
        6, 
        &TreeNode{1, nil, nil},
        nil,
    }
    fmt.Println(isSubStructure(tree11,tree12)) // false
    // 示例 2：
    // <img src="https://pic.leetcode.cn/1694685602-myWXCv-two_tree_2.png" />
    // 输入：tree1 = [3,6,7,1,8], tree2 = [6,1]
    // 输出：true
    // 解释：tree2 与 tree1 的一个子树拥有相同的结构和节点值。即 6 - > 1。
    tree21 := &TreeNode{
        3, 
        &TreeNode{6, &TreeNode{1, nil, nil}, &TreeNode{8, nil, nil}, },
        &TreeNode{7, nil, nil},
    }
    tree22 := &TreeNode{
        6, 
        &TreeNode{1, nil, nil},
        nil,
    }
    fmt.Println(isSubStructure(tree21,tree22)) // true

    tree31 := &TreeNode{
        10, 
        &TreeNode{12, &TreeNode{8, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{6,  &TreeNode{11, nil, nil}, nil},
    }
    tree32 := &TreeNode{
        10, 
        &TreeNode{12, &TreeNode{8, nil, nil}, nil, },
        &TreeNode{6,  nil, nil, },
    }
    fmt.Println(isSubStructure(tree31,tree32)) // true

    fmt.Println(isSubStructure(nil,nil)) // false
}