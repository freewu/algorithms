package main

// 1376. Time Needed to Inform All Employees
// A company has n employees with a unique ID for each employee from 0 to n - 1. 
// The head of the company is the one with headID.

// Each employee has one direct manager given in the manager array where manager[i] is the direct manager of the i-th employee, manager[headID] = -1. 
// Also, it is guaranteed that the subordination relationships have a tree structure.

// The head of the company wants to inform all the company employees of an urgent piece of news. 
// He will inform his direct subordinates, and they will inform their subordinates, 
// and so on until all employees know about the urgent news.

// The i-th employee needs informTime[i] minutes to inform all of his direct subordinates 
// (i.e., After informTime[i] minutes, all his direct subordinates can start spreading the news).

// Return the number of minutes needed to inform all the employees about the urgent news.

// Example 1:
// Input: n = 1, headID = 0, manager = [-1], informTime = [0]
// Output: 0
// Explanation: The head of the company is the only employee in the company.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/02/27/graph.png" />
// Input: n = 6, headID = 2, manager = [2,2,-1,2,2,2], informTime = [0,0,1,0,0,0]
// Output: 1
// Explanation: The head of the company with id = 2 is the direct manager of all the employees in the company and needs 1 minute to inform them all.
// The tree structure of the employees in the company is shown.

// Constraints:
//     1 <= n <= 10^5
//     0 <= headID < n
//     manager.length == n
//     0 <= manager[i] < n
//     manager[headID] == -1
//     informTime.length == n
//     0 <= informTime[i] <= 1000
//     informTime[i] == 0 if employee i has no subordinates.
//     It is guaranteed that all the employees can be informed.

import "fmt"

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    m := make(map[int][]int, n)
    for i, v := range manager { // 将 map[管理者] = []int{员工ID}
        if v != -1 {
            m[v] = append(m[v], i)
        }
    }
    var dfs func(id int) int
    dfs = func (id int) int {
        if len(m[id]) == 0 {
            return 0
        }
        curMaxTime := 0
        for _, v := range m[id] { // 通知下属员工
            if current := dfs(v) + informTime[id]; current > curMaxTime {
                curMaxTime = current
            }
        }
        return curMaxTime
    }
    return dfs(headID)
}

func numOfMinutes1(n int, headID int, manager, informTime []int) int {
    res := 0
    for i, m := range manager {
        if m < 0 {
            continue
        }
        // 计算从 i 向上的累加值
        s, x := 0, i
        for ; manager[x] >= 0; x = manager[x] {
            s += informTime[x]
        }
        // 此时 x 要么是 headID，要么是一个计算过的节点
        s += informTime[x]
        res = max(res, s)
        // 记录从 i 向上的每个未被计算的节点值的对应累加值
        for x = i; manager[x] >= 0; {
            informTime[x], s = s, s-informTime[x]
            manager[x], x = -1, manager[x]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: n = 1, headID = 0, manager = [-1], informTime = [0]
    // Output: 0
    // Explanation: The head of the company is the only employee in the company.
    fmt.Println(numOfMinutes(1, 0,[]int{-1},[]int{0})) // 0
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/02/27/graph.png" />
    // Input: n = 6, headID = 2, manager = [2,2,-1,2,2,2], informTime = [0,0,1,0,0,0]
    // Output: 1
    // Explanation: The head of the company with id = 2 is the direct manager of all the employees in the company and needs 1 minute to inform them all.
    // The tree structure of the employees in the company is shown.
    fmt.Println(numOfMinutes(6, 2,[]int{-2,2,-1,2,2,2},[]int{0,0,1,0,0,0})) // 0

    fmt.Println(numOfMinutes1(1, 0,[]int{-1},[]int{0})) // 0
    fmt.Println(numOfMinutes1(6, 2,[]int{-2,2,-1,2,2,2},[]int{0,0,1,0,0,0})) // 0
}