package main

// 2589. Minimum Time to Complete All Tasks
// There is a computer that can run an unlimited number of tasks at the same time. 
// You are given a 2D integer array tasks where tasks[i] = [starti, endi, durationi] indicates 
// that the ith task should run for a total of durationi seconds (not necessarily continuous) within the inclusive time range [starti, endi].

// You may turn on the computer only when it needs to run a task. You can also turn it off if it is idle.
// Return the minimum time during which the computer should be turned on to complete all tasks.

// Example 1:
// Input: tasks = [[2,3,1],[4,5,1],[1,5,2]]
// Output: 2
// Explanation: 
// - The first task can be run in the inclusive time range [2, 2].
// - The second task can be run in the inclusive time range [5, 5].
// - The third task can be run in the two inclusive time ranges [2, 2] and [5, 5].
// The computer will be on for a total of 2 seconds.

// Example 2:
// Input: tasks = [[1,3,2],[2,5,3],[5,6,2]]
// Output: 4
// Explanation: 
// - The first task can be run in the inclusive time range [2, 3].
// - The second task can be run in the inclusive time ranges [2, 3] and [5, 5].
// - The third task can be run in the two inclusive time range [5, 6].
// The computer will be on for a total of 4 seconds.

// Constraints:
//     1 <= tasks.length <= 2000
//     tasks[i].length == 3
//     1 <= starti, endi <= 2000
//     1 <= durationi <= endi - starti + 1

import "fmt"
import "sort"
import "math"
import "slices"

func findMinimumTime(tasks [][]int) int {
    sort.SliceStable(tasks, func(i, j int) bool {
        return tasks[i][1] < tasks[j][1]
    })
    min := func (x, y int) int { if x < y { return x; }; return y; }
    sum := func (nums []int) int {
        res := 0
        for _, v := range nums {
            res += v
        }
        return res
    }
    min_start := math.MaxInt
    for _,task := range(tasks){
        min_start = min(min_start,task[0])
    }
    max_end := tasks[len(tasks)-1][1]
    time := make([]int, max_end-min_start + 1)
    for _,task := range(tasks){
        start,end, duration := task[0] - min_start, task[1] - min_start, task[2]
        used := sum(time[start: end + 1])
        duration -= used
        for duration > 0 {
            if time[end] == 0 {
                time[end] = 1
                duration -= 1
            }
            end -= 1
        }
    }
    return sum(time)
}

func findMinimumTime1(tasks [][]int) int {
    slices.SortFunc(tasks, func(a, b []int) int { return a[1] - b[1] })
    type tuple struct{ l, r, s int } // 栈中保存闭区间左右端点，栈底到栈顶的区间长度的和
    stack := []tuple{{-2, -2, 0}} // 哨兵
    for _, p := range tasks {
        start, end, d := p[0], p[1], p[2]
        i := sort.Search(len(stack), func(i int) bool { return stack[i].l >= start }) - 1
        d -= stack[len(stack)-1].s - stack[i].s // 去掉运行中的时间点
        if start <= stack[i].r {          // start 在区间 st[i] 内
            d -= stack[i].r - start + 1 // 去掉运行中的时间点
        }
        if d <= 0 {
            continue
        }
        for end - stack[len(stack)-1].r <= d { // 剩余的 d 填充区间后缀
            top := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            d += top.r - top.l + 1 // 合并区间
        }
        stack = append(stack, tuple{end - d + 1, end, stack[len(stack) - 1].s + d})
    }
    return stack[len(stack)-1].s
}

func main() {
    // Example 1:
    // Input: tasks = [[2,3,1],[4,5,1],[1,5,2]]
    // Output: 2
    // Explanation: 
    // - The first task can be run in the inclusive time range [2, 2].
    // - The second task can be run in the inclusive time range [5, 5].
    // - The third task can be run in the two inclusive time ranges [2, 2] and [5, 5].
    // The computer will be on for a total of 2 seconds.
    fmt.Println(findMinimumTime([][]int{{2,3,1},{4,5,1},{1,5,2}})) // 2
    // Example 2:
    // Input: tasks = [[1,3,2],[2,5,3],[5,6,2]]
    // Output: 4
    // Explanation: 
    // - The first task can be run in the inclusive time range [2, 3].
    // - The second task can be run in the inclusive time ranges [2, 3] and [5, 5].
    // - The third task can be run in the two inclusive time range [5, 6].
    // The computer will be on for a total of 4 seconds.
    fmt.Println(findMinimumTime([][]int{{1,3,2},{2,5,3},{5,6,2}})) // 4

    fmt.Println(findMinimumTime1([][]int{{2,3,1},{4,5,1},{1,5,2}})) // 2
    fmt.Println(findMinimumTime1([][]int{{1,3,2},{2,5,3},{5,6,2}})) // 4
}