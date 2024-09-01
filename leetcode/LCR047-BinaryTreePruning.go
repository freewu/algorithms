package main

// LCR 047. 二叉树剪枝
// 给定一个二叉树 根节点 root ，树的每个节点的值要么是 0，要么是 1。
// 请剪除该二叉树中所有节点的值为 0 的子树。

// 节点 node 的子树为 node 本身，以及所有 node 的后代。

// 示例 1:
// 输入: [1,null,0,0,1]
// 输出: [1,null,0,null,1] 
// 解释: 
// 只有红色节点满足条件“所有不包含 1 的子树”。
// 右图为返回的答案。
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/06/1028_2.png" />

// 示例 2:
// 输入: [1,0,1,0,0,0,1]
// 输出: [1,null,1,null,1]
// 解释: 
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/06/1028_1.png" />

// 示例 3:
// 输入: [1,1,0,1,1,0,1,0]
// 输出: [1,1,0,1,1,null,1]
// 解释: 
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/05/1028.png" />

// 提示:
//     二叉树的节点个数的范围是 [1,200]
//     二叉树节点的值只会是 0 或 1

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
func pruneTree1(root *TreeNode) *TreeNode {
    var dfs func(node *TreeNode) bool
    dfs = func(node *TreeNode) bool {
        if node == nil {
            return false
        }
        left := dfs(node.Left)
        if !left {
            node.Left = nil
        }
        right := dfs(node.Right) 
        if !right {
            node.Right = nil
        }
        return node.Val == 1 || left || right
    }
    if !dfs(root) {
        return nil
    }
    return root
}

func pruneTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    root.Left, root.Right = pruneTree(root.Left), pruneTree(root.Right)
    if root.Val == 0 && root.Left == nil && root.Right == nil { // 后序遍历位置，判断自己是否是值为 0 的叶子节点
        return nil
    }
    return root // 如果不是，正常返回
}

func main() {
    // Example 1:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/06/1028_2.png" />
    // Input: root = [1,null,0,0,1]
    // Output: [1,null,0,null,1]
    // Explanation: 
    // Only the red nodes satisfy the property "every subtree not containing a 1".
    // The diagram on the right represents the answer.
    tree1 := &TreeNode {
        1,
        nil,
        &TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}, },
    }
    fmt.Println(pruneTree(tree1)) 
    // Example 2:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/06/1028_1.png" />
    // Input: root = [1,0,1,0,0,0,1]
    // Output: [1,null,1,null,1]
    tree2 := &TreeNode {
        1,
        &TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}, },
        &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}, },
    }
    fmt.Println(pruneTree(tree2)) 
    // Example 3:
    // <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/05/1028.png" />
    // Input: root = [1,1,0,1,1,0,1,0]
    // Output: [1,1,0,1,1,null,1]
    tree3 := &TreeNode {
        1,
        &TreeNode{1, &TreeNode{1, &TreeNode{0, nil, nil}, nil}, &TreeNode{1, nil, nil}, },
        &TreeNode{0, &TreeNode{0, nil,                    nil}, &TreeNode{1, nil, nil}, },
    }
    fmt.Println(pruneTree(tree3)) 

    tree11 := &TreeNode {
        1,
        nil,
        &TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}, },
    }
    fmt.Println(pruneTree1(tree11)) 
    tree12 := &TreeNode {
        1,
        &TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}, },
        &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}, },
    }
    fmt.Println(pruneTree1(tree12)) 
    tree13 := &TreeNode {
        1,
        &TreeNode{1, &TreeNode{1, &TreeNode{0, nil, nil}, nil}, &TreeNode{1, nil, nil}, },
        &TreeNode{0, &TreeNode{0, nil,                    nil}, &TreeNode{1, nil, nil}, },
    }
    fmt.Println(pruneTree1(tree13)) 
}