package main

// 1136. Parallel Courses
// You are given an integer n, which indicates that there are n courses labeled from 1 to n. 
// You are also given an array relations where relations[i] = [prevCoursei, nextCoursei], 
// representing a prerequisite relationship between course prevCoursei and course nextCoursei: 
//     course prevCoursei has to be taken before course nextCoursei.

// In one semester, you can take any number of courses as long as you have taken all the prerequisites in the previous semester for the courses you are taking.
// Return the minimum number of semesters needed to take all courses. If there is no way to take all the courses, return -1.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/24/course1graph.jpg" />
// Input: n = 3, relations = [[1,3],[2,3]]
// Output: 2
// Explanation: The figure above represents the given graph.
// In the first semester, you can take courses 1 and 2.
// In the second semester, you can take course 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/02/24/course2graph.jpg" />
// Input: n = 3, relations = [[1,2],[2,3],[3,1]]
// Output: -1
// Explanation: No course can be studied because they are prerequisites of each other.

// Constraints:
//     1 <= n <= 5000
//     1 <= relations.length <= 5000
//     relations[i].length == 2
//     1 <= prevCoursei, nextCoursei <= n
//     prevCoursei != nextCoursei
//     All the pairs [prevCoursei, nextCoursei] are unique.

import "fmt"

// topsort
func minimumSemesters(n int, relations [][]int) int {
    graph := make([][]int, n+1) //邻接表
    for i:=0; i<n; i++ {
        graph[i] = []int{}
    }
    indegree := make([]int, n+1) //入度表
    for i:=0; i<len(relations); i++ { //构建邻接表和入度表
        first, second := relations[i][0], relations[i][1]
        graph[first] = append(graph[first], second)
        indegree[second]++
    }
    queue := []int{}
    for i:=1; i<n; i++ { // 入度为0的点写入队列
        if 0 == indegree[i] {
            queue = append(queue, i)
        }
    }
    res, visit := 0, 0 // 课程访问数量
    for 0 != len(queue) {
        res++ // 入度同时为 0 的节点在一个学期内一起修
        size := len(queue)
        for i:=0; i<size; i++ { //将同一批入度为0的节点修完
            cur := queue[i]
            for _, nxt := range graph[cur] {
                indegree[nxt]--
                if 0 == indegree[nxt] { //入度为0，写入队列，为下一批修的课程
                    queue = append(queue, nxt)
                }
            }
            visit++
        }
        queue = queue[size:]
    }
    if visit != n { // 如果课程访问数量 != 课程数，肯定存在环
        return -1
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/02/24/course1graph.jpg" />
    // Input: n = 3, relations = [[1,3],[2,3]]
    // Output: 2
    // Explanation: The figure above represents the given graph.
    // In the first semester, you can take courses 1 and 2.
    // In the second semester, you can take course 3.
    fmt.Println(minimumSemesters(3,[][]int{{1,3},{2,3}})) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/02/24/course2graph.jpg" />
    // Input: n = 3, relations = [[1,2],[2,3],[3,1]]
    // Output: -1
    // Explanation: No course can be studied because they are prerequisites of each other.
    fmt.Println(minimumSemesters(3,[][]int{{1,2},{2,3},{3,1}})) // -1
}