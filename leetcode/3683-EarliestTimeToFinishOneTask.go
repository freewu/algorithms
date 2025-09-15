package main

// 3683. Earliest Time to Finish One Task
// You are given a 2D integer array tasks where tasks[i] = [si, ti].

// Each [si, ti] in tasks represents a task with start time si that takes ti units of time to finish.

// Return the earliest time at which at least one task is finished.

// Example 1:
// Input: tasks = [[1,6],[2,3]]
// Output: 5
// Explanation:
// The first task starts at time t = 1 and finishes at time 1 + 6 = 7. 
// The second task finishes at time 2 + 3 = 5. 
// You can finish one task at time 5.

// Example 2:
// Input: tasks = [[100,100],[100,100],[100,100]]
// Output: 200
// Explanation:
// All three tasks finish at time 100 + 100 = 200.

// Constraints:
//     1 <= tasks.length <= 100
//     tasks[i] = [si, ti]
//     1 <= si, ti <= 100

import "fmt"

func earliestTime(tasks [][]int) int {
    res, sum := 0, 0
    if len(tasks) == 1 {
        for _, t := range tasks[0] {
            sum += t
        }
        return sum
    }
    for _, t := range tasks[0] {
        res += t
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < len(tasks); i++ {
        sum = 0
        for _, t := range tasks[i] {
            sum += t
        }
        res = min(res, sum)
    }
    return res
}
      
func main() {
    // Example 1:
    // Input: tasks = [[1,6],[2,3]]
    // Output: 5
    // Explanation:
    // The first task starts at time t = 1 and finishes at time 1 + 6 = 7. 
    // The second task finishes at time 2 + 3 = 5. 
    // You can finish one task at time 5.
    fmt.Println(earliestTime([][]int{{1,6},{2,3}})) // 5
    // Example 2:
    // Input: tasks = [[100,100],[100,100],[100,100]]
    // Output: 200
    // Explanation:
    // All three tasks finish at time 100 + 100 = 200.
    fmt.Println(earliestTime([][]int{{100,100},{100,100},{100,100}})) // 200
}