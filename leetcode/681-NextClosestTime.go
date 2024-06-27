package main

// 681. Next Closest Time
// Given a time represented in the format "HH:MM", form the next closest time by reusing the current digits. 
// There is no limit on how many times a digit can be reused.

// You may assume the given input string is always valid. 
// For example, "01:34", "12:09" are all valid. "1:34", "12:9" are all invalid.

// Example 1:
// Input: time = "19:34"
// Output: "19:39"
// Explanation: The next closest time choosing from digits 1, 9, 3, 4, is 19:39, which occurs 5 minutes later.
// It is not 19:33, because this occurs 23 hours and 59 minutes later.

// Example 2:
// Input: time = "23:59"
// Output: "22:22"
// Explanation: The next closest time choosing from digits 2, 3, 5, 9, is 22:22.
// It may be assumed that the returned time is next day's time since it is smaller than the input time numerically.

// Constraints:
//     time.length == 5
//     time is a valid time in the form "HH:MM".
//     0 <= HH < 24
//     0 <= MM < 60

import "fmt"
import "sort"

func nextClosestTime(time string) string {
    record := make(map[string]bool)
    for _, ch := range time { // 时间数字拆解出来
        if ch != ':' {
            record[string(ch)] = true
        }
    }
    res, hours, mins := []string{},[]string{},[]string{}
    for key1, _ := range record {
        for key2, _ := range record {
            temp := key1 + key2
            if temp < "24" { // 枚举出所有的小时
                hours = append(hours, temp)
            }
            if temp < "60" { // 枚举出所有的分钟
                mins = append(mins, temp)
            }
        }
    }
    // 合并所有的枚举时间
    for _, str1 := range hours { 
        for _, str2 := range mins { 
            res = append(res, str1+":"+str2)
        }
    }
    sort.Strings(res) // 排序
    indexOf := func(input []string, target string) int {
        for i, str := range input {
            if str == target {
                return i
            }
        }
        return -1
    }
    index := indexOf(res, time) // 找到输入时间在排序后数组的位置
    if index == len(res) - 1 { // 如果是最后一个，取第一个
        return res[0]
    } else { // 否则取下一个
        return res[index+1]
    }
}

func main() {
    // Example 1:
    // Input: time = "19:34"
    // Output: "19:39"
    // Explanation: The next closest time choosing from digits 1, 9, 3, 4, is 19:39, which occurs 5 minutes later.
    // It is not 19:33, because this occurs 23 hours and 59 minutes later.
    fmt.Println(nextClosestTime("19:34")) // "19:39"
    // Example 2:
    // Input: time = "23:59"
    // Output: "22:22"
    // Explanation: The next closest time choosing from digits 2, 3, 5, 9, is 22:22.
    // It may be assumed that the returned time is next day's time since it is smaller than the input time numerically.
    fmt.Println(nextClosestTime("23:59")) // 22:22"
}