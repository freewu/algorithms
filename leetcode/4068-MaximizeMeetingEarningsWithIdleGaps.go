package main

// 4068. Maximize Meeting Earnings with Idle Gaps
// You are given a 2D integer array meetings, where meetings[i] = [starti, endi, revenuei] represents a meeting starting at time starti, ending at time endi, with revenue revenuei.

// All meetings use half-open intervals [start, end), so meetings that only touch at endpoints do not overlap.

// You may select any non-empty subset of meetings such that no two selected meetings overlap. 
// You earn the revenue of each selected meeting.

// Arrange the selected meetings in increasing order of their start times. 
// For each pair of adjacent meetings in this order, you also earn 1 unit of revenue per unit of idle time between them. 
// This idle time equals the later meeting's start time minus the earlier meeting's end time.

// No idle revenue is earned before the earliest selected meeting starts or after the latest selected meeting ends. 
// If only one meeting is selected, no idle revenue is earned.

// Return the maximum total earnings achievable.

// A subset of an array is a selection of elements of the array.

// Example 1:
// Input: meetings = [[2,5,4],[6,8,3]]
// Output: 8
// Explanation:
// Select both meetings. They do not overlap and earn 4 + 3 = 7 units of meeting revenue.
// The first meeting ends at time 5, and the second starts at time 6. This idle gap earns 6 - 5 = 1 additional unit.
// The maximum total earnings are 7 + 1 = 8.

// Example 2:
// Input: meetings = [[3,5,4],[4,7,8],[8,10,3]]
// Output: 12
// Explanation:
// Select the meetings at indices 1 and 2. They do not overlap and earn 8 + 3 = 11 units of meeting revenue.
// In chronological order, these meetings run from time 4 to 7 and from time 8 to 10. The idle gap earns 8 - 7 = 1 additional unit.
// The maximum total earnings are 11 + 1 = 12.

// Example 3:
// Input: meetings = [[1,2,2],[4,5,2],[7,9,3]]
// Output: 11
// Explanation:
// Select all three meetings. They do not overlap and earn 2 + 2 + 3 = 7 units of meeting revenue.
// The idle gap from time 2 to 4 earns 4 - 2 = 2 additional units.
// The idle gap from time 5 to 7 earns 7 - 5 = 2 additional units.
// The maximum total earnings are 7 + 2 + 2 = 11.

// Constraints
//     1 <= meetings.length <= 10^5
//     meetings[i] = [starti, endi, revenuei]
//     0 <= starti < endi <= 10^9
//     1 <= revenuei <= 10^9

import "fmt"
import "sort"
import "slices"

func maxEarnings(meetings [][]int) int64 {
    // 按结束时间从小到大排序
    slices.SortFunc(meetings, func(a, b []int) int { 
        return a[1] - b[1] 
    })
    res, end0 := 0, meetings[0][1]
    // preMax[i+1] = [0,i] 中的 f[j] - end[j] 的前缀最大值
    preMax := make([]int, len(meetings)+1)
    preMax[0] = -1 << 61
    for i, m := range meetings {
        start, end, revenue := m[0], m[1], m[2]
        f := revenue
        if start >= end0 { // 左边有会议
            j := sort.Search(i, func(j int) bool { return meetings[j][1] > start })
            // 为什么是 j 不是 j+1：上面算的是 > start，-1 后得到 <= start，但由于还要 +1，抵消了
            f += preMax[j] + start
        }
        res = max(res, f)
        preMax[i+1] = max(preMax[i], f-end)
    }
    return int64(res)
}

func maxEarnings1(meetings [][]int) int64 {
    type Meeting struct {
        start, end, revenue int64
    }
    ordered := make([]Meeting, len(meetings))
    for i, item := range meetings {
        ordered[i] = Meeting {
            start:   int64(item[0]),
            end:     int64(item[1]),
            revenue: int64(item[2]),
        }
    }
    sort.Slice(ordered, func(i, j int) bool {
        if ordered[i].end == ordered[j].end {
            return ordered[i].start < ordered[j].start
        }
        return ordered[i].end < ordered[j].end
    })
    prefix := make([]int64, len(ordered)+1)
    prefix[0] = -1 << 63
    res := int64(0)
    for i, curr := range ordered {
        compatibleCount := sort.Search(i, func(j int) bool {
            return ordered[j].end > curr.start
        })
        best := curr.revenue // Select only this meeting.
        if compatibleCount > 0 {
            candidate := curr.revenue + curr.start + prefix[compatibleCount]
            if candidate > best {
                best = candidate
            }
        }
        if best > res {
            res = best
        }
        prefix[i+1] = prefix[i]
        if value := best - curr.end; value > prefix[i+1] {
            prefix[i+1] = value
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: meetings = [[2,5,4],[6,8,3]]
    // Output: 8
    // Explanation:
    // Select both meetings. They do not overlap and earn 4 + 3 = 7 units of meeting revenue.
    // The first meeting ends at time 5, and the second starts at time 6. This idle gap earns 6 - 5 = 1 additional unit.
    // The maximum total earnings are 7 + 1 = 8.
    fmt.Println(maxEarnings([][]int{{2,5,4},{6,8,3}})) // 8 
    // Example 2:
    // Input: meetings = [[3,5,4],[4,7,8],[8,10,3]]
    // Output: 12
    // Explanation:
    // Select the meetings at indices 1 and 2. They do not overlap and earn 8 + 3 = 11 units of meeting revenue.
    // In chronological order, these meetings run from time 4 to 7 and from time 8 to 10. The idle gap earns 8 - 7 = 1 additional unit.
    // The maximum total earnings are 11 + 1 = 12.
    fmt.Println(maxEarnings([][]int{{3,5,4},{4,7,8},{8,10,3}})) // 12
    // Example 3:
    // Input: meetings = [[1,2,2],[4,5,2],[7,9,3]]
    // Output: 11
    // Explanation:
    // Select all three meetings. They do not overlap and earn 2 + 2 + 3 = 7 units of meeting revenue.
    // The idle gap from time 2 to 4 earns 4 - 2 = 2 additional units.
    // The idle gap from time 5 to 7 earns 7 - 5 = 2 additional units.
    // The maximum total earnings are 7 + 2 + 2 = 11.
    fmt.Println(maxEarnings([][]int{{1,2,2},{4,5,2},{7,9,3}})) // 11

    fmt.Println(maxEarnings1([][]int{{2,5,4},{6,8,3}})) // 8 
    fmt.Println(maxEarnings1([][]int{{3,5,4},{4,7,8},{8,10,3}})) // 12
    fmt.Println(maxEarnings1([][]int{{1,2,2},{4,5,2},{7,9,3}})) // 11
}