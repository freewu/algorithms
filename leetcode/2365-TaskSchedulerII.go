package main

// 2365. Task Scheduler II
// You are given a 0-indexed array of positive integers tasks, 
// representing tasks that need to be completed in order, 
// where tasks[i] represents the type of the ith task.

// You are also given a positive integer space, 
// which represents the minimum number of days that must pass after the completion of a task before another task of the same type can be performed.

// Each day, until all tasks have been completed, you must either:
//     1. Complete the next task from tasks, or
//     2. Take a break.

// Return the minimum number of days needed to complete all tasks.

// Example 1:
// Input: tasks = [1,2,1,2,3,1], space = 3
// Output: 9
// Explanation:
// One way to complete all tasks in 9 days is as follows:
// Day 1: Complete the 0th task.
// Day 2: Complete the 1st task.
// Day 3: Take a break.
// Day 4: Take a break.
// Day 5: Complete the 2nd task.
// Day 6: Complete the 3rd task.
// Day 7: Take a break.
// Day 8: Complete the 4th task.
// Day 9: Complete the 5th task.
// It can be shown that the tasks cannot be completed in less than 9 days.

// Example 2:
// Input: tasks = [5,8,8,5], space = 2
// Output: 6
// Explanation:
// One way to complete all tasks in 6 days is as follows:
// Day 1: Complete the 0th task.
// Day 2: Complete the 1st task.
// Day 3: Take a break.
// Day 4: Take a break.
// Day 5: Complete the 2nd task.
// Day 6: Complete the 3rd task.
// It can be shown that the tasks cannot be completed in less than 6 days.

// Constraints:
//     1 <= tasks.length <= 10^5
//     1 <= tasks[i] <= 10^9
//     1 <= space <= tasks.length

import "fmt"

func taskSchedulerII(tasks []int, space int) int64 {
    res, mp := 0, make(map[int]int)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, task := range tasks {
        day, ok := mp[task]
        if ok {
            res = max(res, day + space)
        }
        res++
        mp[task] = res
    }
    return int64(res)
}

func taskSchedulerII1(tasks []int, space int) int64 {
    res, mp := int64(1), map[int]int64{}
    for _, task := range tasks {
        if mp[task] > res {
            res = mp[task]
        }
        mp[task] = res + int64(space) + 1
        res++
    }
    return res - 1
}

func main() {
    // Example 1:
    // Input: tasks = [1,2,1,2,3,1], space = 3
    // Output: 9
    // Explanation:
    // One way to complete all tasks in 9 days is as follows:
    // Day 1: Complete the 0th task.
    // Day 2: Complete the 1st task.
    // Day 3: Take a break.
    // Day 4: Take a break.
    // Day 5: Complete the 2nd task.
    // Day 6: Complete the 3rd task.
    // Day 7: Take a break.
    // Day 8: Complete the 4th task.
    // Day 9: Complete the 5th task.
    // It can be shown that the tasks cannot be completed in less than 9 days.
    fmt.Println(taskSchedulerII([]int{1,2,1,2,3,1}, 3)) // 9
    // Example 2:
    // Input: tasks = [5,8,8,5], space = 2
    // Output: 6
    // Explanation:
    // One way to complete all tasks in 6 days is as follows:
    // Day 1: Complete the 0th task.
    // Day 2: Complete the 1st task.
    // Day 3: Take a break.
    // Day 4: Take a break.
    // Day 5: Complete the 2nd task.
    // Day 6: Complete the 3rd task.
    // It can be shown that the tasks cannot be completed in less than 6 days.
    fmt.Println(taskSchedulerII([]int{5,8,8,5}, 2)) // 6

    fmt.Println(taskSchedulerII1([]int{1,2,1,2,3,1}, 3)) // 9
    fmt.Println(taskSchedulerII1([]int{5,8,8,5}, 2)) // 6
}