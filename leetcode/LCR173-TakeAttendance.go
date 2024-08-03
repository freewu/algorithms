package main

// LCR 173. 点名
// 某班级 n 位同学的学号为 0 ~ n-1。点名结果记录于升序数组 records。假定仅有一位同学缺席，请返回他的学号。

// 示例 1:
// 输入: records = [0,1,2,3,5]
// 输出: 4

// 示例 2:
// 输入: records = [0, 1, 2, 3, 4, 5, 6, 8]
// 输出: 7

// 提示：
//     1 <= records.length <= 10000

import "fmt"

func takeAttendance(records []int) int {
    for i, v := range records {
        if i != v {
            return i
        }
    }
    return len(records)
}

func main() {
    // 示例 1:
    // 输入: records = [0,1,2,3,5]
    // 输出: 4
    fmt.Println(takeAttendance([]int{0,1,2,3,5})) // 4
    // 示例 2:
    // 输入: records = [0, 1, 2, 3, 4, 5, 6, 8]
    // 输出: 7
    fmt.Println(takeAttendance([]int{0, 1, 2, 3, 4, 5, 6, 8})) // 7

    fmt.Println(takeAttendance([]int{0})) // 1
    fmt.Println(takeAttendance([]int{0,1})) // 2
}