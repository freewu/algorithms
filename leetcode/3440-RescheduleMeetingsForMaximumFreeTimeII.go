package main

// 3440. Reschedule Meetings for Maximum Free Time II
// You are given an integer eventTime denoting the duration of an event. 
// You are also given two integer arrays startTime and endTime, each of length n.

// These represent the start and end times of n non-overlapping meetings that occur during the event between time t = 0 and time t = eventTime, 
// where the ith meeting occurs during the time [startTime[i], endTime[i]].

// You can reschedule at most one meeting by moving its start time while maintaining the same duration, 
// such that the meetings remain non-overlapping, 
// to maximize the longest continuous period of free time during the event.

// Return the maximum amount of free time possible after rearranging the meetings.

// Note that the meetings can not be rescheduled to a time outside the event and they should remain non-overlapping.

// Note: In this version, it is valid for the relative ordering of the meetings to change after rescheduling one meeting.

// Example 1:
// Input: eventTime = 5, startTime = [1,3], endTime = [2,5]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/12/22/example0_rescheduled.png" />
// Reschedule the meeting at [1, 2] to [2, 3], leaving no meetings during the time [0, 2].

// Example 2:
// Input: eventTime = 10, startTime = [0,7,9], endTime = [1,8,10]
// Output: 7
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/12/22/rescheduled_example0.png" />
// Reschedule the meeting at [0, 1] to [8, 9], leaving no meetings during the time [0, 7].

// Example 3:
// Input: eventTime = 10, startTime = [0,3,7,9], endTime = [1,4,8,10]
// Output: 6
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2025/01/28/image3.png" />
// Reschedule the meeting at [3, 4] to [8, 9], leaving no meetings during the time [1, 7].

// Example 4:
// Input: eventTime = 5, startTime = [0,1,2,3,4], endTime = [1,2,3,4,5]
// Output: 0
// Explanation:
// There is no time during the event not occupied by meetings.

// Constraints:
//     1 <= eventTime <= 10^9
//     n == startTime.length == endTime.length
//     2 <= n <= 10^5
//     0 <= startTime[i] < endTime[i] <= eventTime
//     endTime[i] <= startTime[i + 1] where i lies in the range [0, n - 2].

import "fmt"

// 至多调整一个会议
// 移动之后相对位置可以变动
func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
    n := len(startTime)
    // 获取 i 左边的间隔
    get := func(i int) int {
        if i == 0 {
            return startTime[0]
        } else if i == n {
            return eventTime - endTime[i-1]
        }
        return startTime[i] - endTime[i-1]
    }
    // 最大的三个间隔的下标,求的是左边的间隔
    a, b, c := 0, -1, -1
    for i := 1; i <= n; i++ {
        sz := get(i)
        if sz > get(a) {
            a, b, c = i, a, b
        } else if b < 0 || sz > get(b) {
            b, c = i, b
        } else if c < 0 || sz > get(c) {
            c = i
        }
    }
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        diff := endTime[i] - startTime[i]
        if (i != a && i+1 != a && get(a) >= diff) || (i != b && i+1 != b && get(b) >= diff) || (i != c && i+1 != c && get(c) >= diff) {
            res = max(res, get(i) + get(i + 1) + diff)
        } else {
            res = max(res, get(i) + get(i + 1))
        }
    }
    return res
}

func maxFreeTime1(eventTime int, startTime []int, endTime []int) int {
    res, mx, left, right, curr, n := 0, 0, 0, 0, 0, len(startTime)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < n; i++ {
        if i == 0 {
            left = startTime[i]
        } else {
            left = startTime[i] - endTime[i-1]
        }
        if i == n-1 {
            right = eventTime - endTime[n-1]
        } else {
            right = startTime[i+1] - endTime[i]
        }
        curr = endTime[i] - startTime[i]
        if mx >= curr {
            res = max(res, left + right + curr)
        } else {
            res = max(res, left + right)
        }
        mx = max(mx, left)
    }
    mx = 0
    for i := n - 1; i >= 0; i-- {
        if i == 0 {
            left = startTime[i]
        } else {
            left = startTime[i] - endTime[i-1]
        }
        if i == n - 1 {
            right = eventTime - endTime[n - 1]
        } else {
            right = startTime[i + 1] - endTime[i]
        }
        curr = endTime[i] - startTime[i]
        if mx >= curr {
            res = max(res, left + right + curr)
        } else {
            res = max(res, left + right)
        }
        mx = max(mx, right)
    }
    return res
}

func main() {
    // Example 1:
    // Input: eventTime = 5, startTime = [1,3], endTime = [2,5]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/12/22/example0_rescheduled.png" />
    // Reschedule the meeting at [1, 2] to [2, 3], leaving no meetings during the time [0, 2].
    fmt.Println(maxFreeTime(5, []int{1,3}, []int{2,5})) // 2
    // Example 2:
    // Input: eventTime = 10, startTime = [0,7,9], endTime = [1,8,10]
    // Output: 7
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/12/22/rescheduled_example0.png" />
    // Reschedule the meeting at [0, 1] to [8, 9], leaving no meetings during the time [0, 7].
    fmt.Println(maxFreeTime(10, []int{0,7,9}, []int{1,8,10})) // 7
    // Example 3:
    // Input: eventTime = 10, startTime = [0,3,7,9], endTime = [1,4,8,10]
    // Output: 6
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2025/01/28/image3.png" />
    // Reschedule the meeting at [3, 4] to [8, 9], leaving no meetings during the time [1, 7].
    fmt.Println(maxFreeTime(10, []int{0,3,7,9}, []int{1,4,8,10})) // 6
    // Example 4:
    // Input: eventTime = 5, startTime = [0,1,2,3,4], endTime = [1,2,3,4,5]
    // Output: 0
    // Explanation:
    // There is no time during the event not occupied by meetings.
    fmt.Println(maxFreeTime(5, []int{0,1,2,3,4}, []int{1,2,3,4,5})) // 0

    fmt.Println(maxFreeTime(5, []int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 6
    fmt.Println(maxFreeTime(5, []int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 8

    fmt.Println(maxFreeTime1(5, []int{1,3}, []int{2,5})) // 2
    fmt.Println(maxFreeTime1(10, []int{0,7,9}, []int{1,8,10})) // 7
    fmt.Println(maxFreeTime1(10, []int{0,3,7,9}, []int{1,4,8,10})) // 6
    fmt.Println(maxFreeTime1(5, []int{0,1,2,3,4}, []int{1,2,3,4,5})) // 0
    fmt.Println(maxFreeTime1(5, []int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 6
    fmt.Println(maxFreeTime1(5, []int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 8
}