package main

// 3169. Count Days Without Meetings
// You are given a positive integer days representing the total number of days an employee is available for work (starting from day 1). 
// You are also given a 2D array meetings of size n where, meetings[i] = [start_i, end_i] represents the starting and ending days of meeting i (inclusive).

// Return the count of days when the employee is available for work but no meetings are scheduled.

// Note: The meetings may overlap.

// Example 1:
// Input: days = 10, meetings = [[5,7],[1,3],[9,10]]
// Output: 2
// Explanation:
// There is no meeting scheduled on the 4th and 8th days.

// Example 2:
// Input: days = 5, meetings = [[2,4],[1,3]]
// Output: 1
// Explanation:
// There is no meeting scheduled on the 5th day.

// Example 3:
// Input: days = 6, meetings = [[1,6]]
// Output: 0
// Explanation:
// Meetings are scheduled for all working days.

// Constraints:
//     1 <= days <= 10^9
//     1 <= meetings.length <= 10^5
//     meetings[i].length == 2
//     1 <= meetings[i][0] <= meetings[i][1] <= days

import "fmt"
import "sort"

func countDays(days int, meetings [][]int) int {
    sort.Slice(meetings, func(i,j int) bool {
        return meetings[i][0] < meetings[j][0]
    })
    maxEnd, between := meetings[0][1], 0 // rest between meetings
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 1; i < len(meetings); i++ {
        between += max(meetings[i][0] - maxEnd - 1, 0)
        maxEnd = max(maxEnd, meetings[i][1])
    }
    after,before := days - maxEnd , meetings[0][0] - 1
    return before + after + between
}

func countDays1(days int, meetings [][]int) int {
    // 先将所有区间按照左端点进行升序排序
    sort.Slice(meetings, func(i,j int) bool {
        return meetings[i][0] < meetings[j][0]
    })
    // 遍历所有区间，逐个求取并集
    unionSet := [][]int{meetings[0]} // 并集
    unionLen := meetings[0][1] - meetings[0][0] + 1 // 并集的长度
    for i := 1; i < len(meetings); i++ {
        last := unionSet[len(unionSet)-1] // 并集可能不连续，获取最后一个区间
        if meetings[i][0] <= last[1] {    // 当前区间与最后一个区间相交
            if meetings[i][1] > last[1] { // 当前区间并不完全被最后一个区间包含
                unionLen += meetings[i][1] - last[1] // 先算长度，防止 last[1] 被更新
                last[1] = meetings[i][1]
            }
        } else { // 当前区间与最后一个区间不相交
            unionSet = append(unionSet, meetings[i])
            unionLen += meetings[i][1] - meetings[i][0] + 1
        }
    }
    return days - unionLen
}

func countDays2(days int, meetings [][]int) int {
    sort.Slice(meetings, func(i,j int) bool {
        return meetings[i][0] < meetings[j][0]
    })
    res, last := 0, 0
    for _, meeting := range meetings {
        if meeting[0] > last {
            res += meeting[0] - last - 1
        }
        last = max(last, meeting[1])
        if meeting[1] >= days { break }
    }
    if days > last {
        res += days - last
    }
    return res
}

func main() {
    // Example 1:
    // Input: days = 10, meetings = [[5,7],[1,3],[9,10]]
    // Output: 2
    // Explanation:
    // There is no meeting scheduled on the 4th and 8th days.
    fmt.Println(countDays(10, [][]int{{5,7},{1,3},{9,10}})) // 2
    // Example 2:
    // Input: days = 5, meetings = [[2,4],[1,3]]
    // Output: 1
    // Explanation:
    // There is no meeting scheduled on the 5th day.
    fmt.Println(countDays(5, [][]int{{2,4},{1,3}})) // 1
    // Example 3:
    // Input: days = 6, meetings = [[1,6]]
    // Output: 0
    // Explanation:
    // Meetings are scheduled for all working days.
    fmt.Println(countDays(6, [][]int{{1,6}})) // 0

    fmt.Println(countDays1(10, [][]int{{5,7},{1,3},{9,10}})) // 2
    fmt.Println(countDays1(5, [][]int{{2,4},{1,3}})) // 1
    fmt.Println(countDays1(6, [][]int{{1,6}})) // 0

    fmt.Println(countDays2(10, [][]int{{5,7},{1,3},{9,10}})) // 2
    fmt.Println(countDays2(5, [][]int{{2,4},{1,3}})) // 1
    fmt.Println(countDays2(6, [][]int{{1,6}})) // 0
}