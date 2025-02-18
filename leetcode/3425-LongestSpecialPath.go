package main

// 3425. Longest Special Path
// You are given an undirected tree rooted at node 0 with n nodes numbered from 0 to n - 1, 
// represented by a 2D array edges of length n - 1, 
// where edges[i] = [ui, vi, lengthi] indicates an edge between nodes ui and vi with length lengthi. 
// You are also given an integer array nums, where nums[i] represents the value at node i.

// A special path is defined as a downward path from an ancestor node to a descendant node such that all the values of the nodes in that path are unique.

// Note that a path may start and end at the same node.

// Return an array result of size 2, where result[0] is the length of the longest special path, 
// and result[1] is the minimum number of nodes in all possible longest special paths.

// Example 1:
// Input: edges = [[0,1,2],[1,2,3],[1,3,5],[1,4,4],[2,5,6]], nums = [2,1,2,1,3,1]
// Output: [6,2]
// Explanation:
// In the image below, nodes are colored by their corresponding values in nums
// <img src="https://assets.leetcode.com/uploads/2024/11/02/tree3.jpeg" />
// The longest special paths are 2 -> 5 and 0 -> 1 -> 4, both having a length of 6. 
// The minimum number of nodes across all longest special paths is 2.

// Example 2:
// Input: edges = [[1,0,8]], nums = [2,2]
// Output: [0,1]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/11/02/tree4.jpeg" />
// The longest special paths are 0 and 1, both having a length of 0. 
// The minimum number of nodes across all longest special paths is 1.

// Constraints:
//     2 <= n <= 5 * 10^4
//     edges.length == n - 1
//     edges[i].length == 3
//     0 <= ui, vi < n
//     1 <= lengthi <= 10^3
//     nums.length == n
//     0 <= nums[i] <= 5 * 10^4
//     The input is generated such that edges represents a valid tree.

import "fmt"
import "slices"

func longestSpecialPath(edges [][]int, nums []int) []int {
    n := len(nums)
    res, graph := []int{0, n + 1}, make(map[int][][2]int)
    for _, v := range edges {
        graph[v[0]] = append(graph[v[0]], [2]int{v[1], v[2]})
        graph[v[1]] = append(graph[v[1]], [2]int{v[0], v[2] })
    }
    costs, last := []int{}, make(map[int]int)
    for _, v := range nums {
        last[v] = -1
    }
    var dfs func(node, cur, prev, left int)
    dfs = func(node, cur, prev, left int) {
        index := last[nums[node]]
        last[nums[node]] = len(costs)
        costs = append(costs, cur)
        if cur - costs[left] > res[0] {
            res[0] = cur - costs[left]
            res[1] = len(costs) - left
        } else if cur - costs[left] == res[0] {
            res[1] = min(res[1], len(costs) - left)
        }
        for _, v := range graph[node] {
            nn, nc := v[0], v[1]
            if nn == prev {
                continue
            }
            nl := left
            if last[nums[nn]] != -1 && last[nums[nn]] >= left {
                nl = last[nums[nn]] + 1
            }
            dfs(nn, cur + nc, node, nl)
        }
        last[nums[node]] = index
        costs = costs[:len(costs) - 1]
    }
    dfs(0, 0, -1, 0)
    return res
}

func longestSpecialPath1(edges [][]int, nums []int) []int {
    n := len(nums)
    // Representing the tree as an adjacency list
    type edge struct{ to, weight int }
    tree := make([][]edge, n)
    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        tree[u] = append(tree[u], edge{v, w})
        tree[v] = append(tree[v], edge{u, w})
    }
    // Tracking the total distance at each depth
    distSum := make([]int, n)
    maxVal := slices.Max(nums)
    depthToNode := make([]int, n)
    // To track the last depth where a value appeared
    lastSeenDepth := make([]int, maxVal+1)
    for i := range lastSeenDepth {
        lastSeenDepth[i] = -1
    }
    // Variables to store the final result
    longestLength := 0
    minNodes := n + 1
    // Pointer to track the start of the valid unique path
    ptr := 0
    // DFS function to traverse the tree
    var dfs func(node, parent, depth, dist int)
    dfs = func(node, parent, depth, dist int) {
        // Store node and distance at the current depth
        depthToNode[depth] = node
        distSum[depth] = dist
        // Get the value at this node
        val := nums[node]
        // Store previous depth where this value was seen
        prevDepth := lastSeenDepth[val]
        prevPtr := ptr
        // If value was seen before in this path, move pointer to maintain uniqueness
        if prevDepth >= ptr {
            ptr = prevDepth + 1
        }
        lastSeenDepth[val] = depth // Update last seen depth
        // Compute the current path length
        curLength := distSum[depth] - distSum[ptr]
        if ptr == depth { // If it's a single node, path length is 0
            curLength = 0
        }
        curNodeCount := depth - ptr + 1
        // Update the results if we find a longer special path
        if curLength > longestLength {
            longestLength = curLength
            minNodes = curNodeCount
        } else if curLength == longestLength && curNodeCount < minNodes {
            minNodes = curNodeCount
        }
        // Explore child nodes
        for _, next := range tree[node] {
            if next.to == parent { continue }
            dfs(next.to, node, depth+1, dist+next.weight)
        }
        // Restore previous state for backtracking
        lastSeenDepth[val] = prevDepth
        ptr = prevPtr
    }
    // Start DFS from the root node (0)
    dfs(0, -1, 0, 0)
    return []int{ longestLength, minNodes }
}


func main() {
    // Example 1:
    // Input: edges = [[0,1,2],[1,2,3],[1,3,5],[1,4,4],[2,5,6]], nums = [2,1,2,1,3,1]
    // Output: [6,2]
    // Explanation:
    // In the image below, nodes are colored by their corresponding values in nums
    // <img src="https://assets.leetcode.com/uploads/2024/11/02/tree3.jpeg" />
    // The longest special paths are 2 -> 5 and 0 -> 1 -> 4, both having a length of 6. 
    // The minimum number of nodes across all longest special paths is 2.
    fmt.Println(longestSpecialPath([][]int{{0,1,2},{1,2,3},{1,3,5},{1,4,4},{2,5,6}}, []int{2,1,2,1,3,1})) // [6,2]
    // Example 2:
    // Input: edges = [[1,0,8]], nums = [2,2]
    // Output: [0,1]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/11/02/tree4.jpeg" />
    // The longest special paths are 0 and 1, both having a length of 0. 
    // The minimum number of nodes across all longest special paths is 1.
    fmt.Println(longestSpecialPath([][]int{{1,0,8}}, []int{2,2})) // [0,1]

    fmt.Println(longestSpecialPath1([][]int{{0,1,2},{1,2,3},{1,3,5},{1,4,4},{2,5,6}}, []int{2,1,2,1,3,1})) // [6,2]
    fmt.Println(longestSpecialPath1([][]int{{1,0,8}}, []int{2,2})) // [0,1]
}