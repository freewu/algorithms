package main

// 3439. Reschedule Meetings for Maximum Free Time I
// You are given an integer eventTime denoting the duration of an event, 
// where the event occurs from time t = 0 to time t = eventTime.

// You are also given two integer arrays startTime and endTime, each of length n. 
// These represent the start and end time of n non-overlapping meetings, 
// where the ith meeting occurs during the time [startTime[i], endTime[i]].

// You can reschedule at most k meetings by moving their start time while maintaining the same duration, 
// to maximize the longest continuous period of free time during the event.

// The relative order of all the meetings should stay the same and they should remain non-overlapping.

// Return the maximum amount of free time possible after rearranging the meetings.

// Note that the meetings can not be rescheduled to a time outside the event.

// Example 1:
// Input: eventTime = 5, k = 1, startTime = [1,3], endTime = [2,5]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/12/21/example0_rescheduled.png" />
// Reschedule the meeting at [1, 2] to [2, 3], leaving no meetings during the time [0, 2].

// Example 2:
// Input: eventTime = 10, k = 1, startTime = [0,2,9], endTime = [1,4,10]
// Output: 6
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/12/21/example1_rescheduled.png" />
// Reschedule the meeting at [2, 4] to [1, 3], leaving no meetings during the time [3, 9].

// Example 3:
// Input: eventTime = 5, k = 2, startTime = [0,1,2,3,4], endTime = [1,2,3,4,5]
// Output: 0
// Explanation:
// There is no time during the event not occupied by meetings.

// Constraints:
//     1 <= eventTime <= 10^9
//     n == startTime.length == endTime.length
//     2 <= n <= 10^5
//     1 <= k <= n
//     0 <= startTime[i] < endTime[i] <= eventTime
//     endTime[i] <= startTime[i + 1] where i lies in the range [0, n - 2].

import "fmt"

// 移动之后相对位置不变
// 会议不重复
func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
    pre, n := 0, len(startTime)
    interval := []int{}
    for i := 0; i < n; i++ {
        interval = append(interval, startTime[i] - pre)
        pre = endTime[i]
    }
    if eventTime - pre > 0 {
        interval = append(interval, eventTime - pre)
    }
    if interval[0] == 0 {
        interval = interval[1:]
    }
    k++
    res, wid := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(interval); i++ {
        if i < k {
            wid += interval[i]
        } else {
            wid = wid + interval[i] - interval[i - k]
        }
        res = max(res, wid)
    }
    return res
}

func maxFreeTime1(eventTime int, k int, startTime []int, endTime []int) int {
    prev := startTime[0]
    for i := 1; i <= k; i++ {
        if i == len(startTime) {
            return prev + (eventTime - endTime[i-1])
        }
        prev += (startTime[i] - endTime[i-1])
    }
    res, curr, le := prev, 0, startTime[0]
    for i := k + 1; i < len(startTime); i++ {
        curr = prev + (startTime[i] - endTime[i - 1]) - le
        if curr > res {
            res = curr
        }
        prev, le = curr, startTime[i - k] - endTime[i - k - 1]
    }
    curr = prev + (eventTime - endTime[len(endTime) - 1]) - le
    if curr > res {
        res = curr
    }
    return res
}

func main() {
    // Example 1:
    // Input: eventTime = 5, k = 1, startTime = [1,3], endTime = [2,5]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/12/21/example0_rescheduled.png" />
    // Reschedule the meeting at [1, 2] to [2, 3], leaving no meetings during the time [0, 2].
    fmt.Println(maxFreeTime(5, 1, []int{1,3}, []int{2,5})) // 2
    // Example 2:
    // Input: eventTime = 10, k = 1, startTime = [0,2,9], endTime = [1,4,10]
    // Output: 6
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/12/21/example1_rescheduled.png" />
    // Reschedule the meeting at [2, 4] to [1, 3], leaving no meetings during the time [3, 9].
    fmt.Println(maxFreeTime(10, 1, []int{0,2,9}, []int{1,4,10})) // 6
    // Example 3:
    // Input: eventTime = 5, k = 2, startTime = [0,1,2,3,4], endTime = [1,2,3,4,5]
    // Output: 0
    // Explanation:
    // There is no time during the event not occupied by meetings.
    fmt.Println(maxFreeTime(5, 2, []int{0,1,2,3,4}, []int{1,2,3,4,5})) // 0

    fmt.Println(maxFreeTime(5, 2, []int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 16
    fmt.Println(maxFreeTime(5, 2, []int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 21

    fmt.Println(maxFreeTime1(5, 1, []int{1,3}, []int{2,5})) // 2
    fmt.Println(maxFreeTime1(10, 1, []int{0,2,9}, []int{1,4,10})) // 6
    fmt.Println(maxFreeTime1(5, 2, []int{0,1,2,3,4}, []int{1,2,3,4,5})) // 0
    fmt.Println(maxFreeTime1(5, 2, []int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1})) // 16
    fmt.Println(maxFreeTime1(5, 2, []int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9})) // 21
}