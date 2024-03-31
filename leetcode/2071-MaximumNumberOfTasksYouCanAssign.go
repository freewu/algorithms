package main

// 2071. Maximum Number of Tasks You Can Assign
// You have n tasks and m workers. Each task has a strength requirement stored in a 0-indexed integer array tasks,
// with the ith task requiring tasks[i] strength to complete. 
// The strength of each worker is stored in a 0-indexed integer array workers, 
// with the jth worker having workers[j] strength. 
// Each worker can only be assigned to a single task and must have a strength greater than or equal to the task's strength requirement (i.e., workers[j] >= tasks[i]).
// Additionally, you have pills magical pills that will increase a worker's strength by strength. 
// You can decide which workers receive the magical pills, however, you may only give each worker at most one magical pill.

// Given the 0-indexed integer arrays tasks and workers and the integers pills and strength, 
// return the maximum number of tasks that can be completed.

// Example 1:
// Input: tasks = [3,2,1], workers = [0,3,3], pills = 1, strength = 1
// Output: 3
// Explanation:
// We can assign the magical pill and tasks as follows:
// - Give the magical pill to worker 0.
// - Assign worker 0 to task 2 (0 + 1 >= 1)
// - Assign worker 1 to task 1 (3 >= 2)
// - Assign worker 2 to task 0 (3 >= 3)

// Example 2:
// Input: tasks = [5,4], workers = [0,0,0], pills = 1, strength = 5
// Output: 1
// Explanation:
// We can assign the magical pill and tasks as follows:
// - Give the magical pill to worker 0.
// - Assign worker 0 to task 0 (0 + 5 >= 5)

// Example 3:
// Input: tasks = [10,15,30], workers = [0,10,10,10,10], pills = 3, strength = 10
// Output: 2
// Explanation:
// We can assign the magical pills and tasks as follows:
// - Give the magical pill to worker 0 and worker 1.
// - Assign worker 0 to task 0 (0 + 10 >= 10)
// - Assign worker 1 to task 1 (10 + 10 >= 15)
// The last pill is not given because it will not make any worker strong enough for the last task.

// Constraints:
//     n == tasks.length
//     m == workers.length
//     1 <= n, m <= 5 * 10^4
//     0 <= pills <= m
//     0 <= tasks[i], workers[j], strength <= 10^9

import "fmt"
import "sort"

func maxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
    n,m := len(tasks), len(workers)
    min := func (a, b int) int { if b > a { return a; }; return b; }

    // Sorting the tasks and workers in increasing order
    sort.Ints(tasks)
    sort.Ints(workers)
    low,high, res := 0, min(m, n), -1
    for low <= high {
        mid := low + (high - low) / 2
        count := 0
        flag := true
        // Copying workers to a new slice
        st := make([]int, len(workers))
        copy(st, workers)
        // Checking if the mid smallest tasks can be assigned
        for i := mid - 1; i >= 0; i-- {
            // Case 1: Trying to assign to a worker without the pill
            it := st[len(st)-1]
            if tasks[i] <= it {
                // Case 1 satisfied!
                st = st[:len(st)-1]
            } else {
                // Case 2: Trying to assign to a worker with the pill
                it = sort.SearchInts(st, tasks[i]-strength)
                if it != len(st) {
                    // Case 2 satisfied!
                    count++
                    st = append(st[:it], st[it+1:]...)
                } else {
                    // Case 3: Impossible to assign mid tasks
                    flag = false
                    break
                }
            }
            // If at any moment, the number of pills required for mid tasks exceeds
            // the allotted number of pills, we stop the loop
            if count > pills {
                flag = false
                break
            }
        }
        if flag {
            res = mid
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return res
}

func maxTaskAssign1(tasks []int, workers []int, pills int, strength int) int {
    sort.Ints(tasks)
    sort.Ints(workers)
    min := func (a, b int) int { if b > a { return a; }; return b; }
    
    m, n := len(workers), len(tasks)
    l, r := 0, min(m, n)
    queue := make([]int, r)
    
    res := 0
    for l <= r { // 二分
        mid := (l+r) / 2
        // 每次选最弱的工作，和最强的工人
        if pick(tasks[:mid], workers[m-mid:], pills, strength, queue) {
            res = mid
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return res
}

type worker struct {
    strength int
    pill bool
}

func pick(tasks, workers []int, pills, strength int, queue []int) bool {
    r := 0
    head, tail := 0, 0
    for l := 0; l < len(workers); l++ {
        worker := workers[l]
        for r < len(tasks) && worker >= tasks[r] {
            queue[tail] = tasks[r]
            tail++
            r++
        }
        if tail > head && worker >= queue[head] {
            head++
            continue
        }
        worker += strength
        pills--
        if pills < 0 {
            return false
        }
        for r < len(tasks) && worker >= tasks[r] {
            queue[tail] = tasks[r]
            tail++
            r++
        }
        k := tail-1 
        for ; k >= head && worker < queue[k]; k-- {}
        if k < head {
            return false
        }
        for ; k+1 < tail; k++ {
            queue[k] = queue[k+1]
        }
        tail--
    }
    return true
}

func main() {
    // We can assign the magical pill and tasks as follows:
    // - Give the magical pill to worker 0.
    // - Assign worker 0 to task 2 (0 + 1 >= 1)
    // - Assign worker 1 to task 1 (3 >= 2)
    // - Assign worker 2 to task 0 (3 >= 3)
    fmt.Println(maxTaskAssign([]int{3,2,1},[]int{0,3,3},1,1)) // 3

    // We can assign the magical pill and tasks as follows:
    // - Give the magical pill to worker 0.
    // - Assign worker 0 to task 0 (0 + 5 >= 5)
    fmt.Println(maxTaskAssign([]int{5,4},[]int{0,0,0},1,5)) // 1

    // We can assign the magical pills and tasks as follows:
    // - Give the magical pill to worker 0 and worker 1.
    // - Assign worker 0 to task 0 (0 + 10 >= 10)
    // - Assign worker 1 to task 1 (10 + 10 >= 15)
    // The last pill is not given because it will not make any worker strong enough for the last task.
    fmt.Println(maxTaskAssign([]int{10,15,30},[]int{0,10,10,10,10},3,10)) // 2


    fmt.Println(maxTaskAssign1([]int{3,2,1},[]int{0,3,3},1,1)) // 3
    fmt.Println(maxTaskAssign1([]int{5,4},[]int{0,0,0},1,5)) // 1
    fmt.Println(maxTaskAssign1([]int{10,15,30},[]int{0,10,10,10,10},3,10)) // 2
}