package main

import "fmt"

// 993. Cousins in Binary Tree
// Given the root of a binary tree with unique values and the values of two different nodes of the tree x and y, 
// return true if the nodes corresponding to the values x and y in the tree are cousins, or false otherwise.

// Two nodes of a binary tree are cousins if they have the same depth with different parents.
// Note that in a binary tree, the root node is at the depth 0, 
// and children of each depth k node are at the depth k + 1.

// Example 1:
// <img src = "https://assets.leetcode.com/uploads/2019/02/12/q1248-01.png" />
// Input: root = [1,2,3,4], x = 4, y = 3
// Output: false

// Example 2:
// <img src = "https://assets.leetcode.com/uploads/2019/02/12/q1248-02.png" />
// Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
// Output: true

// Example 3:
// <img src = "https://assets.leetcode.com/uploads/2019/02/12/q1248-03.png" />
// Input: root = [1,2,3,null,4], x = 2, y = 3
// Output: false

// Constraints:
//     The number of nodes in the tree is in the range [2, 100].
//     1 <= Node.val <= 100
//     Each node has a unique value.
//     x != y
//     x and y are exist in the tree.


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
// 解法一 递归
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	levelX, levelY := findLevel(root, x, 1), findLevel(root, y, 1)
	if levelX != levelY {
		return false
	}
	return !haveSameParents(root, x, y)
}

func findLevel(root *TreeNode, x, level int) int {
    if root == nil {
        return 0
    }
    if root.Val != x {
        leftLevel, rightLevel := findLevel(root.Left, x, level+1), findLevel(root.Right, x, level+1)
        if leftLevel == 0 {
            return rightLevel
        }
        return leftLevel
    }
    return level
}

func haveSameParents(root *TreeNode, x, y int) bool {
    if root == nil {
        return false
    }
    if (root.Left != nil && root.Right != nil && root.Left.Val == x && root.Right.Val == y) ||
        (root.Left != nil && root.Right != nil && root.Left.Val == y && root.Right.Val == x) {
        return true
    }
    return haveSameParents(root.Left, x, y) || haveSameParents(root.Right, x, y)
}

// 解法二 BFS
func isCousinsBFS(root *TreeNode, x int, y int) bool {
    if root == nil {
        return false
    }
    type mark struct {
        prev  int
        depth int
    }
    queue := []*TreeNode{root}
    visited := [101]*mark{}
    visited[root.Val] = &mark{prev: -1, depth: 1}

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        depth := visited[node.Val].depth
        if node.Left != nil {
            visited[node.Left.Val] = &mark{prev: node.Val, depth: depth + 1}
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            visited[node.Right.Val] = &mark{prev: node.Val, depth: depth + 1}
            queue = append(queue, node.Right)
        }
    }
    if visited[x] == nil || visited[y] == nil {
        return false
    }
    if visited[x].depth == visited[y].depth && visited[x].prev != visited[y].prev {
        return true
    }
    return false
}

// 解法三 DFS
func isCousinsDFS(root *TreeNode, x int, y int) bool {
	depth1, depth2, parent1, parent2 := 0, 0, 0, 0
    var dfsCousins func(root *TreeNode, val, depth, last int, parent, res *int) 
    dfsCousins = func(root *TreeNode, val, depth, last int, parent, res *int) {
        if root == nil {
            return
        }
        if root.Val == val {
            *res = depth
            *parent = last
            return
        }
        depth++
        dfsCousins(root.Left, val, depth, root.Val, parent, res)
        dfsCousins(root.Right, val, depth, root.Val, parent, res)
    }
	dfsCousins(root, x, 0, -1, &parent1, &depth1)
	dfsCousins(root, y, 0, -1, &parent2, &depth2)
	return depth1 > 1 && depth1 == depth2 && parent1 != parent2
}

func main() {
    tree1 := &TreeNode {
        1,
        &TreeNode { 2, &TreeNode{4, nil, nil},  nil, },
        &TreeNode { 3, nil,                     nil, },
    }
    tree2 := &TreeNode {
        1,
        &TreeNode { 2, nil, &TreeNode{4, nil, nil}, },
        &TreeNode { 3, nil, &TreeNode{5, nil, nil}, },
    }
    tree3 := &TreeNode {
        1,
        &TreeNode { 2, nil, &TreeNode{4, nil, nil}, },
        &TreeNode { 3, nil, nil,                    },
    }
    fmt.Println(isCousins(tree1,4,3)) // false
    fmt.Println(isCousins(tree2,5,4)) // true
    fmt.Println(isCousins(tree3,2,3)) // false

    fmt.Println(isCousinsBFS(tree1,4,3)) // false
    fmt.Println(isCousinsBFS(tree2,5,4)) // true
    fmt.Println(isCousinsBFS(tree3,2,3)) // false

    fmt.Println(isCousinsDFS(tree1,4,3)) // false
    fmt.Println(isCousinsDFS(tree2,5,4)) // true
    fmt.Println(isCousinsDFS(tree3,2,3)) // false
}