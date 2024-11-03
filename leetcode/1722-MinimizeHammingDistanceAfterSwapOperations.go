package main

// 1722. Minimize Hamming Distance After Swap Operations
// You are given two integer arrays, source and target, both of length n. 
// You are also given an array allowedSwaps where each allowedSwaps[i] = [ai, bi] indicates 
// that you are allowed to swap the elements at index ai and index bi (0-indexed) of array source. 
// Note that you can swap elements at a specific pair of indices multiple times and in any order.

// The Hamming distance of two arrays of the same length, source and target, 
// is the number of positions where the elements are different. 
// Formally, it is the number of indices i for 0 <= i <= n-1 where source[i] != target[i] (0-indexed).

// Return the minimum Hamming distance of source 
// and target after performing any amount of swap operations on array source.

// Example 1:
// Input: source = [1,2,3,4], target = [2,1,4,5], allowedSwaps = [[0,1],[2,3]]
// Output: 1
// Explanation: source can be transformed the following way:
// - Swap indices 0 and 1: source = [2,1,3,4]
// - Swap indices 2 and 3: source = [2,1,4,3]
// The Hamming distance of source and target is 1 as they differ in 1 position: index 3.

// Example 2:
// Input: source = [1,2,3,4], target = [1,3,2,4], allowedSwaps = []
// Output: 2
// Explanation: There are no allowed swaps.
// The Hamming distance of source and target is 2 as they differ in 2 positions: index 1 and index 2.

// Example 3:
// Input: source = [5,1,2,4,3], target = [1,5,4,2,3], allowedSwaps = [[0,4],[4,2],[1,3],[1,4]]
// Output: 0

// Constraints:
//     n == source.length == target.length
//     1 <= n <= 10^5
//     1 <= source[i], target[i] <= 10^5
//     0 <= allowedSwaps.length <= 10^5
//     allowedSwaps[i].length == 2
//     0 <= ai, bi <= n - 1
//     ai != bi

import "fmt"

// bfs
func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
    graph := make(map[int][]int)
    for _, v := range allowedSwaps {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    res, visited := 0, make(map[int]bool)
    for i := 0; i < len(source); i++ {
        if visited[i] { continue }
        count, queue := make(map[int]int), []int{ i }
        seen := queue
        for len(queue) > 0 {
            visited[i] = true
            j := queue[0]
            queue = queue[1:] // pop
            for _, next := range graph[j] {
                if !visited[next] {
                    queue = append(queue, next)
                    seen = append(seen, next)
                    visited[next] = true
                }
            }
        }
        for _, v := range seen {
            count[source[v]]++
            count[target[v]]--
        }
        for _, v := range count {
            if v > 0 {
                res += v
            }
        }
    }
    return res
}

// dfs
func minimumHammingDistance1(source []int, target []int, allowedSwaps [][]int) int {
    n := len(source)
    graph := make(map[int][]int)
    for _, v := range allowedSwaps {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var dfs func(i int, scc map[int]bool)
    dfs = func(i int, scc map[int]bool) {
        scc[i] = true
        for _, v := range graph[i] {
            if !scc[v] {
                dfs(v, scc)
            }
        }
    }
    same, visited := 0, make([]bool, n)
    for i := 0; i < n; i++ {
        if !visited[i] {
            scc := make(map[int]bool)
            dfs(i, scc)
            v_s, v_t := make(map[int]int), make(map[int]int)
            for j, _ := range scc {
                visited[j] = true
                v_s[source[j]]++
                v_t[target[j]]++
            }
            for k, v := range v_s {
                same += min(v, v_t[k])
            }
        }
    }
    for i, v := range source {
        if !visited[i] && v == target[i] {
            same++
        }
    }
    return n - same
}

func main() {
    // Example 1:
    // Input: source = [1,2,3,4], target = [2,1,4,5], allowedSwaps = [[0,1],[2,3]]
    // Output: 1
    // Explanation: source can be transformed the following way:
    // - Swap indices 0 and 1: source = [2,1,3,4]
    // - Swap indices 2 and 3: source = [2,1,4,3]
    // The Hamming distance of source and target is 1 as they differ in 1 position: index 3.
    fmt.Println(minimumHammingDistance([]int{1,2,3,4}, []int{2,1,4,5}, [][]int{{0,1},{2,3}})) // 1
    // Example 2:
    // Input: source = [1,2,3,4], target = [1,3,2,4], allowedSwaps = []
    // Output: 2
    // Explanation: There are no allowed swaps.
    // The Hamming distance of source and target is 2 as they differ in 2 positions: index 1 and index 2.
    fmt.Println(minimumHammingDistance([]int{1,2,3,4}, []int{1,3,2,4}, [][]int{})) // 2
    // Example 3:
    // Input: source = [5,1,2,4,3], target = [1,5,4,2,3], allowedSwaps = [[0,4],[4,2],[1,3],[1,4]]
    // Output: 0
    fmt.Println(minimumHammingDistance([]int{5,1,2,4,3}, []int{1,5,4,2,3}, [][]int{{0,4},{4,2},{1,3},{1,4}})) // 0

    fmt.Println(minimumHammingDistance1([]int{1,2,3,4}, []int{2,1,4,5}, [][]int{{0,1},{2,3}})) // 1
    // Example 2:
    // Input: source = [1,2,3,4], target = [1,3,2,4], allowedSwaps = []
    // Output: 2
    // Explanation: There are no allowed swaps.
    // The Hamming distance of source and target is 2 as they differ in 2 positions: index 1 and index 2.
    fmt.Println(minimumHammingDistance1([]int{1,2,3,4}, []int{1,3,2,4}, [][]int{})) // 2
    // Example 3:
    // Input: source = [5,1,2,4,3], target = [1,5,4,2,3], allowedSwaps = [[0,4],[4,2],[1,3],[1,4]]
    // Output: 0
    fmt.Println(minimumHammingDistance1([]int{5,1,2,4,3}, []int{1,5,4,2,3}, [][]int{{0,4},{4,2},{1,3},{1,4}})) // 0
}