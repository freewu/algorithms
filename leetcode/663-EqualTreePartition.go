package main

// 663. Equal Tree Partition
// Given the root of a binary tree, 
// return true if you can partition the tree into two trees with equal sums of values 
// after removing exactly one edge on the original tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/03/split1-tree.jpg" />
// Input: root = [5,10,10,null,null,2,3]
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/05/03/split2-tree.jpg" />
// Input: root = [1,2,10,null,null,2,20]
// Output: false
// Explanation: You cannot split the tree into two trees with equal sums after removing exactly one edge on the tree.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -10^5 <= Node.val <= 10^5

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func checkEqualTree(root *TreeNode) bool {
    m := map[int]int{} // 求 root 的和的时候，就可以把中间的和记录下来。所以解决方案就是找到是否有节点和为n/2的。
    var cal func(*TreeNode) int
    cal = func(n *TreeNode) int {
        if nil == n {
            return 0
        }
        v := n.Val + cal(n.Left) + cal(n.Right)
        m[v]++
        return v
    }
    s := cal(root)
    if 0 == s {
        return m[0] > 1
    }
    return s&1 == 0 && m[s >> 1] != 0 // 两颗树相等 -> 一个树的和是 n/2，
}

func checkEqualTree1(root *TreeNode) bool {
    var dfs func(root *TreeNode)int
    mp := make(map[int]bool)
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        v := dfs(node.Left) + dfs(node.Right) + node.Val
        mp[v] = true
        return v
    }
    sum := dfs(root.Left) + dfs(root.Right) + root.Val
    if sum % 2 != 0 {
        return false
    }
    _, ok := mp[sum / 2] // 两颗树相等 -> 一个树的和是 n/2，
    return ok
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/split1-tree.jpg" />
    // Input: root = [5,10,10,null,null,2,3]
    // Output: true
    tree1 := &TreeNode {
        5,
        &TreeNode{10, nil,                      nil,                      },
        &TreeNode{10, &TreeNode{2, nil, nil, }, &TreeNode{3, nil, nil, }, },
    }
    fmt.Println(checkEqualTree(tree1)) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/05/03/split2-tree.jpg" />
    // Input: root = [1,2,10,null,null,2,20]
    // Output: false
    // Explanation: You cannot split the tree into two trees with equal sums after removing exactly one edge on the tree.
    tree2 := &TreeNode {
        1,
        &TreeNode{2,  nil,                      nil,                       },
        &TreeNode{10, &TreeNode{2, nil, nil, }, &TreeNode{20, nil, nil, }, },
    }
    fmt.Println(checkEqualTree(tree2)) // false

    fmt.Println(checkEqualTree1(tree1)) // true
    fmt.Println(checkEqualTree1(tree2)) // false
}