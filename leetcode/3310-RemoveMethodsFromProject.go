package main

// 3310. Remove Methods From Project
// You are maintaining a project that has n methods numbered from 0 to n - 1.

// You are given two integers n and k, and a 2D integer array invocations, 
// where invocations[i] = [ai, bi] indicates that method ai invokes method bi.

// There is a known bug in method k. Method k, along with any method invoked by it, either directly or indirectly, are considered suspicious and we aim to remove them.

// A group of methods can only be removed if no method outside the group invokes any methods within it.

// Return an array containing all the remaining methods after removing all the suspicious methods. 
// You may return the answer in any order. 
// If it is not possible to remove all the suspicious methods, none should be removed.

// Example 1:
// Input: n = 4, k = 1, invocations = [[1,2],[0,1],[3,2]]
// Output: [0,1,2,3]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/07/18/graph-2.png" />
// Method 2 and method 1 are suspicious, but they are directly invoked by methods 3 and 0, which are not suspicious. 
// We return all elements without removing anything.

// Example 2:
// Input: n = 5, k = 0, invocations = [[1,2],[0,2],[0,1],[3,4]]
// Output: [3,4]
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/07/18/graph-3.png" />
// Methods 0, 1, and 2 are suspicious and they are not directly invoked by any other method. 
// We can remove them.

// Example 3:
// Input: n = 3, k = 2, invocations = [[1,2],[0,1],[2,0]]
// Output: []
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/07/20/graph.png" />
// All methods are suspicious. We can remove them.

// Constraints:
//     1 <= n <= 10^5
//     0 <= k <= n - 1
//     0 <= invocations.length <= 2 * 10^5
//     invocations[i] == [ai, bi]
//     0 <= ai, bi <= n - 1
//     ai != bi
//     invocations[i] != invocations[j]

import "fmt"

func remainingMethods(n int, k int, invocations [][]int) []int {
    adj := map[int][]int{}
    for _, v := range invocations {
        adj[v[0]] = append(adj[v[0]], v[1])
    }
    res, stack, within := []int{}, []int{k}, make(map[int]bool)
    for len(stack) > 0 {
        top := stack[len(stack) - 1]
        within[top] = true
        stack = stack[:len(stack) - 1]
        for _,v := range(adj[top]) {
            if !within[v] {
                stack = append(stack, v)
            }
        }
    }
    for _, v  := range invocations {
        if !within[v[0]] && within[v[1]] {
            curr := []int{}
            for i := 0 ; i < n ; i ++ {
                curr = append(curr, i)
            }
            return curr
        }
    }
    for i := 0 ; i < n ; i ++ {
        if !within[i] {
            res = append(res, i)
        }
    }
    return res 
}

func remainingMethods1(n int, k int, invocations [][]int) []int {
    res, adj := make([]int, 0, n), make([][]int, n)
    for _, v := range invocations {
        adj[v[0]] = append(adj[v[0]], v[1])
    }
    queue, suspicious, count := []int{ k }, make([]bool, n), 1
    suspicious[k] = true
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        for _, next := range adj[cur] {
            if suspicious[next] { continue }
            suspicious[next], count = true, count + 1
            queue = append(queue, next)
        }
    }
    if count == n { return res }
    visited, remove := make([]bool, n), true
    for i := 0; i < n && remove; i++ {
        if suspicious[i] { continue }
        queue = append(queue, i)
        for len(queue) > 0 && remove {
            cur := queue[0]
            queue = queue[1:]
            for _, next := range adj[cur] {
                if visited[next] {
                    continue
                } else if suspicious[next] {
                    remove = false
                    break
                }
                visited[next] = true
                queue = append(queue, next)
            }
        }
    }
    for i := 0; i < n; i++ {
        if !remove || !suspicious[i] {
            res = append(res, i)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 4, k = 1, invocations = [[1,2],[0,1],[3,2]]
    // Output: [0,1,2,3]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/07/18/graph-2.png" />
    // Method 2 and method 1 are suspicious, but they are directly invoked by methods 3 and 0, which are not suspicious. 
    // We return all elements without removing anything.
    fmt.Println(remainingMethods(4, 1, [][]int{{1,2},{0,1},{3,2}})) // [0,1,2,3]
    // Example 2:
    // Input: n = 5, k = 0, invocations = [[1,2],[0,2],[0,1],[3,4]]
    // Output: [3,4]
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/07/18/graph-3.png" />
    // Methods 0, 1, and 2 are suspicious and they are not directly invoked by any other method. 
    // We can remove them.
    fmt.Println(remainingMethods(5, 0, [][]int{{1,2},{0,2},{0,1},{3,4}})) // [3,4]
    // Example 3:
    // Input: n = 3, k = 2, invocations = [[1,2],[0,1],[2,0]]
    // Output: []
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/07/20/graph.png" />
    // All methods are suspicious. We can remove them.
    fmt.Println(remainingMethods(3, 2, [][]int{{1,2},{0,1},{2,0}})) // []

    fmt.Println(remainingMethods1(4, 1, [][]int{{1,2},{0,1},{3,2}})) // [0,1,2,3]
    fmt.Println(remainingMethods1(5, 0, [][]int{{1,2},{0,2},{0,1},{3,4}})) // [3,4]
    fmt.Println(remainingMethods1(3, 2, [][]int{{1,2},{0,1},{2,0}})) // []
}