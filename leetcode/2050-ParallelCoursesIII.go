package main

// 2050. Parallel Courses III
// You are given an integer n, which indicates that there are n courses labeled from 1 to n. 
// You are also given a 2D integer array relations where relations[j] = [prevCoursej, nextCoursej] denotes 
// that course prevCoursej has to be completed before course nextCoursej (prerequisite relationship). 
// Furthermore, you are given a 0-indexed integer array time where time[i] denotes how many months it takes to complete the (i+1)th course.

// You must find the minimum number of months needed to complete all the courses following these rules:
//     You may start taking a course at any time if the prerequisites are met.
//     Any number of courses can be taken at the same time.

// Return the minimum number of months needed to complete all the courses.

// Note: The test cases are generated such that it is possible to complete every course (i.e., the graph is a directed acyclic graph).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/10/07/ex1.png" />
// Input: n = 3, relations = [[1,3],[2,3]], time = [3,2,5]
// Output: 8
// Explanation: The figure above represents the given graph and the time required to complete each course. 
// We start course 1 and course 2 simultaneously at month 0.
// Course 1 takes 3 months and course 2 takes 2 months to complete respectively.
// Thus, the earliest time we can start course 3 is at month 3, and the total time required is 3 + 5 = 8 months.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/10/07/ex2.png" />
// Input: n = 5, relations = [[1,5],[2,5],[3,5],[3,4],[4,5]], time = [1,2,3,4,5]
// Output: 12
// Explanation: The figure above represents the given graph and the time required to complete each course.
// You can start courses 1, 2, and 3 at month 0.
// You can complete them after 1, 2, and 3 months respectively.
// Course 4 can be taken only after course 3 is completed, i.e., after 3 months. It is completed after 3 + 4 = 7 months.
// Course 5 can be taken only after courses 1, 2, 3, and 4 have been completed, i.e., after max(1,2,3,7) = 7 months.
// Thus, the minimum time needed to complete all the courses is 7 + 5 = 12 months.

// Constraints:
//     1 <= n <= 5 * 10^4
//     0 <= relations.length <= min(n * (n - 1) / 2, 5 * 10^4)
//     relations[j].length == 2
//     1 <= prevCoursej, nextCoursej <= n
//     prevCoursej != nextCoursej
//     All the pairs [prevCoursej, nextCoursej] are unique.
//     time.length == n
//     1 <= time[i] <= 10^4
//     The given graph is a directed acyclic graph.

import "fmt"

// bfs
func minimumTime(n int, relations [][]int, time []int) int {
    if n == 0 { return 0 }
    l, graph, degree := len(relations), make([][]int,n), make([]int,n)
    for i := 0; i < l; i++ {
        graph[relations[i][0]-1] = append(graph[relations[i][0]-1],relations[i][1]-1)
        degree[relations[i][1]-1]++
    }
    res, queue, maxTime := 0, []int{}, make([]int, n)
    for i := 0; i < n; i++ {
        if degree[i] == 0 {
            queue = append(queue,i)
            maxTime[i]= time[i]
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for len(queue) > 0 {
        node := queue[0] // pop
        queue = queue[1:]
        for _, adj := range graph[node] {
            maxTime[adj] = max(maxTime[adj], maxTime[node] + time[adj])
            degree[adj]--
            if degree[adj] == 0 {
                queue = append(queue,adj)
            }
        }
    }
    for i := 0; i < n; i++ {
        res = max(res, maxTime[i])
    }
    return res
}

func minimumTime1(n int, relations [][]int, time []int) int {
    graph, inDegree := make([][]int, n+1), make([]int, n+1)
    for _, relation := range relations {
        prev, next := relation[0], relation[1]
        graph[prev] = append(graph[prev], next)
        inDegree[next]++
    }
    res, dp, queue := 0, make([]int, n+1), make([]int, 0)
    for i := 1; i <= n; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
            dp[i] = time[i-1]
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for len(queue) > 0 {
        curr := queue[0] // pop
        queue = queue[1:]
        for _, next := range graph[curr] {
            inDegree[next]--
            dp[next] = max(dp[next], dp[curr]+time[next-1])
            if inDegree[next] == 0 {
                queue = append(queue, next)
            }
        }
        res = max(res, dp[curr])
    }
    return res
}

// dfs
func minimumTime2(n int, relations [][]int, time []int) int {
    res, graph := 0, make([][]int, n+1)
    for _, relation := range relations {
        prev, next := relation[0], relation[1]
        graph[prev] = append(graph[prev], next)
    }
    memo := make([]int, n+1)
    for i := range memo {
        memo[i] = -1
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(x int) int
    dfs = func(x int) int {
        if memo[x] != -1 {
            return memo[x]
        }
        cur := 0
        for _, y := range graph[x] {
            cur = max(cur, dfs(y))
        }
        cur += time[x-1]
        memo[x] = cur
        return cur
    }
    for i := 1; i <= n; i++ {
        res = max(res, dfs(i))
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/10/07/ex1.png" />
    // Input: n = 3, relations = [[1,3],[2,3]], time = [3,2,5]
    // Output: 8
    // Explanation: The figure above represents the given graph and the time required to complete each course. 
    // We start course 1 and course 2 simultaneously at month 0.
    // Course 1 takes 3 months and course 2 takes 2 months to complete respectively.
    // Thus, the earliest time we can start course 3 is at month 3, and the total time required is 3 + 5 = 8 months.
    fmt.Println(minimumTime(3, [][]int{{1,3},{2,3}}, []int{3,2,5})) // 8
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/10/07/ex2.png" />
    // Input: n = 5, relations = [[1,5],[2,5],[3,5],[3,4],[4,5]], time = [1,2,3,4,5]
    // Output: 12
    // Explanation: The figure above represents the given graph and the time required to complete each course.
    // You can start courses 1, 2, and 3 at month 0.
    // You can complete them after 1, 2, and 3 months respectively.
    // Course 4 can be taken only after course 3 is completed, i.e., after 3 months. It is completed after 3 + 4 = 7 months.
    // Course 5 can be taken only after courses 1, 2, 3, and 4 have been completed, i.e., after max(1,2,3,7) = 7 months.
    // Thus, the minimum time needed to complete all the courses is 7 + 5 = 12 months.
    fmt.Println(minimumTime(5, [][]int{{1,5},{2,5},{3,5},{3,4},{4,5}}, []int{1,2,3,4,5})) // 12

    fmt.Println(minimumTime1(3, [][]int{{1,3},{2,3}}, []int{3,2,5})) // 8
    fmt.Println(minimumTime1(5, [][]int{{1,5},{2,5},{3,5},{3,4},{4,5}}, []int{1,2,3,4,5})) // 12

    fmt.Println(minimumTime2(3, [][]int{{1,3},{2,3}}, []int{3,2,5})) // 8
    fmt.Println(minimumTime2(5, [][]int{{1,5},{2,5},{3,5},{3,4},{4,5}}, []int{1,2,3,4,5})) // 12
}