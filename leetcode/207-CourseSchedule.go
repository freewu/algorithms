package main

// 207. Course Schedule
// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. 
// You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.
//     For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

// Return true if you can finish all courses. Otherwise, return false.

// Example 1:
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: true
// Explanation: There are a total of 2 courses to take. 
// To take course 1 you should have finished course 0. So it is possible.

// Example 2:
// Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
// Output: false
// Explanation: There are a total of 2 courses to take. 
// To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.

// Constraints:
//     1 <= numCourses <= 2000
//     0 <= prerequisites.length <= 5000
//     prerequisites[i].length == 2
//     0 <= ai, bi < numCourses
//     All the pairs prerequisites[i] are unique.

import "fmt"

// func canFinish(numCourses int, prerequisites [][]int) bool {
//     if len()
//     // 保存每个课程的前置要求
//     m := make(map[int]int,numCourses)
//     return true
// }

// bfs
func canFinish(numCourses int, prerequisites [][]int) bool {
    graph, degree := make([][]int, numCourses), make([]int, numCourses)
    for _, prerequisite := range prerequisites {
        graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
        degree[prerequisite[0]] += 1
    }
    res := make([]int, 0)
    for course, v := range degree {
        if v == 0 { // 没有前置要求的课程直接加入
            res = append(res, course)
        }
    }
    for i := 0; i < len(res); i ++{
        for _, j := range graph[res[i]] {
            degree[j] -= 1
            if degree[j] == 0 {
                res = append(res, j)
            }
        }
    }
    if len(res) == numCourses {
        return true
    }
    return false
}

func canFinish1(numCourses int, prerequisites [][]int) bool {
    queue, rudu, graph := []int{}, make([]int, numCourses), make(map[int][]int)
    for _,v := range prerequisites {
        rudu[v[0]]++
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    for i, v := range rudu {
        if v == 0 {
            queue = append(queue, i)
        }
    }
    res := 0
    for len(queue) > 0 {
        c := queue[0]
        queue = queue[1:]
        res++
        for _,v := range graph[c] {
            rudu[v]--
            if rudu[v] == 0 {
                queue = append(queue, v)
            }
        }
    }
    return res == numCourses
}

func main() {
    // Explanation: There are a total of 2 courses to take. 
    // To take course 1 you should have finished course 0. So it is possible.
    fmt.Println(canFinish(2,[][]int{{1,0}})) // true
    // Explanation: There are a total of 2 courses to take. 
    // To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
    fmt.Println(canFinish(2,[][]int{{1,0},{0,1}})) // true

    fmt.Println(canFinish1(2,[][]int{{1,0}})) // true
    fmt.Println(canFinish1(2,[][]int{{1,0},{0,1}})) // true
}