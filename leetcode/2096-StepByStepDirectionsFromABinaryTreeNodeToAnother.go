package main

// 2096. Step-By-Step Directions From a Binary Tree Node to Another
// You are given the root of a binary tree with n nodes. 
// Each node is uniquely assigned a value from 1 to n. 
// You are also given an integer startValue representing the value of the start node s, 
// and a different integer destValue representing the value of the destination node t.

// Find the shortest path starting from node s and ending at node t. 
// Generate step-by-step directions of such path as a string consisting of only the uppercase letters 'L', 'R', and 'U'. 
// Each letter indicates a specific direction:
//     'L' means to go from a node to its left child node.
//     'R' means to go from a node to its right child node.
//     'U' means to go from a node to its parent node.

// Return the step-by-step directions of the shortest path from node s to node t.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/11/15/eg1.png" />
// Input: root = [5,1,2,3,null,6,4], startValue = 3, destValue = 6
// Output: "UURL"
// Explanation: The shortest path is: 3 → 1 → 5 → 2 → 6.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/11/15/eg2.png" />
// Input: root = [2,1], startValue = 2, destValue = 1
// Output: "L"
// Explanation: The shortest path is: 2 → 1.

// Constraints:
//     The number of nodes in the tree is n.
//     2 <= n <= 10^5
//     1 <= Node.val <= n
//     All the values in the tree are unique.
//     1 <= startValue, destValue <= n
//     startValue != destValue

import "fmt"
import "strings"

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
// dfs + bfs
func getDirections(root *TreeNode, startValue, destValue int) string {
    queue, parents := []*TreeNode{nil}, map[*TreeNode]*TreeNode{}
    var dfs func(node, pa *TreeNode)
    dfs = func(node, pa *TreeNode) {
        if node == nil { return }
        parents[node] = pa
        if node.Val == startValue {
            queue[0] = node // 只有一个起点
        }
        dfs(node.Left, node)
        dfs(node.Right, node)
    }
    dfs(root, nil)
    res, vis := []byte{}, map[*TreeNode]bool{nil: true, queue[0]: true}
    type pair struct {
        from *TreeNode
        dir  byte
    }
    from := map[*TreeNode]pair{}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        if node.Val == destValue {
            for ; from[node].from != nil; node = from[node].from {
                res = append(res, from[node].dir)
            }
            break
        }
        if !vis[node.Left] {
            vis[node.Left] = true
            from[node.Left] = pair{node, 'L'}
            queue = append(queue, node.Left)
        }
        if !vis[node.Right] {
            vis[node.Right] = true
            from[node.Right] = pair{node, 'R'}
            queue = append(queue, node.Right)
        }
        if !vis[parents[node]] {
            vis[parents[node]] = true
            from[parents[node]] = pair{node, 'U'}
            queue = append(queue, parents[node])
        }
    }
    for i, n := 0, len(res); i < n/2; i++ {
        res[i], res[n-1-i] = res[n-1-i], res[i]
    }
    return string(res)
}

// 最近公共祖先
func getDirections1(root *TreeNode, startValue, destValue int) string {
    path := []byte{}
    var dfs func(*TreeNode, int) bool
    dfs = func(node *TreeNode, target int) bool {
        if node == nil { return false }
        if node.Val == target { return true }
        path = append(path, 'L')
        if dfs(node.Left, target) { return true }
        path[len(path)-1] = 'R'
        if dfs(node.Right, target) { return true }
        path = path[:len(path)-1]
        return false
    }
    dfs(root, startValue)
    pathToStart := path

    path = nil
    dfs(root, destValue)
    pathToDest := path

    for len(pathToStart) > 0 && len(pathToDest) > 0 && pathToStart[0] == pathToDest[0] {
        pathToStart = pathToStart[1:] // 去掉前缀相同的部分
        pathToDest = pathToDest[1:]
    }
    return strings.Repeat("U", len(pathToStart)) + string(pathToDest)
}

func getDirections2(root *TreeNode, startValue int, destValue int) string {
    var findTarget func(root *TreeNode, targetValue int, path *[]byte) bool
    findTarget = func(root *TreeNode, targetValue int, path *[]byte) bool {
        if root.Val == targetValue {
            return true
        }
        if root.Left != nil && findTarget(root.Left, targetValue, path) {
            *path = append(*path, 'L')
            return true
        }
        if root.Right != nil && findTarget(root.Right, targetValue, path) {
            *path = append(*path, 'R')
            return true
        }
        return false
    }
    reverse := func (path []byte) {
        left, right := 0, len(path)-1
        for left < right {
            path[left], path[right] = path[right], path[left]
            left++
            right--
        }
    }
    replace := func(path []byte) {
        for i := 0; i < len(path); i++ {
            path[i] = 'U'
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    sPath, dPath := make([]byte, 0), make([]byte, 0)
    findTarget(root, startValue, &sPath)
    findTarget(root, destValue, &dPath)
    size, i := min(len(sPath), len(dPath)), 0
    for i < size {
        if sPath[len(sPath)-1-i] == dPath[len(dPath)-1-i] {
            i++
        } else {
            break
        }
    }
    sPath = sPath[:len(sPath)-i]
    replace(sPath)
    dPath = dPath[:len(dPath)-i]
    reverse(dPath)
    sPath = append(sPath, dPath...)
    return string(sPath)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/11/15/eg1.png" />
    // Input: root = [5,1,2,3,null,6,4], startValue = 3, destValue = 6
    // Output: "UURL"
    // Explanation: The shortest path is: 3 → 1 → 5 → 2 → 6.
    tree1 := &TreeNode {
        5,
        &TreeNode { 1, &TreeNode { 3, nil, nil }, nil },
        &TreeNode { 2, &TreeNode { 6, nil, nil }, &TreeNode { 4, nil, nil } },
    }
    fmt.Println(getDirections(tree1,3, 6)) // "UURL"
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/11/15/eg2.png" />
    // Input: root = [2,1], startValue = 2, destValue = 1
    // Output: "L"
    // Explanation: The shortest path is: 2 → 1.
    tree2 := &TreeNode {
        1,
        &TreeNode { 2, nil, nil },
        nil,
    }
    fmt.Println(getDirections(tree2,2, 1)) // "L"

    fmt.Println(getDirections1(tree1,3, 6)) // "UURL"
    fmt.Println(getDirections1(tree2,2, 1)) // "L"

    fmt.Println(getDirections2(tree1,3, 6)) // "UURL"
    fmt.Println(getDirections2(tree2,2, 1)) // "L"
}