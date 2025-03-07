package main

// 3476. Maximize Profit from Task Assignment
// You are given an integer array workers, where workers[i] represents the skill level of the ith worker. 
// You are also given a 2D integer array tasks, where:
//     1. tasks[i][0] represents the skill requirement needed to complete the task.
//     2. tasks[i][1] represents the profit earned from completing the task.

// Each worker can complete at most one task, and they can only take a task if their skill level is equal to the task's skill requirement. 
// An additional worker joins today who can take up any task, regardless of the skill requirement.

// Return the maximum total profit that can be earned by optimally assigning the tasks to the workers.

// Example 1:
// Input: workers = [1,2,3,4,5], tasks = [[1,100],[2,400],[3,100],[3,400]]
// Output: 1000
// Explanation:
// Worker 0 completes task 0.
// Worker 1 completes task 1.
// Worker 2 completes task 3.
// The additional worker completes task 2.

// Example 2:
// Input: workers = [10,10000,100000000], tasks = [[1,100]]
// Output: 100
// Explanation:
// Since no worker matches the skill requirement, only the additional worker can complete task 0.

// Example 3:
// Input: workers = [7], tasks = [[3,3],[3,3]]
// Output: 3
// Explanation:
// The additional worker completes task 1. Worker 0 cannot work since no task has a skill requirement of 7.

// Constraints:
//     1 <= workers.length <= 10^5
//     1 <= workers[i] <= 10^9
//     1 <= tasks.length <= 10^5
//     tasks[i].length == 2
//     1 <= tasks[i][0], tasks[i][1] <= 10^9

import "fmt"
import "sort"

func maxProfit(workers []int, tasks [][]int) int64 {
    mp := make(map[int]int)
    for _, v := range workers {
        mp[v]++
    }
    mp_task := make(map[int][]int)
    for _, t := range tasks {
        mp_task[t[0]] = append(mp_task[t[0]], t[1])
    }
    sum, mx := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for k, v := range mp_task {
        tmp := v
        worker_num := mp[k]
        sort.Ints(tmp)
        index := len(tmp) - 1
        for index >= 0 && worker_num > 0 {
            sum += tmp[index]
            index--
            worker_num--
        }
        if index >= 0 {
            mx = max(mx, tmp[index])
        }
    }
    return int64(sum + mx)
}

func main() {
    // Example 1:
    // Input: workers = [1,2,3,4,5], tasks = [[1,100],[2,400],[3,100],[3,400]]
    // Output: 1000
    // Explanation:
    // Worker 0 completes task 0.
    // Worker 1 completes task 1.
    // Worker 2 completes task 3.
    // The additional worker completes task 2.
    fmt.Println(maxProfit([]int{1,2,3,4,5}, [][]int{{1,100},{2,400},{3,100},{3,400}})) // 1000
    // Example 2:
    // Input: workers = [10,10000,100000000], tasks = [[1,100]]
    // Output: 100
    // Explanation:
    // Since no worker matches the skill requirement, only the additional worker can complete task 0.
    fmt.Println(maxProfit([]int{10,10000,100000000}, [][]int{{1,100}})) // 100
    // Example 3:
    // Input: workers = [7], tasks = [[3,3],[3,3]]
    // Output: 3
    // Explanation:
    // The additional worker completes task 1. Worker 0 cannot work since no task has a skill requirement of 7.
    fmt.Println(maxProfit([]int{7}, [][]int{{3,3},{3,3}})) // 3

    fmt.Println(maxProfit([]int{1,2,3,4,5,6,7,8,9}, [][]int{{3,3},{3,3}})) // 6
    fmt.Println(maxProfit([]int{9,8,7,6,5,4,3,2,1}, [][]int{{3,3},{3,3}})) // 6
}