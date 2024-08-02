package main

// LCR 172. 统计目标成绩的出现次数
// 某班级考试成绩按非严格递增顺序记录于整数数组 scores，请返回目标成绩 target 的出现次数。

// 示例 1：
// 输入: scores = [2, 2, 3, 4, 4, 4, 5, 6, 6, 8], target = 4
// 输出: 3

// 示例 2：
// 输入: scores = [1, 2, 3, 5, 7, 9], target = 6
// 输出: 0

// 提示：
//     0 <= scores.length <= 10^5
//     -10^9 <= scores[i] <= 10^9
//     scores 是一个非递减数组
//     -10^9 <= target <= 10^9

import "fmt"

func countTarget(scores []int, target int) int {
    res := 0
    for _, v := range scores {
        if v > target { 
            break 
        }
        if v == target {
            res++
        }
    }
    return res
}

// 二分法
func countTarget1(scores []int, target int) int {
    // 二分找左边界 在找右边界
    n := len(scores)
    left, right := 0, n
    start, end := 0, 0
    // 左闭右开
    for left < right {
        mid := left + (right - left) / 2 
        if scores[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
        start = left
    }
    left, right = 0, n
    for left < right {
        mid := left + (right - left) / 2 
        if scores[mid] < target + 1{
            left = mid + 1
        } else {
            right = mid
        }
        end = left
    }
    return end - start
}

func main() {
    // 示例 1：
    // 输入: scores = [2, 2, 3, 4, 4, 4, 5, 6, 6, 8], target = 4
    // 输出: 3
    fmt.Println(countTarget([]int{2, 2, 3, 4, 4, 4, 5, 6, 6, 8}, 4)) // 3
    // 示例 2：
    // 输入: scores = [1, 2, 3, 5, 7, 9], target = 6
    // 输出: 0
    fmt.Println(countTarget([]int{1, 2, 3, 5, 7, 9}, 6)) // 0

    fmt.Println(countTarget1([]int{2, 2, 3, 4, 4, 4, 5, 6, 6, 8}, 4)) // 3
    fmt.Println(countTarget1([]int{1, 2, 3, 5, 7, 9}, 6)) // 0
}