package main

// 1743. Restore the Array From Adjacent Pairs
// There is an integer array nums that consists of n unique elements, but you have forgotten it. 
// However, you do remember every pair of adjacent elements in nums.

// You are given a 2D integer array adjacentPairs of size n - 1 
// where each adjacentPairs[i] = [ui, vi] indicates that the elements ui and vi are adjacent in nums.

// It is guaranteed that every adjacent pair of elements nums[i] 
// and nums[i+1] will exist in adjacentPairs, either as [nums[i], nums[i+1]] or [nums[i+1], nums[i]]. 
// The pairs can appear in any order.

// Return the original array nums. 
// If there are multiple solutions, return any of them.

// Example 1:
// Input: adjacentPairs = [[2,1],[3,4],[3,2]]
// Output: [1,2,3,4]
// Explanation: This array has all its adjacent pairs in adjacentPairs.
// Notice that adjacentPairs[i] may not be in left-to-right order.

// Example 2:
// Input: adjacentPairs = [[4,-2],[1,4],[-3,1]]
// Output: [-2,4,1,-3]
// Explanation: There can be negative numbers.
// Another solution is [-3,1,4,-2], which would also be accepted.

// Example 3:
// Input: adjacentPairs = [[100000,-100000]]
// Output: [100000,-100000]

// Constraints:
//     nums.length == n
//     adjacentPairs.length == n - 1
//     adjacentPairs[i].length == 2
//     2 <= n <= 10^5
//     -10^5 <= nums[i], ui, vi <= 10^5
//     There exists some nums that has adjacentPairs as its pairs.

import "fmt"

// bfs
func restoreArray(adjacentPairs [][]int) []int {
    // 1 Create adjacency matrix
    adj := make(map[int][]int)
    for _, v := range adjacentPairs {
        adj[v[0]] = append(adj[v[0]], v[1])
        adj[v[1]] = append(adj[v[1]], v[0])
    }
    // 2 Find any node with 1 edge
    start := 0
    for k, v := range adj {
        if len(v) == 1 {
            start = k
            break
        }
    }
    // 3 Iterate
    res, queue, visited := []int{}, []int{ start }, map[int]bool{start: true}
    for len(queue) > 0 {
        n := queue[0]
        queue = queue[1:] // pop
        res = append(res, n)
        for _, v := range adj[n] {
            if visited[v] { continue }
            visited[v] = true
            queue = append(queue, v)
        }
    }
    return res
}

// dfs
func restoreArray1(adjacentPairs [][]int) []int {
    adj, count := make(map[int][]int), make(map[int]int)
    for _, v := range adjacentPairs {
        adj[v[0]] = append(adj[v[0]], v[1])
        adj[v[1]] = append(adj[v[1]], v[0])
        count[v[0]]++
        count[v[1]]++
    }
    start := -1
    for k, v := range count {
        if v == 1 {
            start = k
            break
        }
    }
    res, visited := make([]int, 0), make(map[int]bool)
    var dfs func(n int)
    dfs = func(n int) {
        if visited[n] { return }
        res = append(res, n)
        visited[n] = true
        for _, v := range adj[n] {
            dfs(v)        
        }
        
    }
    dfs(start)
    return res
}

func main() {
    // Example 1:
    // Input: adjacentPairs = [[2,1],[3,4],[3,2]]
    // Output: [1,2,3,4]
    // Explanation: This array has all its adjacent pairs in adjacentPairs.
    // Notice that adjacentPairs[i] may not be in left-to-right order.
    fmt.Println(restoreArray([][]int{{2,1},{3,4},{3,2}})) // [1,2,3,4]
    // Example 2:
    // Input: adjacentPairs = [[4,-2],[1,4],[-3,1]]
    // Output: [-2,4,1,-3]
    // Explanation: There can be negative numbers.
    // Another solution is [-3,1,4,-2], which would also be accepted.
    fmt.Println(restoreArray([][]int{{4,-2},{1,4},{-3,1}})) // [-2,4,1,-3]
    // Example 3:
    // Input: adjacentPairs = [[100000,-100000]]
    // Output: [100000,-100000]
    fmt.Println(restoreArray([][]int{{100000,-100000}})) // [100000,-100000]

    fmt.Println(restoreArray1([][]int{{2,1},{3,4},{3,2}})) // [1,2,3,4]
    fmt.Println(restoreArray1([][]int{{4,-2},{1,4},{-3,1}})) // [-2,4,1,-3]
    fmt.Println(restoreArray1([][]int{{100000,-100000}})) // [100000,-100000]
}