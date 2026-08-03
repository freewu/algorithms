package main

// 4012. Count of Unfinished Tasks After Each Shift
// You are given two integer arrays tasks and shifts.
//     1. tasks[i] represents the time required to complete the ith task.
//     2. shifts[j] represents the amount of time available during the jth shift.

// The tasks must be processed in order from left to right.
//     1. Carry-over: If a task is not completed during a shift, processing continues from the same point in that task during the next shift.
//     2. Restart: If all tasks are completed during a shift, the shift ends immediately. 
//                 Any unused time in that shift is discarded, and the next shift begins again from task 0.

// A task is unfinished if it has not been fully completed. This includes a task that is currently in progress.

// Return an integer array ans where ans[j] is the number of unfinished tasks immediately after the jth shift.

// Example 1:
// Input: tasks = [1,4,4], shifts = [9,1,4]
// Output: [0,2,1]
// Explanation:
// Shift 0: The tasks require 1 + 4 + 4 = 9 units of time, so all tasks are completed. There are 0 unfinished tasks.
// Shift 1: Processing restarts from task 0. The shift has time 1, so task 0 is completed. There are 2 unfinished tasks.
// Shift 2: Processing continues from task 1. The shift has time 4, so task 1 is completed. There is 1 unfinished task.

// Example 2:
// Input: tasks = [2,3,4], shifts = [20,4,5]
// Output: [0,2,0]
// Explanation:
// Shift 0: The tasks require 2 + 3 + 4 = 9 units of time, so all tasks are completed. The remaining time in this shift is ignored. There are 0 unfinished tasks.
// Shift 1: Processing restarts from task 0. The shift has time 4, so task 0 is completed and task 1 is partially completed. There are 2 unfinished tasks.
// Shift 2: Processing continues from task 1. The remaining time needed is 1 + 4 = 5, so all tasks are completed. There are 0 unfinished tasks.

// Example 3:
// Input: tasks = [4,2], shifts = [3,6,1]
// Output: [2,0,2]
// Explanation:
// Shift 0: The shift has time 3, so task 0 is partially completed with 1 unit of work remaining. There are 2 unfinished tasks.
// Shift 1: Processing continues from task 0. The remaining time needed is 1 + 2 = 3, so all tasks are completed. There are 0 unfinished tasks.
// Shift 2: Processing restarts from task 0. The shift has time 1, so task 0 is partially completed. There are 2 unfinished tasks.

// Constraints:
//     1 <= tasks.length <= 10^5
//     1 <= shifts.length <= 10^5
//     1 <= tasks[i] <= 10^9
//     1 <= shifts[i] <= 10^9

import "fmt"

func countTasks(tasks []int, shifts []int) []int {
    for i := range tasks {
        if i == 0 { continue }
        tasks[i] += tasks[i - 1]
    }
    res, done, taskAt := make([]int, 0), 0, 0
    binarySearch := func(arr []int, key, i int) int {
        res, j :=  i - 1, len(arr) - 1
        for i <= j {
            mid := i + (j - i) / 2
            if arr[mid] <= key {
                res = mid
                i = mid + 1
            } else {
                j = mid - 1
            }
        }
        return res
    } 
    for _, shift := range shifts {
        done += shift
        till := binarySearch(tasks, done, taskAt)
        if till == len(tasks) - 1 {
            taskAt = 0
            done   = 0
        } else {
            taskAt = till + 1
        }
        res = append(res, len(tasks) - till - 1)
    }
    return res
}

func countTasks1(tasks []int, shifts []int) []int {
    n := len(tasks)
    suffix := make([]int, n)
    suffix[n-1] = tasks[n-1]
    for i := n - 2; i >= 0; i-- {
        suffix[i] = suffix[i+1] + tasks[i]
    }
    m := len(shifts)
    res := make([]int, m)
    currTask, currDone := 0, 0
    for i := 0; i < m; i++ {
        shift := shifts[i]
        if suffix[currTask] - currDone <= shift {
            res[i], currTask, currDone = 0, 0, 0
            continue
        }
        for currTask < n && shift > 0 {
            currTaskRemain := tasks[currTask] - currDone
            finished := min(currTaskRemain, shift)
            shift -= finished
            currDone += finished
            if tasks[currTask] == currDone {
                currTask++
                currDone = 0
            }
        }
        res[i] = n - currTask
        if currTask == n {
            currTask, currDone = 0, 0
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: tasks = [1,4,4], shifts = [9,1,4]
    // Output: [0,2,1]
    // Explanation:
    // Shift 0: The tasks require 1 + 4 + 4 = 9 units of time, so all tasks are completed. There are 0 unfinished tasks.
    // Shift 1: Processing restarts from task 0. The shift has time 1, so task 0 is completed. There are 2 unfinished tasks.
    // Shift 2: Processing continues from task 1. The shift has time 4, so task 1 is completed. There is 1 unfinished task.
    fmt.Println(countTasks([]int{1,4,4}, []int{9,1,4})) // [0,2,1]
    // Example 2:
    // Input: tasks = [2,3,4], shifts = [20,4,5]
    // Output: [0,2,0]
    // Explanation:
    // Shift 0: The tasks require 2 + 3 + 4 = 9 units of time, so all tasks are completed. The remaining time in this shift is ignored. There are 0 unfinished tasks.
    // Shift 1: Processing restarts from task 0. The shift has time 4, so task 0 is completed and task 1 is partially completed. There are 2 unfinished tasks.
    // Shift 2: Processing continues from task 1. The remaining time needed is 1 + 4 = 5, so all tasks are completed. There are 0 unfinished tasks.
    fmt.Println(countTasks([]int{2,3,4}, []int{20,4,5})) // [0,2,0]
    // Example 3:
    // Input: tasks = [4,2], shifts = [3,6,1]
    // Output: [2,0,2]
    // Explanation:
    // Shift 0: The shift has time 3, so task 0 is partially completed with 1 unit of work remaining. There are 2 unfinished tasks.
    // Shift 1: Processing continues from task 0. The remaining time needed is 1 + 2 = 3, so all tasks are completed. There are 0 unfinished tasks.
    // Shift 2: Processing restarts from task 0. The shift has time 1, so task 0 is partially completed. There are 2 unfinished tasks.
    fmt.Println(countTasks([]int{4,2}, []int{3,6,1})) // [2,0,2]

    fmt.Println(countTasks([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [8 7 6 5 4 3 2 1 0]
    fmt.Println(countTasks([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [6 4 3 2 2 1 1 1 0]
    fmt.Println(countTasks([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // [9 9 9 8 8 7 6 4 0]
    fmt.Println(countTasks([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // [8 7 6 5 4 3 2 1 0]

    fmt.Println(countTasks1([]int{1,4,4}, []int{9,1,4})) // [0,2,1]
    fmt.Println(countTasks1([]int{2,3,4}, []int{20,4,5})) // [0,2,0]
    fmt.Println(countTasks1([]int{4,2}, []int{3,6,1})) // [2,0,2]
    fmt.Println(countTasks1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9})) // [8 7 6 5 4 3 2 1 0]
    fmt.Println(countTasks1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // [6 4 3 2 2 1 1 1 0]
    fmt.Println(countTasks1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // [9 9 9 8 8 7 6 4 0]
    fmt.Println(countTasks1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1})) // [8 7 6 5 
}