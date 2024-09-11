package main

// 1462. Course Schedule IV
// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. 
// You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course ai first if you want to take course bi.
//     For example, the pair [0, 1] indicates that you have to take course 0 before you can take course 1.

// Prerequisites can also be indirect. 
// If course a is a prerequisite of course b, and course b is a prerequisite of course c, then course a is a prerequisite of course c.

// You are also given an array queries where queries[j] = [uj, vj]. 
// For the jth query, you should answer whether course uj is a prerequisite of course vj or not.

// Return a boolean array answer, where answer[j] is the answer to the jth query.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/05/01/courses4-1-graph.jpg" />
// Input: numCourses = 2, prerequisites = [[1,0]], queries = [[0,1],[1,0]]
// Output: [false,true]
// Explanation: The pair [1, 0] indicates that you have to take course 1 before you can take course 0.
// Course 0 is not a prerequisite of course 1, but the opposite is true.

// Example 2:
// Input: numCourses = 2, prerequisites = [], queries = [[1,0],[0,1]]
// Output: [false,false]
// Explanation: There are no prerequisites, and each course is independent.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/05/01/courses4-3-graph.jpg" />
// Input: numCourses = 3, prerequisites = [[1,2],[1,0],[2,0]], queries = [[1,0],[1,2]]
// Output: [true,true]

// Constraints:
//     2 <= numCourses <= 100
//     0 <= prerequisites.length <= (numCourses * (numCourses - 1) / 2)
//     prerequisites[i].length == 2
//     0 <= ai, bi <= numCourses - 1
//     ai != bi
//     All the pairs [ai, bi] are unique.
//     The prerequisites graph has no cycles.
//     1 <= queries.length <= 10^4
//     0 <= ui, vi <= numCourses - 1
//     ui != vi

import "fmt"

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    graph := make([][]int, numCourses)
    for i := 0; i < numCourses; i++ {
        graph[i] = []int{}
    }

    for i := 0; i < len(prerequisites); i++ {
        u, v := prerequisites[i][1], prerequisites[i][0]
        graph[u] = append(graph[u], v)
    }
    mp := map[int]map[int]bool{}
    var dfs func(u int) map[int]bool
    dfs = func(u int) map[int]bool {
        if v, ok := mp[u]; ok { return v }
        mp[u] = map[int]bool{}
        for _, neighbor := range graph[u] {
            mp[u][neighbor] = true
            for m, _ := range dfs(neighbor) {
                mp[u][m] = true
            }
        }
        return mp[u]
    }
    for i := 0; i < numCourses; i++ { dfs(i) }
    // fmt.Println(mp)
    res := make([]bool, len(queries))
    for i, q := range queries {
        u, v := q[0], q[1]
        if edges, ok := mp[v]; ok && edges[u] {
            res[i] = true
        } else {
            res[i] = false
        }
    }
    return res
}

func checkIfPrerequisite1(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    graph := make([][]int, numCourses)
    for _, pre := range prerequisites {
        f, t := pre[0], pre[1]
        graph[f] = append(graph[f], t)
    }
    connected := make([][]bool, numCourses) // numCourses * numCourses
    for i := range connected {
        connected[i] = make([]bool, numCourses)
    }
    visited := make([]bool, numCourses)
    var dfs func(node int, start int)
    dfs = func(node int, start int) {
        visited[node] = true
        connected[start][node] = true
        for _, next := range graph[node] {
            if !visited[next] {
                dfs(next, start)
            }
        }
    }
    for course := 0; course < numCourses; course++ {
        for i := range visited {
            visited[i] = false
        }
        dfs(course, course)
    }
    res := make([]bool, len(queries))
    for i := range queries {
        x, y := queries[i][0], queries[i][1]
        res[i] = connected[x][y]
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/05/01/courses4-1-graph.jpg" />
    // Input: numCourses = 2, prerequisites = [[1,0]], queries = [[0,1],[1,0]]
    // Output: [false,true]
    // Explanation: The pair [1, 0] indicates that you have to take course 1 before you can take course 0.
    // Course 0 is not a prerequisite of course 1, but the opposite is true.
    fmt.Println(checkIfPrerequisite(2,[][]int{{1,0}},[][]int{{0,1},{1,0}})) // [false,true]
    // Example 2:
    // Input: numCourses = 2, prerequisites = [], queries = [[1,0],[0,1]]
    // Output: [false,false]
    // Explanation: There are no prerequisites, and each course is independent.
    fmt.Println(checkIfPrerequisite(2,[][]int{},[][]int{{1,0},{0,1}})) // [false,false]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/05/01/courses4-3-graph.jpg" />
    // Input: numCourses = 3, prerequisites = [[1,2],[1,0],[2,0]], queries = [[1,0],[1,2]]
    // Output: [true,true]]
    fmt.Println(checkIfPrerequisite(3,[][]int{{1,2},{1,0},{2,0}},[][]int{{1,0},{1,2}})) // [true,true]]

    fmt.Println(checkIfPrerequisite1(2,[][]int{{1,0}},[][]int{{0,1},{1,0}})) // [false,true]
    fmt.Println(checkIfPrerequisite1(2,[][]int{},[][]int{{1,0},{0,1}})) // [false,false]
    fmt.Println(checkIfPrerequisite1(3,[][]int{{1,2},{1,0},{2,0}},[][]int{{1,0},{1,2}})) // [true,true]]
}