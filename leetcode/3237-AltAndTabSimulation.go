package main

// 3237. Alt and Tab Simulation
// There are n windows open numbered from 1 to n, we want to simulate using alt + tab to navigate between the windows.
// You are given an array windows which contains the initial order of the windows (the first element is at the top and the last one is at the bottom).
// You are also given an array queries where for each query, the window queries[i] is brought to the top.
// Return the final state of the array windows.

// Example 1:
// Input: windows = [1,2,3], queries = [3,3,2]
// Output: [2,3,1]
// Explanation:
// Here is the window array after each query:
// Initial order: [1,2,3]
// After the first query: [3,1,2]
// After the second query: [3,1,2]
// After the last query: [2,3,1]

// Example 2:
// Input: windows = [1,4,2,3], queries = [4,1,3]
// Output: [3,1,4,2]
// Explanation:
// Here is the window array after each query:
// Initial order: [1,4,2,3]
// After the first query: [4,1,2,3]
// After the second query: [1,4,2,3]
// After the last query: [3,1,4,2]

// Constraints:
//     1 <= n == windows.length <= 10^5
//     windows is a permutation of [1, n].
//     1 <= queries.length <= 10^5
//     1 <= queries[i] <= n

import "fmt"

func simulationResult(windows []int, queries []int) []int {
    n, order, activated := len(windows), []int{}, map[int]bool{}
    for i := len(queries) - 1; i >= 0; i-- {
        window := queries[i]
        if !activated[window] {
            activated[window] = true
            order = append(order, window)
        }
    }
    for i := 0; i < n; i++ {
        window := windows[i]
        if !activated[window] {
            order = append(order, window)
        }
    }
    return order
}

func simulationResult1(windows []int, queries []int) []int {
    n, m := len(windows), len(queries)
    res, visited := []int{}, make([]bool, n + 1)
    for i := m - 1; i >= 0; i-- {
        one := queries[i]
        if !visited[one] {
            visited[one] = true
            res = append(res, one)
        }
    }
    for _, v := range windows {
        if !visited[v] {
            res = append(res, v)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: windows = [1,2,3], queries = [3,3,2]
    // Output: [2,3,1]
    // Explanation:
    // Here is the window array after each query:
    // Initial order: [1,2,3]
    // After the first query: [3,1,2]
    // After the second query: [3,1,2]
    // After the last query: [2,3,1]
    fmt.Println(simulationResult([]int{1,2,3},[]int{3,3,2})) // [2,3,1]
    // Example 2:
    // Input: windows = [1,4,2,3], queries = [4,1,3]
    // Output: [3,1,4,2]
    // Explanation:
    // Here is the window array after each query:
    // Initial order: [1,4,2,3]
    // After the first query: [4,1,2,3]
    // After the second query: [1,4,2,3]
    // After the last query: [3,1,4,2]
    fmt.Println(simulationResult([]int{1,4,2,3},[]int{4,1,3})) // [3,1,4,2]

    fmt.Println(simulationResult1([]int{1,2,3},[]int{3,3,2})) // [2,3,1]
    fmt.Println(simulationResult1([]int{1,4,2,3},[]int{4,1,3})) // [3,1,4,2]
}