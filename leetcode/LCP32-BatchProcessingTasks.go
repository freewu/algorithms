package main

// LCP 32. 批量处理任务
// 某实验室计算机待处理任务以 [start,end,period] 格式记于二维数组 tasks，
// 表示完成该任务的时间范围为起始时间 start 至结束时间 end 之间，需要计算机投入 period 的时长，注意：
//     1. period 可为不连续时间
//     2. 首尾时间均包含在内

// 处于开机状态的计算机可同时处理任意多个任务，请返回电脑最少开机多久，可处理完所有任务。

// 示例 1：
// 输入：tasks = [[1,3,2],[2,5,3],[5,6,2]]
// 输出：4
// 解释： tasks[0] 选择时间点 2、3； tasks[1] 选择时间点 2、3、5； tasks[2] 选择时间点 5、6； 因此计算机仅需在时间点 2、3、5、6 四个时刻保持开机即可完成任务。

// 示例 2：
// 输入：tasks = [[2,3,1],[5,5,1],[5,6,2]]
// 输出：3
// 解释： tasks[0] 选择时间点 2 或 3； tasks[1] 选择时间点 5； tasks[2] 选择时间点 5、6； 因此计算机仅需在时间点 2、5、6 或 3、5、6 三个时刻保持开机即可完成任务。

// 提示：
//     2 <= tasks.length <= 10^5
//     tasks[i].length == 3
//     0 <= tasks[i][0] <= tasks[i][1] <= 10^9
//     1 <= tasks[i][2] <= tasks[i][1]-tasks[i][0] + 1

import "fmt"
import "sort"

// 超出时间限制 67 / 72
func processTasks(tasks [][]int) int {
    sort.Slice(tasks, func(i, j int) bool { 
        return tasks[i][1] < tasks[j][1] 
    })
    res, mp := 0, make([]bool, tasks[len(tasks)-1][1] + 1)
    for _, t := range tasks {
        start, end, d := t[0], t[1], t[2]
        for _, b := range mp[start : end + 1] { // 去掉运行中的时间点
            if b {
                d--
            }
        }
        for i := end; d > 0; i-- { // 剩余的 d 填充区间后缀
            if !mp[i] {
                mp[i] = true
                d--
                res++
            }
        }
    }
    return res
}

func processTasks1(tasks [][]int) int {
    sort.Slice(tasks, func(i, j int) bool { 
        return tasks[i][1] < tasks[j][1] 
    })
    type Tuple struct{ l, r, s int }
    stack := []Tuple{{-2, -2, 0}} // 闭区间左右端点，栈底到栈顶的区间长度的和
    for _, p := range tasks {
        start, end, d := p[0], p[1], p[2]
        i := sort.Search(len(stack), func(i int) bool { return stack[i].l >= start }) - 1
        d -= stack[len(stack) - 1].s - stack[i].s // 去掉运行中的时间点
        if start <= stack[i].r { // start 在区间 st[i] 内
            d -= stack[i].r - start + 1 // 去掉运行中的时间点
        }
        if d <= 0 { continue }
        for end - stack[len(stack) - 1].r <= d { // 剩余的 d 填充区间后缀
            top := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            d += top.r - top.l + 1 // 合并区间
        }
        stack = append(stack, Tuple{end - d + 1, end, stack[len(stack) - 1].s + d})
    }
    return stack[len(stack) - 1].s
}

func main() {
    // 示例 1：
    // 输入：tasks = [[1,3,2],[2,5,3],[5,6,2]]
    // 输出：4
    // 解释： tasks[0] 选择时间点 2、3； tasks[1] 选择时间点 2、3、5； tasks[2] 选择时间点 5、6； 因此计算机仅需在时间点 2、3、5、6 四个时刻保持开机即可完成任务。
    fmt.Println(processTasks([][]int{{1,3,2},{2,5,3},{5,6,2}})) // 4
    // 示例 2：
    // 输入：tasks = [[2,3,1],[5,5,1],[5,6,2]]
    // 输出：3
    // 解释： tasks[0] 选择时间点 2 或 3； tasks[1] 选择时间点 5； tasks[2] 选择时间点 5、6； 因此计算机仅需在时间点 2、5、6 或 3、5、6 三个时刻保持开机即可完成任务。
    fmt.Println(processTasks([][]int{{2,3,1},{5,5,1},{5,6,2}})) // 3

    fmt.Println(processTasks1([][]int{{1,3,2},{2,5,3},{5,6,2}})) // 4
    fmt.Println(processTasks1([][]int{{2,3,1},{5,5,1},{5,6,2}})) // 3
}