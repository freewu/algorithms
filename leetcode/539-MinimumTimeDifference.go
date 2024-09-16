package main

// 539. Minimum Time Difference
// Given a list of 24-hour clock time points in "HH:MM" format, 
// return the minimum minutes difference between any two time-points in the list.
 
// Example 1:
// Input: timePoints = ["23:59","00:00"]
// Output: 1

// Example 2:
// Input: timePoints = ["00:00","23:59","00:00"]
// Output: 0
 
// Constraints:
//     2 <= timePoints.length <= 2 * 10^4
//     timePoints[i] is in the format "HH:MM".

import "fmt"
import "sort"
import "strconv"
import "strings"

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

func findMinDifference1(timePoints []string) int {
    minutes := make([]bool, 24* 60)
    mx, mn := 0, 24 * 60
    for _, timePoint := range(timePoints) {
        parts := strings.Split(timePoint, ":")
        h, _ := strconv.Atoi(parts[0])
        m, _ := strconv.Atoi(parts[1])
        temp := h * 60 + m
        if temp > mx { mx = temp }
        if temp < mn { mn = temp  }
        if minutes[temp] {  return 0 }
        minutes[temp] = true
    }
    res, prev := mn - mx + 1440, mn
    for i := mn + 1; i <= mx; i++ {
        if minutes[i]{
            diff := i - prev
            if diff < res {
                res = diff
            }
            prev = i
        }
    }
    return res
}

func findMinDifference2(timePoints []string) int {
    res, date :=  60*24, []int{}
    convert := func(timePoint string) int { // 转换成分钟数
        vals := strings.Split(timePoint, ":")
        h := int(vals[0][0] - '0') * 10 + int(vals[0][1] - '0')
        m := int(vals[1][0] - '0') * 10 + int(vals[1][1] - '0')
        return h * 60 + m
    }
    for i:=range timePoints{
        date = append(date, convert(timePoints[i]))
    }
    sort.Ints(date)
    for i := 1; i <len(date); i++ {
        res = min(res, date[i] - date[i-1])
    }
    res = min(res, 60 * 24 - date[len(date)-1] + date[0]) // 处理跨天的情况
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

    fmt.Println(findMinDifference1([]string{"23:59","00:00"})) // 1
    fmt.Println(findMinDifference1([]string{"00:00","23:59","00:00"})) // 0

    fmt.Println(findMinDifference2([]string{"23:59","00:00"})) // 1
    fmt.Println(findMinDifference2([]string{"00:00","23:59","00:00"})) // 0
}