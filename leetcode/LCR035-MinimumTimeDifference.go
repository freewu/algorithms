package main

// LCR 035. 最小时间差
// 给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。

// 示例 1：
// 输入：timePoints = ["23:59","00:00"]
// 输出：1

// 示例 2：
// 输入：timePoints = ["00:00","23:59","00:00"]
// 输出：0

// 提示：
//     2 <= timePoints.length <= 2 * 10^4
//     timePoints[i] 格式为 "HH:MM"

import "fmt"
import "sort"
import "strconv"

func findMinDifference(timePoints []string) int {
    res, minutes := 1 << 32 -1, []int{}
    for _, tp := range timePoints { // 把时间转换成秒
        minute, _ := strconv.Atoi(tp[:2])
        second, _ := strconv.Atoi(tp[3:])
        minutes = append(minutes, minute * 60 + second)
    }
    sort.Ints(minutes) // 从小到大排序
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 1; i < len(minutes); i++ { // 找到最小的时间差
        res = min(res, minutes[i] - minutes[i-1])
    }
    res = min(res, 24 * 60 - (minutes[len(minutes)-1] - minutes[0])) // 处理 23:59 00:00 这种情况 开头结束相差的情况
    return res
}

func main() {
    // Example 1:
    // Input: timePoints = ["23:59","00:00"]
    // Output: 1
    fmt.Println(findMinDifference([]string{"23:59","00:00"})) // 1
    // Example 2:
    // Input: timePoints = ["00:00","23:59","00:00"]
    // Output: 0
    fmt.Println(findMinDifference([]string{"00:00","23:59","00:00"})) // 0
}