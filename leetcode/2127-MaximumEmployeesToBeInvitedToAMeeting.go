package main

// 2127. Maximum Employees to Be Invited to a Meeting
// A company is organizing a meeting and has a list of n employees, waiting to be invited. 
// They have arranged for a large circular table, capable of seating any number of employees.

// The employees are numbered from 0 to n - 1. 
// Each employee has a favorite person and they will attend the meeting only if they can sit next to their favorite person at the table. 
// The favorite person of an employee is not themself.

// Given a 0-indexed integer array favorite, where favorite[i] denotes the favorite person of the ith employee, 
// return the maximum number of employees that can be invited to the meeting.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/12/14/ex1.png" />
// Input: favorite = [2,2,1,2]
// Output: 3
// Explanation:
// The above figure shows how the company can invite employees 0, 1, and 2, and seat them at the round table.
// All employees cannot be invited because employee 2 cannot sit beside employees 0, 1, and 3, simultaneously.
// Note that the company can also invite employees 1, 2, and 3, and give them their desired seats.
// The maximum number of employees that can be invited to the meeting is 3. 

// Example 2:
// Input: favorite = [1,2,0]
// Output: 3
// Explanation: 
// Each employee is the favorite person of at least one other employee, and the only way the company can invite them is if they invite every employee.
// The seating arrangement will be the same as that in the figure given in example 1:
// - Employee 0 will sit between employees 2 and 1.
// - Employee 1 will sit between employees 0 and 2.
// - Employee 2 will sit between employees 1 and 0.
// The maximum number of employees that can be invited to the meeting is 3.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/12/14/ex2.png" />
// Input: favorite = [3,0,1,4,1]
// Output: 4
// Explanation:
// The above figure shows how the company will invite employees 0, 1, 3, and 4, and seat them at the round table.
// Employee 2 cannot be invited because the two spots next to their favorite employee 1 are taken.
// So the company leaves them out of the meeting.
// The maximum number of employees that can be invited to the meeting is 4.

// Constraints:
//     n == favorite.length
//     2 <= n <= 10^5
//     0 <= favorite[i] <= n - 1
//     favorite[i] != i

import "fmt"

// topological sort
func maximumInvitations(favorite []int) int {
    n := len(favorite)
    indegrees, queue := make([]int, n), []int{}
    for _, v := range favorite {
        indegrees[v]++
    }
    for i, v := range indegrees {
        if v == 0 {
            queue = append(queue, i)
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    i, dp := 0, make([]int, n) // dp[i] is the longest path leading to i exclusively.
    for len(queue) != 0 {
        i, queue = queue[0], queue[1:] // remove front
        j := favorite[i]
        dp[j] = max(dp[j], dp[i] + 1)
        indegrees[j]--
        if indegrees[j] == 0 {
            queue = append(queue, j)
        }
    }
    res, res2 := 0, 0
    for i = 0; i < n; i++ {
        if indegrees[i] != 0 {
            count := 0
            for j := i; indegrees[j] > 0; j = favorite[j] {
                indegrees[j] = 0
                count++
            }
            if count == 2 {
                res2 += 2 + dp[i] + dp[favorite[i]]
            } else {
                res = max(res, count)
            }
        }
    }
    return max(res, res2)
}

// 基环树和基环树森林
// n 节点 n 条边的有向图必定有环，分别考虑基环长度为2和大于2的情况
// 1. 基环为2：左右两边插入从基环两个节点出发的最长路径上的节点
// 2. 基环大于2：无法插入其他节点
func maximumInvitations1(favorite []int) int {
    n := len(favorite) // 节点数量
    deg := make([]int, n) // 每个节点上的入度数
    maxDepth := make([]int, n) // 从每个节点出发反向的搜索的最长路径上的节点数
    //拓扑排序
    queue := []int{}
    for _, v := range favorite {
        deg[v]++
    }
    for i, in := range deg {
        if in == 0 {
            queue = append(queue, i)
        }
    }
    for len(queue) != 0 {
        pre := queue[0]
        queue = queue[1:] // pop
        next := favorite[pre]
        maxDepth[next] = maxDepth[pre] + 1
        if deg[next]--; deg[next] == 0 {
            queue = append(queue, next)
        }
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    maxRingSize, sumChainSize := 0, 0
    for i,in := range deg {
        if in == 0 { continue }
        // 遍历当前连通分量上基环上的节点
        deg[i]=0
        ringSize := 1
        for x := favorite[i]; x != i ; x = favorite[x] {
            deg[x]=0
            ringSize++
        }
        if ringSize == 2 {
            sumChainSize += maxDepth[i] + maxDepth[favorite[i]] + 2 // 累加两条最长链的长度
        } else {
            maxRingSize = max(maxRingSize, ringSize) // 更新最长基环长度
        }
    }
    return max(maxRingSize, sumChainSize)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/12/14/ex1.png" />
    // Input: favorite = [2,2,1,2]
    // Output: 3
    // Explanation:
    // The above figure shows how the company can invite employees 0, 1, and 2, and seat them at the round table.
    // All employees cannot be invited because employee 2 cannot sit beside employees 0, 1, and 3, simultaneously.
    // Note that the company can also invite employees 1, 2, and 3, and give them their desired seats.
    // The maximum number of employees that can be invited to the meeting is 3. 
    fmt.Println(maximumInvitations([]int{2,2,1,2})) // 3
    // Example 2:
    // Input: favorite = [1,2,0]
    // Output: 3
    // Explanation: 
    // Each employee is the favorite person of at least one other employee, and the only way the company can invite them is if they invite every employee.
    // The seating arrangement will be the same as that in the figure given in example 1:
    // - Employee 0 will sit between employees 2 and 1.
    // - Employee 1 will sit between employees 0 and 2.
    // - Employee 2 will sit between employees 1 and 0.
    // The maximum number of employees that can be invited to the meeting is 3.
    fmt.Println(maximumInvitations([]int{1,2,0})) // 3
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/12/14/ex2.png" />
    // Input: favorite = [3,0,1,4,1]
    // Output: 4
    // Explanation:
    // The above figure shows how the company will invite employees 0, 1, 3, and 4, and seat them at the round table.
    // Employee 2 cannot be invited because the two spots next to their favorite employee 1 are taken.
    // So the company leaves them out of the meeting.
    // The maximum number of employees that can be invited to the meeting is 4.
    fmt.Println(maximumInvitations([]int{3,0,1,4,1})) // 4

    fmt.Println(maximumInvitations1([]int{2,2,1,2})) // 3
    fmt.Println(maximumInvitations1([]int{1,2,0})) // 3
    fmt.Println(maximumInvitations1([]int{3,0,1,4,1})) // 4
}