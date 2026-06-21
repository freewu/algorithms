package main

// 3965. Finish Time of Tasks I
// You are given an integer n representing the number of tasks in a project, numbered from 0 to n - 1. 
// These tasks are connected as a tree rooted at task 0. 
// This is represented by a 2D integer array edges of length n - 1, where edges[i] = [ui, vi] indicates that task ui is the parent of task vi.

// You are also given an array baseTime of length n, where baseTime[i] represents the time to complete task i.

// The finish time of each task is calculated as follows:
//     1. Leaf task: The finish time is baseTime[i].
//     2. Non-leaf task:
//         2.1 Let earliest be the minimum finish time among its children, and latest be the maximum finish time among its children.
//         2.2 Let ownDuration be (latest - earliest) + baseTime[i].
//         2.3 The finish time of task i is latest + ownDuration.

// Return the finish time of the root task 0.

// Example 1:
// Input: n = 3, edges = [[0,1],[1,2]], baseTime = [9,5,3]
// Output: 17
// Explanation:

//   0 ---- 1 ---- 2
//  (9)    (5)    (3)
 
// Task 2 is a leaf, so its finish time is baseTime[2] = 3.
// Task 1 has one child task 2:
// earliest = latest = 3
// ownDuration = (latest - earliest) + baseTime[1] = 5
// Finish time of task 1 is 3 + 5 = 8
// Task 0 has one child with finish time 8:
// earliest = latest = 8
// ownDuration = (latest - earliest) + baseTime[0] = 9
// Finish time of task 0 is 8 + 9 = 17

// Example 2:
// Input: n = 3, edges = [[0,1],[0,2]], baseTime = [4,7,6]
// Output: 12
// Explanation:

//         0
//        (4)
//     /     \
//     1      2
//    (7)    (6)

// Task 1 is a leaf, so its finish time is baseTime[1] = 7.
// Task 2 is a leaf, so its finish time is baseTime[2] = 6.
// Task 0 has two children with finish times 7 and 6:
// earliest = 6, latest = 7
// ownDuration = (latest - earliest) + baseTime[0] = (7 - 6) + 4 = 5
// Finish time of task 0 is latest + ownDuration = 7 + 5 = 12

// Example 3:
// Input: n = 4, edges = [[0,1],[0,2],[2,3]], baseTime = [5,8,2,1]
// Output: 18
// Explanation:
// Task 1 is a leaf, so its finish time is baseTime[1] = 8.
// Task 3 is a leaf, so its finish time is baseTime[3] = 1.
// Task 2 has one child task 3:
// earliest = latest = 1
// ownDuration = (latest - earliest) + baseTime[2] = 0 + 2 = 2
// Finish time of task 2 is latest + ownDuration = 1 + 2 = 3
// Task 0 has two children with finish times 8 and 3:
// earliest = 3, latest = 8
// ownDuration = (latest - earliest) + baseTime[0] = (8 - 3) + 5 = 10
// Finish time of task 0 is latest + ownDuration = 8 + 10 = 18

// Constraints:
//     1 <= n <= 10^5
//     edges.length = n - 1
//     edges[i] == [ui, vi]
//     0 <= ui, vi <= n - 1
//     ui != vi
//     The input is generated such that edges represents a valid tree.
//     baseTime.length == n
//     1 <= baseTime[i] <= 10^5​​​​​​​

import "fmt"

func finishTime(n int, edges [][]int, baseTime []int) int64 {
    graph := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1]
        graph[x] = append(graph[x], y) // 题目保证 x 是 y 的父节点
    }
    var dfs func(x int) int
    dfs = func(x int) int {
        if graph[x] == nil { // x 是叶子
            return baseTime[x]
        }
        earliest, latest := 1 << 61, 0
        for _, y := range graph[x] {
            t := dfs(y)
            earliest = min(earliest, t)
            latest = max(latest, t)
        }
        return latest * 2 - earliest + baseTime[x]
    }
    return int64(dfs(0))
}

func finishTime1(n int, edges [][]int, baseTime []int) int64 {
    children, parents, outDegree := make([][]int, n), make([][]int, n), make([]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        children[u], parents[v] = append(children[u], v), append(parents[v], u)
        outDegree[u]++
    }
    finish, maxChild, minChild := make([]int64, n), make([]int64, n), make([]int64, n)
    for i := 0; i < n; i++ {
        minChild[i] = 1 << 61
    }
    res, queue := int64(0), make([]int, 0)
    // 叶子
    for i := 0; i < n; i++ {
        if outDegree[i] == 0 {
            finish[i] = int64(baseTime[i])
            queue = append(queue, i)
        }
    }
    for head := 0; head < len(queue); head++ {
        curr := queue[head]
        for _, p := range parents[curr] {
            if finish[curr] > maxChild[p] {
                maxChild[p] = finish[curr]
            }
            if finish[curr] < minChild[p] {
                minChild[p] = finish[curr]
            }
            outDegree[p]--
            // 所有孩子处理完
            if outDegree[p] == 0 {
                finish[p] = 2 * maxChild[p] - minChild[p] + int64(baseTime[p])
                queue = append(queue, p)
            }
        }
    }
    // 如果要求整个工程完成时间
    // 取所有根节点最大值
    hasParent := make([]bool, n)
    for _, e := range edges {
        hasParent[e[1]] = true
    }
    for i := 0; i < n; i++ {
        if !hasParent[i] {
            if finish[i] > res {
                res = finish[i]
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, edges = [[0,1],[1,2]], baseTime = [9,5,3]
    // Output: 17
    // Explanation:
    //   0 ---- 1 ---- 2
    //  (9)    (5)    (3)
    // Task 2 is a leaf, so its finish time is baseTime[2] = 3.
    // Task 1 has one child task 2:
    // earliest = latest = 3
    // ownDuration = (latest - earliest) + baseTime[1] = 5
    // Finish time of task 1 is 3 + 5 = 8
    // Task 0 has one child with finish time 8:
    // earliest = latest = 8
    // ownDuration = (latest - earliest) + baseTime[0] = 9
    // Finish time of task 0 is 8 + 9 = 17
    fmt.Println(finishTime(3, [][]int{{0,1},{1,2}}, []int{9,5,3})) // 17
    // Example 2:
    // Input: n = 3, edges = [[0,1],[0,2]], baseTime = [4,7,6]
    // Output: 12
    // Explanation:
    //         0
    //        (4)
    //     /     \
    //     1      2
    //    (7)    (6)
    // Task 1 is a leaf, so its finish time is baseTime[1] = 7.
    // Task 2 is a leaf, so its finish time is baseTime[2] = 6.
    // Task 0 has two children with finish times 7 and 6:
    // earliest = 6, latest = 7
    // ownDuration = (latest - earliest) + baseTime[0] = (7 - 6) + 4 = 5
    // Finish time of task 0 is latest + ownDuration = 7 + 5 = 12
    fmt.Println(finishTime(3, [][]int{{0,1},{0,2}}, []int{4,7,6})) // 12
    // Example 3:
    // Input: n = 4, edges = [[0,1],[0,2],[2,3]], baseTime = [5,8,2,1]
    // Output: 18
    // Explanation:
    // Task 1 is a leaf, so its finish time is baseTime[1] = 8.
    // Task 3 is a leaf, so its finish time is baseTime[3] = 1.
    // Task 2 has one child task 3:
    // earliest = latest = 1
    // ownDuration = (latest - earliest) + baseTime[2] = 0 + 2 = 2
    // Finish time of task 2 is latest + ownDuration = 1 + 2 = 3
    // Task 0 has two children with finish times 8 and 3:
    // earliest = 3, latest = 8
    // ownDuration = (latest - earliest) + baseTime[0] = (8 - 3) + 5 = 10
    // Finish time of task 0 is latest + ownDuration = 8 + 10 = 18
    fmt.Println(finishTime(4, [][]int{{0,1},{0,2},{2,3},}, []int{5,8,2,1})) // 18

    fmt.Println(finishTime1(3, [][]int{{0,1},{1,2}}, []int{9,5,3})) // 17
    fmt.Println(finishTime1(3, [][]int{{0,1},{0,2}}, []int{4,7,6})) // 12
    fmt.Println(finishTime1(4, [][]int{{0,1},{0,2},{2,3},}, []int{5,8,2,1})) // 18
}