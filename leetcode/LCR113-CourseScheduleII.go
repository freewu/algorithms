package main

// LCR 113. 课程表 II
// 现在总共有 numCourses 门课需要选，记为 0 到 numCourses-1。

// 给定一个数组 prerequisites ，它的每一个元素 prerequisites[i] 表示两门课程之间的先修顺序。 
// 例如 prerequisites[i] = [ai, bi] 表示想要学习课程 ai ，需要先完成课程 bi 。

// 请根据给出的总课程数  numCourses 和表示先修顺序的 prerequisites 得出一个可行的修课序列。

// 可能会有多个正确的顺序，只要任意返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

// 示例 1:
// 输入: numCourses = 2, prerequisites = [[1,0]] 
// 输出: [0,1]
// 解释: 总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。

// 示例 2:
// 输入: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
// 输出: [0,1,2,3] or [0,2,1,3]
// 解释: 总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
// 因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。

// 示例 3:
// 输入: numCourses = 1, prerequisites = [] 
// 输出: [0]
// 解释: 总共 1 门课，直接修第一门课就可。

// 提示:
//     1 <= numCourses <= 2000
//     0 <= prerequisites.length <= numCourses * (numCourses - 1)
//     prerequisites[i].length == 2
//     0 <= ai, bi < numCourses
//     ai != bi
//     prerequisites 中不存在重复元素

import "fmt"

// Kahn's Algorithm
func findOrder(numCourses int, prerequisites [][]int) []int {
    graph, inDegrees := make([][]int, numCourses), make([]int, numCourses)
    for _, prerequisite := range(prerequisites) {
        graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
        inDegrees[prerequisite[0]]++
    }
    queue, res := []int{}, []int{}
    for i, inDegree := range(inDegrees) {
        if inDegree == 0 {
            queue = append(queue, i)
        }
    }
    for len(queue) > 0 {
        lenQueue := len(queue)
        for i := 0; i < lenQueue; i++ {
            course := queue[0]
            queue = queue[1:]
            res = append(res, course)
            for _, nextCourse := range(graph[course]) {
                inDegrees[nextCourse]--
                if inDegrees[nextCourse] == 0 {
                    queue = append(queue, nextCourse)
                }
            }
        }
    }
    if len(res) == numCourses {
        return res
    }
    return []int{}
}

// bfs
func findOrder1(numCourses int, prerequisites [][]int) []int {
    buildGrapth := func(numCourses int, prerequisites [][]int) (grapth [][]int, indegree []int) {
        grapth = make([][]int, numCourses)
        indegree = make([]int, numCourses)
        for _, prerequisite := range prerequisites{
            from, to := prerequisite[1], prerequisite[0]
            indegree[to] += 1
            grapth[from] = append(grapth[from], to)
        }
        return grapth, indegree
    }
    // 使用 BFS 入度来处理
    grapth, indegree := buildGrapth(numCourses, prerequisites)
    res := make([]int, 0)
    var bfs func(grapth [][]int, indegree []int)
    bfs = func(grapth [][]int, indegree []int) {
        // 先得到入度为0的所有节点列表
        queue := make([]int, 0)
        for node, val := range indegree{
            if val == 0 {
                queue = append(queue, node)
            }
        }
        // 入度为0的节点，都可以选择
        for len(queue) > 0 {
            // 可选课程
            course := queue[0]
            queue = queue[1:]
            res = append(res, course)
            // 将course相关的所有连接节点的入度设置为0
            for _, edgeNode := range grapth[course] {
                indegree[edgeNode] -= 1
                if indegree[edgeNode] == 0 {
                    queue = append(queue, edgeNode)
                }
            }

        }
    } 
    bfs(grapth, indegree)
    if len(res) != numCourses {
        return []int{}
    }
    return res
}

func main() {
    // Example 1:
    // Input: numCourses = 2, prerequisites = [[1,0]]
    // Output: [0,1]
    // Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].
    fmt.Println(findOrder(2,[][]int{{1,0}})) // [0,1]
    // Example 2:
    // Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
    // Output: [0,2,1,3]
    // Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
    // So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].
    fmt.Println(findOrder(4,[][]int{{1,0},{2,0},{3,1},{3,2}})) // [0,2,1,3]
    // Example 3:
    // Input: numCourses = 1, prerequisites = []
    // Output: [0]
    fmt.Println(findOrder(1,[][]int{})) // [0]

    
    fmt.Println(findOrder1(2,[][]int{{1,0}})) // [0,1]
    fmt.Println(findOrder1(4,[][]int{{1,0},{2,0},{3,1},{3,2}})) // [0,2,1,3]
    fmt.Println(findOrder1(1,[][]int{})) // [0]
}