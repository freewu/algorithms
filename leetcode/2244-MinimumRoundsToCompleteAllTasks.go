package main

// 2244. Minimum Rounds to Complete All Tasks
// You are given a 0-indexed integer array tasks, where tasks[i] represents the difficulty level of a task. 
// In each round, you can complete either 2 or 3 tasks of the same difficulty level.
// Return the minimum rounds required to complete all the tasks, or -1 if it is not possible to complete all the tasks.
 
// Example 1:
// Input: tasks = [2,2,3,3,2,4,4,4,4,4]
// Output: 4
// Explanation: To complete all the tasks, a possible plan is:
// - In the first round, you complete 3 tasks of difficulty level 2. 
// - In the second round, you complete 2 tasks of difficulty level 3. 
// - In the third round, you complete 3 tasks of difficulty level 4. 
// - In the fourth round, you complete 2 tasks of difficulty level 4.  
// It can be shown that all the tasks cannot be completed in fewer than 4 rounds, so the answer is 4.

// Example 2:
// Input: tasks = [2,3,3]
// Output: -1
// Explanation: There is only 1 task of difficulty level 2, but in each round, you can only complete either 2 or 3 tasks of the same difficulty level. Hence, you cannot complete all the tasks, and the answer is -1.

// Constraints:
//     1 <= tasks.length <= 10^5
//     1 <= tasks[i] <= 10^9

import "fmt"
import "math"

func minimumRounds(tasks []int) int {
    res, m := 0, map[int]int{}
    for _, v := range tasks { // 累加出任务不同难度任务的数量
        m[v]++
    }
    for _, v := range m {
        if v == 1 { // 一次只能完成 2 或 3 个任务 有 1完成不了直接返回 -1
            return -1 
        }
        res += int(math.Ceil(float64(v) / 3.0)) // 因为要最小轮数 向上取除3的值就勤够 余1的情况可以 3+1 分成 2+2 来处理所以向上取整数
    }
    return res
}

func minimumRounds1(tasks []int) int {
    res, mp := 0, make(map[int]int)
    for _, v := range tasks {
        mp[v]++
    }
    for _, v := range mp {
        if v == 1 {
            return -1
        }
        if v % 3 != 1 {
            res += (v + 1) / 3
        } else { // 余1的情况 3 + 1 分成 2 + 2 来处理
            res += (v + 1) / 3 + 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: tasks = [2,2,3,3,2,4,4,4,4,4]
    // Output: 4
    // Explanation: To complete all the tasks, a possible plan is:
    // - In the first round, you complete 3 tasks of difficulty level 2. 
    // - In the second round, you complete 2 tasks of difficulty level 3. 
    // - In the third round, you complete 3 tasks of difficulty level 4. 
    // - In the fourth round, you complete 2 tasks of difficulty level 4.  
    // It can be shown that all the tasks cannot be completed in fewer than 4 rounds, so the answer is 4.
    fmt.Println(minimumRounds([]int{2,2,3,3,2,4,4,4,4,4})) // 4
    // Example 2:
    // Input: tasks = [2,3,3]
    // Output: -1
    // Explanation: There is only 1 task of difficulty level 2, but in each round, you can only complete either 2 or 3 tasks of the same difficulty level. Hence, you cannot complete all the tasks, and the answer is -1.
    fmt.Println(minimumRounds([]int{2,3,3})) // -1
    fmt.Println(minimumRounds([]int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(minimumRounds([]int{9,8,7,6,5,4,3,2,1})) // -1

    fmt.Println(minimumRounds1([]int{2,2,3,3,2,4,4,4,4,4})) // 4
    fmt.Println(minimumRounds1([]int{2,3,3})) // -1
    fmt.Println(minimumRounds1([]int{1,2,3,4,5,6,7,8,9})) // -1
    fmt.Println(minimumRounds1([]int{9,8,7,6,5,4,3,2,1})) // -1
}