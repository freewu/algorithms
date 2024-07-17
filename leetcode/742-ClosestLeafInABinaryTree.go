package main

// 742. Closest Leaf in a Binary Tree
// Given the root of a binary tree where every node has a unique value and a target integer k, 
// return the value of the nearest leaf node to the target k in the tree.

// Nearest to a leaf means the least number of edges traveled on the binary tree to reach any leaf of the tree. 
// Also, a node is called a leaf if it has no children.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/13/closest1-tree.jpg" />
// Input: root = [1,3,2], k = 1
// Output: 2
// Explanation: Either 2 or 3 is the nearest leaf node to the target of 1.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/13/closest2-tree.jpg" />
// Input: root = [1], k = 1
// Output: 1
// Explanation: The nearest leaf node is the root node itself.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/06/13/closest3-tree.jpg" />
// Input: root = [1,2,3,4,null,null,null,5,null,6], k = 2
// Output: 3
// Explanation: The leaf node with value 3 (and not the leaf node with value 6) is nearest to the node with value 2.

// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     1 <= Node.val <= 1000
//     All the values of the tree are unique.
//     There exist some node in the tree where Node.val == k.

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
func findClosestLeaf(root *TreeNode, k int) int {
    res, deep, pre := 0, 1001, []*TreeNode{}
    var findK func(*TreeNode, int)
    findK = func(root *TreeNode, k int) {
        var pLast *TreeNode
        nd := root
        for nd != nil {
            pre = append(pre, nd)
            nd = nd.Left
        }
        for len(pre) > 0 {
            nd = pre[len(pre)-1]
            if nd.Right == nil || nd.Right == pLast {
                if nd.Val == k { return }
                pLast = nd
                pre = pre[:len(pre)-1]
                
            } else {
                nd = nd.Right
                for nd != nil {
                    pre = append(pre, nd)
                    nd = nd.Left
                }
                
            }
        }       
    }
    // find k node and its ancestor
    findK(root, k)
    var findTarget func(*TreeNode, int)
    findTarget = func(nd *TreeNode, level int) {
        if nd == nil { return }
        if nd.Left == nil && nd.Right == nil && level < deep {
            res = nd.Val
            deep = level
        }
        findTarget(nd.Left, level+1)
        findTarget(nd.Right, level+1)
    }
    for i := len(pre)-1; i >= 0; i-- {
        findTarget(pre[i], len(pre)-1-i)
    }
    return res
}


func findClosestLeaf1(root *TreeNode, k int) int {
    parentMap := make(map[*TreeNode]*TreeNode) // 找到目标节点并构建所需的子节点到父节点的映射
    // 遍历函数，找到值为 k 的那个节点，同时记录子节点到父节点的映射
    var traverse func(root *TreeNode, k int, parent *TreeNode, parentMap map[*TreeNode]*TreeNode) *TreeNode
    traverse = func (root *TreeNode, k int, parent *TreeNode, parentMap map[*TreeNode]*TreeNode) *TreeNode {
        if root == nil { return nil  }
        parentMap[root] = parent // 记录子节点到父节点的映射
        if root.Val == k { return root }
        left := traverse(root.Left, k, root, parentMap)
        if left != nil { return left }
        return traverse(root.Right, k, root, parentMap)
    }
    var bfs func(target *TreeNode, parentMap map[*TreeNode]*TreeNode) int 
    bfs = func(target *TreeNode, parentMap map[*TreeNode]*TreeNode) int {
        queue, visited := []*TreeNode{target}, make(map[*TreeNode]bool)
        visited[target] = true
        for len(queue) > 0 {
            size := len(queue)
            for i := 0; i < size; i++ {
                cur := queue[0]
                queue = queue[1:]
                if cur.Left == nil && cur.Right == nil { // 首次到达的叶子结点就是最近的叶子结点
                    return cur.Val
                }
                if cur.Left != nil && !visited[cur.Left] {
                    queue = append(queue, cur.Left)
                    visited[cur.Left] = true
                }
                if cur.Right != nil && !visited[cur.Right] {
                    queue = append(queue, cur.Right)
                    visited[cur.Right] = true
                }
                parentNode := parentMap[cur]
                if parentNode != nil && !visited[parentNode] {
                    queue = append(queue, parentNode)
                    visited[parentNode] = true
                }
            }
        }
        return -1
    }
    target := traverse(root, k, nil, parentMap) 
    return bfs(target, parentMap)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/13/closest1-tree.jpg" />
    // Input: root = [1,3,2], k = 1
    // Output: 2
    // Explanation: Either 2 or 3 is the nearest leaf node to the target of 1.
    tree1 := &TreeNode{
        1, 
        &TreeNode{3, nil, nil, },
        &TreeNode{2, nil, nil, },
    }
    fmt.Println(findClosestLeaf(tree1,1)) // 2 | 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/13/closest2-tree.jpg" />
    // Input: root = [1], k = 1
    // Output: 1
    // Explanation: The nearest leaf node is the root node itself.
    tree2 := &TreeNode{1, nil, nil, } 
    fmt.Println(findClosestLeaf(tree2,1)) // 1
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/06/13/closest3-tree.jpg" />
    // Input: root = [1,2,3,4,null,null,null,5,null,6], k = 2
    // Output: 3
    // Explanation: The leaf node with value 3 (and not the leaf node with value 6) is nearest to the node with value 2.
    tree3 := &TreeNode{
        1, 
        &TreeNode{2, &TreeNode{4, &TreeNode{5, &TreeNode{ 6, nil, nil, }, nil, }, nil, }, nil, },
        &TreeNode{3, nil, nil, },
    }
    fmt.Println(findClosestLeaf(tree3,2)) // 3

    tree11 := &TreeNode{
        1, 
        &TreeNode{3, nil, nil, },
        &TreeNode{2, nil, nil, },
    }
    fmt.Println(findClosestLeaf1(tree11,1)) // 2
    tree12 := &TreeNode{1, nil, nil, } 
    fmt.Println(findClosestLeaf1(tree12,1)) // 1
    tree13 := &TreeNode{
        1, 
        &TreeNode{2, &TreeNode{4, &TreeNode{5, &TreeNode{ 6, nil, nil, }, nil, }, nil, }, nil, },
        &TreeNode{3, nil, nil, },
    }
    fmt.Println(findClosestLeaf1(tree13,2)) // 3
}