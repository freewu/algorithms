package main

// 3967. Finish Time of Tasks II
// You are given an integer n representing the number of tasks in a project, numbered from 0 to n - 1. 
// These tasks are connected as an undirected tree. 
// This is represented by a 2D integer array edges of length n - 1, where edges[i] = [ui, vi] indicates an undirected connection between task ui and task vi.

// You are also given an array baseTime of length n, where baseTime[i] represents the time to complete task i.

// For any chosen task as the root, the finish time of each task is calculated as follows:
//     1. Leaf task: The finish time is baseTime[i].
//     2. Non-leaf task:
//         2.1 Let earliest be the minimum finish time among its children, and latest be the maximum finish time among its children.
//         2.2 Let ownDuration be (latest - earliest) + baseTime[i].
//         2.3 Finish time of task i is latest + ownDuration.

// Choose any task as the root and compute the finish time of that root based on the rules above.

// Return the minimum possible finish time among all choices of root.

// Example 1:
// Input: n = 3, edges = [[0,1],[1,2]], baseTime = [9,1,5]
// Output: 14
// Explanation:

// //   0 ---- 1 ---- 2
// //  (9)    (1)    (3)

// The optimal choice is to treat task 1 as the root.
// Task 0 is a leaf, so its finish time is baseTime[0] = 9.
// Task 2 is a leaf, so its finish time is baseTime[2] = 5.
// Task 1 has two children with finish times 9 and 5:
// earliest = 5, latest = 9
// ownDuration = (latest - earliest) + baseTime[1] = (9 - 5) + 1 = 5
// Finish time of task 1 is latest + ownDuration = 9 + 5 = 14
// Thus, the minimum possible finish time among all choices of root is 14.

// Example 2:
// Input: n = 3, edges = [[0,1],[0,2]], baseTime = [4,7,6]
// Output: 12
// Explanation:

// //        0
// //       (4)
// //     /     \
// //     1      2
// //    (7)    (6)

// The optimal choice is to treat task 0 as the root.
// Task 1 is a leaf, so its finish time is baseTime[1] = 7.
// Task 2 is a leaf, so its finish time is baseTime[2] = 6.
// Task 0 has two children with finish times 7 and 6:
// earliest = 6, latest = 7
// ownDuration = (latest - earliest) + baseTime[0] = (7 - 6) + 4 = 5
// Finish time of task 0 is latest + ownDuration = 7 + 5 = 12
// Thus, the minimum possible finish time among all choices of root is 12.

// Example 3:
// Input: n = 4, edges = [[0,1],[0,2],[2,3]], baseTime = [5,8,2,1]
// Output: 16
// Explanation:

// //        0
// //       (5)
// //     /     \
// //     1      2
// //    (8)    (2)
//                 \
//                  3
//                 (1)

// The optimal choice is to treat task 1 as the root.
// Task 3 is a leaf, so its finish time is baseTime[3] = 1.
// Task 2 has one child task 3:
// earliest = latest = 1
// ownDuration = (latest - earliest) + baseTime[2] = 0 + 2 = 2
// Finish time of task 2 is latest + ownDuration = 1 + 2 = 3
// Task 0 has one child task 2:
// earliest = latest = 3
// ownDuration = (latest - earliest) + baseTime[0] = 0 + 5 = 5
// Finish time of task 0 is latest + ownDuration = 3 + 5 = 8
// Task 1 has one child task 0:
// earliest = latest = 8
// ownDuration = (latest - earliest) + baseTime[1] = 0 + 8 = 8
// Finish time of task 1 is latest + ownDuration = 8 + 8 = 16
// Thus, the minimum possible finish time among all choices of root is 16.

// Constraints:
//     1 <= n <= 10^5
//     edges.length = n - 1
//     edges[i] == [ui, vi]
//     0 <= ui, vi <= n - 1
//     ui != vi
//     The input is generated such that edges represents a valid undirected tree.
//     baseTime.length == n
//     1 <= baseTime[i] <= 10^5

import "fmt"

func finishTime(n int, edges [][]int, baseTime []int) int64 {
    if n == 1 {
        return int64(baseTime[0])
    }
    // 建无向邻接表
    adj := make([][]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    res, childrenTime, parentTime := int64(1 << 61), make([]int64, n), make([]int64, n)
    // DFS1：自底向上计算仅子树的完成时间 childrenTime
    var dfsDown func(task, p int) int64
    dfsDown = func(task, p int) int64 {
        earliest, latest := int64(1 << 61), int64(-1 << 61)
        for _, next := range adj[task] {
            if next == p {
                continue
            }
            ct := dfsDown(next, task)
            if ct < earliest {
                earliest = ct
            }
            if ct > latest {
                latest = ct
            }
        }
        if earliest > latest {
            // 叶子，无子女
            childrenTime[task] = int64(baseTime[task])
        } else {
            ownDur := latest - earliest + int64(baseTime[task])
            childrenTime[task] = latest + ownDur
        }
        return childrenTime[task]
    }
    dfsDown(0, -1)
    // DFS2：换根，计算父侧时间并更新全局最小完成时间
    var dfsUp func(task, p int)
    dfsUp = func(task, p int) {
        e1, e2, l1, l2 := int64(1 << 61), int64(1 << 61), int64(-1 << 61), int64(-1 << 61) // 最小第一、第二 最大第一、 第二
        // 遍历所有邻接点，收集有效时间（子树用childrenTime，父侧用parentTime[task]）
        for _, next := range adj[task] {
            t := int64(0)
            if next == p {
                t = parentTime[task]
            } else {
                t = childrenTime[next]
            }
            // 更新最小双值
            if t < e1 {
                e2 = e1
                e1 = t
            } else if t < e2 {
                e2 = t
            }
            // 更新最大双值
            if t > l1 {
                l2 = l1
                l1 = t
            } else if t > l2 {
                l2 = t
            }
        }
        // 当前task作为全局根的完成时间，更新最小值
        ownDur := l1 - e1 + int64(baseTime[task])
        totalFT := l1 + ownDur
        if totalFT < res {
            res = totalFT
        }
        // 遍历所有子节点，预计算子节点的parentTime，递归
        for _, child := range adj[task] {
            if child == p {
                continue
            }
            ct := childrenTime[child]
            var pe, pl int64
            if len(adj[task]) == 1 {
                // 当前节点只有这一个子，移除后无其他分支，父侧等效为自身baseTime
                parentTime[child] = int64(baseTime[task])
            } else {
                // 剔除当前child，取剩余集合的min/max
                if ct != e1 {
                    pe = e1
                } else {
                    pe = e2
                }
                if ct != l1 {
                    pl = l1
                } else {
                    pl = l2
                }
                childOwnDur := pl - pe + int64(baseTime[task])
                parentTime[child] = pl + childOwnDur
            }
            dfsUp(child, task)
        }
    }
    dfsUp(0, -1)
    return res
}

func main() {
    // Example 1:
    // Input: n = 3, edges = [[0,1],[1,2]], baseTime = [9,1,5]
    // Output: 14
    // Explanation:
    // //   0 ---- 1 ---- 2
    // //  (9)    (1)    (3)
    // The optimal choice is to treat task 1 as the root.
    // Task 0 is a leaf, so its finish time is baseTime[0] = 9.
    // Task 2 is a leaf, so its finish time is baseTime[2] = 5.
    // Task 1 has two children with finish times 9 and 5:
    // earliest = 5, latest = 9
    // ownDuration = (latest - earliest) + baseTime[1] = (9 - 5) + 1 = 5
    // Finish time of task 1 is latest + ownDuration = 9 + 5 = 14
    // Thus, the minimum possible finish time among all choices of root is 14.
    fmt.Println(finishTime(3, [][]int{{0,1},{1,2}}, []int{9,1,5})) // 14
    // Example 2:
    // Input: n = 3, edges = [[0,1],[0,2]], baseTime = [4,7,6]
    // Output: 12
    // Explanation:
    // //        0
    // //       (4)
    // //     /     \
    // //     1      2
    // //    (7)    (6)
    // The optimal choice is to treat task 0 as the root.
    // Task 1 is a leaf, so its finish time is baseTime[1] = 7.
    // Task 2 is a leaf, so its finish time is baseTime[2] = 6.
    // Task 0 has two children with finish times 7 and 6:
    // earliest = 6, latest = 7
    // ownDuration = (latest - earliest) + baseTime[0] = (7 - 6) + 4 = 5
    // Finish time of task 0 is latest + ownDuration = 7 + 5 = 12
    // Thus, the minimum possible finish time among all choices of root is 12.
    fmt.Println(finishTime(3, [][]int{{0,1},{0,2}}, []int{4,7,6})) // 12
    // Example 3:
    // Input: n = 4, edges = [[0,1],[0,2],[2,3]], baseTime = [5,8,2,1]
    // Output: 16
    // Explanation:
    // //        0
    // //       (5)
    // //     /     \
    // //     1      2
    // //    (8)    (2)
    //                 \
    //                  3
    //                 (1)
    // The optimal choice is to treat task 1 as the root.
    // Task 3 is a leaf, so its finish time is baseTime[3] = 1.
    // Task 2 has one child task 3:
    // earliest = latest = 1
    // ownDuration = (latest - earliest) + baseTime[2] = 0 + 2 = 2
    // Finish time of task 2 is latest + ownDuration = 1 + 2 = 3
    // Task 0 has one child task 2:
    // earliest = latest = 3
    // ownDuration = (latest - earliest) + baseTime[0] = 0 + 5 = 5
    // Finish time of task 0 is latest + ownDuration = 3 + 5 = 8
    // Task 1 has one child task 0:
    // earliest = latest = 8
    // ownDuration = (latest - earliest) + baseTime[1] = 0 + 8 = 8
    // Finish time of task 1 is latest + ownDuration = 8 + 8 = 16
    // Thus, the minimum possible finish time among all choices of root is 16.
    fmt.Println(finishTime(4, [][]int{{0,1},{0,2},{2,3}}, []int{5,8,2,1})) // 16
}