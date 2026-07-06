package main

// 1288. Remove Covered Intervals
// Given an array intervals where intervals[i] = [li, ri] represent the interval [li, ri), 
// remove all intervals that are covered by another interval in the list.

// The interval [a, b) is covered by the interval [c, d) if and only if c <= a and b <= d.
// Return the number of remaining intervals.

// Example 1:
// Input: intervals = [[1,4],[3,6],[2,8]]
// Output: 2
// Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.

// Example 2:
// Input: intervals = [[1,4],[2,3]]
// Output: 1
 
// Constraints:
//     1 <= intervals.length <= 1000
//     intervals[i].length == 2
//     0 <= li < ri <= 10^5
//     All the given intervals are unique.

import "fmt"
import "sort"
import "cmp"
import "slices"

func removeCoveredIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool { // sort by increasing 'start' then by decreasing 'end'
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] <= intervals[j][0]
    })
    i, count := 0, 0
    for j := 1; j < len(intervals); j++ {
        if intervals[j][1] > intervals[i][1] {
            i = j
            continue
        }
        count++// remove jth interval 
    }
    return len(intervals) - count
}

func removeCoveredIntervals1(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })
    res, left, right := 0, intervals[0][0], intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        if left <= intervals[i][0] && right >= intervals[i][1] { // 被覆盖
            res++
        } else if right >= intervals[i][0] && right <= intervals[i][1] { // 有交集
            right = intervals[i][1]
        } else { // 无交集
            left = intervals[i][0]
            right = intervals[i][1]
        }
    }
    return len(intervals) - res
}

func removeCoveredIntervals2(intervals [][]int) int {
    // 按区间左端点从小到大排序
    // 区间左端点相同时，按区间右端点从大到小排序，这样会先遍历大区间，再遍历被大区间覆盖的小区间
    slices.SortFunc(intervals, func(a, b []int) int {
        return cmp.Or(a[0] - b[0], b[1] - a[1])
    })
    res, maxRight := 0, 0 // 已遍历区间中的最大右端点
    for _, p := range intervals {
        // 由于区间左端点是从小到大排序的，遍历过的区间的左端点都 <= 当前区间的左端点
        // 如果当前区间右端点 <= maxRight，说明当前区间被另一个区间覆盖，否则没被覆盖
        if p[1] > maxRight {
            maxRight = p[1]
            res++ // 当前区间没被覆盖
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: intervals = [[1,4],[3,6],[2,8]]
    // Output: 2
    // Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.
    fmt.Println(removeCoveredIntervals([][]int{{1,4},{3,6},{2,8}})) // 2
    // Example 2:
    // Input: intervals = [[1,4],[2,3]]
    // Output: 1
    fmt.Println(removeCoveredIntervals([][]int{{1,4},{2,3}})) // 1

    fmt.Println(removeCoveredIntervals1([][]int{{1,4},{3,6},{2,8}})) // 2
    fmt.Println(removeCoveredIntervals1([][]int{{1,4},{2,3}})) // 1

    fmt.Println(removeCoveredIntervals2([][]int{{1,4},{3,6},{2,8}})) // 2
    fmt.Println(removeCoveredIntervals2([][]int{{1,4},{2,3}})) // 1
}