package main

// 252. Meeting Rooms
// Given an array of meeting time intervals where intervals[i] = [starti, endi], 
// determine if a person could attend all meetings.

// Example 1:
// Input: intervals = [[0,30],[5,10],[15,20]]
// Output: false

// Example 2:
// Input: intervals = [[7,10],[2,4]]
// Output: true

// Constraints:
//     0 <= intervals.length <= 10^4
//     intervals[i].length == 2
//     0 <= starti < endi <= 10^6

import "fmt"
import "sort"

func canAttendMeetings(intervals [][]int) bool {
    // 先从小到大排序
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })
    for i := 1; i < len(intervals); i++ {
        // 存在包含关系，有时间冲突
        if intervals[i - 1][1] > intervals[i][0] {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println(canAttendMeetings([][]int{ {0,30}, {5,10}, {15,20}})) // false
    fmt.Println(canAttendMeetings([][]int{ {7,10}, {2,4} })) // true
}